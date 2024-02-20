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
func (r *ZarazService) Update(ctx context.Context, zoneID string, body ZarazUpdateParams, opts ...option.RequestOption) (res *ZarazUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz workflow
type ZarazUpdateResponse string

const (
	ZarazUpdateResponseRealtime ZarazUpdateResponse = "realtime"
	ZarazUpdateResponsePreview  ZarazUpdateResponse = "preview"
)

type ZarazUpdateParams struct {
	// Zaraz workflow
	Body param.Field[ZarazUpdateParamsBody] `json:"body,required"`
}

func (r ZarazUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Zaraz workflow
type ZarazUpdateParamsBody string

const (
	ZarazUpdateParamsBodyRealtime ZarazUpdateParamsBody = "realtime"
	ZarazUpdateParamsBodyPreview  ZarazUpdateParamsBody = "preview"
)

type ZarazUpdateResponseEnvelope struct {
	Errors   []ZarazUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result ZarazUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	JSON    zarazUpdateResponseEnvelopeJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazUpdateResponseEnvelope]
type zarazUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazUpdateResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zarazUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZarazUpdateResponseEnvelopeErrors]
type zarazUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazUpdateResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zarazUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazUpdateResponseEnvelopeMessages]
type zarazUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
