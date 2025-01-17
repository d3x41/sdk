// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

// CreateRecordRequestBodyDNSRequest10Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest10Type string

const (
	CreateRecordRequestBodyDNSRequest10TypeA     CreateRecordRequestBodyDNSRequest10Type = "A"
	CreateRecordRequestBodyDNSRequest10TypeAaaa  CreateRecordRequestBodyDNSRequest10Type = "AAAA"
	CreateRecordRequestBodyDNSRequest10TypeAlias CreateRecordRequestBodyDNSRequest10Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest10TypeCaa   CreateRecordRequestBodyDNSRequest10Type = "CAA"
	CreateRecordRequestBodyDNSRequest10TypeCname CreateRecordRequestBodyDNSRequest10Type = "CNAME"
	CreateRecordRequestBodyDNSRequest10TypeHTTPS CreateRecordRequestBodyDNSRequest10Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest10TypeMx    CreateRecordRequestBodyDNSRequest10Type = "MX"
	CreateRecordRequestBodyDNSRequest10TypeSrv   CreateRecordRequestBodyDNSRequest10Type = "SRV"
	CreateRecordRequestBodyDNSRequest10TypeTxt   CreateRecordRequestBodyDNSRequest10Type = "TXT"
	CreateRecordRequestBodyDNSRequest10TypeNs    CreateRecordRequestBodyDNSRequest10Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest10Type) ToPointer() *CreateRecordRequestBodyDNSRequest10Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest10Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest10Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest10Type: %v", v)
	}
}

type RequestBodyHTTPS struct {
	Priority *float64 `json:"priority"`
	Target   string   `json:"target"`
	Params   *string  `json:"params,omitempty"`
}

func (o *RequestBodyHTTPS) GetPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *RequestBodyHTTPS) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

func (o *RequestBodyHTTPS) GetParams() *string {
	if o == nil {
		return nil
	}
	return o.Params
}

type Ten struct {
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest10Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL   *float64         `json:"ttl,omitempty"`
	HTTPS RequestBodyHTTPS `json:"https"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *Ten) GetType() CreateRecordRequestBodyDNSRequest10Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest10Type("")
	}
	return o.Type
}

func (o *Ten) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Ten) GetHTTPS() RequestBodyHTTPS {
	if o == nil {
		return RequestBodyHTTPS{}
	}
	return o.HTTPS
}

func (o *Ten) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequest9Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest9Type string

const (
	CreateRecordRequestBodyDNSRequest9TypeA     CreateRecordRequestBodyDNSRequest9Type = "A"
	CreateRecordRequestBodyDNSRequest9TypeAaaa  CreateRecordRequestBodyDNSRequest9Type = "AAAA"
	CreateRecordRequestBodyDNSRequest9TypeAlias CreateRecordRequestBodyDNSRequest9Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest9TypeCaa   CreateRecordRequestBodyDNSRequest9Type = "CAA"
	CreateRecordRequestBodyDNSRequest9TypeCname CreateRecordRequestBodyDNSRequest9Type = "CNAME"
	CreateRecordRequestBodyDNSRequest9TypeHTTPS CreateRecordRequestBodyDNSRequest9Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest9TypeMx    CreateRecordRequestBodyDNSRequest9Type = "MX"
	CreateRecordRequestBodyDNSRequest9TypeSrv   CreateRecordRequestBodyDNSRequest9Type = "SRV"
	CreateRecordRequestBodyDNSRequest9TypeTxt   CreateRecordRequestBodyDNSRequest9Type = "TXT"
	CreateRecordRequestBodyDNSRequest9TypeNs    CreateRecordRequestBodyDNSRequest9Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest9Type) ToPointer() *CreateRecordRequestBodyDNSRequest9Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest9Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest9Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest9Type: %v", v)
	}
}

type Nine struct {
	// A subdomain name.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest9Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// An NS domain value.
	Value *string `json:"value,omitempty"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *Nine) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Nine) GetType() CreateRecordRequestBodyDNSRequest9Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest9Type("")
	}
	return o.Type
}

