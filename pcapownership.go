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
	Options []option.RequestOption
}

// NewPcapOwnershipService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPcapOwnershipService(opts ...option.RequestOption) (r *PcapOwnershipService) {
	r = &PcapOwnershipService{}
	r.Options = opts
	return
}

// Adds an AWS or GCP bucket to use with full packet captures.
func (r *PcapOwnershipService) New(ctx context.Context, accountIdentifier string, body PcapOwnershipNewParams, opts ...option.RequestOption) (res *PcapOwnershipNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// List all buckets configured for use with PCAPs API.
func (r *PcapOwnershipService) Get(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]PcapOwnershipGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates buckets added to the packet captures API.
func (r *PcapOwnershipService) Validate(ctx context.Context, accountIdentifier string, body PcapOwnershipValidateParams, opts ...option.RequestOption) (res *PcapOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PcapOwnershipValidateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/pcaps/ownership/validate", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PcapOwnershipNewResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipNewResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                       `json:"validated"`
	JSON      pcapOwnershipNewResponseJSON `json:"-"`
}

// pcapOwnershipNewResponseJSON contains the JSON metadata for the struct
// [PcapOwnershipNewResponse]
type pcapOwnershipNewResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipNewResponseStatus string

const (
	PcapOwnershipNewResponseStatusPending PcapOwnershipNewResponseStatus = "pending"
	PcapOwnershipNewResponseStatusSuccess PcapOwnershipNewResponseStatus = "success"
	PcapOwnershipNewResponseStatusFailed  PcapOwnershipNewResponseStatus = "failed"
)

type PcapOwnershipGetResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipGetResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                       `json:"validated"`
	JSON      pcapOwnershipGetResponseJSON `json:"-"`
}

// pcapOwnershipGetResponseJSON contains the JSON metadata for the struct
// [PcapOwnershipGetResponse]
type pcapOwnershipGetResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipGetResponseStatus string

const (
	PcapOwnershipGetResponseStatusPending PcapOwnershipGetResponseStatus = "pending"
	PcapOwnershipGetResponseStatusSuccess PcapOwnershipGetResponseStatus = "success"
	PcapOwnershipGetResponseStatusFailed  PcapOwnershipGetResponseStatus = "failed"
)

type PcapOwnershipValidateResponse struct {
	// The bucket ID associated with the packet captures API.
	ID string `json:"id,required"`
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf string `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	Filename string `json:"filename,required"`
	// The status of the ownership challenge. Can be pending, success or failed.
	Status PcapOwnershipValidateResponseStatus `json:"status,required"`
	// The RFC 3339 timestamp when the bucket was added to packet captures API.
	Submitted string `json:"submitted,required"`
	// The RFC 3339 timestamp when the bucket was validated.
	Validated string                            `json:"validated"`
	JSON      pcapOwnershipValidateResponseJSON `json:"-"`
}

// pcapOwnershipValidateResponseJSON contains the JSON metadata for the struct
// [PcapOwnershipValidateResponse]
type pcapOwnershipValidateResponseJSON struct {
	ID              apijson.Field
	DestinationConf apijson.Field
	Filename        apijson.Field
	Status          apijson.Field
	Submitted       apijson.Field
	Validated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PcapOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the ownership challenge. Can be pending, success or failed.
type PcapOwnershipValidateResponseStatus string

const (
	PcapOwnershipValidateResponseStatusPending PcapOwnershipValidateResponseStatus = "pending"
	PcapOwnershipValidateResponseStatusSuccess PcapOwnershipValidateResponseStatus = "success"
	PcapOwnershipValidateResponseStatusFailed  PcapOwnershipValidateResponseStatus = "failed"
)

type PcapOwnershipNewParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
}

func (r PcapOwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapOwnershipNewResponseEnvelope struct {
	Errors   []PcapOwnershipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapOwnershipNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapOwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipNewResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PcapOwnershipNewResponseEnvelope]
type pcapOwnershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pcapOwnershipNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PcapOwnershipNewResponseEnvelopeErrors]
type pcapOwnershipNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    pcapOwnershipNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PcapOwnershipNewResponseEnvelopeMessages]
type pcapOwnershipNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipNewResponseEnvelopeSuccess bool

const (
	PcapOwnershipNewResponseEnvelopeSuccessTrue PcapOwnershipNewResponseEnvelopeSuccess = true
)

type PcapOwnershipGetResponseEnvelope struct {
	Errors   []PcapOwnershipGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []PcapOwnershipGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PcapOwnershipGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PcapOwnershipGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pcapOwnershipGetResponseEnvelopeJSON       `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PcapOwnershipGetResponseEnvelope]
type pcapOwnershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pcapOwnershipGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PcapOwnershipGetResponseEnvelopeErrors]
type pcapOwnershipGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    pcapOwnershipGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PcapOwnershipGetResponseEnvelopeMessages]
type pcapOwnershipGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipGetResponseEnvelopeSuccess bool

const (
	PcapOwnershipGetResponseEnvelopeSuccessTrue PcapOwnershipGetResponseEnvelopeSuccess = true
)

type PcapOwnershipGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       pcapOwnershipGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// pcapOwnershipGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PcapOwnershipGetResponseEnvelopeResultInfo]
type pcapOwnershipGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipValidateParams struct {
	// The full URI for the bucket. This field only applies to `full` packet captures.
	DestinationConf param.Field[string] `json:"destination_conf,required"`
	// The ownership challenge filename stored in the bucket.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r PcapOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PcapOwnershipValidateResponseEnvelope struct {
	Errors   []PcapOwnershipValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PcapOwnershipValidateResponseEnvelopeMessages `json:"messages,required"`
	Result   PcapOwnershipValidateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PcapOwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeJSON    `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeJSON contains the JSON metadata for the
// struct [PcapOwnershipValidateResponseEnvelope]
type pcapOwnershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipValidateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeErrorsJSON `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PcapOwnershipValidateResponseEnvelopeErrors]
type pcapOwnershipValidateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PcapOwnershipValidateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    pcapOwnershipValidateResponseEnvelopeMessagesJSON `json:"-"`
}

// pcapOwnershipValidateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PcapOwnershipValidateResponseEnvelopeMessages]
type pcapOwnershipValidateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PcapOwnershipValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PcapOwnershipValidateResponseEnvelopeSuccess bool

const (
	PcapOwnershipValidateResponseEnvelopeSuccessTrue PcapOwnershipValidateResponseEnvelopeSuccess = true
)
