// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

type GetConfigurationRequest struct {
	// ID of the configuration to check
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetConfigurationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetConfigurationRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetConfigurationRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// ProjectSelection - A string representing the permission for projects. Possible values are `all` or `selected`.
type ProjectSelection string

const (
	ProjectSelectionSelected ProjectSelection = "selected"
	ProjectSelectionAll      ProjectSelection = "all"
)

func (e ProjectSelection) ToPointer() *ProjectSelection {
	return &e
}
func (e *ProjectSelection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "selected":
		fallthrough
	case "all":
		*e = ProjectSelection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProjectSelection: %v", v)
	}
}

type GetConfigurationLevel string

const (
	GetConfigurationLevelInfo  GetConfigurationLevel = "info"
	GetConfigurationLevelWarn  GetConfigurationLevel = "warn"
	GetConfigurationLevelError GetConfigurationLevel = "error"
)

func (e GetConfigurationLevel) ToPointer() *GetConfigurationLevel {
	return &e
}
func (e *GetConfigurationLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "info":
		fallthrough
	case "warn":
		fallthrough
	case "error":
		*e = GetConfigurationLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationLevel: %v", v)
	}
}

type GetConfigurationNotification struct {
	Level   GetConfigurationLevel `json:"level"`
	Title   string                `json:"title"`
	Message *string               `json:"message,omitempty"`
	Href    *string               `json:"href,omitempty"`
}

func (o *GetConfigurationNotification) GetLevel() GetConfigurationLevel {
	if o == nil {
		return GetConfigurationLevel("")
	}
	return o.Level
}

func (o *GetConfigurationNotification) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *GetConfigurationNotification) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetConfigurationNotification) GetHref() *string {
	if o == nil {
		return nil
	}
	return o.Href
}

type KindTransferFromMarketplace string

const (
	KindTransferFromMarketplaceTransferFromMarketplace KindTransferFromMarketplace = "transfer-from-marketplace"
)

func (e KindTransferFromMarketplace) ToPointer() *KindTransferFromMarketplace {
	return &e
}
func (e *KindTransferFromMarketplace) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "transfer-from-marketplace":
		*e = KindTransferFromMarketplace(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KindTransferFromMarketplace: %v", v)
	}
}

type Requester2 struct {
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
}

func (o *Requester2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Requester2) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

type TransferRequestTransferFromMarketplace struct {
	Kind            KindTransferFromMarketplace `json:"kind"`
	RequestID       string                      `json:"requestId"`
	TransferID      string                      `json:"transferId"`
	Requester       Requester2                  `json:"requester"`
	CreatedAt       float64                     `json:"createdAt"`
	ExpiresAt       float64                     `json:"expiresAt"`
	DiscardedAt     *float64                    `json:"discardedAt,omitempty"`
	DiscardedBy     *string                     `json:"discardedBy,omitempty"`
	ApprovedAt      *float64                    `json:"approvedAt,omitempty"`
	ApprovedBy      *string                     `json:"approvedBy,omitempty"`
	AuthorizationID *string                     `json:"authorizationId,omitempty"`
}

func (o *TransferRequestTransferFromMarketplace) GetKind() KindTransferFromMarketplace {
	if o == nil {
		return KindTransferFromMarketplace("")
	}
	return o.Kind
}

func (o *TransferRequestTransferFromMarketplace) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *TransferRequestTransferFromMarketplace) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

func (o *TransferRequestTransferFromMarketplace) GetRequester() Requester2 {
	if o == nil {
		return Requester2{}
	}
	return o.Requester
}

func (o *TransferRequestTransferFromMarketplace) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *TransferRequestTransferFromMarketplace) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

func (o *TransferRequestTransferFromMarketplace) GetDiscardedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscardedAt
}

func (o *TransferRequestTransferFromMarketplace) GetDiscardedBy() *string {
	if o == nil {
		return nil
	}
	return o.DiscardedBy
}

