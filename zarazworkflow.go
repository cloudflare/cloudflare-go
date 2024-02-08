// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZarazWorkflowService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazWorkflowService] method
// instead.
type ZarazWorkflowService struct {
	Options []option.RequestOption
}

// NewZarazWorkflowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazWorkflowService(opts ...option.RequestOption) (r *ZarazWorkflowService) {
	r = &ZarazWorkflowService{}
	r.Options = opts
	return
}

// Gets Zaraz workflow for a zone.
func (r *ZarazWorkflowService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZarazWorkflowGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazWorkflowGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz workflow
type ZarazWorkflowGetResponse string

const (
	ZarazWorkflowGetResponseRealtime ZarazWorkflowGetResponse = "realtime"
	ZarazWorkflowGetResponsePreview  ZarazWorkflowGetResponse = "preview"
)

type ZarazWorkflowGetResponseEnvelope struct {
	Errors   []ZarazWorkflowGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazWorkflowGetResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result ZarazWorkflowGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                 `json:"success,required"`
	JSON    zarazWorkflowGetResponseEnvelopeJSON `json:"-"`
}

// zarazWorkflowGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazWorkflowGetResponseEnvelope]
type zarazWorkflowGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zarazWorkflowGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazWorkflowGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazWorkflowGetResponseEnvelopeErrors]
type zarazWorkflowGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zarazWorkflowGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazWorkflowGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazWorkflowGetResponseEnvelopeMessages]
type zarazWorkflowGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
