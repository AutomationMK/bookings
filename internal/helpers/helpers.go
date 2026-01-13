package helpers

import "github.com/AutomationMK/bookings/internal/config"

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}
