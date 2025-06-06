// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

type Events string

const (
	EventsBudgetReached                                      Events = "budget.reached"
	EventsBudgetReset                                        Events = "budget.reset"
	EventsDomainCreated                                      Events = "domain.created"
	EventsDeploymentCreated                                  Events = "deployment.created"
	EventsDeploymentError                                    Events = "deployment.error"
	EventsDeploymentCanceled                                 Events = "deployment.canceled"
	EventsDeploymentSucceeded                                Events = "deployment.succeeded"
	EventsDeploymentReady                                    Events = "deployment.ready"
	EventsDeploymentCheckRerequested                         Events = "deployment.check-rerequested"
	EventsDeploymentPromoted                                 Events = "deployment.promoted"
	EventsDeploymentIntegrationActionStart                   Events = "deployment.integration.action.start"
	EventsDeploymentIntegrationActionCancel                  Events = "deployment.integration.action.cancel"
	EventsDeploymentIntegrationActionCleanup                 Events = "deployment.integration.action.cleanup"
	EventsEdgeConfigCreated                                  Events = "edge-config.created"
	EventsEdgeConfigDeleted                                  Events = "edge-config.deleted"
	EventsEdgeConfigItemsUpdated                             Events = "edge-config.items.updated"
	EventsFirewallAttack                                     Events = "firewall.attack"
	EventsIntegrationConfigurationPermissionUpgraded         Events = "integration-configuration.permission-upgraded"
	EventsIntegrationConfigurationRemoved                    Events = "integration-configuration.removed"
	EventsIntegrationConfigurationScopeChangeConfirmed       Events = "integration-configuration.scope-change-confirmed"
	EventsIntegrationResourceProjectConnected                Events = "integration-resource.project-connected"
	EventsIntegrationResourceProjectDisconnected             Events = "integration-resource.project-disconnected"
	EventsProjectCreated                                     Events = "project.created"
	EventsProjectRemoved                                     Events = "project.removed"
	EventsProjectDomainVerified                              Events = "project.domain.verified"
	EventsProjectRollingReleaseStarted                       Events = "project.rolling-release.started"
	EventsProjectRollingReleaseAborted                       Events = "project.rolling-release.aborted"
	EventsProjectRollingReleaseCompleted                     Events = "project.rolling-release.completed"
	EventsProjectRollingReleaseApproved                      Events = "project.rolling-release.approved"
	EventsDeploymentChecksCompletedLegacy                    Events = "deployment-checks-completed"
	EventsDeploymentReadyLegacy                              Events = "deployment-ready"
	EventsDeploymentPreparedLegacy                           Events = "deployment-prepared"
	EventsDeploymentErrorLegacy                              Events = "deployment-error"
	EventsDeploymentCheckRerequestedLegacy                   Events = "deployment-check-rerequested"
	EventsDeploymentCanceledLegacy                           Events = "deployment-canceled"
	EventsProjectCreatedLegacy                               Events = "project-created"
	EventsProjectRemovedLegacy                               Events = "project-removed"
	EventsDomainCreatedLegacy                                Events = "domain-created"
	EventsDeploymentLegacy                                   Events = "deployment"
	EventsIntegrationConfigurationPermissionUpdatedLegacy    Events = "integration-configuration-permission-updated"
	EventsIntegrationConfigurationRemovedLegacy              Events = "integration-configuration-removed"
	EventsIntegrationConfigurationScopeChangeConfirmedLegacy Events = "integration-configuration-scope-change-confirmed"
	EventsMarketplaceInvoiceCreated                          Events = "marketplace.invoice.created"
	EventsMarketplaceInvoicePaid                             Events = "marketplace.invoice.paid"
	EventsMarketplaceInvoiceNotpaid                          Events = "marketplace.invoice.notpaid"
	EventsMarketplaceInvoiceRefunded                         Events = "marketplace.invoice.refunded"
	EventsObservabilityAnomaly                               Events = "observability.anomaly"
	EventsTestWebhook                                        Events = "test-webhook"
)

