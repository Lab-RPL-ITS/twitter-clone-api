# Twitter Clone API 🐦

A modern Twitter clone API built with Go, Gin, and GORM following Clean Architecture principles. This project provides a complete backend solution for a Twitter-like social media platform with user authentication, posts, likes, and real-time logging.

## Introduction 👋

This project implements Clean Architecture principles in Go, providing a well-structured, maintainable, and scalable API for a Twitter clone application. The architecture separates concerns into distinct layers (controllers, services, repositories, entities) making the codebase easy to understand, test, and extend.

### Key Features ✨

- **User Management**: Registration, login, profile management
- **Post System**: Create, read, update, delete posts
- **Like System**: Like and unlike posts
- **JWT Authentication**: Secure user authentication
- **Real-time Logging**: Built-in logging system with web interface
- **Clean Architecture**: Well-structured codebase following best practices
- **Docker Support**: Easy deployment with Docker and Docker Compose
- **Database Migrations**: Automated database schema management
- **Testing**: Comprehensive test coverage

## API Endpoints 📋

### User Endpoints (`/api/user`)
- `POST /register` - User registration
- `POST /login` - User authentication
- `POST /check-username` - Check username availability
- `GET /me` - Get current user profile (authenticated)
- `GET /:username` - Get user by username
- `GET /:username/posts` - Get posts by user
- `PATCH /update` - Update user profile (authenticated)

### Post Endpoints (`/api/post`)
- `POST /` - Create new post (authenticated)
- `GET /:post_id` - Get post by ID
- `DELETE /:post_id` - Delete post (authenticated)
- `PUT /:post_id` - Update post (authenticated)
- `GET /` - Get all posts

### Like Endpoints (`/api/likes`)
- `PUT /:post_id` - Like a post (authenticated)
- `DELETE /:post_id` - Unlike a post (authenticated)

## Logs Feature 📊

The application includes a built-in logging system that allows you to monitor and track system queries. You can access the logs through a modern, user-friendly interface.

### Accessing Logs
To view the logs:
1. Make sure the application is running
2. Set `IS_LOGGER=true` in your environment variables
3. Open your browser and navigate to:
```bash
http://your-domain/logs
```

### Features
- **Monthly Filtering**: Filter logs by selecting different months
- **Real-time Refresh**: Instantly refresh logs with the refresh button
- **Expandable Entries**: Click on any log entry to view its full content
- **Modern UI**: Clean and responsive interface with glass-morphism design

## Prerequisites 🏆
- Go Version `>= 1.23.0`
- PostgreSQL Version `>= 15.0`
- Docker and Docker Compose (for containerized setup)

## Quick Start 🚀

### Option 1: Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/Lab-RPL-ITS/twitter-clone-api.git
   cd twitter-clone-api
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start the application**
   ```bash
   make up
   ```

4. **Initialize database**
   ```bash
   make init-uuid
   make migrate-seed
   ```

5. **Access the API**
   - API: `http://localhost:8888`
   - Nginx: `http://localhost:81`
   - Logs: `http://localhost:8888/logs` (if IS_LOGGER=true)

### Option 2: Local Development

1. **Clone and navigate**
   ```bash
   git clone https://github.com/Lab-RPL-ITS/twitter-clone-api.git
   cd twitter-clone-api
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Configure your PostgreSQL credentials in .env
   ```

3. **Set up PostgreSQL**
   ```bash
   # Create database and enable UUID extension
   psql -U postgres
   CREATE DATABASE your_database;
   \c your_database
   CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
   \q
   ```

4. **Install dependencies and run**
   ```bash
   go mod tidy
   go run main.go
   ```

## Available Commands 🛠️

### Docker Commands
```bash
make up          # Start all services
make down        # Stop all services
make logs        # View logs
make init-docker # Build and start services
```

### Database Commands
```bash
make migrate     # Run database migrations
make seed        # Seed database with initial data
make migrate-seed # Run both migrations and seeding
make init-uuid   # Initialize UUID extension
```

### Development Commands
```bash
make run         # Run the application locally
make build       # Build the application
make test        # Run tests
make dep         # Install dependencies
```

### Command Line Arguments
The application supports various command line arguments:

```bash
go run main.go --migrate --seed --run --script:example_script
```

- `--migrate` - Apply database migrations
- `--seed` - Seed database with initial data
- `--script:script_name` - Run a specific script
- `--run` - Keep the application running after executing commands

## Project Structure 📁

```
twitter-clone-api/
├── config/          # Configuration files
├── controller/      # HTTP controllers
├── dto/            # Data Transfer Objects
├── entity/         # Database entities/models
├── helpers/        # Helper functions
├── middleware/     # HTTP middleware
├── migrations/     # Database migrations
├── provider/       # Dependency injection
├── repository/     # Data access layer
├── routes/         # API route definitions
├── script/         # Utility scripts
├── service/        # Business logic layer
├── tests/          # Test files
├── utils/          # Utility functions
├── docker/         # Docker configuration
├── main.go         # Application entry point
├── go.mod          # Go module file
├── docker-compose.yml
└── Makefile        # Build and deployment commands
```

## Environment Variables 🔧

Create a `.env` file with the following variables:

```env
# Application
APP_NAME=twitter-clone-api
APP_ENV=development
PORT=8888
IS_LOGGER=true

# Database
DB_HOST=localhost
DB_USER=postgres
DB_PASS=your_password
DB_NAME=twitter_clone
DB_PORT=5432

# JWT
JWT_SECRET=your_jwt_secret_key

# Email (optional)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=your_email@gmail.com
SMTP_PASS=your_email_password
```

## API Documentation 📚

### Postman Collection
You can explore the available endpoints and their usage in the [Postman Documentation](https://documenter.getpostman.com/view/28076994/2sB2cYbf38). This documentation provides a comprehensive overview of the API endpoints, including request and response examples.

### Authentication
Most endpoints require JWT authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

## Development 🔧

### Hot Reload
For development with hot reload, the project includes Air configuration:
```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

### Testing
Run the test suite:
```bash
make test
```

### Code Quality
The project follows Go best practices and Clean Architecture principles:
- Separation of concerns
- Dependency injection
- Interface-based design
- Comprehensive error handling

## Contributing 🤝

We welcome contributions! Please see our contributing guidelines:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

### Issue Templates
The repository includes templates for issues and pull requests to standardize contributions and improve the quality of discussions and code reviews.

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support 💬

If you have any questions or need help, please:
- Open an issue on GitHub
- Check the existing issues for solutions
- Review the API documentation

---

**Happy Coding! 🚀**