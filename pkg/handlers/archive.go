package handlers

import (
	"net/http"

	"github.com/k8s-community/codegen/pkg/router"
)

// Archive returns archives with generated services
func (h *Handler) Archive(c router.Control) {
	h.logger.Infof("Trying to get archive by path: %s", "/tmp"+c.Request().URL.Path)

	http.ServeFile(c, c.Request(), "/tmp"+c.Request().URL.Path)
}
