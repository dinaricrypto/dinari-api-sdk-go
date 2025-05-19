// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/dinaricrypto/dinari-api-sdk-go/internal/apijson"
	"github.com/dinaricrypto/dinari-api-sdk-go/internal/requestconfig"
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
	"github.com/dinaricrypto/dinari-api-sdk-go/packages/respjson"
)

// V2AccountWalletService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountWalletService] method instead.
type V2AccountWalletService struct {
	Options  []option.RequestOption
	External V2AccountWalletExternalService
}

// NewV2AccountWalletService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV2AccountWalletService(opts ...option.RequestOption) (r V2AccountWalletService) {
	r = V2AccountWalletService{}
	r.Options = opts
	r.External = NewV2AccountWalletExternalService(opts...)
	return
}

// Get the wallet connected to the `Account`.
func (r *V2AccountWalletService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/wallet", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about a blockchain `Wallet`.
type Wallet struct {
	// Address of the `Wallet`.
	Address string `json:"address,required"`
	// Indicates whether the `Wallet` is flagged for AML violation.
	IsAmlFlagged bool `json:"is_aml_flagged,required"`
	// Indicates whether the `Wallet` is a Dinari-managed wallet.
	IsManagedWallet bool `json:"is_managed_wallet,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address         respjson.Field
		IsAmlFlagged    respjson.Field
		IsManagedWallet respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Wallet) RawJSON() string { return r.JSON.raw }
func (r *Wallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
