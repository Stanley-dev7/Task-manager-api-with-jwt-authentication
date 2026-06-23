# Task Manager API

A lightweight Go baTckend for managing personal tasks with secure authentication.

This project demonstrates a practical REST API for user registration, login, and user-scoped task management using JWT authentication and SQLite.

## Table of contents

- [Overview](#overview)
- [Features](#features)
- [Technology stack](#technology-stack)
- [Prerequisites](#prerequisites)
- [Configuration](#configuration)
- [Getting started](#getting-started)
- [API reference](#api-reference)
- [Repository structure](#repository-structure)
- [Security notes](#security-notes)

## Overview

This API allows users to:

- register a new account
- login and receive a JWT token
- create, read, update, and delete their own tasks

Each task belongs to a single user, and protected endpoints require a valid bearer token.

## Features

- user sign-up and login
- password hashing with bcrypt
- JWT authentication for protected routes
- middleware-based authorization
- task CRUD operations scoped to the authenticated user
- SQLite database for development

## Technology stack

- Go (Golang)
- Gorilla Mux for routing
- SQLite for persistent storage
- JWT for stateless authentication
- bcrypt for secure password hashing

## Prerequisites

- Go 1.20 or later
- git

## Configuration

This project uses environment variables for configuration.

Example values:

- `PORT` — server port (default: `8080`)
- `DB_PATH` — SQLite database file path (default: `./data.db`)
- `JWT_SECRET` — secret key for signing JWT tokens (must be set for production)

If you do not set these values, the application falls back to sensible defaults for local development.

## Getting started

Clone the repository and run the service locally.

```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd Task-manager-api-with-jwt-authentication
go mod tidy
go run .
```

After startup, the API is available at `http://localhost:8080`.

## API reference

All endpoints use JSON request and response payloads unless noted.

### Register user

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"s3cr3t"}'
```

Success response:

```json
{
  "id": 1,
  "email": "alice@example.com"
}
```

### Login user

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"alice@example.com","password":"s3cr3t"}'
```

Success response:

```json
{
  "token": "<JWT_TOKEN>",
  "expires_in": 3600
}
```

### Create task

Protected route — include the token in the Authorization header.

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{"title":"Buy groceries","description":"Milk and eggs"}'
```

### Get tasks

```bash
curl -X GET http://localhost:8080/api/tasks \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

### Update task

```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{"title":"Buy groceries and fruit"}'
```

### Delete task

```bash
curl -X DELETE http://localhost:8080/api/tasks/1 \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

## Repository structure

- `main.go` — application entry point
- `database/` — database initialization and connection logic
- `handlers/` — HTTP handler functions for auth and tasks
- `middleware/` — JWT authentication middleware
- `models/` — request and database models
- `README.md` — project documentation

## Security notes

- Keep `JWT_SECRET` private and do not commit it to source control.
- Use HTTPS for production deployments.
- Add rate limiting and input validation before exposing this API publicly.

## Challenges I Solved

During development, I worked through:

- Implementing JWT authentication and protected routes
- Creating middleware for authorization
- Managing user-specific task access
- Connecting Go handlers with database operations
- Debugging API routing issues
- Deploying and managing the project with Git/GitHub

## Future improvements

- add automated tests and CI integration
- add OpenAPI/Swagger documentation
- improve error handling and request validation
- support role-based access control
- add docker/docker-compose support
