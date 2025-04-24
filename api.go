// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"github.com/dinaricrypto/dinari-api-sdk-go/option"
)

// APIService contains methods and other services that help with interacting with
// the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIService] method instead.
type APIService struct {
	Options []option.RequestOption
	V2      APIV2Service
}

// NewAPIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIService(opts ...option.RequestOption) (r APIService) {
	r = APIService{}
	r.Options = opts
	r.V2 = NewAPIV2Service(opts...)
	return
}
