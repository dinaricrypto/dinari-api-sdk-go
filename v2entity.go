// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/respjson"
)

// V2EntityService contains methods and other services that help with interacting
// with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityService] method instead.
type V2EntityService struct {
	Options  []option.RequestOption
	Accounts V2EntityAccountService
	KYC      V2EntityKYCService
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
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of all direct `Entities` your organization manages. These `Entities`
// represent individual customers of your organization.
func (r *V2EntityService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a specific customer `Entity` of your organization by their ID.
func (r *V2EntityService) GetByID(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the current authenticated `Entity`, which represents your organization.
func (r *V2EntityService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about an `Entity`, which can be either an individual or an
// organization.
type Entity struct {
	// Unique ID of the `Entity`.
	ID string `json:"id,required" format:"uuid"`
	// Type of `Entity`. `ORGANIZATION` for Dinari Partners and `INDIVIDUAL` for their
	// individual customers.
	//
	// Any of "INDIVIDUAL", "ORGANIZATION".
	EntityType EntityEntityType `json:"entity_type,required"`
	// Indicates if `Entity` completed KYC.
	IsKYCComplete bool `json:"is_kyc_complete,required"`
	// Name of `Entity`.
	Name string `json:"name"`
	// Nationality or home country of the `Entity`.
	Nationality string `json:"nationality"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		EntityType    respjson.Field
		IsKYCComplete respjson.Field
		Name          respjson.Field
		Nationality   respjson.Field
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

type V2EntityNewParams struct {
	// Name of the `Entity`.
	Name string `json:"name,required"`
	paramObj
}

func (r V2EntityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
