// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// R2ConfigurationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewR2ConfigurationService] method
// instead.
type R2ConfigurationService struct {
	Options []option.RequestOption
	Queues  *R2ConfigurationQueueService
}

// NewR2ConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewR2ConfigurationService(opts ...option.RequestOption) (r *R2ConfigurationService) {
	r = &R2ConfigurationService{}
	r.Options = opts
	r.Queues = NewR2ConfigurationQueueService(opts...)
	return
}

// Returns all notification rules for each queue for which bucket notifications are
// produced.
func (r *R2ConfigurationService) Get(ctx context.Context, bucketName string, query R2ConfigurationGetParams, opts ...option.RequestOption) (res *R2ConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env R2ConfigurationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type R2ConfigurationGetResponse map[string]map[string]R2ConfigurationGetResponse

type R2ConfigurationGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2ConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   R2ConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success R2ConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2ConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// r2ConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponseEnvelope]
type r2ConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type R2ConfigurationGetResponseEnvelopeSuccess bool

const (
	R2ConfigurationGetResponseEnvelopeSuccessTrue R2ConfigurationGetResponseEnvelopeSuccess = true
)

func (r R2ConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case R2ConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
