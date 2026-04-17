// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apiquery"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// **`Accounts` represent the financial accounts of an `Entity`.**
//
// `Orders`, dividends, and other transactions are associated with an `Account`.
//
// V2EntityAccountService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityAccountService] method instead.
type V2EntityAccountService struct {
	Options []option.RequestOption
}

// NewV2EntityAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2EntityAccountService(opts ...option.RequestOption) (r V2EntityAccountService) {
	r = V2EntityAccountService{}
	r.Options = opts
	return
}

// Create a new `Account` for a specific `Entity`. This `Entity` represents your
// organization itself, or an individual customer of your organization.
func (r *V2EntityAccountService) New(ctx context.Context, entityID string, body V2EntityAccountNewParams, opts ...option.RequestOption) (res *V2EntityAccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a list of all `Accounts` that belong to a specific `Entity`. This `Entity`
// represents your organization itself, or an individual customer of your
// organization.
func (r *V2EntityAccountService) List(ctx context.Context, entityID string, query V2EntityAccountListParams, opts ...option.RequestOption) (res *V2EntityAccountListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Information about an `Account` owned by an `Entity`.
type Account struct {
	// Unique ID for the `Account`.
	ID string `json:"id" api:"required" format:"uuid"`
	// Datetime when the `Account` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// ID for the `Entity` that owns the `Account`.
	EntityID string `json:"entity_id" api:"required" format:"uuid"`
	// Indicates whether the `Account` is active.
	IsActive bool `json:"is_active" api:"required"`
	// Jurisdiction of the `Account`.
	//
	// Any of "BASELINE", "US".
	Jurisdiction AccountJurisdiction `json:"jurisdiction" api:"required"`
	// ID of the brokerage account associated with the `Account`.
	BrokerageAccountID string `json:"brokerage_account_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedDt          respjson.Field
		EntityID           respjson.Field
		IsActive           respjson.Field
		Jurisdiction       respjson.Field
		BrokerageAccountID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Jurisdiction of the `Account`.
type AccountJurisdiction string

const (
	AccountJurisdictionBaseline AccountJurisdiction = "BASELINE"
	AccountJurisdictionUs       AccountJurisdiction = "US"
)

type Jurisdiction string

const (
	JurisdictionBaseline Jurisdiction = "BASELINE"
	JurisdictionUs       Jurisdiction = "US"
)

// Information about an `Account` owned by an `Entity`.
type V2EntityAccountNewResponse struct {
	// Unique ID for the `Account`.
	ID string `json:"id" api:"required" format:"uuid"`
	// Datetime when the `Account` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// ID for the `Entity` that owns the `Account`.
	EntityID string `json:"entity_id" api:"required" format:"uuid"`
	// Indicates whether the `Account` is active.
	IsActive bool `json:"is_active" api:"required"`
	// Jurisdiction of the `Account`.
	//
	// Any of "BASELINE", "US".
	Jurisdiction Jurisdiction `json:"jurisdiction" api:"required"`
	// ID of the brokerage account associated with the `Account`.
	BrokerageAccountID string `json:"brokerage_account_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedDt          respjson.Field
		EntityID           respjson.Field
		IsActive           respjson.Field
		Jurisdiction       respjson.Field
		BrokerageAccountID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V2EntityAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// V2EntityAccountListResponseUnion contains all possible properties and values
// from [[]Account], [V2EntityAccountListResponsePaginatedAccountResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAccountArray]
type V2EntityAccountListResponseUnion struct {
	// This field will be present if the value is a [[]Account] instead of an object.
	OfAccountArray []Account `json:",inline"`
	// This field is from variant
	// [V2EntityAccountListResponsePaginatedAccountResponse].
	Data []Account `json:"data"`
	// This field is from variant
	// [V2EntityAccountListResponsePaginatedAccountResponse].
	PaginationMetadata V2EntityAccountListResponsePaginatedAccountResponsePaginationMetadata `json:"pagination_metadata"`
	// This field is from variant
	// [V2EntityAccountListResponsePaginatedAccountResponse].
	Sv   string `json:"_sv"`
	JSON struct {
		OfAccountArray     respjson.Field
		Data               respjson.Field
		PaginationMetadata respjson.Field
		Sv                 respjson.Field
		raw                string
	} `json:"-"`
}

func (u V2EntityAccountListResponseUnion) AsAccountArray() (v []Account) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2EntityAccountListResponseUnion) AsV2EntityAccountListResponsePaginatedAccountResponse() (v V2EntityAccountListResponsePaginatedAccountResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2EntityAccountListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2EntityAccountListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityAccountListResponsePaginatedAccountResponse struct {
	// List of Account
	Data []Account `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata V2EntityAccountListResponsePaginatedAccountResponsePaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedAccountResponse:v1".
	Sv string `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data               respjson.Field
		PaginationMetadata respjson.Field
		Sv                 respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityAccountListResponsePaginatedAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *V2EntityAccountListResponsePaginatedAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type V2EntityAccountListResponsePaginatedAccountResponsePaginationMetadata struct {
	// Cursor for next page
	Next string `json:"next"`
	// Cursor for previous page
	Previous string `json:"previous"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityAccountListResponsePaginatedAccountResponsePaginationMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *V2EntityAccountListResponsePaginatedAccountResponsePaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityAccountNewParams struct {
	// Jurisdiction of the `Account`.
	//
	// Any of "BASELINE", "US".
	Jurisdiction Jurisdiction `json:"jurisdiction,omitzero"`
	paramObj
}

func (r V2EntityAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityAccountListParams struct {
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Number of results to return
	Limit    param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2EntityAccountListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2EntityAccountListParams]'s query parameters as
// `url.Values`.
func (r V2EntityAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2EntityAccountListParamsOrder string

const (
	V2EntityAccountListParamsOrderAsc  V2EntityAccountListParamsOrder = "asc"
	V2EntityAccountListParamsOrderDesc V2EntityAccountListParamsOrder = "desc"
)
