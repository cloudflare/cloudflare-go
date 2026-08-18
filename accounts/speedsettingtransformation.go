// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SpeedSettingTransformationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeedSettingTransformationService] method instead.
type SpeedSettingTransformationService struct {
	Options []option.RequestOption
}

// NewSpeedSettingTransformationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpeedSettingTransformationService(opts ...option.RequestOption) (r *SpeedSettingTransformationService) {
	r = &SpeedSettingTransformationService{}
	r.Options = opts
	return
}

// Returns a list of Image Resizing configurations across all zones for the
// account. This endpoint is useful for retrieving the transformations
// (image_resizing) state for all zones belonging to an account.
func (r *SpeedSettingTransformationService) Get(ctx context.Context, query SpeedSettingTransformationGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[TransformationsConfig], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/settings/transformations", query.AccountID)
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

// Returns a list of Image Resizing configurations across all zones for the
// account. This endpoint is useful for retrieving the transformations
// (image_resizing) state for all zones belonging to an account.
func (r *SpeedSettingTransformationService) GetAutoPaging(ctx context.Context, query SpeedSettingTransformationGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TransformationsConfig] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

// A configuration item for a specific zone and feature.
type TransformationsConfig struct {
	// Feature identifier.
	ID string `json:"id"`
	// Zone tag identifier.
	CfZoneTag string `json:"cf_zone_tag"`
	// Whether this setting can be modified.
	Editable bool `json:"editable"`
	// When this setting was last modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Current value of the feature setting.
	Value string                    `json:"value"`
	JSON  transformationsConfigJSON `json:"-"`
}

// transformationsConfigJSON contains the JSON metadata for the struct
// [TransformationsConfig]
type transformationsConfigJSON struct {
	ID          apijson.Field
	CfZoneTag   apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsConfigJSON) RawJSON() string {
	return r.raw
}

type SpeedSettingTransformationGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
