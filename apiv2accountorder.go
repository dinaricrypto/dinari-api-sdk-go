// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apiquery"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// APIV2AccountOrderService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountOrderService] method instead.
type APIV2AccountOrderService struct {
	Options []option.RequestOption
}

// NewAPIV2AccountOrderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV2AccountOrderService(opts ...option.RequestOption) (r APIV2AccountOrderService) {
	r = APIV2AccountOrderService{}
	r.Options = opts
	return
}

// Get a specific `Order` by its ID.
func (r *APIV2AccountOrderService) Get(ctx context.Context, orderID string, query APIV2AccountOrderGetParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/%s", query.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all `Orders` under the `Account`.
func (r *APIV2AccountOrderService) List(ctx context.Context, accountID string, query APIV2AccountOrderListParams, opts ...option.RequestOption) (res *[]Order, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel an `Order` by its ID. Note that this requires the `Order` ID, not the
// `OrderRequest` ID. Once you submit a cancellation request, it cannot be undone.
// Be advised that orders with a status of PENDING_FILL, PENDING_ESCROW, FILLED,
// REJECTED, or CANCELLED cannot be cancelled.
//
// `Order` cancellation is not guaranteed nor is it immediate. The `Order` may
// still be executed if the cancellation request is not received in time.
//
// Check the status using the "Get Order by ID" endpoint to confirm whether the
// `Order` has been cancelled.
func (r *APIV2AccountOrderService) Cancel(ctx context.Context, orderID string, body APIV2AccountOrderCancelParams, opts ...option.RequestOption) (res *Order, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/%s/cancel", body.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get `OrderFulfillments` for a specific `Order`.
func (r *APIV2AccountOrderService) GetFulfillments(ctx context.Context, orderID string, params APIV2AccountOrderGetFulfillmentsParams, opts ...option.RequestOption) (res *[]OrderFulfillment, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/%s/fulfillments", params.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type Order struct {
	// ID of the `Order`.
	ID string `json:"id,required" format:"uuid"`
	// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
	// on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866".
	ChainID OrderChainID `json:"chain_id,required"`
	// Datetime at which the `Order` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Smart contract address that `Order` was created from.
	OrderContractAddress string `json:"order_contract_address,required" format:"eth_address"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderOrderSide `json:"order_side,required"`
	// Time in force. Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderOrderTif `json:"order_tif,required"`
	// Transaction hash for the `Order` creation.
	OrderTransactionHash string `json:"order_transaction_hash,required" format:"hex_string"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderOrderType `json:"order_type,required"`
	// Status of the `Order`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "FILLED", "REJECTED", "REQUIRING_CONTACT",
	// "ERROR".
	Status OrderStatus `json:"status,required"`
	// The `Stock` ID associated with the `Order`
	StockID string `json:"stock_id,required" format:"uuid"`
	// The dShare asset token address.
	AssetToken string `json:"asset_token" format:"eth_address"`
	// Total amount of assets involved.
	AssetTokenQuantity float64 `json:"asset_token_quantity"`
	// Transaction hash for cancellation of `Order`, if the `Order` was cancelled.
	CancelTransactionHash string `json:"cancel_transaction_hash" format:"hex_string"`
	// Fee amount associated with `Order`.
	Fee float64 `json:"fee"`
	// For limit `Orders`, the price per asset, specified in the `Stock`'s native
	// currency (USD for US equities and ETFs).
	LimitPrice float64 `json:"limit_price"`
	// The payment token (stablecoin) address.
	PaymentToken string `json:"payment_token" format:"eth_address"`
	// Total amount of payment involved.
	PaymentTokenQuantity float64 `json:"payment_token_quantity"`
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
		Status                respjson.Field
		StockID               respjson.Field
		AssetToken            respjson.Field
		AssetTokenQuantity    respjson.Field
		CancelTransactionHash respjson.Field
		Fee                   respjson.Field
		LimitPrice            respjson.Field
		PaymentToken          respjson.Field
		PaymentTokenQuantity  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
// on.
type OrderChainID string

const (
	OrderChainIDEip155_1     OrderChainID = "eip155:1"
	OrderChainIDEip155_42161 OrderChainID = "eip155:42161"
	OrderChainIDEip155_8453  OrderChainID = "eip155:8453"
	OrderChainIDEip155_81457 OrderChainID = "eip155:81457"
	OrderChainIDEip155_7887  OrderChainID = "eip155:7887"
	OrderChainIDEip155_98866 OrderChainID = "eip155:98866"
)

// Indicates whether `Order` is a buy or sell.
type OrderOrderSide string

const (
	OrderOrderSideBuy  OrderOrderSide = "BUY"
	OrderOrderSideSell OrderOrderSide = "SELL"
)

// Time in force. Indicates how long `Order` is valid for.
type OrderOrderTif string

const (
	OrderOrderTifDay OrderOrderTif = "DAY"
	OrderOrderTifGtc OrderOrderTif = "GTC"
	OrderOrderTifIoc OrderOrderTif = "IOC"
	OrderOrderTifFok OrderOrderTif = "FOK"
)

// Type of `Order`.
type OrderOrderType string

const (
	OrderOrderTypeMarket OrderOrderType = "MARKET"
	OrderOrderTypeLimit  OrderOrderType = "LIMIT"
)

// Status of the `Order`.
type OrderStatus string

const (
	OrderStatusPendingSubmit    OrderStatus = "PENDING_SUBMIT"
	OrderStatusPendingCancel    OrderStatus = "PENDING_CANCEL"
	OrderStatusPendingEscrow    OrderStatus = "PENDING_ESCROW"
	OrderStatusPendingFill      OrderStatus = "PENDING_FILL"
	OrderStatusEscrowed         OrderStatus = "ESCROWED"
	OrderStatusSubmitted        OrderStatus = "SUBMITTED"
	OrderStatusCancelled        OrderStatus = "CANCELLED"
	OrderStatusFilled           OrderStatus = "FILLED"
	OrderStatusRejected         OrderStatus = "REJECTED"
	OrderStatusRequiringContact OrderStatus = "REQUIRING_CONTACT"
	OrderStatusError            OrderStatus = "ERROR"
)

type APIV2AccountOrderGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type APIV2AccountOrderListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2AccountOrderListParams]'s query parameters as
// `url.Values`.
func (r APIV2AccountOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV2AccountOrderCancelParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type APIV2AccountOrderGetFulfillmentsParams struct {
	AccountID string           `path:"account_id,required" format:"uuid" json:"-"`
	Page      param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2AccountOrderGetFulfillmentsParams]'s query parameters
// as `url.Values`.
func (r APIV2AccountOrderGetFulfillmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
