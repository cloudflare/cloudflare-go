// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// TSIGService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTSIGService] method instead.
type TSIGService struct {
	Options []option.RequestOption
}

// NewTSIGService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTSIGService(opts ...option.RequestOption) (r *TSIGService) {
	r = &TSIGService{}
	r.Options = opts
	return
}

// Create TSIG.
func (r *TSIGService) New(ctx context.Context, params TSIGNewParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env TSIGNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify TSIG.
func (r *TSIGService) Update(ctx context.Context, tsigID string, params TSIGUpdateParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env TSIGUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", params.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List TSIGs.
func (r *TSIGService) List(ctx context.Context, query TSIGListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TSIG], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List TSIGs.
func (r *TSIGService) ListAutoPaging(ctx context.Context, query TSIGListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TSIG] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete TSIG.
func (r *TSIGService) Delete(ctx context.Context, tsigID string, body TSIGDeleteParams, opts ...option.RequestOption) (res *TSIGDeleteResponse, err error) {
	var env TSIGDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", body.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get TSIG.
func (r *TSIGService) Get(ctx context.Context, tsigID string, query TSIGGetParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env TSIGGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tsigID == "" {
		err = errors.New("missing required tsig_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/tsigs/%s", query.AccountID, tsigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TSIG struct {
	ID string `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string   `json:"secret,required"`
	JSON   tsigJSON `json:"-"`
}

// tsigJSON contains the JSON metadata for the struct [TSIG]
type tsigJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIG) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigJSON) RawJSON() string {
	return r.raw
}

type TSIGParam struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r TSIGParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TSIGDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON tsigDeleteResponseJSON `json:"-"`
}

// tsigDeleteResponseJSON contains the JSON metadata for the struct
// [TSIGDeleteResponse]
type tsigDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TSIGNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	TSIG      TSIGParam           `json:"tsig,required"`
}

func (r TSIGNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TSIG)
}

type TSIGNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TSIGNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                           `json:"result"`
	JSON    tsigNewResponseEnvelopeJSON    `json:"-"`
}

// tsigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGNewResponseEnvelope]
type tsigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGNewResponseEnvelopeSuccess bool

const (
	TSIGNewResponseEnvelopeSuccessTrue TSIGNewResponseEnvelopeSuccess = true
)

func (r TSIGNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	TSIG      TSIGParam           `json:"tsig,required"`
}

func (r TSIGUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TSIG)
}

type TSIGUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TSIGUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                              `json:"result"`
	JSON    tsigUpdateResponseEnvelopeJSON    `json:"-"`
}

// tsigUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGUpdateResponseEnvelope]
type tsigUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGUpdateResponseEnvelopeSuccess bool

const (
	TSIGUpdateResponseEnvelopeSuccessTrue TSIGUpdateResponseEnvelopeSuccess = true
)

func (r TSIGUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TSIGDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIGDeleteResponse                `json:"result"`
	JSON    tsigDeleteResponseEnvelopeJSON    `json:"-"`
}

// tsigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGDeleteResponseEnvelope]
type tsigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGDeleteResponseEnvelopeSuccess bool

const (
	TSIGDeleteResponseEnvelopeSuccessTrue TSIGDeleteResponseEnvelopeSuccess = true
)

func (r TSIGDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TSIGGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type TSIGGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TSIGGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                           `json:"result"`
	JSON    tsigGetResponseEnvelopeJSON    `json:"-"`
}

// tsigGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TSIGGetResponseEnvelope]
type tsigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TSIGGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tsigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TSIGGetResponseEnvelopeSuccess bool

const (
	TSIGGetResponseEnvelopeSuccessTrue TSIGGetResponseEnvelopeSuccess = true
)

func (r TSIGGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TSIGGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
