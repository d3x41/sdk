// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

// AddProjectDomainRedirectStatusCode - Status code for domain redirect
type AddProjectDomainRedirectStatusCode int64

const (
	AddProjectDomainRedirectStatusCodeThreeHundredAndOne   AddProjectDomainRedirectStatusCode = 301
	AddProjectDomainRedirectStatusCodeThreeHundredAndTwo   AddProjectDomainRedirectStatusCode = 302
	AddProjectDomainRedirectStatusCodeThreeHundredAndSeven AddProjectDomainRedirectStatusCode = 307
	AddProjectDomainRedirectStatusCodeThreeHundredAndEight AddProjectDomainRedirectStatusCode = 308
)

func (e AddProjectDomainRedirectStatusCode) ToPointer() *AddProjectDomainRedirectStatusCode {
	return &e
}
func (e *AddProjectDomainRedirectStatusCode) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 301:
		fallthrough
	case 302:
		fallthrough
	case 307:
		fallthrough
	case 308:
		*e = AddProjectDomainRedirectStatusCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AddProjectDomainRedirectStatusCode: %v", v)
	}
}

type AddProjectDomainRequestBody struct {
	// The project domain name
	Name string `json:"name"`
	// Git branch to link the project domain
	GitBranch *string `json:"gitBranch,omitempty"`
	// Target destination domain for redirect
	Redirect *string `json:"redirect,omitempty"`
	// Status code for domain redirect
	RedirectStatusCode *AddProjectDomainRedirectStatusCode `json:"redirectStatusCode,omitempty"`
}

func (o *AddProjectDomainRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddProjectDomainRequestBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *AddProjectDomainRequestBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *AddProjectDomainRequestBody) GetRedirectStatusCode() *AddProjectDomainRedirectStatusCode {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

type AddProjectDomainRequest struct {
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                     `queryParam:"style=form,explode=true,name=slug"`
	RequestBody AddProjectDomainRequestBody `request:"mediaType=application/json"`
}

func (o *AddProjectDomainRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *AddProjectDomainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *AddProjectDomainRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *AddProjectDomainRequest) GetRequestBody() AddProjectDomainRequestBody {
	if o == nil {
		return AddProjectDomainRequestBody{}
	}
	return o.RequestBody
}

// AddProjectDomainVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type AddProjectDomainVerification struct {
	Type   string `json:"type"`
	Domain string `json:"domain"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

func (o *AddProjectDomainVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AddProjectDomainVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *AddProjectDomainVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *AddProjectDomainVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

// AddProjectDomainResponseBody - The domain was successfully added to the project
type AddProjectDomainResponseBody struct {
	Name                string   `json:"name"`
	ApexName            string   `json:"apexName"`
	ProjectID           string   `json:"projectId"`
	Redirect            *string  `json:"redirect,omitempty"`
	RedirectStatusCode  *float64 `json:"redirectStatusCode,omitempty"`
	GitBranch           *string  `json:"gitBranch,omitempty"`
	CustomEnvironmentID *string  `json:"customEnvironmentId,omitempty"`
	UpdatedAt           *float64 `json:"updatedAt,omitempty"`
	CreatedAt           *float64 `json:"createdAt,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []AddProjectDomainVerification `json:"verification,omitempty"`
}

func (o *AddProjectDomainResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AddProjectDomainResponseBody) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *AddProjectDomainResponseBody) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *AddProjectDomainResponseBody) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *AddProjectDomainResponseBody) GetRedirectStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *AddProjectDomainResponseBody) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *AddProjectDomainResponseBody) GetCustomEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.CustomEnvironmentID
}

func (o *AddProjectDomainResponseBody) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AddProjectDomainResponseBody) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AddProjectDomainResponseBody) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

func (o *AddProjectDomainResponseBody) GetVerification() []AddProjectDomainVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

type AddProjectDomainResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The domain was successfully added to the project
	Object *AddProjectDomainResponseBody
}

func (o *AddProjectDomainResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *AddProjectDomainResponse) GetObject() *AddProjectDomainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
