// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"errors"
	"fmt"
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

// V2AccountService contains methods and other services that help with interacting
// with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountService] method instead.
type V2AccountService struct {
	Options            []option.RequestOption
	Wallet             V2AccountWalletService
	Orders             V2AccountOrderService
	OrderFulfillments  V2AccountOrderFulfillmentService
	OrderRequests      V2AccountOrderRequestService
	WithdrawalRequests V2AccountWithdrawalRequestService
	Withdrawals        V2AccountWithdrawalService
	TokenTransfers     V2AccountTokenTransferService
}

// NewV2AccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AccountService(opts ...option.RequestOption) (r V2AccountService) {
	r = V2AccountService{}
	r.Options = opts
	r.Wallet = NewV2AccountWalletService(opts...)
	r.Orders = NewV2AccountOrderService(opts...)
	r.OrderFulfillments = NewV2AccountOrderFulfillmentService(opts...)
	r.OrderRequests = NewV2AccountOrderRequestService(opts...)
	r.WithdrawalRequests = NewV2AccountWithdrawalRequestService(opts...)
	r.Withdrawals = NewV2AccountWithdrawalService(opts...)
	r.TokenTransfers = NewV2AccountTokenTransferService(opts...)
	return
}

// Get a specific `Account` by its ID.
func (r *V2AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the `Account` to be inactive. Inactive accounts cannot be used for trading.
func (r *V2AccountService) Deactivate(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/deactivate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get the cash balances of the `Account`, including stablecoins and other cash
// equivalents.
func (r *V2AccountService) GetCashBalances(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]V2AccountGetCashBalancesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/cash", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get dividend payments made to the `Account` from dividend-bearing stock
// holdings.
func (r *V2AccountService) GetDividendPayments(ctx context.Context, accountID string, query V2AccountGetDividendPaymentsParams, opts ...option.RequestOption) (res *[]V2AccountGetDividendPaymentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/dividend_payments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get interest payments made to the `Account` from yield-bearing cash holdings.
//
// Currently, the only yield-bearing stablecoin accepted by Dinari is
// [USD+](https://usd.dinari.com/).
func (r *V2AccountService) GetInterestPayments(ctx context.Context, accountID string, query V2AccountGetInterestPaymentsParams, opts ...option.RequestOption) (res *[]V2AccountGetInterestPaymentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/interest_payments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the portfolio of the `Account`, excluding cash equivalents such as
// stablecoins.
func (r *V2AccountService) GetPortfolio(ctx context.Context, accountID string, opts ...option.RequestOption) (res *V2AccountGetPortfolioResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/portfolio", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Mints 1,000 mockUSD sandbox payment tokens to the `Wallet` connected to the
// `Account`.
//
// This feature is only supported in sandbox mode.
func (r *V2AccountService) MintSandboxTokens(ctx context.Context, accountID string, body V2AccountMintSandboxTokensParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/faucet", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type Chain string

const (
	ChainEip155_1         Chain = "eip155:1"
	ChainEip155_42161     Chain = "eip155:42161"
	ChainEip155_8453      Chain = "eip155:8453"
	ChainEip155_81457     Chain = "eip155:81457"
	ChainEip155_98866     Chain = "eip155:98866"
	ChainEip155_11155111  Chain = "eip155:11155111"
	ChainEip155_421614    Chain = "eip155:421614"
	ChainEip155_84532     Chain = "eip155:84532"
	ChainEip155_168587773 Chain = "eip155:168587773"
	ChainEip155_98867     Chain = "eip155:98867"
	ChainEip155_202110    Chain = "eip155:202110"
	ChainEip155_179205    Chain = "eip155:179205"
	ChainEip155_179202    Chain = "eip155:179202"
	ChainEip155_98865     Chain = "eip155:98865"
	ChainEip155_7887      Chain = "eip155:7887"
)

// Balance of a payment token in an `Account`.
type V2AccountGetCashBalancesResponse struct {
	// Total amount of the payment token in the `Account`.
	Amount float64 `json:"amount,required"`
	// CAIP-2 chain ID of the payment token.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,required"`
	// Symbol of the payment token.
	Symbol string `json:"symbol,required"`
	// Address of the payment token.
	TokenAddress string `json:"token_address,required" format:"eth_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		ChainID      respjson.Field
		Symbol       respjson.Field
		TokenAddress respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountGetCashBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountGetCashBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a dividend payment event for an `Account`.
type V2AccountGetDividendPaymentsResponse struct {
	// Amount of the dividend paid.
	Amount float64 `json:"amount,required"`
	// Currency in which the dividend was paid. (e.g. USD)
	Currency string `json:"currency,required"`
	// Date the dividend was distributed to the account. ISO 8601 format, YYYY-MM-DD.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// ID of the `Stock` for which the dividend was paid.
	StockID string `json:"stock_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		PaymentDate respjson.Field
		StockID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountGetDividendPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountGetDividendPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object representing an interest payment from stablecoin holdings.
type V2AccountGetInterestPaymentsResponse struct {
	// Amount of interest paid.
	Amount float64 `json:"amount,required"`
	// Currency in which the interest was paid (e.g. USD).
	Currency string `json:"currency,required"`
	// Date of interest payment in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		PaymentDate respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountGetInterestPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountGetInterestPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance information of `Stock` assets in your `Account`.
type V2AccountGetPortfolioResponse struct {
	// Balance details for all owned `Stocks`.
	Assets []V2AccountGetPortfolioResponseAsset `json:"assets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountGetPortfolioResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountGetPortfolioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance of a dShare in an `Account`.
type V2AccountGetPortfolioResponseAsset struct {
	// Total amount of the dShare asset token in the `Account`.
	Amount float64 `json:"amount,required"`
	// CAIP-2 chain ID of the blockchain where the dShare asset token exists.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,required"`
	// ID of the underlying `Stock` represented by the dShare asset token.
	StockID string `json:"stock_id,required" format:"uuid"`
	// Token symbol of the dShare asset token.
	Symbol string `json:"symbol,required"`
	// Address of the dShare asset token.
	TokenAddress string `json:"token_address,required" format:"eth_address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount       respjson.Field
		ChainID      respjson.Field
		StockID      respjson.Field
		Symbol       respjson.Field
		TokenAddress respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2AccountGetPortfolioResponseAsset) RawJSON() string { return r.JSON.raw }
func (r *V2AccountGetPortfolioResponseAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountGetDividendPaymentsParams struct {
	// End date, exclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	EndDate time.Time `query:"end_date,required" format:"date" json:"-"`
	// Start date, inclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	StartDate time.Time        `query:"start_date,required" format:"date" json:"-"`
	Page      param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Optional ID of the `Stock` to filter by
	StockID param.Opt[string] `query:"stock_id,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountGetDividendPaymentsParams]'s query parameters as
// `url.Values`.
func (r V2AccountGetDividendPaymentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2AccountGetInterestPaymentsParams struct {
	// End date, exclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	EndDate time.Time `query:"end_date,required" format:"date" json:"-"`
	// Start date, inclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	StartDate time.Time        `query:"start_date,required" format:"date" json:"-"`
	Page      param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountGetInterestPaymentsParams]'s query parameters as
// `url.Values`.
func (r V2AccountGetInterestPaymentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2AccountMintSandboxTokensParams struct {
	// CAIP-2 chain ID of blockchain in which to mint the sandbox payment tokens. If
	// none specified, defaults to eip155:421614. If the `Account` is linked to a
	// Dinari-managed `Wallet`, only eip155:42161 is allowed.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:202110", "eip155:179205",
	// "eip155:179202", "eip155:98865", "eip155:7887".
	ChainID Chain `json:"chain_id,omitzero"`
	paramObj
}

func (r V2AccountMintSandboxTokensParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountMintSandboxTokensParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountMintSandboxTokensParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
