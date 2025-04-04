// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPPayloadLogService] method instead.
type DLPPayloadLogService struct {
	Options []option.RequestOption
}

// NewDLPPayloadLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPayloadLogService(opts ...option.RequestOption) (r *DLPPayloadLogService) {
	r = &DLPPayloadLogService{}
	r.Options = opts
	return
}

// Set payload log settings
func (r *DLPPayloadLogService) Update(ctx context.Context, params DLPPayloadLogUpdateParams, opts ...option.RequestOption) (res *DLPPayloadLogUpdateResponse, err error) {
	var env DLPPayloadLogUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get payload log settings
func (r *DLPPayloadLogService) Get(ctx context.Context, query DLPPayloadLogGetParams, opts ...option.RequestOption) (res *DLPPayloadLogGetResponse, err error) {
	var env DLPPayloadLogGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DLPPayloadLogUpdateResponse struct {
	UpdatedAt time.Time                       `json:"updated_at,required" format:"date-time"`
	PublicKey string                          `json:"public_key,nullable"`
	JSON      dlpPayloadLogUpdateResponseJSON `json:"-"`
}

// dlpPayloadLogUpdateResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogUpdateResponse]
type dlpPayloadLogUpdateResponseJSON struct {
	UpdatedAt   apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogGetResponse struct {
	UpdatedAt time.Time                    `json:"updated_at,required" format:"date-time"`
	PublicKey string                       `json:"public_key,nullable"`
	JSON      dlpPayloadLogGetResponseJSON `json:"-"`
}

// dlpPayloadLogGetResponseJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponse]
type dlpPayloadLogGetResponseJSON struct {
	UpdatedAt   apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DLPPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPayloadLogUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPPayloadLogUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPPayloadLogUpdateResponse                `json:"result"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPPayloadLogUpdateResponseEnvelope]
type dlpPayloadLogUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPPayloadLogUpdateResponseEnvelopeSuccess bool

const (
	DLPPayloadLogUpdateResponseEnvelopeSuccessTrue DLPPayloadLogUpdateResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPPayloadLogGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPPayloadLogGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DLPPayloadLogGetResponse                `json:"result"`
	JSON    dlpPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponseEnvelope]
type dlpPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPPayloadLogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpPayloadLogGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DLPPayloadLogGetResponseEnvelopeSuccess bool

const (
	DLPPayloadLogGetResponseEnvelopeSuccessTrue DLPPayloadLogGetResponseEnvelopeSuccess = true
)

func (r DLPPayloadLogGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPPayloadLogGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