func (o *Nine) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Nine) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Nine) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequest8Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest8Type string

const (
	CreateRecordRequestBodyDNSRequest8TypeA     CreateRecordRequestBodyDNSRequest8Type = "A"
	CreateRecordRequestBodyDNSRequest8TypeAaaa  CreateRecordRequestBodyDNSRequest8Type = "AAAA"
	CreateRecordRequestBodyDNSRequest8TypeAlias CreateRecordRequestBodyDNSRequest8Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest8TypeCaa   CreateRecordRequestBodyDNSRequest8Type = "CAA"
	CreateRecordRequestBodyDNSRequest8TypeCname CreateRecordRequestBodyDNSRequest8Type = "CNAME"
	CreateRecordRequestBodyDNSRequest8TypeHTTPS CreateRecordRequestBodyDNSRequest8Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest8TypeMx    CreateRecordRequestBodyDNSRequest8Type = "MX"
	CreateRecordRequestBodyDNSRequest8TypeSrv   CreateRecordRequestBodyDNSRequest8Type = "SRV"
	CreateRecordRequestBodyDNSRequest8TypeTxt   CreateRecordRequestBodyDNSRequest8Type = "TXT"
	CreateRecordRequestBodyDNSRequest8TypeNs    CreateRecordRequestBodyDNSRequest8Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest8Type) ToPointer() *CreateRecordRequestBodyDNSRequest8Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest8Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest8Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest8Type: %v", v)
	}
}

type Eight struct {
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest8Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// A TXT record containing arbitrary text.
	Value string `json:"value"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *Eight) GetType() CreateRecordRequestBodyDNSRequest8Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest8Type("")
	}
	return o.Type
}

func (o *Eight) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Eight) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Eight) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequest7Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest7Type string

const (
	CreateRecordRequestBodyDNSRequest7TypeA     CreateRecordRequestBodyDNSRequest7Type = "A"
	CreateRecordRequestBodyDNSRequest7TypeAaaa  CreateRecordRequestBodyDNSRequest7Type = "AAAA"
	CreateRecordRequestBodyDNSRequest7TypeAlias CreateRecordRequestBodyDNSRequest7Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest7TypeCaa   CreateRecordRequestBodyDNSRequest7Type = "CAA"
	CreateRecordRequestBodyDNSRequest7TypeCname CreateRecordRequestBodyDNSRequest7Type = "CNAME"
	CreateRecordRequestBodyDNSRequest7TypeHTTPS CreateRecordRequestBodyDNSRequest7Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest7TypeMx    CreateRecordRequestBodyDNSRequest7Type = "MX"
	CreateRecordRequestBodyDNSRequest7TypeSrv   CreateRecordRequestBodyDNSRequest7Type = "SRV"
	CreateRecordRequestBodyDNSRequest7TypeTxt   CreateRecordRequestBodyDNSRequest7Type = "TXT"
	CreateRecordRequestBodyDNSRequest7TypeNs    CreateRecordRequestBodyDNSRequest7Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest7Type) ToPointer() *CreateRecordRequestBodyDNSRequest7Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest7Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest7Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest7Type: %v", v)
	}
}

type RequestBodySrv struct {
	Priority *float64 `json:"priority"`
	Weight   *float64 `json:"weight"`
	Port     *float64 `json:"port"`
	Target   string   `json:"target"`
}

func (o *RequestBodySrv) GetPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *RequestBodySrv) GetWeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

func (o *RequestBodySrv) GetPort() *float64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *RequestBodySrv) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}

type Seven struct {
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest7Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64       `json:"ttl,omitempty"`
	Srv RequestBodySrv `json:"srv"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *Seven) GetType() CreateRecordRequestBodyDNSRequest7Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest7Type("")
	}
	return o.Type
}

func (o *Seven) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Seven) GetSrv() RequestBodySrv {
	if o == nil {
		return RequestBodySrv{}
	}
	return o.Srv
}

func (o *Seven) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequest6Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest6Type string

