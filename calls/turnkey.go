// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// TurnKeyService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTurnKeyService] method instead.
type TurnKeyService struct {
	Options []option.RequestOption
}

// NewTurnKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTurnKeyService(opts ...option.RequestOption) (r *TurnKeyService) {
	r = &TurnKeyService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare Calls TURN key.
func (r *TurnKeyService) New(ctx context.Context, params TurnKeyNewParams, opts ...option.RequestOption) (res *TurnKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Edit details for a single TURN key.
func (r *TurnKeyService) Update(ctx context.Context, keyID string, params TurnKeyUpdateParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env TurnKeyUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", params.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all TURN keys in the Cloudflare account
func (r *TurnKeyService) List(ctx context.Context, query TurnKeyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/calls/turn_keys", query.AccountID)
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

// Lists all TURN keys in the Cloudflare account
func (r *TurnKeyService) ListAutoPaging(ctx context.Context, query TurnKeyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a TURN key from Cloudflare Calls
func (r *TurnKeyService) Delete(ctx context.Context, keyID string, body TurnKeyDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env TurnKeyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", body.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single TURN key.
func (r *TurnKeyService) Get(ctx context.Context, keyID string, query TurnKeyGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env TurnKeyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/turn_keys/%s", query.AccountID, keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TurnKeyNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// Bearer token
	Key string `json:"key"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of a TURN key, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                 `json:"uid"`
	JSON turnKeyNewResponseJSON `json:"-"`
}

// turnKeyNewResponseJSON contains the JSON metadata for the struct
// [TurnKeyNewResponse]
type turnKeyNewResponseJSON struct {
	Created     apijson.Field
	Key         apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type TurnKeyNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TurnKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TurnKeyUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of a TURN key, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r TurnKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TurnKeyUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TurnKeyUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                            `json:"result"`
	JSON   turnKeyUpdateResponseEnvelopeJSON `json:"-"`
}

// turnKeyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [TurnKeyUpdateResponseEnvelope]
type turnKeyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnKeyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TurnKeyUpdateResponseEnvelopeSuccess bool

const (
	TurnKeyUpdateResponseEnvelopeSuccessTrue TurnKeyUpdateResponseEnvelopeSuccess = true
)

func (r TurnKeyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TurnKeyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TurnKeyListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TurnKeyDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TurnKeyDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TurnKeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                            `json:"result"`
	JSON   turnKeyDeleteResponseEnvelopeJSON `json:"-"`
}

// turnKeyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TurnKeyDeleteResponseEnvelope]
type turnKeyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnKeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TurnKeyDeleteResponseEnvelopeSuccess bool

const (
	TurnKeyDeleteResponseEnvelopeSuccessTrue TurnKeyDeleteResponseEnvelopeSuccess = true
)

func (r TurnKeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TurnKeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TurnKeyGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TurnKeyGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TurnKeyGetResponseEnvelopeSuccess `json:"success,required"`
	// Bearer token
	Result string                         `json:"result"`
	JSON   turnKeyGetResponseEnvelopeJSON `json:"-"`
}

// turnKeyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TurnKeyGetResponseEnvelope]
type turnKeyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TurnKeyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r turnKeyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TurnKeyGetResponseEnvelopeSuccess bool

const (
	TurnKeyGetResponseEnvelopeSuccessTrue TurnKeyGetResponseEnvelopeSuccess = true
)

func (r TurnKeyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TurnKeyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
