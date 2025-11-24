// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	shimjson "github.com/dinaricrypto/dinari-api-sdk-go/internal/encoding/json"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// V2AccountOrderRequestEip155Service contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderRequestEip155Service] method instead.
type V2AccountOrderRequestEip155Service struct {
	Options []option.RequestOption
}

// NewV2AccountOrderRequestEip155Service generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV2AccountOrderRequestEip155Service(opts ...option.RequestOption) (r V2AccountOrderRequestEip155Service) {
	r = V2AccountOrderRequestEip155Service{}
	r.Options = opts
	return
}

// Generates a permit that can be signed and used to create an `OrderRequest` using
// Dinari's EVM smart contracts.
//
// This is a convenience method to prepare the transactions needed to create an
// `OrderRequest` using Dinari's EVM smart contracts. Once signed, the transactions
// can be sent to the EVM network to create the order. Note that the fee quote is
// already included in the transactions, so no additional fee quote lookup is
// needed.
func (r *V2AccountOrderRequestEip155Service) NewPermit(ctx context.Context, accountID string, body V2AccountOrderRequestEip155NewPermitParams, opts ...option.RequestOption) (res *V2AccountOrderRequestEip155NewPermitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/eip155/permit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Prepare a transaction to be placed on EVM. The returned structure contains the
// necessary data to create an `EIP155Transaction` object.
func (r *V2AccountOrderRequestEip155Service) NewPermitTransaction(ctx context.Context, accountID string, body V2AccountOrderRequestEip155NewPermitTransactionParams, opts ...option.RequestOption) (res *V2AccountOrderRequestEip155NewPermitTransactionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/eip155/permit_transaction", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submits a transaction for an EIP155 Order Request given the EIP155OrderRequest
// ID and Permit Signature.
//
// An `EIP155OrderRequest` representing the proxied order is returned.
func (r *V2AccountOrderRequestEip155Service) Submit(ctx context.Context, accountID string, body V2AccountOrderRequestEip155SubmitParams, opts ...option.RequestOption) (res *V2AccountOrderRequestEip155SubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/eip155", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Input parameters for creating a proxied `EIP155OrderRequestPermitTransaction`.
//
// The properties OrderRequestID, PermitSignature are required.
type Eip155OrderRequestPermitTransactionParam struct {
	// ID of the prepared proxied order to be submitted as a proxied order.
	OrderRequestID string `json:"order_request_id,required" format:"uuid"`
	// Signature of the permit typed data, allowing Dinari to spend the payment token
	// or dShare asset token on behalf of the owner.
	PermitSignature string `json:"permit_signature,required" format:"hex_string"`
	paramObj
}

func (r Eip155OrderRequestPermitTransactionParam) MarshalJSON() (data []byte, err error) {
	type shadow Eip155OrderRequestPermitTransactionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Eip155OrderRequestPermitTransactionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token permit to be signed by the smart contract submitter.
type V2AccountOrderRequestEip155NewPermitResponse struct {
	// ID representing the EIP155 `OrderRequest`
	OrderRequestID string `json:"order_request_id,required" format:"uuid"`
	// Token permit that is to be signed by smart contract submitter for authorizing
	// token transfer for the `OrderRequest`
	Permit map[string]any `json:"permit,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrderRequestID respjson.Field
		Permit         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestEip155NewPermitResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderRequestEip155NewPermitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestEip155NewPermitTransactionResponse struct {
	// [JSON ABI](https://docs.soliditylang.org/en/v0.8.30/abi-spec.html#json) of the
	// smart contract function encoded in the transaction. Provided for informational
	// purposes.
	Abi any `json:"abi,required"`
	// Arguments to the smart contract function encoded in the transaction. Provided
	// for informational purposes.
	Args any `json:"args,required"`
	// Smart contract address that the transaction should call.
	ContractAddress string `json:"contract_address,required" format:"eth_address"`
	// Hex-encoded function call.
	Data string `json:"data,required" format:"hex_string"`
	// Transaction value estimate in Wei.
	Value string `json:"value,required" format:"bigint"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi             respjson.Field
		Args            respjson.Field
		ContractAddress respjson.Field
		Data            respjson.Field
		Value           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestEip155NewPermitTransactionResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderRequestEip155NewPermitTransactionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to create an `Order`.
//
// An `EIP155OrderRequest` is created when a user places an order through the
// Dinari API. The `EIP155OrderRequest` is then fulfilled by creating an `Order`
// on-chain.
//
// The `EIP155OrderRequest` is a record of the user's intent to place an order,
// while the `Order` is the actual transaction that occurs on the blockchain.
type V2AccountOrderRequestEip155SubmitResponse struct {
	// ID of `EIP155OrderRequest`. This is the primary identifier for the
	// `/order_requests` routes.
	ID string `json:"id,required" format:"uuid"`
	// ID of `Account` placing the `EIP155OrderRequest`.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Datetime at which the `EIP155OrderRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderSide `json:"order_side,required"`
	// Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderTif `json:"order_tif,required"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderType `json:"order_type,required"`
	// Status of `EIP155OrderRequest`.
	//
	// Any of "QUOTED", "PENDING", "PENDING_BRIDGE", "SUBMITTED", "ERROR", "CANCELLED",
	// "EXPIRED".
	Status OrderRequestStatus `json:"status,required"`
	// ID of `Order` created from the `EIP155OrderRequest`. This is the primary
	// identifier for the `/orders` routes.
	OrderID string `json:"order_id,nullable" format:"uuid"`
	// ID of recipient `Account`.
	RecipientAccountID string `json:"recipient_account_id,nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		OrderSide          respjson.Field
		OrderTif           respjson.Field
		OrderType          respjson.Field
		Status             respjson.Field
		OrderID            respjson.Field
		RecipientAccountID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestEip155SubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderRequestEip155SubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestEip155NewPermitParams struct {
	// CAIP-2 chain ID of the blockchain where the `Order` will be placed.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,omitzero,required"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderSide `json:"order_side,omitzero,required"`
	// Time in force. Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderTif `json:"order_tif,omitzero,required"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderType `json:"order_type,omitzero,required"`
	// Address of payment token.
	PaymentToken string `json:"payment_token,required" format:"eth_address"`
	// Amount of dShare asset tokens involved. Required for limit `Orders` and market
	// sell `Orders`.
	AssetTokenQuantity param.Opt[float64] `json:"asset_token_quantity,omitzero"`
	// Customer-supplied unique identifier to map this `Order` to an order in the
	// customer's systems.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Price per asset in the asset's native currency. USD for US equities and ETFs.
	// Required for limit `Orders`.
	LimitPrice param.Opt[float64] `json:"limit_price,omitzero"`
	// Amount of payment tokens involved. Required for market buy `Orders`.
	PaymentTokenQuantity param.Opt[float64] `json:"payment_token_quantity,omitzero"`
	// The ID of the `Stock` for which the `Order` is being placed.
	StockID param.Opt[string] `json:"stock_id,omitzero" format:"uuid"`
	// The ID of the `Token` for which the `Order` is being placed.
	TokenID param.Opt[string] `json:"token_id,omitzero" format:"uuid"`
	paramObj
}

func (r V2AccountOrderRequestEip155NewPermitParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderRequestEip155NewPermitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderRequestEip155NewPermitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestEip155NewPermitTransactionParams struct {
	// Input parameters for creating a proxied `EIP155OrderRequestPermitTransaction`.
	Eip155OrderRequestPermitTransaction Eip155OrderRequestPermitTransactionParam
	paramObj
}

func (r V2AccountOrderRequestEip155NewPermitTransactionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Eip155OrderRequestPermitTransaction)
}
func (r *V2AccountOrderRequestEip155NewPermitTransactionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Eip155OrderRequestPermitTransaction)
}

type V2AccountOrderRequestEip155SubmitParams struct {
	// Input parameters for creating a proxied `EIP155OrderRequestPermitTransaction`.
	Eip155OrderRequestPermitTransaction Eip155OrderRequestPermitTransactionParam
	paramObj
}

func (r V2AccountOrderRequestEip155SubmitParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Eip155OrderRequestPermitTransaction)
}
func (r *V2AccountOrderRequestEip155SubmitParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Eip155OrderRequestPermitTransaction)
}