const (
	CreateRecordRequestBodyDNSRequest6TypeA     CreateRecordRequestBodyDNSRequest6Type = "A"
	CreateRecordRequestBodyDNSRequest6TypeAaaa  CreateRecordRequestBodyDNSRequest6Type = "AAAA"
	CreateRecordRequestBodyDNSRequest6TypeAlias CreateRecordRequestBodyDNSRequest6Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest6TypeCaa   CreateRecordRequestBodyDNSRequest6Type = "CAA"
	CreateRecordRequestBodyDNSRequest6TypeCname CreateRecordRequestBodyDNSRequest6Type = "CNAME"
	CreateRecordRequestBodyDNSRequest6TypeHTTPS CreateRecordRequestBodyDNSRequest6Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest6TypeMx    CreateRecordRequestBodyDNSRequest6Type = "MX"
	CreateRecordRequestBodyDNSRequest6TypeSrv   CreateRecordRequestBodyDNSRequest6Type = "SRV"
	CreateRecordRequestBodyDNSRequest6TypeTxt   CreateRecordRequestBodyDNSRequest6Type = "TXT"
	CreateRecordRequestBodyDNSRequest6TypeNs    CreateRecordRequestBodyDNSRequest6Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest6Type) ToPointer() *CreateRecordRequestBodyDNSRequest6Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest6Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest6Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest6Type: %v", v)
	}
}

