// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ContentScanningPayloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentScanningPayloadService] method instead.
type ContentScanningPayloadService struct {
	Options []option.RequestOption
}

// NewContentScanningPayloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContentScanningPayloadService(opts ...option.RequestOption) (r *ContentScanningPayloadService) {
	r = &ContentScanningPayloadService{}
	r.Options = opts
	return
}

// Add custom scan expressions for Content Scanning
func (r *ContentScanningPayloadService) New(ctx context.Context, params ContentScanningPayloadNewParams, opts ...option.RequestOption) (res *[]ContentScanningPayloadNewResponse, err error) {
	var env ContentScanningPayloadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of existing custom scan expressions for Content Scanning
func (r *ContentScanningPayloadService) List(ctx context.Context, query ContentScanningPayloadListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ContentScanningPayloadListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of existing custom scan expressions for Content Scanning
func (r *ContentScanningPayloadService) ListAutoPaging(ctx context.Context, query ContentScanningPayloadListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ContentScanningPayloadListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Content Scan Custom Expression
func (r *ContentScanningPayloadService) Delete(ctx context.Context, expressionID string, body ContentScanningPayloadDeleteParams, opts ...option.RequestOption) (res *[]ContentScanningPayloadDeleteResponse, err error) {
	var env ContentScanningPayloadDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if expressionID == "" {
		err = errors.New("missing required expression_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads/%s", body.ZoneID, expressionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A custom scan expression to match Content Scanning on
type ContentScanningPayloadNewResponse struct {
	// The unique ID for this custom scan expression
	ID string `json:"id"`
	// Ruleset expression to use in matching content objects
	Payload string                                `json:"payload"`
	JSON    contentScanningPayloadNewResponseJSON `json:"-"`
}

// contentScanningPayloadNewResponseJSON contains the JSON metadata for the struct
// [ContentScanningPayloadNewResponse]
type contentScanningPayloadNewResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningPayloadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningPayloadNewResponseJSON) RawJSON() string {
	return r.raw
}

// A custom scan expression to match Content Scanning on
type ContentScanningPayloadListResponse struct {
	// The unique ID for this custom scan expression
	ID string `json:"id"`
	// Ruleset expression to use in matching content objects
	Payload string                                 `json:"payload"`
	JSON    contentScanningPayloadListResponseJSON `json:"-"`
}

// contentScanningPayloadListResponseJSON contains the JSON metadata for the struct
// [ContentScanningPayloadListResponse]
type contentScanningPayloadListResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningPayloadListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningPayloadListResponseJSON) RawJSON() string {
	return r.raw
}

// A custom scan expression to match Content Scanning on
type ContentScanningPayloadDeleteResponse struct {
	// The unique ID for this custom scan expression
	ID string `json:"id"`
	// Ruleset expression to use in matching content objects
	Payload string                                   `json:"payload"`
	JSON    contentScanningPayloadDeleteResponseJSON `json:"-"`
}

// contentScanningPayloadDeleteResponseJSON contains the JSON metadata for the
// struct [ContentScanningPayloadDeleteResponse]
type contentScanningPayloadDeleteResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningPayloadDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningPayloadDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ContentScanningPayloadNewParams struct {
	// Identifier
	ZoneID param.Field[string]                   `path:"zone_id,required"`
	Body   []ContentScanningPayloadNewParamsBody `json:"body,required"`
}

func (r ContentScanningPayloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ContentScanningPayloadNewParamsBody struct {
	// Ruleset expression to use in matching content objects
	Payload param.Field[string] `json:"payload,required"`
}

func (r ContentScanningPayloadNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ContentScanningPayloadNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []shared.ResponseInfo               `json:"messages,required"`
	Result   []ContentScanningPayloadNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ContentScanningPayloadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningPayloadNewResponseEnvelopeJSON    `json:"-"`
}

// contentScanningPayloadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningPayloadNewResponseEnvelope]
type contentScanningPayloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningPayloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningPayloadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ContentScanningPayloadNewResponseEnvelopeSuccess bool

const (
	ContentScanningPayloadNewResponseEnvelopeSuccessTrue ContentScanningPayloadNewResponseEnvelopeSuccess = true
)

func (r ContentScanningPayloadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningPayloadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningPayloadListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningPayloadDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningPayloadDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                  `json:"errors,required"`
	Messages []shared.ResponseInfo                  `json:"messages,required"`
	Result   []ContentScanningPayloadDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ContentScanningPayloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningPayloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// contentScanningPayloadDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ContentScanningPayloadDeleteResponseEnvelope]
type contentScanningPayloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningPayloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningPayloadDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ContentScanningPayloadDeleteResponseEnvelopeSuccess bool

const (
	ContentScanningPayloadDeleteResponseEnvelopeSuccessTrue ContentScanningPayloadDeleteResponseEnvelopeSuccess = true
)

func (r ContentScanningPayloadDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningPayloadDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
