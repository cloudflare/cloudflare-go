// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// ContentScanningService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentScanningService] method instead.
type ContentScanningService struct {
	Options  []option.RequestOption
	Payloads *PayloadService
	Settings *SettingService
}

// NewContentScanningService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContentScanningService(opts ...option.RequestOption) (r *ContentScanningService) {
	r = &ContentScanningService{}
	r.Options = opts
	r.Payloads = NewPayloadService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Update the Content Scanning status.
func (r *ContentScanningService) New(ctx context.Context, params ContentScanningNewParams, opts ...option.RequestOption) (res *ContentScanningNewResponse, err error) {
	var env ContentScanningNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the Content Scanning status.
func (r *ContentScanningService) Update(ctx context.Context, params ContentScanningUpdateParams, opts ...option.RequestOption) (res *ContentScanningUpdateResponse, err error) {
	var env ContentScanningUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable Content Scanning.
func (r *ContentScanningService) Disable(ctx context.Context, body ContentScanningDisableParams, opts ...option.RequestOption) (res *ContentScanningDisableResponse, err error) {
	var env ContentScanningDisableResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/disable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Content Scanning.
func (r *ContentScanningService) Enable(ctx context.Context, body ContentScanningEnableParams, opts ...option.RequestOption) (res *ContentScanningEnableResponse, err error) {
	var env ContentScanningEnableResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the current status of Content Scanning.
func (r *ContentScanningService) Get(ctx context.Context, query ContentScanningGetParams, opts ...option.RequestOption) (res *ContentScanningGetResponse, err error) {
	var env ContentScanningGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Defines the status for Content Scanning.
type ContentScanningNewResponse struct {
	// Defines the last modification date (ISO 8601) of the Content Scanning status.
	Modified string `json:"modified"`
	// Defines the status of Content Scanning.
	Value string                         `json:"value"`
	JSON  contentScanningNewResponseJSON `json:"-"`
}

// contentScanningNewResponseJSON contains the JSON metadata for the struct
// [ContentScanningNewResponse]
type contentScanningNewResponseJSON struct {
	Modified    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningNewResponseJSON) RawJSON() string {
	return r.raw
}

// Defines the status for Content Scanning.
type ContentScanningUpdateResponse struct {
	// Defines the last modification date (ISO 8601) of the Content Scanning status.
	Modified string `json:"modified"`
	// Defines the status of Content Scanning.
	Value string                            `json:"value"`
	JSON  contentScanningUpdateResponseJSON `json:"-"`
}

// contentScanningUpdateResponseJSON contains the JSON metadata for the struct
// [ContentScanningUpdateResponse]
type contentScanningUpdateResponseJSON struct {
	Modified    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ContentScanningDisableResponse = interface{}

type ContentScanningEnableResponse = interface{}

// Defines the status for Content Scanning.
type ContentScanningGetResponse struct {
	// Defines the last modification date (ISO 8601) of the Content Scanning status.
	Modified string `json:"modified"`
	// Defines the status of Content Scanning.
	Value string                         `json:"value"`
	JSON  contentScanningGetResponseJSON `json:"-"`
}

// contentScanningGetResponseJSON contains the JSON metadata for the struct
// [ContentScanningGetResponse]
type contentScanningGetResponseJSON struct {
	Modified    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningGetResponseJSON) RawJSON() string {
	return r.raw
}

type ContentScanningNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The status value for Content Scanning.
	Value param.Field[ContentScanningNewParamsValue] `json:"value,required"`
}

func (r ContentScanningNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status value for Content Scanning.
type ContentScanningNewParamsValue string

const (
	ContentScanningNewParamsValueEnabled  ContentScanningNewParamsValue = "enabled"
	ContentScanningNewParamsValueDisabled ContentScanningNewParamsValue = "disabled"
)

func (r ContentScanningNewParamsValue) IsKnown() bool {
	switch r {
	case ContentScanningNewParamsValueEnabled, ContentScanningNewParamsValueDisabled:
		return true
	}
	return false
}

type ContentScanningNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Defines the status for Content Scanning.
	Result ContentScanningNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningNewResponseEnvelopeJSON    `json:"-"`
}

// contentScanningNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ContentScanningNewResponseEnvelope]
type contentScanningNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningNewResponseEnvelopeSuccess bool

const (
	ContentScanningNewResponseEnvelopeSuccessTrue ContentScanningNewResponseEnvelopeSuccess = true
)

func (r ContentScanningNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The status value for Content Scanning.
	Value param.Field[ContentScanningUpdateParamsValue] `json:"value,required"`
}

func (r ContentScanningUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status value for Content Scanning.
type ContentScanningUpdateParamsValue string

const (
	ContentScanningUpdateParamsValueEnabled  ContentScanningUpdateParamsValue = "enabled"
	ContentScanningUpdateParamsValueDisabled ContentScanningUpdateParamsValue = "disabled"
)

func (r ContentScanningUpdateParamsValue) IsKnown() bool {
	switch r {
	case ContentScanningUpdateParamsValueEnabled, ContentScanningUpdateParamsValueDisabled:
		return true
	}
	return false
}

type ContentScanningUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Defines the status for Content Scanning.
	Result ContentScanningUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningUpdateResponseEnvelopeJSON    `json:"-"`
}

// contentScanningUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningUpdateResponseEnvelope]
type contentScanningUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningUpdateResponseEnvelopeSuccess bool

const (
	ContentScanningUpdateResponseEnvelopeSuccessTrue ContentScanningUpdateResponseEnvelopeSuccess = true
)

func (r ContentScanningUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningDisableParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningDisableResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ContentScanningDisableResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningDisableResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningDisableResponseEnvelopeJSON    `json:"-"`
}

// contentScanningDisableResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningDisableResponseEnvelope]
type contentScanningDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningDisableResponseEnvelopeSuccess bool

const (
	ContentScanningDisableResponseEnvelopeSuccessTrue ContentScanningDisableResponseEnvelopeSuccess = true
)

func (r ContentScanningDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningEnableParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningEnableResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   ContentScanningEnableResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningEnableResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningEnableResponseEnvelopeJSON    `json:"-"`
}

// contentScanningEnableResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningEnableResponseEnvelope]
type contentScanningEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningEnableResponseEnvelopeSuccess bool

const (
	ContentScanningEnableResponseEnvelopeSuccessTrue ContentScanningEnableResponseEnvelopeSuccess = true
)

func (r ContentScanningEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Defines the status for Content Scanning.
	Result ContentScanningGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningGetResponseEnvelopeJSON    `json:"-"`
}

// contentScanningGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ContentScanningGetResponseEnvelope]
type contentScanningGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningGetResponseEnvelopeSuccess bool

const (
	ContentScanningGetResponseEnvelopeSuccessTrue ContentScanningGetResponseEnvelopeSuccess = true
)

func (r ContentScanningGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