type Six struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest6Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// An MX record specifying the mail server responsible for accepting messages on behalf of the domain name.
	Value      string  `json:"value"`
	MxPriority float64 `json:"mxPriority"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *Six) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Six) GetType() CreateRecordRequestBodyDNSRequest6Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest6Type("")
	}
	return o.Type
}

func (o *Six) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *Six) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *Six) GetMxPriority() float64 {
	if o == nil {
		return 0.0
	}
	return o.MxPriority
}

func (o *Six) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequest5Type - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequest5Type string

const (
	CreateRecordRequestBodyDNSRequest5TypeA     CreateRecordRequestBodyDNSRequest5Type = "A"
	CreateRecordRequestBodyDNSRequest5TypeAaaa  CreateRecordRequestBodyDNSRequest5Type = "AAAA"
	CreateRecordRequestBodyDNSRequest5TypeAlias CreateRecordRequestBodyDNSRequest5Type = "ALIAS"
	CreateRecordRequestBodyDNSRequest5TypeCaa   CreateRecordRequestBodyDNSRequest5Type = "CAA"
	CreateRecordRequestBodyDNSRequest5TypeCname CreateRecordRequestBodyDNSRequest5Type = "CNAME"
	CreateRecordRequestBodyDNSRequest5TypeHTTPS CreateRecordRequestBodyDNSRequest5Type = "HTTPS"
	CreateRecordRequestBodyDNSRequest5TypeMx    CreateRecordRequestBodyDNSRequest5Type = "MX"
	CreateRecordRequestBodyDNSRequest5TypeSrv   CreateRecordRequestBodyDNSRequest5Type = "SRV"
	CreateRecordRequestBodyDNSRequest5TypeTxt   CreateRecordRequestBodyDNSRequest5Type = "TXT"
	CreateRecordRequestBodyDNSRequest5TypeNs    CreateRecordRequestBodyDNSRequest5Type = "NS"
)

func (e CreateRecordRequestBodyDNSRequest5Type) ToPointer() *CreateRecordRequestBodyDNSRequest5Type {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequest5Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequest5Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequest5Type: %v", v)
	}
}

type RequestBody5 struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequest5Type `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// A CNAME record mapping to another domain name.
	Value *string `json:"value,omitempty"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *RequestBody5) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RequestBody5) GetType() CreateRecordRequestBodyDNSRequest5Type {
	if o == nil {
		return CreateRecordRequestBodyDNSRequest5Type("")
	}
	return o.Type
}

func (o *RequestBody5) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RequestBody5) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *RequestBody5) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSRequestType - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSRequestType string

const (
	CreateRecordRequestBodyDNSRequestTypeA     CreateRecordRequestBodyDNSRequestType = "A"
	CreateRecordRequestBodyDNSRequestTypeAaaa  CreateRecordRequestBodyDNSRequestType = "AAAA"
	CreateRecordRequestBodyDNSRequestTypeAlias CreateRecordRequestBodyDNSRequestType = "ALIAS"
	CreateRecordRequestBodyDNSRequestTypeCaa   CreateRecordRequestBodyDNSRequestType = "CAA"
	CreateRecordRequestBodyDNSRequestTypeCname CreateRecordRequestBodyDNSRequestType = "CNAME"
	CreateRecordRequestBodyDNSRequestTypeHTTPS CreateRecordRequestBodyDNSRequestType = "HTTPS"
	CreateRecordRequestBodyDNSRequestTypeMx    CreateRecordRequestBodyDNSRequestType = "MX"
	CreateRecordRequestBodyDNSRequestTypeSrv   CreateRecordRequestBodyDNSRequestType = "SRV"
	CreateRecordRequestBodyDNSRequestTypeTxt   CreateRecordRequestBodyDNSRequestType = "TXT"
	CreateRecordRequestBodyDNSRequestTypeNs    CreateRecordRequestBodyDNSRequestType = "NS"
)

func (e CreateRecordRequestBodyDNSRequestType) ToPointer() *CreateRecordRequestBodyDNSRequestType {
	return &e
}
func (e *CreateRecordRequestBodyDNSRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSRequestType: %v", v)
	}
}

type RequestBody4 struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSRequestType `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// A CAA record to specify which Certificate Authorities (CAs) are allowed to issue certificates for the domain.
	Value string `json:"value"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *RequestBody4) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RequestBody4) GetType() CreateRecordRequestBodyDNSRequestType {
	if o == nil {
		return CreateRecordRequestBodyDNSRequestType("")
	}
	return o.Type
}

func (o *RequestBody4) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RequestBody4) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *RequestBody4) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyDNSType - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyDNSType string

const (
	CreateRecordRequestBodyDNSTypeA     CreateRecordRequestBodyDNSType = "A"
	CreateRecordRequestBodyDNSTypeAaaa  CreateRecordRequestBodyDNSType = "AAAA"
	CreateRecordRequestBodyDNSTypeAlias CreateRecordRequestBodyDNSType = "ALIAS"
	CreateRecordRequestBodyDNSTypeCaa   CreateRecordRequestBodyDNSType = "CAA"
	CreateRecordRequestBodyDNSTypeCname CreateRecordRequestBodyDNSType = "CNAME"
	CreateRecordRequestBodyDNSTypeHTTPS CreateRecordRequestBodyDNSType = "HTTPS"
	CreateRecordRequestBodyDNSTypeMx    CreateRecordRequestBodyDNSType = "MX"
	CreateRecordRequestBodyDNSTypeSrv   CreateRecordRequestBodyDNSType = "SRV"
	CreateRecordRequestBodyDNSTypeTxt   CreateRecordRequestBodyDNSType = "TXT"
	CreateRecordRequestBodyDNSTypeNs    CreateRecordRequestBodyDNSType = "NS"
)

func (e CreateRecordRequestBodyDNSType) ToPointer() *CreateRecordRequestBodyDNSType {
	return &e
}
func (e *CreateRecordRequestBodyDNSType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyDNSType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyDNSType: %v", v)
	}
}

type RequestBody3 struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyDNSType `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// An ALIAS virtual record pointing to a hostname resolved to an A record on server side.
	Value string `json:"value"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *RequestBody3) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RequestBody3) GetType() CreateRecordRequestBodyDNSType {
	if o == nil {
		return CreateRecordRequestBodyDNSType("")
	}
	return o.Type
}

func (o *RequestBody3) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RequestBody3) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *RequestBody3) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// CreateRecordRequestBodyType - The type of record, it could be one of the valid DNS records.
type CreateRecordRequestBodyType string

const (
	CreateRecordRequestBodyTypeA     CreateRecordRequestBodyType = "A"
	CreateRecordRequestBodyTypeAaaa  CreateRecordRequestBodyType = "AAAA"
	CreateRecordRequestBodyTypeAlias CreateRecordRequestBodyType = "ALIAS"
	CreateRecordRequestBodyTypeCaa   CreateRecordRequestBodyType = "CAA"
	CreateRecordRequestBodyTypeCname CreateRecordRequestBodyType = "CNAME"
	CreateRecordRequestBodyTypeHTTPS CreateRecordRequestBodyType = "HTTPS"
	CreateRecordRequestBodyTypeMx    CreateRecordRequestBodyType = "MX"
	CreateRecordRequestBodyTypeSrv   CreateRecordRequestBodyType = "SRV"
	CreateRecordRequestBodyTypeTxt   CreateRecordRequestBodyType = "TXT"
	CreateRecordRequestBodyTypeNs    CreateRecordRequestBodyType = "NS"
)

func (e CreateRecordRequestBodyType) ToPointer() *CreateRecordRequestBodyType {
	return &e
}
func (e *CreateRecordRequestBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = CreateRecordRequestBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRecordRequestBodyType: %v", v)
	}
}

type RequestBody2 struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type CreateRecordRequestBodyType `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// An AAAA record pointing to an IPv6 address.
	Value string `json:"value"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *RequestBody2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RequestBody2) GetType() CreateRecordRequestBodyType {
	if o == nil {
		return CreateRecordRequestBodyType("")
	}
	return o.Type
}

func (o *RequestBody2) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RequestBody2) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *RequestBody2) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

// RequestBodyType - The type of record, it could be one of the valid DNS records.
type RequestBodyType string

const (
	RequestBodyTypeA     RequestBodyType = "A"
	RequestBodyTypeAaaa  RequestBodyType = "AAAA"
	RequestBodyTypeAlias RequestBodyType = "ALIAS"
	RequestBodyTypeCaa   RequestBodyType = "CAA"
	RequestBodyTypeCname RequestBodyType = "CNAME"
	RequestBodyTypeHTTPS RequestBodyType = "HTTPS"
	RequestBodyTypeMx    RequestBodyType = "MX"
	RequestBodyTypeSrv   RequestBodyType = "SRV"
	RequestBodyTypeTxt   RequestBodyType = "TXT"
	RequestBodyTypeNs    RequestBodyType = "NS"
)

func (e RequestBodyType) ToPointer() *RequestBodyType {
	return &e
}
func (e *RequestBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "A":
		fallthrough
	case "AAAA":
		fallthrough
	case "ALIAS":
		fallthrough
	case "CAA":
		fallthrough
	case "CNAME":
		fallthrough
	case "HTTPS":
		fallthrough
	case "MX":
		fallthrough
	case "SRV":
		fallthrough
	case "TXT":
		fallthrough
	case "NS":
		*e = RequestBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestBodyType: %v", v)
	}
}

type RequestBody1 struct {
	// A subdomain name or an empty string for the root domain.
	Name string `json:"name"`
	// The type of record, it could be one of the valid DNS records.
	Type RequestBodyType `json:"type"`
	// The TTL value. Must be a number between 60 and 2147483647. Default value is 60.
	TTL *float64 `json:"ttl,omitempty"`
	// The record value must be a valid IPv4 address.
	Value string `json:"value"`
	// A comment to add context on what this DNS record is for
	Comment *string `json:"comment,omitempty"`
}

func (o *RequestBody1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RequestBody1) GetType() RequestBodyType {
	if o == nil {
		return RequestBodyType("")
	}
	return o.Type
}

func (o *RequestBody1) GetTTL() *float64 {
	if o == nil {
		return nil
	}
	return o.TTL
}

func (o *RequestBody1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *RequestBody1) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

type CreateRecordRequestBodyUnionType string

const (
	CreateRecordRequestBodyUnionTypeRequestBody1 CreateRecordRequestBodyUnionType = "requestBody_1"
	CreateRecordRequestBodyUnionTypeRequestBody2 CreateRecordRequestBodyUnionType = "requestBody_2"
	CreateRecordRequestBodyUnionTypeRequestBody3 CreateRecordRequestBodyUnionType = "requestBody_3"
	CreateRecordRequestBodyUnionTypeRequestBody4 CreateRecordRequestBodyUnionType = "requestBody_4"
	CreateRecordRequestBodyUnionTypeRequestBody5 CreateRecordRequestBodyUnionType = "requestBody_5"
	CreateRecordRequestBodyUnionTypeSix          CreateRecordRequestBodyUnionType = "6"
	CreateRecordRequestBodyUnionTypeSeven        CreateRecordRequestBodyUnionType = "7"
	CreateRecordRequestBodyUnionTypeEight        CreateRecordRequestBodyUnionType = "8"
	CreateRecordRequestBodyUnionTypeNine         CreateRecordRequestBodyUnionType = "9"
	CreateRecordRequestBodyUnionTypeTen          CreateRecordRequestBodyUnionType = "10"
)

type CreateRecordRequestBody struct {
	RequestBody1 *RequestBody1
	RequestBody2 *RequestBody2
	RequestBody3 *RequestBody3
	RequestBody4 *RequestBody4
	RequestBody5 *RequestBody5
	Six          *Six
	Seven        *Seven
	Eight        *Eight
	Nine         *Nine
	Ten          *Ten

	Type CreateRecordRequestBodyUnionType
}

func CreateCreateRecordRequestBodyRequestBody1(requestBody1 RequestBody1) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeRequestBody1

	return CreateRecordRequestBody{
		RequestBody1: &requestBody1,
		Type:         typ,
	}
}

func CreateCreateRecordRequestBodyRequestBody2(requestBody2 RequestBody2) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeRequestBody2

	return CreateRecordRequestBody{
		RequestBody2: &requestBody2,
		Type:         typ,
	}
}

func CreateCreateRecordRequestBodyRequestBody3(requestBody3 RequestBody3) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeRequestBody3

	return CreateRecordRequestBody{
		RequestBody3: &requestBody3,
		Type:         typ,
	}
}

func CreateCreateRecordRequestBodyRequestBody4(requestBody4 RequestBody4) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeRequestBody4

	return CreateRecordRequestBody{
		RequestBody4: &requestBody4,
		Type:         typ,
	}
}

func CreateCreateRecordRequestBodyRequestBody5(requestBody5 RequestBody5) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeRequestBody5

	return CreateRecordRequestBody{
		RequestBody5: &requestBody5,
		Type:         typ,
	}
}

func CreateCreateRecordRequestBodySix(six Six) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeSix

	return CreateRecordRequestBody{
		Six:  &six,
		Type: typ,
	}
}

func CreateCreateRecordRequestBodySeven(seven Seven) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeSeven

	return CreateRecordRequestBody{
		Seven: &seven,
		Type:  typ,
	}
}

func CreateCreateRecordRequestBodyEight(eight Eight) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeEight

	return CreateRecordRequestBody{
		Eight: &eight,
		Type:  typ,
	}
}

func CreateCreateRecordRequestBodyNine(nine Nine) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeNine

	return CreateRecordRequestBody{
		Nine: &nine,
		Type: typ,
	}
}

func CreateCreateRecordRequestBodyTen(ten Ten) CreateRecordRequestBody {
	typ := CreateRecordRequestBodyUnionTypeTen

	return CreateRecordRequestBody{
		Ten:  &ten,
		Type: typ,
	}
}

func (u *CreateRecordRequestBody) UnmarshalJSON(data []byte) error {

	var seven Seven = Seven{}
	if err := utils.UnmarshalJSON(data, &seven, "", true, true); err == nil {
		u.Seven = &seven
		u.Type = CreateRecordRequestBodyUnionTypeSeven
		return nil
	}

	var eight Eight = Eight{}
	if err := utils.UnmarshalJSON(data, &eight, "", true, true); err == nil {
		u.Eight = &eight
		u.Type = CreateRecordRequestBodyUnionTypeEight
		return nil
	}

	var ten Ten = Ten{}
	if err := utils.UnmarshalJSON(data, &ten, "", true, true); err == nil {
		u.Ten = &ten
		u.Type = CreateRecordRequestBodyUnionTypeTen
		return nil
	}

	var requestBody1 RequestBody1 = RequestBody1{}
	if err := utils.UnmarshalJSON(data, &requestBody1, "", true, true); err == nil {
		u.RequestBody1 = &requestBody1
		u.Type = CreateRecordRequestBodyUnionTypeRequestBody1
		return nil
	}

	var requestBody2 RequestBody2 = RequestBody2{}
	if err := utils.UnmarshalJSON(data, &requestBody2, "", true, true); err == nil {
		u.RequestBody2 = &requestBody2
		u.Type = CreateRecordRequestBodyUnionTypeRequestBody2
		return nil
	}

	var requestBody3 RequestBody3 = RequestBody3{}
	if err := utils.UnmarshalJSON(data, &requestBody3, "", true, true); err == nil {
		u.RequestBody3 = &requestBody3
		u.Type = CreateRecordRequestBodyUnionTypeRequestBody3
		return nil
	}

	var requestBody4 RequestBody4 = RequestBody4{}
	if err := utils.UnmarshalJSON(data, &requestBody4, "", true, true); err == nil {
		u.RequestBody4 = &requestBody4
		u.Type = CreateRecordRequestBodyUnionTypeRequestBody4
		return nil
	}

	var requestBody5 RequestBody5 = RequestBody5{}
	if err := utils.UnmarshalJSON(data, &requestBody5, "", true, true); err == nil {
		u.RequestBody5 = &requestBody5
		u.Type = CreateRecordRequestBodyUnionTypeRequestBody5
		return nil
	}

	var nine Nine = Nine{}
	if err := utils.UnmarshalJSON(data, &nine, "", true, true); err == nil {
		u.Nine = &nine
		u.Type = CreateRecordRequestBodyUnionTypeNine
		return nil
	}

	var six Six = Six{}
	if err := utils.UnmarshalJSON(data, &six, "", true, true); err == nil {
		u.Six = &six
		u.Type = CreateRecordRequestBodyUnionTypeSix
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateRecordRequestBody", string(data))
}

func (u CreateRecordRequestBody) MarshalJSON() ([]byte, error) {
	if u.RequestBody1 != nil {
		return utils.MarshalJSON(u.RequestBody1, "", true)
	}

	if u.RequestBody2 != nil {
		return utils.MarshalJSON(u.RequestBody2, "", true)
	}

	if u.RequestBody3 != nil {
		return utils.MarshalJSON(u.RequestBody3, "", true)
	}

	if u.RequestBody4 != nil {
		return utils.MarshalJSON(u.RequestBody4, "", true)
	}

	if u.RequestBody5 != nil {
		return utils.MarshalJSON(u.RequestBody5, "", true)
	}

	if u.Six != nil {
		return utils.MarshalJSON(u.Six, "", true)
	}

	if u.Seven != nil {
		return utils.MarshalJSON(u.Seven, "", true)
	}

	if u.Eight != nil {
		return utils.MarshalJSON(u.Eight, "", true)
	}

	if u.Nine != nil {
		return utils.MarshalJSON(u.Nine, "", true)
	}

	if u.Ten != nil {
		return utils.MarshalJSON(u.Ten, "", true)
	}

	return nil, errors.New("could not marshal union type CreateRecordRequestBody: all fields are null")
}

type CreateRecordRequest struct {
	// The domain used to create the DNS record.
	Domain string `pathParam:"style=simple,explode=false,name=domain"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                 `queryParam:"style=form,explode=true,name=slug"`
	RequestBody CreateRecordRequestBody `request:"mediaType=application/json"`
}

