// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/models/operations"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathPutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfig(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "put_/v1/installations/{integrationConfigurationId}/resources/{resourceId}/experimentation/edge-config[0]":
			dir.HandlerFunc("put_/v1/installations/{integrationConfigurationId}/resources/{resourceId}/experimentation/edge-config", testPutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigPutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfig0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testPutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigPutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfig0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.ContentType(req, "application/json", false); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var respBody *operations.PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody = &operations.PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody{
		Items: map[string]*components.EdgeConfigItemValue{
			"key": types.Pointer(components.CreateEdgeConfigItemValueStr(
				"<value>",
			)),
		},
		UpdatedAt: 1217.32,
		Digest:    "<value>",
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}
