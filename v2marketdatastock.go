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

// V2MarketDataStockService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MarketDataStockService] method instead.
type V2MarketDataStockService struct {
	Options []option.RequestOption
	// **Corporate actions are events that affect the ownership of a `Stock`.**
	//
	// Corporate actions include dividends and stock splits.
	Splits V2MarketDataStockSplitService
}

// NewV2MarketDataStockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2MarketDataStockService(opts ...option.RequestOption) (r V2MarketDataStockService) {
	r = V2MarketDataStockService{}
	r.Options = opts
	r.Splits = NewV2MarketDataStockSplitService(opts...)
	return
}

// Get a list of `Stocks`.
func (r *V2MarketDataStockService) List(ctx context.Context, query V2MarketDataStockListParams, opts ...option.RequestOption) (res *V2MarketDataStockListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/market_data/stocks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get current price for a specified `Stock`.
func (r *V2MarketDataStockService) GetCurrentPrice(ctx context.Context, stockID string, opts ...option.RequestOption) (res *V2MarketDataStockGetCurrentPriceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/current_price", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get quote for a specified `Stock`.
func (r *V2MarketDataStockService) GetCurrentQuote(ctx context.Context, stockID string, opts ...option.RequestOption) (res *V2MarketDataStockGetCurrentQuoteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/current_quote", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a list of announced stock dividend details for a specified `Stock`.
//
// Note that this data applies only to actual stocks. Yield received for holding
// tokenized shares may differ from this.
func (r *V2MarketDataStockService) GetDividends(ctx context.Context, stockID string, opts ...option.RequestOption) (res *[]V2MarketDataStockGetDividendsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/dividends", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get historical price data for a specified `Stock`. Each index in the array
// represents a single tick in a price chart.
func (r *V2MarketDataStockService) GetHistoricalPrices(ctx context.Context, stockID string, query V2MarketDataStockGetHistoricalPricesParams, opts ...option.RequestOption) (res *[]V2MarketDataStockGetHistoricalPricesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/historical_prices/", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the most recent news articles relating to a `Stock`, including a summary of
// the article and a link to the original source.
func (r *V2MarketDataStockService) GetNews(ctx context.Context, stockID string, query V2MarketDataStockGetNewsParams, opts ...option.RequestOption) (res *[]V2MarketDataStockGetNewsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/news", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V2MarketDataStockListResponse struct {
	// List of Stock
	Data []V2MarketDataStockListResponseData `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata V2MarketDataStockListResponsePaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedStockResponse:v1".
	Sv V2MarketDataStockListResponse_Sv `json:"_sv"`
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
func (r V2MarketDataStockListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about stock available for trading.
type V2MarketDataStockListResponseData struct {
	// ID of the `Stock`
	ID string `json:"id" api:"required" format:"uuid"`
	// Whether the `Stock` allows for fractional trading. If it is not fractionable,
	// Dinari only supports limit orders for the `Stock`.
	IsFractionable bool `json:"is_fractionable" api:"required"`
	// Whether the `Stock` is available for trading.
	IsTradable bool `json:"is_tradable" api:"required"`
	// Company name
	Name string `json:"name" api:"required"`
	// Ticker symbol
	Symbol string `json:"symbol" api:"required"`
	// List of CAIP-10 formatted token addresses.
	Tokens []string `json:"tokens" api:"required"`
	// SEC Central Index Key. Refer to
	// [this link](https://www.sec.gov/submit-filings/filer-support-resources/how-do-i-guides/understand-utilize-edgar-ciks-passphrases-access-codes)
	// for more information.
	Cik string `json:"cik" api:"nullable"`
	// Composite FIGI ID. Refer to [this link](https://www.openfigi.com/about/figi) for
	// more information.
	CompositeFigi string `json:"composite_figi" api:"nullable"`
	// CUSIP ID. Refer to [this link](https://www.cusip.com/identifiers.html) for more
	// information. A license agreement with CUSIP Global Services is required to
	// receive this value.
	Cusip string `json:"cusip" api:"nullable"`
	// Description of the company and their services.
	Description string `json:"description" api:"nullable"`
	// Name of `Stock` for application display. If defined, this supercedes the `name`
	// field for displaying the name.
	DisplayName string `json:"display_name" api:"nullable"`
	// URL of the company's logo. Supported formats are SVG and PNG.
	LogoURL string `json:"logo_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		IsFractionable respjson.Field
		IsTradable     respjson.Field
		Name           respjson.Field
		Symbol         respjson.Field
		Tokens         respjson.Field
		Cik            respjson.Field
		CompositeFigi  respjson.Field
		Cusip          respjson.Field
		Description    respjson.Field
		DisplayName    respjson.Field
		LogoURL        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination metadata
type V2MarketDataStockListResponsePaginationMetadata struct {
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
func (r V2MarketDataStockListResponsePaginationMetadata) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockListResponsePaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version
type V2MarketDataStockListResponse_Sv string

const (
	V2MarketDataStockListResponse_SvPaginatedStockResponseV1 V2MarketDataStockListResponse_Sv = "PaginatedStockResponse:v1"
)

type V2MarketDataStockGetCurrentPriceResponse struct {
	// The price (fair market value) of the asset at the given time period.
	Price float64 `json:"price" api:"required"`
	// ID of the `Stock`
	StockID string `json:"stock_id" api:"required" format:"uuid"`
	// When the `StockPrice` was generated.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The change in price from the previous close.
	Change float64 `json:"change" api:"nullable"`
	// The percentage change in price from the previous close.
	ChangePercent float64 `json:"change_percent" api:"nullable"`
	// The close price from the given time period.
	Close float64 `json:"close" api:"nullable"`
	// The highest price from the given time period
	High float64 `json:"high" api:"nullable"`
	// The lowest price from the given time period.
	Low float64 `json:"low" api:"nullable"`
	// The market capitalization of the `Stock` calculated at the most recent close
	// price.
	MarketCap int64 `json:"market_cap" api:"nullable"`
	// The open price from the given time period.
	Open float64 `json:"open" api:"nullable"`
	// The close price for the `Stock` from the previous trading session.
	PreviousClose float64 `json:"previous_close" api:"nullable"`
	// The trading volume in shares from the given time period.
	Volume float64 `json:"volume" api:"nullable"`
	// The number of shares outstanding in the given time period.
	WeightedSharesOutstanding int64 `json:"weighted_shares_outstanding" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price                     respjson.Field
		StockID                   respjson.Field
		Timestamp                 respjson.Field
		Change                    respjson.Field
		ChangePercent             respjson.Field
		Close                     respjson.Field
		High                      respjson.Field
		Low                       respjson.Field
		MarketCap                 respjson.Field
		Open                      respjson.Field
		PreviousClose             respjson.Field
		Volume                    respjson.Field
		WeightedSharesOutstanding respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetCurrentPriceResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetCurrentPriceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stock Quote
type V2MarketDataStockGetCurrentQuoteResponse struct {
	// The ask price. 0 if there is no active ask.
	AskPrice float64 `json:"ask_price" api:"required"`
	// The ask size in shares.
	AskSize float64 `json:"ask_size" api:"required"`
	// The bid price. 0 if there is no active bid.
	BidPrice float64 `json:"bid_price" api:"required"`
	// The bid size in shares.
	BidSize float64 `json:"bid_size" api:"required"`
	// ID of the `Stock`
	StockID string `json:"stock_id" api:"required" format:"uuid"`
	// When the `StockQuote` was generated.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Schema version
	//
	// Any of "StockQuote:v1".
	Sv V2MarketDataStockGetCurrentQuoteResponse_Sv `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AskPrice    respjson.Field
		AskSize     respjson.Field
		BidPrice    respjson.Field
		BidSize     respjson.Field
		StockID     respjson.Field
		Timestamp   respjson.Field
		Sv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetCurrentQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetCurrentQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema version
type V2MarketDataStockGetCurrentQuoteResponse_Sv string

const (
	V2MarketDataStockGetCurrentQuoteResponse_SvStockQuoteV1 V2MarketDataStockGetCurrentQuoteResponse_Sv = "StockQuote:v1"
)

// Information about a dividend announcement for a `Stock`.
type V2MarketDataStockGetDividendsResponse struct {
	// Cash amount of the dividend per share owned.
	CashAmount float64 `json:"cash_amount" api:"nullable"`
	// Currency in which the dividend is paid.
	Currency string `json:"currency" api:"nullable"`
	// Type of dividend. Dividends that have been paid and/or are expected to be paid
	// on consistent schedules are denoted as `CD`. Special Cash dividends that have
	// been paid that are infrequent or unusual, and/or can not be expected to occur in
	// the future are denoted as `SC`.
	DividendType string `json:"dividend_type" api:"nullable"`
	// Date on or after which a `Stock` is traded without the right to receive the next
	// dividend payment. If you purchase a `Stock` on or after the ex-dividend date,
	// you will not receive the upcoming dividend. In ISO 8601 format, YYYY-MM-DD.
	ExDividendDate time.Time `json:"ex_dividend_date" api:"nullable" format:"date"`
	// Date on which the dividend is paid out. In ISO 8601 format, YYYY-MM-DD.
	PayDate time.Time `json:"pay_date" api:"nullable" format:"date"`
	// Date that the shares must be held to receive the dividend; set by the company.
	// In ISO 8601 format, YYYY-MM-DD.
	RecordDate time.Time `json:"record_date" api:"nullable" format:"date"`
	// Ticker symbol of the `Stock`.
	Ticker string `json:"ticker" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CashAmount     respjson.Field
		Currency       respjson.Field
		DividendType   respjson.Field
		ExDividendDate respjson.Field
		PayDate        respjson.Field
		RecordDate     respjson.Field
		Ticker         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetDividendsResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetDividendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Datapoint of historical price data for a `Stock`.
type V2MarketDataStockGetHistoricalPricesResponse struct {
	// Close price from the given time period.
	Close float64 `json:"close" api:"required"`
	// Highest price from the given time period.
	High float64 `json:"high" api:"required"`
	// Lowest price from the given time period.
	Low float64 `json:"low" api:"required"`
	// Open price from the given time period.
	Open float64 `json:"open" api:"required"`
	// The UNIX timestamp in seconds for the start of the aggregate window.
	Timestamp int64 `json:"timestamp" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		High        respjson.Field
		Low         respjson.Field
		Open        respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetHistoricalPricesResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetHistoricalPricesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A news article relating to a `Stock` which includes a summary of the article and
// a link to the original source.
type V2MarketDataStockGetNewsResponse struct {
	// URL of the news article
	ArticleURL string `json:"article_url" api:"required"`
	// Description of the news article
	Description string `json:"description" api:"required"`
	// URL of the image for the news article
	ImageURL string `json:"image_url" api:"required"`
	// Datetime when the article was published. ISO 8601 timestamp.
	PublishedDt time.Time `json:"published_dt" api:"required" format:"date-time"`
	// The publisher of the news article
	Publisher string `json:"publisher" api:"required"`
	// Mobile-friendly Accelerated Mobile Page (AMP) URL of the news article, if
	// available
	AmpURL string `json:"amp_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArticleURL  respjson.Field
		Description respjson.Field
		ImageURL    respjson.Field
		PublishedDt respjson.Field
		Publisher   respjson.Field
		AmpURL      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2MarketDataStockListParams struct {
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2MarketDataStockListParamsOrder `query:"order,omitzero" json:"-"`
	// List of `Stock` symbols to query. If not provided, all `Stocks` are returned.
	Symbols []string `query:"symbols,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockListParams]'s query parameters as
// `url.Values`.
func (r V2MarketDataStockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2MarketDataStockListParamsOrder string

const (
	V2MarketDataStockListParamsOrderAsc  V2MarketDataStockListParamsOrder = "asc"
	V2MarketDataStockListParamsOrderDesc V2MarketDataStockListParamsOrder = "desc"
)

type V2MarketDataStockGetHistoricalPricesParams struct {
	// The timespan of the historical prices to query.
	//
	// Any of "DAY", "WEEK", "MONTH", "YEAR".
	Timespan V2MarketDataStockGetHistoricalPricesParamsTimespan `query:"timespan,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockGetHistoricalPricesParams]'s query
// parameters as `url.Values`.
func (r V2MarketDataStockGetHistoricalPricesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The timespan of the historical prices to query.
type V2MarketDataStockGetHistoricalPricesParamsTimespan string

const (
	V2MarketDataStockGetHistoricalPricesParamsTimespanDay   V2MarketDataStockGetHistoricalPricesParamsTimespan = "DAY"
	V2MarketDataStockGetHistoricalPricesParamsTimespanWeek  V2MarketDataStockGetHistoricalPricesParamsTimespan = "WEEK"
	V2MarketDataStockGetHistoricalPricesParamsTimespanMonth V2MarketDataStockGetHistoricalPricesParamsTimespan = "MONTH"
	V2MarketDataStockGetHistoricalPricesParamsTimespanYear  V2MarketDataStockGetHistoricalPricesParamsTimespan = "YEAR"
)

type V2MarketDataStockGetNewsParams struct {
	// The number of articles to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockGetNewsParams]'s query parameters as
// `url.Values`.
func (r V2MarketDataStockGetNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
