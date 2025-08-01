// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

// UpdateResourceSecretsByIDEnvironmentOverrides - A map of environments to override values for the secret, used for setting different values across deployments in production, preview, and development environments. Note: the same value will be used for all deployments in the given environment.
type UpdateResourceSecretsByIDEnvironmentOverrides struct {
	// Value used for development environment.
	Development *string `json:"development,omitempty"`
	// Value used for preview environment.
	Preview *string `json:"preview,omitempty"`
	// Value used for production environment.
	Production *string `json:"production,omitempty"`
}

func (o *UpdateResourceSecretsByIDEnvironmentOverrides) GetDevelopment() *string {
	if o == nil {
		return nil
	}
	return o.Development
}

func (o *UpdateResourceSecretsByIDEnvironmentOverrides) GetPreview() *string {
	if o == nil {
		return nil
	}
	return o.Preview
}

func (o *UpdateResourceSecretsByIDEnvironmentOverrides) GetProduction() *string {
	if o == nil {
		return nil
	}
	return o.Production
}

type UpdateResourceSecretsByIDSecret struct {
	Name   string  `json:"name"`
	Value  string  `json:"value"`
	Prefix *string `json:"prefix,omitempty"`
	// A map of environments to override values for the secret, used for setting different values across deployments in production, preview, and development environments. Note: the same value will be used for all deployments in the given environment.
	EnvironmentOverrides *UpdateResourceSecretsByIDEnvironmentOverrides `json:"environmentOverrides,omitempty"`
}

func (o *UpdateResourceSecretsByIDSecret) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateResourceSecretsByIDSecret) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *UpdateResourceSecretsByIDSecret) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *UpdateResourceSecretsByIDSecret) GetEnvironmentOverrides() *UpdateResourceSecretsByIDEnvironmentOverrides {
	if o == nil {
		return nil
	}
	return o.EnvironmentOverrides
}

type UpdateResourceSecretsByIDRequestBody struct {
	Secrets []UpdateResourceSecretsByIDSecret `json:"secrets"`
	// If true, will only overwrite the provided secrets instead of replacing all secrets.
	Partial *bool `json:"partial,omitempty"`
}

func (o *UpdateResourceSecretsByIDRequestBody) GetSecrets() []UpdateResourceSecretsByIDSecret {
	if o == nil {
		return []UpdateResourceSecretsByIDSecret{}
	}
	return o.Secrets
}

func (o *UpdateResourceSecretsByIDRequestBody) GetPartial() *bool {
	if o == nil {
		return nil
	}
	return o.Partial
}

type UpdateResourceSecretsByIDRequest struct {
	IntegrationConfigurationID string                                `pathParam:"style=simple,explode=false,name=integrationConfigurationId"`
	ResourceID                 string                                `pathParam:"style=simple,explode=false,name=resourceId"`
	RequestBody                *UpdateResourceSecretsByIDRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateResourceSecretsByIDRequest) GetIntegrationConfigurationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationConfigurationID
}

func (o *UpdateResourceSecretsByIDRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *UpdateResourceSecretsByIDRequest) GetRequestBody() *UpdateResourceSecretsByIDRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type UpdateResourceSecretsByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateResourceSecretsByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
