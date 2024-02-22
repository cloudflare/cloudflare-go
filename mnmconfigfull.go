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

// MNMConfigFullService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMNMConfigFullService] method
// instead.
type MNMConfigFullService struct {
	Options []option.RequestOption
}

// NewMNMConfigFullService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMNMConfigFullService(opts ...option.RequestOption) (r *MNMConfigFullService) {
	r = &MNMConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *MNMConfigFullService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMConfigFullListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMConfigFullListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MNMConfigFullListResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                        `json:"name,required"`
	RouterIPs []string                      `json:"router_ips,required"`
	JSON      mnmConfigFullListResponseJSON `json:"-"`
}

// mnmConfigFullListResponseJSON contains the JSON metadata for the struct
// [MNMConfigFullListResponse]
type mnmConfigFullListResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MNMConfigFullListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigFullListResponseEnvelope struct {
	Errors   []MNMConfigFullListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMConfigFullListResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMConfigFullListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MNMConfigFullListResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmConfigFullListResponseEnvelopeJSON    `json:"-"`
}

// mnmConfigFullListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMConfigFullListResponseEnvelope]
type mnmConfigFullListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigFullListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigFullListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    mnmConfigFullListResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmConfigFullListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMConfigFullListResponseEnvelopeErrors]
type mnmConfigFullListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigFullListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMConfigFullListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    mnmConfigFullListResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmConfigFullListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMConfigFullListResponseEnvelopeMessages]
type mnmConfigFullListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMConfigFullListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMConfigFullListResponseEnvelopeSuccess bool

const (
	MNMConfigFullListResponseEnvelopeSuccessTrue MNMConfigFullListResponseEnvelopeSuccess = true
)
