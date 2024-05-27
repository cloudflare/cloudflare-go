// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// R2ConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2ConfigurationService] method instead.
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
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type R2ConfigurationGetResponse map[string]map[string]R2ConfigurationGetResponseItem

type R2ConfigurationGetResponseItem struct {
	// Queue ID that will receive notifications based on the configured rules
	Queue string `json:"queue,required"`
	// Array of rules to drive notifications
	Rules []R2ConfigurationGetResponseItemRule `json:"rules,required"`
	JSON  r2ConfigurationGetResponseItemJSON   `json:"-"`
}

// r2ConfigurationGetResponseItemJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponseItem]
type r2ConfigurationGetResponseItemJSON struct {
	Queue       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseItemJSON) RawJSON() string {
	return r.raw
}

type R2ConfigurationGetResponseItemRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions []R2ConfigurationGetResponseItemRulesAction `json:"actions,required"`
	// Notifications will be sent only for objects with this prefix
	Prefix string `json:"prefix"`
	// Notifications will be sent only for objects with this suffix
	Suffix string                                 `json:"suffix"`
	JSON   r2ConfigurationGetResponseItemRuleJSON `json:"-"`
}

// r2ConfigurationGetResponseItemRuleJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponseItemRule]
type r2ConfigurationGetResponseItemRuleJSON struct {
	Actions     apijson.Field
	Prefix      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseItemRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseItemRuleJSON) RawJSON() string {
	return r.raw
}

type R2ConfigurationGetResponseItemRulesAction string

const (
	R2ConfigurationGetResponseItemRulesActionPutObject               R2ConfigurationGetResponseItemRulesAction = "PutObject"
	R2ConfigurationGetResponseItemRulesActionCopyObject              R2ConfigurationGetResponseItemRulesAction = "CopyObject"
	R2ConfigurationGetResponseItemRulesActionDeleteObject            R2ConfigurationGetResponseItemRulesAction = "DeleteObject"
	R2ConfigurationGetResponseItemRulesActionCompleteMultipartUpload R2ConfigurationGetResponseItemRulesAction = "CompleteMultipartUpload"
	R2ConfigurationGetResponseItemRulesActionAbortMultipartUpload    R2ConfigurationGetResponseItemRulesAction = "AbortMultipartUpload"
)

func (r R2ConfigurationGetResponseItemRulesAction) IsKnown() bool {
	switch r {
	case R2ConfigurationGetResponseItemRulesActionPutObject, R2ConfigurationGetResponseItemRulesActionCopyObject, R2ConfigurationGetResponseItemRulesActionDeleteObject, R2ConfigurationGetResponseItemRulesActionCompleteMultipartUpload, R2ConfigurationGetResponseItemRulesActionAbortMultipartUpload:
		return true
	}
	return false
}

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
