// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

type GetWebhookRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetWebhookRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetWebhookRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetWebhookRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// GetWebhookEvents - The webhooks events
type GetWebhookEvents string

const (
	GetWebhookEventsBudgetReached                                      GetWebhookEvents = "budget.reached"
	GetWebhookEventsBudgetReset                                        GetWebhookEvents = "budget.reset"
	GetWebhookEventsDomainCreated                                      GetWebhookEvents = "domain.created"
	GetWebhookEventsDeploymentCreated                                  GetWebhookEvents = "deployment.created"
	GetWebhookEventsDeploymentError                                    GetWebhookEvents = "deployment.error"
	GetWebhookEventsDeploymentCanceled                                 GetWebhookEvents = "deployment.canceled"
	GetWebhookEventsDeploymentSucceeded                                GetWebhookEvents = "deployment.succeeded"
	GetWebhookEventsDeploymentReady                                    GetWebhookEvents = "deployment.ready"
	GetWebhookEventsDeploymentCheckRerequested                         GetWebhookEvents = "deployment.check-rerequested"
	GetWebhookEventsDeploymentPromoted                                 GetWebhookEvents = "deployment.promoted"
	GetWebhookEventsDeploymentIntegrationActionStart                   GetWebhookEvents = "deployment.integration.action.start"
	GetWebhookEventsDeploymentIntegrationActionCancel                  GetWebhookEvents = "deployment.integration.action.cancel"
	GetWebhookEventsDeploymentIntegrationActionCleanup                 GetWebhookEvents = "deployment.integration.action.cleanup"
	GetWebhookEventsEdgeConfigCreated                                  GetWebhookEvents = "edge-config.created"
	GetWebhookEventsEdgeConfigDeleted                                  GetWebhookEvents = "edge-config.deleted"
	GetWebhookEventsEdgeConfigItemsUpdated                             GetWebhookEvents = "edge-config.items.updated"
	GetWebhookEventsFirewallAttack                                     GetWebhookEvents = "firewall.attack"
	GetWebhookEventsIntegrationConfigurationPermissionUpgraded         GetWebhookEvents = "integration-configuration.permission-upgraded"
	GetWebhookEventsIntegrationConfigurationRemoved                    GetWebhookEvents = "integration-configuration.removed"
	GetWebhookEventsIntegrationConfigurationScopeChangeConfirmed       GetWebhookEvents = "integration-configuration.scope-change-confirmed"
	GetWebhookEventsIntegrationResourceProjectConnected                GetWebhookEvents = "integration-resource.project-connected"
	GetWebhookEventsIntegrationResourceProjectDisconnected             GetWebhookEvents = "integration-resource.project-disconnected"
	GetWebhookEventsProjectCreated                                     GetWebhookEvents = "project.created"
	GetWebhookEventsProjectRemoved                                     GetWebhookEvents = "project.removed"
	GetWebhookEventsProjectDomainVerified                              GetWebhookEvents = "project.domain.verified"
	GetWebhookEventsProjectRollingReleaseStarted                       GetWebhookEvents = "project.rolling-release.started"
	GetWebhookEventsProjectRollingReleaseAborted                       GetWebhookEvents = "project.rolling-release.aborted"
	GetWebhookEventsProjectRollingReleaseCompleted                     GetWebhookEvents = "project.rolling-release.completed"
	GetWebhookEventsProjectRollingReleaseApproved                      GetWebhookEvents = "project.rolling-release.approved"
	GetWebhookEventsDeploymentChecksCompletedLegacy                    GetWebhookEvents = "deployment-checks-completed"
	GetWebhookEventsDeploymentReadyLegacy                              GetWebhookEvents = "deployment-ready"
	GetWebhookEventsDeploymentPreparedLegacy                           GetWebhookEvents = "deployment-prepared"
	GetWebhookEventsDeploymentErrorLegacy                              GetWebhookEvents = "deployment-error"
	GetWebhookEventsDeploymentCheckRerequestedLegacy                   GetWebhookEvents = "deployment-check-rerequested"
	GetWebhookEventsDeploymentCanceledLegacy                           GetWebhookEvents = "deployment-canceled"
	GetWebhookEventsProjectCreatedLegacy                               GetWebhookEvents = "project-created"
	GetWebhookEventsProjectRemovedLegacy                               GetWebhookEvents = "project-removed"
	GetWebhookEventsDomainCreatedLegacy                                GetWebhookEvents = "domain-created"
	GetWebhookEventsDeploymentLegacy                                   GetWebhookEvents = "deployment"
	GetWebhookEventsIntegrationConfigurationPermissionUpdatedLegacy    GetWebhookEvents = "integration-configuration-permission-updated"
	GetWebhookEventsIntegrationConfigurationRemovedLegacy              GetWebhookEvents = "integration-configuration-removed"
	GetWebhookEventsIntegrationConfigurationScopeChangeConfirmedLegacy GetWebhookEvents = "integration-configuration-scope-change-confirmed"
	GetWebhookEventsMarketplaceInvoiceCreated                          GetWebhookEvents = "marketplace.invoice.created"
	GetWebhookEventsMarketplaceInvoicePaid                             GetWebhookEvents = "marketplace.invoice.paid"
	GetWebhookEventsMarketplaceInvoiceNotpaid                          GetWebhookEvents = "marketplace.invoice.notpaid"
	GetWebhookEventsMarketplaceInvoiceRefunded                         GetWebhookEvents = "marketplace.invoice.refunded"
	GetWebhookEventsObservabilityAnomaly                               GetWebhookEvents = "observability.anomaly"
	GetWebhookEventsTestWebhook                                        GetWebhookEvents = "test-webhook"
)

