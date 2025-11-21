# API Documentation

## Server Endpoints

Base URL: `http://localhost:8080/api`

### Users

- **POST /api/users** - Create a new user
  ```json
  {
    "id": "string",
    "name": "string",
    "email": "string",
    "phone": "string"
  }
  ```

- **GET /api/users/:id** - Get user by ID

- **GET /api/users** - Get all users

### Payments

- **POST /api/payments/initiate** - Initiate a payment
  ```json
  {
    "id": "string",
    "user_id": "string",
    "travel_id": "string",
    "amount": 0.0,
    "status": "string",
    "payment_method": "string"
  }
  ```

- **PUT /api/payments/:id** - Update payment
  ```json
  {
    "status": "string"
  }
  ```

- **PUT /api/payments/:id/status** - Update payment status
  ```json
  {
    "status": "string"
  }
  ```

- **GET /api/payments?user_id=xxx** - Get all payments for a user

### Places

- **POST /api/places** - Create a new place
  ```json
  {
    "id": "string",
    "name": "string",
    "address": "string",
    "latitude": 0.0,
    "longitude": 0.0
  }
  ```

- **GET /api/places/:id** - Get place by ID

- **GET /api/places** - Get all places

### Drivers

- **POST /api/drivers** - Create a new driver
  ```json
  {
    "id": "string",
    "name": "string",
    "phone": "string",
    "license_no": "string",
    "vehicle_type": "string",
    "vehicle_no": "string",
    "rating": 0.0,
    "is_available": true
  }
  ```

- **GET /api/drivers/:id** - Get driver by ID

- **GET /api/drivers** - Get all available drivers

### Transportation Fare

- **POST /api/calculate-fare** - Calculate transportation fare
  ```json
  {
    "distance": 0.0,
    "vehicle_type": "standard"
  }
  ```
  
  Response:
  ```json
  {
    "distance": 0.0,
    "vehicle_type": "standard",
    "fare": 0.0
  }
  ```

  Fare calculation:
  - Base fare: 500
  - Standard rate: 200 per km
  - Premium rate: 300 per km

## Running the Server

```bash
go run cmd/main.go
```

Environment variables:
- `DATABASE_URL` - PostgreSQL connection string (default: `postgres://postgres:postgres@localhost:5432/clando?sslmode=disable`)
- `PORT` - Server port (default: `8080`)
