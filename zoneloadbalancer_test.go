// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestZoneLoadBalancerNewWithOptionalParams(t *testing.T) {
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
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.LoadBalancers.New(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneLoadBalancerNewParams{
			DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsAdaptiveRouting{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F[any](map[string]interface{}{
				"GB": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
				},
				"US": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Description: cloudflare.F("Load Balancer for www.example.com"),
			LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsLocationStrategy{
				Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsLocationStrategyPreferEcsAlways),
			}),
			PopPools: cloudflare.F[any](map[string]interface{}{
				"LAX": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
				"LHR": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					"1": "f9138c5d07c2e2f4df57b1f61d4196",
				},
				"SJC": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRandomSteering{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F[any](map[string]interface{}{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F[any](map[string]interface{}{
				"ENAM": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
				"WNAM": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
			}),
			Rules: cloudflare.F([]cloudflare.ZoneLoadBalancerNewParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerNewParamsSteeringPolicyDynamicLatency),
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

func TestZoneLoadBalancerGet(t *testing.T) {
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
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.LoadBalancers.Get(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneLoadBalancerUpdateWithOptionalParams(t *testing.T) {
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
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.LoadBalancers.Update(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
		cloudflare.ZoneLoadBalancerUpdateParams{
			DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsAdaptiveRouting{
				FailoverAcrossPools: cloudflare.F(true),
			}),
			CountryPools: cloudflare.F[any](map[string]interface{}{
				"GB": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
				},
				"US": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Description: cloudflare.F("Load Balancer for www.example.com"),
			Enabled:     cloudflare.F(true),
			LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsLocationStrategy{
				Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcsAlways),
			}),
			PopPools: cloudflare.F[any](map[string]interface{}{
				"LAX": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
				"LHR": map[string]interface{}{
					"0": "abd90f38ced07c2e2f4df50b1f61d4194",
					"1": "f9138c5d07c2e2f4df57b1f61d4196",
				},
				"SJC": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
			}),
			Proxied: cloudflare.F(true),
			RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRandomSteering{
				DefaultWeight: cloudflare.F(0.200000),
				PoolWeights: cloudflare.F[any](map[string]interface{}{
					"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
					"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
				}),
			}),
			RegionPools: cloudflare.F[any](map[string]interface{}{
				"ENAM": map[string]interface{}{
					"0": "00920f38ce07c2e2f4df50b1f61d4194",
				},
				"WNAM": map[string]interface{}{
					"0": "de90f38ced07c2e2f4df50b1f61d4194",
					"1": "9290f38c5d07c2e2f4df57b1f61d4196",
				},
			}),
			Rules: cloudflare.F([]cloudflare.ZoneLoadBalancerUpdateParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
						FailoverAcrossPools: cloudflare.F(true),
					}),
					CountryPools: cloudflare.F[any](map[string]interface{}{
						"GB": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
						},
						"US": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
					FallbackPool: cloudflare.F[any](map[string]interface{}{}),
					LocationStrategy: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
					}),
					PopPools: cloudflare.F[any](map[string]interface{}{
						"LAX": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
						"LHR": map[string]interface{}{
							"0": "abd90f38ced07c2e2f4df50b1f61d4194",
							"1": "f9138c5d07c2e2f4df57b1f61d4196",
						},
						"SJC": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
					}),
					RandomSteering: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering{
						DefaultWeight: cloudflare.F(0.200000),
						PoolWeights: cloudflare.F[any](map[string]interface{}{
							"9290f38c5d07c2e2f4df57b1f61d4196": 0.500000,
							"de90f38ced07c2e2f4df50b1f61d4194": 0.300000,
						}),
					}),
					RegionPools: cloudflare.F[any](map[string]interface{}{
						"ENAM": map[string]interface{}{
							"0": "00920f38ce07c2e2f4df50b1f61d4194",
						},
						"WNAM": map[string]interface{}{
							"0": "de90f38ced07c2e2f4df50b1f61d4194",
							"1": "9290f38c5d07c2e2f4df57b1f61d4196",
						},
					}),
					SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(cloudflare.ZoneLoadBalancerUpdateParamsSteeringPolicyDynamicLatency),
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

func TestZoneLoadBalancerList(t *testing.T) {
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
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.LoadBalancers.List(context.TODO(), "699d98642c564d2e855e9661899b7252")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestZoneLoadBalancerDelete(t *testing.T) {
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
		option.WithEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
	)
	_, err := client.Zones.LoadBalancers.Delete(
		context.TODO(),
		"699d98642c564d2e855e9661899b7252",
		"699d98642c564d2e855e9661899b7252",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
