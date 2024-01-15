// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAPIGatewaySettingSchemaValidationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneAPIGatewaySettingSchemaValidationService] method instead.
type ZoneAPIGatewaySettingSchemaValidationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewaySettingSchemaValidationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneAPIGatewaySettingSchemaValidationService(opts ...option.RequestOption) (r *ZoneAPIGatewaySettingSchemaValidationService) {
	r = &ZoneAPIGatewaySettingSchemaValidationService{}
	r.Options = opts
	return
}

// Retrieves zone level schema validation settings currently set on the zone
func (r *ZoneAPIGatewaySettingSchemaValidationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAPIGatewaySettingSchemaValidationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates zone level schema validation settings on the zone
func (r *ZoneAPIGatewaySettingSchemaValidationService) Update(ctx context.Context, zoneID string, body ZoneAPIGatewaySettingSchemaValidationUpdateParams, opts ...option.RequestOption) (res *ZoneAPIGatewaySettingSchemaValidationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates zone level schema validation settings on the zone
func (r *ZoneAPIGatewaySettingSchemaValidationService) Patch(ctx context.Context, zoneID string, body ZoneAPIGatewaySettingSchemaValidationPatchParams, opts ...option.RequestOption) (res *ZoneAPIGatewaySettingSchemaValidationPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneAPIGatewaySettingSchemaValidationGetResponse struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction ZoneAPIGatewaySettingSchemaValidationGetResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               zoneAPIGatewaySettingSchemaValidationGetResponseJSON                               `json:"-"`
}

// zoneAPIGatewaySettingSchemaValidationGetResponseJSON contains the JSON metadata
// for the struct [ZoneAPIGatewaySettingSchemaValidationGetResponse]
type zoneAPIGatewaySettingSchemaValidationGetResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ZoneAPIGatewaySettingSchemaValidationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationActionNone  ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationActionLog   ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationAction = "log"
	ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationActionBlock ZoneAPIGatewaySettingSchemaValidationGetResponseValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type ZoneAPIGatewaySettingSchemaValidationGetResponseValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationGetResponseValidationOverrideMitigationActionNone ZoneAPIGatewaySettingSchemaValidationGetResponseValidationOverrideMitigationAction = "none"
)

type ZoneAPIGatewaySettingSchemaValidationUpdateResponse struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               zoneAPIGatewaySettingSchemaValidationUpdateResponseJSON                               `json:"-"`
}

// zoneAPIGatewaySettingSchemaValidationUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneAPIGatewaySettingSchemaValidationUpdateResponse]
type zoneAPIGatewaySettingSchemaValidationUpdateResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ZoneAPIGatewaySettingSchemaValidationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationActionNone  ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationActionLog   ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationAction = "log"
	ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationActionBlock ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationOverrideMitigationActionNone ZoneAPIGatewaySettingSchemaValidationUpdateResponseValidationOverrideMitigationAction = "none"
)

type ZoneAPIGatewaySettingSchemaValidationPatchResponse struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               zoneAPIGatewaySettingSchemaValidationPatchResponseJSON                               `json:"-"`
}

// zoneAPIGatewaySettingSchemaValidationPatchResponseJSON contains the JSON
// metadata for the struct [ZoneAPIGatewaySettingSchemaValidationPatchResponse]
type zoneAPIGatewaySettingSchemaValidationPatchResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ZoneAPIGatewaySettingSchemaValidationPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationActionNone  ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationActionLog   ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationAction = "log"
	ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationActionBlock ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationOverrideMitigationActionNone ZoneAPIGatewaySettingSchemaValidationPatchResponseValidationOverrideMitigationAction = "none"
)

type ZoneAPIGatewaySettingSchemaValidationUpdateParams struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction param.Field[ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action,required"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	//
	// To clear any override, use the special value `disable_override` or `null`
	ValidationOverrideMitigationAction param.Field[ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r ZoneAPIGatewaySettingSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionNone  ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionLog   ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "log"
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionBlock ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
//
// To clear any override, use the special value `disable_override` or `null`
type ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone            ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "disable_override"
)

type ZoneAPIGatewaySettingSchemaValidationPatchParams struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	//
	// `null` will have no effect.
	ValidationDefaultMitigationAction param.Field[ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	//
	// To clear any override, use the special value `disable_override`
	//
	// `null` will have no effect.
	ValidationOverrideMitigationAction param.Field[ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r ZoneAPIGatewaySettingSchemaValidationPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used when there is no mitigation action defined on
// the operation Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
//
// `null` will have no effect.
type ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationActionNone  ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationActionLog   ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationAction = "log"
	ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationActionBlock ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
//
// To clear any override, use the special value `disable_override`
//
// `null` will have no effect.
type ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationActionNone            ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationActionDisableOverride ZoneAPIGatewaySettingSchemaValidationPatchParamsValidationOverrideMitigationAction = "disable_override"
)
