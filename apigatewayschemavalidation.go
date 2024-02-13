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

// APIGatewaySchemaValidationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAPIGatewaySchemaValidationService] method instead.
type APIGatewaySchemaValidationService struct {
	Options []option.RequestOption
}

// NewAPIGatewaySchemaValidationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAPIGatewaySchemaValidationService(opts ...option.RequestOption) (r *APIGatewaySchemaValidationService) {
	r = &APIGatewaySchemaValidationService{}
	r.Options = opts
	return
}

// Updates zone level schema validation settings on the zone
func (r *APIGatewaySchemaValidationService) Update(ctx context.Context, zoneID string, body APIGatewaySchemaValidationUpdateParams, opts ...option.RequestOption) (res *APIShieldZoneSchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type APIGatewaySchemaValidationUpdateParams struct {
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
	ValidationDefaultMitigationAction param.Field[APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	//
	// To clear any override, use the special value `disable_override`
	//
	// `null` will have no effect.
	ValidationOverrideMitigationAction param.Field[APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r APIGatewaySchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
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
type APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationAction string

const (
	APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationActionNone  APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationAction = "none"
	APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationActionLog   APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationAction = "log"
	APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationActionBlock APIGatewaySchemaValidationUpdateParamsValidationDefaultMitigationAction = "block"
)

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
//
// To clear any override, use the special value `disable_override`
//
// `null` will have no effect.
type APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationAction string

const (
	APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationActionNone            APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationAction = "none"
	APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride APIGatewaySchemaValidationUpdateParamsValidationOverrideMitigationAction = "disable_override"
)
