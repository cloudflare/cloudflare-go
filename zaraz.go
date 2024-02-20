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
func (r *ZarazService) Replace(ctx context.Context, zoneID string, body ZarazReplaceParams, opts ...option.RequestOption) (res *ZarazReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zaraz workflow
type ZarazReplaceResponse string

const (
	ZarazReplaceResponseRealtime ZarazReplaceResponse = "realtime"
	ZarazReplaceResponsePreview  ZarazReplaceResponse = "preview"
)

type ZarazReplaceParams struct {
	// Zaraz workflow
	Body param.Field[ZarazReplaceParamsBody] `json:"body,required"`
}

func (r ZarazReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Zaraz workflow
type ZarazReplaceParamsBody string

const (
	ZarazReplaceParamsBodyRealtime ZarazReplaceParamsBody = "realtime"
	ZarazReplaceParamsBodyPreview  ZarazReplaceParamsBody = "preview"
)

type ZarazReplaceResponseEnvelope struct {
	Errors   []ZarazReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazReplaceResponseEnvelopeMessages `json:"messages,required"`
	// Zaraz workflow
	Result ZarazReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success bool                             `json:"success,required"`
	JSON    zarazReplaceResponseEnvelopeJSON `json:"-"`
}

// zarazReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazReplaceResponseEnvelope]
type zarazReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazReplaceResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zarazReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ZarazReplaceResponseEnvelopeErrors]
type zarazReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazReplaceResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zarazReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazReplaceResponseEnvelopeMessages]
type zarazReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
