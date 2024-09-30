// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// ScreenshotService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScreenshotService] method instead.
type ScreenshotService struct {
	Options []option.RequestOption
}

// NewScreenshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScreenshotService(opts ...option.RequestOption) (r *ScreenshotService) {
	r = &ScreenshotService{}
	r.Options = opts
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *ScreenshotService) Get(ctx context.Context, accountID string, scanID string, query ScreenshotGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/screenshots/%s.png", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ScreenshotGetParams struct {
	// Target device type.
	Resolution param.Field[ScreenshotGetParamsResolution] `query:"resolution"`
}

// URLQuery serializes [ScreenshotGetParams]'s query parameters as `url.Values`.
func (r ScreenshotGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Target device type.
type ScreenshotGetParamsResolution string

const (
	ScreenshotGetParamsResolutionDesktop ScreenshotGetParamsResolution = "desktop"
	ScreenshotGetParamsResolutionMobile  ScreenshotGetParamsResolution = "mobile"
	ScreenshotGetParamsResolutionTablet  ScreenshotGetParamsResolution = "tablet"
)

func (r ScreenshotGetParamsResolution) IsKnown() bool {
	switch r {
	case ScreenshotGetParamsResolutionDesktop, ScreenshotGetParamsResolutionMobile, ScreenshotGetParamsResolutionTablet:
		return true
	}
	return false
}
