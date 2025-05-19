// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"net/http"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// APIV2MarketDataService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2MarketDataService] method instead.
type APIV2MarketDataService struct {
	Options []option.RequestOption
	Stocks  APIV2MarketDataStockService
}

// NewAPIV2MarketDataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV2MarketDataService(opts ...option.RequestOption) (r APIV2MarketDataService) {
	r = APIV2MarketDataService{}
	r.Options = opts
	r.Stocks = NewAPIV2MarketDataStockService(opts...)
	return
}

// Get the market hours for the current trading session and next open trading
// session.
func (r *APIV2MarketDataService) GetMarketHours(ctx context.Context, opts ...option.RequestOption) (res *Apiv2MarketDataGetMarketHoursResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/market_hours/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Apiv2MarketDataGetMarketHoursResponse struct {
	// Whether or not the market is open
	IsMarketOpen bool `json:"is_market_open,required"`
	// Datetime at which the next session closes. ISO 8601 timestamp.
	NextSessionCloseDt time.Time `json:"next_session_close_dt,required" format:"date-time"`
	// Datetime at which the next session opens. ISO 8601 timestamp.
	NextSessionOpenDt time.Time `json:"next_session_open_dt,required" format:"date-time"`
	// Datetime at which the current session closes. `null` if the market is currently
	// closed. ISO 8601 timestamp.
	CurrentSessionCloseDt time.Time `json:"current_session_close_dt" format:"date-time"`
	// Datetime at which the current session opened. `null` if the market is currently
	// closed. ISO 8601 timestamp.
	CurrentSessionOpenDt time.Time `json:"current_session_open_dt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsMarketOpen          respjson.Field
		NextSessionCloseDt    respjson.Field
		NextSessionOpenDt     respjson.Field
		CurrentSessionCloseDt respjson.Field
		CurrentSessionOpenDt  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2MarketDataGetMarketHoursResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataGetMarketHoursResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
