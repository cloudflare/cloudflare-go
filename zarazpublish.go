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
func (r *ZarazPublishService) New(ctx context.Context, zoneID string, body ZarazPublishNewParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env ZarazPublishNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/zaraz/publish", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZarazPublishNewParams struct {
	// Zaraz configuration description.
	Body param.Field[string] `json:"body,required"`
}

func (r ZarazPublishNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZarazPublishNewResponseEnvelope struct {
	Errors   []ZarazPublishNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZarazPublishNewResponseEnvelopeMessages `json:"messages,required"`
	Result   string                                    `json:"result,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	JSON    zarazPublishNewResponseEnvelopeJSON `json:"-"`
}

// zarazPublishNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZarazPublishNewResponseEnvelope]
type zarazPublishNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazPublishNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zarazPublishNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zarazPublishNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZarazPublishNewResponseEnvelopeErrors]
type zarazPublishNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZarazPublishNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zarazPublishNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zarazPublishNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZarazPublishNewResponseEnvelopeMessages]
type zarazPublishNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZarazPublishNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
