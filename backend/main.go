package main

import (
	"fmt"
	"log"
	"os"
	controllers "spring-assessment-backend/controller"

	"github.com/go-pg/pg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgAddr := os.Getenv("POSTGRES_HOST")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDb := os.Getenv("POSTGRES_DATABASE")
	db := pg.Connect(&pg.Options{
		Addr:     pgAddr,
		User:     pgUser,
		Password: pgPass,
		Database: pgDb,
	})

	fmt.Printf("postgres://%s:%s@%s/%s?sslmode=disable", pgUser, pgPass, pgAddr, pgDb)

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pgUser, pgPass, pgAddr, pgDb))

	if err != nil {
		panic(err)
	}
	m.Up()

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	controller := controllers.NewProductController(db)

	app.Get("/products", controller.ListProducts)
	app.Get("/products/search", controller.SearchProducts)
	app.Post("/products/generate", controller.InsertProducts)

	log.Fatal(app.Listen(":3001"))
}
