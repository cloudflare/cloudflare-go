package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/cloudflare/cloudflare-go"
	"github.com/codegangsta/cli"
)

var api *cloudflare.API

// Map type used for printing a table
type table map[string]string

// Print a nicely-formatted table
func makeTable(zones []table, cols ...string) {
	// Store the maximum length of all columns
	// The default is the length of the title
	lens := make(map[string]int)
	for _, col := range cols {
		lens[col] = len(col)
	}
	// Increase the size of the column if it is larger than the current value
	for _, z := range zones {
		for col, val := range z {
			if _, ok := lens[col]; ok && len(val) > lens[col] {
				lens[col] = len(val)
			}
		}
	}
	// Print the headings and an underline for each heading
	for _, col := range cols {
		fmt.Printf("%s%s ", strings.Title(col), strings.Repeat(" ", lens[col]-len(col)))
	}
	fmt.Println()
	for _, col := range cols {
		fmt.Printf("%s ", strings.Repeat("-", lens[col]))
	}
	fmt.Println()
	// And finally print the table data
	for _, z := range zones {
		for _, col := range cols {
			fmt.Printf("%s%s ", z[col], strings.Repeat(" ", lens[col]-len(z[col])))
		}
		fmt.Println()
	}

}

func checkEnv() error {
	if api == nil {
		var err error
		api, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		if err != nil {
			log.Fatal(err)
		}
	}

	if api.APIKey == "" {
		return errors.New("API key not defined")
	}
	if api.APIEmail == "" {
		return errors.New("API email not defined")
	}

	return nil
}

