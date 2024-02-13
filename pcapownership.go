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

// PcapOwnershipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPcapOwnershipService] method
// instead.
type PcapOwnershipService struct {
	Options   []option.RequestOption
	Validates *PcapOwnershipValidateService
}

// NewPcapOwnershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPcapOwnershipService(opts ...option.RequestOption) (r *PcapOwnershipService) {
	r = &PcapOwnershipService{}
	r.Options = opts
	r.Validates = NewPcapOwnershipValidateService(opts...)
	return
}

// Deletes buckets added to the packet captures API.
func (r *PcapOwnershipService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *PcapOwnershipService) MagicPcapCollectionAddBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all buckets configured for use with PCAPs API.
func (r *PcapOwnershipService) MagicPcapCollectionListPcaPsBucketOwnership(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                                      `json:"validated"`
	JSON      pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse]
type pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatus string

const (
	PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatusPending PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatus = "pending"
	PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatusSuccess PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatus = "success"
	PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatusFailed  PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseStatus = "failed"
)

type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                               `json:"validated"`
	JSON      pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON contains
// the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse]
type pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatus string

const (
	PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatusPending PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatus = "pending"
	PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatusSuccess PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatus = "success"
	PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatusFailed  PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseStatus = "failed"
)

type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelope struct {
	Errors   []PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelope]
type pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrors struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrors]
type pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessages struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessages]
type pcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeSuccess bool

const (
	PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeSuccessTrue PcapOwnershipMagicPcapCollectionAddBucketsForFullPacketCapturesResponseEnvelopeSuccess = true
)

type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelope struct {
	Errors   []PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessages `json:"messages,required"`
	Result   []PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeJSON       `json:"-"`
}

// pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelope]
type pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrors]
type pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessages]
type pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeSuccess bool

const (
	PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeSuccessTrue PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeSuccess = true
)

type PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                `json:"total_count"`
	JSON       pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfo]
type pcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipMagicPcapCollectionListPcaPsBucketOwnershipResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
