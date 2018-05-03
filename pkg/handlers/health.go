// Copyright 2017 Kubernetes Community Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"net/http"

	"github.com/k8s-community/codegen/pkg/router"
)

// Health returns "OK" if service is alive
func (h *Handler) Health(c router.Control) {
	c.Code(http.StatusOK)
	c.Body(http.StatusText(http.StatusOK))
}
