// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type ArtifactExistsRequest struct {
	// The artifact hash
	Hash string `pathParam:"style=simple,explode=false,name=hash"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *ArtifactExistsRequest) GetHash() string {
	if o == nil {
		return ""
	}
	return o.Hash
}

func (o *ArtifactExistsRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *ArtifactExistsRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type ArtifactExistsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ArtifactExistsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
