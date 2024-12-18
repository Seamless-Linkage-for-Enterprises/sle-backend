# Golang Backend with PostgreSQL and Gin

This project is a backend API built in Golang using the Gin framework. It utilizes PostgreSQL for data storage and follows a clean architecture with well-organized modules. The project is designed to manage four core entities: Seller, Buyer, Product, and Bookmark.

## 📂 Project Structure

```bash 
.
├── Dockerfile              # Dockerfile to build the image for deployment
├── Makefile                # Makefile to automate tasks (build, run, test)
├── cmd/
│   └── main.go            # Main entry point for the application
├── config/
│   └── config.go          # Configuration management (loading environment variables)
├── database/
│   ├── database.go        # Database connection setup
│   └── migrations/        # SQL migration files
├── dburls.txt             # Database URLs or references for migrations
├── go.mod                 # Go module file
├── go.sum                 # Go sum file for dependency management
├── helpers/
│   ├── hash_password.go   # Utility functions (e.g., hashing passwords)
│   └── helpers.go         # General helper functions
├── internal/
│   ├── apiservice.go      # Main API service logic
│   ├── bookmark/          # Bookmarks-related business logic
│   ├── buyer/             # Buyers-related business logic
│   ├── product/           # Products-related business logic
│   └── seller/            # Sellers-related business logic
├── middleware/
│   └── middleware.go      # Custom middleware functions
├── routes/
│   └── routes.go          # Routes definition
└── utils/
    └── utils.go           # Utility functions

```

## Features

### Seller Features:

- Signup: Allows sellers to create an account.
- Login: Allows sellers to log in to their accounts.
- Get All Sellers: Retrieve a list of all sellers in the system.
- Get Seller by ID: Retrieve a specific seller by their ID.
- Delete Seller: Remove a seller's account from the system.
- Forgot Password: Allow sellers to reset their password.

### Buyer Features:

- Signup: Allows buyers to create an account.
- Login: Allows buyers to log in to their accounts.
- Get All Buyers: Retrieve a list of all buyers in the system.
- Get Buyer by ID: Retrieve a specific buyer by their ID.
- Delete Buyer: Remove a buyer's account from the system.
- Get Buyer by Phone Number: Retrieve a buyer by their phone number.

### Product Features:

- Add Product: Allows sellers to add new products to the system.
- Get All Products: Retrieve a list of all products in the system.
- Get Product by ID: Retrieve a specific product by its ID.
- Delete Product: Allows sellers to delete a product from the system.
- Update Product: Allows sellers to update the product details.

### Bookmark Features:

- Add Bookmark: Allows users to bookmark a product for later reference.
- Get All Bookmarks: Retrieve a list of all bookmarks.
- Delete Bookmark: Allows users to delete a specific bookmark.

## Setup & Installation

### Prerequisites
- Go 1.23 or above
- Docker (for local PostgreSQL)
- PostgreSQL (for local or cloud database)
- Make (for automating common tasks)

## Docker Deployment (Render.com)

1. Docker Hub Image: The Docker image for this project is available on Docker Hub: ```bash manankoyawala/sleapp:v1.2```.

2. Deploy on Render.com:

- Link your Render account to your GitHub repository.
- Set the environment variables as mentioned earlier in the Render settings.
- Render will automatically build the project using the Dockerfile and deploy it.

## Database Schema

The application handles four main tables:

- Sellers: Stores seller information.
- Buyers: Stores buyer information.
- Products: Stores product listings.
- Bookmarks: Stores user bookmarks for products.

Each table has associated migration files to manage schema changes (up.sql for creation and down.sql for deletion).

## Environment Variables

The following environment variables are required to connect to the database:

- ```DB_USERNAME```: The username for your PostgreSQL database.
- ```DB_PASSWORD```: The password for your PostgreSQL database.
- ```DB_NAME```: The name of the database.
- ```DB_HOST```: The host where the database is running (e.g., ```localhost``` for local, AWS endpoint for production).
- ```DB_PORT```: The port on which PostgreSQL is running (default is ```5432```).
- ```DB_SSLMODE```: SSL mode for database connection (```disable```, ```require```, etc.).

## 📜 [LICENSE](LICENSE)

## Acknowledgements

**Gin**: Web framework used for building the API.
**PostgreSQL**: Database used for persistence.
**Docker**: Used for containerization and local development.