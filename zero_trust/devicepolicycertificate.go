// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DevicePolicyCertificateService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCertificateService] method instead.
type DevicePolicyCertificateService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyCertificateService(opts ...option.RequestOption) (r *DevicePolicyCertificateService) {
	r = &DevicePolicyCertificateService{}
	r.Options = opts
	return
}

// Enable Zero Trust Clients to provision a certificate, containing a x509 subject,
// and referenced by Access device posture policies when the client visits MTLS
// protected domains. This facilitates device posture without a WARP session.
func (r *DevicePolicyCertificateService) Update(ctx context.Context, zoneTag string, body DevicePolicyCertificateUpdateParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env DevicePolicyCertificateUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneTag == "" {
		err = errors.New("missing required zone_tag parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", zoneTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device certificate provisioning
func (r *DevicePolicyCertificateService) Get(ctx context.Context, zoneTag string, opts ...option.RequestOption) (res *interface{}, err error) {
	var env DevicePolicyCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if zoneTag == "" {
		err = errors.New("missing required zone_tag parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", zoneTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyCertificatesParam struct {
	// The current status of the device policy certificate provisioning feature for
	// WARP clients.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r DevicePolicyCertificatesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyCertificateUpdateParams struct {
	DevicePolicyCertificates DevicePolicyCertificatesParam `json:"device_policy_certificates,required"`
}

func (r DevicePolicyCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DevicePolicyCertificates)
}

type DevicePolicyCertificateUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyCertificateUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyCertificateUpdateResponseEnvelope]
type devicePolicyCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCertificateUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyCertificateUpdateResponseEnvelopeSuccessTrue DevicePolicyCertificateUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyCertificateUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCertificateUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyCertificateGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyCertificateGetResponseEnvelope]
type devicePolicyCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCertificateGetResponseEnvelopeSuccess bool

const (
	DevicePolicyCertificateGetResponseEnvelopeSuccessTrue DevicePolicyCertificateGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
