package handler

import (
	"net/http"

	"github.com/user/cost-optimizer/internal/api"
)

// corsMiddleware allows browser clients to call API routes from the hosted UI.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handler is invoked by Vercel for API routes mapped to api/index.go.
func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/cost", api.GetCostHandler)
	mux.HandleFunc("/ec2", api.GetEC2InstancesHandler)
	mux.HandleFunc("/services", api.GetAllServicesHandler)
	mux.HandleFunc("/security", api.GetSecurityHandler)
	mux.HandleFunc("/security-details", api.GetSecurityDetailsHandler)

	corsMiddleware(mux).ServeHTTP(w, r)
}
