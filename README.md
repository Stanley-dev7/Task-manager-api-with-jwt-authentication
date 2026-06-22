Task Manager API 
A secure REST API built with Go that allows users to register, login, and manage personal tasks using JWT authentication.

 Features
- User registration
- Secure password hashing with bcrypt
- User login
- JWT authentication
- Protected API routes
- Create tasks
- View user-specific tasks only
- Update tasks
- Delete tasks
- Database storage (SQLite)

 Tech Stack
- Go (Golang)
- Gorilla Mux
- SQLite
- JWT (Authentication)
- bcrypt (Password hashing)

 Authentication Flow
1. User registers
2. Password is hashed
3. User logs in
4. Server returns JWT token
5. Token is used for all protected routes

How to Run
1.go mod tidy

2.go run .

3.server now runs at : http://localhost:8080

API Endpoints

Auth

Register

POST /register

Login

POST /login

Tasks (Protected Routes)

All task routes require:

Authorization: Bearer YOUR_TOKEN

Create Task

POST /api/tasks

Get Tasks

GET /api/tasks

Update Task

PUT /api/tasks/{id}

Delete Task

DELETE /api/tasks/{id}

 What I Learned
 
1.Building REST APIs in Go

2.JWT authentication

3.Middleware usage

4.Secure password storage with bcrypt

5.CRUD operations with SQLite

6.Structuring backend projects

Project Goal
This project demonstrates a secure backend system with authentication, authorization, and full CRUD functionality.
