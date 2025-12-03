package main

import (
	"log"

	"net/http"

	"TP_Andreev/internal/config"
	"TP_Andreev/internal/db"
	"TP_Andreev/internal/db/migrations"
	"TP_Andreev/internal/repo/business_trip_repo"
	"TP_Andreev/internal/repo/employee_repo"
	"TP_Andreev/internal/service"
	"TP_Andreev/internal/transport/http/controller/employee_controller"
	"TP_Andreev/internal/transport/http/controller/main_controller"
	"TP_Andreev/internal/transport/http/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("congif load failed: %v", err)
	}

	db, err := db.Connect(&cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}

	service := service.New(
		employee_repo.New(db),
		business_trip_repo.New(db),
	)

	// Initialize router
	r := router.New()

	// Initialize controller
	pageCtrl := main_controller.New(*service)
	employeeCtrl := employee_controller.New(*service)

	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register routes
	r.GET("/", pageCtrl.GetMainPage)
	r.GET("/employee/:id", employeeCtrl.GetEmployee)

	// Start server with both router and static handler
	http.Handle("/", r)
	http.ListenAndServe(":"+cfg.Server.Port, nil)
}
