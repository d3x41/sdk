// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/operations"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathGetV1SecurityFirewallAttackStatus(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "getActiveAttackStatus[0]":
			dir.HandlerFunc("getActiveAttackStatus", testGetActiveAttackStatusGetActiveAttackStatus0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testGetActiveAttackStatusGetActiveAttackStatus0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
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
	respBody := types.Pointer(operations.CreateGetActiveAttackStatusResponseBodyGetActiveAttackStatusResponseBody2(
		operations.GetActiveAttackStatusResponseBody2{
			Anomalies: []operations.Anomalies{
				operations.Anomalies{
					OwnerID:         "<id>",
					ProjectID:       "<id>",
					StartTime:       9556.58,
					EndTime:         types.Float64(3001.16),
					AtMinute:        5447.77,
					AffectedHostMap: map[string]operations.AffectedHostMap{},
				},
				operations.Anomalies{
					OwnerID:         "<id>",
					ProjectID:       "<id>",
					StartTime:       7786.06,
					EndTime:         types.Float64(9758.22),
					AtMinute:        7118.69,
					AffectedHostMap: map[string]operations.AffectedHostMap{},
				},
			},
		},
	))
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
