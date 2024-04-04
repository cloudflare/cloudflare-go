// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DLPPayloadLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDLPPayloadLogService] method
// instead.
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

// Updates the DLP payload log settings for this account.
func (r *DLPPayloadLogService) Update(ctx context.Context, params DLPPayloadLogUpdateParams, opts ...option.RequestOption) (res *UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the current DLP payload log settings for this account.
func (r *DLPPayloadLogService) Get(ctx context.Context, query DLPPayloadLogGetParams, opts ...option.RequestOption) (res *UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5, err error) {
	opts = append(r.Options[:], opts...)
	var env DLPPayloadLogGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/payload_log", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5 struct {
	PublicKey string                                               `json:"public_key,required,nullable"`
	JSON      unnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5JSON `json:"-"`
}

// unnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5]
type unnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5JSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5JSON) RawJSON() string {
	return r.raw
}

type DLPPayloadLogUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The public key to use when encrypting extracted payloads, as a base64 string
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r DLPPayloadLogUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPPayloadLogUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DLPPayloadLogUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPPayloadLogUpdateResponseEnvelope]
type dlpPayloadLogUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DLPPayloadLogGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   UnnamedSchemaRefE31ff4936b1b42746e8cb62bbc87f2e5          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DLPPayloadLogGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dlpPayloadLogGetResponseEnvelopeJSON    `json:"-"`
}

// dlpPayloadLogGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPPayloadLogGetResponseEnvelope]
type dlpPayloadLogGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
