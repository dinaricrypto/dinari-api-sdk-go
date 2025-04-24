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
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2EntityKYCService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2EntityKYCService] method instead.
type APIV2EntityKYCService struct {
	Options []option.RequestOption
}

// NewAPIV2EntityKYCService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAPIV2EntityKYCService(opts ...option.RequestOption) (r APIV2EntityKYCService) {
	r = APIV2EntityKYCService{}
	r.Options = opts
	return
}

// Retrieves KYC data of the entity.
func (r *APIV2EntityKYCService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *KYCInfo, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets an iframe URL for managed (self-service) KYC.
func (r *APIV2EntityKYCService) GetURL(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Apiv2EntityKYCGetURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc/url", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Submits KYC data manually (for Partner KYC-enabled entities).
func (r *APIV2EntityKYCService) Submit(ctx context.Context, entityID string, body APIV2EntityKYCSubmitParams, opts ...option.RequestOption) (res *KYCInfo, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Uploads KYC-related documentation (for Partner KYC-enabled entities).
func (r *APIV2EntityKYCService) UploadDocument(ctx context.Context, entityID string, kycID string, body APIV2EntityKYCUploadDocumentParams, opts ...option.RequestOption) (res *Apiv2EntityKYCUploadDocumentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	if kycID == "" {
		err = errors.New("missing required kyc_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc/%s/document", entityID, kycID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Object consisting of KYC data for an entity
type KYCData struct {
	// ISO 3166-1 alpha 2 country code of citizenship or the country the organization
	// is based out of.
	CountryCode string `json:"country_code,required"`
	// Last name of the person
	LastName string `json:"last_name,required"`
	// City of address. Not all international addresses use this attribute.
	AddressCity string `json:"address_city"`
	// ZIP or postal code of residence address. Not all international addresses use
	// this attribute.
	AddressPostalCode string `json:"address_postal_code"`
	// Street name of address.
	AddressStreet1 string `json:"address_street_1"`
	// Extension of address, usually apartment or suite number.
	AddressStreet2 string `json:"address_street_2"`
	// State or subdivision of address. In the US, this should be the unabbreviated
	// name. Not all international addresses use this attribute.
	AddressSubdivision string `json:"address_subdivision"`
	// Birth date of the individual
	BirthDate time.Time `json:"birth_date" format:"date"`
	// Email address
	Email string `json:"email,nullable"`
	// First name of the person, or name of the organization
	FirstName string `json:"first_name,nullable"`
	// Middle name of the user
	MiddleName string `json:"middle_name,nullable"`
	// ID number of the official tax document of the country the entity belongs to
	TaxIDNumber string `json:"tax_id_number"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		CountryCode        resp.Field
		LastName           resp.Field
		AddressCity        resp.Field
		AddressPostalCode  resp.Field
		AddressStreet1     resp.Field
		AddressStreet2     resp.Field
		AddressSubdivision resp.Field
		BirthDate          resp.Field
		Email              resp.Field
		FirstName          resp.Field
		MiddleName         resp.Field
		TaxIDNumber        resp.Field
		ExtraFields        map[string]resp.Field
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
// KYCDataParam.IsOverridden()
func (r KYCData) ToParam() KYCDataParam {
	return param.OverrideObj[KYCDataParam](r.RawJSON())
}

// Object consisting of KYC data for an entity
//
// The properties CountryCode, LastName are required.
type KYCDataParam struct {
	// ISO 3166-1 alpha 2 country code of citizenship or the country the organization
	// is based out of.
	CountryCode string `json:"country_code,required"`
	// Last name of the person
	LastName string `json:"last_name,required"`
	// Email address
	Email param.Opt[string] `json:"email,omitzero"`
	// First name of the person, or name of the organization
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	// Middle name of the user
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// City of address. Not all international addresses use this attribute.
	AddressCity param.Opt[string] `json:"address_city,omitzero"`
	// ZIP or postal code of residence address. Not all international addresses use
	// this attribute.
	AddressPostalCode param.Opt[string] `json:"address_postal_code,omitzero"`
	// Street name of address.
	AddressStreet1 param.Opt[string] `json:"address_street_1,omitzero"`
	// Extension of address, usually apartment or suite number.
	AddressStreet2 param.Opt[string] `json:"address_street_2,omitzero"`
	// State or subdivision of address. In the US, this should be the unabbreviated
	// name. Not all international addresses use this attribute.
	AddressSubdivision param.Opt[string] `json:"address_subdivision,omitzero"`
	// Birth date of the individual
	BirthDate param.Opt[time.Time] `json:"birth_date,omitzero" format:"date"`
	// ID number of the official tax document of the country the entity belongs to
	TaxIDNumber param.Opt[string] `json:"tax_id_number,omitzero"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f KYCDataParam) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }
func (r KYCDataParam) MarshalJSON() (data []byte, err error) {
	type shadow KYCDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type KYCDocumentType string

const (
	KYCDocumentTypeGovernmentID KYCDocumentType = "GOVERNMENT_ID"
	KYCDocumentTypeSelfie       KYCDocumentType = "SELFIE"
	KYCDocumentTypeResidency    KYCDocumentType = "RESIDENCY"
	KYCDocumentTypeUnknown      KYCDocumentType = "UNKNOWN"
)

// KYC information for an entity
type KYCInfo struct {
	// Unique identifier for the KYC check
	ID string `json:"id,required" format:"bigint"`
	// KYC status
	//
	// Any of "PASS", "FAIL", "PENDING", "INCOMPLETE".
	Status KYCInfoStatus `json:"status,required"`
	// Timestamp when the KYC was last checked
	CheckedDt time.Time `json:"checked_dt" format:"date-time"`
	// Object consisting of KYC data for an entity
	Data KYCData `json:"data"`
	// Name of the KYC provider that provided the KYC check
	ProviderName string `json:"provider_name"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		Status       resp.Field
		CheckedDt    resp.Field
		Data         resp.Field
		ProviderName resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCInfo) RawJSON() string { return r.JSON.raw }
func (r *KYCInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC status
type KYCInfoStatus string

const (
	KYCInfoStatusPass       KYCInfoStatus = "PASS"
	KYCInfoStatusFail       KYCInfoStatus = "FAIL"
	KYCInfoStatusPending    KYCInfoStatus = "PENDING"
	KYCInfoStatusIncomplete KYCInfoStatus = "INCOMPLETE"
)

// URL for a managed KYC flow for the entity that can be shown in an iframe
type Apiv2EntityKYCGetURLResponse struct {
	// URL of a managed KYC flow interface for the entity. This URL is unique per KYC
	// attempt.
	EmbedURL string `json:"embed_url,required"`
	// Timestamp at which the KYC request will be expired
	ExpirationDt time.Time `json:"expiration_dt,required" format:"date-time"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		EmbedURL     resp.Field
		ExpirationDt resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2EntityKYCGetURLResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2EntityKYCGetURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Document associated with KYC for an entity
type Apiv2EntityKYCUploadDocumentResponse struct {
	// ID of the document
	ID string `json:"id,required" format:"uuid"`
	// Type of the document
	//
	// Any of "GOVERNMENT_ID", "SELFIE", "RESIDENCY", "UNKNOWN".
	DocumentType KYCDocumentType `json:"document_type,required"`
	// Filename of the document
	Filename string `json:"filename,required"`
	// URL to access the document. Expires in 1 hour
	URL string `json:"url,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		ID           resp.Field
		DocumentType resp.Field
		Filename     resp.Field
		URL          resp.Field
		ExtraFields  map[string]resp.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2EntityKYCUploadDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2EntityKYCUploadDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2EntityKYCSubmitParams struct {
	// Object consisting of KYC data for an entity
	Data KYCDataParam `json:"data,omitzero,required"`
	// Name of the KYC provider that provided the KYC information
	ProviderName string `json:"provider_name,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2EntityKYCSubmitParams) IsPresent() bool { return !param.IsOmitted(f) && !f.IsNull() }

func (r APIV2EntityKYCSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2EntityKYCSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type APIV2EntityKYCUploadDocumentParams struct {
	// Type of the document to be uploaded
	//
	// Any of "GOVERNMENT_ID", "SELFIE", "RESIDENCY", "UNKNOWN".
	DocumentType KYCDocumentType `json:"document_type,omitzero,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2EntityKYCUploadDocumentParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2EntityKYCUploadDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2EntityKYCUploadDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
