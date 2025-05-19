// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

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

// V2AccountWalletExternalService contains methods and other services that help
// with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountWalletExternalService] method instead.
type V2AccountWalletExternalService struct {
	Options []option.RequestOption
}

// NewV2AccountWalletExternalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountWalletExternalService(opts ...option.RequestOption) (r V2AccountWalletExternalService) {
	r = V2AccountWalletExternalService{}
	r.Options = opts
	return
}

// Connect a `Wallet` to the `Account` after verifying the signature.
func (r *V2AccountWalletExternalService) Connect(ctx context.Context, accountID string, body V2AccountWalletExternalConnectParams, opts ...option.RequestOption) (res *Wallet, err error) {
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
func (r *V2AccountWalletExternalService) GetNonce(ctx context.Context, accountID string, query V2AccountWalletExternalGetNonceParams, opts ...option.RequestOption) (res *V2AccountWalletExternalGetNonceResponse, err error) {
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
type V2AccountWalletExternalGetNonceResponse struct {
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
func (r V2AccountWalletExternalGetNonceResponse) RawJSON() string { return r.JSON.raw }
func (r *V2AccountWalletExternalGetNonceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWalletExternalConnectParams struct {
	// CAIP-2 formatted chain ID of the blockchain the `Wallet` to link is on.
	//
	// Any of "eip155:1", "eip155:42161", "eip155:8453", "eip155:81457", "eip155:7887",
	// "eip155:98866", "eip155:11155111", "eip155:421614", "eip155:84532",
	// "eip155:168587773", "eip155:98867", "eip155:31337", "eip155:1337".
	ChainID Chain `json:"chain_id,omitzero,required"`
	// Nonce contained within the connection message.
	Nonce string `json:"nonce,required" format:"uuid"`
	// Signature payload from signing the connection message with the `Wallet`.
	Signature string `json:"signature,required" format:"hex_string"`
	// Address of the `Wallet`.
	WalletAddress string `json:"wallet_address,required" format:"eth_address"`
	paramObj
}

func (r V2AccountWalletExternalConnectParams) MarshalJSON() (data []byte, err error) {
	type shadow V2AccountWalletExternalConnectParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V2AccountWalletExternalConnectParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V2AccountWalletExternalGetNonceParams struct {
	// Address of the `Wallet` to connect.
	WalletAddress string `query:"wallet_address,required" format:"eth_address" json:"-"`
	paramObj
}

// URLQuery serializes [V2AccountWalletExternalGetNonceParams]'s query parameters
// as `url.Values`.
func (r V2AccountWalletExternalGetNonceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
