// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// V2AccountOrderRequestStockEip155Service contains methods and other services that
// help with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderRequestStockEip155Service] method instead.
type V2AccountOrderRequestStockEip155Service struct {
	Options []option.RequestOption
}

// NewV2AccountOrderRequestStockEip155Service generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV2AccountOrderRequestStockEip155Service(opts ...option.RequestOption) (r V2AccountOrderRequestStockEip155Service) {
	r = V2AccountOrderRequestStockEip155Service{}
	r.Options = opts
	return
}

// Prepare a proxied order to be placed on EVM. The returned structure contains the
// necessary data to create an `OrderRequest` with a `Wallet` that is not
// Dinari-managed.
//
// **⚠️ This endpoint will be deprecated on 2025-12-15.**
//
// Deprecated: deprecated
func (r *V2AccountOrderRequestStockEip155Service) PrepareProxiedOrder(ctx context.Context, accountID string, body V2AccountOrderRequestStockEip155PrepareProxiedOrderParams, opts ...option.RequestOption) (res *V2AccountOrderRequestStockEip155PrepareProxiedOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/stocks/eip155/prepare", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [EIP-712](https://eips.ethereum.org/EIPS/eip-712) typed data to be signed with a
// wallet.
type EvmTypedData struct {
	// Domain separator for the typed data.
	Domain any `json:"domain,required"`
	// Message to be signed. Contains the actual data that will be signed with the
	// wallet.
	Message any `json:"message,required"`
	// Primary type of the typed data.
	PrimaryType string `json:"primaryType,required"`
	// Types used in the typed data.
	Types any `json:"types,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Message     respjson.Field
		PrimaryType respjson.Field
		Types       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvmTypedData) RawJSON() string { return r.JSON.raw }
func (r *EvmTypedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prepared data for creating an `OrderRequest` through the EVM proxied order API
// route.
type V2AccountOrderRequestStockEip155PrepareProxiedOrderResponse struct {
	// ID of the EvmPreparedProxiedOrder.
	ID string `json:"id,required" format:"uuid"`
	// Deadline for the prepared order to be placed.
	Deadline time.Time `json:"deadline,required" format:"date-time"`
	// Fees involved in the order. Provided here as a reference.
	Fees []OrderFeeAmount `json:"fees,required"`
	// [EIP-712](https://eips.ethereum.org/EIPS/eip-712) typed data to be signed with a
	// wallet.
	OrderTypedData EvmTypedData `json:"order_typed_data,required"`
	// [EIP-712](https://eips.ethereum.org/EIPS/eip-712) typed data to be signed with a
	// wallet.
	PermitTypedData EvmTypedData `json:"permit_typed_data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Deadline        respjson.Field
		Fees            respjson.Field
		OrderTypedData  respjson.Field
		PermitTypedData respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestStockEip155PrepareProxiedOrderResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderRequestStockEip155PrepareProxiedOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestStockEip155PrepareProxiedOrderParams struct {
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

func (r V2AccountOrderRequestStockEip155PrepareProxiedOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderRequestStockEip155PrepareProxiedOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderRequestStockEip155PrepareProxiedOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
