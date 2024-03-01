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
func (r *MagicNetworkMonitoringConfigFullService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringConfigFullListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringConfigFullListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/config/full", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringConfigFullListResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name      string                                           `json:"name,required"`
	RouterIPs []string                                         `json:"router_ips,required"`
	JSON      magicNetworkMonitoringConfigFullListResponseJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullListResponseJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringConfigFullListResponse]
type magicNetworkMonitoringConfigFullListResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringConfigFullListResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringConfigFullListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringConfigFullListResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringConfigFullListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringConfigFullListResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringConfigFullListResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringConfigFullListResponseEnvelopeJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringConfigFullListResponseEnvelope]
type magicNetworkMonitoringConfigFullListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringConfigFullListResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    magicNetworkMonitoringConfigFullListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringConfigFullListResponseEnvelopeErrors]
type magicNetworkMonitoringConfigFullListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MagicNetworkMonitoringConfigFullListResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    magicNetworkMonitoringConfigFullListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringConfigFullListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [MagicNetworkMonitoringConfigFullListResponseEnvelopeMessages]
type magicNetworkMonitoringConfigFullListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringConfigFullListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MagicNetworkMonitoringConfigFullListResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringConfigFullListResponseEnvelopeSuccessTrue MagicNetworkMonitoringConfigFullListResponseEnvelopeSuccess = true
)
