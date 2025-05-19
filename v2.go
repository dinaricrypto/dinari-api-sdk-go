// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dinari

import (
	"github.com/stainless-sdks/dinari-go/option"
)

// V2Service contains methods and other services that help with interacting with
// the dinari API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2Service] method instead.
type V2Service struct {
	Options    []option.RequestOption
	MarketData V2MarketDataService
	Entities   V2EntityService
	Accounts   V2AccountService
}

// NewV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2Service(opts ...option.RequestOption) (r V2Service) {
	r = V2Service{}
	r.Options = opts
	r.MarketData = NewV2MarketDataService(opts...)
	r.Entities = NewV2EntityService(opts...)
	r.Accounts = NewV2AccountService(opts...)
	return
}
