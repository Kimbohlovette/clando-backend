package models

import "time"

type Config struct {
	Port string `conf:"default:8080"`
	DB   struct {
		Host     string `conf:"env:DB_HOST,required"`
		Port     string `conf:"env:DB_PORT,required"`
		User     string `conf:"env:DB_USER,required"`
		Password string `conf:"env:DB_PASSWORD,required"`
		DBName   string `conf:"env:DB_NAME,required"`
		SSLMode  bool   `conf:"env:DB_SSL_MODE,default:false"`
	}
}

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payment struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	UserID        string    `json:"user_id"`
	TravelID      string    `json:"travel_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
}

type Driver struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	LicenseNo   string    `json:"license_no" gorm:"unique"`
	VehicleType string    `json:"vehicle_type"`
	VehicleNo   string    `json:"vehicle_no"`
	Rating      float64   `json:"rating"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TravelHistory struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	UserID     string    `json:"user_id"`
	DriverID   string    `json:"driver_id"`
	PickupLoc  string    `json:"pickup_loc"`
	DropoffLoc string    `json:"dropoff_loc"`
	Distance   float64   `json:"distance"`
	Fare       float64   `json:"fare"`
	Status     string    `json:"status"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	CreatedAt  time.Time `json:"created_at"`
}

type Place struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
}
