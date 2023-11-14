# Golang CRUD Rest API Book using PostgreSQL

This repository contains a simple CRUD (Create, Read, Update, Delete) Rest API for managing books written in Golang. The application uses PostgreSQL as its database to store book information. This project serves as a practical example for building a scalable and efficient API with Go and integrating it with a PostgreSQL database.

## Features

- **Create:** Add new books to the database with relevant details such as title, author, and publication date.
- **Read:** Retrieve a list of all books or fetch details of a specific book using its unique identifier.
- **Update:** Modify existing book information, including title, author, and publication date.
- **Delete:** Remove a book from the database based on its identifier.

## Technologies Used

- **Golang:** The backend is developed in Go, a statically typed, compiled programming language known for its simplicity and efficiency.
- **PostgreSQL:** The database management system used to store and retrieve book data.
- **Gorilla Mux:** A powerful URL router and dispatcher for matching incoming requests to their respective handler functions.

## Tools Used
- **fl0:** For deploy backend applications and databases in minutes.
- **Postman:** Testing API

## Getting Started

Follow these steps to set up and run the Golang CRUD Rest API on your local machine:

1. Clone the repository: `git clone https://github.com/zoelabbb/gocrud_psql.git`
2. Install dependencies: `go mod tidy`
3. Configure PostgreSQL connection in the `.env` file.
4. Run the application: `go run main.go`
5. Access the API at `http://localhost:8080`

Feel free to explore and extend this project to suit your requirements. Contributions and suggestions are welcome!

## Contributing

If you'd like to contribute to this project, please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

