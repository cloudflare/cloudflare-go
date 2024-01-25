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

// ZoneDcvDelegationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDcvDelegationService] method
// instead.
type ZoneDcvDelegationService struct {
	Options []option.RequestOption
}

// NewZoneDcvDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDcvDelegationService(opts ...option.RequestOption) (r *ZoneDcvDelegationService) {
	r = &ZoneDcvDelegationService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *ZoneDcvDelegationService) Uuid(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneDcvDelegationUuidResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneDcvDelegationUuidResponse struct {
	Errors   []ZoneDcvDelegationUuidResponseError   `json:"errors"`
	Messages []ZoneDcvDelegationUuidResponseMessage `json:"messages"`
	Result   ZoneDcvDelegationUuidResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneDcvDelegationUuidResponseSuccess `json:"success"`
	JSON    zoneDcvDelegationUuidResponseJSON    `json:"-"`
}

// zoneDcvDelegationUuidResponseJSON contains the JSON metadata for the struct
// [ZoneDcvDelegationUuidResponse]
type zoneDcvDelegationUuidResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationUuidResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDcvDelegationUuidResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneDcvDelegationUuidResponseErrorJSON `json:"-"`
}

// zoneDcvDelegationUuidResponseErrorJSON contains the JSON metadata for the struct
// [ZoneDcvDelegationUuidResponseError]
type zoneDcvDelegationUuidResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationUuidResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDcvDelegationUuidResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneDcvDelegationUuidResponseMessageJSON `json:"-"`
}

// zoneDcvDelegationUuidResponseMessageJSON contains the JSON metadata for the
// struct [ZoneDcvDelegationUuidResponseMessage]
type zoneDcvDelegationUuidResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationUuidResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDcvDelegationUuidResponseResult struct {
	// The DCV Delegation unique identifier.
	Uuid string                                  `json:"uuid"`
	JSON zoneDcvDelegationUuidResponseResultJSON `json:"-"`
}

// zoneDcvDelegationUuidResponseResultJSON contains the JSON metadata for the
// struct [ZoneDcvDelegationUuidResponseResult]
type zoneDcvDelegationUuidResponseResultJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDcvDelegationUuidResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneDcvDelegationUuidResponseSuccess bool

const (
	ZoneDcvDelegationUuidResponseSuccessTrue ZoneDcvDelegationUuidResponseSuccess = true
)
