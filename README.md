# Company Profile API with Golang and NuxtJS 3

This project is a modern and high-performance web application for a company profile, built using Golang (Echo) for the backend and NuxtJS 3 for the frontend. The backend handles authentication, content management, and API services, while the frontend provides a dynamic and user-friendly interface.

## Features
- User authentication (JWT-based)
- Company profile management
- File upload (Supabase or other cloud storage integration)
- RESTful API development
- Database management with PostgreSQL

## Technologies Used
### Backend:
- Golang (Echo Framework)
- GORM (ORM for Golang)
- PostgreSQL (Database)
- Supabase (Cloud Storage for file uploads)
- JWT Authentication
- Docker
### Backend Setup

# Clone the repository
git clone https://github.com/prudencadaffa/latihan-compro.git
cd latihan-compro

# Create an environment file
cp .env.example .env

# Update the .env file with database credentials

# Install dependencies
go mod tidy

# Run the application
go run main.go
