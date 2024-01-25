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

// AccountPcapOwnershipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountPcapOwnershipService]
// method instead.
type AccountPcapOwnershipService struct {
	Options   []option.RequestOption
	Validates *AccountPcapOwnershipValidateService
}

// NewAccountPcapOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPcapOwnershipService(opts ...option.RequestOption) (r *AccountPcapOwnershipService) {
	r = &AccountPcapOwnershipService{}
	r.Options = opts
	r.Validates = NewAccountPcapOwnershipValidateService(opts...)
	return
}

// Deletes buckets added to the packet captures API.
func (r *AccountPcapOwnershipService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *AccountPcapOwnershipService) MagicPcapCollectionAddBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *AccountPcapOwnershipService) MagicPcapCollectionListPcaPsBucketOwnership(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse struct {
	Errors   []AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseError   `json:"errors"`
	Messages []AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessage `json:"messages"`
	Result   AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseSuccess `json:"success"`
	JSON    accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON    `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse]
type accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseError struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseErrorJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseError]
type accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessage struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessageJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessage]
type accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResult struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                                                   `json:"validated"`
	JSON      accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResult]
type accountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatus string

const (
	AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatusPending AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatus = "pending"
	AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatusSuccess AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatus = "success"
	AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatusFailed  AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseResultStatus = "failed"
)

// Whether the API call was successful
type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseSuccess bool

const (
	AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseSuccessTrue AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseSuccess = true
)

type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse struct {
	Errors     []AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseError    `json:"errors"`
	Messages   []AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessage  `json:"messages"`
	Result     []AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResult   `json:"result,nullable"`
	ResultInfo AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseSuccess `json:"success"`
	JSON    accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON    `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse]
type accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseError struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseErrorJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseError]
type accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessage struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessageJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessage]
type accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResult struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                                            `json:"validated"`
	JSON      accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResult]
type accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatus string

const (
	AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatusPending AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatus = "pending"
	AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatusSuccess AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatus = "success"
	AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatusFailed  AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultStatus = "failed"
)

type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfoJSON `json:"-"`
}

// accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfo]
type accountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseSuccess bool

const (
	AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseSuccessTrue AccountPcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseSuccess = true
)

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
