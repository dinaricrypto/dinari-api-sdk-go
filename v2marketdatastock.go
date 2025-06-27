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

// V2MarketDataStockService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MarketDataStockService] method instead.
type V2MarketDataStockService struct {
	Options []option.RequestOption
	Splits  V2MarketDataStockSplitService
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
func (r *V2MarketDataStockService) List(ctx context.Context, query V2MarketDataStockListParams, opts ...option.RequestOption) (res *[]V2MarketDataStockListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/stocks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get current price for a specified `Stock`.
func (r *V2MarketDataStockService) GetCurrentPrice(ctx context.Context, stockID string, opts ...option.RequestOption) (res *V2MarketDataStockGetCurrentPriceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/current_price", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get quote for a specified `Stock`.
func (r *V2MarketDataStockService) GetCurrentQuote(ctx context.Context, stockID string, opts ...option.RequestOption) (res *V2MarketDataStockGetCurrentQuoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/current_quote", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of announced stock dividend details for a specified `Stock`.
//
// Note that this data applies only to actual stocks. Yield received for holding
// tokenized shares may differ from this.
func (r *V2MarketDataStockService) GetDividends(ctx context.Context, stockID string, opts ...option.RequestOption) (res *[]V2MarketDataStockGetDividendsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/dividends", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get historical price data for a specified `Stock`. Each index in the array
// represents a single tick in a price chart.
func (r *V2MarketDataStockService) GetHistoricalPrices(ctx context.Context, stockID string, query V2MarketDataStockGetHistoricalPricesParams, opts ...option.RequestOption) (res *[]V2MarketDataStockGetHistoricalPricesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/historical_prices/", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the most recent news articles relating to a `Stock`, including a summary of
// the article and a link to the original source.
func (r *V2MarketDataStockService) GetNews(ctx context.Context, stockID string, query V2MarketDataStockGetNewsParams, opts ...option.RequestOption) (res *[]V2MarketDataStockGetNewsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/news", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about stock available for trading.
type V2MarketDataStockListResponse struct {
	// ID of the `Stock`
	ID string `json:"id,required" format:"uuid"`
	// Whether the `Stock` allows for fractional trading. If it is not fractionable,
	// Dinari only supports limit orders for the `Stock`.
	IsFractionable bool `json:"is_fractionable,required"`
	// Whether the `Stock` is available for trading.
	IsTradable bool `json:"is_tradable,required"`
	// Company name
	Name string `json:"name,required"`
	// Ticker symbol
	Symbol string `json:"symbol,required"`
	// List of CAIP-10 formatted token addresses.
	Tokens []string `json:"tokens,required"`
	// SEC Central Index Key. Refer to
	// [this link](https://www.sec.gov/submit-filings/filer-support-resources/how-do-i-guides/understand-utilize-edgar-ciks-passphrases-access-codes)
	// for more information.
	Cik string `json:"cik,nullable"`
	// Composite FIGI ID. Refer to [this link](https://www.openfigi.com/about/figi) for
	// more information.
	CompositeFigi string `json:"composite_figi,nullable"`
	// CUSIP ID. Refer to [this link](https://www.cusip.com/identifiers.html) for more
	// information.
	Cusip string `json:"cusip,nullable"`
	// Description of the company and their services.
	Description string `json:"description,nullable"`
	// Name of `Stock` for application display. If defined, this supercedes the `name`
	// field for displaying the name.
	DisplayName string `json:"display_name,nullable"`
	// URL of the company's logo. Supported formats are SVG and PNG.
	LogoURL string `json:"logo_url,nullable"`
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
func (r V2MarketDataStockListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2MarketDataStockGetCurrentPriceResponse struct {
	// The ask price.
	Price float64 `json:"price,required"`
	// ID of the `Stock`
	StockID string `json:"stock_id,required" format:"uuid"`
	// When the Stock Quote was generated.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The change in price from the previous close.
	Change float64 `json:"change"`
	// The percentage change in price from the previous close.
	ChangePercent float64 `json:"change_percent"`
	// The close price from the given time period.
	Close float64 `json:"close"`
	// The highest price from the given time period
	High float64 `json:"high"`
	// The lowest price from the given time period.
	Low float64 `json:"low"`
	// The most recent close price of the ticker multiplied by weighted outstanding
	// shares.
	MarketCap int64 `json:"market_cap"`
	// The open price from the given time period.
	Open float64 `json:"open"`
	// The close price for the `Stock` from the previous trading session.
	PreviousClose float64 `json:"previous_close"`
	// The trading volume from the given time period.
	Volume float64 `json:"volume"`
	// The number of shares outstanding in the given time period.
	WeightedSharesOutstanding int64 `json:"weighted_shares_outstanding"`
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

type V2MarketDataStockGetCurrentQuoteResponse struct {
	// The ask price.
	AskPrice float64 `json:"ask_price,required"`
	// The ask size.
	AskSize float64 `json:"ask_size,required"`
	// The bid price.
	BidPrice float64 `json:"bid_price,required"`
	// The bid size.
	BidSize float64 `json:"bid_size,required"`
	// ID of the `Stock`
	StockID string `json:"stock_id,required" format:"uuid"`
	// When the Stock Quote was generated.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AskPrice    respjson.Field
		AskSize     respjson.Field
		BidPrice    respjson.Field
		BidSize     respjson.Field
		StockID     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataStockGetCurrentQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataStockGetCurrentQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about a dividend announcement for a `Stock`.
type V2MarketDataStockGetDividendsResponse struct {
	// Cash amount of the dividend per share owned.
	CashAmount float64 `json:"cash_amount"`
	// Currency in which the dividend is paid.
	Currency string `json:"currency"`
	// Date on which the dividend was announced. In ISO 8601 format, YYYY-MM-DD.
	DeclarationDate time.Time `json:"declaration_date" format:"date"`
	// Type of dividend. Dividends that have been paid and/or are expected to be paid
	// on consistent schedules are denoted as `CD`. Special Cash dividends that have
	// been paid that are infrequent or unusual, and/or can not be expected to occur in
	// the future are denoted as `SC`. Long-term and short-term capital gain
	// distributions are denoted as `LT` and `ST`, respectively.
	DividendType string `json:"dividend_type"`
	// Date on or after which a `Stock` is traded without the right to receive the next
	// dividend payment. If you purchase a `Stock` on or after the ex-dividend date,
	// you will not receive the upcoming dividend. In ISO 8601 format, YYYY-MM-DD.
	ExDividendDate time.Time `json:"ex_dividend_date" format:"date"`
	// Frequency of the dividend. The following values are possible:
	//
	// - `1` - Annual
	// - `2` - Semi-Annual
	// - `4` - Quarterly
	// - `12` - Monthly
	// - `52` - Weekly
	// - `365` - Daily
	Frequency int64 `json:"frequency"`
	// Date on which the dividend is paid out. In ISO 8601 format, YYYY-MM-DD.
	PayDate time.Time `json:"pay_date" format:"date"`
	// Date that the shares must be held to receive the dividend; set by the company.
	// In ISO 8601 format, YYYY-MM-DD.
	RecordDate time.Time `json:"record_date" format:"date"`
	// Ticker symbol of the `Stock`.
	Ticker string `json:"ticker"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CashAmount      respjson.Field
		Currency        respjson.Field
		DeclarationDate respjson.Field
		DividendType    respjson.Field
		ExDividendDate  respjson.Field
		Frequency       respjson.Field
		PayDate         respjson.Field
		RecordDate      respjson.Field
		Ticker          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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
	Close float64 `json:"close,required"`
	// Highest price from the given time period.
	High float64 `json:"high,required"`
	// Lowest price from the given time period.
	Low float64 `json:"low,required"`
	// Open price from the given time period.
	Open float64 `json:"open,required"`
	// The UNIX timestamp in seconds for the start of the aggregate window.
	Timestamp int64 `json:"timestamp,required"`
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
	ArticleURL string `json:"article_url,required"`
	// Description of the news article
	Description string `json:"description,required"`
	// URL of the image for the news article
	ImageURL string `json:"image_url,required"`
	// Datetime when the article was published. ISO 8601 timestamp.
	PublishedDt time.Time `json:"published_dt,required" format:"date-time"`
	// The publisher of the news article
	Publisher string `json:"publisher,required"`
	// Mobile-friendly Accelerated Mobile Page (AMP) URL of the news article, if
	// available
	AmpURL string `json:"amp_url"`
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
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// List of `Stock` symbols to query. If not provided, all `Stocks` are returned.
	Symbols []string `query:"symbols,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockListParams]'s query parameters as
// `url.Values`.
func (r V2MarketDataStockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2MarketDataStockGetHistoricalPricesParams struct {
	// The timespan of the historical prices to query.
	//
	// Any of "DAY", "WEEK", "MONTH", "YEAR".
	Timespan V2MarketDataStockGetHistoricalPricesParamsTimespan `query:"timespan,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockGetHistoricalPricesParams]'s query
// parameters as `url.Values`.
func (r V2MarketDataStockGetHistoricalPricesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
