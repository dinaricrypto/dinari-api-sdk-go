// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/dinari-go/internal/apiform"
	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/apiquery"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/respjson"
)

// V2EntityKYCDocumentService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2EntityKYCDocumentService] method instead.
type V2EntityKYCDocumentService struct {
	Options []option.RequestOption
}

// NewV2EntityKYCDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2EntityKYCDocumentService(opts ...option.RequestOption) (r V2EntityKYCDocumentService) {
	r = V2EntityKYCDocumentService{}
	r.Options = opts
	return
}

// Get uploaded documents for a KYC check
func (r *V2EntityKYCDocumentService) Get(ctx context.Context, kycID string, query V2EntityKYCDocumentGetParams, opts ...option.RequestOption) (res *[]KYCDocument, err error) {
	opts = append(r.Options[:], opts...)
	if query.EntityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	if kycID == "" {
		err = errors.New("missing required kyc_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc/%s/document", query.EntityID, kycID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload KYC-related documentation for partners that are provisioned to provide
// their own KYC data.
func (r *V2EntityKYCDocumentService) Upload(ctx context.Context, kycID string, params V2EntityKYCDocumentUploadParams, opts ...option.RequestOption) (res *KYCDocument, err error) {
	opts = append(r.Options[:], opts...)
	if params.EntityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	if kycID == "" {
		err = errors.New("missing required kyc_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/entities/%s/kyc/%s/document", params.EntityID, kycID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// A document associated with KYC for an `Entity`.
type KYCDocument struct {
	// ID of the document.
	ID string `json:"id,required" format:"uuid"`
	// Type of document.
	//
	// Any of "GOVERNMENT_ID", "SELFIE", "RESIDENCY", "UNKNOWN".
	DocumentType KYCDocumentType `json:"document_type,required"`
	// Filename of document.
	Filename string `json:"filename,required"`
	// Temporary URL to access the document. Expires in 1 hour.
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DocumentType respjson.Field
		Filename     respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KYCDocument) RawJSON() string { return r.JSON.raw }
func (r *KYCDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KYCDocumentType string

const (
	KYCDocumentTypeGovernmentID KYCDocumentType = "GOVERNMENT_ID"
	KYCDocumentTypeSelfie       KYCDocumentType = "SELFIE"
	KYCDocumentTypeResidency    KYCDocumentType = "RESIDENCY"
	KYCDocumentTypeUnknown      KYCDocumentType = "UNKNOWN"
)

type V2EntityKYCDocumentGetParams struct {
	EntityID string `path:"entity_id,required" format:"uuid" json:"-"`
	paramObj
}

type V2EntityKYCDocumentUploadParams struct {
	EntityID string `path:"entity_id,required" format:"uuid" json:"-"`
	// Type of `KYCDocument` to be uploaded.
	//
	// Any of "GOVERNMENT_ID", "SELFIE", "RESIDENCY", "UNKNOWN".
	DocumentType KYCDocumentType `query:"document_type,omitzero,required" json:"-"`
	// File to be uploaded. Must be a valid image or PDF file (jpg, jpeg, png, pdf)
	// less than 10MB in size.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	paramObj
}

func (r V2EntityKYCDocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [V2EntityKYCDocumentUploadParams]'s query parameters as
// `url.Values`.
func (r V2EntityKYCDocumentUploadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
