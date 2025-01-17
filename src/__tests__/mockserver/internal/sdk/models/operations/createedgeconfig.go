// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

type Two5 struct {
}

type Items2Type string

const (
	Items2TypeStr     Items2Type = "str"
	Items2TypeNumber  Items2Type = "number"
	Items2TypeBoolean Items2Type = "boolean"
	Items2TypeAny     Items2Type = "any"
	Items2TypeTwo5    Items2Type = "2_5"
)

type Items2 struct {
	Str     *string
	Number  *float64
	Boolean *bool
	Any     any
	Two5    *Two5

	Type Items2Type
}

func CreateItems2Str(str string) Items2 {
	typ := Items2TypeStr

	return Items2{
		Str:  &str,
		Type: typ,
	}
}

func CreateItems2Number(number float64) Items2 {
	typ := Items2TypeNumber

	return Items2{
		Number: &number,
		Type:   typ,
	}
}

func CreateItems2Boolean(boolean bool) Items2 {
	typ := Items2TypeBoolean

	return Items2{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateItems2Any(any any) Items2 {
	typ := Items2TypeAny

	return Items2{
		Any:  any,
		Type: typ,
	}
}

func CreateItems2Two5(two5 Two5) Items2 {
	typ := Items2TypeTwo5

	return Items2{
		Two5: &two5,
		Type: typ,
	}
}

func (u *Items2) UnmarshalJSON(data []byte) error {

	var two5 Two5 = Two5{}
	if err := utils.UnmarshalJSON(data, &two5, "", true, true); err == nil {
		u.Two5 = &two5
		u.Type = Items2TypeTwo5
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = Items2TypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = Items2TypeNumber
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = Items2TypeBoolean
		return nil
	}

	var any any = nil
	if err := utils.UnmarshalJSON(data, &any, "", true, true); err == nil {
		u.Any = any
		u.Type = Items2TypeAny
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Items2", string(data))
}

func (u Items2) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Any != nil {
		return utils.MarshalJSON(u.Any, "", true)
	}

	if u.Two5 != nil {
		return utils.MarshalJSON(u.Two5, "", true)
	}

	return nil, errors.New("could not marshal union type Items2: all fields are null")
}

type One5 struct {
}

type Items1Type string

const (
	Items1TypeStr     Items1Type = "str"
	Items1TypeNumber  Items1Type = "number"
	Items1TypeBoolean Items1Type = "boolean"
	Items1TypeAny     Items1Type = "any"
	Items1TypeOne5    Items1Type = "1_5"
)

type Items1 struct {
	Str     *string
	Number  *float64
	Boolean *bool
	Any     any
	One5    *One5

	Type Items1Type
}

func CreateItems1Str(str string) Items1 {
	typ := Items1TypeStr

	return Items1{
		Str:  &str,
		Type: typ,
	}
}

func CreateItems1Number(number float64) Items1 {
	typ := Items1TypeNumber

	return Items1{
		Number: &number,
		Type:   typ,
	}
}

func CreateItems1Boolean(boolean bool) Items1 {
	typ := Items1TypeBoolean

	return Items1{
		Boolean: &boolean,
		Type:    typ,
	}
}

func CreateItems1Any(any any) Items1 {
	typ := Items1TypeAny

	return Items1{
		Any:  any,
		Type: typ,
	}
}

func CreateItems1One5(one5 One5) Items1 {
	typ := Items1TypeOne5

	return Items1{
		One5: &one5,
		Type: typ,
	}
}

func (u *Items1) UnmarshalJSON(data []byte) error {

	var one5 One5 = One5{}
	if err := utils.UnmarshalJSON(data, &one5, "", true, true); err == nil {
		u.One5 = &one5
		u.Type = Items1TypeOne5
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = Items1TypeStr
		return nil
	}

	var number float64 = float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = Items1TypeNumber
		return nil
	}

	var boolean bool = false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = Items1TypeBoolean
		return nil
	}

	var any any = nil
	if err := utils.UnmarshalJSON(data, &any, "", true, true); err == nil {
		u.Any = any
		u.Type = Items1TypeAny
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Items1", string(data))
}

func (u Items1) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	if u.Any != nil {
		return utils.MarshalJSON(u.Any, "", true)
	}

	if u.One5 != nil {
		return utils.MarshalJSON(u.One5, "", true)
	}

	return nil, errors.New("could not marshal union type Items1: all fields are null")
}

type ItemsType string

const (
	ItemsTypeItems1        ItemsType = "items_1"
	ItemsTypeArrayOfItems2 ItemsType = "arrayOfItems2"
)

type Items struct {
	Items1        *Items1
	ArrayOfItems2 []Items2

	Type ItemsType
}

func CreateItemsItems1(items1 Items1) Items {
	typ := ItemsTypeItems1

	return Items{
		Items1: &items1,
		Type:   typ,
	}
}

func CreateItemsArrayOfItems2(arrayOfItems2 []Items2) Items {
	typ := ItemsTypeArrayOfItems2

	return Items{
		ArrayOfItems2: arrayOfItems2,
		Type:          typ,
	}
}

