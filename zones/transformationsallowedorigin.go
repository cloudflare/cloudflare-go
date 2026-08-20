// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// TransformationsAllowedOriginService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransformationsAllowedOriginService] method instead.
type TransformationsAllowedOriginService struct {
	Options []option.RequestOption
}

// NewTransformationsAllowedOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransformationsAllowedOriginService(opts ...option.RequestOption) (r *TransformationsAllowedOriginService) {
	r = &TransformationsAllowedOriginService{}
	r.Options = opts
	return
}

// Media Transformations Allowed Origins restricts transformations for images and
// video served through Cloudflare's network to requests originating from specified
// domains. Refer to the Image Transformations and Video Transformations
// documentation for more information.
func (r *TransformationsAllowedOriginService) Edit(ctx context.Context, params TransformationsAllowedOriginEditParams, opts ...option.RequestOption) (res *TransformationsAllowedOrigins, err error) {
	var env TransformationsAllowedOriginEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/transformations_allowed_origins", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Media Transformations Allowed Origins restricts transformations for images and
// video served through Cloudflare's network to requests originating from specified
// domains. Refer to the Image Transformations and Video Transformations
// documentation for more information.
func (r *TransformationsAllowedOriginService) Get(ctx context.Context, query TransformationsAllowedOriginGetParams, opts ...option.RequestOption) (res *TransformationsAllowedOrigins, err error) {
	var env TransformationsAllowedOriginGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/transformations_allowed_origins", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Controls which origins are allowed to request image and video transformations.
type TransformationsAllowedOrigins struct {
	// ID of the zone setting.
	ID TransformationsAllowedOriginsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TransformationsAllowedOriginsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Comma-separated list of allowed origin domains for image and video
	// transformations. Use "\*" to allow all origins (default).
	Value TransformationsAllowedOriginsValue `json:"value"`
	JSON  transformationsAllowedOriginsJSON  `json:"-"`
}

// transformationsAllowedOriginsJSON contains the JSON metadata for the struct
// [TransformationsAllowedOrigins]
type transformationsAllowedOriginsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsAllowedOrigins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsAllowedOriginsJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TransformationsAllowedOriginsID string

const (
	TransformationsAllowedOriginsIDImageResizingAllowedOrigins TransformationsAllowedOriginsID = "image_resizing_allowed_origins"
)

func (r TransformationsAllowedOriginsID) IsKnown() bool {
	switch r {
	case TransformationsAllowedOriginsIDImageResizingAllowedOrigins:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TransformationsAllowedOriginsEditable bool

const (
	TransformationsAllowedOriginsEditableTrue  TransformationsAllowedOriginsEditable = true
	TransformationsAllowedOriginsEditableFalse TransformationsAllowedOriginsEditable = false
)

func (r TransformationsAllowedOriginsEditable) IsKnown() bool {
	switch r {
	case TransformationsAllowedOriginsEditableTrue, TransformationsAllowedOriginsEditableFalse:
		return true
	}
	return false
}

// Comma-separated list of allowed origin domains for image and video
// transformations. Use "\*" to allow all origins (default).
type TransformationsAllowedOriginsValue string

const (
	TransformationsAllowedOriginsValueOn  TransformationsAllowedOriginsValue = "on"
	TransformationsAllowedOriginsValueOff TransformationsAllowedOriginsValue = "off"
)

func (r TransformationsAllowedOriginsValue) IsKnown() bool {
	switch r {
	case TransformationsAllowedOriginsValueOn, TransformationsAllowedOriginsValueOff:
		return true
	}
	return false
}

type TransformationsAllowedOriginEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Comma-separated list of allowed origin domains for image and video
	// transformations. Use "\*" to allow all origins (default).
	Value param.Field[string] `json:"value" api:"required"`
}

func (r TransformationsAllowedOriginEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransformationsAllowedOriginEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Controls which origins are allowed to request image and video transformations.
	Result TransformationsAllowedOrigins                        `json:"result"`
	JSON   transformationsAllowedOriginEditResponseEnvelopeJSON `json:"-"`
}

// transformationsAllowedOriginEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [TransformationsAllowedOriginEditResponseEnvelope]
type transformationsAllowedOriginEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsAllowedOriginEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsAllowedOriginEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformationsAllowedOriginGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type TransformationsAllowedOriginGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Controls which origins are allowed to request image and video transformations.
	Result TransformationsAllowedOrigins                       `json:"result"`
	JSON   transformationsAllowedOriginGetResponseEnvelopeJSON `json:"-"`
}

// transformationsAllowedOriginGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [TransformationsAllowedOriginGetResponseEnvelope]
type transformationsAllowedOriginGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsAllowedOriginGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsAllowedOriginGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
