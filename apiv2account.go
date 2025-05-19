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

// APIV2AccountService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountService] method instead.
type APIV2AccountService struct {
	Options           []option.RequestOption
	Wallet            APIV2AccountWalletService
	Orders            APIV2AccountOrderService
	OrderFulfillments APIV2AccountOrderFulfillmentService
	OrderRequests     APIV2AccountOrderRequestService
}

// NewAPIV2AccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV2AccountService(opts ...option.RequestOption) (r APIV2AccountService) {
	r = APIV2AccountService{}
	r.Options = opts
	r.Wallet = NewAPIV2AccountWalletService(opts...)
	r.Orders = NewAPIV2AccountOrderService(opts...)
	r.OrderFulfillments = NewAPIV2AccountOrderFulfillmentService(opts...)
	r.OrderRequests = NewAPIV2AccountOrderRequestService(opts...)
	return
}

// Get a specific `Account` by its ID.
func (r *APIV2AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the `Account` to be inactive. Inactive accounts cannot be used for trading.
func (r *APIV2AccountService) Deactivate(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *APIV2AccountService) GetCash(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]Apiv2AccountGetCashResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *APIV2AccountService) GetDividendPayments(ctx context.Context, accountID string, query APIV2AccountGetDividendPaymentsParams, opts ...option.RequestOption) (res *[]Apiv2AccountGetDividendPaymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *APIV2AccountService) GetInterestPayments(ctx context.Context, accountID string, query APIV2AccountGetInterestPaymentsParams, opts ...option.RequestOption) (res *[]Apiv2AccountGetInterestPaymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *APIV2AccountService) GetPortfolio(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Apiv2AccountGetPortfolioResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/portfolio", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Balance of a payment token in an `Account`.
type Apiv2AccountGetCashResponse struct {
	// Total amount of the payment token in the `Account`.
	Amount float64 `json:"amount,required"`
	// CAIP-2 chain ID of the payment token.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866".
	ChainID Apiv2AccountGetCashResponseChainID `json:"chain_id,required"`
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
func (r Apiv2AccountGetCashResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetCashResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CAIP-2 chain ID of the payment token.
type Apiv2AccountGetCashResponseChainID string

const (
	Apiv2AccountGetCashResponseChainIDEip155_1     Apiv2AccountGetCashResponseChainID = "eip155:1"
	Apiv2AccountGetCashResponseChainIDEip155_42161 Apiv2AccountGetCashResponseChainID = "eip155:42161"
	Apiv2AccountGetCashResponseChainIDEip155_8453  Apiv2AccountGetCashResponseChainID = "eip155:8453"
	Apiv2AccountGetCashResponseChainIDEip155_81457 Apiv2AccountGetCashResponseChainID = "eip155:81457"
	Apiv2AccountGetCashResponseChainIDEip155_7887  Apiv2AccountGetCashResponseChainID = "eip155:7887"
	Apiv2AccountGetCashResponseChainIDEip155_98866 Apiv2AccountGetCashResponseChainID = "eip155:98866"
)

// Represents a dividend payment event for an `Account`.
type Apiv2AccountGetDividendPaymentsResponse struct {
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
func (r Apiv2AccountGetDividendPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetDividendPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object representing an interest payment from stablecoin holdings.
type Apiv2AccountGetInterestPaymentsResponse struct {
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
func (r Apiv2AccountGetInterestPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetInterestPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance information of `Stock` assets in your `Account`.
type Apiv2AccountGetPortfolioResponse struct {
	// Balance details for all owned `Stocks`.
	Assets []Apiv2AccountGetPortfolioResponseAsset `json:"assets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assets      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetPortfolioResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetPortfolioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance of a dShare in an `Account`.
type Apiv2AccountGetPortfolioResponseAsset struct {
	// Total amount of the dShare asset token in the `Account`.
	Amount float64 `json:"amount,required"`
	// CAIP-2 chain ID of the blockchain where the dShare asset token exists.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866".
	ChainID string `json:"chain_id,required"`
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
func (r Apiv2AccountGetPortfolioResponseAsset) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetPortfolioResponseAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2AccountGetDividendPaymentsParams struct {
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

// URLQuery serializes [APIV2AccountGetDividendPaymentsParams]'s query parameters
// as `url.Values`.
func (r APIV2AccountGetDividendPaymentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV2AccountGetInterestPaymentsParams struct {
	// End date, exclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	EndDate time.Time `query:"end_date,required" format:"date" json:"-"`
	// Start date, inclusive, in US Eastern time zone. ISO 8601 format, YYYY-MM-DD.
	StartDate time.Time        `query:"start_date,required" format:"date" json:"-"`
	Page      param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize  param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2AccountGetInterestPaymentsParams]'s query parameters
// as `url.Values`.
func (r APIV2AccountGetInterestPaymentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
