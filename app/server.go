package app

import (
	"net/http"

	internalapi "github.com/user/cost-optimizer/internal/api"
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

// APIHandler returns the HTTP handler for all API endpoints.
func APIHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/cost", internalapi.GetCostHandler)
	mux.HandleFunc("/ec2", internalapi.GetEC2InstancesHandler)
	mux.HandleFunc("/services", internalapi.GetAllServicesHandler)
	mux.HandleFunc("/security", internalapi.GetSecurityHandler)
	mux.HandleFunc("/security-details", internalapi.GetSecurityDetailsHandler)

	return corsMiddleware(mux)
}

// LocalHandler returns API + static web content for local runtime.
func LocalHandler() http.Handler {
	apiHandler := APIHandler()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web")))
	mux.Handle("/cost", apiHandler)
	mux.Handle("/ec2", apiHandler)
	mux.Handle("/services", apiHandler)
	mux.Handle("/security", apiHandler)
	mux.Handle("/security-details", apiHandler)

	return mux
}
