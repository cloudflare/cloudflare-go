// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// EmailService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewEmailService] method instead.
type EmailService struct {
	Options  []option.RequestOption
	Routing  *EmailRoutingService
	Security *EmailSecurityService
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r *EmailService) {
	r = &EmailService{}
	r.Options = opts
	r.Routing = NewEmailRoutingService(opts...)
	r.Security = NewEmailSecurityService(opts...)
	return
}

type UnnamedSchemaRef149 struct {
	Fail string                  `json:"FAIL,required"`
	None string                  `json:"NONE,required"`
	Pass string                  `json:"PASS,required"`
	JSON unnamedSchemaRef149JSON `json:"-"`
}

// unnamedSchemaRef149JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef149]
type unnamedSchemaRef149JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef149) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef149JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef150 struct {
	Fail []string                `json:"FAIL,required"`
	None []string                `json:"NONE,required"`
	Pass []string                `json:"PASS,required"`
	JSON unnamedSchemaRef150JSON `json:"-"`
}

// unnamedSchemaRef150JSON contains the JSON metadata for the struct
// [UnnamedSchemaRef150]
type unnamedSchemaRef150JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef150) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef150JSON) RawJSON() string {
	return r.raw
}