func (u *Items) UnmarshalJSON(data []byte) error {

	var items1 Items1 = Items1{}
	if err := utils.UnmarshalJSON(data, &items1, "", true, true); err == nil {
		u.Items1 = &items1
		u.Type = ItemsTypeItems1
		return nil
	}

	var arrayOfItems2 []Items2 = []Items2{}
	if err := utils.UnmarshalJSON(data, &arrayOfItems2, "", true, true); err == nil {
		u.ArrayOfItems2 = arrayOfItems2
		u.Type = ItemsTypeArrayOfItems2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Items", string(data))
}

func (u Items) MarshalJSON() ([]byte, error) {
	if u.Items1 != nil {
		return utils.MarshalJSON(u.Items1, "", true)
	}

	if u.ArrayOfItems2 != nil {
		return utils.MarshalJSON(u.ArrayOfItems2, "", true)
	}

	return nil, errors.New("could not marshal union type Items: all fields are null")
}

type CreateEdgeConfigRequestBody struct {
	Slug  string           `json:"slug"`
	Items map[string]Items `json:"items,omitempty"`
}

func (o *CreateEdgeConfigRequestBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateEdgeConfigRequestBody) GetItems() map[string]Items {
	if o == nil {
		return nil
	}
	return o.Items
}

type CreateEdgeConfigRequest struct {
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                     `queryParam:"style=form,explode=true,name=slug"`
	RequestBody CreateEdgeConfigRequestBody `request:"mediaType=application/json"`
}

func (o *CreateEdgeConfigRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateEdgeConfigRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateEdgeConfigRequest) GetRequestBody() CreateEdgeConfigRequestBody {
	if o == nil {
		return CreateEdgeConfigRequestBody{}
	}
	return o.RequestBody
}

// CreateEdgeConfigTransfer - Keeps track of the current state of the Edge Config while it gets transferred.
type CreateEdgeConfigTransfer struct {
	FromAccountID string   `json:"fromAccountId"`
	StartedAt     float64  `json:"startedAt"`
	DoneAt        *float64 `json:"doneAt"`
}

func (o *CreateEdgeConfigTransfer) GetFromAccountID() string {
	if o == nil {
		return ""
	}
	return o.FromAccountID
}

func (o *CreateEdgeConfigTransfer) GetStartedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.StartedAt
}

func (o *CreateEdgeConfigTransfer) GetDoneAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DoneAt
}

type CreateEdgeConfigSchema struct {
}

type CreateEdgeConfigType string

const (
	CreateEdgeConfigTypeFlags CreateEdgeConfigType = "flags"
)

func (e CreateEdgeConfigType) ToPointer() *CreateEdgeConfigType {
	return &e
}
func (e *CreateEdgeConfigType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "flags":
		*e = CreateEdgeConfigType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEdgeConfigType: %v", v)
	}
}

type CreateEdgeConfigPurpose struct {
	Type      CreateEdgeConfigType `json:"type"`
	ProjectID string               `json:"projectId"`
}

func (o *CreateEdgeConfigPurpose) GetType() CreateEdgeConfigType {
	if o == nil {
		return CreateEdgeConfigType("")
	}
	return o.Type
}

func (o *CreateEdgeConfigPurpose) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

// CreateEdgeConfigResponseBody - An Edge Config
type CreateEdgeConfigResponseBody struct {
	CreatedAt *float64 `json:"createdAt,omitempty"`
	UpdatedAt *float64 `json:"updatedAt,omitempty"`
	ID        *string  `json:"id,omitempty"`
	// Name for the Edge Config Names are not unique. Must start with an alphabetic character and can contain only alphanumeric characters and underscores).
	Slug    *string `json:"slug,omitempty"`
	OwnerID *string `json:"ownerId,omitempty"`
	Digest  *string `json:"digest,omitempty"`
	// Keeps track of the current state of the Edge Config while it gets transferred.
	Transfer    *CreateEdgeConfigTransfer `json:"transfer,omitempty"`
	Schema      *CreateEdgeConfigSchema   `json:"schema,omitempty"`
	Purpose     *CreateEdgeConfigPurpose  `json:"purpose,omitempty"`
	SizeInBytes float64                   `json:"sizeInBytes"`
	ItemCount   float64                   `json:"itemCount"`
}

func (o *CreateEdgeConfigResponseBody) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CreateEdgeConfigResponseBody) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CreateEdgeConfigResponseBody) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CreateEdgeConfigResponseBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateEdgeConfigResponseBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *CreateEdgeConfigResponseBody) GetDigest() *string {
	if o == nil {
		return nil
	}
	return o.Digest
}

func (o *CreateEdgeConfigResponseBody) GetTransfer() *CreateEdgeConfigTransfer {
	if o == nil {
		return nil
	}
	return o.Transfer
}

func (o *CreateEdgeConfigResponseBody) GetSchema() *CreateEdgeConfigSchema {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *CreateEdgeConfigResponseBody) GetPurpose() *CreateEdgeConfigPurpose {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *CreateEdgeConfigResponseBody) GetSizeInBytes() float64 {
	if o == nil {
		return 0.0
	}
	return o.SizeInBytes
}

func (o *CreateEdgeConfigResponseBody) GetItemCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.ItemCount
}

type CreateEdgeConfigResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateEdgeConfigResponseBody
}

func (o *CreateEdgeConfigResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateEdgeConfigResponse) GetObject() *CreateEdgeConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