func (o *TransferRequestTransferFromMarketplace) GetApprovedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ApprovedAt
}

func (o *TransferRequestTransferFromMarketplace) GetApprovedBy() *string {
	if o == nil {
		return nil
	}
	return o.ApprovedBy
}

func (o *TransferRequestTransferFromMarketplace) GetAuthorizationID() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationID
}

type KindTransferToMarketplace string

const (
	KindTransferToMarketplaceTransferToMarketplace KindTransferToMarketplace = "transfer-to-marketplace"
)

func (e KindTransferToMarketplace) ToPointer() *KindTransferToMarketplace {
	return &e
}
func (e *KindTransferToMarketplace) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "transfer-to-marketplace":
		*e = KindTransferToMarketplace(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KindTransferToMarketplace: %v", v)
	}
}

type TransferRequestType string

const (
	TransferRequestTypeSubscription TransferRequestType = "subscription"
	TransferRequestTypePrepayment   TransferRequestType = "prepayment"
)

func (e TransferRequestType) ToPointer() *TransferRequestType {
	return &e
}
func (e *TransferRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "subscription":
		fallthrough
	case "prepayment":
		*e = TransferRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransferRequestType: %v", v)
	}
}

type GetConfigurationScope string

const (
	GetConfigurationScopeInstallation GetConfigurationScope = "installation"
	GetConfigurationScopeResource     GetConfigurationScope = "resource"
)

func (e GetConfigurationScope) ToPointer() *GetConfigurationScope {
	return &e
}
func (e *GetConfigurationScope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "installation":
		fallthrough
	case "resource":
		*e = GetConfigurationScope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationScope: %v", v)
	}
}

type GetConfigurationBillingPlan struct {
	ID                     string                 `json:"id"`
	Type                   TransferRequestType    `json:"type"`
	Scope                  *GetConfigurationScope `json:"scope,omitempty"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description"`
	PaymentMethodRequired  *bool                  `json:"paymentMethodRequired,omitempty"`
	PreauthorizationAmount *float64               `json:"preauthorizationAmount,omitempty"`
}

func (o *GetConfigurationBillingPlan) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetConfigurationBillingPlan) GetType() TransferRequestType {
	if o == nil {
		return TransferRequestType("")
	}
	return o.Type
}

func (o *GetConfigurationBillingPlan) GetScope() *GetConfigurationScope {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *GetConfigurationBillingPlan) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetConfigurationBillingPlan) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *GetConfigurationBillingPlan) GetPaymentMethodRequired() *bool {
	if o == nil {
		return nil
	}
	return o.PaymentMethodRequired
}

func (o *GetConfigurationBillingPlan) GetPreauthorizationAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.PreauthorizationAmount
}

type Requester1 struct {
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
}

func (o *Requester1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Requester1) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

type TransferRequestTransferToMarketplace struct {
	Kind            KindTransferToMarketplace    `json:"kind"`
	Metadata        map[string]any               `json:"metadata,omitempty"`
	BillingPlan     *GetConfigurationBillingPlan `json:"billingPlan,omitempty"`
	RequestID       string                       `json:"requestId"`
	TransferID      string                       `json:"transferId"`
	Requester       Requester1                   `json:"requester"`
	CreatedAt       float64                      `json:"createdAt"`
	ExpiresAt       float64                      `json:"expiresAt"`
	DiscardedAt     *float64                     `json:"discardedAt,omitempty"`
	DiscardedBy     *string                      `json:"discardedBy,omitempty"`
	ApprovedAt      *float64                     `json:"approvedAt,omitempty"`
	ApprovedBy      *string                      `json:"approvedBy,omitempty"`
	AuthorizationID *string                      `json:"authorizationId,omitempty"`
}

func (o *TransferRequestTransferToMarketplace) GetKind() KindTransferToMarketplace {
	if o == nil {
		return KindTransferToMarketplace("")
	}
	return o.Kind
}

