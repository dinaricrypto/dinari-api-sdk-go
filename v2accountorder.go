// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

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

// V2AccountOrderService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderService] method instead.
type V2AccountOrderService struct {
	Options []option.RequestOption
}

// NewV2AccountOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AccountOrderService(opts ...option.RequestOption) (r V2AccountOrderService) {
	r = V2AccountOrderService{}
	r.Options = opts
	return
}

// Get a specific `Order` by its ID.
func (r *V2AccountOrderService) Get(ctx context.Context, orderID string, query V2AccountOrderGetParams, opts ...option.RequestOption) (res *Order, err error) {
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
func (r *V2AccountOrderService) List(ctx context.Context, accountID string, query V2AccountOrderListParams, opts ...option.RequestOption) (res *[]Order, err error) {
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
func (r *V2AccountOrderService) Cancel(ctx context.Context, orderID string, body V2AccountOrderCancelParams, opts ...option.RequestOption) (res *Order, err error) {
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

// Get fee quote data for an `Order`.
//
// The `order_fee_contract_object` property contains the fee quote structure to be
// used verbatim when placing an `Order` directly through our Contracts.
func (r *V2AccountOrderService) GetFeeQuote(ctx context.Context, accountID string, body V2AccountOrderGetFeeQuoteParams, opts ...option.RequestOption) (res *V2AccountOrderGetFeeQuoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/fee_quote", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get `OrderFulfillments` for a specific `Order`.
func (r *V2AccountOrderService) GetFulfillments(ctx context.Context, orderID string, params V2AccountOrderGetFulfillmentsParams, opts ...option.RequestOption) (res *[]Fulfillment, err error) {
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

type BrokerageOrderStatus string

const (
	BrokerageOrderStatusPendingSubmit    BrokerageOrderStatus = "PENDING_SUBMIT"
	BrokerageOrderStatusPendingCancel    BrokerageOrderStatus = "PENDING_CANCEL"
	BrokerageOrderStatusPendingEscrow    BrokerageOrderStatus = "PENDING_ESCROW"
	BrokerageOrderStatusPendingFill      BrokerageOrderStatus = "PENDING_FILL"
	BrokerageOrderStatusEscrowed         BrokerageOrderStatus = "ESCROWED"
	BrokerageOrderStatusSubmitted        BrokerageOrderStatus = "SUBMITTED"
	BrokerageOrderStatusCancelled        BrokerageOrderStatus = "CANCELLED"
	BrokerageOrderStatusFilled           BrokerageOrderStatus = "FILLED"
	BrokerageOrderStatusRejected         BrokerageOrderStatus = "REJECTED"
	BrokerageOrderStatusRequiringContact BrokerageOrderStatus = "REQUIRING_CONTACT"
	BrokerageOrderStatusError            BrokerageOrderStatus = "ERROR"
)

type Order struct {
	// ID of the `Order`.
	ID string `json:"id,required" format:"uuid"`
	// CAIP-2 formatted chain ID of the blockchain that the `Order` transaction was run
	// on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
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
	// Status of the `Order`.
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "FILLED", "REJECTED", "REQUIRING_CONTACT",
	// "ERROR".
	Status BrokerageOrderStatus `json:"status,required"`
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

type OrderSide string

const (
	OrderSideBuy  OrderSide = "BUY"
	OrderSideSell OrderSide = "SELL"
)

type OrderTif string

const (
	OrderTifDay OrderTif = "DAY"
	OrderTifGtc OrderTif = "GTC"
	OrderTifIoc OrderTif = "IOC"
	OrderTifFok OrderTif = "FOK"
)

type OrderType string

const (
	OrderTypeMarket OrderType = "MARKET"
	OrderTypeLimit  OrderType = "LIMIT"
)

type V2AccountOrderGetFeeQuoteResponse struct {
	// CAIP-2 chain ID of the blockchain where the `Order` will be placed
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
	ChainID Chain `json:"chain_id,required"`
	// The total quantity of the fees paid via payment token.
	Fee float64 `json:"fee,required"`
	// Opaque fee quote object to pass into the contract when creating an `Order`
	// directly through Dinari's smart contracts.
	OrderFeeContractObject V2AccountOrderGetFeeQuoteResponseOrderFeeContractObject `json:"order_fee_contract_object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID                respjson.Field
		Fee                    respjson.Field
		OrderFeeContractObject respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderGetFeeQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderGetFeeQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opaque fee quote object to pass into the contract when creating an `Order`
// directly through Dinari's smart contracts.
type V2AccountOrderGetFeeQuoteResponseOrderFeeContractObject struct {
	// EVM chain ID where the order is placed
	ChainID int64 `json:"chain_id,required"`
	// `FeeQuote` structure to pass into contracts.
	FeeQuote V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFeeQuote `json:"fee_quote,required"`
	// Signed `FeeQuote` structure to pass into contracts.
	FeeQuoteSignature string `json:"fee_quote_signature,required" format:"hex_string"`
	// Breakdown of fees
	Fees []V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFee `json:"fees,required"`
	// Address of payment token used for fees
	PaymentToken string `json:"payment_token,required" format:"eth_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID           respjson.Field
		FeeQuote          respjson.Field
		FeeQuoteSignature respjson.Field
		Fees              respjson.Field
		PaymentToken      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderGetFeeQuoteResponseOrderFeeContractObject) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderGetFeeQuoteResponseOrderFeeContractObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `FeeQuote` structure to pass into contracts.
type V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFeeQuote struct {
	Deadline  int64  `json:"deadline,required"`
	Fee       string `json:"fee,required" format:"bigint"`
	OrderID   string `json:"orderId,required" format:"bigint"`
	Requester string `json:"requester,required" format:"eth_address"`
	Timestamp int64  `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deadline    respjson.Field
		Fee         respjson.Field
		OrderID     respjson.Field
		Requester   respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFeeQuote) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFeeQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFee struct {
	// The quantity of the fee paid via payment token in
	// [ETH](https://ethereum.org/en/developers/docs/intro-to-ether/#what-is-ether).
	FeeInEth float64 `json:"fee_in_eth,required"`
	// The quantity of the fee paid via payment token in
	// [wei](https://ethereum.org/en/developers/docs/intro-to-ether/#denominations).
	FeeInWei string `json:"fee_in_wei,required" format:"bigint"`
	// Type of fee.
	//
	// Any of "SPONSORED_NETWORK", "NETWORK", "TRADING", "ORDER", "PARTNER_ORDER",
	// "PARTNER_TRADING".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeeInEth    respjson.Field
		FeeInWei    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFee) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderGetFeeQuoteResponseOrderFeeContractObjectFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountOrderListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountOrderListParams]'s query parameters as
// `url.Values`.
func (r V2AccountOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2AccountOrderCancelParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2AccountOrderGetFeeQuoteParams struct {
	// CAIP-2 chain ID of the blockchain where the `Order` will be placed.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
	ChainID Chain `json:"chain_id,omitzero,required"`
	// Address of the smart contract that will create the `Order`.
	ContractAddress string `json:"contract_address,required" format:"eth_address"`
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
	// The Stock ID associated with the Order
	StockID string `json:"stock_id,required" format:"uuid"`
	// Amount of dShare asset tokens involved. Required for limit `Orders` and market
	// sell `Orders`.
	AssetTokenQuantity param.Opt[float64] `json:"asset_token_quantity,omitzero"`
	// Price per asset in the asset's native currency. USD for US equities and ETFs.
	// Required for limit `Orders`.
	LimitPrice param.Opt[float64] `json:"limit_price,omitzero"`
	// Address of payment token.
	PaymentToken param.Opt[string] `json:"payment_token,omitzero" format:"eth_address"`
	// Amount of payment tokens involved. Required for market buy `Orders`.
	PaymentTokenQuantity param.Opt[float64] `json:"payment_token_quantity,omitzero"`
	paramObj
}

func (r V2AccountOrderGetFeeQuoteParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderGetFeeQuoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderGetFeeQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderGetFulfillmentsParams struct {
	AccountID string           `path:"account_id,required" format:"uuid" json:"-"`
	Page      param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountOrderGetFulfillmentsParams]'s query parameters as
// `url.Values`.
func (r V2AccountOrderGetFulfillmentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
