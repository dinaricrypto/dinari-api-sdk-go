// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"net/http"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// APIV2Service contains methods and other services that help with interacting with
// the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2Service] method instead.
type APIV2Service struct {
	Options    []option.RequestOption
	MarketData APIV2MarketDataService
	Entities   APIV2EntityService
	Accounts   APIV2AccountService
}

// NewAPIV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIV2Service(opts ...option.RequestOption) (r APIV2Service) {
	r = APIV2Service{}
	r.Options = opts
	r.MarketData = NewAPIV2MarketDataService(opts...)
	r.Entities = NewAPIV2EntityService(opts...)
	r.Accounts = NewAPIV2AccountService(opts...)
	return
}

// Get Health Status
func (r *APIV2Service) GetHealth(ctx context.Context, opts ...option.RequestOption) (res *Apiv2GetHealthResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/_health/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Apiv2GetHealthResponse struct {
	// Status of server
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2GetHealthResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2GetHealthResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
