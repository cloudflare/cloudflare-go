package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/jamesog/cloudflare"
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

// Utility function to check if CLI flags were given.
func checkFlags(c *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if c.String(flag) == "" {
			fmt.Println(flag, "not specified")
			return fmt.Errorf("%s not specified", flag)
		}
	}
	return nil
}

func userInfo(*cli.Context) {
	user, err := api.UserDetails()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", user)
	var output []table
	output = append(output, table{
		"ID":       user.ID,
		"Email":    user.Email,
		"Username": user.Username,
		"Name":     user.FirstName + " " + user.LastName,
		"2FA":      fmt.Sprintf("%t", user.TwoFA),
	})
	makeTable(output, "ID", "Email", "Username", "Name", "2FA")
}

func userUpdate(*cli.Context) {
}

func zoneList(c *cli.Context) {
	zones, err := api.ListZones()
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	for _, z := range zones {
		output = append(output, table{
			"ID":     z.ID,
			"Name":   z.Name,
			"Plan":   z.Plan.LegacyID,
			"Status": z.Status,
		})
	}
	makeTable(output, "ID", "Name", "Plan", "Status")
}

func zoneInfo(c *cli.Context) {
	var zone string
	if len(c.Args()) > 0 {
		zone = c.Args()[0]
	} else {
		fmt.Println("No zone specified")
		return
	}
	zones, err := api.ListZones(zone)
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	for _, z := range zones {
		output = append(output, table{
			"ID":           z.ID,
			"Zone":         z.Name,
			"Plan":         z.Plan.LegacyID,
			"Status":       z.Status,
			"Name Servers": strings.Join(z.NameServers, ", "),
			"Paused":       fmt.Sprintf("%t", z.Paused),
			"Type":         z.Type,
		})
	}
	makeTable(output, "ID", "Zone", "Plan", "Status", "Name Servers", "Paused", "Type")
}

func zonePlan(*cli.Context) {
}

func zoneSettings(*cli.Context) {
}

func zoneRecords(c *cli.Context) {
	var zone string
	if len(c.Args()) > 0 {
		zone = c.Args()[0]
	} else if c.String("zone") != "" {
		zone = c.String("zone")
	} else {
		fmt.Println("No zone specified")
		return
	}
	records, err := api.DNSRecords(zone)
	if err != nil {
		fmt.Println(err)
		return
	}
	var output []table
	for _, r := range records {
		switch r.Type {
		case "MX":
			r.Content = fmt.Sprintf("%d %s", r.Priority, r.Content)
		case "SRV":
			dp := reflect.ValueOf(r.Data).Interface().(map[string]interface{})
			r.Content = fmt.Sprintf("%.f %s", dp["priority"], r.Content)
			// CloudFlare's API, annoyingly, automatically prepends the weight
			// and port into content, separated by tabs.
			// XXX: File this as a bug. LOC doesn't do this.
			r.Content = strings.Replace(r.Content, "\t", " ", -1)
		}
		output = append(output, table{
			"ID":      r.ID,
			"Type":    r.Type,
			"Name":    r.Name,
			"Content": r.Content,
			"Proxied": fmt.Sprintf("%t", r.Proxied),
			"TTL":     fmt.Sprintf("%d", r.TTL),
		})
	}
	makeTable(output, "ID", "Type", "Name", "Content", "Proxied", "TTL")
}

func dnsCreate(c *cli.Context) {
	if err := checkFlags(c, "zone", "name", "type", "content"); err != nil {
		return
	}
	zone := c.String("zone")
	name := c.String("name")
	rtype := c.String("type")
	content := c.String("content")
	ttl := c.Int("ttl")
	proxy := c.Bool("proxy")

	record := cloudflare.DNSRecord{
		Name:    name,
		Type:    strings.ToUpper(rtype),
		Content: content,
		TTL:     ttl,
		Proxied: proxy,
	}
	err := api.CreateDNSRecord(zone, record)
	if err != nil {
		fmt.Println("Error creating DNS record:", err)
	}
}

func dnsUpdate(*cli.Context) {
}

func dnsDelete(c *cli.Context) {
	if err := checkFlags(c, "zone", "id"); err != nil {
		return
	}
	zone := c.String("zone")
	id := c.String("id")

	err := api.DeleteDNSRecord(zone, id)
	if err != nil {
		fmt.Println("Error deleting DNS record:", err)
	}
}

func zoneCerts(*cli.Context) {
}

func zoneKeyless(*cli.Context) {
}

func zoneRailgun(*cli.Context) {
}

func railgun(*cli.Context) {
}

func main() {
	// zone := cloudflare.Zone{Name: "example.com"}
	// zone.CreateZone()

	// apiKey := os.Getenv("CF_API_KEY")
	// apiEmail := os.Getenv("CF_API_EMAIL")
	api = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))

	if api.APIKey == "" {
		fmt.Println("API key not defined")
		return
	}
	if api.APIEmail == "" {
		fmt.Println("API email not defined")
		return
	}

	// userCommand := cli.Command{
	// railgunCommand := cli.Command{
	app := cli.NewApp()
	app.Name = "cloudflare"
	app.Commands = []cli.Command{
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
							Name:  "zone",
							Usage: "zone name",
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
							Usage: "proxy through CloudFlare (orange cloud)",
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
			Name:    "railgun",
			Aliases: []string{"r"},
			Usage:   "Railgun information",
			Action:  railgun,
		},
	}
	app.Run(os.Args)
}
