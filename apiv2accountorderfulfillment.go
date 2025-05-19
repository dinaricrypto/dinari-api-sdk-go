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

// Query `OrderFulfillments` under the `Account`.
func (r *APIV2AccountOrderFulfillmentService) Query(ctx context.Context, accountID string, query APIV2AccountOrderFulfillmentQueryParams, opts ...option.RequestOption) (res *[]OrderFulfillment, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/order_fulfillments", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about a fulfillment of an `Order`. An order may be fulfilled in
// multiple transactions.
type OrderFulfillment struct {
	// ID of the `OrderFulfillment`.
	ID string `json:"id,required" format:"uuid"`
	// Amount of dShare asset token filled for `BUY` orders.
	AssetTokenFilled float64 `json:"asset_token_filled,required"`
	// Amount of dShare asset token spent for `SELL` orders.
	AssetTokenSpent float64 `json:"asset_token_spent,required"`
	// Blockchain that the transaction was run on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866".
	ChainID OrderFulfillmentChainID `json:"chain_id,required"`
	// ID of the `Order` this `OrderFulfillment` is for.
	OrderID string `json:"order_id,required" format:"uuid"`
	// Amount of payment token filled for `SELL` orders.
	PaymentTokenFilled float64 `json:"payment_token_filled,required"`
	// Amount of payment token spent for `BUY` orders.
	PaymentTokenSpent float64 `json:"payment_token_spent,required"`
	// Time when transaction occurred.
	TransactionDt time.Time `json:"transaction_dt,required" format:"date-time"`
	// Transaction hash for this fulfillment.
	TransactionHash string `json:"transaction_hash,required" format:"hex_string"`
	// Fee amount, in payment tokens.
	PaymentTokenFee float64 `json:"payment_token_fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AssetTokenFilled   respjson.Field
		AssetTokenSpent    respjson.Field
		ChainID            respjson.Field
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

// Blockchain that the transaction was run on.
type OrderFulfillmentChainID string

const (
	OrderFulfillmentChainIDEip155_1     OrderFulfillmentChainID = "eip155:1"
	OrderFulfillmentChainIDEip155_42161 OrderFulfillmentChainID = "eip155:42161"
	OrderFulfillmentChainIDEip155_8453  OrderFulfillmentChainID = "eip155:8453"
	OrderFulfillmentChainIDEip155_81457 OrderFulfillmentChainID = "eip155:81457"
	OrderFulfillmentChainIDEip155_7887  OrderFulfillmentChainID = "eip155:7887"
	OrderFulfillmentChainIDEip155_98866 OrderFulfillmentChainID = "eip155:98866"
)

type APIV2AccountOrderFulfillmentQueryParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// List of `Order` IDs to query `OrderFulfillments` for.
	OrderIDs []string `query:"order_ids,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2AccountOrderFulfillmentQueryParams]'s query parameters
// as `url.Values`.
func (r APIV2AccountOrderFulfillmentQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
