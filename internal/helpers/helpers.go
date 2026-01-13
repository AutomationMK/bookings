package helpers

import (
	"net/http"

	"github.com/AutomationMK/bookings/internal/config"
)

var app *config.AppConfig

// NewHelpers sets up appConfig for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {

}