func (o *CreateRecordRequest) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *CreateRecordRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateRecordRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateRecordRequest) GetRequestBody() CreateRecordRequestBody {
	if o == nil {
		return CreateRecordRequestBody{}
	}
	return o.RequestBody
}

type CreateRecordResponseBody2 struct {
	// The id of the newly created DNS record
	UID string `json:"uid"`
}

func (o *CreateRecordResponseBody2) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

type CreateRecordResponseBody1 struct {
	UID     string  `json:"uid"`
	Updated float64 `json:"updated"`
}

func (o *CreateRecordResponseBody1) GetUID() string {
	if o == nil {
		return ""
	}
	return o.UID
}

func (o *CreateRecordResponseBody1) GetUpdated() float64 {
	if o == nil {
		return 0.0
	}
	return o.Updated
}

type CreateRecordResponseBodyType string

const (
	CreateRecordResponseBodyTypeCreateRecordResponseBody1 CreateRecordResponseBodyType = "createRecord_responseBody_1"
	CreateRecordResponseBodyTypeCreateRecordResponseBody2 CreateRecordResponseBodyType = "createRecord_responseBody_2"
)

// CreateRecordResponseBody - Successful response showing the uid of the newly created DNS record.
type CreateRecordResponseBody struct {
	CreateRecordResponseBody1 *CreateRecordResponseBody1
	CreateRecordResponseBody2 *CreateRecordResponseBody2

	Type CreateRecordResponseBodyType
}

