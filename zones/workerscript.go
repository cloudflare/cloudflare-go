// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// WorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptService] method
// instead.
type WorkerScriptService struct {
	Options []option.RequestOption
}

// NewWorkerScriptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerScriptService(opts ...option.RequestOption) (r *WorkerScriptService) {
	r = &WorkerScriptService{}
	r.Options = opts
	return
}

// Upload a worker, or a new version of a worker.
func (r *WorkerScriptService) Update(ctx context.Context, body WorkerScriptUpdateParams, opts ...option.RequestOption) (res *WorkerScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete your Worker. This call has no response body on a successful delete.
func (r *WorkerScriptService) Delete(ctx context.Context, body WorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *WorkerScriptService) Get(ctx context.Context, query WorkerScriptGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by [zones.WorkerScriptUpdateResponseUnknown] or
// [shared.UnionString].
type WorkerScriptUpdateResponse interface {
	ImplementsZonesWorkerScriptUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerScriptUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WorkerScriptUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerScriptUpdateResponseEnvelope struct {
	Errors   []WorkerScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerScriptUpdateResponseEnvelope]
type workerScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerScriptUpdateResponseEnvelopeErrors]
type workerScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerScriptUpdateResponseEnvelopeMessages]
type workerScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptUpdateResponseEnvelopeSuccessTrue WorkerScriptUpdateResponseEnvelopeSuccess = true
)

func (r WorkerScriptUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkerScriptUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkerScriptDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerScriptGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
