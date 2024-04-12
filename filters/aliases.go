// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filters

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apierror"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type AuditLog = shared.AuditLog

// This is an alias to an internal type.
type AuditLogAction = shared.AuditLogAction

// This is an alias to an internal type.
type AuditLogActor = shared.AuditLogActor

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
//
// This is an alias to an internal type.
type AuditLogActorType = shared.AuditLogActorType

// This is an alias to an internal value.
const AuditLogActorTypeUser = shared.AuditLogActorTypeUser

// This is an alias to an internal value.
const AuditLogActorTypeAdmin = shared.AuditLogActorTypeAdmin

// This is an alias to an internal value.
const AuditLogActorTypeCloudflare = shared.AuditLogActorTypeCloudflare

// This is an alias to an internal type.
type AuditLogOwner = shared.AuditLogOwner

// This is an alias to an internal type.
type AuditLogResource = shared.AuditLogResource

// This is an alias to an internal type.
type ErrorData = shared.ErrorData

// This is an alias to an internal type.
type ResponseInfo = shared.ResponseInfo

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// This is an alias to an internal type.
type Tunnel = shared.Tunnel

// This is an alias to an internal type.
type TunnelConnection = shared.TunnelConnection

// The type of tunnel.
//
// This is an alias to an internal type.
type TunnelTunType = shared.TunnelTunType

// This is an alias to an internal value.
const TunnelTunTypeCfdTunnel = shared.TunnelTunTypeCfdTunnel

// This is an alias to an internal value.
const TunnelTunTypeWARPConnector = shared.TunnelTunTypeWARPConnector

// This is an alias to an internal value.
const TunnelTunTypeIPSec = shared.TunnelTunTypeIPSec

// This is an alias to an internal value.
const TunnelTunTypeGRE = shared.TunnelTunTypeGRE

// This is an alias to an internal value.
const TunnelTunTypeCni = shared.TunnelTunTypeCni

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// This is an alias to an internal type.
type TunnelParam = shared.TunnelParam

// This is an alias to an internal type.
type TunnelConnectionParam = shared.TunnelConnectionParam
