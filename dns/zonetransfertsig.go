// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ZoneTransferTSIGService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferTSIGService] method instead.
type ZoneTransferTSIGService struct {
	Options []option.RequestOption
}

// NewZoneTransferTSIGService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferTSIGService(opts ...option.RequestOption) (r *ZoneTransferTSIGService) {
	r = &ZoneTransferTSIGService{}
	r.Options = opts
	return
}

// Create TSIG.
func (r *ZoneTransferTSIGService) New(ctx context.Context, params ZoneTransferTSIGNewParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env ZoneTransferTSIGNewResponseEnvelope
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
func (r *ZoneTransferTSIGService) Update(ctx context.Context, tsigID string, params ZoneTransferTSIGUpdateParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env ZoneTransferTSIGUpdateResponseEnvelope
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
func (r *ZoneTransferTSIGService) List(ctx context.Context, query ZoneTransferTSIGListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TSIG], err error) {
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
func (r *ZoneTransferTSIGService) ListAutoPaging(ctx context.Context, query ZoneTransferTSIGListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TSIG] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete TSIG.
func (r *ZoneTransferTSIGService) Delete(ctx context.Context, tsigID string, body ZoneTransferTSIGDeleteParams, opts ...option.RequestOption) (res *ZoneTransferTSIGDeleteResponse, err error) {
	var env ZoneTransferTSIGDeleteResponseEnvelope
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
func (r *ZoneTransferTSIGService) Get(ctx context.Context, tsigID string, query ZoneTransferTSIGGetParams, opts ...option.RequestOption) (res *TSIG, err error) {
	var env ZoneTransferTSIGGetResponseEnvelope
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

type ZoneTransferTSIGDeleteResponse struct {
	ID   string                             `json:"id"`
	JSON zoneTransferTSIGDeleteResponseJSON `json:"-"`
}

// zoneTransferTSIGDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferTSIGDeleteResponse]
type zoneTransferTSIGDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferTSIGDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferTSIGDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferTSIGNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	TSIG      TSIGParam           `json:"tsig,required"`
}

func (r ZoneTransferTSIGNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TSIG)
}

type ZoneTransferTSIGNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferTSIGNewResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                                       `json:"result"`
	JSON    zoneTransferTSIGNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferTSIGNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferTSIGNewResponseEnvelope]
type zoneTransferTSIGNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferTSIGNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferTSIGNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferTSIGNewResponseEnvelopeSuccess bool

const (
	ZoneTransferTSIGNewResponseEnvelopeSuccessTrue ZoneTransferTSIGNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferTSIGNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferTSIGNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferTSIGUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	TSIG      TSIGParam           `json:"tsig,required"`
}

func (r ZoneTransferTSIGUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TSIG)
}

type ZoneTransferTSIGUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferTSIGUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                                          `json:"result"`
	JSON    zoneTransferTSIGUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferTSIGUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferTSIGUpdateResponseEnvelope]
type zoneTransferTSIGUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferTSIGUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferTSIGUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferTSIGUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferTSIGUpdateResponseEnvelopeSuccessTrue ZoneTransferTSIGUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferTSIGUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferTSIGUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferTSIGListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferTSIGDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferTSIGDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferTSIGDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferTSIGDeleteResponse                `json:"result"`
	JSON    zoneTransferTSIGDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferTSIGDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferTSIGDeleteResponseEnvelope]
type zoneTransferTSIGDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferTSIGDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferTSIGDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferTSIGDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferTSIGDeleteResponseEnvelopeSuccessTrue ZoneTransferTSIGDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferTSIGDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferTSIGDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferTSIGGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferTSIGGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferTSIGGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TSIG                                       `json:"result"`
	JSON    zoneTransferTSIGGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferTSIGGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferTSIGGetResponseEnvelope]
type zoneTransferTSIGGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferTSIGGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferTSIGGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferTSIGGetResponseEnvelopeSuccess bool

const (
	ZoneTransferTSIGGetResponseEnvelopeSuccessTrue ZoneTransferTSIGGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferTSIGGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferTSIGGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
