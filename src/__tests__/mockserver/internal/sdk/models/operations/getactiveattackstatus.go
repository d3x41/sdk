// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

type GetActiveAttackStatusRequest struct {
	ProjectID string `queryParam:"style=form,explode=true,name=projectId"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetActiveAttackStatusRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetActiveAttackStatusRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetActiveAttackStatusRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type AnomalyAlerts struct {
	AtMinute            string  `json:"at_minute"`
	Zscore              float64 `json:"zscore"`
	TotalRequestsMinute float64 `json:"total_requests_minute"`
	AvgRequests         float64 `json:"avg_requests"`
	StddevRequests      float64 `json:"stddev_requests"`
}

func (o *AnomalyAlerts) GetAtMinute() string {
	if o == nil {
		return ""
	}
	return o.AtMinute
}

func (o *AnomalyAlerts) GetZscore() float64 {
	if o == nil {
		return 0.0
	}
	return o.Zscore
}

func (o *AnomalyAlerts) GetTotalRequestsMinute() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalRequestsMinute
}

func (o *AnomalyAlerts) GetAvgRequests() float64 {
	if o == nil {
		return 0.0
	}
	return o.AvgRequests
}

func (o *AnomalyAlerts) GetStddevRequests() float64 {
	if o == nil {
		return 0.0
	}
	return o.StddevRequests
}

type DdosAlerts struct {
	AtMinute  string  `json:"atMinute"`
	TotalReqs float64 `json:"totalReqs"`
}

func (o *DdosAlerts) GetAtMinute() string {
	if o == nil {
		return ""
	}
	return o.AtMinute
}

func (o *DdosAlerts) GetTotalReqs() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalReqs
}

type AffectedHostMap struct {
	AnomalyAlerts map[string]AnomalyAlerts `json:"anomalyAlerts,omitempty"`
	DdosAlerts    map[string]DdosAlerts    `json:"ddosAlerts,omitempty"`
}

func (o *AffectedHostMap) GetAnomalyAlerts() map[string]AnomalyAlerts {
	if o == nil {
		return nil
	}
	return o.AnomalyAlerts
}

func (o *AffectedHostMap) GetDdosAlerts() map[string]DdosAlerts {
	if o == nil {
		return nil
	}
	return o.DdosAlerts
}

type Anomaly struct {
	OwnerID         string                     `json:"ownerId"`
	ProjectID       string                     `json:"projectId"`
	StartTime       float64                    `json:"startTime"`
	EndTime         *float64                   `json:"endTime"`
	AtMinute        float64                    `json:"atMinute"`
	State           *string                    `json:"state,omitempty"`
	AffectedHostMap map[string]AffectedHostMap `json:"affectedHostMap"`
}

func (o *Anomaly) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *Anomaly) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *Anomaly) GetStartTime() float64 {
	if o == nil {
		return 0.0
	}
	return o.StartTime
}

func (o *Anomaly) GetEndTime() *float64 {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *Anomaly) GetAtMinute() float64 {
	if o == nil {
		return 0.0
	}
	return o.AtMinute
}

func (o *Anomaly) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *Anomaly) GetAffectedHostMap() map[string]AffectedHostMap {
	if o == nil {
		return map[string]AffectedHostMap{}
	}
	return o.AffectedHostMap
}

type GetActiveAttackStatusResponseBody2 struct {
	Anomalies []Anomaly `json:"anomalies"`
}

func (o *GetActiveAttackStatusResponseBody2) GetAnomalies() []Anomaly {
	if o == nil {
		return []Anomaly{}
	}
	return o.Anomalies
}

type GetActiveAttackStatusResponseBody1 struct {
}

type GetActiveAttackStatusResponseBodyType string

const (
	GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody1 GetActiveAttackStatusResponseBodyType = "getActiveAttackStatus_ResponseBody_1"
	GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody2 GetActiveAttackStatusResponseBodyType = "getActiveAttackStatus_ResponseBody_2"
)

type GetActiveAttackStatusResponseBody struct {
	GetActiveAttackStatusResponseBody1 *GetActiveAttackStatusResponseBody1 `queryParam:"inline"`
	GetActiveAttackStatusResponseBody2 *GetActiveAttackStatusResponseBody2 `queryParam:"inline"`

	Type GetActiveAttackStatusResponseBodyType
}

func CreateGetActiveAttackStatusResponseBodyGetActiveAttackStatusResponseBody1(getActiveAttackStatusResponseBody1 GetActiveAttackStatusResponseBody1) GetActiveAttackStatusResponseBody {
	typ := GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody1

	return GetActiveAttackStatusResponseBody{
		GetActiveAttackStatusResponseBody1: &getActiveAttackStatusResponseBody1,
		Type:                               typ,
	}
}

func CreateGetActiveAttackStatusResponseBodyGetActiveAttackStatusResponseBody2(getActiveAttackStatusResponseBody2 GetActiveAttackStatusResponseBody2) GetActiveAttackStatusResponseBody {
	typ := GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody2

	return GetActiveAttackStatusResponseBody{
		GetActiveAttackStatusResponseBody2: &getActiveAttackStatusResponseBody2,
		Type:                               typ,
	}
}

func (u *GetActiveAttackStatusResponseBody) UnmarshalJSON(data []byte) error {

	var getActiveAttackStatusResponseBody1 GetActiveAttackStatusResponseBody1 = GetActiveAttackStatusResponseBody1{}
	if err := utils.UnmarshalJSON(data, &getActiveAttackStatusResponseBody1, "", true, true); err == nil {
		u.GetActiveAttackStatusResponseBody1 = &getActiveAttackStatusResponseBody1
		u.Type = GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody1
		return nil
	}

	var getActiveAttackStatusResponseBody2 GetActiveAttackStatusResponseBody2 = GetActiveAttackStatusResponseBody2{}
	if err := utils.UnmarshalJSON(data, &getActiveAttackStatusResponseBody2, "", true, true); err == nil {
		u.GetActiveAttackStatusResponseBody2 = &getActiveAttackStatusResponseBody2
		u.Type = GetActiveAttackStatusResponseBodyTypeGetActiveAttackStatusResponseBody2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetActiveAttackStatusResponseBody", string(data))
}

func (u GetActiveAttackStatusResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetActiveAttackStatusResponseBody1 != nil {
		return utils.MarshalJSON(u.GetActiveAttackStatusResponseBody1, "", true)
	}

	if u.GetActiveAttackStatusResponseBody2 != nil {
		return utils.MarshalJSON(u.GetActiveAttackStatusResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type GetActiveAttackStatusResponseBody: all fields are null")
}

type GetActiveAttackStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	OneOf    *GetActiveAttackStatusResponseBody
}

func (o *GetActiveAttackStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetActiveAttackStatusResponse) GetOneOf() *GetActiveAttackStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
