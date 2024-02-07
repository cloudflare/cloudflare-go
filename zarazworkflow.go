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
	path := fmt.Sprintf("zones/%s/settings/zaraz/workflow", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZarazWorkflowGetResponse struct {
	Errors   []ZarazWorkflowGetResponseError   `json:"errors"`
	Messages []ZarazWorkflowGetResponseMessage `json:"messages"`
	// Zaraz workflow
	Result ZarazWorkflowGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                         `json:"success"`
	JSON    zarazWorkflowGetResponseJSON `json:"-"`
}

// zarazWorkflowGetResponseJSON contains the JSON metadata for the struct
// [ZarazWorkflowGetResponse]
type zarazWorkflowGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zarazWorkflowGetResponseErrorJSON `json:"-"`
}

// zarazWorkflowGetResponseErrorJSON contains the JSON metadata for the struct
// [ZarazWorkflowGetResponseError]
type zarazWorkflowGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazWorkflowGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zarazWorkflowGetResponseMessageJSON `json:"-"`
}

// zarazWorkflowGetResponseMessageJSON contains the JSON metadata for the struct
// [ZarazWorkflowGetResponseMessage]
type zarazWorkflowGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazWorkflowGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zaraz workflow
type ZarazWorkflowGetResponseResult string

const (
	ZarazWorkflowGetResponseResultRealtime ZarazWorkflowGetResponseResult = "realtime"
	ZarazWorkflowGetResponseResultPreview  ZarazWorkflowGetResponseResult = "preview"
)
