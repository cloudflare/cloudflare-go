// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v7/internal/apierror"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

type Error = apierror.Error

// ASN this is an alias to an internal type.
type ASN = shared.ASN

// ASNParam this is an alias to an internal type.
type ASNParam = shared.ASNParam

// AuditLog this is an alias to an internal type.
type AuditLog = shared.AuditLog

// AuditLogAction this is an alias to an internal type.
type AuditLogAction = shared.AuditLogAction

// AuditLogActor this is an alias to an internal type.
type AuditLogActor = shared.AuditLogActor

// AuditLogActorType is the type of actor, whether a User, Cloudflare Admin, or an Automated System.
//
// This is an alias to an internal type.
type AuditLogActorType = shared.AuditLogActorType

// This is an alias to an internal value.
const AuditLogActorTypeUser = shared.AuditLogActorTypeUser

// This is an alias to an internal value.
const AuditLogActorTypeAdmin = shared.AuditLogActorTypeAdmin

// This is an alias to an internal value.
const AuditLogActorTypeCloudflare = shared.AuditLogActorTypeCloudflare

// AuditLogOwner this is an alias to an internal type.
type AuditLogOwner = shared.AuditLogOwner

// AuditLogResource this is an alias to an internal type.
type AuditLogResource = shared.AuditLogResource

// CertificateCA is the Certificate Authority that will issue the certificate.
//
// This is an alias to an internal type.
type CertificateCA = shared.CertificateCA

// This is an alias to an internal value.
const CertificateCADigicert = shared.CertificateCADigicert

// This is an alias to an internal value.
const CertificateCAGoogle = shared.CertificateCAGoogle

// This is an alias to an internal value.
const CertificateCALetsEncrypt = shared.CertificateCALetsEncrypt

// This is an alias to an internal value.
const CertificateCASSLCom = shared.CertificateCASSLCom

// CertificateRequestType signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
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

// CloudflareTunnel is a Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// This is an alias to an internal type.
type CloudflareTunnel = shared.CloudflareTunnel

// CloudflareTunnelConfigSrc indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
//
// This is an alias to an internal type.
type CloudflareTunnelConfigSrc = shared.CloudflareTunnelConfigSrc

// This is an alias to an internal value.
const CloudflareTunnelConfigSrcLocal = shared.CloudflareTunnelConfigSrcLocal

// This is an alias to an internal value.
const CloudflareTunnelConfigSrcCloudflare = shared.CloudflareTunnelConfigSrcCloudflare

// CloudflareTunnelConnection this is an alias to an internal type.
type CloudflareTunnelConnection = shared.CloudflareTunnelConnection

// CloudflareTunnelStatus is the status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
//
// This is an alias to an internal type.
type CloudflareTunnelStatus = shared.CloudflareTunnelStatus

// This is an alias to an internal value.
const CloudflareTunnelStatusInactive = shared.CloudflareTunnelStatusInactive

// This is an alias to an internal value.
const CloudflareTunnelStatusDegraded = shared.CloudflareTunnelStatusDegraded

// This is an alias to an internal value.
const CloudflareTunnelStatusHealthy = shared.CloudflareTunnelStatusHealthy

// This is an alias to an internal value.
const CloudflareTunnelStatusDown = shared.CloudflareTunnelStatusDown

// CloudflareTunnelTunType is the type of tunnel.
//
// This is an alias to an internal type.
type CloudflareTunnelTunType = shared.CloudflareTunnelTunType

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCfdTunnel = shared.CloudflareTunnelTunTypeCfdTunnel

// This is an alias to an internal value.
const CloudflareTunnelTunTypeWARPConnector = shared.CloudflareTunnelTunTypeWARPConnector

// This is an alias to an internal value.
const CloudflareTunnelTunTypeWARP = shared.CloudflareTunnelTunTypeWARP

// This is an alias to an internal value.
const CloudflareTunnelTunTypeMagic = shared.CloudflareTunnelTunTypeMagic

// This is an alias to an internal value.
const CloudflareTunnelTunTypeIPSec = shared.CloudflareTunnelTunTypeIPSec

// This is an alias to an internal value.
const CloudflareTunnelTunTypeGRE = shared.CloudflareTunnelTunTypeGRE

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCNI = shared.CloudflareTunnelTunTypeCNI

// ErrorData this is an alias to an internal type.
type ErrorData = shared.ErrorData

// ErrorDataSource this is an alias to an internal type.
type ErrorDataSource = shared.ErrorDataSource

// Member this is an alias to an internal type.
type Member = shared.Member

// MemberPolicy this is an alias to an internal type.
type MemberPolicy = shared.MemberPolicy

// MemberPoliciesAccess allow or deny operations against the resources.
//
// This is an alias to an internal type.
type MemberPoliciesAccess = shared.MemberPoliciesAccess

