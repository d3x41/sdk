// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteProjectsProjectIDLogsPresetsIDRequest struct {
	// projectId of the preset
	ProjectID string `pathParam:"style=simple,explode=false,name=projectId"`
	// id of the preset
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteProjectsProjectIDLogsPresetsIDRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *DeleteProjectsProjectIDLogsPresetsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteProjectsProjectIDLogsPresetsIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteProjectsProjectIDLogsPresetsIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
