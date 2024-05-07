// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apierror"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type ASN = shared.ASN

// This is an alias to an internal type.
type ASNParam = shared.ASNParam

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

// The Certificate Authority that will issue the certificate
//
// This is an alias to an internal type.
type CertificateCA = shared.CertificateCA

// This is an alias to an internal value.
const CertificateCADigicert = shared.CertificateCADigicert

// This is an alias to an internal value.
const CertificateCAGoogle = shared.CertificateCAGoogle

// This is an alias to an internal value.
const CertificateCALetsEncrypt = shared.CertificateCALetsEncrypt

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
//
// This is an alias to an internal type.
type CertificateRequestType = shared.CertificateRequestType

// This is an alias to an internal value.
const CertificateRequestTypeOriginRSA = shared.CertificateRequestTypeOriginRSA

// This is an alias to an internal value.
const CertificateRequestTypeOriginECC = shared.CertificateRequestTypeOriginECC

// This is an alias to an internal value.
const CertificateRequestTypeKeylessCertificate = shared.CertificateRequestTypeKeylessCertificate

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
type Member = shared.Member

// This is an alias to an internal type.
type MemberRole = shared.MemberRole

// This is an alias to an internal type.
type MemberRolesPermissions = shared.MemberRolesPermissions

// This is an alias to an internal type.
type MemberUser = shared.MemberUser

// This is an alias to an internal type.
type MemberParam = shared.MemberParam

// This is an alias to an internal type.
type MemberRoleParam = shared.MemberRoleParam

// This is an alias to an internal type.
type MemberRolesPermissionsParam = shared.MemberRolesPermissionsParam

// This is an alias to an internal type.
type MemberUserParam = shared.MemberUserParam

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
