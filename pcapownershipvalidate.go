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

// PcapOwnershipValidateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPcapOwnershipValidateService]
// method instead.
type PcapOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewPcapOwnershipValidateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPcapOwnershipValidateService(opts ...option.RequestOption) (r *PcapOwnershipValidateService) {
	r = &PcapOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates buckets added to the packet captures API.
func (r *PcapOwnershipValidateService) MagicPcapCollectionValidateBucketsForFullPacketCaptures(ctx context.Context, accountIdentifier string, body PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams, opts ...option.RequestOption) (res *PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                                                                                   `json:"validated"`
	JSON      pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON `json:"-"`
}

// pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON
// contains the JSON metadata for the struct
// [PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse]
type pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatus string

const (
	PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatusPending PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatus = "pending"
	PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatusSuccess PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatus = "success"
	PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatusFailed  PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseStatus = "failed"
)

type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelope struct {
	Errors   []PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelope]
type pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrors]
type pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessages]
type pcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeSuccess bool

const (
	PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeSuccessTrue PcapOwnershipValidateMagicPcapCollectionValidateBucketsForFullPacketCapturesResponseEnvelopeSuccess = true
)
