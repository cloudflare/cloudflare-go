# Zones

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Type">Type</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Type">Type</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Zone">Zone</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneDeleteResponse">ZoneDeleteResponse</a>

Methods:

- <code title="post /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneNewParams">ZoneNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Zone">Zone</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneListParams">ZoneListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#V4PagePaginationArray">V4PagePaginationArray</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Zone">Zone</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneDeleteParams">ZoneDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneDeleteResponse">ZoneDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneEditParams">ZoneEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Zone">Zone</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}">client.Zones.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneGetParams">ZoneGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Zone">Zone</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ActivationCheck

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ActivationCheckTriggerResponse">ActivationCheckTriggerResponse</a>

Methods:

- <code title="put /zones/{zone_id}/activation_check">client.Zones.ActivationCheck.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ActivationCheckService.Trigger">Trigger</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ActivationCheckTriggerParams">ActivationCheckTriggerParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ActivationCheckTriggerResponse">ActivationCheckTriggerResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Settings

Params Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AlwaysUseHTTPSParam">AlwaysUseHTTPSParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AutomaticHTTPSRewritesParam">AutomaticHTTPSRewritesParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AutomaticPlatformOptimizationParam">AutomaticPlatformOptimizationParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#BrowserCacheTTLParam">BrowserCacheTTLParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#BrowserCheckParam">BrowserCheckParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CacheLevelParam">CacheLevelParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EmailObfuscationParam">EmailObfuscationParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#IPGeolocationParam">IPGeolocationParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#MirageParam">MirageParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OpportunisticEncryptionParam">OpportunisticEncryptionParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OriginErrorPagePassThruParam">OriginErrorPagePassThruParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PolishParam">PolishParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ResponseBufferingParam">ResponseBufferingParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RocketLoaderParam">RocketLoaderParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SecurityLevelParam">SecurityLevelParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SortQueryStringForCacheParam">SortQueryStringForCacheParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SSLParam">SSLParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#TrueClientIPHeaderParam">TrueClientIPHeaderParam</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#WAFParam">WAFParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AdvancedDDoS">AdvancedDDoS</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AlwaysOnline">AlwaysOnline</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AlwaysUseHTTPS">AlwaysUseHTTPS</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AutomaticHTTPSRewrites">AutomaticHTTPSRewrites</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AutomaticPlatformOptimization">AutomaticPlatformOptimization</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Brotli">Brotli</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#BrowserCacheTTL">BrowserCacheTTL</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#BrowserCheck">BrowserCheck</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CacheLevel">CacheLevel</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ChallengeTTL">ChallengeTTL</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Ciphers">Ciphers</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#DevelopmentMode">DevelopmentMode</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EarlyHints">EarlyHints</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EmailObfuscation">EmailObfuscation</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#H2Prioritization">H2Prioritization</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HotlinkProtection">HotlinkProtection</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HTTP2">HTTP2</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HTTP3">HTTP3</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ImageResizing">ImageResizing</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#IPGeolocation">IPGeolocation</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#IPV6">IPV6</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#MinTLSVersion">MinTLSVersion</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Mirage">Mirage</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#NEL">NEL</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OpportunisticEncryption">OpportunisticEncryption</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OpportunisticOnion">OpportunisticOnion</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OrangeToOrange">OrangeToOrange</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#OriginErrorPagePassThru">OriginErrorPagePassThru</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Polish">Polish</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PrefetchPreload">PrefetchPreload</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ProxyReadTimeout">ProxyReadTimeout</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PseudoIPV4">PseudoIPV4</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ResponseBuffering">ResponseBuffering</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RocketLoader">RocketLoader</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SecurityHeaders">SecurityHeaders</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SecurityLevel">SecurityLevel</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ServerSideExcludes">ServerSideExcludes</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SortQueryStringForCache">SortQueryStringForCache</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SSL">SSL</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SSLRecommender">SSLRecommender</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#TLS1_3">TLS1_3</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#TLSClientAuth">TLSClientAuth</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#TrueClientIPHeader">TrueClientIPHeader</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#WAF">WAF</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#WebP">WebP</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#Websocket">Websocket</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZeroRTT">ZeroRTT</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingEditResponse">SettingEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingGetResponse">SettingGetResponse</a>

