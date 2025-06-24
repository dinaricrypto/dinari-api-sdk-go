// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

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

// V2AccountOrderStockEip155Service contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderStockEip155Service] method instead.
type V2AccountOrderStockEip155Service struct {
	Options []option.RequestOption
}

// NewV2AccountOrderStockEip155Service generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV2AccountOrderStockEip155Service(opts ...option.RequestOption) (r V2AccountOrderStockEip155Service) {
	r = V2AccountOrderStockEip155Service{}
	r.Options = opts
	return
}

// Get fee quote data for an `Order` to be placed on Dinari's EVM smart contracts.
//
// Dinari's EVM smart contracts require a fee quote to be provided when placing an
// `Order`. Use this method to retrieve the quote.
//
// The `order_fee_contract_object` property contains the opaque fee quote structure
// to be used.
func (r *V2AccountOrderStockEip155Service) GetFeeQuote(ctx context.Context, accountID string, body V2AccountOrderStockEip155GetFeeQuoteParams, opts ...option.RequestOption) (res *V2AccountOrderStockEip155GetFeeQuoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/stocks/eip155/fee_quote", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create a set of transactions to create an `Order` using Dinari's EVM smart
// contracts.
//
// This is a convenience method to prepare the transactions needed to create an
// `Order` using Dinari's EVM smart contracts. Once signed, the transactions can be
// sent to the EVM network to create the order. Note that the fee quote is already
// included in the transactions, so no additional fee quote lookup is needed.
func (r *V2AccountOrderStockEip155Service) PrepareOrder(ctx context.Context, accountID string, body V2AccountOrderStockEip155PrepareOrderParams, opts ...option.RequestOption) (res *V2AccountOrderStockEip155PrepareOrderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/orders/stocks/eip155/prepare", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrderFeeAmount struct {
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
	Type OrderFeeAmountType `json:"type,required"`
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
func (r OrderFeeAmount) RawJSON() string { return r.JSON.raw }
func (r *OrderFeeAmount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of fee.
type OrderFeeAmountType string

const (
	OrderFeeAmountTypeSponsoredNetwork OrderFeeAmountType = "SPONSORED_NETWORK"
	OrderFeeAmountTypeNetwork          OrderFeeAmountType = "NETWORK"
	OrderFeeAmountTypeTrading          OrderFeeAmountType = "TRADING"
	OrderFeeAmountTypeOrder            OrderFeeAmountType = "ORDER"
	OrderFeeAmountTypePartnerOrder     OrderFeeAmountType = "PARTNER_ORDER"
	OrderFeeAmountTypePartnerTrading   OrderFeeAmountType = "PARTNER_TRADING"
)

type V2AccountOrderStockEip155GetFeeQuoteResponse struct {
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
	OrderFeeContractObject V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObject `json:"order_fee_contract_object,required"`
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
func (r V2AccountOrderStockEip155GetFeeQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderStockEip155GetFeeQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Opaque fee quote object to pass into the contract when creating an `Order`
// directly through Dinari's smart contracts.
type V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObject struct {
	// EVM chain ID of the blockchain where the `Order` will be placed.
	//
	// Any of 42161, 1, 8453, 81457, 7887, 98866.
	ChainID int64 `json:"chain_id,required"`
	// `FeeQuote` structure to pass into contracts.
	FeeQuote V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObjectFeeQuote `json:"fee_quote,required"`
	// Signed `FeeQuote` structure to pass into contracts.
	FeeQuoteSignature string `json:"fee_quote_signature,required" format:"hex_string"`
	// Breakdown of fees
	Fees []OrderFeeAmount `json:"fees,required"`
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
func (r V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObject) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// `FeeQuote` structure to pass into contracts.
type V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObjectFeeQuote struct {
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
func (r V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObjectFeeQuote) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderStockEip155GetFeeQuoteResponseOrderFeeContractObjectFeeQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prepared transactions to place an order on an EIP-155 (EVM) chain.
type V2AccountOrderStockEip155PrepareOrderResponse struct {
	// Fees included in the order transaction. Provided here as a reference.
	Fees []OrderFeeAmount `json:"fees,required"`
	// List of contract addresses and call data for building transactions to be signed
	// and submitted on chain. Transactions should be submitted on chain in the order
	// provided in this property.
	TransactionData []V2AccountOrderStockEip155PrepareOrderResponseTransactionData `json:"transaction_data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fees            respjson.Field
		TransactionData respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderStockEip155PrepareOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountOrderStockEip155PrepareOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about a transaction to be signed with a wallet and submitted on
// chain.
//
// Typically the transactions will include an ERC20 approval transaction to allow
// the Dinari smart contract to spend the payment token or Dshare asset tokens on
// behalf of the user, followed by a transaction to call the Dinari smart contract
// to create an `Order`.
type V2AccountOrderStockEip155PrepareOrderResponseTransactionData struct {
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Abi             respjson.Field
		Args            respjson.Field
		ContractAddress respjson.Field
		Data            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountOrderStockEip155PrepareOrderResponseTransactionData) RawJSON() string {
	return r.JSON.raw
}
func (r *V2AccountOrderStockEip155PrepareOrderResponseTransactionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderStockEip155GetFeeQuoteParams struct {
	// CAIP-2 chain ID of the blockchain where the `Order` will be placed.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
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
	// The ID of the `Stock` for which the `Order` is being placed.
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

func (r V2AccountOrderStockEip155GetFeeQuoteParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderStockEip155GetFeeQuoteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderStockEip155GetFeeQuoteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountOrderStockEip155PrepareOrderParams struct {
	// CAIP-2 chain ID of the blockchain where the `Order` will be placed.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
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
	// The ID of the `Stock` for which the `Order` is being placed.
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

func (r V2AccountOrderStockEip155PrepareOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountOrderStockEip155PrepareOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountOrderStockEip155PrepareOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
