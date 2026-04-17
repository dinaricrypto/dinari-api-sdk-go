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

// **`Withdrawals` represent the transfer of stablecoins from an `Account`
// connected to a managed `Wallet` to another `Account` that is owned by the
// `Entity`.**
//
// Since the `Account` is backed by a managed `Wallet`, the `Withdrawal` must be
// processed by Dinari and the corresponding transaction is submitted on chain.
//
// Upon requesting a withdrawal, a `WithdrawalRequest` is created, which is then
// submitted on chain by Dinari. Once the transfer is submitted on chain, the
// corresponding `Withdrawal` is created.
//
// Currently, withdrawals are made in USDC on the Arbitrum network (Chain ID
// `eip155:42161`).
//
// V2AccountWithdrawalRequestService contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountWithdrawalRequestService] method instead.
type V2AccountWithdrawalRequestService struct {
	Options []option.RequestOption
}

// NewV2AccountWithdrawalRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV2AccountWithdrawalRequestService(opts ...option.RequestOption) (r V2AccountWithdrawalRequestService) {
	r = V2AccountWithdrawalRequestService{}
	r.Options = opts
	return
}

// Request to withdraw USD+ payment tokens from a managed `Account` and send the
// equivalent amount of USDC to the specified recipient `Account`.
//
// The recipient `Account` must belong to the same `Entity` as the managed
// `Account`.
func (r *V2AccountWithdrawalRequestService) New(ctx context.Context, accountID string, body V2AccountWithdrawalRequestNewParams, opts ...option.RequestOption) (res *V2AccountWithdrawalRequestNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a specific `WithdrawalRequest` by its ID.
func (r *V2AccountWithdrawalRequestService) Get(ctx context.Context, withdrawalRequestID string, query V2AccountWithdrawalRequestGetParams, opts ...option.RequestOption) (res *V2AccountWithdrawalRequestGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if withdrawalRequestID == "" {
		err = errors.New("missing required withdrawal_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests/%s", query.AccountID, withdrawalRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List `WithdrawalRequests` under the `Account`, sorted by most recent.
func (r *V2AccountWithdrawalRequestService) List(ctx context.Context, accountID string, query V2AccountWithdrawalRequestListParams, opts ...option.RequestOption) (res *V2AccountWithdrawalRequestListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Information for a withdrawal request of payment tokens from an `Account` backed
// by a Dinari-managed `Wallet`.
type WithdrawalRequest struct {
	// ID of the `WithdrawalRequest`.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of the `Account` of the `WithdrawalRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `WithdrawalRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Amount of USD+ payment tokens submitted for withdrawal.
	PaymentTokenAmount float64 `json:"payment_token_amount" api:"required"`
	// ID of the `Account` that will receive USDC payment tokens from the `Withdrawal`.
	// This `Account` must be connected to a non-managed `Wallet` and belong to the
	// same `Entity`.
	RecipientAccountID string `json:"recipient_account_id" api:"required" format:"uuid"`
	// Status of the `WithdrawalRequest`
	//
	// Any of "PENDING", "SUBMITTED", "ERROR", "CANCELLED".
	Status WithdrawalRequestStatus `json:"status" api:"required"`
	// Datetime at which the `WithdrawalRequest` was updated. ISO 8601 timestamp.
	UpdatedDt time.Time `json:"updated_dt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		PaymentTokenAmount respjson.Field
		RecipientAccountID respjson.Field
		Status             respjson.Field
		UpdatedDt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WithdrawalRequest) RawJSON() string { return r.JSON.raw }
func (r *WithdrawalRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the `WithdrawalRequest`
type WithdrawalRequestStatus string

const (
	WithdrawalRequestStatusPending   WithdrawalRequestStatus = "PENDING"
	WithdrawalRequestStatusSubmitted WithdrawalRequestStatus = "SUBMITTED"
	WithdrawalRequestStatusError     WithdrawalRequestStatus = "ERROR"
	WithdrawalRequestStatusCancelled WithdrawalRequestStatus = "CANCELLED"
)

// Information for a withdrawal request of payment tokens from an `Account` backed
// by a Dinari-managed `Wallet`.
type V2AccountWithdrawalRequestNewResponse struct {
	// ID of the `WithdrawalRequest`.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of the `Account` of the `WithdrawalRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `WithdrawalRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Amount of USD+ payment tokens submitted for withdrawal.
	PaymentTokenAmount float64 `json:"payment_token_amount" api:"required"`
	// ID of the `Account` that will receive USDC payment tokens from the `Withdrawal`.
	// This `Account` must be connected to a non-managed `Wallet` and belong to the
	// same `Entity`.
	RecipientAccountID string `json:"recipient_account_id" api:"required" format:"uuid"`
	// Status of the `WithdrawalRequest`
	//
	// Any of "PENDING", "SUBMITTED", "ERROR", "CANCELLED".
	Status V2AccountWithdrawalRequestNewResponseStatus `json:"status" api:"required"`
	// Datetime at which the `WithdrawalRequest` was updated. ISO 8601 timestamp.
	UpdatedDt time.Time `json:"updated_dt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		PaymentTokenAmount respjson.Field
		RecipientAccountID respjson.Field
		Status             respjson.Field
		UpdatedDt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountWithdrawalRequestNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountWithdrawalRequestNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the `WithdrawalRequest`
type V2AccountWithdrawalRequestNewResponseStatus string

const (
	V2AccountWithdrawalRequestNewResponseStatusPending   V2AccountWithdrawalRequestNewResponseStatus = "PENDING"
	V2AccountWithdrawalRequestNewResponseStatusSubmitted V2AccountWithdrawalRequestNewResponseStatus = "SUBMITTED"
	V2AccountWithdrawalRequestNewResponseStatusError     V2AccountWithdrawalRequestNewResponseStatus = "ERROR"
	V2AccountWithdrawalRequestNewResponseStatusCancelled V2AccountWithdrawalRequestNewResponseStatus = "CANCELLED"
)

// Information for a withdrawal request of payment tokens from an `Account` backed
// by a Dinari-managed `Wallet`.
type V2AccountWithdrawalRequestGetResponse struct {
	// ID of the `WithdrawalRequest`.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of the `Account` of the `WithdrawalRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `WithdrawalRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Amount of USD+ payment tokens submitted for withdrawal.
	PaymentTokenAmount float64 `json:"payment_token_amount" api:"required"`
	// ID of the `Account` that will receive USDC payment tokens from the `Withdrawal`.
	// This `Account` must be connected to a non-managed `Wallet` and belong to the
	// same `Entity`.
	RecipientAccountID string `json:"recipient_account_id" api:"required" format:"uuid"`
	// Status of the `WithdrawalRequest`
	//
	// Any of "PENDING", "SUBMITTED", "ERROR", "CANCELLED".
	Status V2AccountWithdrawalRequestGetResponseStatus `json:"status" api:"required"`
	// Datetime at which the `WithdrawalRequest` was updated. ISO 8601 timestamp.
	UpdatedDt time.Time `json:"updated_dt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		PaymentTokenAmount respjson.Field
		RecipientAccountID respjson.Field
		Status             respjson.Field
		UpdatedDt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountWithdrawalRequestGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountWithdrawalRequestGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the `WithdrawalRequest`
type V2AccountWithdrawalRequestGetResponseStatus string

const (
	V2AccountWithdrawalRequestGetResponseStatusPending   V2AccountWithdrawalRequestGetResponseStatus = "PENDING"
	V2AccountWithdrawalRequestGetResponseStatusSubmitted V2AccountWithdrawalRequestGetResponseStatus = "SUBMITTED"
	V2AccountWithdrawalRequestGetResponseStatusError     V2AccountWithdrawalRequestGetResponseStatus = "ERROR"
	V2AccountWithdrawalRequestGetResponseStatusCancelled V2AccountWithdrawalRequestGetResponseStatus = "CANCELLED"
)

// V2AccountWithdrawalRequestListResponseUnion contains all possible properties and
// values from [[]WithdrawalRequest],
// [V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfWithdrawalRequestArray]
type V2AccountWithdrawalRequestListResponseUnion struct {
	// This field will be present if the value is a [[]WithdrawalRequest] instead of an
	// object.
	OfWithdrawalRequestArray []WithdrawalRequest `json:",inline"`
	// This field is from variant
	// [V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse].
	Data []WithdrawalRequest `json:"data"`
	// This field is from variant
	// [V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse].
	PaginationMetadata V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponsePaginationMetadata `json:"pagination_metadata"`
	// This field is from variant
	// [V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse].
	Sv   string `json:"_sv"`
	JSON struct {
		OfWithdrawalRequestArray respjson.Field
		Data                     respjson.Field
		PaginationMetadata       respjson.Field
		Sv                       respjson.Field
		raw                      string
	} `json:"-"`
}

func (u V2AccountWithdrawalRequestListResponseUnion) AsWithdrawalRequestArray() (v []WithdrawalRequest) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2AccountWithdrawalRequestListResponseUnion) AsV2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse() (v V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2AccountWithdrawalRequestListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2AccountWithdrawalRequestListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse struct {
	// List of WithdrawalRequest
	Data []WithdrawalRequest `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponsePaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedWithdrawalRequestResponse:v1".
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
func (r V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponsePaginationMetadata struct {
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
func (r V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponsePaginationMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountWithdrawalRequestListResponsePaginatedWithdrawalRequestResponsePaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWithdrawalRequestNewParams struct {
	// Amount of USD+ payment tokens to be withdrawn. Must be greater than 0 and have
	// at most 6 decimal places.
	PaymentTokenQuantity float64 `json:"payment_token_quantity" api:"required"`
	// ID of the `Account` that will receive payment tokens from the `Withdrawal`.
	RecipientAccountID string `json:"recipient_account_id" api:"required" format:"uuid"`
	paramObj
}

func (r V2AccountWithdrawalRequestNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountWithdrawalRequestNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountWithdrawalRequestNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWithdrawalRequestGetParams struct {
	AccountID string `path:"account_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountWithdrawalRequestListParams struct {
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
	Order V2AccountWithdrawalRequestListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountWithdrawalRequestListParams]'s query parameters as
// `url.Values`.
func (r V2AccountWithdrawalRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2AccountWithdrawalRequestListParamsOrder string

const (
	V2AccountWithdrawalRequestListParamsOrderAsc  V2AccountWithdrawalRequestListParamsOrder = "asc"
	V2AccountWithdrawalRequestListParamsOrderDesc V2AccountWithdrawalRequestListParamsOrder = "desc"
)
