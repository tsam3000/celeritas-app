package handlers

import (
	"myapp/data"
	"net/http"


	"github.com/CloudyKit/jet/v6"
	"github.com/tsam3000/celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
