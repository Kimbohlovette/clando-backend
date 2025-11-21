# Clando Backend API Documentation

Base URL: `http://localhost:8080/api`

## Table of Contents
- [Setup](#setup)
- [Users](#users)
- [Drivers](#drivers)
- [Places](#places)
- [Payments](#payments)
- [Fare Calculation](#fare-calculation)

---

## Setup

### Prerequisites
- Go 1.23.0 or higher
- Docker and Docker Compose
- sqlc (for generating Go code from SQL)
- golang-migrate (for database migrations)

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/kimbohlovette/clando-backend.git
```

2. **Install dependencies**
```bash
go mod download
```

3. **Install required tools**
```bash
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

4. **Start PostgreSQL with Docker**
```bash
docker run -d \
  --name clando-postgres \
  -e POSTGRES_USER=clando \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=clando \
  -p 5432:5432 \
  postgres:15-alpine
```

5. **Configure environment variables**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

Example `.env` file:
```
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=clando
DB_PASSWORD=secret
DB_NAME=clando
DB_SSL_MODE=false
```

6. **Run the application**
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### Database Schema

The application uses the following tables:
- `users`: User accounts
- `drivers`: Driver profiles and vehicle information
- `places`: Locations and addresses
- `payments`: Payment transactions
- `travel_histories`: Trip records

### Development

**Run with hot reload (optional)**
```bash
# Install air for hot reload
go install github.com/air-verse/air@latest

# Run with air
air
```

**Run migrations**
```bash
# Up
migrate -path db/migrations -database "postgresql://clando:secret@localhost:5432/clando?sslmode=disable" up

# Down
migrate -path db/migrations -database "postgresql://clando:secret@localhost:5432/clando?sslmode=disable" down
```

**Regenerate sqlc code**
```bash
sqlc generate
```

---

## Users

### Create User
Creates a new user in the system.

**Endpoint:** `POST /api/users`

**Request Body:**
```json
{
  "username": "string",
  "phone": "string"
}
```

**Request Fields:**
- `username` (string, required): The user's display name
- `phone` (string, required): The user's phone number

**Response:** `201 Created`
```json
{
  "id": "uuid-string",
  "username": "string",
  "phone": "string"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "phone": "+237670000000"
  }'
```

---

### Get User by ID
Retrieves a specific user by their ID.

**Endpoint:** `GET /api/users/:id`

**URL Parameters:**
- `id` (string, required): The user's UUID

**Response:** `200 OK`
```json
{
  "id": "uuid-string",
  "username": "string",
  "phone": "string"
}
```

**Error Responses:**
- `404 Not Found`: User not found
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/users/123e4567-e89b-12d3-a456-426614174000
```

---

### Get All Users
Retrieves a list of all users.

**Endpoint:** `GET /api/users`

**Response:** `200 OK`
```json
[
  {
    "id": "uuid-string",
    "username": "string",
    "phone": "string"
  }
]
```

**Error Responses:**
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/users
```

---

## Drivers

### Create Driver
Creates a new driver in the system.

**Endpoint:** `POST /api/drivers`

**Request Body:**
```json
{
  "name": "string",
  "phone": "string",
  "license_no": "string",
  "vehicle_type": "string",
  "vehicle_no": "string",
  "rating": 0.0,
  "is_available": true
}
```

**Request Fields:**
- `name` (string, required): Driver's full name
- `phone` (string, required): Driver's phone number
- `license_no` (string, required): Driver's license number (must be unique)
- `vehicle_type` (string, required): Type of vehicle (e.g., "sedan", "suv", "premium")
- `vehicle_no` (string, required): Vehicle registration number
- `rating` (float, optional): Driver's rating (0.0-5.0, default: 0.0)
- `is_available` (boolean, optional): Driver availability status (default: true)

**Response:** `201 Created`
```json
{
  "id": "uuid-string",
  "name": "string",
  "phone": "string",
  "license_no": "string",
  "vehicle_type": "string",
  "vehicle_no": "string",
  "rating": 0.0,
  "is_available": true,
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error (e.g., duplicate license number)

**Example:**
```bash
curl -X POST http://localhost:8080/api/drivers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jean Dupont",
    "phone": "+237670000001",
    "license_no": "DL123456",
    "vehicle_type": "sedan",
    "vehicle_no": "ABC-123-XY",
    "rating": 4.5,
    "is_available": true
  }'
```

---

### Get Driver by ID
Retrieves a specific driver by their ID.

**Endpoint:** `GET /api/drivers/:id`

**URL Parameters:**
- `id` (string, required): The driver's UUID

**Response:** `200 OK`
```json
{
  "id": "uuid-string",
  "name": "string",
  "phone": "string",
  "license_no": "string",
  "vehicle_type": "string",
  "vehicle_no": "string",
  "rating": 0.0,
  "is_available": true,
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

**Error Responses:**
- `404 Not Found`: Driver not found
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/drivers/123e4567-e89b-12d3-a456-426614174000
```

---

### Get All Available Drivers
Retrieves a list of all available drivers.

**Endpoint:** `GET /api/drivers`

**Response:** `200 OK`
```json
[
  {
    "id": "uuid-string",
    "name": "string",
    "phone": "string",
    "license_no": "string",
    "vehicle_type": "string",
    "vehicle_no": "string",
    "rating": 0.0,
    "is_available": true,
    "created_at": "timestamp",
    "updated_at": "timestamp"
  }
]
```

**Error Responses:**
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/drivers
```

---

## Places

### Create Place
Creates a new place/location in the system.

**Endpoint:** `POST /api/places`

**Request Body:**
```json
{
  "name": "string",
  "address": "string",
  "latitude": 0.0,
  "longitude": 0.0
}
```

**Request Fields:**
- `name` (string, required): Name of the place
- `address` (string, required): Full address of the place
- `latitude` (float, required): Latitude coordinate
- `longitude` (float, required): Longitude coordinate

**Response:** `201 Created`
```json
{
  "id": "uuid-string",
  "name": "string",
  "address": "string",
  "latitude": 0.0,
  "longitude": 0.0,
  "created_at": "timestamp"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl -X POST http://localhost:8080/api/places \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Douala Airport",
    "address": "Douala International Airport, Cameroon",
    "latitude": 4.0061,
    "longitude": 9.7195
  }'
```

---

### Get Place by ID
Retrieves a specific place by its ID.

**Endpoint:** `GET /api/places/:id`

**URL Parameters:**
- `id` (string, required): The place's UUID

**Response:** `200 OK`
```json
{
  "id": "uuid-string",
  "name": "string",
  "address": "string",
  "latitude": 0.0,
  "longitude": 0.0,
  "created_at": "timestamp"
}
```

**Error Responses:**
- `404 Not Found`: Place not found
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/places/123e4567-e89b-12d3-a456-426614174000
```

---

### Get All Places
Retrieves a list of all places.

**Endpoint:** `GET /api/places`

**Response:** `200 OK`
```json
[
  {
    "id": "uuid-string",
    "name": "string",
    "address": "string",
    "latitude": 0.0,
    "longitude": 0.0,
    "created_at": "timestamp"
  }
]
```

**Error Responses:**
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl http://localhost:8080/api/places
```

---

## Payments

### Initiate Payment
Creates a new payment record for a travel transaction.

**Endpoint:** `POST /api/payments/initiate`

**Request Body:**
```json
{
  "user_id": "string",
  "travel_id": "string",
  "amount": 0.0,
  "status": "string",
  "payment_method": "string"
}
```

**Request Fields:**
- `user_id` (string, required): UUID of the user making the payment
- `travel_id` (string, required): UUID of the travel/trip
- `amount` (float, required): Payment amount
- `status` (string, required): Payment status (e.g., "pending", "completed", "failed")
- `payment_method` (string, required): Payment method (e.g., "mobile_money", "card", "cash")

**Response:** `201 Created`
```json
{
  "id": "uuid-string",
  "user_id": "string",
  "travel_id": "string",
  "amount": 0.0,
  "status": "string",
  "payment_method": "string",
  "created_at": "timestamp"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl -X POST http://localhost:8080/api/payments/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "travel_id": "987e6543-e21b-12d3-a456-426614174000",
    "amount": 5000.0,
    "status": "pending",
    "payment_method": "mobile_money"
  }'
```

---

### Update Payment
Updates a payment's status.

**Endpoint:** `PUT /api/payments/:id`

**URL Parameters:**
- `id` (string, required): The payment's UUID

**Request Body:**
```json
{
  "status": "string"
}
```

**Request Fields:**
- `status` (string, required): New payment status (e.g., "completed", "failed", "refunded")

**Response:** `200 OK`
```json
{
  "message": "payment updated"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl -X PUT http://localhost:8080/api/payments/123e4567-e89b-12d3-a456-426614174000 \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }'
```

---

### Update Payment Status
Updates a payment's status (alternative endpoint).

**Endpoint:** `PUT /api/payments/:id/status`

**URL Parameters:**
- `id` (string, required): The payment's UUID

**Request Body:**
```json
{
  "status": "string"
}
```

**Request Fields:**
- `status` (string, required): New payment status

**Response:** `200 OK`
```json
{
  "message": "payment status updated"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl -X PUT http://localhost:8080/api/payments/123e4567-e89b-12d3-a456-426614174000/status \
  -H "Content-Type: application/json" \
  -d '{
    "status": "completed"
  }'
```

---

### Get User Payments
Retrieves all payments for a specific user.

**Endpoint:** `GET /api/payments?user_id={user_id}`

**Query Parameters:**
- `user_id` (string, required): The user's UUID

**Response:** `200 OK`
```json
[
  {
    "id": "uuid-string",
    "user_id": "string",
    "travel_id": "string",
    "amount": 0.0,
    "status": "string",
    "payment_method": "string",
    "created_at": "timestamp"
  }
]
```

**Error Responses:**
- `400 Bad Request`: Missing user_id query parameter
- `500 Internal Server Error`: Database error

**Example:**
```bash
curl "http://localhost:8080/api/payments?user_id=123e4567-e89b-12d3-a456-426614174000"
```

---

## Fare Calculation

### Calculate Fare
Calculates the fare for a trip based on distance and vehicle type.

**Endpoint:** `POST /api/calculate-fare`

**Request Body:**
```json
{
  "distance": 0.0,
  "vehicle_type": "string"
}
```

**Request Fields:**
- `distance` (float, required): Distance in kilometers
- `vehicle_type` (string, required): Type of vehicle ("standard" or "premium")

**Fare Calculation Logic:**
- Base fare: 500 XAF
- Standard rate: 200 XAF per km
- Premium rate: 300 XAF per km
- Formula: `fare = base_fare + (distance * per_km_rate)`

**Response:** `200 OK`
```json
{
  "distance": 0.0,
  "vehicle_type": "string",
  "fare": 0.0
}
```

**Error Responses:**
- `400 Bad Request`: Invalid request body

**Example:**
```bash
curl -X POST http://localhost:8080/api/calculate-fare \
  -H "Content-Type: application/json" \
  -d '{
    "distance": 10.5,
    "vehicle_type": "standard"
  }'
```

**Example Response:**
```json
{
  "distance": 10.5,
  "vehicle_type": "standard",
  "fare": 2600.0
}
```

---

## Error Response Format

All error responses follow this format:

```json
{
  "error": "error message description"
}
```

## Common HTTP Status Codes

- `200 OK`: Request succeeded
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request data
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

## Notes

- All IDs are UUIDs generated automatically by the server
- Timestamps are in ISO 8601 format
- All monetary amounts are in XAF (Central African Franc)
- Phone numbers should include country code (e.g., +237 for Cameroon)