func CreateCreateRecordResponseBodyCreateRecordResponseBody1(createRecordResponseBody1 CreateRecordResponseBody1) CreateRecordResponseBody {
	typ := CreateRecordResponseBodyTypeCreateRecordResponseBody1

	return CreateRecordResponseBody{
		CreateRecordResponseBody1: &createRecordResponseBody1,
		Type:                      typ,
	}
}

func CreateCreateRecordResponseBodyCreateRecordResponseBody2(createRecordResponseBody2 CreateRecordResponseBody2) CreateRecordResponseBody {
	typ := CreateRecordResponseBodyTypeCreateRecordResponseBody2

	return CreateRecordResponseBody{
		CreateRecordResponseBody2: &createRecordResponseBody2,
		Type:                      typ,
	}
}

func (u *CreateRecordResponseBody) UnmarshalJSON(data []byte) error {

	var createRecordResponseBody2 CreateRecordResponseBody2 = CreateRecordResponseBody2{}
	if err := utils.UnmarshalJSON(data, &createRecordResponseBody2, "", true, true); err == nil {
		u.CreateRecordResponseBody2 = &createRecordResponseBody2
		u.Type = CreateRecordResponseBodyTypeCreateRecordResponseBody2
		return nil
	}

	var createRecordResponseBody1 CreateRecordResponseBody1 = CreateRecordResponseBody1{}
	if err := utils.UnmarshalJSON(data, &createRecordResponseBody1, "", true, true); err == nil {
		u.CreateRecordResponseBody1 = &createRecordResponseBody1
		u.Type = CreateRecordResponseBodyTypeCreateRecordResponseBody1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateRecordResponseBody", string(data))
}

func (u CreateRecordResponseBody) MarshalJSON() ([]byte, error) {
	if u.CreateRecordResponseBody1 != nil {
		return utils.MarshalJSON(u.CreateRecordResponseBody1, "", true)
	}

	if u.CreateRecordResponseBody2 != nil {
		return utils.MarshalJSON(u.CreateRecordResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type CreateRecordResponseBody: all fields are null")
}

type CreateRecordResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful response showing the uid of the newly created DNS record.
	OneOf *CreateRecordResponseBody
}

func (o *CreateRecordResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateRecordResponse) GetOneOf() *CreateRecordResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
