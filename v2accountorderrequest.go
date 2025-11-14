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
	shimjson "github.com/dinaricrypto/dinari-api-sdk-go/internal/encoding/json"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// V2AccountOrderRequestService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderRequestService] method instead.
type V2AccountOrderRequestService struct {
	Options []option.RequestOption
	Stocks  V2AccountOrderRequestStockService
	Eip155  V2AccountOrderRequestEip155Service
}

// NewV2AccountOrderRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountOrderRequestService(opts ...option.RequestOption) (r V2AccountOrderRequestService) {
	r = V2AccountOrderRequestService{}
	r.Options = opts
	r.Stocks = NewV2AccountOrderRequestStockService(opts...)
	r.Eip155 = NewV2AccountOrderRequestEip155Service(opts...)
	return
}

// Get a specific `OrderRequest` by its ID.
func (r *V2AccountOrderRequestService) Get(ctx context.Context, orderRequestID string, query V2AccountOrderRequestGetParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderRequestID == "" {
		err = errors.New("missing required order_request_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/%s", query.AccountID, orderRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists `OrderRequests`. Optionally `OrderRequests` can be filtered by certain
// parameters.
func (r *V2AccountOrderRequestService) List(ctx context.Context, accountID string, query V2AccountOrderRequestListParams, opts ...option.RequestOption) (res *[]OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create a managed `OrderRequest` to place a limit buy `Order`.
//
// Fees for the `Order` are included in the transaction. Refer to our
// [Fee Quote API](https://docs.dinari.com/reference/createproxiedorderfeequote#/)
// for fee estimation.
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the existing `OrderRequest` will be returned instead of creating
// a new one.
func (r *V2AccountOrderRequestService) NewLimitBuy(ctx context.Context, accountID string, body V2AccountOrderRequestNewLimitBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a managed `OrderRequest` to place a limit sell `Order`.
//
// Fees for the `Order` are included in the transaction. Refer to our
// [Fee Quote API](https://docs.dinari.com/reference/createproxiedorderfeequote#/)
// for fee estimation.
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the existing `OrderRequest` will be returned instead of creating
// a new one.
func (r *V2AccountOrderRequestService) NewLimitSell(ctx context.Context, accountID string, body V2AccountOrderRequestNewLimitSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a managed `OrderRequest` to place a market buy `Order`.
//
// Fees for the `Order` are included in the transaction. Refer to our
// [Fee Quote API](https://docs.dinari.com/reference/createproxiedorderfeequote#/)
// for fee estimation.
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the existing `OrderRequest` will be returned instead of creating
// a new one.
func (r *V2AccountOrderRequestService) NewMarketBuy(ctx context.Context, accountID string, body V2AccountOrderRequestNewMarketBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a managed `OrderRequest` to place a market sell `Order`.
//
// Fees for the `Order` are included in the transaction. Refer to our
// [Fee Quote API](https://docs.dinari.com/reference/createproxiedorderfeequote#/)
// for fee estimation.
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the existing `OrderRequest` will be returned instead of creating
// a new one.
func (r *V2AccountOrderRequestService) NewMarketSell(ctx context.Context, accountID string, body V2AccountOrderRequestNewMarketSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get fee quote data for an `Order Request`. This is provided primarily for
// informational purposes.
//
// For market buy orders, the notional amount of the order includes the fees. For
// market and limit sell orders, fees are deducted from the proceeds of the sale.
// For limit buy orders, the fees are added to the total cost of the order.
func (r *V2AccountOrderRequestService) GetFeeQuote(ctx context.Context, accountID string, body V2AccountOrderRequestGetFeeQuoteParams, opts ...option.RequestOption) (res *V2AccountOrderRequestGetFeeQuoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/fee_quote", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Input parameters for creating a limit buy `OrderRequest`.
//
// The properties AssetQuantity, LimitPrice, StockID are required.
type CreateLimitBuyOrderInputParam struct {
	// Amount of dShare asset involved. Required for limit `Orders` and market sell
	// `Orders`.
	AssetQuantity float64 `json:"asset_quantity,required"`
	// Price at which to execute the order. Must be a positive number with a precision
	// of up to 2 decimal places.
	LimitPrice float64 `json:"limit_price,required"`
	// ID of `Stock`.
	StockID string `json:"stock_id,required" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	paramObj
}

func (r CreateLimitBuyOrderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateLimitBuyOrderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateLimitBuyOrderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input parameters for creating a limit sell `OrderRequest`.
//
// The properties AssetQuantity, LimitPrice, StockID are required.
type CreateLimitSellOrderInputParam struct {
	// Amount of dShare asset involved. Required for limit `Orders` and market sell
	// `Orders`.
	AssetQuantity float64 `json:"asset_quantity,required"`
	// Price at which to execute the order. Must be a positive number with a precision
	// of up to 2 decimal places.
	LimitPrice float64 `json:"limit_price,required"`
	// ID of `Stock`.
	StockID string `json:"stock_id,required" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Address of the payment token to be used for the sell order. If not provided, the
	// default payment token (USD+) will be used. Should only be specified if
	// `recipient_account_id` for a non-managed wallet account is also provided.
	PaymentTokenAddress param.Opt[string] `json:"payment_token_address,omitzero" format:"eth_address"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	paramObj
}

func (r CreateLimitSellOrderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateLimitSellOrderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateLimitSellOrderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input parameters for creating a market buy `OrderRequest`.
//
// The properties PaymentAmount, StockID are required.
type CreateMarketBuyOrderInputParam struct {
	// Amount of currency (USD for US equities and ETFs) to pay for the order. Must be
	// a positive number with a precision of up to 2 decimal places.
	PaymentAmount float64 `json:"payment_amount,required"`
	// ID of `Stock`.
	StockID string `json:"stock_id,required" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	paramObj
}

func (r CreateMarketBuyOrderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMarketBuyOrderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMarketBuyOrderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input parameters for creating a market sell `OrderRequest`.
//
// The properties AssetQuantity, StockID are required.
type CreateMarketSellOrderInputParam struct {
	// Quantity of shares to trade. Must be a positive number with a precision of up to
	// 9 decimal places.
	AssetQuantity float64 `json:"asset_quantity,required"`
	// ID of `Stock`.
	StockID string `json:"stock_id,required" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Address of the payment token to be used for the sell order. If not provided, the
	// default payment token (USD+) will be used. Should only be specified if
	// `recipient_account_id` for a non-managed wallet account is also provided.
	PaymentTokenAddress param.Opt[string] `json:"payment_token_address,omitzero" format:"eth_address"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	paramObj
}

func (r CreateMarketSellOrderInputParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateMarketSellOrderInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateMarketSellOrderInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to create an `Order`.
//
// An `OrderRequest` is created when a user places an order through the Dinari API.
// The `OrderRequest` is then fulfilled by creating an `Order` on-chain.
//
// The `OrderRequest` is a record of the user's intent to place an order, while the
// `Order` is the actual transaction that occurs on the blockchain.
type OrderRequest struct {
	// ID of `OrderRequest`. This is the primary identifier for the `/order_requests`
	// routes.
	ID string `json:"id,required" format:"uuid"`
	// ID of `Account` placing the `OrderRequest`.
	AccountID string `json:"account_id,required" format:"uuid"`
	// Datetime at which the `OrderRequest` was created. ISO 8601 timestamp.
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
	// Status of `OrderRequest`.
	//
	// Any of "QUOTED", "PENDING", "PENDING_BRIDGE", "SUBMITTED", "ERROR", "CANCELLED",
	// "EXPIRED".
	Status OrderRequestStatus `json:"status,required"`
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID string `json:"client_order_id,nullable"`
	// ID of `Order` created from the `OrderRequest`. This is the primary identifier
	// for the `/orders` routes.
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
		ClientOrderID      respjson.Field
		OrderID            respjson.Field
		RecipientAccountID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderRequest) RawJSON() string { return r.JSON.raw }
func (r *OrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of `OrderRequest`.
type OrderRequestStatus string

const (
	OrderRequestStatusQuoted        OrderRequestStatus = "QUOTED"
	OrderRequestStatusPending       OrderRequestStatus = "PENDING"
	OrderRequestStatusPendingBridge OrderRequestStatus = "PENDING_BRIDGE"
	OrderRequestStatusSubmitted     OrderRequestStatus = "SUBMITTED"
	OrderRequestStatusError         OrderRequestStatus = "ERROR"
	OrderRequestStatusCancelled     OrderRequestStatus = "CANCELLED"
	OrderRequestStatusExpired       OrderRequestStatus = "EXPIRED"
)

// A preview of the fee that would be collected when placing an Order Request.
type V2AccountOrderRequestGetFeeQuoteResponse struct {
	// Cash amount in USD paid for fees for the Order Request.
	Fee float64 `json:"fee,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fee         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestGetFeeQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderRequestGetFeeQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountOrderRequestListParams struct {
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID param.Opt[string] `query:"client_order_id,omitzero" json:"-"`
	// Order ID for the `OrderRequest`
	OrderID param.Opt[string] `query:"order_id,omitzero" format:"uuid" json:"-"`
	// Order Request ID for the `OrderRequest`
	OrderRequestID param.Opt[string] `query:"order_request_id,omitzero" format:"uuid" json:"-"`
	Page           param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize       param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountOrderRequestListParams]'s query parameters as
// `url.Values`.
func (r V2AccountOrderRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2AccountOrderRequestNewLimitBuyParams struct {
	// Input parameters for creating a limit buy `OrderRequest`.
	CreateLimitBuyOrderInput CreateLimitBuyOrderInputParam
	paramObj
}

func (r V2AccountOrderRequestNewLimitBuyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateLimitBuyOrderInput)
}
func (r *V2AccountOrderRequestNewLimitBuyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateLimitBuyOrderInput)
}

type V2AccountOrderRequestNewLimitSellParams struct {
	// Input parameters for creating a limit sell `OrderRequest`.
	CreateLimitSellOrderInput CreateLimitSellOrderInputParam
	paramObj
}

func (r V2AccountOrderRequestNewLimitSellParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateLimitSellOrderInput)
}
func (r *V2AccountOrderRequestNewLimitSellParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateLimitSellOrderInput)
}

type V2AccountOrderRequestNewMarketBuyParams struct {
	// Input parameters for creating a market buy `OrderRequest`.
	CreateMarketBuyOrderInput CreateMarketBuyOrderInputParam
	paramObj
}

func (r V2AccountOrderRequestNewMarketBuyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMarketBuyOrderInput)
}
func (r *V2AccountOrderRequestNewMarketBuyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateMarketBuyOrderInput)
}

type V2AccountOrderRequestNewMarketSellParams struct {
	// Input parameters for creating a market sell `OrderRequest`.
	CreateMarketSellOrderInput CreateMarketSellOrderInputParam
	paramObj
}

func (r V2AccountOrderRequestNewMarketSellParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateMarketSellOrderInput)
}
func (r *V2AccountOrderRequestNewMarketSellParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateMarketSellOrderInput)
}

type V2AccountOrderRequestGetFeeQuoteParams struct {
	// Indicates whether `Order Request` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderSide `json:"order_side,omitzero,required"`
	// Type of `Order Request`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderType `json:"order_type,omitzero,required"`
	// The Stock ID associated with the Order Request
	StockID string `json:"stock_id,required" format:"uuid"`
	// Amount of dShare asset tokens involved. Required for limit `Orders` and market
	// sell `Order Requests`.
	AssetTokenQuantity param.Opt[float64] `json:"asset_token_quantity,omitzero"`
	// Price per asset in the asset's native currency. USD for US equities and ETFs.
	// Required for limit `Order Requests`.
	LimitPrice param.Opt[float64] `json:"limit_price,omitzero"`
	// Address of the payment token to be used for an order. If not provided, the
	// default payment token (USD+) will be used.
	PaymentTokenAddress param.Opt[string] `json:"payment_token_address,omitzero" format:"eth_address"`
	// Amount of payment tokens involved. Required for market buy `Order Requests`.
	PaymentTokenQuantity param.Opt[float64] `json:"payment_token_quantity,omitzero"`
	// CAIP-2 chain ID of the blockchain where the `Order Request` will be placed. If
	// not provided, the default chain ID (eip155:42161) will be used.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,omitzero"`
	paramObj
}

func (r V2AccountOrderRequestGetFeeQuoteParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderRequestGetFeeQuoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderRequestGetFeeQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
