# Go Chi Cursor Example

This project is a simple RESTful API for managing student records using Go, Chi router, and PostgreSQL.

## Prerequisites

-   Docker
-   Docker Compose

## Getting Started

1. Clone the repository:

    ```
    git clone https://github.com/khoilu-1512/go-chi-cursor-example.git
    cd go-chi-cursor-example
    ```

2. Create a `.env` file in the root directory with the following content:

    ```
    DB_HOST=db
    DB_USER=your_username
    DB_PASSWORD=your_password
    DB_NAME=your_database_name
    DB_SSLMODE=disable
    ```

    Replace `your_username`, `your_password`, and `your_database_name` with your desired values.

3. Build and run the application using Docker Compose:

    ```
    docker-compose up --build
    ```

4. The API will be available at `http://localhost:8080`

## API Endpoints

-   `GET /students`: Retrieve all students
-   `POST /students`: Create a new student
-   `GET /students/{id}`: Retrieve a specific student
-   `PUT /students/{id}`: Update a specific student
-   `DELETE /students/{id}`: Delete a specific student

## Project Structure

-   `main.go`: Entry point of the application
-   `student_handlers.go`: Contains handlers for student-related operations
-   `Dockerfile`: Defines the Docker image for the Go application
-   `docker-compose.yml`: Defines the multi-container Docker environment
-   `init.sql`: Initializes the database schema and inserts sample data

## Stopping the Application

To stop the application, use:

```
docker-compose down
```

This will stop and remove the containers. To also remove the volumes, use:

```
docker-compose down -v
```

This README provides instructions on how to set up and run the project locally using Docker and Docker Compose. It includes information about prerequisites, getting started steps, API endpoints, project structure, and how to stop the application.
