# Marketplace Enabler API
This project aims to provide a powerful API that enables seamless integration with various online marketplaces

## Description
Marketplace Enabler API provides a unified interface to interact with different online marketplaces, streamlining the integration process for various e-commerce platforms

## System Requirement
- Go 1.18 or higher
- MySQL 8.0 or higher

## Setup and Installation
1. Clone the repository and navigate to the root folder

2. Create a `.env` file and set your MySQL DB configurations. Refer to `./infras/mysql.go` for the required parameters

3. Set up the database tables by running the SQL scripts located in the `./migrations` folder

4. Generate the necessary wire code:
   ```
   go generate ./...
   ```

5. Build the application:
   ```
   go build
   ```

**Note:** Make sure to replace the placeholders in the `.env` file with your actual MySQL configurations

## Run and Test
1. Ensure that you have completed the setup and installation steps mentioned above

2. To start the application, run the following command in the project root folder:
   ```
   go run .
   ```

3. The API will be accessible at [http://localhost:8080](http://localhost:8080)

4. Once the application is up and running, you can interact with the API using the following endpoints:

### Get Orders by Market

- **Endpoint:** GET /v1/orders/{market}
- **Description:** Retrieve orders by market with pagination
- **Path Parameter:** `market`
- **Query Parameters:** `page` (default: 1), `size` (default: 10)
- **Example:** GET http://localhost:8080/v1/orders/shopee?page=1&size=5

### Count Total Orders by Market

- **Endpoint:** GET /v1/orders/{market}/count
- **Description:** Count the total number of orders by market
- **Path Parameter:** `market`
- **Example:** GET http://localhost:8080/v1/orders/blibli/count

### Get Brands by Market

- **Endpoint:** GET /v1/orders/{market}/brand
- **Description:** Retrieve brands by market
- **Path Parameter:** `market`
- **Example:** GET http://localhost:8080/v1/orders/blibli/brand

### Get Stores by Market

- **Endpoint:** GET /v1/orders/{market}/store
- **Description:** Retrieve stores by market
- **Path Parameter:** `market`
- **Example:** GET http://localhost:8080/v1/orders/blibli/store

### Download Orders by Market

- **Endpoint:** GET /v1/orders/{market}/download
- **Description:** Download orders data by market (max 6000 rows)
- **Path Parameter:** `market`
- **Query Parameters:** `brand` (required), `store` (required), `dateFrom` (RFC3339 datetime format, optional), `dateTo` (RFC3339 datetime format, optional)
- **Example:** GET http://localhost:8080/v1/orders/shopee/download?brand=Abadi%20Logam&store=Abadi&dateFrom=2023-07-15T00:00:00Z&dateTo=2023-07-31T22:59:59Z


5. Replace `{market}` in the URLs with the specific market value (e.g., "shopee" or "blibli").

6. You can use tools like `curl`, Postman, or any other API testing tool to make requests to the endpoints and observe the API's functionality.

**Note:** Keep the application running while testing the endpoints.

## Swagger Documentation
To access the API documentation using Swagger, follow these steps:

1. Make sure the server is running locally.
2. Open your web browser and go to [http://localhost:8080/swagger/doc.json/](http://localhost:8080/swagger/doc.json/).
3. You'll see the Swagger UI interface with a list of endpoints, request parameters, and example requests/responses.
4. You can interact with the API directly from the Swagger interface.

## Contributing
Contributions are welcome! If you want to contribute, please follow these steps:
 - Create an issue detailing the feature or bug fix you intend to work on.
 - Fork the repository and create a new branch for your feature.
 - Implement your changes.
 - Create a pull request and reference the issue.