// This is an alias to an internal value.
const MemberPoliciesAccessAllow = shared.MemberPoliciesAccessAllow

// This is an alias to an internal value.
const MemberPoliciesAccessDeny = shared.MemberPoliciesAccessDeny

// MemberPoliciesPermissionGroup is a named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type MemberPoliciesPermissionGroup = shared.MemberPoliciesPermissionGroup

// MemberPoliciesPermissionGroupsMeta attributes associated to the permission group.
//
// This is an alias to an internal type.
type MemberPoliciesPermissionGroupsMeta = shared.MemberPoliciesPermissionGroupsMeta

// MemberPoliciesResourceGroup is a group of scoped resources.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroup = shared.MemberPoliciesResourceGroup

// MemberPoliciesResourceGroupsScope is a scope is a combination of scope objects which provides additional context.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsScope = shared.MemberPoliciesResourceGroupsScope

// MemberPoliciesResourceGroupsScopeObject is a scope object represents any resource that can have actions applied against
// invite.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsScopeObject = shared.MemberPoliciesResourceGroupsScopeObject

// MemberPoliciesResourceGroupsMeta attributes associated to the resource group.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsMeta = shared.MemberPoliciesResourceGroupsMeta

// MemberStatus is a member's status in the account.
//
// This is an alias to an internal type.
type MemberStatus = shared.MemberStatus

// This is an alias to an internal value.
const MemberStatusAccepted = shared.MemberStatusAccepted

// This is an alias to an internal value.
const MemberStatusPending = shared.MemberStatusPending

// MemberUser details of the user associated to the membership.
//
// This is an alias to an internal type.
type MemberUser = shared.MemberUser

// Permission this is an alias to an internal type.
type Permission = shared.Permission

// PermissionGrant this is an alias to an internal type.
type PermissionGrant = shared.PermissionGrant

// PermissionGrantParam this is an alias to an internal type.
type PermissionGrantParam = shared.PermissionGrantParam

// RatePlan is the rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlan = shared.RatePlan

// RatePlanID is the ID of the rate plan.
//
// This is an alias to an internal type.
type RatePlanID = shared.RatePlanID

// This is an alias to an internal value.
const RatePlanIDFree = shared.RatePlanIDFree

// This is an alias to an internal value.
const RatePlanIDLite = shared.RatePlanIDLite

// This is an alias to an internal value.
const RatePlanIDPro = shared.RatePlanIDPro

// This is an alias to an internal value.
const RatePlanIDProPlus = shared.RatePlanIDProPlus

// This is an alias to an internal value.
const RatePlanIDBusiness = shared.RatePlanIDBusiness

// This is an alias to an internal value.
const RatePlanIDEnterprise = shared.RatePlanIDEnterprise

// This is an alias to an internal value.
const RatePlanIDPartnersFree = shared.RatePlanIDPartnersFree

// This is an alias to an internal value.
const RatePlanIDPartnersPro = shared.RatePlanIDPartnersPro

// This is an alias to an internal value.
const RatePlanIDPartnersBusiness = shared.RatePlanIDPartnersBusiness

// This is an alias to an internal value.
const RatePlanIDPartnersEnterprise = shared.RatePlanIDPartnersEnterprise

// RatePlanParam is the rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlanParam = shared.RatePlanParam

// ResponseInfo this is an alias to an internal type.
type ResponseInfo = shared.ResponseInfo

// ResponseInfoSource this is an alias to an internal type.
type ResponseInfoSource = shared.ResponseInfoSource

// Role this is an alias to an internal type.
type Role = shared.Role

// RolePermissions this is an alias to an internal type.
type RolePermissions = shared.RolePermissions

// RoleParam this is an alias to an internal type.
type RoleParam = shared.RoleParam

// RolePermissionsParam this is an alias to an internal type.
type RolePermissionsParam = shared.RolePermissionsParam

// SortDirection direction to order DNS records in.
//
// This is an alias to an internal type.
type SortDirection = shared.SortDirection

// This is an alias to an internal value.
const SortDirectionAsc = shared.SortDirectionAsc

// This is an alias to an internal value.
const SortDirectionDesc = shared.SortDirectionDesc

// Subscription this is an alias to an internal type.
type Subscription = shared.Subscription

// SubscriptionFrequency how often the subscription is renewed automatically.
//
// This is an alias to an internal type.
type SubscriptionFrequency = shared.SubscriptionFrequency

// This is an alias to an internal value.
const SubscriptionFrequencyWeekly = shared.SubscriptionFrequencyWeekly

// This is an alias to an internal value.
const SubscriptionFrequencyMonthly = shared.SubscriptionFrequencyMonthly

// This is an alias to an internal value.
const SubscriptionFrequencyQuarterly = shared.SubscriptionFrequencyQuarterly

// This is an alias to an internal value.
const SubscriptionFrequencyYearly = shared.SubscriptionFrequencyYearly

