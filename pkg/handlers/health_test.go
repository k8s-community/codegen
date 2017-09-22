package handlers

import (
	"net/http"
	"testing"

	"github.com/k8s-community/codegen/pkg/config"
	"github.com/k8s-community/codegen/pkg/logger"
	"github.com/k8s-community/codegen/pkg/logger/standard"
	"github.com/k8s-community/codegen/pkg/router/bitroute"
)

func TestHealth(t *testing.T) {
	h := New(standard.New(&logger.Config{}), new(config.Config))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Base(h.Health)(bitroute.NewControl(w, r))
	})

	testHandler(t, handler, http.StatusOK, http.StatusText(http.StatusOK))
}
