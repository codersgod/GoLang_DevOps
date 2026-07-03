package api

import (
	"encoding/json"
	"net/http"

	"github.com/user/cost-optimizer/internal/service"
)

const jsonContentType = "application/json"
const contentTypeHeader = "Content-Type"

func GetCostHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetCostData()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	json.NewEncoder(w).Encode(data)
}

func GetEC2InstancesHandler(w http.ResponseWriter, r *http.Request) {
	instances, count, err := service.GetEC2InstancesWithCPU()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"TotalCount": count,
		"Instances":  instances,
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	json.NewEncoder(w).Encode(response)
}

func GetAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	data, err := service.GetAllServices()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	json.NewEncoder(w).Encode(data)
}

func GetSecurityHandler(w http.ResponseWriter, r *http.Request) {
	sgCount, err := service.GetSecurityGroupsCount()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	kpCount, err := service.GetKeyPairsCount()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"SecurityGroups": sgCount,
		"KeyPairs":       kpCount,
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	json.NewEncoder(w).Encode(response)
}

func GetSecurityDetailsHandler(w http.ResponseWriter, r *http.Request) {
	sgs, err := service.GetSecurityGroupsDetails()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	kps, err := service.GetKeyPairsDetails()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := map[string]interface{}{
		"SecurityGroupsCount": len(sgs),
		"SecurityGroups":      sgs,
		"KeyPairsCount":       len(kps),
		"KeyPairs":            kps,
	}

	w.Header().Set(contentTypeHeader, jsonContentType)
	json.NewEncoder(w).Encode(response)
}