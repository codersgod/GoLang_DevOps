package handler

import (
	"net/http"

	"github.com/user/cost-optimizer/app"
)

// Handler is invoked by Vercel for API routes mapped to api/index.go.
func Handler(w http.ResponseWriter, r *http.Request) {
	app.APIHandler().ServeHTTP(w, r)
}
