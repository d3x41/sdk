// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/utils"
)

type ScopesType string

const (
	ScopesTypeTeam ScopesType = "team"
)

func (e ScopesType) ToPointer() *ScopesType {
	return &e
}
func (e *ScopesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "team":
		*e = ScopesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScopesType: %v", v)
	}
}

type AuthTokenScopesOrigin string

const (
	AuthTokenScopesOriginSaml      AuthTokenScopesOrigin = "saml"
	AuthTokenScopesOriginGithub    AuthTokenScopesOrigin = "github"
	AuthTokenScopesOriginGitlab    AuthTokenScopesOrigin = "gitlab"
	AuthTokenScopesOriginBitbucket AuthTokenScopesOrigin = "bitbucket"
	AuthTokenScopesOriginEmail     AuthTokenScopesOrigin = "email"
	AuthTokenScopesOriginManual    AuthTokenScopesOrigin = "manual"
	AuthTokenScopesOriginPasskey   AuthTokenScopesOrigin = "passkey"
	AuthTokenScopesOriginOtp       AuthTokenScopesOrigin = "otp"
	AuthTokenScopesOriginSms       AuthTokenScopesOrigin = "sms"
	AuthTokenScopesOriginInvite    AuthTokenScopesOrigin = "invite"
)

func (e AuthTokenScopesOrigin) ToPointer() *AuthTokenScopesOrigin {
	return &e
}
func (e *AuthTokenScopesOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saml":
		fallthrough
	case "github":
		fallthrough
	case "gitlab":
		fallthrough
	case "bitbucket":
		fallthrough
	case "email":
		fallthrough
	case "manual":
		fallthrough
	case "passkey":
		fallthrough
	case "otp":
		fallthrough
	case "sms":
		fallthrough
	case "invite":
		*e = AuthTokenScopesOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopesOrigin: %v", v)
	}
}

// Scopes2 - The access scopes granted to the token.
type Scopes2 struct {
	Type      ScopesType            `json:"type"`
	TeamID    string                `json:"teamId"`
	Origin    AuthTokenScopesOrigin `json:"origin"`
	CreatedAt float64               `json:"createdAt"`
	ExpiresAt *float64              `json:"expiresAt,omitempty"`
}

func (o *Scopes2) GetType() ScopesType {
	if o == nil {
		return ScopesType("")
	}
	return o.Type
}

func (o *Scopes2) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *Scopes2) GetOrigin() AuthTokenScopesOrigin {
	if o == nil {
		return AuthTokenScopesOrigin("")
	}
	return o.Origin
}

func (o *Scopes2) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *Scopes2) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

type AuthTokenScopesType string

const (
	AuthTokenScopesTypeUser AuthTokenScopesType = "user"
)

func (e AuthTokenScopesType) ToPointer() *AuthTokenScopesType {
	return &e
}
func (e *AuthTokenScopesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = AuthTokenScopesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthTokenScopesType: %v", v)
	}
}

type ScopesOrigin string

const (
	ScopesOriginSaml      ScopesOrigin = "saml"
	ScopesOriginGithub    ScopesOrigin = "github"
	ScopesOriginGitlab    ScopesOrigin = "gitlab"
	ScopesOriginBitbucket ScopesOrigin = "bitbucket"
	ScopesOriginEmail     ScopesOrigin = "email"
	ScopesOriginManual    ScopesOrigin = "manual"
	ScopesOriginPasskey   ScopesOrigin = "passkey"
	ScopesOriginOtp       ScopesOrigin = "otp"
	ScopesOriginSms       ScopesOrigin = "sms"
	ScopesOriginInvite    ScopesOrigin = "invite"
)

func (e ScopesOrigin) ToPointer() *ScopesOrigin {
	return &e
}
func (e *ScopesOrigin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "saml":
		fallthrough
	case "github":
		fallthrough
	case "gitlab":
		fallthrough
	case "bitbucket":
		fallthrough
	case "email":
		fallthrough
	case "manual":
		fallthrough
	case "passkey":
		fallthrough
	case "otp":
		fallthrough
	case "sms":
		fallthrough
	case "invite":
		*e = ScopesOrigin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScopesOrigin: %v", v)
	}
}

// Scopes1 - The access scopes granted to the token.
type Scopes1 struct {
	Type      AuthTokenScopesType `json:"type"`
	Origin    ScopesOrigin        `json:"origin"`
	CreatedAt float64             `json:"createdAt"`
	ExpiresAt *float64            `json:"expiresAt,omitempty"`
}

func (o *Scopes1) GetType() AuthTokenScopesType {
	if o == nil {
		return AuthTokenScopesType("")
	}
	return o.Type
}

func (o *Scopes1) GetOrigin() ScopesOrigin {
	if o == nil {
		return ScopesOrigin("")
	}
	return o.Origin
}

func (o *Scopes1) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *Scopes1) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

type ScopesUnionType string

const (
	ScopesUnionTypeScopes1 ScopesUnionType = "scopes_1"
	ScopesUnionTypeScopes2 ScopesUnionType = "scopes_2"
)

type Scopes struct {
	Scopes1 *Scopes1
	Scopes2 *Scopes2

	Type ScopesUnionType
}

func CreateScopesScopes1(scopes1 Scopes1) Scopes {
	typ := ScopesUnionTypeScopes1

	return Scopes{
		Scopes1: &scopes1,
		Type:    typ,
	}
}

func CreateScopesScopes2(scopes2 Scopes2) Scopes {
	typ := ScopesUnionTypeScopes2

	return Scopes{
		Scopes2: &scopes2,
		Type:    typ,
	}
}

func (u *Scopes) UnmarshalJSON(data []byte) error {

	var scopes1 Scopes1 = Scopes1{}
	if err := utils.UnmarshalJSON(data, &scopes1, "", true, true); err == nil {
		u.Scopes1 = &scopes1
		u.Type = ScopesUnionTypeScopes1
		return nil
	}

	var scopes2 Scopes2 = Scopes2{}
	if err := utils.UnmarshalJSON(data, &scopes2, "", true, true); err == nil {
		u.Scopes2 = &scopes2
		u.Type = ScopesUnionTypeScopes2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Scopes", string(data))
}

func (u Scopes) MarshalJSON() ([]byte, error) {
	if u.Scopes1 != nil {
		return utils.MarshalJSON(u.Scopes1, "", true)
	}

	if u.Scopes2 != nil {
		return utils.MarshalJSON(u.Scopes2, "", true)
	}

	return nil, errors.New("could not marshal union type Scopes: all fields are null")
}

// AuthToken - Authentication token metadata.
type AuthToken struct {
	// The unique identifier of the token.
	ID string `json:"id"`
	// The human-readable name of the token.
	Name string `json:"name"`
	// The type of the token.
	Type string `json:"type"`
	// The origin of how the token was created.
	Origin *string `json:"origin,omitempty"`
	// The access scopes granted to the token.
	Scopes []Scopes `json:"scopes,omitempty"`
	// Timestamp (in milliseconds) of when the token expires.
	ExpiresAt *float64 `json:"expiresAt,omitempty"`
	// Timestamp (in milliseconds) of when the token was most recently used.
	ActiveAt float64 `json:"activeAt"`
	// Timestamp (in milliseconds) of when the token was created.
	CreatedAt float64 `json:"createdAt"`
}

func (o *AuthToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthToken) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AuthToken) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *AuthToken) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *AuthToken) GetScopes() []Scopes {
	if o == nil {
		return nil
	}
	return o.Scopes
}

func (o *AuthToken) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *AuthToken) GetActiveAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ActiveAt
}

func (o *AuthToken) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}
