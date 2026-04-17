// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"encoding/json"
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

// **`Orders` represent the buying and selling of assets under an `Account`.**
//
// For `Accounts` using self-custodied `Wallets`, `Orders` are created and
// fulfilled by making calls to Dinari's smart contracts, or using the _Proxied
// Orders_ methods.
//
// For `Accounts` using managed `Wallets`, `Orders` are created and fulfilled by
// using the `Managed Orders` methods, which then create the corresponding
// transactions on the blockchain.
//
// V2Service contains methods and other services that help with interacting with
// the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2Service] method instead.
type V2Service struct {
	Options []option.RequestOption
	// **Dinari provides basic market data for `Stocks` and `Alloys` that are available
	// to transact on.**
	//
	// This data is provided on a best-effort basis and we recommend using a dedicated
	// provider for more intensive market data needs.
	MarketData V2MarketDataService
	// **`Entities` represent a business or organization that uses the API, and their
	// customers.**
	//
	// Dinari Partners are represented as an organization `Entity` in the API, with
	// their own accounts. Individual customers of Partner `Entities` are also
	// represented as `Entities` in the API, which are managed by the Partner `Entity`.
	Entities V2EntityService
	// **`Accounts` represent the financial accounts of an `Entity`.**
	//
	// `Orders`, dividends, and other transactions are associated with an `Account`.
	Accounts V2AccountService
}

// NewV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2Service(opts ...option.RequestOption) (r V2Service) {
	r = V2Service{}
	r.Options = opts
	r.MarketData = NewV2MarketDataService(opts...)
	r.Entities = NewV2EntityService(opts...)
	r.Accounts = NewV2AccountService(opts...)
	return
}

