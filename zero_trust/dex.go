// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DEXService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDEXService] method instead.
type DEXService struct {
	Options               []option.RequestOption
	Colos                 *DEXColoService
	FleetStatus           *DEXFleetStatusService
	HTTPTests             *DEXHTTPTestService
	Tests                 *DEXTestService
	TracerouteTestResults *DEXTracerouteTestResultService
	TracerouteTests       *DEXTracerouteTestService
}

// NewDEXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXService(opts ...option.RequestOption) (r *DEXService) {
	r = &DEXService{}
	r.Options = opts
	r.Colos = NewDEXColoService(opts...)
	r.FleetStatus = NewDEXFleetStatusService(opts...)
	r.HTTPTests = NewDEXHTTPTestService(opts...)
	r.Tests = NewDEXTestService(opts...)
	r.TracerouteTestResults = NewDEXTracerouteTestResultService(opts...)
	r.TracerouteTests = NewDEXTracerouteTestService(opts...)
	return
}

type UnnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11 struct {
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                                                 `json:"default,required"`
	Name    string                                               `json:"name,required"`
	JSON    unnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11JSON `json:"-"`
}

// unnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11]
type unnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11JSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefBf9e2abcf1b78a6cab8e6e29e2228a11JSON) RawJSON() string {
	return r.raw
}
