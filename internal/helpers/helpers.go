package helpers

import "github.com/AutomationMK/bookings/internal/config"

var app *config.AppConfig

// NewHelpers sets up appConfig for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}