func (o *TransferRequestTransferToMarketplace) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *TransferRequestTransferToMarketplace) GetBillingPlan() *GetConfigurationBillingPlan {
	if o == nil {
		return nil
	}
	return o.BillingPlan
}

func (o *TransferRequestTransferToMarketplace) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *TransferRequestTransferToMarketplace) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

func (o *TransferRequestTransferToMarketplace) GetRequester() Requester1 {
	if o == nil {
		return Requester1{}
	}
	return o.Requester
}

func (o *TransferRequestTransferToMarketplace) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *TransferRequestTransferToMarketplace) GetExpiresAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.ExpiresAt
}

func (o *TransferRequestTransferToMarketplace) GetDiscardedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscardedAt
}

func (o *TransferRequestTransferToMarketplace) GetDiscardedBy() *string {
	if o == nil {
		return nil
	}
	return o.DiscardedBy
}

func (o *TransferRequestTransferToMarketplace) GetApprovedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ApprovedAt
}

func (o *TransferRequestTransferToMarketplace) GetApprovedBy() *string {
	if o == nil {
		return nil
	}
	return o.ApprovedBy
}

func (o *TransferRequestTransferToMarketplace) GetAuthorizationID() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizationID
}

type TransferRequestUnionType string

const (
	TransferRequestUnionTypeTransferRequestTransferToMarketplace   TransferRequestUnionType = "transferRequest_TransferToMarketplace"
	TransferRequestUnionTypeTransferRequestTransferFromMarketplace TransferRequestUnionType = "transferRequest_TransferFromMarketplace"
)

type TransferRequest struct {
	TransferRequestTransferToMarketplace   *TransferRequestTransferToMarketplace   `queryParam:"inline"`
	TransferRequestTransferFromMarketplace *TransferRequestTransferFromMarketplace `queryParam:"inline"`

	Type TransferRequestUnionType
}

func CreateTransferRequestTransferRequestTransferToMarketplace(transferRequestTransferToMarketplace TransferRequestTransferToMarketplace) TransferRequest {
	typ := TransferRequestUnionTypeTransferRequestTransferToMarketplace

	return TransferRequest{
		TransferRequestTransferToMarketplace: &transferRequestTransferToMarketplace,
		Type:                                 typ,
	}
}

func CreateTransferRequestTransferRequestTransferFromMarketplace(transferRequestTransferFromMarketplace TransferRequestTransferFromMarketplace) TransferRequest {
	typ := TransferRequestUnionTypeTransferRequestTransferFromMarketplace

	return TransferRequest{
		TransferRequestTransferFromMarketplace: &transferRequestTransferFromMarketplace,
		Type:                                   typ,
	}
}

func (u *TransferRequest) UnmarshalJSON(data []byte) error {

	var transferRequestTransferFromMarketplace TransferRequestTransferFromMarketplace = TransferRequestTransferFromMarketplace{}
	if err := utils.UnmarshalJSON(data, &transferRequestTransferFromMarketplace, "", true, true); err == nil {
		u.TransferRequestTransferFromMarketplace = &transferRequestTransferFromMarketplace
		u.Type = TransferRequestUnionTypeTransferRequestTransferFromMarketplace
		return nil
	}

	var transferRequestTransferToMarketplace TransferRequestTransferToMarketplace = TransferRequestTransferToMarketplace{}
	if err := utils.UnmarshalJSON(data, &transferRequestTransferToMarketplace, "", true, true); err == nil {
		u.TransferRequestTransferToMarketplace = &transferRequestTransferToMarketplace
		u.Type = TransferRequestUnionTypeTransferRequestTransferToMarketplace
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TransferRequest", string(data))
}

func (u TransferRequest) MarshalJSON() ([]byte, error) {
	if u.TransferRequestTransferToMarketplace != nil {
		return utils.MarshalJSON(u.TransferRequestTransferToMarketplace, "", true)
	}

	if u.TransferRequestTransferFromMarketplace != nil {
		return utils.MarshalJSON(u.TransferRequestTransferFromMarketplace, "", true)
	}

	return nil, errors.New("could not marshal union type TransferRequest: all fields are null")
}

// GetConfigurationSource2 - Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
type GetConfigurationSource2 string

const (
	GetConfigurationSource2Marketplace  GetConfigurationSource2 = "marketplace"
	GetConfigurationSource2DeployButton GetConfigurationSource2 = "deploy-button"
	GetConfigurationSource2External     GetConfigurationSource2 = "external"
	GetConfigurationSource2V0           GetConfigurationSource2 = "v0"
)

func (e GetConfigurationSource2) ToPointer() *GetConfigurationSource2 {
	return &e
}
func (e *GetConfigurationSource2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "deploy-button":
		fallthrough
	case "external":
		fallthrough
	case "v0":
		*e = GetConfigurationSource2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationSource2: %v", v)
	}
}

