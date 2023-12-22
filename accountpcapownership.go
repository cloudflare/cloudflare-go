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
func (r *AccountPcapOwnershipService) MagicPcapCollectionAddBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *PcapsOwnershipSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all buckets configured for use with PCAPs API.
func (r *AccountPcapOwnershipService) MagicPcapCollectionListPcaPsBucketOwnership(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *PcapsOwnershipCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PcapsOwnershipCollection struct {
	Errors     []PcapsOwnershipCollectionError    `json:"errors"`
	Messages   []PcapsOwnershipCollectionMessage  `json:"messages"`
	Result     []PcapsOwnershipCollectionResult   `json:"result,nullable"`
	ResultInfo PcapsOwnershipCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PcapsOwnershipCollectionSuccess `json:"success"`
	JSON    pcapsOwnershipCollectionJSON    `json:"-"`
}

// pcapsOwnershipCollectionJSON contains the JSON metadata for the struct
// [PcapsOwnershipCollection]
type pcapsOwnershipCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pcapsOwnershipCollectionErrorJSON `json:"-"`
}

// pcapsOwnershipCollectionErrorJSON contains the JSON metadata for the struct
// [PcapsOwnershipCollectionError]
type pcapsOwnershipCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pcapsOwnershipCollectionMessageJSON `json:"-"`
}

// pcapsOwnershipCollectionMessageJSON contains the JSON metadata for the struct
// [PcapsOwnershipCollectionMessage]
type pcapsOwnershipCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipCollectionResult struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapsOwnershipCollectionResultStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                             `json:"validated"`
	JSON      pcapsOwnershipCollectionResultJSON `json:"-"`
}

// pcapsOwnershipCollectionResultJSON contains the JSON metadata for the struct
// [PcapsOwnershipCollectionResult]
type pcapsOwnershipCollectionResultJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapsOwnershipCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapsOwnershipCollectionResultStatus string

const (
	PcapsOwnershipCollectionResultStatusPending PcapsOwnershipCollectionResultStatus = "pending"
	PcapsOwnershipCollectionResultStatusSuccess PcapsOwnershipCollectionResultStatus = "success"
	PcapsOwnershipCollectionResultStatusFailed  PcapsOwnershipCollectionResultStatus = "failed"
)

type PcapsOwnershipCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       pcapsOwnershipCollectionResultInfoJSON `json:"-"`
}

// pcapsOwnershipCollectionResultInfoJSON contains the JSON metadata for the struct
// [PcapsOwnershipCollectionResultInfo]
type pcapsOwnershipCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapsOwnershipCollectionSuccess bool

const (
	PcapsOwnershipCollectionSuccessTrue PcapsOwnershipCollectionSuccess = true
)

type PcapsOwnershipSingleResponse struct {
	Errors   []PcapsOwnershipSingleResponseError   `json:"errors"`
	Messages []PcapsOwnershipSingleResponseMessage `json:"messages"`
	Result   PcapsOwnershipSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PcapsOwnershipSingleResponseSuccess `json:"success"`
	JSON    pcapsOwnershipSingleResponseJSON    `json:"-"`
}

// pcapsOwnershipSingleResponseJSON contains the JSON metadata for the struct
// [PcapsOwnershipSingleResponse]
type pcapsOwnershipSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipSingleResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    pcapsOwnershipSingleResponseErrorJSON `json:"-"`
}

// pcapsOwnershipSingleResponseErrorJSON contains the JSON metadata for the struct
// [PcapsOwnershipSingleResponseError]
type pcapsOwnershipSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipSingleResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    pcapsOwnershipSingleResponseMessageJSON `json:"-"`
}

// pcapsOwnershipSingleResponseMessageJSON contains the JSON metadata for the
// struct [PcapsOwnershipSingleResponseMessage]
type pcapsOwnershipSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapsOwnershipSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapsOwnershipSingleResponseResult struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapsOwnershipSingleResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                 `json:"validated"`
	JSON      pcapsOwnershipSingleResponseResultJSON `json:"-"`
}

// pcapsOwnershipSingleResponseResultJSON contains the JSON metadata for the struct
// [PcapsOwnershipSingleResponseResult]
type pcapsOwnershipSingleResponseResultJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapsOwnershipSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapsOwnershipSingleResponseResultStatus string

const (
	PcapsOwnershipSingleResponseResultStatusPending PcapsOwnershipSingleResponseResultStatus = "pending"
	PcapsOwnershipSingleResponseResultStatusSuccess PcapsOwnershipSingleResponseResultStatus = "success"
	PcapsOwnershipSingleResponseResultStatusFailed  PcapsOwnershipSingleResponseResultStatus = "failed"
)

// Whether the API call was successful
type PcapsOwnershipSingleResponseSuccess bool

const (
	PcapsOwnershipSingleResponseSuccessTrue PcapsOwnershipSingleResponseSuccess = true
)

type AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r AccountPcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
