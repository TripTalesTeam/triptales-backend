# TripTales Backend

A Golang backend service for the TripTales application that provides API functionality for sharing and exploring travel experiences.

## 📋 Prerequisites

- Docker and Docker Compose
- Go 1.18+ (for local development)

## 🔧 Installation

### Step 1: Clone the repository

```bash
git clone https://github.com/yourusername/triptales-backend.git
cd triptales-backend
```

### Step 2: Set up environment variables

Create a `.env` file in the root directory with the following variables:

```
MYSQL_ROOT_PASSWORD=your_mysql_password
MYSQL_DATABASE=triptales_db
# Add any other environment variables needed by your Go application
```

### Step 3: Start the Docker containers

```bash
docker-compose up -d
```

This will start:
- MySQL database (accessible on port 3306)
- phpMyAdmin (accessible on http://localhost:8081)
- Go backend application (accessible on http://localhost:8080)


## 🚀 Running and Managing the Backend

### Start all services
```bash
docker-compose up -d
```

### View logs
```bash
docker-compose logs -f
```

### View logs for a specific service
```bash
docker-compose logs -f go-app
```

### Stop all services
```bash
docker-compose down
```

### Rebuild and restart services
```bash
docker-compose up -d --build
```

## 🔒 Database Management

The MySQL database is accessible:
- From within the Docker network at `db:3306`
- From your host machine at `localhost:3306`

phpMyAdmin is available at `http://localhost:8081`
- Username: root
- Password: The value of `MYSQL_ROOT_PASSWORD` in your .env file
---

Made with ❤️ 
