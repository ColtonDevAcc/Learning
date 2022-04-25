package handlers

import (
	"net/http"

	voo "github.com/VooDooStack/Voo"
)

type Handlers struct {
	App *voo.Voo
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
