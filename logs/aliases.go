// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

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

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// This is an alias to an internal type.
type CloudflareTunnel = shared.CloudflareTunnel

// This is an alias to an internal type.
type CloudflareTunnelConnection = shared.CloudflareTunnelConnection

// The type of tunnel.
//
// This is an alias to an internal type.
type CloudflareTunnelTunType = shared.CloudflareTunnelTunType

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCfdTunnel = shared.CloudflareTunnelTunTypeCfdTunnel

// This is an alias to an internal value.
const CloudflareTunnelTunTypeWARPConnector = shared.CloudflareTunnelTunTypeWARPConnector

// This is an alias to an internal value.
const CloudflareTunnelTunTypeIPSec = shared.CloudflareTunnelTunTypeIPSec

// This is an alias to an internal value.
const CloudflareTunnelTunTypeGRE = shared.CloudflareTunnelTunTypeGRE

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCNI = shared.CloudflareTunnelTunTypeCNI

// This is an alias to an internal type.
type ErrorData = shared.ErrorData

// This is an alias to an internal type.
type Permission = shared.Permission

// This is an alias to an internal type.
type PermissionGrant = shared.PermissionGrant

// This is an alias to an internal type.
type PermissionGrantParam = shared.PermissionGrantParam

// This is an alias to an internal type.
type ResponseInfo = shared.ResponseInfo

// This is an alias to an internal type.
type Role = shared.Role

// This is an alias to an internal type.
type User = shared.User

// This is an alias to an internal type.
type UserRole = shared.UserRole

// This is an alias to an internal type.
type UserRolesPermissions = shared.UserRolesPermissions

// This is an alias to an internal type.
type UserUser = shared.UserUser

// This is an alias to an internal type.
type UserParam = shared.UserParam

// This is an alias to an internal type.
type UserRoleParam = shared.UserRoleParam

// This is an alias to an internal type.
type UserRolesPermissionsParam = shared.UserRolesPermissionsParam

// This is an alias to an internal type.
type UserUserParam = shared.UserUserParam
