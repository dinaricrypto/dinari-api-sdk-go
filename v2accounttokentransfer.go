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

// V2AccountTokenTransferService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountTokenTransferService] method instead.
type V2AccountTokenTransferService struct {
	Options []option.RequestOption
}

// NewV2AccountTokenTransferService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountTokenTransferService(opts ...option.RequestOption) (r V2AccountTokenTransferService) {
	r = V2AccountTokenTransferService{}
	r.Options = opts
	return
}

// Creates a `TokenTransfer` from this `Account`.
//
// A `TokenTransfer` represents a transfer of tokens through the Dinari platform
// from one `Account` to another. As such, only `Account`s that are connected to
// Dinari-managed `Wallet`s can initiate `TokenTransfer`s.
func (r *V2AccountTokenTransferService) New(ctx context.Context, accountID string, body V2AccountTokenTransferNewParams, opts ...option.RequestOption) (res *TokenTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/token_transfers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific `TokenTransfer` made from this `Account` by its ID.
//
// A `TokenTransfer` represents a transfer of tokens through the Dinari platform
// from one `Account` to another. As such, only `Account`s that are connected to
// Dinari-managed `Wallet`s can initiate `TokenTransfer`s.
func (r *V2AccountTokenTransferService) Get(ctx context.Context, transferID string, query V2AccountTokenTransferGetParams, opts ...option.RequestOption) (res *TokenTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if transferID == "" {
		err = errors.New("missing required transfer_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/token_transfers/%s", query.AccountID, transferID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get `TokenTransfer`s made from this `Account`.
//
// A `TokenTransfer` represents a transfer of tokens through the Dinari platform
// from one `Account` to another. As such, only `Account`s that are connected to
// Dinari-managed `Wallet`s can initiate `TokenTransfer`s.
func (r *V2AccountTokenTransferService) List(ctx context.Context, accountID string, query V2AccountTokenTransferListParams, opts ...option.RequestOption) (res *[]TokenTransfer, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/token_transfers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about a token transfer between accounts.
type TokenTransfer struct {
	// ID of the token transfer.
	ID string `json:"id,required" format:"uuid"`
	// CAIP-2 chain ID of the blockchain that the transfer is made on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,required"`
	// Datetime at which the transfer was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Quantity of the token being transferred.
	Quantity float64 `json:"quantity,required"`
	// ID of the account to which the tokens are transferred.
	RecipientAccountID string `json:"recipient_account_id,required" format:"uuid"`
	// ID of the account from which the tokens are transferred.
	SenderAccountID string `json:"sender_account_id,required" format:"uuid"`
	// Status of the token transfer.
	//
	// Any of "PENDING", "IN_PROGRESS", "COMPLETE", "FAILED".
	Status TokenTransferStatus `json:"status,required"`
	// Address of the token being transferred.
	TokenAddress string `json:"token_address,required" format:"eth_address"`
	// Datetime at which the transfer was last updated. ISO 8601 timestamp.
	UpdatedDt time.Time `json:"updated_dt,required" format:"date-time"`
	// Transaction hash of the transfer on the blockchain, if applicable. This is only
	// present if the transfer has been executed on-chain.
	TransactionHash string `json:"transaction_hash,nullable" format:"hex_string"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		ChainID            respjson.Field
		CreatedDt          respjson.Field
		Quantity           respjson.Field
		RecipientAccountID respjson.Field
		SenderAccountID    respjson.Field
		Status             respjson.Field
		TokenAddress       respjson.Field
		UpdatedDt          respjson.Field
		TransactionHash    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenTransfer) RawJSON() string { return r.JSON.raw }
func (r *TokenTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the token transfer.
type TokenTransferStatus string

const (
	TokenTransferStatusPending    TokenTransferStatus = "PENDING"
	TokenTransferStatusInProgress TokenTransferStatus = "IN_PROGRESS"
	TokenTransferStatusComplete   TokenTransferStatus = "COMPLETE"
	TokenTransferStatusFailed     TokenTransferStatus = "FAILED"
)

type V2AccountTokenTransferNewParams struct {
	// Quantity of the token to transfer.
	Quantity float64 `json:"quantity,required"`
	// ID of the recipient account to which the tokens will be transferred.
	RecipientAccountID string `json:"recipient_account_id,required" format:"uuid"`
	// Address of the token to transfer.
	TokenAddress string `json:"token_address,required" format:"eth_address"`
	paramObj
}

func (r V2AccountTokenTransferNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountTokenTransferNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountTokenTransferNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountTokenTransferGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountTokenTransferListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountTokenTransferListParams]'s query parameters as
// `url.Values`.
func (r V2AccountTokenTransferListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
