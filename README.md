# Spring financial assessment

## Setup

- Run docker-compose up
- Access the frontend at http://localhost:3000

## Seeding
Below is an example CURL request to seed data to the database:
```
curl --location 'localhost:3001/products/generate' \
--header 'Content-Type: application/json' \
--data '[{
        "name": "product-1",
        "description": "test description",
        "category": "test category",
        "brand": "test brand",
        "stock_quantity": 100,
        "sku": "TEST-SKU"
    }]'
```
