// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// APIV2AccountOrderFulfillmentService contains methods and other services that
// help with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountOrderFulfillmentService] method instead.
type APIV2AccountOrderFulfillmentService struct {
	Options []option.RequestOption
}

// NewAPIV2AccountOrderFulfillmentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIV2AccountOrderFulfillmentService(opts ...option.RequestOption) (r APIV2AccountOrderFulfillmentService) {
	r = APIV2AccountOrderFulfillmentService{}
	r.Options = opts
	return
}

// Retrieves details of a specific order fulfillment by its ID.
func (r *APIV2AccountOrderFulfillmentService) Get(ctx context.Context, fulfillmentID string, query APIV2AccountOrderFulfillmentGetParams, opts ...option.RequestOption) (res *OrderFulfillment, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if fulfillmentID == "" {
		err = errors.New("missing required fulfillment_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_fulfillments/%s", query.AccountID, fulfillmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Queries all order fulfillments under the account.
func (r *APIV2AccountOrderFulfillmentService) Query(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]OrderFulfillment, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_fulfillments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about a fulfillment of an order. An order may be fulfilled in
// multiple transactions.
type OrderFulfillment struct {
	// Identifier of the order fulfillment
	ID string `json:"id,required" format:"uuid"`
	// Amount of asset token filled
	AssetTokenFilled float64 `json:"asset_token_filled,required"`
	// Amount of asset token spent
	AssetTokenSpent float64 `json:"asset_token_spent,required"`
	// Identifier of the order this fulfillment is for
	OrderID string `json:"order_id,required" format:"uuid"`
	// Amount of payment token filled
	PaymentTokenFilled float64 `json:"payment_token_filled,required"`
	// Amount of payment token spent
	PaymentTokenSpent float64 `json:"payment_token_spent,required"`
	// Time when transaction occurred
	TransactionDt time.Time `json:"transaction_dt,required" format:"date-time"`
	// Transaction hash for this fulfillment
	TransactionHash string `json:"transaction_hash,required" format:"hex_string"`
	// Fee amount of payment token spent
	PaymentTokenFee float64 `json:"payment_token_fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AssetTokenFilled   respjson.Field
		AssetTokenSpent    respjson.Field
		OrderID            respjson.Field
		PaymentTokenFilled respjson.Field
		PaymentTokenSpent  respjson.Field
		TransactionDt      respjson.Field
		TransactionHash    respjson.Field
		PaymentTokenFee    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderFulfillment) RawJSON() string { return r.JSON.raw }
func (r *OrderFulfillment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2AccountOrderFulfillmentGetParams struct {
	AccountID string `path:"account_id,required" format:"uuid" json:"-"`
	paramObj
}
