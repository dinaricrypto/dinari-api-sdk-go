// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinariapisdkgo

import (
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

// V2AccountOrderStockService contains methods and other services that help with
// interacting with the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2AccountOrderStockService] method instead.
type V2AccountOrderStockService struct {
	Options []option.RequestOption
	Eip155  V2AccountOrderStockEip155Service
}

// NewV2AccountOrderStockService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV2AccountOrderStockService(opts ...option.RequestOption) (r V2AccountOrderStockService) {
	r = V2AccountOrderStockService{}
	r.Options = opts
	r.Eip155 = NewV2AccountOrderStockEip155Service(opts...)
	return
}
