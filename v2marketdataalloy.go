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

// **Dinari provides basic market data for `Stocks` and `Alloys` that are available
// to transact on.**
//
// This data is provided on a best-effort basis and we recommend using a dedicated
// provider for more intensive market data needs.
//
// V2MarketDataAlloyService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MarketDataAlloyService] method instead.
type V2MarketDataAlloyService struct {
	Options []option.RequestOption
}

// NewV2MarketDataAlloyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2MarketDataAlloyService(opts ...option.RequestOption) (r V2MarketDataAlloyService) {
	r = V2MarketDataAlloyService{}
	r.Options = opts
	return
}

// Returns available `Alloys` with cursor-based pagination.
func (r *V2MarketDataAlloyService) List(ctx context.Context, query V2MarketDataAlloyListParams, opts ...option.RequestOption) (res *V2MarketDataAlloyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/market_data/alloys/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get the current price for a specified `Alloy`.
func (r *V2MarketDataAlloyService) GetCurrentPrice(ctx context.Context, alloyID string, opts ...option.RequestOption) (res *V2MarketDataAlloyGetCurrentPriceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if alloyID == "" {
		err = errors.New("missing required alloy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/alloys/%s/current_price", alloyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get historical price data for a specified `Alloy`. Each index in the array
// represents a single tick in a price chart.
func (r *V2MarketDataAlloyService) GetHistoricalPrices(ctx context.Context, alloyID string, query V2MarketDataAlloyGetHistoricalPricesParams, opts ...option.RequestOption) (res *[]V2MarketDataAlloyGetHistoricalPricesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if alloyID == "" {
		err = errors.New("missing required alloy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/market_data/alloys/%s/historical_prices/", alloyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type PaginationMetadata struct {
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
func (r PaginationMetadata) RawJSON() string { return r.JSON.raw }
func (r *PaginationMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Paginated response containing a list of Alloys.
type V2MarketDataAlloyListResponse struct {
	// List of Alloys
	Data []V2MarketDataAlloyListResponseData `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata PaginationMetadata `json:"pagination_metadata" api:"required"`
	// Schema version
	//
	// Any of "PaginatedAlloyResponse:v1".
	Sv V2MarketDataAlloyListResponse_Sv `json:"_sv"`
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
func (r V2MarketDataAlloyListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataAlloyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Alloy details.
type V2MarketDataAlloyListResponseData struct {
	// Unique identifier of the Alloy asset
	ID string `json:"id" api:"required" format:"uuid"`
	// Indicates if tradable on Dinari platform
	IsTradable bool `json:"is_tradable" api:"required"`
	// Name of the Alloy
	Name string `json:"name" api:"required"`
	// Symbol of the Alloy
	Symbol string `json:"symbol" api:"required"`
	// Schema version
	//
	// Any of "Alloy:v1".
	Sv string `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsTradable  respjson.Field
		Name        respjson.Field
		Symbol      respjson.Field
		Sv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataAlloyListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataAlloyListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema version
type V2MarketDataAlloyListResponse_Sv string

const (
	V2MarketDataAlloyListResponse_SvPaginatedAlloyResponseV1 V2MarketDataAlloyListResponse_Sv = "PaginatedAlloyResponse:v1"
)

type V2MarketDataAlloyGetCurrentPriceResponse struct {
	// ID of the `Alloy` asset.
	ID string `json:"id" api:"required" format:"uuid"`
	// Current trade price.
	Price float64 `json:"price" api:"required"`
	// When the price was generated.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// Schema version
	//
	// Any of "AlloyPrice:v1".
	Sv V2MarketDataAlloyGetCurrentPriceResponse_Sv `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Price       respjson.Field
		Timestamp   respjson.Field
		Sv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataAlloyGetCurrentPriceResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataAlloyGetCurrentPriceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema version
type V2MarketDataAlloyGetCurrentPriceResponse_Sv string

const (
	V2MarketDataAlloyGetCurrentPriceResponse_SvAlloyPriceV1 V2MarketDataAlloyGetCurrentPriceResponse_Sv = "AlloyPrice:v1"
)

// Datapoint of historical price data for an `Alloy`.
type V2MarketDataAlloyGetHistoricalPricesResponse struct {
	// Close price from the given time period.
	Close float64 `json:"close" api:"required"`
	// High price from the given time period.
	High float64 `json:"high" api:"required"`
	// Low price from the given time period.
	Low float64 `json:"low" api:"required"`
	// Open price from the given time period.
	Open float64 `json:"open" api:"required"`
	// UNIX timestamp in seconds for the start of the aggregate window.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Schema version
	//
	// Any of "AlloyHistoricalPriceDataPointV1:v1".
	Sv V2MarketDataAlloyGetHistoricalPricesResponse_Sv `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Close       respjson.Field
		High        respjson.Field
		Low         respjson.Field
		Open        respjson.Field
		Timestamp   respjson.Field
		Sv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataAlloyGetHistoricalPricesResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataAlloyGetHistoricalPricesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema version
type V2MarketDataAlloyGetHistoricalPricesResponse_Sv string

const (
	V2MarketDataAlloyGetHistoricalPricesResponse_SvAlloyHistoricalPriceDataPointV1V1 V2MarketDataAlloyGetHistoricalPricesResponse_Sv = "AlloyHistoricalPriceDataPointV1:v1"
)

type V2MarketDataAlloyListParams struct {
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2MarketDataAlloyListParamsOrder `query:"order,omitzero" json:"-"`
	// If set, this endpoint will return Alloys that match the symbols specified
	Symbols []string `query:"symbols,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataAlloyListParams]'s query parameters as
// `url.Values`.
func (r V2MarketDataAlloyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2MarketDataAlloyListParamsOrder string

const (
	V2MarketDataAlloyListParamsOrderAsc  V2MarketDataAlloyListParamsOrder = "asc"
	V2MarketDataAlloyListParamsOrderDesc V2MarketDataAlloyListParamsOrder = "desc"
)

type V2MarketDataAlloyGetHistoricalPricesParams struct {
	// The timespan of the historical prices to query.
	//
	// Any of "DAY", "WEEK".
	Timespan V2MarketDataAlloyGetHistoricalPricesParamsTimespan `query:"timespan,omitzero" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V2MarketDataAlloyGetHistoricalPricesParams]'s query
// parameters as `url.Values`.
func (r V2MarketDataAlloyGetHistoricalPricesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The timespan of the historical prices to query.
type V2MarketDataAlloyGetHistoricalPricesParamsTimespan string

const (
	V2MarketDataAlloyGetHistoricalPricesParamsTimespanDay  V2MarketDataAlloyGetHistoricalPricesParamsTimespan = "DAY"
	V2MarketDataAlloyGetHistoricalPricesParamsTimespanWeek V2MarketDataAlloyGetHistoricalPricesParamsTimespan = "WEEK"
)
