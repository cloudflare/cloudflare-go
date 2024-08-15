// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.New(context.TODO(), load_balancers.LoadBalancerNewParams{
		ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
		DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
		FallbackPool: cloudflare.F("fallback_pool"),
		Name:         cloudflare.F("www.example.com"),
		AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
			FailoverAcrossPools: cloudflare.F(true),
		}),
		CountryPools: cloudflare.F(load_balancers.LoadBalancerNewParamsCountryPools{
			CountryCode: cloudflare.F([]string{"string", "string", "string"}),
		}),
		Description: cloudflare.F("Load Balancer for www.example.com"),
		LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
			Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
			PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
		}),
		PopPools: cloudflare.F(load_balancers.LoadBalancerNewParamsPopPools{
			Pop: cloudflare.F([]string{"string", "string", "string"}),
		}),
		Proxied: cloudflare.F(true),
		RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
			DefaultWeight: cloudflare.F(0.200000),
			PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
				Key:   cloudflare.F("key"),
				Value: cloudflare.F(0.000000),
			}),
		}),
		RegionPools: cloudflare.F(load_balancers.LoadBalancerNewParamsRegionPools{
			RegionCode: cloudflare.F([]string{"string", "string", "string"}),
		}),
		Rules: cloudflare.F([]load_balancers.RulesParam{{
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
					CountryCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F("fallback_pool"),
				LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
					Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
					PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
					Pop: cloudflare.F([]string{"string", "string", "string"}),
				}),
				RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   cloudflare.F("key"),
						Value: cloudflare.F(0.000000),
					}),
				}),
				RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
					RegionCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
					CountryCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F("fallback_pool"),
				LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
					Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
					PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
					Pop: cloudflare.F([]string{"string", "string", "string"}),
				}),
				RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   cloudflare.F("key"),
						Value: cloudflare.F(0.000000),
					}),
				}),
				RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
					RegionCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
				AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
					FailoverAcrossPools: cloudflare.F(true),
				}),
				CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
					CountryCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
				FallbackPool: cloudflare.F("fallback_pool"),
				LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
					Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
					PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
				}),
				PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
					Pop: cloudflare.F([]string{"string", "string", "string"}),
				}),
				RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
					DefaultWeight: cloudflare.F(0.200000),
					PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
						Key:   cloudflare.F("key"),
						Value: cloudflare.F(0.000000),
					}),
				}),
				RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
					RegionCode: cloudflare.F([]string{"string", "string", "string"}),
				}),
				SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
				SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}}),
		SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
		SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
			DrainDuration:        cloudflare.F(100.000000),
			Headers:              cloudflare.F([]string{"x"}),
			RequireAllHeaders:    cloudflare.F(true),
			Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
			Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
			ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
		}),
		SessionAffinityTTL: cloudflare.F(1800.000000),
		SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
		TTL:                cloudflare.F(30.000000),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerUpdateParams{
			ZoneID:       cloudflare.F("699d98642c564d2e855e9661899b7252"),
			DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F("fallback_pool"),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F(load_balancers.LoadBalancerUpdateParamsCountryPools{
				CountryCode: cloudflare.F([]string{"string", "string", "string"}),
			}),
			Description: cloudflare.F("Load Balancer for www.example.com"),
			Enabled:     cloudflare.F(true),
			LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
				Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
				PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			PopPools: cloudflare.F(load_balancers.LoadBalancerUpdateParamsPopPools{
				Pop: cloudflare.F([]string{"string", "string", "string"}),
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F(0.000000),
				}),
			}),
			RegionPools: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRegionPools{
				RegionCode: cloudflare.F([]string{"string", "string", "string"}),
			}),
			Rules: cloudflare.F([]load_balancers.RulesParam{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
			TTL:                cloudflare.F(30.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.List(context.TODO(), load_balancers.LoadBalancerListParams{
		ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerDeleteParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerEditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Edit(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerEditParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
			AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F(load_balancers.LoadBalancerEditParamsCountryPools{
				CountryCode: cloudflare.F([]string{"string", "string", "string"}),
			}),
			DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			Description:  cloudflare.F("Load Balancer for www.example.com"),
			Enabled:      cloudflare.F(true),
			FallbackPool: cloudflare.F("fallback_pool"),
			LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
				Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
				PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
			}),
			Name: cloudflare.F("www.example.com"),
			PopPools: cloudflare.F(load_balancers.LoadBalancerEditParamsPopPools{
				Pop: cloudflare.F([]string{"string", "string", "string"}),
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
					Key:   cloudflare.F("key"),
					Value: cloudflare.F(0.000000),
				}),
			}),
			RegionPools: cloudflare.F(load_balancers.LoadBalancerEditParamsRegionPools{
				RegionCode: cloudflare.F([]string{"string", "string", "string"}),
			}),
			Rules: cloudflare.F([]load_balancers.RulesParam{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.RulesFixedResponseParam{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.RulesOverridesParam{
					AdaptiveRouting: cloudflare.F(load_balancers.AdaptiveRoutingParam{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F(load_balancers.RulesOverridesCountryPoolsParam{
						CountryCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					DefaultPools: cloudflare.F([]load_balancers.DefaultPoolsParam{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F("fallback_pool"),
					LocationStrategy: cloudflare.F(load_balancers.LocationStrategyParam{
						Mode:      cloudflare.F(load_balancers.LocationStrategyModePop),
						PreferECS: cloudflare.F(load_balancers.LocationStrategyPreferECSAlways),
					}),
					PopPools: cloudflare.F(load_balancers.RulesOverridesPopPoolsParam{
						Pop: cloudflare.F([]string{"string", "string", "string"}),
					}),
					RandomSteering: cloudflare.F(load_balancers.RandomSteeringParam{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F(load_balancers.RandomSteeringPoolWeightsParam{
							Key:   cloudflare.F("key"),
							Value: cloudflare.F(0.000000),
						}),
					}),
					RegionPools: cloudflare.F(load_balancers.RulesOverridesRegionPoolsParam{
						RegionCode: cloudflare.F([]string{"string", "string", "string"}),
					}),
					SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
					SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.SessionAffinityNone),
			SessionAffinityAttributes: cloudflare.F(load_balancers.SessionAffinityAttributesParam{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.SessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.SessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.SessionAffinityAttributesZeroDowntimeFailoverNone),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.SteeringPolicyOff),
			TTL:                cloudflare.F(30.000000),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		load_balancers.LoadBalancerGetParams{
			ZoneID: cloudflare.F("699d98642c564d2e855e9661899b7252"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
