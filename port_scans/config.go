// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package port_scans

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

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	return
}

// Create a new Scan Config
func (r *ConfigService) New(ctx context.Context, params ConfigNewParams, opts ...option.RequestOption) (res *ConfigNewResponse, err error) {
	var env ConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("%s/scans/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the Scan Config for an Account
func (r *ConfigService) Delete(ctx context.Context, body ConfigDeleteParams, opts ...option.RequestOption) (res *ConfigDeleteResponse, err error) {
	var env ConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("%s/scans/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Scan Config for An Account
func (r *ConfigService) Get(ctx context.Context, query ConfigGetParams, opts ...option.RequestOption) (res *ConfigGetResponse, err error) {
	var env ConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("%s/scans/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigNewResponse struct {
	AccountID string                `json:"account_id,required"`
	Frequency float64               `json:"frequency,required"`
	IPs       []string              `json:"ips,required"`
	JSON      configNewResponseJSON `json:"-"`
}

// configNewResponseJSON contains the JSON metadata for the struct
// [ConfigNewResponse]
type configNewResponseJSON struct {
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteResponse = interface{}

type ConfigGetResponse struct {
	AccountID string                `json:"account_id,required"`
	Frequency float64               `json:"frequency,required"`
	IPs       []string              `json:"ips,required"`
	JSON      configGetResponseJSON `json:"-"`
}

// configGetResponseJSON contains the JSON metadata for the struct
// [ConfigGetResponse]
type configGetResponseJSON struct {
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigNewParams struct {
	// Account ID
	AccountID param.Field[string]   `path:"account_id,required"`
	Frequency param.Field[float64]  `json:"frequency,required"`
	IPs       param.Field[[]string] `json:"ips,required"`
}

func (r ConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigNewResponseEnvelope struct {
	Errors   []string                      `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   ConfigNewResponse             `json:"result,required"`
	Success  bool                          `json:"success,required"`
	JSON     configNewResponseEnvelopeJSON `json:"-"`
}

// configNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigNewResponseEnvelope]
type configNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigDeleteResponseEnvelope struct {
	Errors   []string                         `json:"errors,required"`
	Messages []string                         `json:"messages,required"`
	Result   ConfigDeleteResponse             `json:"result,required"`
	Success  bool                             `json:"success,required"`
	JSON     configDeleteResponseEnvelopeJSON `json:"-"`
}

// configDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigDeleteResponseEnvelope]
type configDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConfigGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigGetResponseEnvelope struct {
	Errors   []string                      `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   ConfigGetResponse             `json:"result,required"`
	Success  bool                          `json:"success,required"`
	JSON     configGetResponseEnvelopeJSON `json:"-"`
}

// configGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigGetResponseEnvelope]
type configGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
