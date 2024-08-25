# GoBooks

## Project Overview

**GoBooks** is an API and CLI tool for managing a collection of books. The project provides a RESTful API for CRUD operations. Additionally, the CLI offers similar functionality for those who prefer a command-line interface.

### Features

- **API**: A RESTful service built with Go for creating, reading, updating, and deleting book records.
- **CLI**: A command-line tool for managing your book collection directly from the terminal.
- **Database**: MySQL integration with automatic migrations and seed data setup.
- **Docker**: Fully containerized using Docker for easy setup and deployment.

## Prerequisites

Before getting started, ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/) (optional, for running the CLI)

## Getting Started

Follow these steps to get started with GoBooks.

### Running the API with Docker

1. **Clone the repository**:
   ```bash
   git clone https://github.com/mwives/gobooks.git
   cd gobooks
   ```
2. **Build and run the Docker container**:
   ```bash
   docker-compose up --build
   ```
   This command will build the Docker image, start the container, automatically apply database migrations, and insert seed data.

### Accessing the API

Once the container is up, the API will be accessible at http://localhost:8080. You can test the API endpoints using VS Code's REST Client extension with the `requests.http` file, or by using an API client like Postman.

#### API Endpoints

- **GET** `/books?title=<book_title>`: Retrieve a list of all books. You can also search for books by title using a query parameter.
- **GET** `/books/{id}`: Retrieve a specific book by ID.
- **POST** `/books`: Add a new book by providing the title, author, and genre in the request body.
- **POST** `/books/read`: Simulate reading one or more books concurrently by sending a list of `book_ids` in the request body.
- **PUT** `/books/{id}`: Update an existing book by ID.
- **DELETE** `/books/{id}`: Delete a book by ID.

### Using the CLI:

You can also use the CLI inside the Docker container to manage books:

1. **Ensure the Database is Running**:

   ```bash
   docker-compose up -d db
   ```

2. **Run the CLI**:

   ```bash
   go run cmd/gobooks/main.go <command>
   ```

Replace `<command>` with one of the following:

- ~~add "Book Title" "Author" "Genre": Add a new book to the collection.~~ (TODO)
- ~~list: List all books in the collection.~~ (TODO)
- search "<book_name>": Search for books by name.
- simulate <book_id(s)>: Simulate reading one or more books concurrently.
