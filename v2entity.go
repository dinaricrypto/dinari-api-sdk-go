// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apiquery"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// **`Entities` represent a business or organization that uses the API, and their
// customers.**
//
// Dinari Partners are represented as an organization `Entity` in the API, with
// their own accounts. Individual customers of Partner `Entities` are also
// represented as `Entities` in the API, which are managed by the Partner `Entity`.
//
// V2EntityService contains methods and other services that help with interacting
// with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityService] method instead.
type V2EntityService struct {
	Options []option.RequestOption
	// **`Accounts` represent the financial accounts of an `Entity`.**
	//
	// `Orders`, dividends, and other transactions are associated with an `Account`.
	Accounts V2EntityAccountService
	// **KYC (Know Your Customer) is a process of verifying the identity of customer
	// `Entities`.**
	//
	// KYC is required for all customer `Entities` that transact on Dinari's platform.
	//
	// Dinari provides a managed KYC process for its Partners, which provides a
	// convenient KYC flow URL to present to the end customer.
	//
	// For Dinari Partners that supply their own KYC data, the API provides a way to
	// record a customer's KYC information using the Partner's KYC data. This requires
	// an existing KYC agreement between Dinari and the Partner.
	KYC V2EntityKYCService
}

// NewV2EntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2EntityService(opts ...option.RequestOption) (r V2EntityService) {
	r = V2EntityService{}
	r.Options = opts
	r.Accounts = NewV2EntityAccountService(opts...)
	r.KYC = NewV2EntityKYCService(opts...)
	return
}

// Create a new `Entity` to be managed by your organization. This `Entity`
// represents an individual customer of your organization.
func (r *V2EntityService) New(ctx context.Context, body V2EntityNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a specific customer `Entity` of your organization.
func (r *V2EntityService) Update(ctx context.Context, entityID string, body V2EntityUpdateParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a list of direct `Entities` your organization manages. These `Entities`
// represent individual customers of your organization.
func (r *V2EntityService) List(ctx context.Context, query V2EntityListParams, opts ...option.RequestOption) (res *V2EntityListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a specific customer `Entity` of your organization by their ID.
func (r *V2EntityService) GetByID(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v2/entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the current authenticated `Entity`, which represents your organization.
func (r *V2EntityService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *Entity, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v2/entities/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Information about an `Entity`, which can be either an individual or an
// organization.
type Entity struct {
	// Unique ID of the `Entity`.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of `Entity`. `ORGANIZATION` for Dinari Partners and `INDIVIDUAL` for their
	// individual customers.
	//
	// Any of "INDIVIDUAL", "ORGANIZATION".
	EntityType EntityEntityType `json:"entity_type" api:"required"`
	// Indicates if `Entity` completed KYC.
	IsKYCComplete bool `json:"is_kyc_complete" api:"required"`
	// Name of `Entity`.
	Name string `json:"name" api:"nullable"`
	// Nationality or home country of the `Entity`.
	Nationality string `json:"nationality" api:"nullable"`
	// Case sensitive unique reference ID that you can set for the `Entity`. We
	// recommend setting this to the unique ID of the `Entity` in your system.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		EntityType    respjson.Field
		IsKYCComplete respjson.Field
		Name          respjson.Field
		Nationality   respjson.Field
		ReferenceID   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Entity) RawJSON() string { return r.JSON.raw }
func (r *Entity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of `Entity`. `ORGANIZATION` for Dinari Partners and `INDIVIDUAL` for their
// individual customers.
type EntityEntityType string

const (
	EntityEntityTypeIndividual   EntityEntityType = "INDIVIDUAL"
	EntityEntityTypeOrganization EntityEntityType = "ORGANIZATION"
)

type V2EntityListResponse struct {
	// List of Entity
	Data []V2EntityListResponseData `json:"data" api:"required"`
	// Pagination metadata
	PaginationMetadata PaginationMetadata `json:"pagination_metadata" api:"required"`
	// Version
	//
	// Any of "PaginatedEntityResponse:v1".
	Sv V2EntityListResponse_Sv `json:"_sv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data               respjson.Field
		PaginationMetadata respjson.Field
		Sv                 respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V2EntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about an `Entity`, which can be either an individual or an
// organization.
type V2EntityListResponseData struct {
	// Unique ID of the `Entity`.
	ID string `json:"id" api:"required" format:"uuid"`
	// Type of `Entity`. `ORGANIZATION` for Dinari Partners and `INDIVIDUAL` for their
	// individual customers.
	//
	// Any of "INDIVIDUAL", "ORGANIZATION".
	EntityType string `json:"entity_type" api:"required"`
	// Indicates if `Entity` completed KYC.
	IsKYCComplete bool `json:"is_kyc_complete" api:"required"`
	// Name of `Entity`.
	Name string `json:"name" api:"nullable"`
	// Nationality or home country of the `Entity`.
	Nationality string `json:"nationality" api:"nullable"`
	// Case sensitive unique reference ID that you can set for the `Entity`. We
	// recommend setting this to the unique ID of the `Entity` in your system.
	ReferenceID string `json:"reference_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		EntityType    respjson.Field
		IsKYCComplete respjson.Field
		Name          respjson.Field
		Nationality   respjson.Field
		ReferenceID   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V2EntityListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Version
type V2EntityListResponse_Sv string

const (
	V2EntityListResponse_SvPaginatedEntityResponseV1 V2EntityListResponse_Sv = "PaginatedEntityResponse:v1"
)

type V2EntityNewParams struct {
	// Name of the `Entity`.
	Name string `json:"name" api:"required"`
	// Case sensitive unique reference ID for the `Entity`. We recommend setting this
	// to the unique ID of the `Entity` in your system.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	paramObj
}

func (r V2EntityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityUpdateParams struct {
	// Case sensitive unique reference ID for the `Entity`. We recommend setting this
	// to the unique ID of the `Entity` in your system.
	ReferenceID param.Opt[string] `json:"reference_id,omitzero"`
	paramObj
}

func (r V2EntityUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityListParams struct {
	// Cursor for next page
	Next param.Opt[string] `query:"next,omitzero" json:"-"`
	// Cursor for previous page
	Previous param.Opt[string] `query:"previous,omitzero" json:"-"`
	// Case sensitive unique reference ID for the `Entity`.
	ReferenceID param.Opt[string] `query:"reference_id,omitzero" json:"-"`
	// Number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order
	//
	// Any of "asc", "desc".
	Order V2EntityListParamsOrder `query:"order,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V2EntityListParams]'s query parameters as `url.Values`.
func (r V2EntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order
type V2EntityListParamsOrder string

const (
	V2EntityListParamsOrderAsc  V2EntityListParamsOrder = "asc"
	V2EntityListParamsOrderDesc V2EntityListParamsOrder = "desc"
)
