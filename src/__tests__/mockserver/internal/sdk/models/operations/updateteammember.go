// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

// UpdateTeamMemberRole - The project role of the member that will be added. \"null\" will remove this project level role.
type UpdateTeamMemberRole string

const (
	UpdateTeamMemberRoleAdmin            UpdateTeamMemberRole = "ADMIN"
	UpdateTeamMemberRoleProjectViewer    UpdateTeamMemberRole = "PROJECT_VIEWER"
	UpdateTeamMemberRoleProjectDeveloper UpdateTeamMemberRole = "PROJECT_DEVELOPER"
)

func (e UpdateTeamMemberRole) ToPointer() *UpdateTeamMemberRole {
	return &e
}
func (e *UpdateTeamMemberRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN":
		fallthrough
	case "PROJECT_VIEWER":
		fallthrough
	case "PROJECT_DEVELOPER":
		*e = UpdateTeamMemberRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateTeamMemberRole: %v", v)
	}
}

type UpdateTeamMemberProjects struct {
	// The ID of the project.
	ProjectID string `json:"projectId"`
	// The project role of the member that will be added. \"null\" will remove this project level role.
	Role *UpdateTeamMemberRole `json:"role"`
}

func (o *UpdateTeamMemberProjects) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *UpdateTeamMemberProjects) GetRole() *UpdateTeamMemberRole {
	if o == nil {
		return nil
	}
	return o.Role
}

type UpdateTeamMemberJoinedFrom struct {
	SsoUserID any `json:"ssoUserId,omitempty"`
}

func (o *UpdateTeamMemberJoinedFrom) GetSsoUserID() any {
	if o == nil {
		return nil
	}
	return o.SsoUserID
}

type UpdateTeamMemberRequestBody struct {
	// Accept a user who requested access to the team.
	Confirmed *bool `json:"confirmed,omitempty"`
	// The role in the team of the member.
	Role       *string                     `default:"MEMBER" json:"role"`
	Projects   []UpdateTeamMemberProjects  `json:"projects,omitempty"`
	JoinedFrom *UpdateTeamMemberJoinedFrom `json:"joinedFrom,omitempty"`
}

func (u UpdateTeamMemberRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateTeamMemberRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateTeamMemberRequestBody) GetConfirmed() *bool {
	if o == nil {
		return nil
	}
	return o.Confirmed
}

func (o *UpdateTeamMemberRequestBody) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *UpdateTeamMemberRequestBody) GetProjects() []UpdateTeamMemberProjects {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *UpdateTeamMemberRequestBody) GetJoinedFrom() *UpdateTeamMemberJoinedFrom {
	if o == nil {
		return nil
	}
	return o.JoinedFrom
}

type UpdateTeamMemberRequest struct {
	// The ID of the member.
	UID         string                      `pathParam:"style=simple,explode=false,name=uid"`
	TeamID      string                      `pathParam:"style=simple,explode=false,name=teamId"`
	RequestBody UpdateTeamMemberRequestBody `request:"mediaType=application/json"`
}

func (o *UpdateTeamMemberRequest) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *UpdateTeamMemberRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *UpdateTeamMemberRequest) GetRequestBody() UpdateTeamMemberRequestBody {
	if o == nil {
		return UpdateTeamMemberRequestBody{}
	}
	return o.RequestBody
}

// UpdateTeamMemberResponseBody - Successfully updated the membership.
type UpdateTeamMemberResponseBody struct {
	// ID of the team.
	ID string `json:"id"`
}

func (o *UpdateTeamMemberResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type UpdateTeamMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successfully updated the membership.
	Object *UpdateTeamMemberResponseBody
}

func (o *UpdateTeamMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateTeamMemberResponse) GetObject() *UpdateTeamMemberResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
