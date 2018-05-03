package handlers

import (
	"net/http"

	"github.com/k8s-community/codegen/pkg/router"
)

// Static handles static files like css, javascript, img files
func (h *Handler) Static(c router.Control) {
	http.ServeFile(c, c.Request(), c.Request().URL.Path[1:])
}
