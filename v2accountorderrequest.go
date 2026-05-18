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
	// **`Order Requests` represent requests for Dinari to create `Orders` on behalf of
	// an `Account`.**
	//
	// `Order Requests` are created when placing **proxied orders** or **managed
	// orders**. See their respective descriptions for more details.
	Eip155 V2AccountOrderRequestEip155Service
}

// NewV2AccountOrderRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountOrderRequestService(opts ...option.RequestOption) (r V2AccountOrderRequestService) {
	r = V2AccountOrderRequestService{}
	r.Options = opts
	r.Eip155 = NewV2AccountOrderRequestEip155Service(opts...)
	return
}

// Get a specific `OrderRequest` by its ID.
func (r *V2AccountOrderRequestService) Get(ctx context.Context, orderRequestID string, query V2AccountOrderRequestGetParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if orderRequestID == "" {
		err = errors.New("missing required order_request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/%s", query.AccountID, orderRequestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists `OrderRequests`. Optionally `OrderRequests` can be filtered by certain
// parameters.
func (r *V2AccountOrderRequestService) List(ctx context.Context, accountID string, query V2AccountOrderRequestListParams, opts ...option.RequestOption) (res *V2AccountOrderRequestListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Create a managed `OrderRequest` to place a limit buy `Order`.
//
// Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN
// orders in USD, supporting up to 6 decimal places
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the creation call will fail.
func (r *V2AccountOrderRequestService) NewLimitBuy(ctx context.Context, accountID string, body V2AccountOrderRequestNewLimitBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a managed `OrderRequest` to place a limit sell `Order`.
//
// Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN
// orders in USD, supporting up to 6 decimal places
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the creation call will fail.
func (r *V2AccountOrderRequestService) NewLimitSell(ctx context.Context, accountID string, body V2AccountOrderRequestNewLimitSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a managed `OrderRequest` to place a market buy `Order`.
//
// Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN
// orders in USD, supporting up to 6 decimal places
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the creation call will fail.
func (r *V2AccountOrderRequestService) NewMarketBuy(ctx context.Context, accountID string, body V2AccountOrderRequestNewMarketBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create a managed `OrderRequest` to place a market sell `Order`.
//
// Fees for the `Order` can optionally be specified in the `OrderRequest` for DFN
// orders in USD, supporting up to 6 decimal places
//
// If an `OrderRequest` with the same `client_order_id` already exists for the
// given account, the creation call will fail.
func (r *V2AccountOrderRequestService) NewMarketSell(ctx context.Context, accountID string, body V2AccountOrderRequestNewMarketSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Input parameters for creating a limit buy `OrderRequest`.
//
// The properties AssetQuantity, LimitPrice are required.
type CreateLimitBuyOrderInputParam struct {
	// Amount of dShare asset involved. Required for limit `Order Requests` and market
	// sell `Order Requests`. Must be a positive number with a precision of up to 4
	// decimal places.
	AssetQuantity float64 `json:"asset_quantity" api:"required"`
	// Price at which to execute the order. Must be a positive number with a precision
	// of up to 2 decimal places.
	LimitPrice float64 `json:"limit_price" api:"required"`
	// ID of `Alloy`.
	AlloyID param.Opt[string] `json:"alloy_id,omitzero" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Optional fee amount associated with `Order` in USD for DFN orders. Must be a
	// positive number with a precision of up to 6 decimal places.
	Fee param.Opt[float64] `json:"fee,omitzero"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	// ID of `Stock`.
	StockID param.Opt[string] `json:"stock_id,omitzero" format:"uuid"`
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
// The properties AssetQuantity, LimitPrice are required.
type CreateLimitSellOrderInputParam struct {
	// Amount of dShare asset involved. Required for limit `Order Requests` and market
	// sell `Order Requests`. Must be a positive number with a precision of up to 4
	// decimal places.
	AssetQuantity float64 `json:"asset_quantity" api:"required"`
	// Price at which to execute the order. Must be a positive number with a precision
	// of up to 2 decimal places.
	LimitPrice float64 `json:"limit_price" api:"required"`
	// ID of `Alloy`.
	AlloyID param.Opt[string] `json:"alloy_id,omitzero" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Optional fee amount associated with `Order` in USD for DFN orders. Must be a
	// positive number with a precision of up to 6 decimal places.
	Fee param.Opt[float64] `json:"fee,omitzero"`
	// Address of the payment token to be used for the sell order. If not provided, the
	// default payment token (USD+) will be used. Should only be specified if
	// `recipient_account_id` for a non-managed wallet account is also provided.
	PaymentTokenAddress param.Opt[string] `json:"payment_token_address,omitzero" format:"eth_address"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	// ID of `Stock`.
	StockID param.Opt[string] `json:"stock_id,omitzero" format:"uuid"`
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
// The property PaymentAmount is required.
type CreateMarketBuyOrderInputParam struct {
	// Amount of currency (USD for US equities and ETFs) to pay for the order. Must be
	// a positive number with a precision of up to 2 decimal places.
	PaymentAmount float64 `json:"payment_amount" api:"required"`
	// ID of `Alloy`.
	AlloyID param.Opt[string] `json:"alloy_id,omitzero" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Optional fee amount associated with `Order` in USD for DFN orders. Must be a
	// positive number with a precision of up to 6 decimal places.
	Fee param.Opt[float64] `json:"fee,omitzero"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	// ID of `Stock`.
	StockID param.Opt[string] `json:"stock_id,omitzero" format:"uuid"`
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
// The property AssetQuantity is required.
type CreateMarketSellOrderInputParam struct {
	// Quantity of shares to trade. Must be a positive number with a precision of up to
	// 6 decimal places.
	AssetQuantity float64 `json:"asset_quantity" api:"required"`
	// ID of `Alloy`.
	AlloyID param.Opt[string] `json:"alloy_id,omitzero" format:"uuid"`
	// Customer-supplied ID to map this order to an order in their own systems. Must be
	// unique within the entity.
	ClientOrderID param.Opt[string] `json:"client_order_id,omitzero"`
	// Optional fee amount associated with `Order` in USD for DFN orders. Must be a
	// positive number with a precision of up to 6 decimal places.
	Fee param.Opt[float64] `json:"fee,omitzero"`
	// Address of the payment token to be used for the sell order. If not provided, the
	// default payment token (USD+) will be used. Should only be specified if
	// `recipient_account_id` for a non-managed wallet account is also provided.
	PaymentTokenAddress param.Opt[string] `json:"payment_token_address,omitzero" format:"eth_address"`
	// ID of `Account` to receive the `Order`.
	RecipientAccountID param.Opt[string] `json:"recipient_account_id,omitzero" format:"uuid"`
	// ID of `Stock`.
	StockID param.Opt[string] `json:"stock_id,omitzero" format:"uuid"`
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
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of `Account` placing the `OrderRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `OrderRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderSide `json:"order_side" api:"required"`
	// Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderTif `json:"order_tif" api:"required"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderType `json:"order_type" api:"required"`
	// Status of `OrderRequest`. Possible values:
	//
	// - `QUOTED`: Order request created with fee quote provided, ready for processing
	// - `PENDING`: Order request is being prepared for submission
	// - `PENDING_BRIDGE`: Order is waiting for bridge transaction to complete
	// - `SUBMITTED`: Order has been successfully submitted to the order book
	// - `ERROR`: An error occurred during order processing
	// - `CANCELLED`: Order request was cancelled
	// - `EXPIRED`: Order request expired due to deadline passing
	// - `REJECTED`: Order request was rejected
	//
	// Any of "QUOTED", "PENDING", "PENDING_BRIDGE", "SUBMITTED", "ERROR", "CANCELLED",
	// "EXPIRED", "REJECTED".
	Status OrderRequestStatus `json:"status" api:"required"`
	// Reason for the order cancellation if the order status is CANCELLED
	CancelMessage string `json:"cancel_message" api:"nullable"`
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID string `json:"client_order_id" api:"nullable"`
	// ID of `Order` created from the `OrderRequest`. This is the primary identifier
	// for the `/orders` routes.
	OrderID string `json:"order_id" api:"nullable" format:"uuid"`
	// ID of recipient `Account`.
	RecipientAccountID string `json:"recipient_account_id" api:"nullable" format:"uuid"`
	// Reason for the order rejection if the order status is REJECTED
	RejectMessage string `json:"reject_message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		OrderSide          respjson.Field
		OrderTif           respjson.Field
		OrderType          respjson.Field
		Status             respjson.Field
		CancelMessage      respjson.Field
		ClientOrderID      respjson.Field
		OrderID            respjson.Field
		RecipientAccountID respjson.Field
		RejectMessage      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderRequest) RawJSON() string { return r.JSON.raw }
func (r *OrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderRequestStatus string

const (
	OrderRequestStatusQuoted        OrderRequestStatus = "QUOTED"
	OrderRequestStatusPending       OrderRequestStatus = "PENDING"
	OrderRequestStatusPendingBridge OrderRequestStatus = "PENDING_BRIDGE"
	OrderRequestStatusSubmitted     OrderRequestStatus = "SUBMITTED"
	OrderRequestStatusError         OrderRequestStatus = "ERROR"
	OrderRequestStatusCancelled     OrderRequestStatus = "CANCELLED"
	OrderRequestStatusExpired       OrderRequestStatus = "EXPIRED"
	OrderRequestStatusRejected      OrderRequestStatus = "REJECTED"
)

// V2AccountOrderRequestListResponseUnion contains all possible properties and
// values from [[]V2AccountOrderRequestListResponseArrayItem],
// [V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfV2AccountOrderRequestListResponseArray]
type V2AccountOrderRequestListResponseUnion struct {
	// This field will be present if the value is a
	// [[]V2AccountOrderRequestListResponseArrayItem] instead of an object.
	OfV2AccountOrderRequestListResponseArray []V2AccountOrderRequestListResponseArrayItem `json:",inline"`
	// This field is from variant
	// [V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse].
	Data []V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponseData `json:"data"`
	// This field is from variant
	// [V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse].
	PaginationMetadata V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponsePaginationMetadata `json:"pagination_metadata"`
	// This field is from variant
	// [V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse].
	Sv   string `json:"_sv"`
	JSON struct {
		OfV2AccountOrderRequestListResponseArray respjson.Field
		Data                                     respjson.Field
		PaginationMetadata                       respjson.Field
		Sv                                       respjson.Field
		raw                                      string
	} `json:"-"`
}

func (u V2AccountOrderRequestListResponseUnion) AsV2AccountOrderRequestListResponseArray() (v []V2AccountOrderRequestListResponseArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u V2AccountOrderRequestListResponseUnion) AsV2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse() (v V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u V2AccountOrderRequestListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *V2AccountOrderRequestListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to create an `Order`.
//
// An `OrderRequest` is created when a user places an order through the Dinari API.
// The `OrderRequest` is then fulfilled by creating an `Order` on-chain.
//
// The `OrderRequest` is a record of the user's intent to place an order, while the
// `Order` is the actual transaction that occurs on the blockchain.
type V2AccountOrderRequestListResponseArrayItem struct {
	// ID of `OrderRequest`. This is the primary identifier for the `/order_requests`
	// routes.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of `Account` placing the `OrderRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `OrderRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide string `json:"order_side" api:"required"`
	// Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif string `json:"order_tif" api:"required"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType string `json:"order_type" api:"required"`
	// Status of `OrderRequest`. Possible values:
	//
	// - `QUOTED`: Order request created with fee quote provided, ready for processing
	// - `PENDING`: Order request is being prepared for submission
	// - `PENDING_BRIDGE`: Order is waiting for bridge transaction to complete
	// - `SUBMITTED`: Order has been successfully submitted to the order book
	// - `ERROR`: An error occurred during order processing
	// - `CANCELLED`: Order request was cancelled
	// - `EXPIRED`: Order request expired due to deadline passing
	// - `REJECTED`: Order request was rejected
	//
	// Any of "QUOTED", "PENDING", "PENDING_BRIDGE", "SUBMITTED", "ERROR", "CANCELLED",
	// "EXPIRED", "REJECTED".
	Status string `json:"status" api:"required"`
	// Reason for the order cancellation if the order status is CANCELLED
	CancelMessage string `json:"cancel_message" api:"nullable"`
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID string `json:"client_order_id" api:"nullable"`
	// ID of `Order` created from the `OrderRequest`. This is the primary identifier
	// for the `/orders` routes.
	OrderID string `json:"order_id" api:"nullable" format:"uuid"`
	// ID of recipient `Account`.
	RecipientAccountID string `json:"recipient_account_id" api:"nullable" format:"uuid"`
	// Reason for the order rejection if the order status is REJECTED
	RejectMessage string `json:"reject_message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		OrderSide          respjson.Field
		OrderTif           respjson.Field
		OrderType          respjson.Field
		Status             respjson.Field
		CancelMessage      respjson.Field
		ClientOrderID      respjson.Field
		OrderID            respjson.Field
		RecipientAccountID respjson.Field
		RejectMessage      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestListResponseArrayItem) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderRequestListResponseArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse struct {
	// List of AccountOrder
	Data []V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponseData `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponsePaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedAccountOrderRequestResponse:v1".
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
func (r V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to create an `Order`.
//
// An `OrderRequest` is created when a user places an order through the Dinari API.
// The `OrderRequest` is then fulfilled by creating an `Order` on-chain.
//
// The `OrderRequest` is a record of the user's intent to place an order, while the
// `Order` is the actual transaction that occurs on the blockchain.
type V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponseData struct {
	// ID of `OrderRequest`. This is the primary identifier for the `/order_requests`
	// routes.
	ID string `json:"id" api:"required" format:"uuid"`
	// ID of `Account` placing the `OrderRequest`.
	AccountID string `json:"account_id" api:"required" format:"uuid"`
	// Datetime at which the `OrderRequest` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt" api:"required" format:"date-time"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide string `json:"order_side" api:"required"`
	// Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif string `json:"order_tif" api:"required"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType string `json:"order_type" api:"required"`
	// Status of `OrderRequest`. Possible values:
	//
	// - `QUOTED`: Order request created with fee quote provided, ready for processing
	// - `PENDING`: Order request is being prepared for submission
	// - `PENDING_BRIDGE`: Order is waiting for bridge transaction to complete
	// - `SUBMITTED`: Order has been successfully submitted to the order book
	// - `ERROR`: An error occurred during order processing
	// - `CANCELLED`: Order request was cancelled
	// - `EXPIRED`: Order request expired due to deadline passing
	// - `REJECTED`: Order request was rejected
	//
	// Any of "QUOTED", "PENDING", "PENDING_BRIDGE", "SUBMITTED", "ERROR", "CANCELLED",
	// "EXPIRED", "REJECTED".
	Status string `json:"status" api:"required"`
	// Reason for the order cancellation if the order status is CANCELLED
	CancelMessage string `json:"cancel_message" api:"nullable"`
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID string `json:"client_order_id" api:"nullable"`
	// ID of `Order` created from the `OrderRequest`. This is the primary identifier
	// for the `/orders` routes.
	OrderID string `json:"order_id" api:"nullable" format:"uuid"`
	// ID of recipient `Account`.
	RecipientAccountID string `json:"recipient_account_id" api:"nullable" format:"uuid"`
	// Reason for the order rejection if the order status is REJECTED
	RejectMessage string `json:"reject_message" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccountID          respjson.Field
		CreatedDt          respjson.Field
		OrderSide          respjson.Field
		OrderTif           respjson.Field
		OrderType          respjson.Field
		Status             respjson.Field
		CancelMessage      respjson.Field
		ClientOrderID      respjson.Field
		OrderID            respjson.Field
		RecipientAccountID respjson.Field
		RejectMessage      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponseData) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponsePaginationMetadata struct {
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
func (r V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponsePaginationMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderRequestListResponsePaginatedAccountOrderRequestResponsePaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderRequestGetParams struct {
	AccountID string `path:"account_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountOrderRequestListParams struct {
	// Customer-supplied ID to map this `OrderRequest` to an order in their own
	// systems.
	ClientOrderID param.Opt[string] `query:"client_order_id,omitzero" json:"-"`
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Order ID for the `OrderRequest`
	OrderID param.Opt[string] `query:"order_id,omitzero" format:"uuid" json:"-"`
	// Order Request ID for the `OrderRequest`
	OrderRequestID param.Opt[string] `query:"order_request_id,omitzero" format:"uuid" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Number of results to return
	Limit    param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2AccountOrderRequestListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountOrderRequestListParams]'s query parameters as
// `url.Values`.
func (r V2AccountOrderRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2AccountOrderRequestListParamsOrder string

const (
	V2AccountOrderRequestListParamsOrderAsc  V2AccountOrderRequestListParamsOrder = "asc"
	V2AccountOrderRequestListParamsOrderDesc V2AccountOrderRequestListParamsOrder = "desc"
)

type V2AccountOrderRequestNewLimitBuyParams struct {
	// Input parameters for creating a limit buy `OrderRequest`.
	CreateLimitBuyOrderInput CreateLimitBuyOrderInputParam
	paramObj
}

func (r V2AccountOrderRequestNewLimitBuyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateLimitBuyOrderInput)
}
func (r *V2AccountOrderRequestNewLimitBuyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	return apijson.UnmarshalRoot(data, r)
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
	return apijson.UnmarshalRoot(data, r)
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
	return apijson.UnmarshalRoot(data, r)
}
