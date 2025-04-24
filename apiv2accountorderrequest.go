// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2AccountOrderRequestService contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountOrderRequestService] method instead.
type APIV2AccountOrderRequestService struct {
	Options []option.RequestOption
}

// NewAPIV2AccountOrderRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIV2AccountOrderRequestService(opts ...option.RequestOption) (r APIV2AccountOrderRequestService) {
	r = APIV2AccountOrderRequestService{}
	r.Options = opts
	return
}

// Retrieves details of a specific managed order request by its ID.
func (r *APIV2AccountOrderRequestService) Get(ctx context.Context, accountID string, requestID string, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/%s", accountID, requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists managed order requests.
func (r *APIV2AccountOrderRequestService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a managed limit buy request.
func (r *APIV2AccountOrderRequestService) NewLimitBuy(ctx context.Context, accountID string, body APIV2AccountOrderRequestNewLimitBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a managed limit sell request.
func (r *APIV2AccountOrderRequestService) NewLimitSell(ctx context.Context, accountID string, body APIV2AccountOrderRequestNewLimitSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/limit_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a managed market buy request.
func (r *APIV2AccountOrderRequestService) NewMarketBuy(ctx context.Context, accountID string, body APIV2AccountOrderRequestNewMarketBuyParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_buy", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates a managed market sell request.
func (r *APIV2AccountOrderRequestService) NewMarketSell(ctx context.Context, accountID string, body APIV2AccountOrderRequestNewMarketSellParams, opts ...option.RequestOption) (res *OrderRequest, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_requests/market_sell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Input parameters for placing a limit order.
//
// The properties AssetQuantity, LimitPrice, StockID are required.
type LimitOrderRequestInputParam struct {
	// Quantity of stock to trade. Must be a positive integer.
	AssetQuantity int64 `json:"asset_quantity,required"`
	// Price at which to execute the order. Must be a positive number with a precision
	// of up to 2 decimal places.
	LimitPrice float64 `json:"limit_price,required"`
	// ID of stock, as returned by the `/stocks` endpoint, e.g. 1
	StockID string `json:"stock_id,required" format:"bigint"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f LimitOrderRequestInputParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r LimitOrderRequestInputParam) MarshalJSON() (data []byte, err error) {
	type shadow LimitOrderRequestInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Request to create an order
type OrderRequest struct {
	// ID of account placing the order
	AccountID string `json:"account_id,required" format:"uuid"`
	// Confirmation code of order request. This is the primary identifier for the
	// `/order_requests` endpoint
	ConfirmationCode string `json:"confirmation_code,required" format:"uuid"`
	// Timestamp at which the order request was created.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Status of order request
	//
	// Any of "PENDING", "SUBMITTED", "ERROR", "CANCELLED".
	Status OrderRequestStatus `json:"status,required"`
	// ID of order created from the order request. This is the primary identifier for
	// the `/orders` endpoint
	OrderID string `json:"order_id" format:"uuid"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		AccountID        resp.Field
		ConfirmationCode resp.Field
		CreatedDt        resp.Field
		Status           resp.Field
		OrderID          resp.Field
		ExtraFields      map[string]resp.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderRequest) RawJSON() string { return r.JSON.raw }
func (r *OrderRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of order request
type OrderRequestStatus string

const (
	OrderRequestStatusPending   OrderRequestStatus = "PENDING"
	OrderRequestStatusSubmitted OrderRequestStatus = "SUBMITTED"
	OrderRequestStatusError     OrderRequestStatus = "ERROR"
	OrderRequestStatusCancelled OrderRequestStatus = "CANCELLED"
)

type APIV2AccountOrderRequestNewLimitBuyParams struct {
	// Input parameters for placing a limit order.
	LimitOrderRequestInput LimitOrderRequestInputParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountOrderRequestNewLimitBuyParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2AccountOrderRequestNewLimitBuyParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.LimitOrderRequestInput)
}

type APIV2AccountOrderRequestNewLimitSellParams struct {
	// Input parameters for placing a limit order.
	LimitOrderRequestInput LimitOrderRequestInputParam
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountOrderRequestNewLimitSellParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2AccountOrderRequestNewLimitSellParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.LimitOrderRequestInput)
}

type APIV2AccountOrderRequestNewMarketBuyParams struct {
	// Amount of USD to pay or receive for the order. Must be a positive number with a
	// precision of up to 2 decimal places.
	PaymentAmount float64 `json:"payment_amount,required"`
	// ID of stock, as returned by the `/stocks` endpoint, e.g. 1
	StockID string `json:"stock_id,required" format:"bigint"`
	// Whether to include fees in the `payment_amount` input field.
	IncludeFees param.Opt[bool] `json:"include_fees,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountOrderRequestNewMarketBuyParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2AccountOrderRequestNewMarketBuyParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2AccountOrderRequestNewMarketBuyParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type APIV2AccountOrderRequestNewMarketSellParams struct {
	// Quantity of stock to trade. Must be a positive number with a precision of up to
	// 9 decimal places.
	AssetQuantity float64 `json:"asset_quantity,required"`
	// ID of stock, as returned by the `/stocks` endpoint, e.g. 1
	StockID string `json:"stock_id,required" format:"bigint"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountOrderRequestNewMarketSellParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2AccountOrderRequestNewMarketSellParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2AccountOrderRequestNewMarketSellParams
	return param.MarshalObject(r, (*shadow)(&r))
}
