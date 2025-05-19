// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/respjson"
)

// V2EntityKYCService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityKYCService] method instead.
type V2EntityKYCService struct {
	Options  []option.RequestOption
	Document V2EntityKYCDocumentService
}

// NewV2EntityKYCService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2EntityKYCService(opts ...option.RequestOption) (r V2EntityKYCService) {
	r = V2EntityKYCService{}
	r.Options = opts
	r.Document = NewV2EntityKYCDocumentService(opts...)
	return
}

// Get most recent KYC data of the `Entity`.
//
// If there are any completed KYC checks, data from the most recent one will be
// returned. If there are no completed KYC checks, the most recent KYC check
// information, regardless of status, will be returned.
func (r *V2EntityKYCService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *KYCInfo, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create a Dinari-managed KYC Check and get a URL for your end customer to
// interact with.
//
// The URL points to a web-based KYC interface that can be presented to the end
// customer for KYC verification. Once the customer completes this KYC flow, the
// KYC check will be created and available in the KYC API.
func (r *V2EntityKYCService) NewManagedCheck(ctx context.Context, entityID string, opts ...option.RequestOption) (res *V2EntityKYCNewManagedCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc/url", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Submit KYC data directly, for partners that are provisioned to provide their own
// KYC data.
//
// This feature is available for everyone in sandbox mode, and for specifically
// provisioned partners in production.
func (r *V2EntityKYCService) Submit(ctx context.Context, entityID string, body V2EntityKYCSubmitParams, opts ...option.RequestOption) (res *KYCInfo, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// KYC data for an `Entity`.
type KYCData struct {
	// Country of citizenship or home country of the organization. ISO 3166-1 alpha 2
	// country code.
	CountryCode string `json:"country_code,required"`
	// Last name of the person.
	LastName string `json:"last_name,required"`
	// City of address. Not all international addresses use this attribute.
	AddressCity string `json:"address_city"`
	// Postal code of residence address. Not all international addresses use this
	// attribute.
	AddressPostalCode string `json:"address_postal_code"`
	// Street address of address.
	AddressStreet1 string `json:"address_street_1"`
	// Extension of address, usually apartment or suite number.
	AddressStreet2 string `json:"address_street_2"`
	// State or subdivision of address. In the US, this should be the unabbreviated
	// name of the state. Not all international addresses use this attribute.
	AddressSubdivision string `json:"address_subdivision"`
	// Birth date of the individual. In ISO 8601 format, YYYY-MM-DD.
	BirthDate time.Time `json:"birth_date" format:"date"`
	// Email address.
	Email string `json:"email"`
	// First name of the person.
	FirstName string `json:"first_name"`
	// Middle name of the user
	MiddleName string `json:"middle_name"`
	// ID number of the official tax document of the country the entity belongs to.
	TaxIDNumber string `json:"tax_id_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryCode        respjson.Field
		LastName           respjson.Field
		AddressCity        respjson.Field
		AddressPostalCode  respjson.Field
		AddressStreet1     respjson.Field
		AddressStreet2     respjson.Field
		AddressSubdivision respjson.Field
		BirthDate          respjson.Field
		Email              respjson.Field
		FirstName          respjson.Field
		MiddleName         respjson.Field
		TaxIDNumber        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCData) RawJSON() string { return r.JSON.raw }
func (r *KYCData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this KYCData to a KYCDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// KYCDataParam.Overrides()
func (r KYCData) ToParam() KYCDataParam {
	return param.Override[KYCDataParam](r.RawJSON())
}

// KYC data for an `Entity`.
//
// The properties CountryCode, LastName are required.
type KYCDataParam struct {
	// Country of citizenship or home country of the organization. ISO 3166-1 alpha 2
	// country code.
	CountryCode string `json:"country_code,required"`
	// Last name of the person.
	LastName string `json:"last_name,required"`
	// City of address. Not all international addresses use this attribute.
	AddressCity param.Opt[string] `json:"address_city,omitzero"`
	// Postal code of residence address. Not all international addresses use this
	// attribute.
	AddressPostalCode param.Opt[string] `json:"address_postal_code,omitzero"`
	// Street address of address.
	AddressStreet1 param.Opt[string] `json:"address_street_1,omitzero"`
	// Extension of address, usually apartment or suite number.
	AddressStreet2 param.Opt[string] `json:"address_street_2,omitzero"`
	// State or subdivision of address. In the US, this should be the unabbreviated
	// name of the state. Not all international addresses use this attribute.
	AddressSubdivision param.Opt[string] `json:"address_subdivision,omitzero"`
	// Birth date of the individual. In ISO 8601 format, YYYY-MM-DD.
	BirthDate param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	// Email address.
	Email param.Opt[string] `json:"email,omitzero"`
	// First name of the person.
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Middle name of the user
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// ID number of the official tax document of the country the entity belongs to.
	TaxIDNumber param.Opt[string] `json:"tax_id_number,omitzero"`
	paramObj
}

func (r KYCDataParam) MarshalJSON() (data []byte, err error) {
	type shadow KYCDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KYCDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC information for an `Entity`.
type KYCInfo struct {
	// ID of the KYC check.
	ID string `json:"id,required" format:"uuid"`
	// KYC check status.
	//
	// Any of "PASS", "FAIL", "PENDING", "INCOMPLETE".
	Status KYCInfoStatus `json:"status,required"`
	// Datetime when the KYC was last checked. ISO 8601 timestamp.
	CheckedDt time.Time `json:"checked_dt" format:"date-time"`
	// KYC data for an `Entity`.
	Data KYCData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		CheckedDt   respjson.Field
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCInfo) RawJSON() string { return r.JSON.raw }
func (r *KYCInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC check status.
type KYCInfoStatus string

const (
	KYCInfoStatusPass       KYCInfoStatus = "PASS"
	KYCInfoStatusFail       KYCInfoStatus = "FAIL"
	KYCInfoStatusPending    KYCInfoStatus = "PENDING"
	KYCInfoStatusIncomplete KYCInfoStatus = "INCOMPLETE"
)

// URL for a managed KYC flow for an `Entity`.
type V2EntityKYCNewManagedCheckResponse struct {
	// URL of a managed KYC flow interface for the `Entity`.
	EmbedURL string `json:"embed_url,required"`
	// Datetime at which the KYC request will expired. ISO 8601 timestamp.
	ExpirationDt time.Time `json:"expiration_dt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbedURL     respjson.Field
		ExpirationDt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V2EntityKYCNewManagedCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *V2EntityKYCNewManagedCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2EntityKYCSubmitParams struct {
	// KYC data for an `Entity`.
	Data KYCDataParam `json:"data,omitzero,required"`
	// Name of the KYC provider that provided the KYC information.
	ProviderName string `json:"provider_name,required"`
	paramObj
}

func (r V2EntityKYCSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityKYCSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityKYCSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
