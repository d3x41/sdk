// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
	"time"
)

// Period for the billing cycle.
type Period struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (p Period) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Period) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Period) GetStart() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Start
}

func (o *Period) GetEnd() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.End
}

type BillingItems struct {
	// Partner's billing plan ID.
	BillingPlanID string `json:"billingPlanId"`
	// Partner's resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	Start *time.Time `json:"start,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	End *time.Time `json:"end,omitempty"`
	// Line item name.
	Name string `json:"name"`
	// Line item details.
	Details *string `json:"details,omitempty"`
	// Price per unit.
	Price string `json:"price"`
	// Quantity of units.
	Quantity float64 `json:"quantity"`
	// Units of the quantity.
	Units string `json:"units"`
	// Total amount.
	Total string `json:"total"`
}

func (b BillingItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BillingItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BillingItems) GetBillingPlanID() string {
	if o == nil {
		return ""
	}
	return o.BillingPlanID
}

func (o *BillingItems) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *BillingItems) GetStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *BillingItems) GetEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *BillingItems) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *BillingItems) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *BillingItems) GetPrice() string {
	if o == nil {
		return ""
	}
	return o.Price
}

func (o *BillingItems) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *BillingItems) GetUnits() string {
	if o == nil {
		return ""
	}
	return o.Units
}

func (o *BillingItems) GetTotal() string {
	if o == nil {
		return ""
	}
	return o.Total
}

type Discounts struct {
	// Partner's billing plan ID.
	BillingPlanID string `json:"billingPlanId"`
	// Partner's resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	Start *time.Time `json:"start,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	End *time.Time `json:"end,omitempty"`
	// Discount name.
	Name string `json:"name"`
	// Discount details.
	Details *string `json:"details,omitempty"`
	// Discount amount.
	Amount string `json:"amount"`
}

func (d Discounts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Discounts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Discounts) GetBillingPlanID() string {
	if o == nil {
		return ""
	}
	return o.BillingPlanID
}

func (o *Discounts) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *Discounts) GetStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *Discounts) GetEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *Discounts) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Discounts) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Discounts) GetAmount() string {
	if o == nil {
		return ""
	}
	return o.Amount
}

type Billing2 struct {
	Items     []BillingItems `json:"items"`
	Discounts []Discounts    `json:"discounts,omitempty"`
}

func (o *Billing2) GetItems() []BillingItems {
	if o == nil {
		return []BillingItems{}
	}
	return o.Items
}

func (o *Billing2) GetDiscounts() []Discounts {
	if o == nil {
		return nil
	}
	return o.Discounts
}

