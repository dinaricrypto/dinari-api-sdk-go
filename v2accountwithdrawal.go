// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/apiquery"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/respjson"
)

// V2AccountWithdrawalService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountWithdrawalService] method instead.
type V2AccountWithdrawalService struct {
	Options []option.RequestOption
}

// NewV2AccountWithdrawalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountWithdrawalService(opts ...option.RequestOption) (r V2AccountWithdrawalService) {
	r = V2AccountWithdrawalService{}
	r.Options = opts
	return
}

// Get a specific `Withdrawal` by its ID.
func (r *V2AccountWithdrawalService) Get(ctx context.Context, withdrawalID string, query V2AccountWithdrawalGetParams, opts ...option.RequestOption) (res *Withdrawal, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if withdrawalID == "" {
		err = errors.New("missing required withdrawal_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawals/%s", query.AccountID, withdrawalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all `Withdrawals` under the `Account`, sorted by most recent.
func (r *V2AccountWithdrawalService) List(ctx context.Context, accountID string, query V2AccountWithdrawalListParams, opts ...option.RequestOption) (res *[]Withdrawal, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/withdrawals", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information for a withdrawal of payment tokens from an `Account` backed by a
// Dinari-managed `Wallet`.
type Withdrawal struct {
	// ID of the `Withdrawal`.
	ID string `json:"id,required" format:"uuid"`
	// ID of the `Account` from which the `Withdrawal` is made.
	AccountID string `json:"account_id,required" format:"uuid"`
	// CAIP-2 chain ID of the blockchain where the `Withdrawal` is made.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
	ChainID Chain `json:"chain_id,required"`
	// Address of USDC payment token that the `Withdrawal` will be received in.
	PaymentTokenAddress string `json:"payment_token_address,required" format:"eth_address"`
	// Amount of USDC payment tokens to be withdrawn.
	PaymentTokenAmount float64 `json:"payment_token_amount,required"`
	// ID of the `Account` that will receive payment tokens from the `Withdrawal`. This
	// `Account` must be connected to a non-managed `Wallet` and belong to the same
	// `Entity`.
	RecipientAccountID string `json:"recipient_account_id,required" format:"uuid"`
	// Status of the `Withdrawal`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "FILLED", "REJECTED", "REQUIRING_CONTACT",
	// "ERROR".
	Status BrokerageOrderStatus `json:"status,required"`
	// Datetime at which the `Withdrawal` was transacted. ISO 8601 timestamp.
	TransactionDt time.Time `json:"transaction_dt,required" format:"date-time"`
	// Hash of the transaction for the `Withdrawal`.
	TransactionHash string `json:"transaction_hash,required" format:"hex_string"`
	// ID of the `WithdrawalRequest` associated with this `Withdrawal`.
	WithdrawalRequestID string `json:"withdrawal_request_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AccountID           respjson.Field
		ChainID             respjson.Field
		PaymentTokenAddress respjson.Field
		PaymentTokenAmount  respjson.Field
		RecipientAccountID  respjson.Field
		Status              respjson.Field
		TransactionDt       respjson.Field
		TransactionHash     respjson.Field
		WithdrawalRequestID respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Withdrawal) RawJSON() string { return r.JSON.raw }
func (r *Withdrawal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWithdrawalGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountWithdrawalListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// ID of the `WithdrawalRequest` to find `Withdrawals` for.
	WithdrawalRequestID param.Opt[string] `query:"withdrawal_request_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountWithdrawalListParams]'s query parameters as
// `url.Values`.
func (r V2AccountWithdrawalListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