type GetConfigurationTypeIntegrationConfiguration2 string

const (
	GetConfigurationTypeIntegrationConfiguration2IntegrationConfiguration GetConfigurationTypeIntegrationConfiguration2 = "integration-configuration"
)

func (e GetConfigurationTypeIntegrationConfiguration2) ToPointer() *GetConfigurationTypeIntegrationConfiguration2 {
	return &e
}
func (e *GetConfigurationTypeIntegrationConfiguration2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "integration-configuration":
		*e = GetConfigurationTypeIntegrationConfiguration2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationTypeIntegrationConfiguration2: %v", v)
	}
}

type GetConfigurationDisabledReason2 string

const (
	GetConfigurationDisabledReason2DisabledByOwner             GetConfigurationDisabledReason2 = "disabled-by-owner"
	GetConfigurationDisabledReason2FeatureNotAvailable         GetConfigurationDisabledReason2 = "feature-not-available"
	GetConfigurationDisabledReason2DisabledByAdmin             GetConfigurationDisabledReason2 = "disabled-by-admin"
	GetConfigurationDisabledReason2OriginalOwnerLeftTheTeam    GetConfigurationDisabledReason2 = "original-owner-left-the-team"
	GetConfigurationDisabledReason2AccountPlanDowngrade        GetConfigurationDisabledReason2 = "account-plan-downgrade"
	GetConfigurationDisabledReason2OriginalOwnerRoleDowngraded GetConfigurationDisabledReason2 = "original-owner-role-downgraded"
)

func (e GetConfigurationDisabledReason2) ToPointer() *GetConfigurationDisabledReason2 {
	return &e
}
func (e *GetConfigurationDisabledReason2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled-by-owner":
		fallthrough
	case "feature-not-available":
		fallthrough
	case "disabled-by-admin":
		fallthrough
	case "original-owner-left-the-team":
		fallthrough
	case "account-plan-downgrade":
		fallthrough
	case "original-owner-role-downgraded":
		*e = GetConfigurationDisabledReason2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationDisabledReason2: %v", v)
	}
}

// GetConfigurationInstallationType2 - Defines the installation type. - 'external' integrations are installed via the existing integrations flow - 'marketplace' integrations are natively installed: - when accepting the TOS of a partner during the store creation process - if undefined, assume 'external'
type GetConfigurationInstallationType2 string

const (
	GetConfigurationInstallationType2Marketplace GetConfigurationInstallationType2 = "marketplace"
	GetConfigurationInstallationType2External    GetConfigurationInstallationType2 = "external"
)

func (e GetConfigurationInstallationType2) ToPointer() *GetConfigurationInstallationType2 {
	return &e
}
func (e *GetConfigurationInstallationType2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "external":
		*e = GetConfigurationInstallationType2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationInstallationType2: %v", v)
	}
}

