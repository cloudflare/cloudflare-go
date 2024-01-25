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

// AccountPcapOwnershipValidateService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountPcapOwnershipValidateService] method instead.
type AccountPcapOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewAccountPcapOwnershipValidateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPcapOwnershipValidateService(opts ...option.RequestOption) (r *AccountPcapOwnershipValidateService) {
	r = &AccountPcapOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates buckets added to the packet captures API.
func (r *AccountPcapOwnershipValidateService) MagicPcapCollectionValidateBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse struct {
	Errors   []AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseError   `json:"errors"`
	Messages []AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessage `json:"messages"`
	Result   AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseSuccess `json:"success"`
	JSON    accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON    `json:"-"`
}

// accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse]
type accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseErrorJSON `json:"-"`
}

// accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseError]
type accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessageJSON `json:"-"`
}

// accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessage]
type accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResult struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                                                                `json:"validated"`
	JSON      accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultJSON `json:"-"`
}

// accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResult]
type accountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatus string

const (
	AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatusPending AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatus = "pending"
	AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatusSuccess AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatus = "success"
	AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatusFailed  AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseResultStatus = "failed"
)

// Whether the API call was successful
type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseSuccess bool

const (
	AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseSuccessTrue AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseSuccess = true
)

type AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r AccountPcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
