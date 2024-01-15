// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountUrlscannerScreenshotService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountUrlscannerScreenshotService] method instead.
type AccountUrlscannerScreenshotService struct {
	Options []option.RequestOption
}

// NewAccountUrlscannerScreenshotService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountUrlscannerScreenshotService(opts ...option.RequestOption) (r *AccountUrlscannerScreenshotService) {
	r = &AccountUrlscannerScreenshotService{}
	r.Options = opts
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *AccountUrlscannerScreenshotService) Get(ctx context.Context, accountID string, scanID string, query AccountUrlscannerScreenshotGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/screenshot", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountUrlscannerScreenshotGetParams struct {
	// Target device type
	Resolution param.Field[AccountUrlscannerScreenshotGetParamsResolution] `query:"resolution"`
}

// URLQuery serializes [AccountUrlscannerScreenshotGetParams]'s query parameters as
// `url.Values`.
func (r AccountUrlscannerScreenshotGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Target device type
type AccountUrlscannerScreenshotGetParamsResolution string

const (
	AccountUrlscannerScreenshotGetParamsResolutionDesktop AccountUrlscannerScreenshotGetParamsResolution = "desktop"
	AccountUrlscannerScreenshotGetParamsResolutionMobile  AccountUrlscannerScreenshotGetParamsResolution = "mobile"
	AccountUrlscannerScreenshotGetParamsResolutionTablet  AccountUrlscannerScreenshotGetParamsResolution = "tablet"
)
