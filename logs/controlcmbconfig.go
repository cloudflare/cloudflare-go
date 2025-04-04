// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ControlCmbConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlCmbConfigService] method instead.
type ControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewControlCmbConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlCmbConfigService(opts ...option.RequestOption) (r *ControlCmbConfigService) {
	r = &ControlCmbConfigService{}
	r.Options = opts
	return
}

// Updates CMB config.
func (r *ControlCmbConfigService) New(ctx context.Context, params ControlCmbConfigNewParams, opts ...option.RequestOption) (res *CmbConfig, err error) {
	var env ControlCmbConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes CMB config.
func (r *ControlCmbConfigService) Delete(ctx context.Context, body ControlCmbConfigDeleteParams, opts ...option.RequestOption) (res *ControlCmbConfigDeleteResponse, err error) {
	var env ControlCmbConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets CMB config.
func (r *ControlCmbConfigService) Get(ctx context.Context, query ControlCmbConfigGetParams, opts ...option.RequestOption) (res *CmbConfig, err error) {
	var env ControlCmbConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CmbConfig struct {
	// Comma-separated list of regions.
	Regions string        `json:"regions"`
	JSON    cmbConfigJSON `json:"-"`
}

// cmbConfigJSON contains the JSON metadata for the struct [CmbConfig]
type cmbConfigJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cmbConfigJSON) RawJSON() string {
	return r.raw
}

type CmbConfigParam struct {
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r CmbConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ControlCmbConfigDeleteResponse = interface{}

type ControlCmbConfigNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	CmbConfig CmbConfigParam      `json:"cmb_config,required"`
}

func (r ControlCmbConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CmbConfig)
}

type ControlCmbConfigNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ControlCmbConfigNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CmbConfig                                  `json:"result,nullable"`
	JSON    controlCmbConfigNewResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigNewResponseEnvelope]
type controlCmbConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigNewResponseEnvelopeSuccess bool

const (
	ControlCmbConfigNewResponseEnvelopeSuccessTrue ControlCmbConfigNewResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlCmbConfigDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ControlCmbConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ControlCmbConfigDeleteResponse                `json:"result,nullable"`
	JSON    controlCmbConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigDeleteResponseEnvelope]
type controlCmbConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigDeleteResponseEnvelopeSuccess bool

const (
	ControlCmbConfigDeleteResponseEnvelopeSuccessTrue ControlCmbConfigDeleteResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlCmbConfigGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ControlCmbConfigGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CmbConfig                                  `json:"result,nullable"`
	JSON    controlCmbConfigGetResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigGetResponseEnvelope]
type controlCmbConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ControlCmbConfigGetResponseEnvelopeSuccess bool

const (
	ControlCmbConfigGetResponseEnvelopeSuccessTrue ControlCmbConfigGetResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
