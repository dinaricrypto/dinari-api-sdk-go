// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"net/http"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// V2MarketDataService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MarketDataService] method instead.
type V2MarketDataService struct {
	Options []option.RequestOption
	Stocks  V2MarketDataStockService
}

// NewV2MarketDataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2MarketDataService(opts ...option.RequestOption) (r V2MarketDataService) {
	r = V2MarketDataService{}
	r.Options = opts
	r.Stocks = NewV2MarketDataStockService(opts...)
	return
}

// Get the market hours for the current trading session and next open trading
// session.
func (r *V2MarketDataService) GetMarketHours(ctx context.Context, opts ...option.RequestOption) (res *V2MarketDataGetMarketHoursResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/market_hours/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V2MarketDataGetMarketHoursResponse struct {
	// Whether or not the market is open
	IsMarketOpen bool `json:"is_market_open,required"`
	// Datetime at which the next session closes. ISO 8601 timestamp.
	NextSessionCloseDt time.Time `json:"next_session_close_dt,required" format:"date-time"`
	// Datetime at which the next session opens. ISO 8601 timestamp.
	NextSessionOpenDt time.Time `json:"next_session_open_dt,required" format:"date-time"`
	// Time at which the current session after-hours end.
	CurrentSessionAfterHoursCloseTimeDt time.Time `json:"current_session_after_hours_close_time_dt" format:"date-time"`
	// Datetime at which the current session closes. `null` if the market is currently
	// closed. ISO 8601 timestamp.
	CurrentSessionCloseDt time.Time `json:"current_session_close_dt" format:"date-time"`
	// Datetime at which the current session opened. `null` if the market is currently
	// closed. ISO 8601 timestamp.
	CurrentSessionOpenDt time.Time `json:"current_session_open_dt" format:"date-time"`
	// Time at which the current session overnight starts.
	CurrentSessionOvernightOpenTimeDt time.Time `json:"current_session_overnight_open_time_dt" format:"date-time"`
	// Time at which the current session pre-market hours start.
	CurrentSessionPreMarketOpenTimeDt time.Time `json:"current_session_pre_market_open_time_dt" format:"date-time"`
	// Time at which the next session after-hours end.
	NextSessionAfterHoursCloseTimeDt time.Time `json:"next_session_after_hours_close_time_dt" format:"date-time"`
	// Time at which the next session overnight starts.
	NextSessionOvernightOpenTimeDt time.Time `json:"next_session_overnight_open_time_dt" format:"date-time"`
	// Time at which the next session pre-market hours start.
	NextSessionPreMarketOpenTimeDt time.Time `json:"next_session_pre_market_open_time_dt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsMarketOpen                        respjson.Field
		NextSessionCloseDt                  respjson.Field
		NextSessionOpenDt                   respjson.Field
		CurrentSessionAfterHoursCloseTimeDt respjson.Field
		CurrentSessionCloseDt               respjson.Field
		CurrentSessionOpenDt                respjson.Field
		CurrentSessionOvernightOpenTimeDt   respjson.Field
		CurrentSessionPreMarketOpenTimeDt   respjson.Field
		NextSessionAfterHoursCloseTimeDt    respjson.Field
		NextSessionOvernightOpenTimeDt      respjson.Field
		NextSessionPreMarketOpenTimeDt      respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2MarketDataGetMarketHoursResponse) RawJSON() string { return r.JSON.raw }
func (r *V2MarketDataGetMarketHoursResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
