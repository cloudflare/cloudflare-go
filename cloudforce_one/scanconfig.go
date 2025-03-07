// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ScanConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanConfigService] method instead.
type ScanConfigService struct {
	Options []option.RequestOption
}

// NewScanConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScanConfigService(opts ...option.RequestOption) (r *ScanConfigService) {
	r = &ScanConfigService{}
	r.Options = opts
	return
}

// Create a new Scan Config
func (r *ScanConfigService) New(ctx context.Context, params ScanConfigNewParams, opts ...option.RequestOption) (res *ScanConfigNewResponse, err error) {
	var env ScanConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Scan Config for An Account
func (r *ScanConfigService) List(ctx context.Context, query ScanConfigListParams, opts ...option.RequestOption) (res *ScanConfigListResponse, err error) {
	var env ScanConfigListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the Scan Config for an Account
func (r *ScanConfigService) Delete(ctx context.Context, body ScanConfigDeleteParams, opts ...option.RequestOption) (res *ScanConfigDeleteResponse, err error) {
	var env ScanConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScanConfigNewResponse struct {
	AccountID string                    `json:"account_id,required"`
	Frequency float64                   `json:"frequency,required"`
	IPs       []string                  `json:"ips,required"`
	JSON      scanConfigNewResponseJSON `json:"-"`
}

// scanConfigNewResponseJSON contains the JSON metadata for the struct
// [ScanConfigNewResponse]
type scanConfigNewResponseJSON struct {
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScanConfigListResponse struct {
	AccountID string                     `json:"account_id,required"`
	Frequency float64                    `json:"frequency,required"`
	IPs       []string                   `json:"ips,required"`
	JSON      scanConfigListResponseJSON `json:"-"`
}

// scanConfigListResponseJSON contains the JSON metadata for the struct
// [ScanConfigListResponse]
type scanConfigListResponseJSON struct {
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanConfigDeleteResponse = interface{}

type ScanConfigNewParams struct {
	// Account ID
	AccountID param.Field[string]   `path:"account_id,required"`
	Frequency param.Field[float64]  `json:"frequency,required"`
	IPs       param.Field[[]string] `json:"ips,required"`
}

func (r ScanConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScanConfigNewResponseEnvelope struct {
	Errors   []string                          `json:"errors,required"`
	Messages []string                          `json:"messages,required"`
	Result   ScanConfigNewResponse             `json:"result,required"`
	Success  bool                              `json:"success,required"`
	JSON     scanConfigNewResponseEnvelopeJSON `json:"-"`
}

// scanConfigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigNewResponseEnvelope]
type scanConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanConfigListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanConfigListResponseEnvelope struct {
	Errors   []string                           `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   ScanConfigListResponse             `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     scanConfigListResponseEnvelopeJSON `json:"-"`
}

// scanConfigListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigListResponseEnvelope]
type scanConfigListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanConfigDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanConfigDeleteResponseEnvelope struct {
	Errors   []string                             `json:"errors,required"`
	Messages []string                             `json:"messages,required"`
	Result   ScanConfigDeleteResponse             `json:"result,required"`
	Success  bool                                 `json:"success,required"`
	JSON     scanConfigDeleteResponseEnvelopeJSON `json:"-"`
}

// scanConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigDeleteResponseEnvelope]
type scanConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