// SubscriptionState is the state that the subscription is in.
//
// This is an alias to an internal type.
type SubscriptionState = shared.SubscriptionState

// This is an alias to an internal value.
const SubscriptionStateTrial = shared.SubscriptionStateTrial

// This is an alias to an internal value.
const SubscriptionStateProvisioned = shared.SubscriptionStateProvisioned

// This is an alias to an internal value.
const SubscriptionStatePaid = shared.SubscriptionStatePaid

// This is an alias to an internal value.
const SubscriptionStateAwaitingPayment = shared.SubscriptionStateAwaitingPayment

// This is an alias to an internal value.
const SubscriptionStateCancelled = shared.SubscriptionStateCancelled

// This is an alias to an internal value.
const SubscriptionStateFailed = shared.SubscriptionStateFailed

// This is an alias to an internal value.
const SubscriptionStateExpired = shared.SubscriptionStateExpired

// SubscriptionParam this is an alias to an internal type.
type SubscriptionParam = shared.SubscriptionParam

// Token this is an alias to an internal type.
type Token = shared.Token

// TokenCondition this is an alias to an internal type.
type TokenCondition = shared.TokenCondition

// TokenConditionRequestIP client IP restrictions.
//
// This is an alias to an internal type.
type TokenConditionRequestIP = shared.TokenConditionRequestIP

// TokenStatus status of the token.
//
// This is an alias to an internal type.
type TokenStatus = shared.TokenStatus

// This is an alias to an internal value.
const TokenStatusActive = shared.TokenStatusActive

// This is an alias to an internal value.
const TokenStatusDisabled = shared.TokenStatusDisabled

// This is an alias to an internal value.
const TokenStatusExpired = shared.TokenStatusExpired

// TokenParam this is an alias to an internal type.
type TokenParam = shared.TokenParam

// TokenConditionParam this is an alias to an internal type.
type TokenConditionParam = shared.TokenConditionParam

// TokenConditionRequestIPParam client IP restrictions.
//
// This is an alias to an internal type.
type TokenConditionRequestIPParam = shared.TokenConditionRequestIPParam

// TokenConditionCIDRList iPv4/IPv6 CIDR.
//
// This is an alias to an internal type.
type TokenConditionCIDRList = shared.TokenConditionCIDRList

// TokenConditionCIDRListParam iPv4/IPv6 CIDR.
//
// This is an alias to an internal type.
type TokenConditionCIDRListParam = shared.TokenConditionCIDRListParam

// TokenPolicy this is an alias to an internal type.
type TokenPolicy = shared.TokenPolicy

// TokenPolicyEffect allow or deny operations against the resources.
//
// This is an alias to an internal type.
type TokenPolicyEffect = shared.TokenPolicyEffect

// This is an alias to an internal value.
const TokenPolicyEffectAllow = shared.TokenPolicyEffectAllow

// This is an alias to an internal value.
const TokenPolicyEffectDeny = shared.TokenPolicyEffectDeny

// TokenPolicyPermissionGroup is a named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroup = shared.TokenPolicyPermissionGroup

// TokenPolicyPermissionGroupsMeta attributes associated to the permission group.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupsMeta = shared.TokenPolicyPermissionGroupsMeta

// TokenPolicyResourcesUnion is a list of resource names that the policy applies to.
//
// This is an alias to an internal type.
type TokenPolicyResourcesUnion = shared.TokenPolicyResourcesUnion

// TokenPolicyResourcesIAMResourcesTypeObjectString map of simple string resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectString = shared.TokenPolicyResourcesIAMResourcesTypeObjectString

// TokenPolicyResourcesIAMResourcesTypeObjectNested map of nested resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectNested = shared.TokenPolicyResourcesIAMResourcesTypeObjectNested

// TokenPolicyParam this is an alias to an internal type.
type TokenPolicyParam = shared.TokenPolicyParam

// TokenPolicyPermissionGroupParam is a named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupParam = shared.TokenPolicyPermissionGroupParam

// TokenPolicyPermissionGroupsMetaParam attributes associated to the permission group.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupsMetaParam = shared.TokenPolicyPermissionGroupsMetaParam

// TokenPolicyResourcesUnionParam is a list of resource names that the policy applies to.
//
// This is an alias to an internal type.
type TokenPolicyResourcesUnionParam = shared.TokenPolicyResourcesUnionParam

// TokenPolicyResourcesIAMResourcesTypeObjectStringParam map of simple string resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectStringParam = shared.TokenPolicyResourcesIAMResourcesTypeObjectStringParam

// TokenPolicyResourcesIAMResourcesTypeObjectNestedParam map of nested resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectNestedParam = shared.TokenPolicyResourcesIAMResourcesTypeObjectNestedParam

// TokenValue is the token value.
//
// This is an alias to an internal type.
type TokenValue = shared.TokenValue