type GetConfigurationIntegrationConfiguration2 struct {
	// A string representing the permission for projects. Possible values are `all` or `selected`.
	ProjectSelection ProjectSelection             `json:"projectSelection"`
	Notification     GetConfigurationNotification `json:"notification"`
	TransferRequest  TransferRequest              `json:"transferRequest"`
	// When a configuration is limited to access certain projects, this will contain each of the project ID it is allowed to access. If it is not defined, the configuration has full access.
	Projects []string `json:"projects,omitempty"`
	// A timestamp that tells you when the configuration was installed successfully
	CompletedAt *float64 `json:"completedAt,omitempty"`
	// A timestamp that tells you when the configuration was created
	CreatedAt float64 `json:"createdAt"`
	// The unique identifier of the configuration
	ID string `json:"id"`
	// The unique identifier of the app the configuration was created for
	IntegrationID string `json:"integrationId"`
	// The user or team ID that owns the configuration
	OwnerID string `json:"ownerId"`
	// Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
	Source *GetConfigurationSource2 `json:"source,omitempty"`
	// The slug of the integration the configuration is created for.
	Slug string `json:"slug"`
	// When the configuration was created for a team, this will show the ID of the team.
	TeamID *string                                       `json:"teamId,omitempty"`
	Type   GetConfigurationTypeIntegrationConfiguration2 `json:"type"`
	// A timestamp that tells you when the configuration was updated.
	UpdatedAt float64 `json:"updatedAt"`
	// The ID of the user that created the configuration.
	UserID string `json:"userId"`
	// The resources that are allowed to be accessed by the configuration.
	Scopes []string `json:"scopes"`
	// A timestamp that tells you when the configuration was disabled. Note: Configurations can be disabled when the associated user loses access to a team. They do not function during this time until the configuration is 'transferred', meaning the associated user is changed to one with access to the team.
	DisabledAt *float64 `json:"disabledAt,omitempty"`
	// A timestamp that tells you when the configuration was deleted.
	DeletedAt *float64 `json:"deletedAt,omitempty"`
	// A timestamp that tells you when the configuration deletion has been started for cases when the deletion needs to be settled/approved by partners, such as when marketplace invoices have been paid.
	DeleteRequestedAt *float64                         `json:"deleteRequestedAt,omitempty"`
	DisabledReason    *GetConfigurationDisabledReason2 `json:"disabledReason,omitempty"`
	// Defines the installation type. - 'external' integrations are installed via the existing integrations flow - 'marketplace' integrations are natively installed: - when accepting the TOS of a partner during the store creation process - if undefined, assume 'external'
	InstallationType          *GetConfigurationInstallationType2 `json:"installationType,omitempty"`
	CanConfigureOpenTelemetry *bool                              `json:"canConfigureOpenTelemetry,omitempty"`
}

func (o *GetConfigurationIntegrationConfiguration2) GetProjectSelection() ProjectSelection {
	if o == nil {
		return ProjectSelection("")
	}
	return o.ProjectSelection
}

func (o *GetConfigurationIntegrationConfiguration2) GetNotification() GetConfigurationNotification {
	if o == nil {
		return GetConfigurationNotification{}
	}
	return o.Notification
}

func (o *GetConfigurationIntegrationConfiguration2) GetTransferRequest() TransferRequest {
	if o == nil {
		return TransferRequest{}
	}
	return o.TransferRequest
}

func (o *GetConfigurationIntegrationConfiguration2) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *GetConfigurationIntegrationConfiguration2) GetCompletedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetConfigurationIntegrationConfiguration2) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *GetConfigurationIntegrationConfiguration2) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetConfigurationIntegrationConfiguration2) GetSource() *GetConfigurationSource2 {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *GetConfigurationIntegrationConfiguration2) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *GetConfigurationIntegrationConfiguration2) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetConfigurationIntegrationConfiguration2) GetType() GetConfigurationTypeIntegrationConfiguration2 {
	if o == nil {
		return GetConfigurationTypeIntegrationConfiguration2("")
	}
	return o.Type
}

