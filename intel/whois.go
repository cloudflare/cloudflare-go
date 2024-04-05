// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WhoisService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWhoisService] method instead.
type WhoisService struct {
	Options []option.RequestOption
}

// NewWhoisService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhoisService(opts ...option.RequestOption) (r *WhoisService) {
	r = &WhoisService{}
	r.Options = opts
	return
}

// Get WHOIS Record
func (r *WhoisService) Get(ctx context.Context, params WhoisGetParams, opts ...option.RequestOption) (res *Whois, err error) {
	opts = append(r.Options[:], opts...)
	var env WhoisGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/whois", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Whois struct {
	CreatedDate       time.Time `json:"created_date" format:"date"`
	Domain            string    `json:"domain"`
	Nameservers       []string  `json:"nameservers"`
	Registrant        string    `json:"registrant"`
	RegistrantCountry string    `json:"registrant_country"`
	RegistrantEmail   string    `json:"registrant_email"`
	RegistrantOrg     string    `json:"registrant_org"`
	Registrar         string    `json:"registrar"`
	UpdatedDate       time.Time `json:"updated_date" format:"date"`
	JSON              whoisJSON `json:"-"`
}

// whoisJSON contains the JSON metadata for the struct [Whois]
type whoisJSON struct {
	CreatedDate       apijson.Field
	Domain            apijson.Field
	Nameservers       apijson.Field
	Registrant        apijson.Field
	RegistrantCountry apijson.Field
	RegistrantEmail   apijson.Field
	RegistrantOrg     apijson.Field
	Registrar         apijson.Field
	UpdatedDate       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Whois) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisJSON) RawJSON() string {
	return r.raw
}

type WhoisGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [WhoisGetParams]'s query parameters as `url.Values`.
func (r WhoisGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhoisGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Whois                                                     `json:"result,required"`
	// Whether the API call was successful
	Success WhoisGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    whoisGetResponseEnvelopeJSON    `json:"-"`
}

// whoisGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelope]
type whoisGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WhoisGetResponseEnvelopeSuccess bool

const (
	WhoisGetResponseEnvelopeSuccessTrue WhoisGetResponseEnvelopeSuccess = true
)

func (r WhoisGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WhoisGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
