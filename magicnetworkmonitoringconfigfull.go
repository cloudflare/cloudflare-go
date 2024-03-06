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

// MagicNetworkMonitoringConfigFullService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewMagicNetworkMonitoringConfigFullService] method instead.
type MagicNetworkMonitoringConfigFullService struct {
	Options []option.RequestOption
}

// NewMagicNetworkMonitoringConfigFullService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicNetworkMonitoringConfigFullService(opts ...option.RequestOption) (r *MagicNetworkMonitoringConfigFullService) {
	r = &MagicNetworkMonitoringConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, and rules for account.
func (r *MagicNetworkMonitoringConfigFullService) Get(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigFullGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigFullGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringConfigFullGetResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                          `json:"name,required"`
	RouterIPs []string                                        `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigFullGetResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullGetResponseJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringConfigFullGetResponse]
type magicNetworkMonitoringConfigFullGetResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigFullGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigFullGetResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigFullGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigFullGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigFullGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigFullGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigFullGetResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigFullGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigFullGetResponseEnvelope]
type magicNetworkMonitoringConfigFullGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigFullGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigFullGetResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    magicNetworkMonitoringConfigFullGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigFullGetResponseEnvelopeErrors]
type magicNetworkMonitoringConfigFullGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigFullGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringConfigFullGetResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    magicNetworkMonitoringConfigFullGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [MagicNetworkMonitoringConfigFullGetResponseEnvelopeMessages]
type magicNetworkMonitoringConfigFullGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringConfigFullGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigFullGetResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigFullGetResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigFullGetResponseEnvelopeSuccess = true
)
