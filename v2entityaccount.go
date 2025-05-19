// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/apiquery"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/respjson"
)

// V2EntityAccountService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityAccountService] method instead.
type V2EntityAccountService struct {
	Options []option.RequestOption
}

// NewV2EntityAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2EntityAccountService(opts ...option.RequestOption) (r V2EntityAccountService) {
	r = V2EntityAccountService{}
	r.Options = opts
	return
}

// Create a new `Account` for a specific `Entity`. This `Entity` represents your
// organization itself, or an individual customer of your organization.
func (r *V2EntityAccountService) New(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get a list of all `Accounts` that belong to a specific `Entity`. This `Entity`
// represents your organization itself, or an individual customer of your
// organization.
func (r *V2EntityAccountService) List(ctx context.Context, entityID string, query V2EntityAccountListParams, opts ...option.RequestOption) (res *[]Account, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about an `Account` owned by an `Entity`.
type Account struct {
	// Unique ID for the `Account`.
	ID string `json:"id,required" format:"uuid"`
	// Datetime when the `Account` was created. ISO 8601 timestamp.
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// ID for the `Entity` that owns the `Account`.
	EntityID string `json:"entity_id,required" format:"uuid"`
	// Indicates whether the `Account` is active.
	IsActive bool `json:"is_active,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedDt   respjson.Field
		EntityID    respjson.Field
		IsActive    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityAccountListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2EntityAccountListParams]'s query parameters as
// `url.Values`.
func (r V2EntityAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
