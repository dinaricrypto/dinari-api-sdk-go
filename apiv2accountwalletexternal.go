// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/dinari-go/internal/apijson"
	"github.com/stainless-sdks/dinari-go/internal/apiquery"
	"github.com/stainless-sdks/dinari-go/internal/requestconfig"
	"github.com/stainless-sdks/dinari-go/option"
	"github.com/stainless-sdks/dinari-go/packages/param"
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2AccountWalletExternalService contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountWalletExternalService] method instead.
type APIV2AccountWalletExternalService struct {
	Options []option.RequestOption
}

// NewAPIV2AccountWalletExternalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIV2AccountWalletExternalService(opts ...option.RequestOption) (r APIV2AccountWalletExternalService) {
	r = APIV2AccountWalletExternalService{}
	r.Options = opts
	return
}

// Connects a wallet to the account using the nonce and signature
func (r *APIV2AccountWalletExternalService) Connect(ctx context.Context, accountID string, body APIV2AccountWalletExternalConnectParams, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/wallet/external", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a nonce and message to be signed in order to verify wallet ownership.
func (r *APIV2AccountWalletExternalService) GetNonce(ctx context.Context, accountID string, query APIV2AccountWalletExternalGetNonceParams, opts ...option.RequestOption) (res *Apiv2AccountWalletExternalGetNonceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/wallet/external/nonce", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Wallet connection message for Dinari-managed wallets
type Apiv2AccountWalletExternalGetNonceResponse struct {
	// Message to be signed by the wallet
	Message string `json:"message,required"`
	// Single-use identifier
	Nonce string `json:"nonce,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Message     resp.Field
		Nonce       resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountWalletExternalGetNonceResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountWalletExternalGetNonceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2AccountWalletExternalConnectParams struct {
	// Blockchain the wallet to link is on
	ChainID int64 `json:"chain_id,required"`
	// Nonce used to sign the wallet connection message
	Nonce string `json:"nonce,required" format:"uuid"`
	// Signature payload from signing the wallet connection message with the wallet
	Signature string `json:"signature,required" format:"hex_string"`
	// Address of the wallet
	WalletAddress string `json:"wallet_address,required" format:"eth_address"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountWalletExternalConnectParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r APIV2AccountWalletExternalConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2AccountWalletExternalConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}

type APIV2AccountWalletExternalGetNonceParams struct {
	// Address of the wallet to connect
	WalletAddress string `query:"wallet_address,required" format:"eth_address" json:"-"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f APIV2AccountWalletExternalGetNonceParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

// URLQuery serializes [APIV2AccountWalletExternalGetNonceParams]'s query
// parameters as `url.Values`.
func (r APIV2AccountWalletExternalGetNonceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