type Billing1 struct {
	// Partner's billing plan ID.
	BillingPlanID string `json:"billingPlanId"`
	// Partner's resource ID.
	ResourceID *string `json:"resourceId,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	Start *time.Time `json:"start,omitempty"`
	// Start and end are only needed if different from the period's start/end.
	End *time.Time `json:"end,omitempty"`
	// Line item name.
	Name string `json:"name"`
	// Line item details.
	Details *string `json:"details,omitempty"`
	// Price per unit.
	Price string `json:"price"`
	// Quantity of units.
	Quantity float64 `json:"quantity"`
	// Units of the quantity.
	Units string `json:"units"`
	// Total amount.
	Total string `json:"total"`
}

func (b Billing1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Billing1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Billing1) GetBillingPlanID() string {
	if o == nil {
		return ""
	}
	return o.BillingPlanID
}

func (o *Billing1) GetResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceID
}

func (o *Billing1) GetStart() *time.Time {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *Billing1) GetEnd() *time.Time {
	if o == nil {
		return nil
	}
	return o.End
}

func (o *Billing1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Billing1) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Billing1) GetPrice() string {
	if o == nil {
		return ""
	}
	return o.Price
}

func (o *Billing1) GetQuantity() float64 {
	if o == nil {
		return 0.0
	}
	return o.Quantity
}

func (o *Billing1) GetUnits() string {
	if o == nil {
		return ""
	}
	return o.Units
}

func (o *Billing1) GetTotal() string {
	if o == nil {
		return ""
	}
	return o.Total
}

type BillingType string

const (
	BillingTypeArrayOfBilling1 BillingType = "arrayOfBilling1"
	BillingTypeBilling2        BillingType = "billing_2"
)

// Billing data (interim invoicing data).
type Billing struct {
	ArrayOfBilling1 []Billing1
	Billing2        *Billing2

	Type BillingType
}

func CreateBillingArrayOfBilling1(arrayOfBilling1 []Billing1) Billing {
	typ := BillingTypeArrayOfBilling1

	return Billing{
		ArrayOfBilling1: arrayOfBilling1,
		Type:            typ,
	}
}

func CreateBillingBilling2(billing2 Billing2) Billing {
	typ := BillingTypeBilling2

	return Billing{
		Billing2: &billing2,
		Type:     typ,
	}
}

func (u *Billing) UnmarshalJSON(data []byte) error {

	var billing2 Billing2 = Billing2{}
	if err := utils.UnmarshalJSON(data, &billing2, "", true, true); err == nil {
		u.Billing2 = &billing2
		u.Type = BillingTypeBilling2
		return nil
	}

	var arrayOfBilling1 []Billing1 = []Billing1{}
	if err := utils.UnmarshalJSON(data, &arrayOfBilling1, "", true, true); err == nil {
		u.ArrayOfBilling1 = arrayOfBilling1
		u.Type = BillingTypeArrayOfBilling1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Billing", string(data))
}

func (u Billing) MarshalJSON() ([]byte, error) {
	if u.ArrayOfBilling1 != nil {
		return utils.MarshalJSON(u.ArrayOfBilling1, "", true)
	}

	if u.Billing2 != nil {
		return utils.MarshalJSON(u.Billing2, "", true)
	}

	return nil, errors.New("could not marshal union type Billing: all fields are null")
}

// SubmitBillingDataType - \n              Type of the metric.\n              - total: measured total value, such as Database size\n              - interval: usage during the period, such as i/o or number of queries.\n              - rate: rate of usage, such as queries per second.\n
type SubmitBillingDataType string

const (
	SubmitBillingDataTypeTotal    SubmitBillingDataType = "total"
	SubmitBillingDataTypeInterval SubmitBillingDataType = "interval"
	SubmitBillingDataTypeRate     SubmitBillingDataType = "rate"
)

func (e SubmitBillingDataType) ToPointer() *SubmitBillingDataType {
	return &e
}
func (e *SubmitBillingDataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total":
		fallthrough
	case "interval":
		fallthrough
	case "rate":
		*e = SubmitBillingDataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SubmitBillingDataType: %v", v)
	}
}

type Usage struct {
	// Partner's resource ID.
	ResourceID string `json:"resourceId"`
	// Metric name.
	Name string `json:"name"`
	// \n              Type of the metric.\n              - total: measured total value, such as Database size\n              - interval: usage during the period, such as i/o or number of queries.\n              - rate: rate of usage, such as queries per second.\n
	Type SubmitBillingDataType `json:"type"`
	// Metric units. Example: \"GB\"
	Units string `json:"units"`
	// Metric value for the day. Could be a final or an interim value for the day.
	DayValue float64 `json:"dayValue"`
	// Metric value for the billing period. Could be a final or an interim value for the period.
	PeriodValue float64 `json:"periodValue"`
	// The limit value of the metric for a billing period, if a limit is defined by the plan.
	PlanValue *float64 `json:"planValue,omitempty"`
}

func (o *Usage) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *Usage) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Usage) GetType() SubmitBillingDataType {
	if o == nil {
		return SubmitBillingDataType("")
	}
	return o.Type
}

func (o *Usage) GetUnits() string {
	if o == nil {
		return ""
	}
	return o.Units
}

func (o *Usage) GetDayValue() float64 {
	if o == nil {
		return 0.0
	}
	return o.DayValue
}

func (o *Usage) GetPeriodValue() float64 {
	if o == nil {
		return 0.0
	}
	return o.PeriodValue
}

func (o *Usage) GetPlanValue() *float64 {
	if o == nil {
		return nil
	}
	return o.PlanValue
}

type SubmitBillingDataRequestBody struct {
	Timestamp time.Time `json:"timestamp"`
	Eod       time.Time `json:"eod"`
	// Period for the billing cycle.
	Period Period `json:"period"`
	// Billing data (interim invoicing data).
	Billing Billing `json:"billing"`
	Usage   []Usage `json:"usage"`
}

func (s SubmitBillingDataRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SubmitBillingDataRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SubmitBillingDataRequestBody) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *SubmitBillingDataRequestBody) GetEod() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Eod
}

func (o *SubmitBillingDataRequestBody) GetPeriod() Period {
	if o == nil {
		return Period{}
	}
	return o.Period
}

func (o *SubmitBillingDataRequestBody) GetBilling() Billing {
	if o == nil {
		return Billing{}
	}
	return o.Billing
}

func (o *SubmitBillingDataRequestBody) GetUsage() []Usage {
	if o == nil {
		return []Usage{}
	}
	return o.Usage
}

type SubmitBillingDataRequest struct {
	IntegrationConfigurationID string                       `pathParam:"style=simple,explode=false,name=integrationConfigurationId"`
	RequestBody                SubmitBillingDataRequestBody `request:"mediaType=application/json"`
}

func (o *SubmitBillingDataRequest) GetIntegrationConfigurationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationConfigurationID
}

func (o *SubmitBillingDataRequest) GetRequestBody() SubmitBillingDataRequestBody {
	if o == nil {
		return SubmitBillingDataRequestBody{}
	}
	return o.RequestBody
}

type SubmitBillingDataResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *SubmitBillingDataResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
