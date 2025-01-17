// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type DeleteAccessGroupProjectRequest struct {
	AccessGroupIDOrName string `pathParam:"style=simple,explode=false,name=accessGroupIdOrName"`
	ProjectID           string `pathParam:"style=simple,explode=false,name=projectId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *DeleteAccessGroupProjectRequest) GetAccessGroupIDOrName() string {
	if o == nil {
		return ""
	}
	return o.AccessGroupIDOrName
}

func (o *DeleteAccessGroupProjectRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *DeleteAccessGroupProjectRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *DeleteAccessGroupProjectRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type DeleteAccessGroupProjectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteAccessGroupProjectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
