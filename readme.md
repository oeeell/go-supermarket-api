# 🛍️ **Supermarket REST API** - A Personal Project for Mastering Golang

A scalable REST API for supermarket transactions, built with **Golang (Gin), SQLite, and GORM**.

## 🚀 Features
✅ Product Management (Add, Update, Delete, List)  
✅ Transaction Processing (Checkout, Stock Management)  
✅ Sales Reports (Daily, Monthly, Best-Selling Products)  
✅ API Documentation (Postman)  
✅ Middleware for Validation & Logging  
✅ Unit and Integration Testing  

## 🏷️ Architecture
This project follows **MVC & Service-Oriented Architecture**:
- **Controllers:** Handle API requests and responses
- **Services:** Contain business logic and interact with repositories
- **Repositories:** Manage database queries and transactions
- **Middleware:** Handle validation, logging, and request processing
- **Database:** SQLite using GORM ORM for easy management
- **Routing:** Centralized API route management

## 📁 Folder Structure
```
supermarket-api/
    │   .env                   # Environment variables (e.g., database connection)
    │   Dockerfile             # Docker configuration for containerization
    │   go.mod                 # Go module dependencies
    │   go.sum                 # Dependency checksums
    │   main.go                # Application entry point
    │   readme.md              # Documentation
    │   supermarket-api.exe     # Compiled binary (if built)
    │   supermarket.db         # SQLite database file
    │
    ├───config
    │       database.go        # Database connection configuration
    │
    ├───controllers           # Handles HTTP requests and responses
    │       product_controller.go  # Product-related API handlers
    │       transaction_controller.go  # Transaction-related API handlers
    │
    ├───middleware            # Middleware for logging, validation, etc.
    │       logger.go         # Logs API requests and errors
    │       validation.go     # Validates request payloads
    │
    ├───models                # Database models (GORM)
    │       product.go        # Product model
    │       transaction.go    # Transaction model
    │
    ├───repositories          # Data persistence and queries
    │       product_repo.go   # Product database operations
    │       transaction_repo.go  # Transaction database operations
    │
    ├───routes                # API route definitions
    │       routes.go         # Centralized API route definitions
    │
    ├───services              # Business logic and service layer
    │       product_service.go  # Product service logic
    │       transaction_service.go  # Transaction service logic
    │
    ├───tests                 # Unit and integration tests
    │       product_test.go   # Tests for product API
    │       transaction_test.go  # Tests for transaction API
    │
    └───utils                 # Utility functions (e.g., API responses)
            response.go       # Standardized API response handling
```

## 📌 API Endpoints
### Product Endpoints
- **GET /api/products** - Get all products
- **POST /api/products** - Add a new product
- **GET /api/products/:id** - Get a product by ID
- **PUT /api/products/:id** - Update a product by ID
- **DELETE /api/products/:id** - Delete a product by ID

### Transaction Endpoints
- **GET /api/transactions** - Get all transactions
- **POST /api/transactions** - Create a new transaction (checkout)

## 🔧 Setup & Installation
### Prerequisites
- Install **Go** (latest version)
- Install **Docker** (optional, for containerization)
- Install **Postman Desktop App** or via **Website** (for API testing)

### Steps
1. **Clone the repository**
   ```sh
   git clone https://github.com/your-username/supermarket-api.git
   cd supermarket-api
   ```
2. **Install dependencies**
   ```sh
   go mod tidy
   ```
3. **Run the application**
   ```sh
   go run main.go
   ```
4. **Test API with Postman or Curl**
   ```sh
   curl -X GET http://localhost:8080/api/products
   ```

## 📜 Documentation
API documentation is available in **Postman Collection**. [Download the Postman Collection](docs/supermarket-api.postman_collection.json) and import the collection to test endpoints.

## 📌 Deployment
### Run with Docker
```sh
docker build -t supermarket-api .
docker run -p 8080:8080 supermarket-api
```

or 

```sh
docker-compose up -d
```


## 🏗️ Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Added new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Create a Pull Request

## 📜 License
This project is licensed under the **MIT License**.

---
**🚀 Supermarket API - A Personal Project for Mastering Golang**

