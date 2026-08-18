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

// TransformationsC2paService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransformationsC2paService] method instead.
type TransformationsC2paService struct {
	Options []option.RequestOption
}

// NewTransformationsC2paService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransformationsC2paService(opts ...option.RequestOption) (r *TransformationsC2paService) {
	r = &TransformationsC2paService{}
	r.Options = opts
	return
}

// C2PA (Coalition for Content Provenance and Authenticity) signing adds
// cryptographic metadata to images processed through Cloudflare Image
// Transformations, enabling verification of image authenticity and provenance.
func (r *TransformationsC2paService) Edit(ctx context.Context, params TransformationsC2paEditParams, opts ...option.RequestOption) (res *TransformationsC2pa, err error) {
	var env TransformationsC2paEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/transformations_c2pa", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// C2PA (Coalition for Content Provenance and Authenticity) signing adds
// cryptographic metadata to images processed through Cloudflare Image
// Transformations, enabling verification of image authenticity and provenance.
func (r *TransformationsC2paService) Get(ctx context.Context, query TransformationsC2paGetParams, opts ...option.RequestOption) (res *TransformationsC2pa, err error) {
	var env TransformationsC2paGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/transformations_c2pa", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Controls C2PA signing for images processed through Cloudflare Image
// Transformations.
type TransformationsC2pa struct {
	// ID of the zone setting.
	ID TransformationsC2paID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TransformationsC2paEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Current value of the zone setting.
	Value TransformationsC2paValue `json:"value"`
	JSON  transformationsC2paJSON  `json:"-"`
}

// transformationsC2paJSON contains the JSON metadata for the struct
// [TransformationsC2pa]
type transformationsC2paJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsC2pa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsC2paJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TransformationsC2paID string

const (
	TransformationsC2paIDImageResizingC2pa TransformationsC2paID = "image_resizing_c2pa"
)

func (r TransformationsC2paID) IsKnown() bool {
	switch r {
	case TransformationsC2paIDImageResizingC2pa:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TransformationsC2paEditable bool

const (
	TransformationsC2paEditableTrue  TransformationsC2paEditable = true
	TransformationsC2paEditableFalse TransformationsC2paEditable = false
)

func (r TransformationsC2paEditable) IsKnown() bool {
	switch r {
	case TransformationsC2paEditableTrue, TransformationsC2paEditableFalse:
		return true
	}
	return false
}

// Current value of the zone setting.
type TransformationsC2paValue string

const (
	TransformationsC2paValueOn  TransformationsC2paValue = "on"
	TransformationsC2paValueOff TransformationsC2paValue = "off"
)

func (r TransformationsC2paValue) IsKnown() bool {
	switch r {
	case TransformationsC2paValueOn, TransformationsC2paValueOff:
		return true
	}
	return false
}

type TransformationsC2paEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Whether C2PA signing is enabled for image transformations.
	Value param.Field[TransformationsC2paEditParamsValue] `json:"value" api:"required"`
}

func (r TransformationsC2paEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether C2PA signing is enabled for image transformations.
type TransformationsC2paEditParamsValue string

const (
	TransformationsC2paEditParamsValueOff TransformationsC2paEditParamsValue = "off"
	TransformationsC2paEditParamsValueOn  TransformationsC2paEditParamsValue = "on"
)

func (r TransformationsC2paEditParamsValue) IsKnown() bool {
	switch r {
	case TransformationsC2paEditParamsValueOff, TransformationsC2paEditParamsValueOn:
		return true
	}
	return false
}

type TransformationsC2paEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Controls C2PA signing for images processed through Cloudflare Image
	// Transformations.
	Result TransformationsC2pa                         `json:"result"`
	JSON   transformationsC2paEditResponseEnvelopeJSON `json:"-"`
}

// transformationsC2paEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [TransformationsC2paEditResponseEnvelope]
type transformationsC2paEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsC2paEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsC2paEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformationsC2paGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type TransformationsC2paGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Controls C2PA signing for images processed through Cloudflare Image
	// Transformations.
	Result TransformationsC2pa                        `json:"result"`
	JSON   transformationsC2paGetResponseEnvelopeJSON `json:"-"`
}

// transformationsC2paGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TransformationsC2paGetResponseEnvelope]
type transformationsC2paGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformationsC2paGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformationsC2paGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