func (e Events) ToPointer() *Events {
	return &e
}
func (e *Events) UnmarshalJSON(data []byte) error {
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
		*e = Events(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Events: %v", v)
	}
}

type CreateWebhookRequestBody struct {
	URL        string   `json:"url"`
	Events     []Events `json:"events"`
	ProjectIds []string `json:"projectIds,omitempty"`
}

func (o *CreateWebhookRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateWebhookRequestBody) GetEvents() []Events {
	if o == nil {
		return []Events{}
	}
	return o.Events
}

func (o *CreateWebhookRequestBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

type CreateWebhookRequest struct {
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                  `queryParam:"style=form,explode=true,name=slug"`
	RequestBody CreateWebhookRequestBody `request:"mediaType=application/json"`
}

func (o *CreateWebhookRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateWebhookRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateWebhookRequest) GetRequestBody() CreateWebhookRequestBody {
	if o == nil {
		return CreateWebhookRequestBody{}
	}
	return o.RequestBody
}

// CreateWebhookEvents - The webhooks events
type CreateWebhookEvents string

const (
	CreateWebhookEventsBudgetReached                                      CreateWebhookEvents = "budget.reached"
	CreateWebhookEventsBudgetReset                                        CreateWebhookEvents = "budget.reset"
	CreateWebhookEventsDomainCreated                                      CreateWebhookEvents = "domain.created"
	CreateWebhookEventsDeploymentCreated                                  CreateWebhookEvents = "deployment.created"
	CreateWebhookEventsDeploymentError                                    CreateWebhookEvents = "deployment.error"
	CreateWebhookEventsDeploymentCanceled                                 CreateWebhookEvents = "deployment.canceled"
	CreateWebhookEventsDeploymentSucceeded                                CreateWebhookEvents = "deployment.succeeded"
	CreateWebhookEventsDeploymentReady                                    CreateWebhookEvents = "deployment.ready"
	CreateWebhookEventsDeploymentCheckRerequested                         CreateWebhookEvents = "deployment.check-rerequested"
	CreateWebhookEventsDeploymentPromoted                                 CreateWebhookEvents = "deployment.promoted"
	CreateWebhookEventsDeploymentIntegrationActionStart                   CreateWebhookEvents = "deployment.integration.action.start"
	CreateWebhookEventsDeploymentIntegrationActionCancel                  CreateWebhookEvents = "deployment.integration.action.cancel"
	CreateWebhookEventsDeploymentIntegrationActionCleanup                 CreateWebhookEvents = "deployment.integration.action.cleanup"
	CreateWebhookEventsEdgeConfigCreated                                  CreateWebhookEvents = "edge-config.created"
	CreateWebhookEventsEdgeConfigDeleted                                  CreateWebhookEvents = "edge-config.deleted"
	CreateWebhookEventsEdgeConfigItemsUpdated                             CreateWebhookEvents = "edge-config.items.updated"
	CreateWebhookEventsFirewallAttack                                     CreateWebhookEvents = "firewall.attack"
	CreateWebhookEventsIntegrationConfigurationPermissionUpgraded         CreateWebhookEvents = "integration-configuration.permission-upgraded"
	CreateWebhookEventsIntegrationConfigurationRemoved                    CreateWebhookEvents = "integration-configuration.removed"
	CreateWebhookEventsIntegrationConfigurationScopeChangeConfirmed       CreateWebhookEvents = "integration-configuration.scope-change-confirmed"
	CreateWebhookEventsIntegrationResourceProjectConnected                CreateWebhookEvents = "integration-resource.project-connected"
	CreateWebhookEventsIntegrationResourceProjectDisconnected             CreateWebhookEvents = "integration-resource.project-disconnected"
	CreateWebhookEventsProjectCreated                                     CreateWebhookEvents = "project.created"
	CreateWebhookEventsProjectRemoved                                     CreateWebhookEvents = "project.removed"
	CreateWebhookEventsProjectDomainVerified                              CreateWebhookEvents = "project.domain.verified"
	CreateWebhookEventsProjectRollingReleaseStarted                       CreateWebhookEvents = "project.rolling-release.started"
	CreateWebhookEventsProjectRollingReleaseAborted                       CreateWebhookEvents = "project.rolling-release.aborted"
	CreateWebhookEventsProjectRollingReleaseCompleted                     CreateWebhookEvents = "project.rolling-release.completed"
	CreateWebhookEventsProjectRollingReleaseApproved                      CreateWebhookEvents = "project.rolling-release.approved"
	CreateWebhookEventsDeploymentChecksCompletedLegacy                    CreateWebhookEvents = "deployment-checks-completed"
	CreateWebhookEventsDeploymentReadyLegacy                              CreateWebhookEvents = "deployment-ready"
	CreateWebhookEventsDeploymentPreparedLegacy                           CreateWebhookEvents = "deployment-prepared"
	CreateWebhookEventsDeploymentErrorLegacy                              CreateWebhookEvents = "deployment-error"
	CreateWebhookEventsDeploymentCheckRerequestedLegacy                   CreateWebhookEvents = "deployment-check-rerequested"
	CreateWebhookEventsDeploymentCanceledLegacy                           CreateWebhookEvents = "deployment-canceled"
	CreateWebhookEventsProjectCreatedLegacy                               CreateWebhookEvents = "project-created"
	CreateWebhookEventsProjectRemovedLegacy                               CreateWebhookEvents = "project-removed"
	CreateWebhookEventsDomainCreatedLegacy                                CreateWebhookEvents = "domain-created"
	CreateWebhookEventsDeploymentLegacy                                   CreateWebhookEvents = "deployment"
	CreateWebhookEventsIntegrationConfigurationPermissionUpdatedLegacy    CreateWebhookEvents = "integration-configuration-permission-updated"
	CreateWebhookEventsIntegrationConfigurationRemovedLegacy              CreateWebhookEvents = "integration-configuration-removed"
	CreateWebhookEventsIntegrationConfigurationScopeChangeConfirmedLegacy CreateWebhookEvents = "integration-configuration-scope-change-confirmed"
	CreateWebhookEventsMarketplaceInvoiceCreated                          CreateWebhookEvents = "marketplace.invoice.created"
	CreateWebhookEventsMarketplaceInvoicePaid                             CreateWebhookEvents = "marketplace.invoice.paid"
	CreateWebhookEventsMarketplaceInvoiceNotpaid                          CreateWebhookEvents = "marketplace.invoice.notpaid"
	CreateWebhookEventsMarketplaceInvoiceRefunded                         CreateWebhookEvents = "marketplace.invoice.refunded"
	CreateWebhookEventsObservabilityAnomaly                               CreateWebhookEvents = "observability.anomaly"
	CreateWebhookEventsTestWebhook                                        CreateWebhookEvents = "test-webhook"
)

func (e CreateWebhookEvents) ToPointer() *CreateWebhookEvents {
	return &e
}
func (e *CreateWebhookEvents) UnmarshalJSON(data []byte) error {
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
		*e = CreateWebhookEvents(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebhookEvents: %v", v)
	}
}

type CreateWebhookResponseBody struct {
	// The webhook secret used to sign the payload
	Secret string `json:"secret"`
	// The webhooks events
	Events []CreateWebhookEvents `json:"events"`
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

func (o *CreateWebhookResponseBody) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

func (o *CreateWebhookResponseBody) GetEvents() []CreateWebhookEvents {
	if o == nil {
		return []CreateWebhookEvents{}
	}
	return o.Events
}

func (o *CreateWebhookResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateWebhookResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateWebhookResponseBody) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *CreateWebhookResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *CreateWebhookResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *CreateWebhookResponseBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

type CreateWebhookResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateWebhookResponseBody
}

func (o *CreateWebhookResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateWebhookResponse) GetObject() *CreateWebhookResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
