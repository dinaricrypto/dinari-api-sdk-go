// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apiquery"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
)

// V2AccountActivityService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountActivityService] method instead.
type V2AccountActivityService struct {
	Options []option.RequestOption
}

// NewV2AccountActivityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountActivityService(opts ...option.RequestOption) (r V2AccountActivityService) {
	r = V2AccountActivityService{}
	r.Options = opts
	return
}

// Get a list of brokerage activities tied to the specified `Account`.
//
// **⚠️ ALPHA: This endpoint is in early development and subject to breaking
// changes.**
func (r *V2AccountActivityService) GetBrokerage(ctx context.Context, accountID string, query V2AccountActivityGetBrokerageParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/activities/brokerage", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type V2AccountActivityGetBrokerageParams struct {
	// The maximum number of entries to return in the response. Defaults to 100.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Pagination token. Set to the `id` field of the last Activity returned in the
	// previous page to get the next page of results.
	PageToken param.Opt[string] `query:"page_token,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountActivityGetBrokerageParams]'s query parameters as
// `url.Values`.
func (r V2AccountActivityGetBrokerageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
