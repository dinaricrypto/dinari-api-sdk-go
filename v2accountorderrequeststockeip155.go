// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

// V2AccountOrderRequestStockEip155Service contains methods and other services that
// help with interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderRequestStockEip155Service] method instead.
type V2AccountOrderRequestStockEip155Service struct {
	Options []option.RequestOption
}

// NewV2AccountOrderRequestStockEip155Service generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV2AccountOrderRequestStockEip155Service(opts ...option.RequestOption) (r V2AccountOrderRequestStockEip155Service) {
	r = V2AccountOrderRequestStockEip155Service{}
	r.Options = opts
	return
}
