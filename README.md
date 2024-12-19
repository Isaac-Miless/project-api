# Project API

A simple RESTful API built in Go (Golang) to manage projects. This API demonstrates:

- Connecting to a PostgreSQL database.
- Handling basic CRUD operations (Create, Read, Update, Delete) for project resources.
- Using environment variables for database credentials.
- Serving JSON responses.

## Features

- **GET /projects**: Retrieve all projects.
- **GET /project?id=<PROJECT_ID>**: Retrieve a single project by its ID.
- **POST /projects**: Create a new project.
- **PUT /project?id=<PROJECT_ID>**: Update an existing project.
- **DELETE /project?id=<PROJECT_ID>**: Delete a project.

## Project Structure

- **database/connection.go**: Handles the database connection to PostgreSQL.
- **handlers/project_handlers.go**: Functions that implement CRUD operations for projects.
- **models/project.go**: Defines the `Project` struct representing a project resource.
- **routes/routes.go**: Sets up HTTP routes and associates them with their handlers.
- **main.go**: The entry point of the application. Connects to the DB, initializes routes, and starts the server.

## Prerequisites

- **Go 1.20+** (or your installed version)
- **PostgreSQL** installed and running locally.

For macOS with Homebrew:

```bash
brew install postgresql@14
brew services start postgresql@14
```

Ensure you’ve created a role and database for this app. For example, in psql:

```sql
-- Connect to psql (adjust username and database as needed)
psql -U <your_username> -d postgres

-- Create a role and database
CREATE ROLE postgres WITH LOGIN SUPERUSER PASSWORD 'psqlPW1';
CREATE DATABASE projects_api;
GRANT ALL PRIVILEGES ON DATABASE projects_api TO postgres;

\c projects_api

-- Create the projects table
CREATE TABLE projects (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    technologies TEXT[],
    github TEXT
);
```

## Running the Application

1.  **Install Dependencies**

    ```bash
    go mod tidy
    ```

2.  **Run the Server**

    ```bash
    go run main.go
    ```

    - Expected Output:

    ```
    Database connected successfully!
    Server started at :8080
    ```

3.  **Test Endpoints**

    - Get all Projects

    ```bash
    curl http://localhost:8080/projects
    ```

    - Create a new Project

    ```bash
    curl -X POST -H "Content-Type: application/json"
    -d '{"name":"MyProject","description":"Testing","tech_stack":["Go","Postgres"],"github_url":"https://github.com/example/myproject"}'
    http://localhost:8080/projects
    ```

    - Retrieve a Single Project

    ```bash
    curl "http://localhost:8080/project?id=<PROJECT_ID>"
    ```

    - Update a Project

    ```bash
    curl -X PUT -H "Content-Type: application/json"
    -d '{"name":"UpdatedName","description":"Updated","tech_stack":["Go","Docker"],"github_url":"https://github.com/example/updated"}'
    "http://localhost:8080/project?id=<PROJECT_ID>"
    ```

    - Delete a Project

    ```bash
    curl -X DELETE "http://localhost:8080/project?id=<PROJECT_ID>"
    ```
