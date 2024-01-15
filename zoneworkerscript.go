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

// ZoneWorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWorkerScriptService] method
// instead.
type ZoneWorkerScriptService struct {
	Options  []option.RequestOption
	Bindings *ZoneWorkerScriptBindingService
}

// NewZoneWorkerScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWorkerScriptService(opts ...option.RequestOption) (r *ZoneWorkerScriptService) {
	r = &ZoneWorkerScriptService{}
	r.Options = opts
	r.Bindings = NewZoneWorkerScriptBindingService(opts...)
	return
}

// Fetch raw script content for your worker. Note this is the original script
// content, not JSON encoded.
func (r *ZoneWorkerScriptService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "undefined")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete your Worker. This call has no response body on a successful delete.
func (r *ZoneWorkerScriptService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/workers/script", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Upload a worker, or a new version of a worker.
func (r *ZoneWorkerScriptService) WorkerScriptDeprecatedUploadWorker(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/script", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponse struct {
	Errors   []ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseError   `json:"errors"`
	Messages []ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessage `json:"messages"`
	Result   interface{}                                                         `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseSuccess `json:"success"`
	JSON    zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseJSON    `json:"-"`
}

// zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseJSON contains the JSON
// metadata for the struct
// [ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponse]
type zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseErrorJSON `json:"-"`
}

// zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseError]
type zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessageJSON `json:"-"`
}

// zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessage]
type zoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseSuccess bool

const (
	ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseSuccessTrue ZoneWorkerScriptWorkerScriptDeprecatedUploadWorkerResponseSuccess = true
)