// Utility function to check if CLI flags were given.
func checkFlags(c *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if c.String(flag) == "" {
			cli.ShowSubcommandHelp(c)
			err := errors.Errorf("error: the required flag %q was empty or not provided", flag)
			fmt.Fprintln(os.Stderr, err)
			return err
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "flarectl"
	app.Usage = "Cloudflare CLI"
	app.Version = "2016.4.0"
	app.Commands = []cli.Command{
		{
			Name:    "ips",
			Aliases: []string{"i"},
			Action:  ips,
			Usage:   "Print Cloudflare IP ranges",
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "User information",
			Subcommands: []cli.Command{
				{
					Name:    "info",
					Aliases: []string{"i"},
					Action:  userInfo,
					Usage:   "User details",
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  userUpdate,
					Usage:   "Update user details",
				},
			},
		},

		{
			Name:    "zone",
			Aliases: []string{"z"},
			Usage:   "Zone information",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  zoneList,
					Usage:   "List all zones on an account",
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  zoneCreate,
					Usage:   "Create a new zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.BoolFlag{
							Name:  "jumpstart",
							Usage: "automatically fetch DNS records",
						},
						cli.StringFlag{
							Name:  "org-id",
							Usage: "organization ID",
						},
					},
				},
				{
					Name:   "check",
					Action: zoneCheck,
					Usage:  "Initiate a zone activation check",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
					},
				},
				{
					Name:    "info",
					Aliases: []string{"i"},
					Action:  zoneInfo,
					Usage:   "Information on one zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
					},
				},
				{
					Name:    "plan",
					Aliases: []string{"p"},
					Action:  zonePlan,
					Usage:   "Plan information for one zone",
				},
				{
					Name:    "settings",
					Aliases: []string{"s"},
					Action:  zoneSettings,
					Usage:   "Settings for one zone",
				},
				{
					Name:    "dns",
					Aliases: []string{"d"},
					Action:  zoneRecords,
					Usage:   "DNS records for a zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
					},
				},
				{
					Name:    "railgun",
					Aliases: []string{"r"},
					Action:  zoneRailgun,
					Usage:   "Railguns for a zone",
				},
				{
					Name:    "certs",
					Aliases: []string{"c"},
					Action:  zoneCerts,
					Usage:   "Custom SSL certificates for a zone",
				},
				{
					Name:    "keyless",
					Aliases: []string{"k"},
					Action:  zoneKeyless,
					Usage:   "Keyless SSL for a zone",
				},
			},
		},

		{
			Name:    "dns",
			Aliases: []string{"d"},
			Usage:   "DNS records",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  zoneRecords,
					Usage:   "List DNS records for a zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
					},
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  dnsCreate,
					Usage:   "Create a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "type",
							Usage: "record type",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  dnsUpdate,
					Usage:   "Update a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "create-or-update",
					Aliases: []string{"o"},
					Action:  dnsCreateOrUpdate,
					Usage:   "Create a DNS record, or update if it exists",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "name",
							Usage: "record name",
						},
						cli.StringFlag{
							Name:  "content",
							Usage: "record content",
						},
						cli.StringFlag{
							Name:  "type",
							Usage: "record type",
						},
						cli.IntFlag{
							Name:  "ttl",
							Usage: "TTL (1 = automatic)",
							Value: 1,
						},
						cli.BoolFlag{
							Name:  "proxy",
							Usage: "proxy through Cloudflare (orange cloud)",
						},
					},
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Action:  dnsDelete,
					Usage:   "Delete a DNS record",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "id",
							Usage: "record id",
						},
					},
				},
			},
		},
		{
			Name:    "user-agents",
			Aliases: []string{"ua"},
			Usage:   "User-Agent blocking",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  userAgentList,
					Usage:   "List User-Agent blocks for a zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.IntFlag{
							Name:  "page",
							Usage: "result page to return",
						},
					},
				},
				{
					Name:    "create",
					Aliases: []string{"c"},
					Action:  userAgentCreate,
					Usage:   "Create a User-Agent blocking rule",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "mode",
							Usage: "the blocking mode: block, challenge, js_challenge, whitelist",
						},
						cli.StringFlag{
							Name:  "value",
							Usage: "the exact User-Agent to block",
						},
						cli.BoolFlag{
							Name:  "paused",
							Usage: "whether the rule should be paused (default: false)",
						},
						cli.StringFlag{
							Name:  "description",
							Usage: "a description for the rule",
						},
					},
				},
				{
					Name:    "update",
					Aliases: []string{"u"},
					Action:  userAgentUpdate,
					Usage:   "Update an existing User-Agent block",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "id",
							Usage: "User-Agent blocking rule ID",
						},
						cli.StringFlag{
							Name:  "mode",
							Usage: "the blocking mode: block, challenge, js_challenge, whitelist",
						},
						cli.StringFlag{
							Name:  "value",
							Usage: "the exact User-Agent to block",
						},
						cli.BoolFlag{
							Name:  "paused",
							Usage: "whether the rule should be paused (default: false)",
						},
						cli.StringFlag{
							Name:  "description",
							Usage: "a description for the rule",
						},
					},
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Action:  userAgentDelete,
					Usage:   "Delete a User-Agent block",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
						cli.StringFlag{
							Name:  "id",
							Usage: "User-Agent blocking rule ID",
						},
					},
				},
			},
		},
		{
			Name:    "pagerules",
			Aliases: []string{"p"},
			Usage:   "Page Rules",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action:  pageRules,
					Usage:   "List Page Rules for a zone",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "zone",
							Usage: "zone name",
						},
					},
				},
			},
		},

		{
			Name:    "railgun",
			Aliases: []string{"r"},
			Usage:   "Railgun information",
			Action:  railgun,
		},

		{
			Name:    "firewall",
			Aliases: []string{"f"},
			Usage:   "Firewall",
			Subcommands: []cli.Command{
				{
					Name:    "rules",
					Aliases: []string{"r"},
					Usage:   "Access Rules",
					Subcommands: []cli.Command{
						{
							Name:    "list",
							Aliases: []string{"l"},
							Action:  firewallAccessRules,
							Usage:   "List firewall access rules",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "zone",
									Usage: "zone name",
								},
								cli.StringFlag{
									Name:  "organization",
									Usage: "organization name",
								},
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "scope-type",
									Usage: "rule scope",
								},

								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "create",
							Aliases: []string{"c"},
							Action:  firewallAccessRuleCreate,
							Usage:   "Create a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "zone",
									Usage: "zone name",
								},
								cli.StringFlag{
									Name:  "organization",
									Usage: "organization name",
								},
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "update",
							Aliases: []string{"u"},
							Action:  firewallAccessRuleUpdate,
							Usage:   "Update a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "id",
									Usage: "rule id",
								},
								cli.StringFlag{
									Name:  "zone",
									Usage: "zone name",
								},
								cli.StringFlag{
									Name:  "organization",
									Usage: "organization name",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "create-or-update",
							Aliases: []string{"o"},
							Action:  firewallAccessRuleCreateOrUpdate,
							Usage:   "Creatae a firewall access rule, or update it if it exists",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "zone",
									Usage: "zone name",
								},
								cli.StringFlag{
									Name:  "organization",
									Usage: "organization name",
								},
								cli.StringFlag{
									Name:  "value",
									Usage: "rule value",
								},
								cli.StringFlag{
									Name:  "mode",
									Usage: "rule mode",
								},
								cli.StringFlag{
									Name:  "notes",
									Usage: "rule notes",
								},
							},
						},
						{
							Name:    "delete",
							Aliases: []string{"d"},
							Action:  firewallAccessRuleDelete,
							Usage:   "Delete a firewall access rule",
							Flags: []cli.Flag{
								cli.StringFlag{
									Name:  "id",
									Usage: "rule id",
								},
								cli.StringFlag{
									Name:  "zone",
									Usage: "zone name",
								},
								cli.StringFlag{
									Name:  "organization",
									Usage: "organization name",
								},
							},
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
