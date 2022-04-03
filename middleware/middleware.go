package middleware

import (
	"myapp/data"

	"github.com/tsam3000/celeritas"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}