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

type UnnamedSchemaRef67c73d4742566cab0909f71b1822e88c struct {
	Fail []string                                             `json:"FAIL,required"`
	None []string                                             `json:"NONE,required"`
	Pass []string                                             `json:"PASS,required"`
	JSON unnamedSchemaRef67c73d4742566cab0909f71b1822e88cJSON `json:"-"`
}

// unnamedSchemaRef67c73d4742566cab0909f71b1822e88cJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef67c73d4742566cab0909f71b1822e88c]
type unnamedSchemaRef67c73d4742566cab0909f71b1822e88cJSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef67c73d4742566cab0909f71b1822e88c) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef67c73d4742566cab0909f71b1822e88cJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef853c157ad369010995e35be614e0343f struct {
	Fail string                                               `json:"FAIL,required"`
	None string                                               `json:"NONE,required"`
	Pass string                                               `json:"PASS,required"`
	JSON unnamedSchemaRef853c157ad369010995e35be614e0343fJSON `json:"-"`
}

// unnamedSchemaRef853c157ad369010995e35be614e0343fJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef853c157ad369010995e35be614e0343f]
type unnamedSchemaRef853c157ad369010995e35be614e0343fJSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef853c157ad369010995e35be614e0343f) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef853c157ad369010995e35be614e0343fJSON) RawJSON() string {
	return r.raw
}
