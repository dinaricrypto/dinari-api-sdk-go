// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

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

// V2MarketDataStockSplitService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MarketDataStockSplitService] method instead.
type V2MarketDataStockSplitService struct {
	Options []option.RequestOption
}

// NewV2MarketDataStockSplitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2MarketDataStockSplitService(opts ...option.RequestOption) (r V2MarketDataStockSplitService) {
	r = V2MarketDataStockSplitService{}
	r.Options = opts
	return
}

// Get a list of stock splits for `Stocks` available for trade via Dinari. The
// splits are ordered by the date they were created, with the most recent split
// first.
//
// In an example 10-for-1 stock split, trading will be halted for the stock at the
// end of the `payable_date`, as the split transitions from `PENDING` to
// `IN_PROGRESS`. This usually occurs over a weekend to minimize trading
// disruptions. Each share of stock owned by a shareholder will then be converted
// into 10 shares, and the split becomes `COMPLETE` as trading resumes on the
// `ex_date` with new split-adjusted prices.
func (r *V2MarketDataStockSplitService) List(ctx context.Context, query V2MarketDataStockSplitListParams, opts ...option.RequestOption) (res *[]StockSplit, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/stocks/splits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of stock splits for a specific `Stock`. The splits are ordered by the
// date they were created, with the most recent split first.
//
// In an example 10-for-1 stock split, trading will be halted for the stock at the
// end of the `payable_date`, as the split transitions from `PENDING` to
// `IN_PROGRESS`. This usually occurs over a weekend to minimize trading
// disruptions. Each share of stock owned by a shareholder will then be converted
// into 10 shares, and the split becomes `COMPLETE` as trading resumes on the
// `ex_date` with new split-adjusted prices.
func (r *V2MarketDataStockSplitService) ListForStock(ctx context.Context, stockID string, query V2MarketDataStockSplitListForStockParams, opts ...option.RequestOption) (res *[]StockSplit, err error) {
	opts = append(r.Options[:], opts...)
	if stockID == "" {
		err = errors.New("missing required stock_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/market_data/stocks/%s/splits", stockID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about a stock split, including the `Stock` ID, the number of shares
// before and after the split, the record date, payable date, ex-date, and the
// status of the split.
type StockSplit struct {
	// ID of the `StockSplit`
	ID string `json:"id,required" format:"uuid"`
	// Ex-date of the split in Eastern Time Zone. First day the stock trades at
	// post-split prices. Typically is last date in the process, and the main important
	// date for investors. In ISO 8601 format, YYYY-MM-DD.
	ExDate time.Time `json:"ex_date,required" format:"date"`
	// Payable date of the split in Eastern Time Zone. This is the date when company
	// will send out the new shares. Mainly for record keeping by brokerages, who
	// forward the shares to eventual owners. Typically is the second date in the
	// process. In ISO 8601 format, YYYY-MM-DD.
	PayableDate time.Time `json:"payable_date,required" format:"date"`
	// Record date of the split in Eastern Time Zone, for company to determine where to
	// send their new shares. Mainly for record keeping by brokerages, who forward the
	// shares to eventual owners. Typically is the first date in the process. In ISO
	// 8601 format, YYYY-MM-DD.
	RecordDate time.Time `json:"record_date,required" format:"date"`
	// The number of shares before the split. In a 10-for-1 split, this would be 1.
	SplitFrom float64 `json:"split_from,required"`
	// The number of shares after the split. In a 10-for-1 split, this would be 10.
	SplitTo float64 `json:"split_to,required"`
	// The status of Dinari's processing of the `StockSplit`. `Stocks` for which this
	// status is `IN_PROGRESS` will not be available for trading.
	//
	// Any of "PENDING", "IN_PROGRESS", "COMPLETE".
	Status StockSplitStatus `json:"status,required"`
	// ID of the `Stock` whose shares are being split.
	StockID string `json:"stock_id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExDate      respjson.Field
		PayableDate respjson.Field
		RecordDate  respjson.Field
		SplitFrom   respjson.Field
		SplitTo     respjson.Field
		Status      respjson.Field
		StockID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StockSplit) RawJSON() string { return r.JSON.raw }
func (r *StockSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of Dinari's processing of the `StockSplit`. `Stocks` for which this
// status is `IN_PROGRESS` will not be available for trading.
type StockSplitStatus string

const (
	StockSplitStatusPending    StockSplitStatus = "PENDING"
	StockSplitStatusInProgress StockSplitStatus = "IN_PROGRESS"
	StockSplitStatusComplete   StockSplitStatus = "COMPLETE"
)

type V2MarketDataStockSplitListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockSplitListParams]'s query parameters as
// `url.Values`.
func (r V2MarketDataStockSplitListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V2MarketDataStockSplitListForStockParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataStockSplitListForStockParams]'s query
// parameters as `url.Values`.
func (r V2MarketDataStockSplitListForStockParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
