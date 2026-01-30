# CertificateAuthorities

## HostnameAssociations

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationParam">HostnameAssociationParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#TLSHostnameAssociationParam">TLSHostnameAssociationParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociation">HostnameAssociation</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationUpdateResponse">HostnameAssociationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationGetResponse">HostnameAssociationGetResponse</a>

Methods:

- <code title="put /zones/{zone_id}/certificate_authorities/hostname_associations">client.CertificateAuthorities.HostnameAssociations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationUpdateParams">HostnameAssociationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationUpdateResponse">HostnameAssociationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/certificate_authorities/hostname_associations">client.CertificateAuthorities.HostnameAssociations.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationGetParams">HostnameAssociationGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities">certificate_authorities</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v6/certificate_authorities#HostnameAssociationGetResponse">HostnameAssociationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
