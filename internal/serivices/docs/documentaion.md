# Services Documentation

## Overview
This document provides an overview and details for the `AuthService`, `UserService`, and `BlogService` used in your application. Each service is designed as an independent microservice, communicating via the API Gateway.

---

## AuthService

**Purpose**:
The `AuthService` handles authentication and authorization tasks, such as user login, registration, token generation, and token validation.

**Service Endpoint**:
- Base URL: `http://localhost:8083`

**Available Routes**:
1. **`GET /health`**
   - **Description**: Health check endpoint to verify that the service is running.
   - **Response**:
     - **200 OK**: `"AuthService is healthy!"`

2. **`POST /login`**
   - **Description**: Handles user login by validating credentials and issuing a JWT.
   - **Request Body** (JSON):
     ```json
     {
       "username": "string",
       "password": "string"
     }
     ```
   - **Response**:
     - **200 OK**: Token issued on successful login.
     - **401 Unauthorized**: Invalid credentials.

3. **`POST /register`**
   - **Description**: Handles user registration.
   - **Request Body** (JSON):
     ```json
     {
       "username": "string",
       "password": "string",
       "email": "string"
     }
     ```
   - **Response**:
     - **201 Created**: User successfully registered.
     - **400 Bad Request**: Invalid input or user already exists.

4. **`POST /validate`**
   - **Description**: Validates JWT tokens.
   - **Request Body** (JSON):
     ```json
     {
       "token": "string"
     }
     ```
   - **Response**:
     - **200 OK**: Token is valid.
     - **401 Unauthorized**: Invalid or expired token.

---

## UserService

**Purpose**:
The `UserService` manages user-related operations, including retrieving and updating user profiles.

**Service Endpoint**:
- Base URL: `http://localhost:8082`

**Available Routes**:
1. **`GET /health`**
   - **Description**: Health check endpoint to verify that the service is running.
   - **Response**:
     - **200 OK**: `"UserService is healthy!"`

2. **`GET /profile/:id`**
   - **Description**: Retrieves user profile by ID.
   - **Path Parameters**:
     - `id`: User ID.
   - **Response**:
     - **200 OK**: User profile details.
     - **404 Not Found**: User not found.

3. **`PUT /profile/:id`**
   - **Description**: Updates user profile by ID.
   - **Path Parameters**:
     - `id`: User ID.
   - **Request Body** (JSON):
     ```json
     {
       "username": "string",
       "email": "string"
     }
     ```
   - **Response**:
     - **200 OK**: Profile successfully updated.
     - **400 Bad Request**: Invalid input.

---

## BlogService

**Purpose**:
The `BlogService` handles operations related to creating, retrieving, updating, and deleting blog posts.

**Service Endpoint**:
- Base URL: `http://localhost:8081`

**Available Routes**:
1. **`GET /health`**
   - **Description**: Health check endpoint to verify that the service is running.
   - **Response**:
     - **200 OK**: `"BlogService is healthy!"`

2. **`GET /blogs`**
   - **Description**: Retrieves a list of all blogs.
   - **Response**:
     - **200 OK**: List of blogs.

3. **`GET /blogs/:id`**
   - **Description**: Retrieves a single blog by ID.
   - **Path Parameters**:
     - `id`: Blog ID.
   - **Response**:
     - **200 OK**: Blog details.
     - **404 Not Found**: Blog not found.

4. **`POST /blogs`**
   - **Description**: Creates a new blog post.
   - **Request Body** (JSON):
     ```json
     {
       "title": "string",
       "content": "string",
       "author": "string"
     }
     ```
   - **Response**:
     - **201 Created**: Blog post successfully created.
     - **400 Bad Request**: Invalid input.

5. **`PUT /blogs/:id`**
   - **Description**: Updates an existing blog post by ID.
   - **Path Parameters**:
     - `id`: Blog ID.
   - **Request Body** (JSON):
     ```json
     {
       "title": "string",
       "content": "string"
     }
     ```
   - **Response**:
     - **200 OK**: Blog post successfully updated.
     - **400 Bad Request**: Invalid input.

6. **`DELETE /blogs/:id`**
   - **Description**: Deletes a blog post by ID.
   - **Path Parameters**:
     - `id`: Blog ID.
   - **Response**:
     - **200 OK**: Blog post successfully deleted.
     - **404 Not Found**: Blog not found.

---

## Notes
- Ensure that each service is running on its respective port before interacting with the API Gateway.
- Each service has its own health check endpoint (`GET /health`) to confirm the service's availability.
- For security, implement authentication and token validation for sensitive routes.

## API DOCUMENTATION
https://documenter.getpostman.com/view/37289771/2sAYQdkAL7