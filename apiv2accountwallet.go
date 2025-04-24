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
	"github.com/stainless-sdks/dinari-go/packages/resp"
)

// APIV2AccountWalletService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIV2AccountWalletService] method instead.
type APIV2AccountWalletService struct {
	Options  []option.RequestOption
	External APIV2AccountWalletExternalService
}

// NewAPIV2AccountWalletService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAPIV2AccountWalletService(opts ...option.RequestOption) (r APIV2AccountWalletService) {
	r = APIV2AccountWalletService{}
	r.Options = opts
	r.External = NewAPIV2AccountWalletExternalService(opts...)
	return
}

// Retrieves details of the wallet connected to the account.
func (r *APIV2AccountWalletService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Wallet, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("api/v2/accounts/%s/wallet", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Information about a digital wallet
type Wallet struct {
	// Address of the wallet
	Address string `json:"address,required"`
	// Indicates whether the wallet is flagged for AML violations
	IsAmlFlagged bool `json:"is_aml_flagged,required"`
	// Indicates whether the wallet is a Dinari-managed wallet
	IsManagedWallet bool `json:"is_managed_wallet,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Address         resp.Field
		IsAmlFlagged    resp.Field
		IsManagedWallet resp.Field
		ExtraFields     map[string]resp.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Wallet) RawJSON() string { return r.JSON.raw }
func (r *Wallet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
