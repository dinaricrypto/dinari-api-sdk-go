// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/resp"
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

// Retrieves a specific account by its ID.
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

// Sets the account to be inactive.
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

// Retrieves the cash amount in the account.
func (r *APIV2AccountService) GetCash(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Apiv2AccountGetCashResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/cash", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves dividend payments made to the account.
func (r *APIV2AccountService) GetDividendPayments(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]Apiv2AccountGetDividendPaymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/dividend_payments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves interest payments made to the account.
func (r *APIV2AccountService) GetInterestPayments(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]Apiv2AccountGetInterestPaymentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/interest_payments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the portfolio of the account, sans cash equivalents.
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

// Balance of cash and cash equivalents currently in the account.
type Apiv2AccountGetCashResponse struct {
	// Total amount of cash and cash equivalents
	Amount float64 `json:"amount,required"`
	// Currency (e.g. USD)
	Currency string `json:"currency,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Amount      resp.Field
		Currency    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetCashResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetCashResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountDividendPayment represents a stock dividend payment event for an account.
type Apiv2AccountGetDividendPaymentsResponse struct {
	// Amount of the dividend paid.
	Amount float64 `json:"amount,required"`
	// Currency in which the dividend was paid. (e.g. USD)
	Currency string `json:"currency,required"`
	// Date the dividend was distributed to the account.
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// ID of the stock for which the dividend was paid.
	StockID string `json:"stock_id,required" format:"bigint"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Amount      resp.Field
		Currency    resp.Field
		PaymentDate resp.Field
		StockID     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetDividendPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetDividendPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object representing interest payment details.
type Apiv2AccountGetInterestPaymentsResponse struct {
	// Amount of interest paid
	Amount float64 `json:"amount,required"`
	// Type of currency (e.g. USD)
	Currency string `json:"currency,required"`
	// Date of interest payment. In US Eastern time zone
	PaymentDate time.Time `json:"payment_date,required" format:"date"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Amount      resp.Field
		Currency    resp.Field
		PaymentDate resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetInterestPaymentsResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetInterestPaymentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// This is an object representing the balance of cash and stock assets in your
// account.
type Apiv2AccountGetPortfolioResponse struct {
	// Stock Balance details for all owned stocks
	Assets []Apiv2AccountGetPortfolioResponseAsset `json:"assets,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Assets      resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetPortfolioResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetPortfolioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Balance of a stock in the account
type Apiv2AccountGetPortfolioResponseAsset struct {
	// Total amount of the stock
	Amount float64 `json:"amount,required"`
	// Total market value of the stock
	MarketValue float64 `json:"market_value,required"`
	// ID of Stock
	StockID string `json:"stock_id,required" format:"bigint"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Amount      resp.Field
		MarketValue resp.Field
		StockID     resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountGetPortfolioResponseAsset) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountGetPortfolioResponseAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
