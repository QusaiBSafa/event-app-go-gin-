# Event App

## Overview

The Event App is a simple application for managing events, built using the Go programming language with the Gin framework. It provides endpoints for user authentication, event creation, and event retrieval. The application includes Swagger documentation for API exploration.

## Features

- **User Authentication**: Register and log in users.
- **Event Management**: Create and view events.
- **Swagger Documentation**: Interactive API documentation.

## Getting Started

### Prerequisites

- Go (1.18 or later)
- PostgreSQL (for database management)
- Swagger (for API documentation generation)

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-username/event-app.git
   cd event-app
   ```
2. **Install Dependencies**
```
go mod tidy
```
3. **Install Swagger**

```
go install github.com/swaggo/swag/cmd/swag@latest
```
4. **Setup Environment**
Create a .env file in the root directory and configure the following environment variables:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your-db-user
DB_PASSWORD=your-db-password
DB_NAME=event_app
```

5. **Generate Swagger Documentation**
```
swag init
```

### Running the Application
```
go run main.go
```

### Access Swagger UI
```
http://localhost:8080/swagger/index.html
```

### Contributing
If you would like to contribute to the project, please fork the repository and submit a pull request with your changes.

### Contact
For any questions or feedback, please reach out to qusi.bassam@gmail.com
