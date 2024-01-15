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

// ZoneZarazWorkflowService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneZarazWorkflowService] method
// instead.
type ZoneZarazWorkflowService struct {
	Options []option.RequestOption
}

// NewZoneZarazWorkflowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneZarazWorkflowService(opts ...option.RequestOption) (r *ZoneZarazWorkflowService) {
	r = &ZoneZarazWorkflowService{}
	r.Options = opts
	return
}

// Gets Zaraz workflow for a zone.
func (r *ZoneZarazWorkflowService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZarazWorkflow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/workflow", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Zaraz workflow for a zone.
func (r *ZoneZarazWorkflowService) Update(ctx context.Context, zoneIdentifier string, body ZoneZarazWorkflowUpdateParams, opts ...option.RequestOption) (res *ZarazWorkflow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/v2/workflow", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZarazWorkflow struct {
	Errors   []ZarazWorkflowError   `json:"errors"`
	Messages []ZarazWorkflowMessage `json:"messages"`
	// Zaraz workflow
	Result ZarazWorkflowResult `json:"result"`
	// Whether the API call was successful
	Success bool              `json:"success"`
	JSON    zarazWorkflowJSON `json:"-"`
}

// zarazWorkflowJSON contains the JSON metadata for the struct [ZarazWorkflow]
type zarazWorkflowJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowError struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    zarazWorkflowErrorJSON `json:"-"`
}

// zarazWorkflowErrorJSON contains the JSON metadata for the struct
// [ZarazWorkflowError]
type zarazWorkflowErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zarazWorkflowMessageJSON `json:"-"`
}

// zarazWorkflowMessageJSON contains the JSON metadata for the struct
// [ZarazWorkflowMessage]
type zarazWorkflowMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zaraz workflow
type ZarazWorkflowResult string

const (
	ZarazWorkflowResultRealtime ZarazWorkflowResult = "realtime"
	ZarazWorkflowResultPreview  ZarazWorkflowResult = "preview"
)

type ZoneZarazWorkflowUpdateParams struct {
	// Zaraz workflow
	Body param.Field[ZoneZarazWorkflowUpdateParamsBody] `json:"body,required"`
}

func (r ZoneZarazWorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Zaraz workflow
type ZoneZarazWorkflowUpdateParamsBody string

const (
	ZoneZarazWorkflowUpdateParamsBodyRealtime ZoneZarazWorkflowUpdateParamsBody = "realtime"
	ZoneZarazWorkflowUpdateParamsBodyPreview  ZoneZarazWorkflowUpdateParamsBody = "preview"
)
