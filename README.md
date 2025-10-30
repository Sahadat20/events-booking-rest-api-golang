# 🎟️ Event Booking REST API (Go)

A Go-powered **Event Booking** REST API for managing events, user authentication, and event registration.  
This project demonstrates how to build a complete backend with Go, REST principles, and JWT authentication.

---

## 🚀 Features

- User authentication (signup & login)
- CRUD operations for events
- Event registration & cancellation
- Secure API routes with authentication middleware
- JSON-based responses
- RESTful route structure

---

## 📚 API Endpoints

| Method | Endpoint | Description |
|--------|-----------|-------------|
| **GET** | `/events` | Get a list of available events |
| **GET** | `/events/{id}` | Get details of a specific event |
| **POST** | `/events` | Create a new bookable event |
| **PUT** | `/events/{id}` | Update an existing event |
| **DELETE** | `/events/{id}` | Delete an event |
| **POST** | `/signup` | Create a new user |
| **POST** | `/login` | Authenticate an existing user |
| **POST** | `/events/{id}/register` | Register a user for an event |
| **DELETE** | `/events/{id}/register` | Cancel a registration |

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** net/http
- **Database:** SQLite / PostgreSQL (depending on configuration)
- **Authentication:** JWT (JSON Web Token)
- **Tools:** Go modules, Gorilla/Mux (for routing)

---

## ⚙️ Installation & Setup

### 1. Clone the repository
```bash
git clone https://github.com/your-username/event-booking-api.git
cd event-booking-api
```

### 2. Initialize Go modules
```bash
go mod tidy
```

### 3. Configure environment variables
Create a `.env` file in the project root:
```env
PORT=8080
JWT_SECRET=your_secret_key
DB_PATH=./events.db
```

### 4. Run the server
```bash
go run main.go
```

Server will start on:  
👉 `http://localhost:8080`

---

## 🧪 Example Requests

### Create an Event
```bash
POST /events
Content-Type: application/json

{
  "title": "Tech Conference 2025",
  "description": "A conference for developers.",
  "date": "2025-12-10",
  "location": "Dhaka, Bangladesh"
}
```

### Register for an Event
```bash
POST /events/1/register
Authorization: Bearer <your_jwt_token>
```

---

## 🧩 Folder Structure

```
├── main.go
├── controllers/
│   ├── event_controller.go
│   ├── user_controller.go
├── models/
│   ├── event.go
│   ├── user.go
├── routes/
│   ├── event_routes.go
│   ├── user_routes.go
├── middlewares/
│   └── auth_middleware.go
├── database/
│   └── db.go
├── go.mod
└── README.md
```

---

## 🔒 Authentication

JWT (JSON Web Token) is used for user authentication.  
Users must include a valid JWT token in the header for protected routes:

```
Authorization: Bearer <token>
```

---

## 🤝 Contributing

Contributions are welcome!  
1. Fork the project  
2. Create your feature branch (`git checkout -b feature/new-feature`)  
3. Commit your changes (`git commit -m "Add new feature"`)  
4. Push to the branch (`git push origin feature/new-feature`)  
5. Open a Pull Request  

---


---

## 👨‍💻 Author

**Sahadat Hossain**  
🌐 [GitHub](https://github.com/sahadat20) 
