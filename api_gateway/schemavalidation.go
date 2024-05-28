// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SchemaValidationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchemaValidationService] method instead.
type SchemaValidationService struct {
	Options []option.RequestOption
}

// NewSchemaValidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSchemaValidationService(opts ...option.RequestOption) (r *SchemaValidationService) {
	r = &SchemaValidationService{}
	r.Options = opts
	return
}

// Updates zone level schema validation settings on the zone
func (r *SchemaValidationService) Edit(ctx context.Context, params SchemaValidationEditParams, opts ...option.RequestOption) (res *Settings, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type SchemaValidationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	ValidationDefaultMitigationAction param.Field[SchemaValidationEditParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	//
	// To clear any override, use the special value `disable_override`
	//
	// `null` will have no effect.
	ValidationOverrideMitigationAction param.Field[SchemaValidationEditParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r SchemaValidationEditParams) MarshalJSON() (data []byte, err error) {
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
type SchemaValidationEditParamsValidationDefaultMitigationAction string

const (
	SchemaValidationEditParamsValidationDefaultMitigationActionNone  SchemaValidationEditParamsValidationDefaultMitigationAction = "none"
	SchemaValidationEditParamsValidationDefaultMitigationActionLog   SchemaValidationEditParamsValidationDefaultMitigationAction = "log"
	SchemaValidationEditParamsValidationDefaultMitigationActionBlock SchemaValidationEditParamsValidationDefaultMitigationAction = "block"
)

func (r SchemaValidationEditParamsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SchemaValidationEditParamsValidationDefaultMitigationActionNone, SchemaValidationEditParamsValidationDefaultMitigationActionLog, SchemaValidationEditParamsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
//
// To clear any override, use the special value `disable_override`
//
// `null` will have no effect.
type SchemaValidationEditParamsValidationOverrideMitigationAction string

const (
	SchemaValidationEditParamsValidationOverrideMitigationActionNone            SchemaValidationEditParamsValidationOverrideMitigationAction = "none"
	SchemaValidationEditParamsValidationOverrideMitigationActionDisableOverride SchemaValidationEditParamsValidationOverrideMitigationAction = "disable_override"
)

func (r SchemaValidationEditParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SchemaValidationEditParamsValidationOverrideMitigationActionNone, SchemaValidationEditParamsValidationOverrideMitigationActionDisableOverride:
		return true
	}
	return false
}
