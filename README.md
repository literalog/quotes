# Quotes API
This is a simple RESTful API that allows you to interact with a collection of quotes stored in MongoDB.

## Features
- Get all quotes
- Get a quote by Id
- Create a new quote
- Update a quote
- Delete a quote

## Routes
The API server has two main routes:
- `/quotes`: This route is used to get all quotes or create a new quote.
- `/quotes/{id}`: This route is used to get, update, or delete a quote by its ID.

## Usage
To run the project, you will need [Go](https://golang.org/) and [MongoDB](https://www.mongodb.com/) installed on your system. Here are the steps to run the project:

1. Clone the repository:
```bash
git clone https://github.com/beatrizhub/quotes.git
```

2. Run the project:
```
go run main.go
```
