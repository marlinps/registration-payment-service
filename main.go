package main

import (
	"github.com/marlinps/registration-payment-service/api/routes"
	"github.com/marlinps/registration-payment-service/pkg/entities"
	"github.com/marlinps/registration-payment-service/pkg/registration"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// TODO: 1. Konfigurasi Database dan inisialisasi repository serta service
	dsn := "root:@tcp(127.0.0.1:3306)/sandbox_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal Terhubung ke Database")
	}
	db.AutoMigrate(&entities.Registration{})

	// TODO: 2. Inisialisasi Repository dan Service
	registrationRepo := registration.NewRegistrationRepository(db)
	registrationService := registration.NewRegistrationService(registrationRepo)

	// TODO: 4. Inisialisasi Fiber dan Setup Routes
	app := fiber.New()

	api := app.Group("/api/v1")
	routes.RegistrationRoutes(api, registrationService)
	app.Listen(":8080")
}
