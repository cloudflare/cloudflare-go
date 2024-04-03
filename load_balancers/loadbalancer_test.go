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
	t.Skip("skipped: tests are disabled for the time being")
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
		DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
		FallbackPool: cloudflare.F[any](map[string]interface{}{}),
		Name:         cloudflare.F("www.example.com"),
		AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerNewParamsAdaptiveRouting{
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
		LocationStrategy: cloudflare.F(load_balancers.LoadBalancerNewParamsLocationStrategy{
			Mode:      cloudflare.F(load_balancers.LoadBalancerNewParamsLocationStrategyModeResolverIP),
			PreferEcs: cloudflare.F(load_balancers.LoadBalancerNewParamsLocationStrategyPreferEcsAlways),
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
		RandomSteering: cloudflare.F(load_balancers.LoadBalancerNewParamsRandomSteering{
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
		Rules: cloudflare.F([]load_balancers.LoadBalancerNewParamsRule{{
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
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
				LocationStrategy: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
				RandomSteering: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesRandomSteering{
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
				SessionAffinity: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
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
				LocationStrategy: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
				RandomSteering: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesRandomSteering{
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
				SessionAffinity: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}, {
			Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
			Disabled:  cloudflare.F(true),
			FixedResponse: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesFixedResponse{
				ContentType: cloudflare.F("application/json"),
				Location:    cloudflare.F("www.example.com"),
				MessageBody: cloudflare.F("Testing Hello"),
				StatusCode:  cloudflare.F(int64(0)),
			}),
			Name: cloudflare.F("route the path /testing to testing datacenter."),
			Overrides: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverrides{
				AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesAdaptiveRouting{
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
				LocationStrategy: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategy{
					Mode:      cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP),
					PreferEcs: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
				RandomSteering: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesRandomSteering{
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
				SessionAffinity: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityCookie),
				SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes{
					DrainDuration:        cloudflare.F(100.000000),
					Headers:              cloudflare.F([]string{"x"}),
					RequireAllHeaders:    cloudflare.F(true),
					Samesite:             cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
					Secure:               cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto),
					ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
				}),
				SessionAffinityTTL: cloudflare.F(1800.000000),
				SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency),
				TTL:                cloudflare.F(30.000000),
			}),
			Priority:   cloudflare.F(int64(0)),
			Terminates: cloudflare.F(true),
		}}),
		SessionAffinity: cloudflare.F(load_balancers.LoadBalancerNewParamsSessionAffinityCookie),
		SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerNewParamsSessionAffinityAttributes{
			DrainDuration:        cloudflare.F(100.000000),
			Headers:              cloudflare.F([]string{"x"}),
			RequireAllHeaders:    cloudflare.F(true),
			Samesite:             cloudflare.F(load_balancers.LoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto),
			Secure:               cloudflare.F(load_balancers.LoadBalancerNewParamsSessionAffinityAttributesSecureAuto),
			ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
		}),
		SessionAffinityTTL: cloudflare.F(1800.000000),
		SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerNewParamsSteeringPolicyDynamicLatency),
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
	t.Skip("skipped: tests are disabled for the time being")
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
			DefaultPools: cloudflare.F([]string{"17b5962d775c646f3f9725cbc7a53df4", "9290f38c5d07c2e2f4df57b1f61d4196", "00920f38ce07c2e2f4df50b1f61d4194"}),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			Name:         cloudflare.F("www.example.com"),
			AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerUpdateParamsAdaptiveRouting{
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
			LocationStrategy: cloudflare.F(load_balancers.LoadBalancerUpdateParamsLocationStrategy{
				Mode:      cloudflare.F(load_balancers.LoadBalancerUpdateParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(load_balancers.LoadBalancerUpdateParamsLocationStrategyPreferEcsAlways),
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
			RandomSteering: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRandomSteering{
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
			Rules: cloudflare.F([]load_balancers.LoadBalancerUpdateParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.LoadBalancerUpdateParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerUpdateParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.LoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerUpdateParamsSteeringPolicyDynamicLatency),
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
	t.Skip("skipped: tests are disabled for the time being")
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
	t.Skip("skipped: tests are disabled for the time being")
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
			Body:   cloudflare.F[any](map[string]interface{}{}),
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
	t.Skip("skipped: tests are disabled for the time being")
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
			AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerEditParamsAdaptiveRouting{
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
			Description:  cloudflare.F("Load Balancer for www.example.com"),
			Enabled:      cloudflare.F(true),
			FallbackPool: cloudflare.F[any](map[string]interface{}{}),
			LocationStrategy: cloudflare.F(load_balancers.LoadBalancerEditParamsLocationStrategy{
				Mode:      cloudflare.F(load_balancers.LoadBalancerEditParamsLocationStrategyModeResolverIP),
				PreferEcs: cloudflare.F(load_balancers.LoadBalancerEditParamsLocationStrategyPreferEcsAlways),
			}),
			Name: cloudflare.F("www.example.com"),
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
			RandomSteering: cloudflare.F(load_balancers.LoadBalancerEditParamsRandomSteering{
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
			Rules: cloudflare.F([]load_balancers.LoadBalancerEditParamsRule{{
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}, {
				Condition: cloudflare.F("http.request.uri.path contains \"/testing\""),
				Disabled:  cloudflare.F(true),
				FixedResponse: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesFixedResponse{
					ContentType: cloudflare.F("application/json"),
					Location:    cloudflare.F("www.example.com"),
					MessageBody: cloudflare.F("Testing Hello"),
					StatusCode:  cloudflare.F(int64(0)),
				}),
				Name: cloudflare.F("route the path /testing to testing datacenter."),
				Overrides: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverrides{
					AdaptiveRouting: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesAdaptiveRouting{
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
					LocationStrategy: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategy{
						Mode:      cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP),
						PreferEcs: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways),
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
					RandomSteering: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesRandomSteering{
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
					SessionAffinity: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityCookie),
					SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes{
						DrainDuration:        cloudflare.F(100.000000),
						Headers:              cloudflare.F([]string{"x"}),
						RequireAllHeaders:    cloudflare.F(true),
						Samesite:             cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto),
						Secure:               cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto),
						ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky),
					}),
					SessionAffinityTTL: cloudflare.F(1800.000000),
					SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency),
					TTL:                cloudflare.F(30.000000),
				}),
				Priority:   cloudflare.F(int64(0)),
				Terminates: cloudflare.F(true),
			}}),
			SessionAffinity: cloudflare.F(load_balancers.LoadBalancerEditParamsSessionAffinityCookie),
			SessionAffinityAttributes: cloudflare.F(load_balancers.LoadBalancerEditParamsSessionAffinityAttributes{
				DrainDuration:        cloudflare.F(100.000000),
				Headers:              cloudflare.F([]string{"x"}),
				RequireAllHeaders:    cloudflare.F(true),
				Samesite:             cloudflare.F(load_balancers.LoadBalancerEditParamsSessionAffinityAttributesSamesiteAuto),
				Secure:               cloudflare.F(load_balancers.LoadBalancerEditParamsSessionAffinityAttributesSecureAuto),
				ZeroDowntimeFailover: cloudflare.F(load_balancers.LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverSticky),
			}),
			SessionAffinityTTL: cloudflare.F(1800.000000),
			SteeringPolicy:     cloudflare.F(load_balancers.LoadBalancerEditParamsSteeringPolicyDynamicLatency),
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
	t.Skip("skipped: tests are disabled for the time being")
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
