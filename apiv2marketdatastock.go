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

// APIV2MarketDataStockService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2MarketDataStockService] method instead.
type APIV2MarketDataStockService struct {
	Options []option.RequestOption
	Splits  APIV2MarketDataStockSplitService
}

// NewAPIV2MarketDataStockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV2MarketDataStockService(opts ...option.RequestOption) (r APIV2MarketDataStockService) {
	r = APIV2MarketDataStockService{}
	r.Options = opts
	r.Splits = NewAPIV2MarketDataStockSplitService(opts...)
	return
}

// Returns a list of stocks available for trading.
func (r *APIV2MarketDataStockService) List(ctx context.Context, query APIV2MarketDataStockListParams, opts ...option.RequestOption) (res *[]Apiv2MarketDataStockListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/stocks/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a list of announced stock dividend details for a specified stock. Note
// that this data applies only to actual stocks. Yield received for holding dShares
// may differ from this.
func (r *APIV2MarketDataStockService) GetDividends(ctx context.Context, stockID string, opts ...option.RequestOption) (res *[]Apiv2MarketDataStockGetDividendsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/dividends", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of historical prices for a specified stock. Each index in the
// array represents a single tick in a price chart.
func (r *APIV2MarketDataStockService) GetHistoricalPrices(ctx context.Context, stockID string, query APIV2MarketDataStockGetHistoricalPricesParams, opts ...option.RequestOption) (res *[]Apiv2MarketDataStockGetHistoricalPricesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/historical_prices/", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the most recent news articles relating to a stock, including a summary of
// the article and a link to the original source
func (r *APIV2MarketDataStockService) GetNews(ctx context.Context, stockID string, query APIV2MarketDataStockGetNewsParams, opts ...option.RequestOption) (res *[]Apiv2MarketDataStockGetNewsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/news", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns a stock quote for a specified stock.
func (r *APIV2MarketDataStockService) GetQuote(ctx context.Context, stockID string, opts ...option.RequestOption) (res *Apiv2MarketDataStockGetQuoteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/quote", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about stock available for trading.
type Apiv2MarketDataStockListResponse struct {
	// Unique identifier for the stock
	ID string `json:"id,required" format:"bigint"`
	// Whether the stock allows for fractional trading. If it is not fractionable,
	// Dinari only supports limit orders for the stock.
	IsFractionable bool `json:"is_fractionable,required"`
	// Stock Name
	Name string `json:"name,required"`
	// Ticker symbol of the stock
	Symbol string `json:"symbol,required"`
	// SEC Central Index Key. Refer to
	// [this link](https://www.sec.gov/submit-filings/filer-support-resources/how-do-i-guides/understand-utilize-edgar-ciks-passphrases-access-codes)
	Cik string `json:"cik,nullable"`
	// Composite FIGI ID. Refer to [this link](https://www.openfigi.com/about/figi)
	CompositeFigi string `json:"composite_figi,nullable"`
	// CUSIP ID. Refer to [this link](https://www.cusip.com/identifiers.html)
	Cusip string `json:"cusip,nullable"`
	// Description of the company and what they do/offer.
	Description string `json:"description,nullable"`
	// Name of Stock for application display
	DisplayName string `json:"display_name,nullable"`
	// The URL of the logo of the stock. The preferred format is svg.
	LogoURL string `json:"logo_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		IsFractionable respjson.Field
		Name           respjson.Field
		Symbol         respjson.Field
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
func (r Apiv2MarketDataStockListResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataStockListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of a dividend announcement for a stock.
type Apiv2MarketDataStockGetDividendsResponse struct {
	// Cash amount of the dividend per share owned.
	CashAmount float64 `json:"cash_amount"`
	// Currency in which the dividend is paid.
	Currency string `json:"currency"`
	// Date on which the dividend was announced.
	DeclarationDate string `json:"declaration_date"`
	// Type of dividend. Dividends that have been paid and/or are expected to be paid
	// on consistent schedules are denoted as CD. Special Cash dividends that have been
	// paid that are infrequent or unusual, and/or can not be expected to occur in the
	// future are denoted as SC. Long-Term and Short-Term capital gain distributions
	// are denoted as LT and ST, respectively.
	DividendType string `json:"dividend_type"`
	// Date on or after which a stock is traded without the right to receive the next
	// dividend payment. (If you purchase a stock on or after the ex-dividend date, you
	// will not receive the upcoming dividend.)
	ExDividendDate string `json:"ex_dividend_date"`
	// Frequency of the dividend. The following values are possible:
	//
	//	1 - Annual
	//
	//	2 - Semi-Annual
	//
	//	4 - Quarterly
	//
	//	12 - Monthly
	//
	//	52 - Weekly
	//
	//	365 - Daily
	Frequency int64 `json:"frequency"`
	// Date that the dividend is paid out.
	PayDate string `json:"pay_date"`
	// Date that the stock must be held to receive the dividend; set by the company.
	RecordDate string `json:"record_date"`
	// Ticker symbol of the stock.
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
func (r Apiv2MarketDataStockGetDividendsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataStockGetDividendsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Datapoint of historical price data for a stock.
type Apiv2MarketDataStockGetHistoricalPricesResponse struct {
	// Close price of the stock in the given time period.
	Close float64 `json:"close,required"`
	// Highest price of the stock in the given time period.
	High float64 `json:"high,required"`
	// Lowest price of the stock in the given time period.
	Low float64 `json:"low,required"`
	// Open price of the stock in the given time period.
	Open float64 `json:"open,required"`
	// The Unix timestamp in seconds for the start of the aggregate window.
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
func (r Apiv2MarketDataStockGetHistoricalPricesResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataStockGetHistoricalPricesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This represents a news article relating to a stock ticker symbol which includes
// a summary of the article and a link to the original source.
type Apiv2MarketDataStockGetNewsResponse struct {
	// URL of the news article
	ArticleURL string `json:"article_url,required"`
	// Description of the news article
	Description string `json:"description,required"`
	// URL of the image for the news article
	ImageURL string `json:"image_url,required"`
	// Timestamp when the article was published
	PublishedDt time.Time `json:"published_dt,required" format:"date-time"`
	// The publisher of the news article
	Publisher string `json:"publisher,required"`
	// The mobile friendly Accelerated Mobile Page (AMP) URL of the news article if
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
func (r Apiv2MarketDataStockGetNewsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataStockGetNewsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Apiv2MarketDataStockGetQuoteResponse struct {
	// The ask price.
	Price   float64 `json:"price,required"`
	StockID string  `json:"stock_id,required" format:"bigint"`
	// The change in price from the previous close.
	Change float64 `json:"change"`
	// The percentage change in price from the previous close.
	ChangePercent float64 `json:"change_percent"`
	// The close price for the stock in the given time period.
	Close float64 `json:"close"`
	// The highest price for the stock in the given time period
	High float64 `json:"high"`
	// The lowest price for the stock in the given time period.
	Low float64 `json:"low"`
	// The most recent close price of the ticker multiplied by weighted outstanding
	// shares
	MarketCap int64 `json:"market_cap"`
	// The open price for the stock in the given time period.
	Open float64 `json:"open"`
	// The close price for the stock for the previous trading day.
	PreviousClose float64 `json:"previous_close"`
	// The trading volume of the stock in the given time period.
	Volume float64 `json:"volume"`
	// The number of shares outstanding in the given time period
	WeightedSharesOutstanding int64 `json:"weighted_shares_outstanding"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price                     respjson.Field
		StockID                   respjson.Field
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
func (r Apiv2MarketDataStockGetQuoteResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataStockGetQuoteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2MarketDataStockListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// List of stock symbols to query. If not provided, all stocks are returned.
	Symbols []string `query:"symbols,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2MarketDataStockListParams]'s query parameters as
// `url.Values`.
func (r APIV2MarketDataStockListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIV2MarketDataStockGetHistoricalPricesParams struct {
	// The timespan of the historical prices to query.
	//
	// Any of "DAY", "WEEK", "MONTH", "YEAR".
	Timespan APIV2MarketDataStockGetHistoricalPricesParamsTimespan `query:"timespan,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2MarketDataStockGetHistoricalPricesParams]'s query
// parameters as `url.Values`.
func (r APIV2MarketDataStockGetHistoricalPricesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The timespan of the historical prices to query.
type APIV2MarketDataStockGetHistoricalPricesParamsTimespan string

const (
	APIV2MarketDataStockGetHistoricalPricesParamsTimespanDay   APIV2MarketDataStockGetHistoricalPricesParamsTimespan = "DAY"
	APIV2MarketDataStockGetHistoricalPricesParamsTimespanWeek  APIV2MarketDataStockGetHistoricalPricesParamsTimespan = "WEEK"
	APIV2MarketDataStockGetHistoricalPricesParamsTimespanMonth APIV2MarketDataStockGetHistoricalPricesParamsTimespan = "MONTH"
	APIV2MarketDataStockGetHistoricalPricesParamsTimespanYear  APIV2MarketDataStockGetHistoricalPricesParamsTimespan = "YEAR"
)

type APIV2MarketDataStockGetNewsParams struct {
	// The number of news articles to return, default is 10 max is 25
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2MarketDataStockGetNewsParams]'s query parameters as
// `url.Values`.
func (r APIV2MarketDataStockGetNewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
