// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
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

// V2Service contains methods and other services that help with interacting with
// the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2Service] method instead.
type V2Service struct {
	Options    []option.RequestOption
	MarketData V2MarketDataService
	Entities   V2EntityService
	Accounts   V2AccountService
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
func (r *V2Service) ListOrders(ctx context.Context, query V2ListOrdersParams, opts ...option.RequestOption) (res *[]V2ListOrdersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/orders/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V2ListOrdersResponse struct {
	// ID of the `Order`.
	ID string `json:"id,required" format:"uuid"`
	// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
	// on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,required"`
	// Datetime at which the `Order` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Smart contract address that `Order` was created from.
	OrderContractAddress string `json:"order_contract_address,required" format:"eth_address"`
	// Indicates whether `Order` is a buy or sell.
	//
	// Any of "BUY", "SELL".
	OrderSide OrderSide `json:"order_side,required"`
	// Time in force. Indicates how long `Order` is valid for.
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderTif `json:"order_tif,required"`
	// Transaction hash for the `Order` creation.
	OrderTransactionHash string `json:"order_transaction_hash,required" format:"hex_string"`
	// Type of `Order`.
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderType `json:"order_type,required"`
	// The payment token (stablecoin) address.
	PaymentToken string `json:"payment_token,required" format:"eth_address"`
	// Status of the `Order`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "FILLED", "REJECTED", "REQUIRING_CONTACT",
	// "ERROR".
	Status BrokerageOrderStatus `json:"status,required"`
	// The `Stock` ID associated with the `Order`
	StockID string `json:"stock_id,required" format:"uuid"`
	// Account ID the order was made for.
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The dShare asset token address.
	AssetToken string `json:"asset_token,nullable" format:"eth_address"`
	// Total amount of assets involved.
	AssetTokenQuantity float64 `json:"asset_token_quantity,nullable"`
	// Transaction hash for cancellation of `Order`, if the `Order` was cancelled.
	CancelTransactionHash string `json:"cancel_transaction_hash,nullable" format:"hex_string"`
	// Entity ID of the Order
	EntityID string `json:"entity_id,nullable" format:"uuid"`
	// Fee amount associated with `Order`.
	Fee float64 `json:"fee,nullable"`
	// For limit `Orders`, the price per asset, specified in the `Stock`'s native
	// currency (USD for US equities and ETFs).
	LimitPrice float64 `json:"limit_price,nullable"`
	// Order Request ID for the `Order`
	OrderRequestID string `json:"order_request_id,nullable" format:"uuid"`
	// Total amount of payment involved.
	PaymentTokenQuantity float64 `json:"payment_token_quantity,nullable"`
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
func (r V2ListOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *V2ListOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2ListOrdersParams struct {
	// Fulfillment transaction hash of the `Order`.
	OrderFulfillmentTransactionHash param.Opt[string] `query:"order_fulfillment_transaction_hash,omitzero" json:"-"`
	// Order Request ID for the `Order`
	OrderRequestID param.Opt[string] `query:"order_request_id,omitzero" format:"uuid" json:"-"`
	// Transaction hash of the `Order`.
	OrderTransactionHash param.Opt[string] `query:"order_transaction_hash,omitzero" json:"-"`
	Page                 param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PageSize             param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	// CAIP-2 formatted chain ID of the blockchain the `Order` was made on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `query:"chain_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2ListOrdersParams]'s query parameters as `url.Values`.
func (r V2ListOrdersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
