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

func pathPatchV1DataCacheProjectsProjectID(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "updateProjectDataCache[0]":
			dir.HandlerFunc("updateProjectDataCache", testUpdateProjectDataCacheUpdateProjectDataCache0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testUpdateProjectDataCacheUpdateProjectDataCache0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.ContentType(req, "application/json", true); err != nil {
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
	respBody := &operations.UpdateProjectDataCacheResponseBody{
		AccountID: "<id>",
		Crons: &operations.Crons{
			EnabledAt:    7135.8,
			DisabledAt:   types.Float64(2989.85),
			UpdatedAt:    6226.23,
			DeploymentID: types.String("<id>"),
			Definitions: []operations.Definitions{
				operations.Definitions{
					Host:     "vercel.com",
					Path:     "/api/crons/sync-something?hello=world",
					Schedule: "0 0 * * *",
				},
				operations.Definitions{
					Host:     "vercel.com",
					Path:     "/api/crons/sync-something?hello=world",
					Schedule: "0 0 * * *",
				},
				operations.Definitions{
					Host:     "vercel.com",
					Path:     "/api/crons/sync-something?hello=world",
					Schedule: "0 0 * * *",
				},
			},
		},
		DirectoryListing: false,
		ID:               "<id>",
		LatestDeployments: []operations.LatestDeployments{
			operations.LatestDeployments{
				ID:        "<id>",
				CreatedAt: 5719.6,
				CreatedIn: "<value>",
				Creator: &operations.UpdateProjectDataCacheProjectsCreator{
					Email:    "Braeden15@gmail.com",
					UID:      "<id>",
					Username: "Teresa84",
				},
				DeploymentHostname:     "<value>",
				Name:                   "<value>",
				Plan:                   operations.UpdateProjectDataCacheProjectsPlanPro,
				PreviewCommentsEnabled: types.Bool(false),
				Private:                false,
				ReadyState:             operations.UpdateProjectDataCacheProjectsReadyStateBuilding,
				Type:                   operations.UpdateProjectDataCacheProjectsTypeLambdas,
				URL:                    "https://unknown-gift.biz",
				UserID:                 "<id>",
			},
			operations.LatestDeployments{
				ID:        "<id>",
				CreatedAt: 7644.5,
				CreatedIn: "<value>",
				Creator: &operations.UpdateProjectDataCacheProjectsCreator{
					Email:    "Erich.Mann@hotmail.com",
					UID:      "<id>",
					Username: "Genoveva89",
				},
				DeploymentHostname:     "<value>",
				Name:                   "<value>",
				Plan:                   operations.UpdateProjectDataCacheProjectsPlanHobby,
				PreviewCommentsEnabled: types.Bool(false),
				Private:                true,
				ReadyState:             operations.UpdateProjectDataCacheProjectsReadyStateError,
				Type:                   operations.UpdateProjectDataCacheProjectsTypeLambdas,
				URL:                    "https://slimy-tuba.name",
				UserID:                 "<id>",
			},
			operations.LatestDeployments{
				ID:        "<id>",
				CreatedAt: 2357.62,
				CreatedIn: "<value>",
				Creator: &operations.UpdateProjectDataCacheProjectsCreator{
					Email:    "Garret.Ferry81@gmail.com",
					UID:      "<id>",
					Username: "Vivianne.Gutkowski",
				},
				DeploymentHostname:     "<value>",
				Name:                   "<value>",
				Plan:                   operations.UpdateProjectDataCacheProjectsPlanPro,
				PreviewCommentsEnabled: types.Bool(false),
				Private:                true,
				ReadyState:             operations.UpdateProjectDataCacheProjectsReadyStateCanceled,
				Type:                   operations.UpdateProjectDataCacheProjectsTypeLambdas,
				URL:                    "https://astonishing-cinema.net/",
				UserID:                 "<id>",
			},
		},
		Name:        "<value>",
		NodeVersion: operations.UpdateProjectDataCacheNodeVersionTenDotX,
		ResourceConfig: operations.UpdateProjectDataCacheResourceConfig{
			FunctionDefaultRegions: []string{},
		},
		DefaultResourceConfig: operations.DefaultResourceConfig{
			FunctionDefaultRegions: []string{},
		},
		Targets: map[string]*operations.Targets{
			"key": &operations.Targets{
				ID:        "<id>",
				CreatedAt: 1048.68,
				CreatedIn: "<value>",
				Creator: &operations.UpdateProjectDataCacheCreator{
					Email:    "Lillie17@gmail.com",
					UID:      "<id>",
					Username: "Herminia_Schowalter50",
				},
				DeploymentHostname:     "<value>",
				Name:                   "<value>",
				Plan:                   operations.UpdateProjectDataCachePlanEnterprise,
				PreviewCommentsEnabled: types.Bool(false),
				Private:                false,
				ReadyState:             operations.UpdateProjectDataCacheReadyStateBuilding,
				Type:                   operations.UpdateProjectDataCacheProjectsResponseTypeLambdas,
				URL:                    "https://quintessential-bidet.com/",
				UserID:                 "<id>",
			},
		},
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
