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
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2EntityService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2EntityService] method instead.
type APIV2EntityService struct {
	Options  []option.RequestOption
	Accounts APIV2EntityAccountService
	KYC      APIV2EntityKYCService
}

// NewAPIV2EntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV2EntityService(opts ...option.RequestOption) (r APIV2EntityService) {
	r = APIV2EntityService{}
	r.Options = opts
	r.Accounts = NewAPIV2EntityAccountService(opts...)
	r.KYC = NewAPIV2EntityKYCService(opts...)
	return
}

// Creates a new Entity to be managed by your organization. The Entity represents
// an individual customer of your organization.
func (r *APIV2EntityService) New(ctx context.Context, body APIV2EntityNewParams, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a specific customer Entity of your organization by their ID.
func (r *APIV2EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all direct Entities your organization manages. An Entity
// represents an individual customer of your organization.
func (r *APIV2EntityService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the current authenticated Entity.
func (r *APIV2EntityService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/v2/entities/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about an entity, which can be either an individual or an
// organization.
type Entity struct {
	// Unique identifier for the entity
	ID string `json:"id,required" format:"uuid"`
	// Type of entity
	//
	// Any of "INDIVIDUAL", "ORGANIZATION".
	EntityType EntityEntityType `json:"entity_type,required"`
	// Indicates if Entity completed KYC
	IsKYCComplete bool `json:"is_kyc_complete,required"`
	// Name of Entity
	Name string `json:"name"`
	// Nationality of the entity
	Nationality string `json:"nationality"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID            resp.Field
		EntityType    resp.Field
		IsKYCComplete resp.Field
		Name          resp.Field
		Nationality   resp.Field
		ExtraFields   map[string]resp.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Entity) RawJSON() string { return r.JSON.raw }
func (r *Entity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of entity
type EntityEntityType string

const (
	EntityEntityTypeIndividual   EntityEntityType = "INDIVIDUAL"
	EntityEntityTypeOrganization EntityEntityType = "ORGANIZATION"
)

type APIV2EntityNewParams struct {
	// Name of the entity
	Name string `json:"name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2EntityNewParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r APIV2EntityNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2EntityNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
