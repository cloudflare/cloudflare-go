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
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ContentScanningSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentScanningSettingService] method instead.
type ContentScanningSettingService struct {
	Options []option.RequestOption
}

// NewContentScanningSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContentScanningSettingService(opts ...option.RequestOption) (r *ContentScanningSettingService) {
	r = &ContentScanningSettingService{}
	r.Options = opts
	return
}

// Retrieve the current status of Content Scanning
func (r *ContentScanningSettingService) Get(ctx context.Context, query ContentScanningSettingGetParams, opts ...option.RequestOption) (res *ContentScanningSettingGetResponse, err error) {
	var env ContentScanningSettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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

// The status for Content Scanning
type ContentScanningSettingGetResponse struct {
	// Last modification date (ISO 8601) of the Content Scanning status
	Modified string `json:"modified"`
	// Status of Content Scanning
	Value string                                `json:"value"`
	JSON  contentScanningSettingGetResponseJSON `json:"-"`
}

// contentScanningSettingGetResponseJSON contains the JSON metadata for the struct
// [ContentScanningSettingGetResponse]
type contentScanningSettingGetResponseJSON struct {
	Modified    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ContentScanningSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The status for Content Scanning
	Result ContentScanningSettingGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ContentScanningSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningSettingGetResponseEnvelopeJSON    `json:"-"`
}

// contentScanningSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningSettingGetResponseEnvelope]
type contentScanningSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ContentScanningSettingGetResponseEnvelopeSuccess bool

const (
	ContentScanningSettingGetResponseEnvelopeSuccessTrue ContentScanningSettingGetResponseEnvelopeSuccess = true
)

func (r ContentScanningSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
