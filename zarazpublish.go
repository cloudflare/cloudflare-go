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

// ZarazPublishService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZarazPublishService] method
// instead.
type ZarazPublishService struct {
	Options []option.RequestOption
}

// NewZarazPublishService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZarazPublishService(opts ...option.RequestOption) (r *ZarazPublishService) {
	r = &ZarazPublishService{}
	r.Options = opts
	return
}

// Publish current Zaraz preview configuration for a zone.
func (r *ZarazPublishService) New(ctx context.Context, zoneID string, body ZarazPublishNewParams, opts ...option.RequestOption) (res *ZarazPublishNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/zaraz/publish", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZarazPublishNewResponse struct {
	Errors   []ZarazPublishNewResponseError   `json:"errors"`
	Messages []ZarazPublishNewResponseMessage `json:"messages"`
	Result   string                           `json:"result"`
	// Whether the API call was successful
	Success bool                        `json:"success"`
	JSON    zarazPublishNewResponseJSON `json:"-"`
}

// zarazPublishNewResponseJSON contains the JSON metadata for the struct
// [ZarazPublishNewResponse]
type zarazPublishNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazPublishNewResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zarazPublishNewResponseErrorJSON `json:"-"`
}

// zarazPublishNewResponseErrorJSON contains the JSON metadata for the struct
// [ZarazPublishNewResponseError]
type zarazPublishNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazPublishNewResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zarazPublishNewResponseMessageJSON `json:"-"`
}

// zarazPublishNewResponseMessageJSON contains the JSON metadata for the struct
// [ZarazPublishNewResponseMessage]
type zarazPublishNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazPublishNewParams struct {
	// Zaraz configuration description.
	Body param.Field[string] `json:"body,required"`
}

func (r ZarazPublishNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
