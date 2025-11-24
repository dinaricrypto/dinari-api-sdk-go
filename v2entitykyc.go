// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
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
func (r *V2EntityKYCService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *KYCInfoUnion, err error) {
	opts = slices.Concat(r.Options, opts)
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
	opts = slices.Concat(r.Options, opts)
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
func (r *V2EntityKYCService) Submit(ctx context.Context, entityID string, body V2EntityKYCSubmitParams, opts ...option.RequestOption) (res *KYCInfoUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// KYC data for an `Entity` in the BASELINE jurisdiction.
type BaselineKYCCheckData struct {
	// Country of residence. ISO 3166-1 alpha 2 country code.
	AddressCountryCode string `json:"address_country_code,required"`
	// Country of citizenship or home country of the organization. ISO 3166-1 alpha 2
	// country code.
	CountryCode string `json:"country_code,required"`
	// Last name of the person.
	LastName string `json:"last_name,required"`
	// City of address. Not all international addresses use this attribute.
	AddressCity string `json:"address_city,nullable"`
	// Postal code of residence address. Not all international addresses use this
	// attribute.
	AddressPostalCode string `json:"address_postal_code,nullable"`
	// Street address of address.
	AddressStreet1 string `json:"address_street_1,nullable"`
	// Extension of address, usually apartment or suite number.
	AddressStreet2 string `json:"address_street_2,nullable"`
	// State or subdivision of address. In the US, this should be the unabbreviated
	// name of the state. Not all international addresses use this attribute.
	AddressSubdivision string `json:"address_subdivision,nullable"`
	// Birth date of the individual. In ISO 8601 format, YYYY-MM-DD.
	BirthDate time.Time `json:"birth_date,nullable" format:"date"`
	// Email address.
	Email string `json:"email,nullable"`
	// First name of the person.
	FirstName string `json:"first_name,nullable"`
	// Middle name of the user
	MiddleName string `json:"middle_name,nullable"`
	// ID number of the official tax document of the country the entity belongs to.
	TaxIDNumber string `json:"tax_id_number,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressCountryCode respjson.Field
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
func (r BaselineKYCCheckData) RawJSON() string { return r.JSON.raw }
func (r *BaselineKYCCheckData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BaselineKYCCheckData to a BaselineKYCCheckDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BaselineKYCCheckDataParam.Overrides()
func (r BaselineKYCCheckData) ToParam() BaselineKYCCheckDataParam {
	return param.Override[BaselineKYCCheckDataParam](json.RawMessage(r.RawJSON()))
}

// KYC data for an `Entity` in the BASELINE jurisdiction.
//
// The properties AddressCountryCode, CountryCode, LastName are required.
type BaselineKYCCheckDataParam struct {
	// Country of residence. ISO 3166-1 alpha 2 country code.
	AddressCountryCode string `json:"address_country_code,required"`
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

func (r BaselineKYCCheckDataParam) MarshalJSON() (data []byte, err error) {
	type shadow BaselineKYCCheckDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaselineKYCCheckDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYCInfoUnion contains all possible properties and values from [KYCInfoBaseline],
// [KYCInfoUs].
//
// Use the [KYCInfoUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type KYCInfoUnion struct {
	ID string `json:"id"`
	// This field is from variant [KYCInfoBaseline].
	Status    KYCStatus `json:"status"`
	CheckedDt time.Time `json:"checked_dt"`
	// This field is a union of [BaselineKYCCheckData], [UsKYCCheckData]
	Data KYCInfoUnionData `json:"data"`
	// Any of "BASELINE", "US".
	Jurisdiction string `json:"jurisdiction"`
	JSON         struct {
		ID           respjson.Field
		Status       respjson.Field
		CheckedDt    respjson.Field
		Data         respjson.Field
		Jurisdiction respjson.Field
		raw          string
	} `json:"-"`
}

// anyKYCInfo is implemented by each variant of [KYCInfoUnion] to add type safety
// for the return type of [KYCInfoUnion.AsAny]
type anyKYCInfo interface {
	implKYCInfoUnion()
}

func (KYCInfoBaseline) implKYCInfoUnion() {}
func (KYCInfoUs) implKYCInfoUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := KYCInfoUnion.AsAny().(type) {
//	case dinariapisdkgo.KYCInfoBaseline:
//	case dinariapisdkgo.KYCInfoUs:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u KYCInfoUnion) AsAny() anyKYCInfo {
	switch u.Jurisdiction {
	case "BASELINE":
		return u.AsBaseline()
	case "US":
		return u.AsUs()
	}
	return nil
}

func (u KYCInfoUnion) AsBaseline() (v KYCInfoBaseline) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u KYCInfoUnion) AsUs() (v KYCInfoUs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u KYCInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *KYCInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYCInfoUnionData is an implicit subunion of [KYCInfoUnion]. KYCInfoUnionData
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [KYCInfoUnion].
type KYCInfoUnionData struct {
	// This field is from variant [BaselineKYCCheckData].
	AddressCountryCode string `json:"address_country_code"`
	// This field is from variant [BaselineKYCCheckData].
	CountryCode string `json:"country_code"`
	// This field is from variant [BaselineKYCCheckData].
	LastName string `json:"last_name"`
	// This field is from variant [BaselineKYCCheckData].
	AddressCity string `json:"address_city"`
	// This field is from variant [BaselineKYCCheckData].
	AddressPostalCode string `json:"address_postal_code"`
	// This field is from variant [BaselineKYCCheckData].
	AddressStreet1 string `json:"address_street_1"`
	// This field is from variant [BaselineKYCCheckData].
	AddressStreet2 string `json:"address_street_2"`
	// This field is from variant [BaselineKYCCheckData].
	AddressSubdivision string `json:"address_subdivision"`
	// This field is from variant [BaselineKYCCheckData].
	BirthDate time.Time `json:"birth_date"`
	// This field is from variant [BaselineKYCCheckData].
	Email string `json:"email"`
	// This field is from variant [BaselineKYCCheckData].
	FirstName string `json:"first_name"`
	// This field is from variant [BaselineKYCCheckData].
	MiddleName string `json:"middle_name"`
	// This field is from variant [BaselineKYCCheckData].
	TaxIDNumber string `json:"tax_id_number"`
	// This field is from variant [UsKYCCheckData].
	AlpacaCustomerAgreement UsKYCCheckDataAlpacaCustomerAgreement `json:"alpaca_customer_agreement"`
	// This field is from variant [UsKYCCheckData].
	AmlCheck UsKYCCheckDataAmlCheck `json:"aml_check"`
	// This field is from variant [UsKYCCheckData].
	DataCitation UsKYCCheckDataDataCitation `json:"data_citation"`
	// This field is from variant [UsKYCCheckData].
	Employment UsKYCCheckDataEmployment `json:"employment"`
	// This field is from variant [UsKYCCheckData].
	FinancialProfile UsKYCCheckDataFinancialProfile `json:"financial_profile"`
	// This field is from variant [UsKYCCheckData].
	Identity UsKYCCheckDataIdentity `json:"identity"`
	// This field is from variant [UsKYCCheckData].
	KYCMetadata UsKYCCheckDataKYCMetadata `json:"kyc_metadata"`
	// This field is from variant [UsKYCCheckData].
	RiskDisclosure UsKYCCheckDataRiskDisclosure `json:"risk_disclosure"`
	// This field is from variant [UsKYCCheckData].
	TrustedContact UsKYCCheckDataTrustedContact `json:"trusted_contact"`
	// This field is from variant [UsKYCCheckData].
	UsImmigrationInfo UsKYCCheckDataUsImmigrationInfo `json:"us_immigration_info"`
	JSON              struct {
		AddressCountryCode      respjson.Field
		CountryCode             respjson.Field
		LastName                respjson.Field
		AddressCity             respjson.Field
		AddressPostalCode       respjson.Field
		AddressStreet1          respjson.Field
		AddressStreet2          respjson.Field
		AddressSubdivision      respjson.Field
		BirthDate               respjson.Field
		Email                   respjson.Field
		FirstName               respjson.Field
		MiddleName              respjson.Field
		TaxIDNumber             respjson.Field
		AlpacaCustomerAgreement respjson.Field
		AmlCheck                respjson.Field
		DataCitation            respjson.Field
		Employment              respjson.Field
		FinancialProfile        respjson.Field
		Identity                respjson.Field
		KYCMetadata             respjson.Field
		RiskDisclosure          respjson.Field
		TrustedContact          respjson.Field
		UsImmigrationInfo       respjson.Field
		raw                     string
	} `json:"-"`
}

func (r *KYCInfoUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC information for an `Entity` in the baseline jurisdiction.
type KYCInfoBaseline struct {
	// ID of the KYC check.
	ID string `json:"id,required" format:"uuid"`
	// KYC check status.
	//
	// Any of "PASS", "FAIL", "PENDING", "INCOMPLETE", "NEEDS_REVIEW".
	Status KYCStatus `json:"status,required"`
	// Datetime when the KYC was last checked. ISO 8601 timestamp.
	CheckedDt time.Time `json:"checked_dt,nullable" format:"date-time"`
	// KYC data for an `Entity` in the BASELINE jurisdiction.
	Data BaselineKYCCheckData `json:"data,nullable"`
	// Jurisdiction of the KYC check.
	//
	// Any of "BASELINE".
	Jurisdiction string `json:"jurisdiction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		CheckedDt    respjson.Field
		Data         respjson.Field
		Jurisdiction respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCInfoBaseline) RawJSON() string { return r.JSON.raw }
func (r *KYCInfoBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC information for an `Entity` in the US jurisdiction.
type KYCInfoUs struct {
	// ID of the KYC check.
	ID string `json:"id,required" format:"uuid"`
	// KYC check status.
	//
	// Any of "PASS", "FAIL", "PENDING", "INCOMPLETE", "NEEDS_REVIEW".
	Status KYCStatus `json:"status,required"`
	// Datetime when the KYC was last checked. ISO 8601 timestamp.
	CheckedDt time.Time `json:"checked_dt,nullable" format:"date-time"`
	// KYC data for an `Entity` in the US jurisdiction.
	Data UsKYCCheckData `json:"data,nullable"`
	// Jurisdiction of the KYC check.
	//
	// Any of "US".
	Jurisdiction string `json:"jurisdiction"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Status       respjson.Field
		CheckedDt    respjson.Field
		Data         respjson.Field
		Jurisdiction respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCInfoUs) RawJSON() string { return r.JSON.raw }
func (r *KYCInfoUs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYCStatus string

const (
	KYCStatusPass        KYCStatus = "PASS"
	KYCStatusFail        KYCStatus = "FAIL"
	KYCStatusPending     KYCStatus = "PENDING"
	KYCStatusIncomplete  KYCStatus = "INCOMPLETE"
	KYCStatusNeedsReview KYCStatus = "NEEDS_REVIEW"
)

// KYC data for an `Entity` in the US jurisdiction.
type UsKYCCheckData struct {
	// Information to affirm that the individual has read, agreed to, and signed
	// Alpaca's customer agreement, found here:
	// https://files.alpaca.markets/disclosures/library/AcctAppMarginAndCustAgmt.pdf
	AlpacaCustomerAgreement UsKYCCheckDataAlpacaCustomerAgreement `json:"alpaca_customer_agreement,required"`
	// AML check information for this individual. If any of the checks have a match,
	// provide details about the matches or hits found. The individual will be marked
	// as high risk and be subject to manual review.
	AmlCheck UsKYCCheckDataAmlCheck `json:"aml_check,required"`
	// Data source citations for a KYC check.
	DataCitation UsKYCCheckDataDataCitation `json:"data_citation,required"`
	// Employment information for the individual
	Employment UsKYCCheckDataEmployment `json:"employment,required"`
	// Financial profile information for the individual <br/><br/> Examples of liquid
	// net worth ranges: <br/> - $0 - $20,000 <br/> - $20,000 - $50,000 <br/> -
	// $50,000 - $100,000 <br/> - $100,000 - $500,000 <br/> - $500,000 - $1,000,000
	FinancialProfile UsKYCCheckDataFinancialProfile `json:"financial_profile,required"`
	// Identity information for the individual
	Identity UsKYCCheckDataIdentity `json:"identity,required"`
	// Metadata about the KYC check.
	KYCMetadata UsKYCCheckDataKYCMetadata `json:"kyc_metadata,required"`
	// Risk information about the individual <br/><br/> Fields denote if the account
	// owner falls under each category defined by FINRA rules. If any of the answers is
	// true (yes), additional verifications may be required before US account approval.
	RiskDisclosure UsKYCCheckDataRiskDisclosure `json:"risk_disclosure,required"`
	// Information for a trusted contact person for the individual. More information:
	// <br/> -
	// <a href="https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-trusted-contact" target="_blank" rel="noopener noreferrer">Investor.gov -
	// Trusted Contact</a> <br/> -
	// <a href="https://www.finra.org/investors/insights/trusted-contact" target="_blank" rel="noopener noreferrer">FINRA -
	// Trusted Contact</a>
	TrustedContact UsKYCCheckDataTrustedContact `json:"trusted_contact,required"`
	// US immigration information for this individual. Required if the individual is
	// not a US citizen.
	UsImmigrationInfo UsKYCCheckDataUsImmigrationInfo `json:"us_immigration_info,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlpacaCustomerAgreement respjson.Field
		AmlCheck                respjson.Field
		DataCitation            respjson.Field
		Employment              respjson.Field
		FinancialProfile        respjson.Field
		Identity                respjson.Field
		KYCMetadata             respjson.Field
		RiskDisclosure          respjson.Field
		TrustedContact          respjson.Field
		UsImmigrationInfo       respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckData) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UsKYCCheckData to a UsKYCCheckDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UsKYCCheckDataParam.Overrides()
func (r UsKYCCheckData) ToParam() UsKYCCheckDataParam {
	return param.Override[UsKYCCheckDataParam](json.RawMessage(r.RawJSON()))
}

// Information to affirm that the individual has read, agreed to, and signed
// Alpaca's customer agreement, found here:
// https://files.alpaca.markets/disclosures/library/AcctAppMarginAndCustAgmt.pdf
type UsKYCCheckDataAlpacaCustomerAgreement struct {
	// The IP address from where the individual signed the agreement.
	IPAddress string `json:"ip_address,required" format:"ip"`
	// The timestamp the agreement was signed.
	SignedAt time.Time `json:"signed_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPAddress   respjson.Field
		SignedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataAlpacaCustomerAgreement) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataAlpacaCustomerAgreement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AML check information for this individual. If any of the checks have a match,
// provide details about the matches or hits found. The individual will be marked
// as high risk and be subject to manual review.
type UsKYCCheckDataAmlCheck struct {
	// Datetime that this AML check was created.
	CheckCreatedAt time.Time `json:"check_created_at,required" format:"date-time"`
	// Whether there was a match in the adverse media check.
	IsAdverseMediaMatch bool `json:"is_adverse_media_match,required"`
	// Whether there was a match in the monitored lists check.
	IsMonitoredListsMatch bool `json:"is_monitored_lists_match,required"`
	// Whether there was a match in the politically exposed person (PEP) check.
	IsPoliticallyExposedPersonMatch bool `json:"is_politically_exposed_person_match,required"`
	// Whether there was a match in the sanctions check.
	IsSanctionsMatch bool `json:"is_sanctions_match,required"`
	// If any of the checks have a match, provide details about the matches or hits
	// found.
	Records []string `json:"records,required"`
	// Your unique identifier for the AML check.
	RefID string `json:"ref_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckCreatedAt                  respjson.Field
		IsAdverseMediaMatch             respjson.Field
		IsMonitoredListsMatch           respjson.Field
		IsPoliticallyExposedPersonMatch respjson.Field
		IsSanctionsMatch                respjson.Field
		Records                         respjson.Field
		RefID                           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataAmlCheck) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataAmlCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data source citations for a KYC check.
type UsKYCCheckDataDataCitation struct {
	// List of sources for address verification
	AddressSources []string `json:"address_sources,required"`
	// List of sources for date of birth verification
	DateOfBirthSources []string `json:"date_of_birth_sources,required"`
	// List of sources for tax ID verification
	TaxIDSources []string `json:"tax_id_sources,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressSources     respjson.Field
		DateOfBirthSources respjson.Field
		TaxIDSources       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataDataCitation) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataDataCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employment information for the individual
type UsKYCCheckDataEmployment struct {
	// One of the following: employed, unemployed, retired, or student.
	//
	// Any of "UNEMPLOYED", "EMPLOYED", "STUDENT", "RETIRED".
	EmploymentStatus string `json:"employment_status,required"`
	// The employer's address if the user is employed.
	EmployerAddress string `json:"employer_address,nullable"`
	// The name of the employer if the user is employed.
	EmployerName string `json:"employer_name,nullable"`
	// The user's position if they are employed.
	EmploymentPosition string `json:"employment_position,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmploymentStatus   respjson.Field
		EmployerAddress    respjson.Field
		EmployerName       respjson.Field
		EmploymentPosition respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataEmployment) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataEmployment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Financial profile information for the individual <br/><br/> Examples of liquid
// net worth ranges: <br/> - $0 - $20,000 <br/> - $20,000 - $50,000 <br/> -
// $50,000 - $100,000 <br/> - $100,000 - $500,000 <br/> - $500,000 - $1,000,000
type UsKYCCheckDataFinancialProfile struct {
	// One or more of the following: employment_income, investments, inheritance,
	// business_income, savings, family.
	//
	// Any of "EMPLOYMENT_INCOME", "INVESTMENTS", "INHERITANCE", "BUSINESS_INCOME",
	// "SAVINGS", "FAMILY".
	FundingSources []string `json:"funding_sources,required"`
	// The upper bound of the user's liquid net worth (USD).
	LiquidNetWorthMax int64 `json:"liquid_net_worth_max,required"`
	// The lower bound of the user's liquid net worth (USD). Can be 0 if max is
	// <=$20,000, but otherwise must be within an order of magnitude of the max value.
	LiquidNetWorthMin int64 `json:"liquid_net_worth_min,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FundingSources    respjson.Field
		LiquidNetWorthMax respjson.Field
		LiquidNetWorthMin respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataFinancialProfile) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataFinancialProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identity information for the individual
type UsKYCCheckDataIdentity struct {
	// City of the applicant.
	City string `json:"city,required"`
	// Nationality of the applicant.
	CountryOfCitizenship string `json:"country_of_citizenship,required"`
	// Country of residency of the applicant. Must be 'US'.
	//
	// Any of "US".
	CountryOfTaxResidence string `json:"country_of_tax_residence,required"`
	// Date of birth of the applicant.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// Email address of the applicant.
	EmailAddress string `json:"email_address,required"`
	// The last name (surname) of the user.
	FamilyName string `json:"family_name,required"`
	// The first/given name of the user.
	GivenName string `json:"given_name,required"`
	// Phone number should include the country code, format: “+15555555555”
	PhoneNumber string `json:"phone_number,required"`
	// Postal code of the applicant.
	PostalCode string `json:"postal_code,required"`
	// Street address of the applicant.
	StreetAddress string `json:"street_address,required"`
	// Social Security Number (SSN) or Tax Identification Number (TIN) of the
	// applicant.
	TaxID string `json:"tax_id,required"`
	// The middle name of the user.
	MiddleName string `json:"middle_name,nullable"`
	// State of the applicant. Required if the applicant resides in the US as a
	// 2-letter abbreviation.
	State string `json:"state,nullable"`
	// The specific apartment number if applicable
	Unit string `json:"unit,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City                  respjson.Field
		CountryOfCitizenship  respjson.Field
		CountryOfTaxResidence respjson.Field
		DateOfBirth           respjson.Field
		EmailAddress          respjson.Field
		FamilyName            respjson.Field
		GivenName             respjson.Field
		PhoneNumber           respjson.Field
		PostalCode            respjson.Field
		StreetAddress         respjson.Field
		TaxID                 respjson.Field
		MiddleName            respjson.Field
		State                 respjson.Field
		Unit                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataIdentity) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataIdentity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata about the KYC check.
type UsKYCCheckDataKYCMetadata struct {
	// Completion datetime of KYC check.
	CheckCompletedAt time.Time `json:"check_completed_at,required" format:"date-time"`
	// Start datetime of KYC check.
	CheckInitiatedAt time.Time `json:"check_initiated_at,required" format:"date-time"`
	// IP address of applicant at time of KYC check.
	IPAddress string `json:"ip_address,required" format:"ip"`
	// Your unique identifier for the KYC check.
	RefID string `json:"ref_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CheckCompletedAt respjson.Field
		CheckInitiatedAt respjson.Field
		IPAddress        respjson.Field
		RefID            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataKYCMetadata) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataKYCMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Risk information about the individual <br/><br/> Fields denote if the account
// owner falls under each category defined by FINRA rules. If any of the answers is
// true (yes), additional verifications may be required before US account approval.
type UsKYCCheckDataRiskDisclosure struct {
	// If the individual's immediate family member (sibling, husband/wife, child,
	// parent) is either politically exposed or holds a control position.
	ImmediateFamilyExposed bool `json:"immediate_family_exposed,required"`
	// Whether the individual is affiliated with any exchanges or FINRA.
	IsAffiliatedExchangeOrFinra bool `json:"is_affiliated_exchange_or_finra,required"`
	// Whether the individual holds a controlling position in a publicly traded
	// company, is a member of the board of directors, or has policy making abilities
	// in a publicly traded company.
	IsControlPerson bool `json:"is_control_person,required"`
	// Whether the individual is politically exposed.
	IsPoliticallyExposed bool `json:"is_politically_exposed,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImmediateFamilyExposed      respjson.Field
		IsAffiliatedExchangeOrFinra respjson.Field
		IsControlPerson             respjson.Field
		IsPoliticallyExposed        respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataRiskDisclosure) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataRiskDisclosure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information for a trusted contact person for the individual. More information:
// <br/> -
// <a href="https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-trusted-contact" target="_blank" rel="noopener noreferrer">Investor.gov -
// Trusted Contact</a> <br/> -
// <a href="https://www.finra.org/investors/insights/trusted-contact" target="_blank" rel="noopener noreferrer">FINRA -
// Trusted Contact</a>
type UsKYCCheckDataTrustedContact struct {
	// The family name of the trusted contact
	FamilyName string `json:"family_name,required"`
	// The given name of the trusted contact
	GivenName string `json:"given_name,required"`
	// The email address of the trusted contact. At least one of email_address or
	// phone_number is required.
	EmailAddress string `json:"email_address,nullable"`
	// The phone number of the trusted contact. At least one of email_address or
	// phone_number is required.
	PhoneNumber string `json:"phone_number,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FamilyName   respjson.Field
		GivenName    respjson.Field
		EmailAddress respjson.Field
		PhoneNumber  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataTrustedContact) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataTrustedContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// US immigration information for this individual. Required if the individual is
// not a US citizen.
type UsKYCCheckDataUsImmigrationInfo struct {
	// Country where the individual was born.
	CountryOfBirth string `json:"country_of_birth,required"`
	// Whether the individual is a US permanent resident (green card holder).
	IsPermanentResident bool `json:"is_permanent_resident,required"`
	// Date the individual is scheduled to leave the US. Required for B1 and B2 visas.
	DepartureFromUsDate time.Time `json:"departure_from_us_date,nullable" format:"date"`
	// Expiration date of the visa. Required if visa_type is provided.
	VisaExpirationDate time.Time `json:"visa_expiration_date,nullable" format:"date"`
	// Type of visa the individual holds. Required if not a permanent resident.
	//
	// Any of "B1", "B2", "DACA", "E1", "E2", "E3", "F1", "G4", "H1B", "J1", "L1",
	// "Other", "O1", "TN1".
	VisaType string `json:"visa_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryOfBirth      respjson.Field
		IsPermanentResident respjson.Field
		DepartureFromUsDate respjson.Field
		VisaExpirationDate  respjson.Field
		VisaType            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsKYCCheckDataUsImmigrationInfo) RawJSON() string { return r.JSON.raw }
func (r *UsKYCCheckDataUsImmigrationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// KYC data for an `Entity` in the US jurisdiction.
//
// The properties AlpacaCustomerAgreement, AmlCheck, DataCitation, Employment,
// FinancialProfile, Identity, KYCMetadata, RiskDisclosure, TrustedContact are
// required.
type UsKYCCheckDataParam struct {
	// Information to affirm that the individual has read, agreed to, and signed
	// Alpaca's customer agreement, found here:
	// https://files.alpaca.markets/disclosures/library/AcctAppMarginAndCustAgmt.pdf
	AlpacaCustomerAgreement UsKYCCheckDataAlpacaCustomerAgreementParam `json:"alpaca_customer_agreement,omitzero,required"`
	// AML check information for this individual. If any of the checks have a match,
	// provide details about the matches or hits found. The individual will be marked
	// as high risk and be subject to manual review.
	AmlCheck UsKYCCheckDataAmlCheckParam `json:"aml_check,omitzero,required"`
	// Data source citations for a KYC check.
	DataCitation UsKYCCheckDataDataCitationParam `json:"data_citation,omitzero,required"`
	// Employment information for the individual
	Employment UsKYCCheckDataEmploymentParam `json:"employment,omitzero,required"`
	// Financial profile information for the individual <br/><br/> Examples of liquid
	// net worth ranges: <br/> - $0 - $20,000 <br/> - $20,000 - $50,000 <br/> -
	// $50,000 - $100,000 <br/> - $100,000 - $500,000 <br/> - $500,000 - $1,000,000
	FinancialProfile UsKYCCheckDataFinancialProfileParam `json:"financial_profile,omitzero,required"`
	// Identity information for the individual
	Identity UsKYCCheckDataIdentityParam `json:"identity,omitzero,required"`
	// Metadata about the KYC check.
	KYCMetadata UsKYCCheckDataKYCMetadataParam `json:"kyc_metadata,omitzero,required"`
	// Risk information about the individual <br/><br/> Fields denote if the account
	// owner falls under each category defined by FINRA rules. If any of the answers is
	// true (yes), additional verifications may be required before US account approval.
	RiskDisclosure UsKYCCheckDataRiskDisclosureParam `json:"risk_disclosure,omitzero,required"`
	// Information for a trusted contact person for the individual. More information:
	// <br/> -
	// <a href="https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-trusted-contact" target="_blank" rel="noopener noreferrer">Investor.gov -
	// Trusted Contact</a> <br/> -
	// <a href="https://www.finra.org/investors/insights/trusted-contact" target="_blank" rel="noopener noreferrer">FINRA -
	// Trusted Contact</a>
	TrustedContact UsKYCCheckDataTrustedContactParam `json:"trusted_contact,omitzero,required"`
	// US immigration information for this individual. Required if the individual is
	// not a US citizen.
	UsImmigrationInfo UsKYCCheckDataUsImmigrationInfoParam `json:"us_immigration_info,omitzero"`
	paramObj
}

func (r UsKYCCheckDataParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information to affirm that the individual has read, agreed to, and signed
// Alpaca's customer agreement, found here:
// https://files.alpaca.markets/disclosures/library/AcctAppMarginAndCustAgmt.pdf
//
// The properties IPAddress, SignedAt are required.
type UsKYCCheckDataAlpacaCustomerAgreementParam struct {
	// The IP address from where the individual signed the agreement.
	IPAddress string `json:"ip_address,required" format:"ip"`
	// The timestamp the agreement was signed.
	SignedAt time.Time `json:"signed_at,required" format:"date-time"`
	paramObj
}

func (r UsKYCCheckDataAlpacaCustomerAgreementParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataAlpacaCustomerAgreementParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataAlpacaCustomerAgreementParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AML check information for this individual. If any of the checks have a match,
// provide details about the matches or hits found. The individual will be marked
// as high risk and be subject to manual review.
//
// The properties CheckCreatedAt, IsAdverseMediaMatch, IsMonitoredListsMatch,
// IsPoliticallyExposedPersonMatch, IsSanctionsMatch, Records, RefID are required.
type UsKYCCheckDataAmlCheckParam struct {
	// Datetime that this AML check was created.
	CheckCreatedAt time.Time `json:"check_created_at,required" format:"date-time"`
	// Whether there was a match in the adverse media check.
	IsAdverseMediaMatch bool `json:"is_adverse_media_match,required"`
	// Whether there was a match in the monitored lists check.
	IsMonitoredListsMatch bool `json:"is_monitored_lists_match,required"`
	// Whether there was a match in the politically exposed person (PEP) check.
	IsPoliticallyExposedPersonMatch bool `json:"is_politically_exposed_person_match,required"`
	// Whether there was a match in the sanctions check.
	IsSanctionsMatch bool `json:"is_sanctions_match,required"`
	// If any of the checks have a match, provide details about the matches or hits
	// found.
	Records []string `json:"records,omitzero,required"`
	// Your unique identifier for the AML check.
	RefID string `json:"ref_id,required"`
	paramObj
}

func (r UsKYCCheckDataAmlCheckParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataAmlCheckParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataAmlCheckParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data source citations for a KYC check.
//
// The properties AddressSources, DateOfBirthSources, TaxIDSources are required.
type UsKYCCheckDataDataCitationParam struct {
	// List of sources for address verification
	AddressSources []string `json:"address_sources,omitzero,required"`
	// List of sources for date of birth verification
	DateOfBirthSources []string `json:"date_of_birth_sources,omitzero,required"`
	// List of sources for tax ID verification
	TaxIDSources []string `json:"tax_id_sources,omitzero,required"`
	paramObj
}

func (r UsKYCCheckDataDataCitationParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataDataCitationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataDataCitationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Employment information for the individual
//
// The property EmploymentStatus is required.
type UsKYCCheckDataEmploymentParam struct {
	// One of the following: employed, unemployed, retired, or student.
	//
	// Any of "UNEMPLOYED", "EMPLOYED", "STUDENT", "RETIRED".
	EmploymentStatus string `json:"employment_status,omitzero,required"`
	// The employer's address if the user is employed.
	EmployerAddress param.Opt[string] `json:"employer_address,omitzero"`
	// The name of the employer if the user is employed.
	EmployerName param.Opt[string] `json:"employer_name,omitzero"`
	// The user's position if they are employed.
	EmploymentPosition param.Opt[string] `json:"employment_position,omitzero"`
	paramObj
}

func (r UsKYCCheckDataEmploymentParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataEmploymentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataEmploymentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsKYCCheckDataEmploymentParam](
		"employment_status", "UNEMPLOYED", "EMPLOYED", "STUDENT", "RETIRED",
	)
}

// Financial profile information for the individual <br/><br/> Examples of liquid
// net worth ranges: <br/> - $0 - $20,000 <br/> - $20,000 - $50,000 <br/> -
// $50,000 - $100,000 <br/> - $100,000 - $500,000 <br/> - $500,000 - $1,000,000
//
// The properties FundingSources, LiquidNetWorthMax, LiquidNetWorthMin are
// required.
type UsKYCCheckDataFinancialProfileParam struct {
	// One or more of the following: employment_income, investments, inheritance,
	// business_income, savings, family.
	//
	// Any of "EMPLOYMENT_INCOME", "INVESTMENTS", "INHERITANCE", "BUSINESS_INCOME",
	// "SAVINGS", "FAMILY".
	FundingSources []string `json:"funding_sources,omitzero,required"`
	// The upper bound of the user's liquid net worth (USD).
	LiquidNetWorthMax int64 `json:"liquid_net_worth_max,required"`
	// The lower bound of the user's liquid net worth (USD). Can be 0 if max is
	// <=$20,000, but otherwise must be within an order of magnitude of the max value.
	LiquidNetWorthMin int64 `json:"liquid_net_worth_min,required"`
	paramObj
}

func (r UsKYCCheckDataFinancialProfileParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataFinancialProfileParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataFinancialProfileParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identity information for the individual
//
// The properties City, CountryOfCitizenship, CountryOfTaxResidence, DateOfBirth,
// EmailAddress, FamilyName, GivenName, PhoneNumber, PostalCode, StreetAddress,
// TaxID are required.
type UsKYCCheckDataIdentityParam struct {
	// City of the applicant.
	City string `json:"city,required"`
	// Nationality of the applicant.
	CountryOfCitizenship string `json:"country_of_citizenship,required"`
	// Country of residency of the applicant. Must be 'US'.
	//
	// Any of "US".
	CountryOfTaxResidence string `json:"country_of_tax_residence,omitzero,required"`
	// Date of birth of the applicant.
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date"`
	// Email address of the applicant.
	EmailAddress string `json:"email_address,required"`
	// The last name (surname) of the user.
	FamilyName string `json:"family_name,required"`
	// The first/given name of the user.
	GivenName string `json:"given_name,required"`
	// Phone number should include the country code, format: “+15555555555”
	PhoneNumber string `json:"phone_number,required"`
	// Postal code of the applicant.
	PostalCode string `json:"postal_code,required"`
	// Street address of the applicant.
	StreetAddress string `json:"street_address,required"`
	// Social Security Number (SSN) or Tax Identification Number (TIN) of the
	// applicant.
	TaxID string `json:"tax_id,required"`
	// The middle name of the user.
	MiddleName param.Opt[string] `json:"middle_name,omitzero"`
	// State of the applicant. Required if the applicant resides in the US as a
	// 2-letter abbreviation.
	State param.Opt[string] `json:"state,omitzero"`
	// The specific apartment number if applicable
	Unit param.Opt[string] `json:"unit,omitzero"`
	paramObj
}

func (r UsKYCCheckDataIdentityParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataIdentityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataIdentityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsKYCCheckDataIdentityParam](
		"country_of_tax_residence", "US",
	)
}

// Metadata about the KYC check.
//
// The properties CheckCompletedAt, CheckInitiatedAt, IPAddress, RefID are
// required.
type UsKYCCheckDataKYCMetadataParam struct {
	// Completion datetime of KYC check.
	CheckCompletedAt time.Time `json:"check_completed_at,required" format:"date-time"`
	// Start datetime of KYC check.
	CheckInitiatedAt time.Time `json:"check_initiated_at,required" format:"date-time"`
	// IP address of applicant at time of KYC check.
	IPAddress string `json:"ip_address,required" format:"ip"`
	// Your unique identifier for the KYC check.
	RefID string `json:"ref_id,required"`
	paramObj
}

func (r UsKYCCheckDataKYCMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataKYCMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataKYCMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Risk information about the individual <br/><br/> Fields denote if the account
// owner falls under each category defined by FINRA rules. If any of the answers is
// true (yes), additional verifications may be required before US account approval.
//
// The properties ImmediateFamilyExposed, IsAffiliatedExchangeOrFinra,
// IsControlPerson, IsPoliticallyExposed are required.
type UsKYCCheckDataRiskDisclosureParam struct {
	// If the individual's immediate family member (sibling, husband/wife, child,
	// parent) is either politically exposed or holds a control position.
	ImmediateFamilyExposed bool `json:"immediate_family_exposed,required"`
	// Whether the individual is affiliated with any exchanges or FINRA.
	IsAffiliatedExchangeOrFinra bool `json:"is_affiliated_exchange_or_finra,required"`
	// Whether the individual holds a controlling position in a publicly traded
	// company, is a member of the board of directors, or has policy making abilities
	// in a publicly traded company.
	IsControlPerson bool `json:"is_control_person,required"`
	// Whether the individual is politically exposed.
	IsPoliticallyExposed bool `json:"is_politically_exposed,required"`
	paramObj
}

func (r UsKYCCheckDataRiskDisclosureParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataRiskDisclosureParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataRiskDisclosureParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information for a trusted contact person for the individual. More information:
// <br/> -
// <a href="https://www.investor.gov/introduction-investing/general-resources/news-alerts/alerts-bulletins/investor-bulletins-trusted-contact" target="_blank" rel="noopener noreferrer">Investor.gov -
// Trusted Contact</a> <br/> -
// <a href="https://www.finra.org/investors/insights/trusted-contact" target="_blank" rel="noopener noreferrer">FINRA -
// Trusted Contact</a>
//
// The properties FamilyName, GivenName are required.
type UsKYCCheckDataTrustedContactParam struct {
	// The family name of the trusted contact
	FamilyName string `json:"family_name,required"`
	// The given name of the trusted contact
	GivenName string `json:"given_name,required"`
	// The email address of the trusted contact. At least one of email_address or
	// phone_number is required.
	EmailAddress param.Opt[string] `json:"email_address,omitzero"`
	// The phone number of the trusted contact. At least one of email_address or
	// phone_number is required.
	PhoneNumber param.Opt[string] `json:"phone_number,omitzero"`
	paramObj
}

func (r UsKYCCheckDataTrustedContactParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataTrustedContactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataTrustedContactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// US immigration information for this individual. Required if the individual is
// not a US citizen.
//
// The properties CountryOfBirth, IsPermanentResident are required.
type UsKYCCheckDataUsImmigrationInfoParam struct {
	// Country where the individual was born.
	CountryOfBirth string `json:"country_of_birth,required"`
	// Whether the individual is a US permanent resident (green card holder).
	IsPermanentResident bool `json:"is_permanent_resident,required"`
	// Date the individual is scheduled to leave the US. Required for B1 and B2 visas.
	DepartureFromUsDate param.Opt[time.Time] `json:"departure_from_us_date,omitzero" format:"date"`
	// Expiration date of the visa. Required if visa_type is provided.
	VisaExpirationDate param.Opt[time.Time] `json:"visa_expiration_date,omitzero" format:"date"`
	// Type of visa the individual holds. Required if not a permanent resident.
	//
	// Any of "B1", "B2", "DACA", "E1", "E2", "E3", "F1", "G4", "H1B", "J1", "L1",
	// "Other", "O1", "TN1".
	VisaType string `json:"visa_type,omitzero"`
	paramObj
}

func (r UsKYCCheckDataUsImmigrationInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow UsKYCCheckDataUsImmigrationInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UsKYCCheckDataUsImmigrationInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[UsKYCCheckDataUsImmigrationInfoParam](
		"visa_type", "B1", "B2", "DACA", "E1", "E2", "E3", "F1", "G4", "H1B", "J1", "L1", "Other", "O1", "TN1",
	)
}

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

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Input
	// parameters for providing KYC information for an `Entity` in the baseline
	// jurisdiction.
	OfBaseline *V2EntityKYCSubmitParamsBodyBaseline `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Input
	// parameters for providing KYC information for an `Entity` in the US jurisdiction.
	OfUs *V2EntityKYCSubmitParamsBodyUs `json:",inline"`

	paramObj
}

func (u V2EntityKYCSubmitParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBaseline, u.OfUs)
}
func (r *V2EntityKYCSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input parameters for providing KYC information for an `Entity` in the baseline
// jurisdiction.
//
// The properties Data, ProviderName are required.
type V2EntityKYCSubmitParamsBodyBaseline struct {
	// KYC data for an `Entity` in the BASELINE jurisdiction.
	Data BaselineKYCCheckDataParam `json:"data,omitzero,required"`
	// Name of the KYC provider that provided the KYC information.
	ProviderName string `json:"provider_name,required"`
	// Jurisdiction of the KYC check.
	//
	// Any of "BASELINE".
	Jurisdiction string `json:"jurisdiction,omitzero"`
	paramObj
}

func (r V2EntityKYCSubmitParamsBodyBaseline) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityKYCSubmitParamsBodyBaseline
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityKYCSubmitParamsBodyBaseline) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2EntityKYCSubmitParamsBodyBaseline](
		"jurisdiction", "BASELINE",
	)
}

// Input parameters for providing KYC information for an `Entity` in the US
// jurisdiction.
//
// The properties Data, ProviderName are required.
type V2EntityKYCSubmitParamsBodyUs struct {
	// KYC data for an `Entity` in the US jurisdiction.
	Data UsKYCCheckDataParam `json:"data,omitzero,required"`
	// Name of the KYC provider that provided the KYC information.
	ProviderName string `json:"provider_name,required"`
	// Jurisdiction of the KYC check.
	//
	// Any of "US".
	Jurisdiction string `json:"jurisdiction,omitzero"`
	paramObj
}

func (r V2EntityKYCSubmitParamsBodyUs) MarshalJSON() (data []byte, err error) {
	type shadow V2EntityKYCSubmitParamsBodyUs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2EntityKYCSubmitParamsBodyUs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V2EntityKYCSubmitParamsBodyUs](
		"jurisdiction", "US",
	)
}
