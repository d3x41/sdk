// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type IssueCertRequestBody struct {
	// The common names the cert should be issued for
	Cns []string `json:"cns,omitempty"`
}

func (o *IssueCertRequestBody) GetCns() []string {
	if o == nil {
		return nil
	}
	return o.Cns
}

type IssueCertRequest struct {
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string              `queryParam:"style=form,explode=true,name=slug"`
	RequestBody IssueCertRequestBody `request:"mediaType=application/json"`
}

func (o *IssueCertRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *IssueCertRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *IssueCertRequest) GetRequestBody() IssueCertRequestBody {
	if o == nil {
		return IssueCertRequestBody{}
	}
	return o.RequestBody
}

type IssueCertResponseBody struct {
	ID        string   `json:"id"`
	CreatedAt float64  `json:"createdAt"`
	ExpiresAt float64  `json:"expiresAt"`
	AutoRenew bool     `json:"autoRenew"`
	Cns       []string `json:"cns"`
}

func (o *IssueCertResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *IssueCertResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *IssueCertResponseBody) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

func (o *IssueCertResponseBody) GetAutoRenew() bool {
	if o == nil {
		return false
	}
	return o.AutoRenew
}

func (o *IssueCertResponseBody) GetCns() []string {
	if o == nil {
		return []string{}
	}
	return o.Cns
}

type IssueCertResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *IssueCertResponseBody
}

func (o *IssueCertResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *IssueCertResponse) GetObject() *IssueCertResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
