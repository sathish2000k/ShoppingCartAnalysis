# ShoppingCartAnalysis

ShoppingCartAnalysis

ShoppingCartAnalysis is a backend service built in Go to analyze shopping cart data using LangChain (Gemini LLM), PostgreSQL, and REST APIs. It provides insights on sales, customers, and products with a scalable, production-ready architecture.

Prerequisites

Go 1.24+
PostgreSQL (locally or via Docker)
Docker
Git

Database

We use the Bike Store Sample Database from Kaggle:

https://www.kaggle.com/datasets/dillonmyrick/bike-store-sample-database

It includes CSV files for:

customers
orders
order_items
products
brands
categories
stores
staffs

These files are used to populate your PostgreSQL database for analytics in the ShoppingCartAnalysis app.

Setup & Installation
Clone the Repository
git clone https://github.com/yourusername/ShoppingCartAnalysis.git
cd ShoppingCartAnalysis

Environment Variables

Create a .env file in the root directory:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=shopping_cart
LLM_API_KEY=your_gemini_key

Running Locally
Install Dependencies
go mod tidy

Run the API
go run cmd/main.go


The API will be available at:

http://localhost:8080

Docker
Build the Image
docker build -t shoppingcart-backend:latest .

Run the Container
docker run -p 8080:8080 --env-file .env shoppingcart-backend:latest

API Reference
POST /query

Description: Ask analytical queries about customers, sales, and products.

Request Body:

{
  "query": "Top 5 products by sales in February"
}


Response:

{
  "result": "Product A: 150 sales, Product B: 120 sales..."
}