func (e GetWebhookEvents) ToPointer() *GetWebhookEvents {
	return &e
}
func (e *GetWebhookEvents) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "budget.reached":
		fallthrough
	case "budget.reset":
		fallthrough
	case "domain.created":
		fallthrough
	case "deployment.created":
		fallthrough
	case "deployment.error":
		fallthrough
	case "deployment.canceled":
		fallthrough
	case "deployment.succeeded":
		fallthrough
	case "deployment.ready":
		fallthrough
	case "deployment.check-rerequested":
		fallthrough
	case "deployment.promoted":
		fallthrough
	case "deployment.integration.action.start":
		fallthrough
	case "deployment.integration.action.cancel":
		fallthrough
	case "deployment.integration.action.cleanup":
		fallthrough
	case "edge-config.created":
		fallthrough
	case "edge-config.deleted":
		fallthrough
	case "edge-config.items.updated":
		fallthrough
	case "firewall.attack":
		fallthrough
	case "integration-configuration.permission-upgraded":
		fallthrough
	case "integration-configuration.removed":
		fallthrough
	case "integration-configuration.scope-change-confirmed":
		fallthrough
	case "integration-resource.project-connected":
		fallthrough
	case "integration-resource.project-disconnected":
		fallthrough
	case "project.created":
		fallthrough
	case "project.removed":
		fallthrough
	case "project.domain.verified":
		fallthrough
	case "project.rolling-release.started":
		fallthrough
	case "project.rolling-release.aborted":
		fallthrough
	case "project.rolling-release.completed":
		fallthrough
	case "project.rolling-release.approved":
		fallthrough
	case "deployment-checks-completed":
		fallthrough
	case "deployment-ready":
		fallthrough
	case "deployment-prepared":
		fallthrough
	case "deployment-error":
		fallthrough
	case "deployment-check-rerequested":
		fallthrough
	case "deployment-canceled":
		fallthrough
	case "project-created":
		fallthrough
	case "project-removed":
		fallthrough
	case "domain-created":
		fallthrough
	case "deployment":
		fallthrough
	case "integration-configuration-permission-updated":
		fallthrough
	case "integration-configuration-removed":
		fallthrough
	case "integration-configuration-scope-change-confirmed":
		fallthrough
	case "marketplace.invoice.created":
		fallthrough
	case "marketplace.invoice.paid":
		fallthrough
	case "marketplace.invoice.notpaid":
		fallthrough
	case "marketplace.invoice.refunded":
		fallthrough
	case "observability.anomaly":
		fallthrough
	case "test-webhook":
		*e = GetWebhookEvents(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetWebhookEvents: %v", v)
	}
}

type GetWebhookResponseBody struct {
	// The webhooks events
	Events []GetWebhookEvents `json:"events"`
	// The webhook id
	ID string `json:"id"`
	// A string with the URL of the webhook
	URL string `json:"url"`
	// The unique ID of the team the webhook belongs to
	OwnerID string `json:"ownerId"`
	// A number containing the date when the webhook was created in in milliseconds
	CreatedAt float64 `json:"createdAt"`
	// A number containing the date when the webhook was updated in in milliseconds
	UpdatedAt float64 `json:"updatedAt"`
	// The ID of the projects the webhook is associated with
	ProjectIds []string `json:"projectIds,omitempty"`
}

func (o *GetWebhookResponseBody) GetEvents() []GetWebhookEvents {
	if o == nil {
		return []GetWebhookEvents{}
	}
	return o.Events
}

func (o *GetWebhookResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetWebhookResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *GetWebhookResponseBody) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetWebhookResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *GetWebhookResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *GetWebhookResponseBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

type GetWebhookResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *GetWebhookResponseBody
}

func (o *GetWebhookResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWebhookResponse) GetObject() *GetWebhookResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
