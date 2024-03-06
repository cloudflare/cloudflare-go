// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneWorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWorkerScriptService] method
// instead.
type ZoneWorkerScriptService struct {
	Options []option.RequestOption
}

// NewZoneWorkerScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWorkerScriptService(opts ...option.RequestOption) (r *ZoneWorkerScriptService) {
	r = &ZoneWorkerScriptService{}
	r.Options = opts
	return
}

// Upload a worker, or a new version of a worker.
func (r *ZoneWorkerScriptService) Update(ctx context.Context, body ZoneWorkerScriptUpdateParams, opts ...option.RequestOption) (res *ZoneWorkerScriptUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneWorkerScriptUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete your Worker. This call has no response body on a successful delete.
func (r *ZoneWorkerScriptService) Delete(ctx context.Context, body ZoneWorkerScriptDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ZoneWorkerScriptService) Get(ctx context.Context, query ZoneWorkerScriptGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [ZoneWorkerScriptUpdateResponseUnknown] or
// [shared.UnionString].
type ZoneWorkerScriptUpdateResponse interface {
	ImplementsZoneWorkerScriptUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneWorkerScriptUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneWorkerScriptUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneWorkerScriptUpdateResponseEnvelope struct {
	Errors   []ZoneWorkerScriptUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneWorkerScriptUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZoneWorkerScriptUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZoneWorkerScriptUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zoneWorkerScriptUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneWorkerScriptUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneWorkerScriptUpdateResponseEnvelope]
type zoneWorkerScriptUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerScriptUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerScriptUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneWorkerScriptUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneWorkerScriptUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneWorkerScriptUpdateResponseEnvelopeErrors]
type zoneWorkerScriptUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerScriptUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneWorkerScriptUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneWorkerScriptUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneWorkerScriptUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneWorkerScriptUpdateResponseEnvelopeMessages]
type zoneWorkerScriptUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWorkerScriptUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneWorkerScriptUpdateResponseEnvelopeSuccess bool

const (
	ZoneWorkerScriptUpdateResponseEnvelopeSuccessTrue ZoneWorkerScriptUpdateResponseEnvelopeSuccess = true
)

type ZoneWorkerScriptDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneWorkerScriptGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
