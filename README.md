# Go Product API

## Overview

This repository contains a REST API built with Go that allows for the management of products. The API supports CRUD operations, including creating, reading, updating, and deleting products. Additionally, it includes user authentication functionalities with login and registration.

## Features

- **CRUD Operations for Products**
  - **Create a Product**: `POST /products`
  - **Get All Products**: `GET /products`
  - **Get a Product by ID**: `GET /products/{id}`
  - **Update a Product**: `PUT /products/{id}`
  - **Delete a Product**: `DELETE /products/{id}`

- **Authentication**
  - **Register a User**: `POST /register`
  - **Login a User**: `POST /login`

## API Endpoints

### Products

- **Create a Product**
  - `POST /products`
  - **Request Body**: JSON containing `name`, `description`, and `price`
  - **Response**: JSON containing the created product with all fields, including `id`, `created_at`, and `updated_at`

- **Get All Products**
  - `GET /products`
  - **Response**: JSON array containing all products

- **Get a Product by ID**
  - `GET /products/{id}`
  - **Response**: JSON containing the product with the specified `id`
  - **Error Response**: 404 Not Found if the product does not exist

- **Update a Product**
  - `PUT /products/{id}`
  - **Request Body**: JSON containing `name`, `description`, and/or `price`
  - **Response**: JSON containing the updated product
  - **Error Response**: 404 Not Found if the product does not exist

- **Delete a Product**
  - `DELETE /products/{id}`
  - **Response**: 204 No Content status if the deletion is successful
  - **Error Response**: 404 Not Found if the product does not exist

### Authentication

- **Register a User**
  - `POST /register`
  - **Request Body**: JSON containing `username`, `password`
  - **Response**: JSON containing a success message or error

- **Login a User**
  - `POST /login`
  - **Request Body**: JSON containing `username`, `password`
  - **Response**: JSON containing authentication token or error

