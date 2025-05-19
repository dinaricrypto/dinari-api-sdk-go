// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apiquery"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/param"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
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

// Connect a `Wallet` to the `Account` after verifying the signature.
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

// Get a nonce and message to be signed in order to verify `Wallet` ownership.
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

// Connection message to sign to prove ownership of the `Wallet`.
type Apiv2AccountWalletExternalGetNonceResponse struct {
	// Message to be signed by the `Wallet`
	Message string `json:"message,required"`
	// Single-use identifier
	Nonce string `json:"nonce,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Nonce       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Apiv2AccountWalletExternalGetNonceResponse) RawJSON() string { return r.JSON.raw }
func (r *Apiv2AccountWalletExternalGetNonceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIV2AccountWalletExternalConnectParams struct {
	// CAIP-2 formatted chain ID of the blockchain the `Wallet` to link is on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866".
	ChainID APIV2AccountWalletExternalConnectParamsChainID `json:"chain_id,omitzero,required"`
	// Nonce contained within the connection message.
	Nonce string `json:"nonce,required" format:"uuid"`
	// Signature payload from signing the connection message with the `Wallet`.
	Signature string `json:"signature,required" format:"hex_string"`
	// Address of the `Wallet`.
	WalletAddress string `json:"wallet_address,required" format:"eth_address"`
	paramObj
}

func (r APIV2AccountWalletExternalConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow APIV2AccountWalletExternalConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIV2AccountWalletExternalConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CAIP-2 formatted chain ID of the blockchain the `Wallet` to link is on.
type APIV2AccountWalletExternalConnectParamsChainID string

const (
	APIV2AccountWalletExternalConnectParamsChainIDEip155_1     APIV2AccountWalletExternalConnectParamsChainID = "eip155:1"
	APIV2AccountWalletExternalConnectParamsChainIDEip155_42161 APIV2AccountWalletExternalConnectParamsChainID = "eip155:42161"
	APIV2AccountWalletExternalConnectParamsChainIDEip155_8453  APIV2AccountWalletExternalConnectParamsChainID = "eip155:8453"
	APIV2AccountWalletExternalConnectParamsChainIDEip155_81457 APIV2AccountWalletExternalConnectParamsChainID = "eip155:81457"
	APIV2AccountWalletExternalConnectParamsChainIDEip155_7887  APIV2AccountWalletExternalConnectParamsChainID = "eip155:7887"
	APIV2AccountWalletExternalConnectParamsChainIDEip155_98866 APIV2AccountWalletExternalConnectParamsChainID = "eip155:98866"
)

type APIV2AccountWalletExternalGetNonceParams struct {
	// Address of the `Wallet` to connect.
	WalletAddress string `query:"wallet_address,required" format:"eth_address" json:"-"`
	paramObj
}

// URLQuery serializes [APIV2AccountWalletExternalGetNonceParams]'s query
// parameters as `url.Values`.
func (r APIV2AccountWalletExternalGetNonceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
