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

// ActivationCheckService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewActivationCheckService] method
// instead.
type ActivationCheckService struct {
	Options []option.RequestOption
}

// NewActivationCheckService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActivationCheckService(opts ...option.RequestOption) (r *ActivationCheckService) {
	r = &ActivationCheckService{}
	r.Options = opts
	return
}

// Triggeres a new activation check for a PENDING Zone. This can be triggered every
// 5 min for paygo/ent customers, every hour for FREE Zones.
func (r *ActivationCheckService) Replace(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ActivationCheckReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ActivationCheckReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/activation_check", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ActivationCheckReplaceResponse struct {
	// Identifier
	ID   string                             `json:"id"`
	JSON activationCheckReplaceResponseJSON `json:"-"`
}

// activationCheckReplaceResponseJSON contains the JSON metadata for the struct
// [ActivationCheckReplaceResponse]
type activationCheckReplaceResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckReplaceResponseEnvelope struct {
	Errors   []ActivationCheckReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ActivationCheckReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   ActivationCheckReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ActivationCheckReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    activationCheckReplaceResponseEnvelopeJSON    `json:"-"`
}

// activationCheckReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [ActivationCheckReplaceResponseEnvelope]
type activationCheckReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckReplaceResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    activationCheckReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// activationCheckReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ActivationCheckReplaceResponseEnvelopeErrors]
type activationCheckReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActivationCheckReplaceResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    activationCheckReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// activationCheckReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ActivationCheckReplaceResponseEnvelopeMessages]
type activationCheckReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivationCheckReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ActivationCheckReplaceResponseEnvelopeSuccess bool

const (
	ActivationCheckReplaceResponseEnvelopeSuccessTrue ActivationCheckReplaceResponseEnvelopeSuccess = true
)
