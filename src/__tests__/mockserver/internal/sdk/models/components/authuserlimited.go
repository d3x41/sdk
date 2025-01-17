// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// AuthUserLimitedVersion - The user's version. Will always be `northstar`.
type AuthUserLimitedVersion string

const (
	AuthUserLimitedVersionNorthstar AuthUserLimitedVersion = "northstar"
)

func (e AuthUserLimitedVersion) ToPointer() *AuthUserLimitedVersion {
	return &e
}
func (e *AuthUserLimitedVersion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "northstar":
		*e = AuthUserLimitedVersion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthUserLimitedVersion: %v", v)
	}
}

// AuthUserLimited - A limited form of data for the currently authenticated User, due to the authentication token missing privileges to read the full User data.
type AuthUserLimited struct {
	// Property indicating that this User data contains only limited information, due to the authentication token missing privileges to read the full User data. Re-login with email, GitHub, GitLab or Bitbucket in order to upgrade the authentication token with the necessary privileges.
	Limited bool `json:"limited"`
	// The User's unique identifier.
	ID string `json:"id"`
	// Email address associated with the User account.
	Email string `json:"email"`
	// Name associated with the User account, or `null` if none has been provided.
	Name *string `json:"name"`
	// Unique username associated with the User account.
	Username string `json:"username"`
	// SHA1 hash of the avatar for the User account. Can be used in conjuction with the ... endpoint to retrieve the avatar image.
	Avatar *string `json:"avatar"`
	// The user's default team.
	DefaultTeamID *string `json:"defaultTeamId"`
	// The user's version. Will always be `northstar`.
	Version AuthUserLimitedVersion `json:"version"`
}

func (o *AuthUserLimited) GetLimited() bool {
	if o == nil {
		return false
	}
	return o.Limited
}

func (o *AuthUserLimited) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AuthUserLimited) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *AuthUserLimited) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AuthUserLimited) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *AuthUserLimited) GetAvatar() *string {
	if o == nil {
		return nil
	}
	return o.Avatar
}

func (o *AuthUserLimited) GetDefaultTeamID() *string {
	if o == nil {
		return nil
	}
	return o.DefaultTeamID
}

func (o *AuthUserLimited) GetVersion() AuthUserLimitedVersion {
	if o == nil {
		return AuthUserLimitedVersion("")
	}
	return o.Version
}
