// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_txt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SecurityTXTService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityTXTService] method instead.
type SecurityTXTService struct {
	Options []option.RequestOption
}

// NewSecurityTXTService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityTXTService(opts ...option.RequestOption) (r *SecurityTXTService) {
	r = &SecurityTXTService{}
	r.Options = opts
	return
}

// Update security.txt
func (r *SecurityTXTService) Update(ctx context.Context, params SecurityTXTUpdateParams, opts ...option.RequestOption) (res *SecurityTXTUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Delete security.txt
func (r *SecurityTXTService) Delete(ctx context.Context, body SecurityTXTDeleteParams, opts ...option.RequestOption) (res *SecurityTXTDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get security.txt
func (r *SecurityTXTService) Get(ctx context.Context, query SecurityTXTGetParams, opts ...option.RequestOption) (res *SecurityTXTGetResponse, err error) {
	var env SecurityTXTGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/securitytxt", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecurityTXTUpdateResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SecurityTXTUpdateResponseSuccess `json:"success,required"`
	JSON    securityTXTUpdateResponseJSON    `json:"-"`
}

// securityTXTUpdateResponseJSON contains the JSON metadata for the struct
// [SecurityTXTUpdateResponse]
type securityTXTUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityTXTUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityTXTUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecurityTXTUpdateResponseSuccess bool

const (
	SecurityTXTUpdateResponseSuccessTrue SecurityTXTUpdateResponseSuccess = true
)

func (r SecurityTXTUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case SecurityTXTUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type SecurityTXTDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SecurityTXTDeleteResponseSuccess `json:"success,required"`
	JSON    securityTXTDeleteResponseJSON    `json:"-"`
}

// securityTXTDeleteResponseJSON contains the JSON metadata for the struct
// [SecurityTXTDeleteResponse]
type securityTXTDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityTXTDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityTXTDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecurityTXTDeleteResponseSuccess bool

const (
	SecurityTXTDeleteResponseSuccessTrue SecurityTXTDeleteResponseSuccess = true
)

func (r SecurityTXTDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case SecurityTXTDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type SecurityTXTGetResponse struct {
	Acknowledgments    []string                   `json:"acknowledgments" format:"uri"`
	Canonical          []string                   `json:"canonical" format:"uri"`
	Contact            []string                   `json:"contact" format:"uri"`
	Enabled            bool                       `json:"enabled"`
	Encryption         []string                   `json:"encryption" format:"uri"`
	Expires            time.Time                  `json:"expires" format:"date-time"`
	Hiring             []string                   `json:"hiring" format:"uri"`
	Policy             []string                   `json:"policy" format:"uri"`
	PreferredLanguages string                     `json:"preferredLanguages"`
	JSON               securityTXTGetResponseJSON `json:"-"`
}

// securityTXTGetResponseJSON contains the JSON metadata for the struct
// [SecurityTXTGetResponse]
type securityTXTGetResponseJSON struct {
	Acknowledgments    apijson.Field
	Canonical          apijson.Field
	Contact            apijson.Field
	Enabled            apijson.Field
	Encryption         apijson.Field
	Expires            apijson.Field
	Hiring             apijson.Field
	Policy             apijson.Field
	PreferredLanguages apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecurityTXTGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityTXTGetResponseJSON) RawJSON() string {
	return r.raw
}

type SecurityTXTUpdateParams struct {
	// Identifier
	ZoneID             param.Field[string]    `path:"zone_id,required"`
	Acknowledgments    param.Field[[]string]  `json:"acknowledgments" format:"uri"`
	Canonical          param.Field[[]string]  `json:"canonical" format:"uri"`
	Contact            param.Field[[]string]  `json:"contact" format:"uri"`
	Enabled            param.Field[bool]      `json:"enabled"`
	Encryption         param.Field[[]string]  `json:"encryption" format:"uri"`
	Expires            param.Field[time.Time] `json:"expires" format:"date-time"`
	Hiring             param.Field[[]string]  `json:"hiring" format:"uri"`
	Policy             param.Field[[]string]  `json:"policy" format:"uri"`
	PreferredLanguages param.Field[string]    `json:"preferredLanguages"`
}

func (r SecurityTXTUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityTXTDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SecurityTXTGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SecurityTXTGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SecurityTXTGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SecurityTXTGetResponse                `json:"result"`
	JSON    securityTXTGetResponseEnvelopeJSON    `json:"-"`
}

// securityTXTGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SecurityTXTGetResponseEnvelope]
type securityTXTGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityTXTGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityTXTGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SecurityTXTGetResponseEnvelopeSuccess bool

const (
	SecurityTXTGetResponseEnvelopeSuccessTrue SecurityTXTGetResponseEnvelopeSuccess = true
)

func (r SecurityTXTGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SecurityTXTGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
