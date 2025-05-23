// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PatchEdgeConfigSchemaRequestBody struct {
	Definition any `json:"definition"`
}

func (o *PatchEdgeConfigSchemaRequestBody) GetDefinition() any {
	if o == nil {
		return nil
	}
	return o.Definition
}

type PatchEdgeConfigSchemaRequest struct {
	EdgeConfigID string  `pathParam:"style=simple,explode=false,name=edgeConfigId"`
	DryRun       *string `queryParam:"style=form,explode=true,name=dryRun"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                          `queryParam:"style=form,explode=true,name=slug"`
	RequestBody PatchEdgeConfigSchemaRequestBody `request:"mediaType=application/json"`
}

func (o *PatchEdgeConfigSchemaRequest) GetEdgeConfigID() string {
	if o == nil {
		return ""
	}
	return o.EdgeConfigID
}

func (o *PatchEdgeConfigSchemaRequest) GetDryRun() *string {
	if o == nil {
		return nil
	}
	return o.DryRun
}

func (o *PatchEdgeConfigSchemaRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *PatchEdgeConfigSchemaRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *PatchEdgeConfigSchemaRequest) GetRequestBody() PatchEdgeConfigSchemaRequestBody {
	if o == nil {
		return PatchEdgeConfigSchemaRequestBody{}
	}
	return o.RequestBody
}

// PatchEdgeConfigSchemaResponseBody - The JSON schema uploaded by the user
type PatchEdgeConfigSchemaResponseBody struct {
}

type PatchEdgeConfigSchemaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *PatchEdgeConfigSchemaResponseBody
}

func (o *PatchEdgeConfigSchemaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchEdgeConfigSchemaResponse) GetObject() *PatchEdgeConfigSchemaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