func (o *GetConfigurationIntegrationConfiguration2) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *GetConfigurationIntegrationConfiguration2) GetScopes() []string {
	if o == nil {
		return []string{}
	}
	return o.Scopes
}

func (o *GetConfigurationIntegrationConfiguration2) GetDisabledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetDeletedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetDeleteRequestedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeleteRequestedAt
}

func (o *GetConfigurationIntegrationConfiguration2) GetDisabledReason() *GetConfigurationDisabledReason2 {
	if o == nil {
		return nil
	}
	return o.DisabledReason
}

func (o *GetConfigurationIntegrationConfiguration2) GetInstallationType() *GetConfigurationInstallationType2 {
	if o == nil {
		return nil
	}
	return o.InstallationType
}

func (o *GetConfigurationIntegrationConfiguration2) GetCanConfigureOpenTelemetry() *bool {
	if o == nil {
		return nil
	}
	return o.CanConfigureOpenTelemetry
}

// GetConfigurationSource1 - Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
type GetConfigurationSource1 string

const (
	GetConfigurationSource1Marketplace  GetConfigurationSource1 = "marketplace"
	GetConfigurationSource1DeployButton GetConfigurationSource1 = "deploy-button"
	GetConfigurationSource1External     GetConfigurationSource1 = "external"
	GetConfigurationSource1V0           GetConfigurationSource1 = "v0"
)

func (e GetConfigurationSource1) ToPointer() *GetConfigurationSource1 {
	return &e
}
func (e *GetConfigurationSource1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "deploy-button":
		fallthrough
	case "external":
		fallthrough
	case "v0":
		*e = GetConfigurationSource1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationSource1: %v", v)
	}
}

type GetConfigurationTypeIntegrationConfiguration1 string

const (
	GetConfigurationTypeIntegrationConfiguration1IntegrationConfiguration GetConfigurationTypeIntegrationConfiguration1 = "integration-configuration"
)

func (e GetConfigurationTypeIntegrationConfiguration1) ToPointer() *GetConfigurationTypeIntegrationConfiguration1 {
	return &e
}
func (e *GetConfigurationTypeIntegrationConfiguration1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "integration-configuration":
		*e = GetConfigurationTypeIntegrationConfiguration1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationTypeIntegrationConfiguration1: %v", v)
	}
}

type GetConfigurationDisabledReason1 string

const (
	GetConfigurationDisabledReason1DisabledByOwner             GetConfigurationDisabledReason1 = "disabled-by-owner"
	GetConfigurationDisabledReason1FeatureNotAvailable         GetConfigurationDisabledReason1 = "feature-not-available"
	GetConfigurationDisabledReason1DisabledByAdmin             GetConfigurationDisabledReason1 = "disabled-by-admin"
	GetConfigurationDisabledReason1OriginalOwnerLeftTheTeam    GetConfigurationDisabledReason1 = "original-owner-left-the-team"
	GetConfigurationDisabledReason1AccountPlanDowngrade        GetConfigurationDisabledReason1 = "account-plan-downgrade"
	GetConfigurationDisabledReason1OriginalOwnerRoleDowngraded GetConfigurationDisabledReason1 = "original-owner-role-downgraded"
)

func (e GetConfigurationDisabledReason1) ToPointer() *GetConfigurationDisabledReason1 {
	return &e
}
func (e *GetConfigurationDisabledReason1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled-by-owner":
		fallthrough
	case "feature-not-available":
		fallthrough
	case "disabled-by-admin":
		fallthrough
	case "original-owner-left-the-team":
		fallthrough
	case "account-plan-downgrade":
		fallthrough
	case "original-owner-role-downgraded":
		*e = GetConfigurationDisabledReason1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationDisabledReason1: %v", v)
	}
}

// GetConfigurationInstallationType1 - Defines the installation type. - 'external' integrations are installed via the existing integrations flow - 'marketplace' integrations are natively installed: - when accepting the TOS of a partner during the store creation process - if undefined, assume 'external'
type GetConfigurationInstallationType1 string

