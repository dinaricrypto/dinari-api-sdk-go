// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
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

// Retrieves details of a specific order by its ID.
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

// Lists all orders under the account.
func (r *APIV2AccountOrderService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]Order, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancels an order by its ID. Note that this requires the order ID, not the order
// request ID.
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

// Gets estimated fee data for an order to be placed directly through our
// contracts.
func (r *APIV2AccountOrderService) GetEstimatedFee(ctx context.Context, accountID string, body APIV2AccountOrderGetEstimatedFeeParams, opts ...option.RequestOption) (res *Apiv2AccountOrderGetEstimatedFeeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/estimated_fee", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves order fulfillments for a specific order.
func (r *APIV2AccountOrderService) GetFulfillments(ctx context.Context, orderID string, query APIV2AccountOrderGetFulfillmentsParams, opts ...option.RequestOption) (res *[]OrderFulfillment, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/%s/fulfillments", query.AccountID, orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Order struct {
	// Identifier of the order
	ID string `json:"id,required" format:"uuid"`
	// Total amount of assets involved
	AssetTokenQuantity float64 `json:"asset_token_quantity,required"`
	// Status of the order
	//
	// Any of "PENDING_SUBMIT", "PENDING_CANCEL", "PENDING_ESCROW", "PENDING_FILL",
	// "ESCROWED", "SUBMITTED", "CANCELLED", "FILLED", "REJECTED", "REQUIRING_CONTACT",
	// "ERROR".
	BrokerageOrderStatus OrderBrokerageOrderStatus `json:"brokerage_order_status,required"`
	// Blockchain that transaction was run on
	ChainID int64 `json:"chain_id,required"`
	// Smart Contract address that order came from
	OrderContractAddress string `json:"order_contract_address,required" format:"eth_address"`
	// Indicates if order is a buy or sell
	//
	// Any of "BUY", "SELL".
	OrderSide OrderOrderSide `json:"order_side,required"`
	// Indicates how long order is valid
	//
	// Any of "DAY", "GTC", "IOC", "FOK".
	OrderTif OrderOrderTif `json:"order_tif,required"`
	// Transaction hash for the order
	OrderTransactionHash string `json:"order_transaction_hash,required" format:"hex_string"`
	// Indicates what type of order
	//
	// Any of "MARKET", "LIMIT".
	OrderType OrderOrderType `json:"order_type,required"`
	// Total amount of payment involved
	PaymentTokenQuantity float64 `json:"payment_token_quantity,required"`
	// Unique identifier of this Order generated by the order contract.
	SmartContractOrderID string `json:"smart_contract_order_id,required"`
	// Transaction hash for cancellation of order
	CancelTransactionHash string `json:"cancel_transaction_hash" format:"hex_string"`
	// List of fees associated with order
	Fees []map[string]any `json:"fees"`
	// Total amount of network fee taken in USD
	NetworkFeeInUsd float64 `json:"network_fee_in_usd"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AssetTokenQuantity    respjson.Field
		BrokerageOrderStatus  respjson.Field
		ChainID               respjson.Field
		OrderContractAddress  respjson.Field
		OrderSide             respjson.Field
		OrderTif              respjson.Field
		OrderTransactionHash  respjson.Field
		OrderType             respjson.Field
		PaymentTokenQuantity  respjson.Field
		SmartContractOrderID  respjson.Field
		CancelTransactionHash respjson.Field
		Fees                  respjson.Field
		NetworkFeeInUsd       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the order
type OrderBrokerageOrderStatus string

const (
	OrderBrokerageOrderStatusPendingSubmit    OrderBrokerageOrderStatus = "PENDING_SUBMIT"
	OrderBrokerageOrderStatusPendingCancel    OrderBrokerageOrderStatus = "PENDING_CANCEL"
	OrderBrokerageOrderStatusPendingEscrow    OrderBrokerageOrderStatus = "PENDING_ESCROW"
	OrderBrokerageOrderStatusPendingFill      OrderBrokerageOrderStatus = "PENDING_FILL"
	OrderBrokerageOrderStatusEscrowed         OrderBrokerageOrderStatus = "ESCROWED"
	OrderBrokerageOrderStatusSubmitted        OrderBrokerageOrderStatus = "SUBMITTED"
	OrderBrokerageOrderStatusCancelled        OrderBrokerageOrderStatus = "CANCELLED"
	OrderBrokerageOrderStatusFilled           OrderBrokerageOrderStatus = "FILLED"
	OrderBrokerageOrderStatusRejected         OrderBrokerageOrderStatus = "REJECTED"
	OrderBrokerageOrderStatusRequiringContact OrderBrokerageOrderStatus = "REQUIRING_CONTACT"
	OrderBrokerageOrderStatusError            OrderBrokerageOrderStatus = "ERROR"
)

// Indicates if order is a buy or sell
type OrderOrderSide string

const (
	OrderOrderSideBuy  OrderOrderSide = "BUY"
	OrderOrderSideSell OrderOrderSide = "SELL"
)

// Indicates how long order is valid
type OrderOrderTif string

const (
	OrderOrderTifDay OrderOrderTif = "DAY"
	OrderOrderTifGtc OrderOrderTif = "GTC"
	OrderOrderTifIoc OrderOrderTif = "IOC"
	OrderOrderTifFok OrderOrderTif = "FOK"
)

// Indicates what type of order
type OrderOrderType string

const (
	OrderOrderTypeMarket OrderOrderType = "MARKET"
	OrderOrderTypeLimit  OrderOrderType = "LIMIT"
)

type Apiv2AccountOrderGetEstimatedFeeResponse struct {
	// Chain where the order is placed
	ChainID int64 `json:"chain_id,required"`
	// FeeQuote structure to pass into contracts
	FeeQuote Apiv2AccountOrderGetEstimatedFeeResponseFeeQuote `json:"fee_quote,required"`
	// Signed FeeQuote structure to pass into contracts
	FeeQuoteSignature string `json:"fee_quote_signature,required" format:"hex_string"`
	// Breakdown of fees
	Fees []Apiv2AccountOrderGetEstimatedFeeResponseFee `json:"fees,required"`
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
func (r Apiv2AccountOrderGetEstimatedFeeResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountOrderGetEstimatedFeeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FeeQuote structure to pass into contracts
type Apiv2AccountOrderGetEstimatedFeeResponseFeeQuote struct {
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
func (r Apiv2AccountOrderGetEstimatedFeeResponseFeeQuote) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountOrderGetEstimatedFeeResponseFeeQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv2AccountOrderGetEstimatedFeeResponseFee struct {
	// The quantity of the fee paid via payment token in ETH
	// <a href='https://ethereum.org/en/developers/docs/intro-to-ether/#what-is-ether' target='_blank'>(what
	// is ETH?)</a>
	FeeInEth float64 `json:"fee_in_eth,required"`
	// The quantity of the fee paid via payment token in wei
	// <a href='https://ethereum.org/en/developers/docs/intro-to-ether/#denominations' target='_blank'>(what
	// is wei?)</a>
	FeeInWei string `json:"fee_in_wei,required" format:"bigint"`
	// Type of fee
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
func (r Apiv2AccountOrderGetEstimatedFeeResponseFee) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountOrderGetEstimatedFeeResponseFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2AccountOrderGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type APIV2AccountOrderCancelParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}

type APIV2AccountOrderGetEstimatedFeeParams struct {
	// Chain where the order is placed
	ChainID int64 `json:"chain_id,required"`
	// Order contract address
	ContractAddress string `json:"contract_address,required" format:"eth_address"`
	// Order data from which to calculate the fees. To be specified in the future
	OrderData map[string]string `json:"order_data,omitzero,required"`
	paramObj
}

func (r APIV2AccountOrderGetEstimatedFeeParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2AccountOrderGetEstimatedFeeParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type APIV2AccountOrderGetFulfillmentsParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}
