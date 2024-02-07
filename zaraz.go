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
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZarazWorkflowUpdateResponse struct {
	Errors   []ZarazWorkflowUpdateResponseError   `json:"errors"`
	Messages []ZarazWorkflowUpdateResponseMessage `json:"messages"`
	// Zaraz workflow
	Result ZarazWorkflowUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                            `json:"success"`
	JSON    zarazWorkflowUpdateResponseJSON `json:"-"`
}

// zarazWorkflowUpdateResponseJSON contains the JSON metadata for the struct
// [ZarazWorkflowUpdateResponse]
type zarazWorkflowUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zarazWorkflowUpdateResponseErrorJSON `json:"-"`
}

// zarazWorkflowUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZarazWorkflowUpdateResponseError]
type zarazWorkflowUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zarazWorkflowUpdateResponseMessageJSON `json:"-"`
}

// zarazWorkflowUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZarazWorkflowUpdateResponseMessage]
type zarazWorkflowUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zaraz workflow
type ZarazWorkflowUpdateResponseResult string

const (
	ZarazWorkflowUpdateResponseResultRealtime ZarazWorkflowUpdateResponseResult = "realtime"
	ZarazWorkflowUpdateResponseResultPreview  ZarazWorkflowUpdateResponseResult = "preview"
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
