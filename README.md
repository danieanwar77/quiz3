# API Documentation

## Introduction

This API provides endpoints for managing users, categories, and books. It allows for CRUD (Create, Read, Update, Delete) operations on these entities.

## Endpoints

### Users

- **Get all users**
  ```
  GET /users
  ```
  Retrieves a list of all users.

- **Get a specific user**
  ```
  GET /users/:id
  ```
  Retrieves details of a user by their ID.

- **Create a new user**
  ```
  POST /users
  ```
  Creates a new user.

- **Update an existing user**
  ```
  PUT /users/:id
  ```
  Updates details of a user by their ID.

- **Delete a user**
  ```
  DELETE /users/:id
  ```
  Deletes a user by their ID.

### Categories

- **Get all categories**
  ```
  GET /categories
  ```
  Retrieves a list of all categories.

- **Get a specific category**
  ```
  GET /categories/:id
  ```
  Retrieves details of a category by its ID.

- **Create a new category**
  ```
  POST /categories
  ```
  Creates a new category.

- **Update an existing category**
  ```
  PUT /categories/:id
  ```
  Updates details of a category by its ID.

- **Delete a category**
  ```
  DELETE /categories/:id
  ```
  Deletes a category by its ID.

### Books

- **Get all books**
  ```
  GET /books
  ```
  Retrieves a list of all books.

- **Get a specific book**
  ```
  GET /books/:id
  ```
  Retrieves details of a book by its ID.

- **Create a new book**
  ```
  POST /books
  ```
  Creates a new book.

- **Update an existing book**
  ```
  PUT /books/:id
  ```
  Updates details of a book by its ID.

- **Delete a book**
  ```
  DELETE /books/:id
  ```
  Deletes a book by its ID.

## Error Handling

Each endpoint may return standard HTTP status codes to indicate the success or failure of an API call. Additionally, error responses may include a message describing the nature of the error.
```
