// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
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
func (r *V2AccountWithdrawalRequestService) New(ctx context.Context, accountID string, body V2AccountWithdrawalRequestNewParams, opts ...option.RequestOption) (res *WithdrawalRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific `WithdrawalRequest` by its ID.
func (r *V2AccountWithdrawalRequestService) Get(ctx context.Context, withdrawalRequestID string, query V2AccountWithdrawalRequestGetParams, opts ...option.RequestOption) (res *WithdrawalRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if withdrawalRequestID == "" {
		err = errors.New("missing required withdrawal_request_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests/%s", query.AccountID, withdrawalRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List `WithdrawalRequests` under the `Account`, sorted by most recent.
func (r *V2AccountWithdrawalRequestService) List(ctx context.Context, accountID string, query V2AccountWithdrawalRequestListParams, opts ...option.RequestOption) (res *[]WithdrawalRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawal_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information for a withdrawal request of payment tokens from an `Account` backed
// by a Dinari-managed `Wallet`.
type WithdrawalRequest struct {
	// ID of the `WithdrawalRequest`.
	ID string `json:"id,required" format:"uuid"`
	// ID of the `Account` of the `WithdrawalRequest`.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Datetime at which the `WithdrawalRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Amount of USD+ payment tokens submitted for withdrawal.
	PaymentTokenAmount float64 `json:"payment_token_amount,required"`
	// ID of the `Account` that will receive USDC payment tokens from the `Withdrawal`.
	// This `Account` must be connected to a non-managed `Wallet` and belong to the
	// same `Entity`.
	RecipientAccountID string `json:"recipient_account_id,required" format:"uuid"`
	// Status of the `WithdrawalRequest`
	//
	// Any of "PENDING", "SUBMITTED", "ERROR", "CANCELLED".
	Status WithdrawalRequestStatus `json:"status,required"`
	// Datetime at which the `WithdrawalRequest` was updated. ISO 8601 timestamp.
	UpdatedDt time.Time `json:"updated_dt,required" format:"date-time"`
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

type V2AccountWithdrawalRequestNewParams struct {
	// Amount of USD+ payment tokens to be withdrawn. Must be greater than 0 and have
	// at most 6 decimal places.
	PaymentTokenQuantity float64 `json:"payment_token_quantity,required"`
	// ID of the `Account` that will receive payment tokens from the `Withdrawal`.
	RecipientAccountID string `json:"recipient_account_id,required" format:"uuid"`
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
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountWithdrawalRequestListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
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
