// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SinkholeIngressService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSinkholeIngressService] method instead.
type SinkholeIngressService struct {
	Options []option.RequestOption
}

// NewSinkholeIngressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSinkholeIngressService(opts ...option.RequestOption) (r *SinkholeIngressService) {
	r = &SinkholeIngressService{}
	r.Options = opts
	return
}

// Create a new ingress rule for the specified sinkhole. The CIDR block must be a
// Cloudflare BYOIP associated with your account. The zone_id must be a zone with
// the ability to create Spectrum Apps. The sinkhole must belong to the same
// account as the zone.
func (r *SinkholeIngressService) New(ctx context.Context, sinkholeID string, params SinkholeIngressNewParams, opts ...option.RequestOption) (res *SinkholeIngressNewResponse, err error) {
	var env SinkholeIngressNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/intel/sinkholes/%s/ingresses", params.ZoneID, sinkholeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the specified ingress rule. The sinkhole must belong to the same
// account as the zone.
func (r *SinkholeIngressService) Update(ctx context.Context, sinkholeID string, ingressID string, params SinkholeIngressUpdateParams, opts ...option.RequestOption) (res *SinkholeIngressUpdateResponse, err error) {
	var env SinkholeIngressUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	if ingressID == "" {
		err = errors.New("missing required ingress_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/intel/sinkholes/%s/ingresses/%s", params.ZoneID, sinkholeID, ingressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete the specified ingress rule. The sinkhole must belong to the same account
// as the zone.
func (r *SinkholeIngressService) Delete(ctx context.Context, sinkholeID string, ingressID string, body SinkholeIngressDeleteParams, opts ...option.RequestOption) (res *SinkholeIngressDeleteResponse, err error) {
	var env SinkholeIngressDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	if ingressID == "" {
		err = errors.New("missing required ingress_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/intel/sinkholes/%s/ingresses/%s", body.ZoneID, sinkholeID, ingressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get the specified ingress rule associated with a sinkhole. The sinkhole must
// belong to the same account as the zone.
func (r *SinkholeIngressService) Get(ctx context.Context, sinkholeID string, ingressID string, query SinkholeIngressGetParams, opts ...option.RequestOption) (res *SinkholeIngressGetResponse, err error) {
	var env SinkholeIngressGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if sinkholeID == "" {
		err = errors.New("missing required sinkhole_id parameter")
		return nil, err
	}
	if ingressID == "" {
		err = errors.New("missing required ingress_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/intel/sinkholes/%s/ingresses/%s", query.ZoneID, sinkholeID, ingressID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SinkholeIngressNewResponse struct {
	// The unique identifier for the ingress rule.
	ID string `json:"id"`
	// The CIDR block for the ingress rule.
	CIDR string `json:"cidr"`
	// The date and time when the ingress rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the ingress rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The sinkhole this ingress rule belongs to.
	SinkholeID string `json:"sinkhole_id"`
	// The zone tag associated with this ingress rule.
	ZoneTag string                         `json:"zone_tag"`
	JSON    sinkholeIngressNewResponseJSON `json:"-"`
}

// sinkholeIngressNewResponseJSON contains the JSON metadata for the struct
// [SinkholeIngressNewResponse]
type sinkholeIngressNewResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SinkholeID  apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressUpdateResponse = interface{}

type SinkholeIngressDeleteResponse = interface{}

type SinkholeIngressGetResponse struct {
	// The unique identifier for the ingress rule.
	ID string `json:"id"`
	// The CIDR block for the ingress rule.
	CIDR string `json:"cidr"`
	// The date and time when the ingress rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the ingress rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The sinkhole this ingress rule belongs to.
	SinkholeID string `json:"sinkhole_id"`
	// The zone tag associated with this ingress rule.
	ZoneTag string                         `json:"zone_tag"`
	JSON    sinkholeIngressGetResponseJSON `json:"-"`
}

// sinkholeIngressGetResponseJSON contains the JSON metadata for the struct
// [SinkholeIngressGetResponse]
type sinkholeIngressGetResponseJSON struct {
	ID          apijson.Field
	CIDR        apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SinkholeID  apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressNewParams struct {
	// An identifier for the resource.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The CIDR block for the ingress rule in IPv4 or IPv6 notation (e.g.,
	// 192.0.2.0/24). Provide a Cloudflare BYOIP CIDR that your account owns.
	CIDR param.Field[string] `json:"cidr" api:"required"`
}

func (r SinkholeIngressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkholeIngressNewResponseEnvelope struct {
	Errors   []SinkholeIngressNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeIngressNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeIngressNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeIngressNewResponse                `json:"result"`
	JSON    sinkholeIngressNewResponseEnvelopeJSON    `json:"-"`
}

// sinkholeIngressNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeIngressNewResponseEnvelope]
type sinkholeIngressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SinkholeIngressNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeIngressNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeIngressNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SinkholeIngressNewResponseEnvelopeErrors]
type sinkholeIngressNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    sinkholeIngressNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeIngressNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressNewResponseEnvelopeErrorsSource]
type sinkholeIngressNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           SinkholeIngressNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeIngressNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeIngressNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SinkholeIngressNewResponseEnvelopeMessages]
type sinkholeIngressNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    sinkholeIngressNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeIngressNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressNewResponseEnvelopeMessagesSource]
type sinkholeIngressNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeIngressNewResponseEnvelopeSuccess bool

const (
	SinkholeIngressNewResponseEnvelopeSuccessTrue SinkholeIngressNewResponseEnvelopeSuccess = true
)

func (r SinkholeIngressNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeIngressNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeIngressUpdateParams struct {
	// An identifier for the resource.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The CIDR block for the ingress rule in IPv4 or IPv6 notation (e.g.,
	// 192.0.2.0/24). Provide a Cloudflare BYOIP CIDR that your account owns.
	CIDR param.Field[string] `json:"cidr" api:"required"`
}

func (r SinkholeIngressUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SinkholeIngressUpdateResponseEnvelope struct {
	Errors   []SinkholeIngressUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeIngressUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeIngressUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeIngressUpdateResponse                `json:"result"`
	JSON    sinkholeIngressUpdateResponseEnvelopeJSON    `json:"-"`
}

// sinkholeIngressUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SinkholeIngressUpdateResponseEnvelope]
type sinkholeIngressUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressUpdateResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SinkholeIngressUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeIngressUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeIngressUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SinkholeIngressUpdateResponseEnvelopeErrors]
type sinkholeIngressUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    sinkholeIngressUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeIngressUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressUpdateResponseEnvelopeErrorsSource]
type sinkholeIngressUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressUpdateResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SinkholeIngressUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeIngressUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeIngressUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SinkholeIngressUpdateResponseEnvelopeMessages]
type sinkholeIngressUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    sinkholeIngressUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeIngressUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SinkholeIngressUpdateResponseEnvelopeMessagesSource]
type sinkholeIngressUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeIngressUpdateResponseEnvelopeSuccess bool

const (
	SinkholeIngressUpdateResponseEnvelopeSuccessTrue SinkholeIngressUpdateResponseEnvelopeSuccess = true
)

func (r SinkholeIngressUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeIngressUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeIngressDeleteParams struct {
	// An identifier for the resource.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SinkholeIngressDeleteResponseEnvelope struct {
	Errors   []SinkholeIngressDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeIngressDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeIngressDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeIngressDeleteResponse                `json:"result"`
	JSON    sinkholeIngressDeleteResponseEnvelopeJSON    `json:"-"`
}

// sinkholeIngressDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SinkholeIngressDeleteResponseEnvelope]
type sinkholeIngressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SinkholeIngressDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeIngressDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeIngressDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SinkholeIngressDeleteResponseEnvelopeErrors]
type sinkholeIngressDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    sinkholeIngressDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeIngressDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressDeleteResponseEnvelopeErrorsSource]
type sinkholeIngressDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SinkholeIngressDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeIngressDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeIngressDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SinkholeIngressDeleteResponseEnvelopeMessages]
type sinkholeIngressDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    sinkholeIngressDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeIngressDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SinkholeIngressDeleteResponseEnvelopeMessagesSource]
type sinkholeIngressDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeIngressDeleteResponseEnvelopeSuccess bool

const (
	SinkholeIngressDeleteResponseEnvelopeSuccessTrue SinkholeIngressDeleteResponseEnvelopeSuccess = true
)

func (r SinkholeIngressDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeIngressDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SinkholeIngressGetParams struct {
	// An identifier for the resource.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SinkholeIngressGetResponseEnvelope struct {
	Errors   []SinkholeIngressGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SinkholeIngressGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SinkholeIngressGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SinkholeIngressGetResponse                `json:"result"`
	JSON    sinkholeIngressGetResponseEnvelopeJSON    `json:"-"`
}

// sinkholeIngressGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SinkholeIngressGetResponseEnvelope]
type sinkholeIngressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SinkholeIngressGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             sinkholeIngressGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// sinkholeIngressGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SinkholeIngressGetResponseEnvelopeErrors]
type sinkholeIngressGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    sinkholeIngressGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// sinkholeIngressGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressGetResponseEnvelopeErrorsSource]
type sinkholeIngressGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           SinkholeIngressGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             sinkholeIngressGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// sinkholeIngressGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SinkholeIngressGetResponseEnvelopeMessages]
type sinkholeIngressGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SinkholeIngressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SinkholeIngressGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    sinkholeIngressGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// sinkholeIngressGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SinkholeIngressGetResponseEnvelopeMessagesSource]
type sinkholeIngressGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SinkholeIngressGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sinkholeIngressGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SinkholeIngressGetResponseEnvelopeSuccess bool

const (
	SinkholeIngressGetResponseEnvelopeSuccessTrue SinkholeIngressGetResponseEnvelopeSuccess = true
)

func (r SinkholeIngressGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SinkholeIngressGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
