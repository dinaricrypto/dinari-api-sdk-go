// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"net/http"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/resp"
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

// Returns an object containing the market hours for the current day and next open
// trading day.
func (r *APIV2MarketDataService) GetMarketHours(ctx context.Context, opts ...option.RequestOption) (res *Apiv2MarketDataGetMarketHoursResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/market_data/market_hours/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Apiv2MarketDataGetMarketHoursResponse struct {
	// Whether or not the market is open
	IsMarketOpen bool `json:"is_market_open,required"`
	// Timestamp in ISO 8601 format at which the next session closes
	NextSessionCloseDt time.Time `json:"next_session_close_dt,required" format:"date-time"`
	// Timestamp in ISO 8601 format at which the next session opens
	NextSessionOpenDt time.Time `json:"next_session_open_dt,required" format:"date-time"`
	// Timestamp in ISO 8601 format at which the current session closes or null if the
	// market is currently closed
	CurrentSessionCloseDt time.Time `json:"current_session_close_dt" format:"date-time"`
	// Timestamp in ISO 8601 format at which the current session opened or null if the
	// market is currently closed
	CurrentSessionOpenDt time.Time `json:"current_session_open_dt" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		IsMarketOpen          resp.Field
		NextSessionCloseDt    resp.Field
		NextSessionOpenDt     resp.Field
		CurrentSessionCloseDt resp.Field
		CurrentSessionOpenDt  resp.Field
		ExtraFields           map[string]resp.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2MarketDataGetMarketHoursResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2MarketDataGetMarketHoursResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