const (
	GetConfigurationInstallationType1Marketplace GetConfigurationInstallationType1 = "marketplace"
	GetConfigurationInstallationType1External    GetConfigurationInstallationType1 = "external"
)

func (e GetConfigurationInstallationType1) ToPointer() *GetConfigurationInstallationType1 {
	return &e
}
func (e *GetConfigurationInstallationType1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "marketplace":
		fallthrough
	case "external":
		*e = GetConfigurationInstallationType1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetConfigurationInstallationType1: %v", v)
	}
}

// GetConfigurationIntegrationConfiguration1 - The configuration with the provided id
type GetConfigurationIntegrationConfiguration1 struct {
	// A timestamp that tells you when the configuration was installed successfully
	CompletedAt *float64 `json:"completedAt,omitempty"`
	// A timestamp that tells you when the configuration was created
	CreatedAt float64 `json:"createdAt"`
	// The unique identifier of the configuration
	ID string `json:"id"`
	// The unique identifier of the app the configuration was created for
	IntegrationID string `json:"integrationId"`
	// The user or team ID that owns the configuration
	OwnerID string `json:"ownerId"`
	// When a configuration is limited to access certain projects, this will contain each of the project ID it is allowed to access. If it is not defined, the configuration has full access.
	Projects []string `json:"projects,omitempty"`
	// Source defines where the configuration was installed from. It is used to analyze user engagement for integration installations in product metrics.
	Source *GetConfigurationSource1 `json:"source,omitempty"`
	// The slug of the integration the configuration is created for.
	Slug string `json:"slug"`
	// When the configuration was created for a team, this will show the ID of the team.
	TeamID *string                                       `json:"teamId,omitempty"`
	Type   GetConfigurationTypeIntegrationConfiguration1 `json:"type"`
	// A timestamp that tells you when the configuration was updated.
	UpdatedAt float64 `json:"updatedAt"`
	// The ID of the user that created the configuration.
	UserID string `json:"userId"`
	// The resources that are allowed to be accessed by the configuration.
	Scopes []string `json:"scopes"`
	// A timestamp that tells you when the configuration was disabled. Note: Configurations can be disabled when the associated user loses access to a team. They do not function during this time until the configuration is 'transferred', meaning the associated user is changed to one with access to the team.
	DisabledAt *float64 `json:"disabledAt,omitempty"`
	// A timestamp that tells you when the configuration was deleted.
	DeletedAt *float64 `json:"deletedAt,omitempty"`
	// A timestamp that tells you when the configuration deletion has been started for cases when the deletion needs to be settled/approved by partners, such as when marketplace invoices have been paid.
	DeleteRequestedAt *float64                         `json:"deleteRequestedAt,omitempty"`
	DisabledReason    *GetConfigurationDisabledReason1 `json:"disabledReason,omitempty"`
	// Defines the installation type. - 'external' integrations are installed via the existing integrations flow - 'marketplace' integrations are natively installed: - when accepting the TOS of a partner during the store creation process - if undefined, assume 'external'
	InstallationType *GetConfigurationInstallationType1 `json:"installationType,omitempty"`
}

func (o *GetConfigurationIntegrationConfiguration1) GetCompletedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetConfigurationIntegrationConfiguration1) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *GetConfigurationIntegrationConfiguration1) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetConfigurationIntegrationConfiguration1) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *GetConfigurationIntegrationConfiguration1) GetSource() *GetConfigurationSource1 {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *GetConfigurationIntegrationConfiguration1) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *GetConfigurationIntegrationConfiguration1) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetConfigurationIntegrationConfiguration1) GetType() GetConfigurationTypeIntegrationConfiguration1 {
	if o == nil {
		return GetConfigurationTypeIntegrationConfiguration1("")
	}
	return o.Type
}

func (o *GetConfigurationIntegrationConfiguration1) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

