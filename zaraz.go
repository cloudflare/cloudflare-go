// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZarazService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewZarazService] method instead.
type ZarazService struct {
	Options  []option.RequestOption
	Config   *ZarazConfigService
	Default  *ZarazDefaultService
	Export   *ZarazExportService
	History  *ZarazHistoryService
	Publish  *ZarazPublishService
	Workflow *ZarazWorkflowService
}

// NewZarazService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZarazService(opts ...option.RequestOption) (r *ZarazService) {
	r = &ZarazService{}
	r.Options = opts
	r.Config = NewZarazConfigService(opts...)
	r.Default = NewZarazDefaultService(opts...)
	r.Export = NewZarazExportService(opts...)
	r.History = NewZarazHistoryService(opts...)
	r.Publish = NewZarazPublishService(opts...)
	r.Workflow = NewZarazWorkflowService(opts...)
	return
}

// Updates Zaraz workflow for a zone.
func (r *ZarazService) WorkflowUpdate(ctx context.Context, zoneID string, body ZarazWorkflowUpdateParams, opts ...option.RequestOption) (res *ZarazWorkflowUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazWorkflowUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz workflow
type ZarazWorkflowUpdateResponse string

const (
	ZarazWorkflowUpdateResponseRealtime ZarazWorkflowUpdateResponse = "realtime"
	ZarazWorkflowUpdateResponsePreview  ZarazWorkflowUpdateResponse = "preview"
)

type ZarazWorkflowUpdateParams struct {
	// Zaraz workflow
	Body param.Field[ZarazWorkflowUpdateParamsBody] `json:"body,required"`
}

func (r ZarazWorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Zaraz workflow
type ZarazWorkflowUpdateParamsBody string

const (
	ZarazWorkflowUpdateParamsBodyRealtime ZarazWorkflowUpdateParamsBody = "realtime"
	ZarazWorkflowUpdateParamsBodyPreview  ZarazWorkflowUpdateParamsBody = "preview"
)

type ZarazWorkflowUpdateResponseEnvelope struct {
	Errors   []ZarazWorkflowUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazWorkflowUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result ZarazWorkflowUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                                    `json:"success,required"`
	JSON    zarazWorkflowUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazWorkflowUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZarazWorkflowUpdateResponseEnvelope]
type zarazWorkflowUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zarazWorkflowUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazWorkflowUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazWorkflowUpdateResponseEnvelopeErrors]
type zarazWorkflowUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zarazWorkflowUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazWorkflowUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZarazWorkflowUpdateResponseEnvelopeMessages]
type zarazWorkflowUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
