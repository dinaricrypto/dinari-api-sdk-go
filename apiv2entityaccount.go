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
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2EntityAccountService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2EntityAccountService] method instead.
type APIV2EntityAccountService struct {
	Options []option.RequestOption
}

// NewAPIV2EntityAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV2EntityAccountService(opts ...option.RequestOption) (r APIV2EntityAccountService) {
	r = APIV2EntityAccountService{}
	r.Options = opts
	return
}

// Creates a new Account for the given Entity.
func (r *APIV2EntityAccountService) New(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieves a list of Accounts that belong to a specific Entity.
func (r *APIV2EntityAccountService) List(ctx context.Context, entityID string, query APIV2EntityAccountListParams, opts ...option.RequestOption) (res *[]Account, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/accounts", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Information about an account owned by an entity
type Account struct {
	// Unique identifier for the account
	ID string `json:"id,required" format:"uuid"`
	// Timestamp when the account was created
	CreatedDt time.Time `json:"created_dt,required" format:"date-time"`
	// Identifier for the Entity that owns the account
	EntityID string `json:"entity_id,required" format:"uuid"`
	// Indicates whether the account is active
	IsActive bool `json:"is_active,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID          resp.Field
		CreatedDt   resp.Field
		EntityID    resp.Field
		IsActive    resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Account) RawJSON() string { return r.JSON.raw }
func (r *Account) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2EntityAccountListParams struct {
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2EntityAccountListParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

// URLQuery serializes [APIV2EntityAccountListParams]'s query parameters as
// `url.Values`.
func (r APIV2EntityAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
