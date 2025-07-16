package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"spring-assessment-backend/db/pg/model"
	"spring-assessment-backend/db/pg/repository"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const DEFAULT_PRODUCT_LIST_COUNT = 100
const PRODUCT_GENERATE_COUNT_QUERY = "count"
const PRODUCT_GENERATE_COUNT_QUERY_DEFAULT_VALUE = "1000"
const PRODUCT_SEARCH_QUERY = "query"

type ProductController interface {
	ListProducts(ctx *fiber.Ctx) error
	InsertProducts(ctx *fiber.Ctx) error
	SearchProducts(ctx *fiber.Ctx) error
}

type productController struct {
	repo repository.ProductRepository
}

func NewProductController(
	db *pg.DB,
) ProductController {
	repo := repository.NewProductRepository(db)
	return &productController{
		repo,
	}
}

func (c *productController) ListProducts(ctx *fiber.Ctx) error {
	products, err := c.repo.ListProducts(ctx.Context(), uuid.Nil, DEFAULT_PRODUCT_LIST_COUNT)

	if err != nil {
		fmt.Println(err.Error())
		ctx.SendString("An unknown error has occured")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(products)
}

func (c *productController) InsertProducts(ctx *fiber.Ctx) error {

	var input []model.Product
	err := json.NewDecoder(bytes.NewReader(ctx.Body())).Decode(&input)
	switch {
	case err == io.EOF:
		// No body
		countStr := ctx.Query(PRODUCT_GENERATE_COUNT_QUERY, PRODUCT_GENERATE_COUNT_QUERY_DEFAULT_VALUE)
		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println(err.Error())
			ctx.SendString("An unknown error has occured")
			ctx.SendStatus(fiber.StatusInternalServerError)
		}

		err = c.repo.CreateProducts(ctx.Context(), count)
		if err != nil {
			fmt.Println(err.Error())
			ctx.SendString("An unknown error has occured")
			ctx.SendStatus(fiber.StatusInternalServerError)
		}
	case err != nil:
		fmt.Println(err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError)
	default:
		err := c.repo.CreateProductsWithBody(ctx.Context(), input)
		if err != nil {
			fmt.Println(err.Error())
			ctx.SendString("An unknown error has occured")
			ctx.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *productController) SearchProducts(ctx *fiber.Ctx) error {
	query := ctx.Query(PRODUCT_SEARCH_QUERY)
	products, err := c.repo.SearchProducts(ctx.Context(), query)

	if query == "" {
		ctx.SendString("Query cannot be empty")
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err != nil {
		fmt.Println(err.Error())
		ctx.SendString("An unknown error has occured")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(products)
}