Methods:

- <code title="patch /zones/{zone_id}/settings/{setting_id}">client.Zones.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, settingID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingEditParams">SettingEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingEditResponse">SettingEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/settings/{setting_id}">client.Zones.Settings.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, settingID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingGetParams">SettingGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SettingGetResponse">SettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Environments

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentNewResponse">EnvironmentNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentUpdateResponse">EnvironmentUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentListResponse">EnvironmentListResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentDeleteResponse">EnvironmentDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentEditResponse">EnvironmentEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentRollbackResponse">EnvironmentRollbackResponse</a>

Methods:

- <code title="post /zones/{zone_id}/environments">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentNewParams">EnvironmentNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentNewResponse">EnvironmentNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/environments">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentUpdateParams">EnvironmentUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentUpdateResponse">EnvironmentUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/environments">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentListParams">EnvironmentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentListResponse">EnvironmentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/environments/{environment_id}">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentDeleteParams">EnvironmentDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentDeleteResponse">EnvironmentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/environments">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentEditParams">EnvironmentEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentEditResponse">EnvironmentEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /zones/{zone_id}/environments/{environment_id}/rollback">client.Zones.Environments.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentService.Rollback">Rollback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, environmentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentRollbackParams">EnvironmentRollbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#EnvironmentRollbackResponse">EnvironmentRollbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CustomNameservers

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverGetResponse">CustomNameserverGetResponse</a>

Methods:

- <code title="put /zones/{zone_id}/custom_ns">client.Zones.CustomNameservers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverUpdateParams">CustomNameserverUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/builtin#string">string</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/custom_ns">client.Zones.CustomNameservers.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverGetParams">CustomNameserverGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CustomNameserverGetResponse">CustomNameserverGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Holds

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneHold">ZoneHold</a>

Methods:

- <code title="post /zones/{zone_id}/hold">client.Zones.Holds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldNewParams">HoldNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneHold">ZoneHold</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /zones/{zone_id}/hold">client.Zones.Holds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldDeleteParams">HoldDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneHold">ZoneHold</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /zones/{zone_id}/hold">client.Zones.Holds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldEditParams">HoldEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneHold">ZoneHold</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/hold">client.Zones.Holds.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#HoldGetParams">HoldGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#ZoneHold">ZoneHold</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Subscriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionNewResponse">SubscriptionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionUpdateResponse">SubscriptionUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionGetResponse">SubscriptionGetResponse</a>

Methods:

- <code title="post /zones/{zone_id}/subscription">client.Zones.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionNewParams">SubscriptionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionNewResponse">SubscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /zones/{zone_id}/subscription">client.Zones.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionUpdateParams">SubscriptionUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionUpdateResponse">SubscriptionUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/subscription">client.Zones.Subscriptions.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionGetParams">SubscriptionGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#SubscriptionGetResponse">SubscriptionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Plans

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AvailableRatePlan">AvailableRatePlan</a>

Methods:

- <code title="get /zones/{zone_id}/available_plans">client.Zones.Plans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PlanService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PlanListParams">PlanListParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AvailableRatePlan">AvailableRatePlan</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/available_plans/{plan_identifier}">client.Zones.Plans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PlanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, planIdentifier <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#PlanGetParams">PlanGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#AvailableRatePlan">AvailableRatePlan</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RatePlans

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RatePlanGetResponse">RatePlanGetResponse</a>

Methods:

- <code title="get /zones/{zone_id}/available_rate_plans">client.Zones.RatePlans.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RatePlanService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RatePlanGetParams">RatePlanGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/packages/pagination#SinglePage">SinglePage</a>[<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#RatePlanGetResponse">RatePlanGetResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## CT

### Alerting

Response Types:

- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingEditResponse">CTAlertingEditResponse</a>
- <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingGetResponse">CTAlertingGetResponse</a>

Methods:

- <code title="patch /zones/{zone_id}/ct/alerting">client.Zones.CT.Alerting.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingService.Edit">Edit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingEditParams">CTAlertingEditParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingEditResponse">CTAlertingEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /zones/{zone_id}/ct/alerting">client.Zones.CT.Alerting.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingGetParams">CTAlertingGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones">zones</a>.<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-go/v7/zones#CTAlertingGetResponse">CTAlertingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