func (o *GetConfigurationIntegrationConfiguration1) GetScopes() []string {
	if o == nil {
		return []string{}
	}
	return o.Scopes
}

func (o *GetConfigurationIntegrationConfiguration1) GetDisabledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetDeletedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetDeleteRequestedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeleteRequestedAt
}

func (o *GetConfigurationIntegrationConfiguration1) GetDisabledReason() *GetConfigurationDisabledReason1 {
	if o == nil {
		return nil
	}
	return o.DisabledReason
}

func (o *GetConfigurationIntegrationConfiguration1) GetInstallationType() *GetConfigurationInstallationType1 {
	if o == nil {
		return nil
	}
	return o.InstallationType
}

type GetConfigurationResponseBodyType string

const (
	GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration1 GetConfigurationResponseBodyType = "getConfiguration_IntegrationConfiguration_1"
	GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration2 GetConfigurationResponseBodyType = "getConfiguration_IntegrationConfiguration_2"
)

// GetConfigurationResponseBody - The configuration with the provided id
type GetConfigurationResponseBody struct {
	GetConfigurationIntegrationConfiguration1 *GetConfigurationIntegrationConfiguration1 `queryParam:"inline"`
	GetConfigurationIntegrationConfiguration2 *GetConfigurationIntegrationConfiguration2 `queryParam:"inline"`

	Type GetConfigurationResponseBodyType
}

func CreateGetConfigurationResponseBodyGetConfigurationIntegrationConfiguration1(getConfigurationIntegrationConfiguration1 GetConfigurationIntegrationConfiguration1) GetConfigurationResponseBody {
	typ := GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration1

	return GetConfigurationResponseBody{
		GetConfigurationIntegrationConfiguration1: &getConfigurationIntegrationConfiguration1,
		Type: typ,
	}
}

func CreateGetConfigurationResponseBodyGetConfigurationIntegrationConfiguration2(getConfigurationIntegrationConfiguration2 GetConfigurationIntegrationConfiguration2) GetConfigurationResponseBody {
	typ := GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration2

	return GetConfigurationResponseBody{
		GetConfigurationIntegrationConfiguration2: &getConfigurationIntegrationConfiguration2,
		Type: typ,
	}
}

func (u *GetConfigurationResponseBody) UnmarshalJSON(data []byte) error {

	var getConfigurationIntegrationConfiguration1 GetConfigurationIntegrationConfiguration1 = GetConfigurationIntegrationConfiguration1{}
	if err := utils.UnmarshalJSON(data, &getConfigurationIntegrationConfiguration1, "", true, true); err == nil {
		u.GetConfigurationIntegrationConfiguration1 = &getConfigurationIntegrationConfiguration1
		u.Type = GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration1
		return nil
	}

	var getConfigurationIntegrationConfiguration2 GetConfigurationIntegrationConfiguration2 = GetConfigurationIntegrationConfiguration2{}
	if err := utils.UnmarshalJSON(data, &getConfigurationIntegrationConfiguration2, "", true, true); err == nil {
		u.GetConfigurationIntegrationConfiguration2 = &getConfigurationIntegrationConfiguration2
		u.Type = GetConfigurationResponseBodyTypeGetConfigurationIntegrationConfiguration2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetConfigurationResponseBody", string(data))
}

func (u GetConfigurationResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetConfigurationIntegrationConfiguration1 != nil {
		return utils.MarshalJSON(u.GetConfigurationIntegrationConfiguration1, "", true)
	}

	if u.GetConfigurationIntegrationConfiguration2 != nil {
		return utils.MarshalJSON(u.GetConfigurationIntegrationConfiguration2, "", true)
	}

	return nil, errors.New("could not marshal union type GetConfigurationResponseBody: all fields are null")
}

type GetConfigurationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configuration with the provided id
	OneOf *GetConfigurationResponseBody
}

func (o *GetConfigurationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetConfigurationResponse) GetOneOf() *GetConfigurationResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
