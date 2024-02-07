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

// DcvDelegationUuidService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDcvDelegationUuidService] method
// instead.
type DcvDelegationUuidService struct {
	Options []option.RequestOption
}

// NewDcvDelegationUuidService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDcvDelegationUuidService(opts ...option.RequestOption) (r *DcvDelegationUuidService) {
	r = &DcvDelegationUuidService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *DcvDelegationUuidService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *DcvDelegationUuidGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DcvDelegationUuidGetResponse struct {
	Errors   []DcvDelegationUuidGetResponseError   `json:"errors"`
	Messages []DcvDelegationUuidGetResponseMessage `json:"messages"`
	Result   DcvDelegationUuidGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DcvDelegationUuidGetResponseSuccess `json:"success"`
	JSON    dcvDelegationUuidGetResponseJSON    `json:"-"`
}

// dcvDelegationUuidGetResponseJSON contains the JSON metadata for the struct
// [DcvDelegationUuidGetResponse]
type dcvDelegationUuidGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    dcvDelegationUuidGetResponseErrorJSON `json:"-"`
}

// dcvDelegationUuidGetResponseErrorJSON contains the JSON metadata for the struct
// [DcvDelegationUuidGetResponseError]
type dcvDelegationUuidGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dcvDelegationUuidGetResponseMessageJSON `json:"-"`
}

// dcvDelegationUuidGetResponseMessageJSON contains the JSON metadata for the
// struct [DcvDelegationUuidGetResponseMessage]
type dcvDelegationUuidGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DcvDelegationUuidGetResponseResult struct {
	// The DCV Delegation unique identifier.
	Uuid string                                 `json:"uuid"`
	JSON dcvDelegationUuidGetResponseResultJSON `json:"-"`
}

// dcvDelegationUuidGetResponseResultJSON contains the JSON metadata for the struct
// [DcvDelegationUuidGetResponseResult]
type dcvDelegationUuidGetResponseResultJSON struct {
	Uuid        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DcvDelegationUuidGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DcvDelegationUuidGetResponseSuccess bool

const (
	DcvDelegationUuidGetResponseSuccessTrue DcvDelegationUuidGetResponseSuccess = true
)
