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

// APIGatewaySettingSchemaValidationService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAPIGatewaySettingSchemaValidationService] method instead.
type APIGatewaySettingSchemaValidationService struct {
	Options []option.RequestOption
}

// NewAPIGatewaySettingSchemaValidationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIGatewaySettingSchemaValidationService(opts ...option.RequestOption) (r *APIGatewaySettingSchemaValidationService) {
	r = &APIGatewaySettingSchemaValidationService{}
	r.Options = opts
	return
}

// Updates zone level schema validation settings on the zone
func (r *APIGatewaySettingSchemaValidationService) Update(ctx context.Context, zoneID string, body APIGatewaySettingSchemaValidationUpdateParams, opts ...option.RequestOption) (res *APIShieldZoneSchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APIShieldZoneSchemaValidationSettings struct {
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
	ValidationDefaultMitigationAction APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction APIShieldZoneSchemaValidationSettingsValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               apiShieldZoneSchemaValidationSettingsJSON                               `json:"-"`
}

// apiShieldZoneSchemaValidationSettingsJSON contains the JSON metadata for the
// struct [APIShieldZoneSchemaValidationSettings]
type apiShieldZoneSchemaValidationSettingsJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *APIShieldZoneSchemaValidationSettings) UnmarshalJSON(data []byte) (err error) {
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
type APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationAction string

const (
	APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationActionNone  APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationAction = "none"
	APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationActionLog   APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationAction = "log"
	APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationActionBlock APIShieldZoneSchemaValidationSettingsValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type APIShieldZoneSchemaValidationSettingsValidationOverrideMitigationAction string

const (
	APIShieldZoneSchemaValidationSettingsValidationOverrideMitigationActionNone APIShieldZoneSchemaValidationSettingsValidationOverrideMitigationAction = "none"
)

type APIGatewaySettingSchemaValidationUpdateParams struct {
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
	ValidationDefaultMitigationAction param.Field[APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action,required"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	//
	// To clear any override, use the special value `disable_override` or `null`
	ValidationOverrideMitigationAction param.Field[APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r APIGatewaySettingSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
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
type APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction string

const (
	APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionNone  APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "none"
	APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionLog   APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "log"
	APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationActionBlock APIGatewaySettingSchemaValidationUpdateParamsValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
//
// To clear any override, use the special value `disable_override` or `null`
type APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction string

const (
	APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone            APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "none"
	APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride APIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "disable_override"
)