// Get a list of all `Orders` under the `Entity`. Optionally `Orders` can be
// transaction hash or fulfillment transaction hash.
func (r *V2Service) ListOrders(ctx context.Context, query V2ListOrdersParams, opts ...option.RequestOption) (res *V2ListOrdersResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/orders/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// V2ListOrdersResponseUnion contains all possible properties and values from
// [[]V2ListOrdersResponseArrayItem],
// [V2ListOrdersResponsePaginatedEntityOrderResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfV2ListOrdersResponseArray]
type V2ListOrdersResponseUnion struct {
	// This field will be present if the value is a [[]V2ListOrdersResponseArrayItem]
	// instead of an object.
	OfV2ListOrdersResponseArray []V2ListOrdersResponseArrayItem `json:",inline"`
	// This field is from variant [V2ListOrdersResponsePaginatedEntityOrderResponse].
	Data []V2ListOrdersResponsePaginatedEntityOrderResponseData `json:"data"`
	// This field is from variant [V2ListOrdersResponsePaginatedEntityOrderResponse].
	PaginationMetadata V2ListOrdersResponsePaginatedEntityOrderResponsePaginationMetadata `json:"pagination_metadata"`
	// This field is from variant [V2ListOrdersResponsePaginatedEntityOrderResponse].
	Sv   string `json:"_sv"`
	JSON struct {
		OfV2ListOrdersResponseArray respjson.Field
		Data                        respjson.Field
		PaginationMetadata          respjson.Field
		Sv                          respjson.Field
		raw                         string
	} `json:"-"`
}

func (u V2ListOrdersResponseUnion) AsV2ListOrdersResponseArray() (v []V2ListOrdersResponseArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2ListOrdersResponseUnion) AsV2ListOrdersResponsePaginatedEntityOrderResponse() (v V2ListOrdersResponsePaginatedEntityOrderResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2ListOrdersResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2ListOrdersResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ListOrdersResponseArrayItem struct {
	// ID of the `Order`.
	ID string `json:"id" api:"required" format:"uuid"`
	// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
	// on.
	ChainID string `json:"chain_id" api:"required"`
	// Datetime at which the `Order` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Smart contract address that `Order` was created from.
	OrderContractAddress string `json:"order_contract_address" api:"required" format:"eth_address"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide string `json:"order_side" api:"required"`
	// Time in force. Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif string `json:"order_tif" api:"required"`
	// Transaction hash for the `Order` creation.
	OrderTransactionHash string `json:"order_transaction_hash" api:"required" format:"hex_string"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType string `json:"order_type" api:"required"`
	// The payment token (stablecoin) address.
	PaymentToken string `json:"payment_token" api:"required" format:"eth_address"`
	// Status of the `Order`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "PARTIALLY_FILLED", "FILLED", "REJECTED",
	// "REQUIRING_CONTACT", "ERROR".
	Status string `json:"status" api:"required"`
	// The `Stock` ID associated with the `Order`
	StockID string `json:"stock_id" api:"required" format:"uuid"`
	// Account ID the order was made for.
	AccountID string `json:"account_id" api:"nullable" format:"uuid"`
	// The dShare asset token address.
	AssetToken string `json:"asset_token" api:"nullable" format:"eth_address"`
	// Total amount of assets involved.
	AssetTokenQuantity float64 `json:"asset_token_quantity" api:"nullable"`
	// Transaction hash for cancellation of `Order`, if the `Order` was cancelled.
	CancelTransactionHash string `json:"cancel_transaction_hash" api:"nullable" format:"hex_string"`
	// Customer-supplied unique identifier to map this `Order` to an order in the
	// customer's systems.
	ClientOrderID string `json:"client_order_id" api:"nullable"`
	// Entity ID of the Order
	EntityID string `json:"entity_id" api:"nullable" format:"uuid"`
	// Fee amount associated with `Order`.
	Fee float64 `json:"fee" api:"nullable"`
	// For limit `Orders`, the price per asset, specified in the `Stock`'s native
	// currency (USD for US equities and ETFs).
	LimitPrice float64 `json:"limit_price" api:"nullable"`
	// Order Request ID for the `Order`
	OrderRequestID string `json:"order_request_id" api:"nullable" format:"uuid"`
	// Total amount of payment involved.
	PaymentTokenQuantity float64 `json:"payment_token_quantity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ChainID               respjson.Field
		CreatedDt             respjson.Field
		OrderContractAddress  respjson.Field
		OrderSide             respjson.Field
		OrderTif              respjson.Field
		OrderTransactionHash  respjson.Field
		OrderType             respjson.Field
		PaymentToken          respjson.Field
		Status                respjson.Field
		StockID               respjson.Field
		AccountID             respjson.Field
		AssetToken            respjson.Field
		AssetTokenQuantity    respjson.Field
		CancelTransactionHash respjson.Field
		ClientOrderID         respjson.Field
		EntityID              respjson.Field
		Fee                   respjson.Field
		LimitPrice            respjson.Field
		OrderRequestID        respjson.Field
		PaymentTokenQuantity  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ListOrdersResponseArrayItem) RawJSON() string { return r.JSON.raw }
func (r *V2ListOrdersResponseArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ListOrdersResponsePaginatedEntityOrderResponse struct {
	// List of EntityOrder
	Data []V2ListOrdersResponsePaginatedEntityOrderResponseData `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata V2ListOrdersResponsePaginatedEntityOrderResponsePaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedEntityOrderResponse:v1".
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
func (r V2ListOrdersResponsePaginatedEntityOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ListOrdersResponsePaginatedEntityOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ListOrdersResponsePaginatedEntityOrderResponseData struct {
	// ID of the `Order`.
	ID string `json:"id" api:"required" format:"uuid"`
	// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
	// on.
	ChainID string `json:"chain_id" api:"required"`
	// Datetime at which the `Order` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Smart contract address that `Order` was created from.
	OrderContractAddress string `json:"order_contract_address" api:"required" format:"eth_address"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide string `json:"order_side" api:"required"`
	// Time in force. Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif string `json:"order_tif" api:"required"`
	// Transaction hash for the `Order` creation.
	OrderTransactionHash string `json:"order_transaction_hash" api:"required" format:"hex_string"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType string `json:"order_type" api:"required"`
	// The payment token (stablecoin) address.
	PaymentToken string `json:"payment_token" api:"required" format:"eth_address"`
	// Status of the `Order`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "PARTIALLY_FILLED", "FILLED", "REJECTED",
	// "REQUIRING_CONTACT", "ERROR".
	Status string `json:"status" api:"required"`
	// The `Stock` ID associated with the `Order`
	StockID string `json:"stock_id" api:"required" format:"uuid"`
	// Account ID the order was made for.
	AccountID string `json:"account_id" api:"nullable" format:"uuid"`
	// The dShare asset token address.
	AssetToken string `json:"asset_token" api:"nullable" format:"eth_address"`
	// Total amount of assets involved.
	AssetTokenQuantity float64 `json:"asset_token_quantity" api:"nullable"`
	// Transaction hash for cancellation of `Order`, if the `Order` was cancelled.
	CancelTransactionHash string `json:"cancel_transaction_hash" api:"nullable" format:"hex_string"`
	// Customer-supplied unique identifier to map this `Order` to an order in the
	// customer's systems.
	ClientOrderID string `json:"client_order_id" api:"nullable"`
	// Entity ID of the Order
	EntityID string `json:"entity_id" api:"nullable" format:"uuid"`
	// Fee amount associated with `Order`.
	Fee float64 `json:"fee" api:"nullable"`
	// For limit `Orders`, the price per asset, specified in the `Stock`'s native
	// currency (USD for US equities and ETFs).
	LimitPrice float64 `json:"limit_price" api:"nullable"`
	// Order Request ID for the `Order`
	OrderRequestID string `json:"order_request_id" api:"nullable" format:"uuid"`
	// Total amount of payment involved.
	PaymentTokenQuantity float64 `json:"payment_token_quantity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		ChainID               respjson.Field
		CreatedDt             respjson.Field
		OrderContractAddress  respjson.Field
		OrderSide             respjson.Field
		OrderTif              respjson.Field
		OrderTransactionHash  respjson.Field
		OrderType             respjson.Field
		PaymentToken          respjson.Field
		Status                respjson.Field
		StockID               respjson.Field
		AccountID             respjson.Field
		AssetToken            respjson.Field
		AssetTokenQuantity    respjson.Field
		CancelTransactionHash respjson.Field
		ClientOrderID         respjson.Field
		EntityID              respjson.Field
		Fee                   respjson.Field
		LimitPrice            respjson.Field
		OrderRequestID        respjson.Field
		PaymentTokenQuantity  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2ListOrdersResponsePaginatedEntityOrderResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2ListOrdersResponsePaginatedEntityOrderResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type V2ListOrdersResponsePaginatedEntityOrderResponsePaginationMetadata struct {
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
func (r V2ListOrdersResponsePaginatedEntityOrderResponsePaginationMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *V2ListOrdersResponsePaginatedEntityOrderResponsePaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ListOrdersParams struct {
	// CAIP-2 formatted chain ID of the blockchain the `Order` was made on.
	ChainID param.Opt[string] `query:"chain_id,omitzero" json:"-"`
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Fulfillment transaction hash of the `Order`.
	OrderFulfillmentTransactionHash param.Opt[string] `query:"order_fulfillment_transaction_hash,omitzero" json:"-"`
	// Order Request ID for the `Order`
	OrderRequestID param.Opt[string] `query:"order_request_id,omitzero" format:"uuid" json:"-"`
	// Transaction hash of the `Order`.
	OrderTransactionHash param.Opt[string] `query:"order_transaction_hash,omitzero" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Number of results to return
	Limit    param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2ListOrdersParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2ListOrdersParams]'s query parameters as `url.Values`.
func (r V2ListOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2ListOrdersParamsOrder string

const (
	V2ListOrdersParamsOrderAsc  V2ListOrdersParamsOrder = "asc"
	V2ListOrdersParamsOrderDesc V2ListOrdersParamsOrder = "desc"
)
