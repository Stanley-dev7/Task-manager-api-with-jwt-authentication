Task Manager API (Go + JWT + SQLite)
A backend REST API built with Go that supports authentication and user-based todo management.

 Features

- User registration and login
- JWT authentication
- Protected routes using middleware
- Create, read, update, delete todos
- Each user only accesses their own data
- SQLite database integration

 What I Learned

- REST API development in Go
- JWT authentication flow
- Middleware design and usage
- Database integration with SQLite
- Structuring backend applications


Challenges Faced

- Debugging invalid JWT tokens
- Fixing SecretKey mismatch issues
- Handling database UNIQUE constraints
- Understanding middleware request flow
- Fixing routing and handler issues


Tech Stack

- Go (Golang)
- SQLite
- JWT
- REST API


 API Routes

 Auth
- POST /register
- POST /login

Todos (Protected)
- GET /todos
- POST /todos
- GET /todos/{id}
- PUT /todos/{id}
- DELETE /todos/{id}
