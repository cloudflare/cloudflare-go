package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func certificatePackCreate(c *cli.Context) error {
	zoneName := c.String("zone")
	if zoneName == "" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: the required flag \"zone\" was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	hosts := c.String("hosts")
	if hosts == "" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: the required flag \"hosts\" was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	validationMethod := c.String("validation")
	if validationMethod != "txt" && validationMethod != "http" && validationMethod != "email" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: invalid validation method")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	validityDays := c.Int("validity-days")

	if validityDays != 14 && validityDays != 30 && validityDays != 90 && validityDays != 365 {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: invalid validity days")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	cloudflareBranding := c.String("authority")
	if cloudflareBranding != "lets_encrypt" && cloudflareBranding != "google" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: invalid Certificate Authority")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	certificatePackRequest := cloudflare.CertificatePackRequest{
		Type:                 "advanced",
		Hosts:                strings.Split(hosts, ","),
		ValidationMethod:     validationMethod,
		ValidityDays:         validityDays,
		CloudflareBranding:   c.Bool("cloudflare-branding"),
		CertificateAuthority: cloudflareBranding,
	}

	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	record, err := api.CreateCertificatePack(context.Background(), zoneID, certificatePackRequest)
	if err != nil {
		fmt.Println(err)
		return err
	}
	output := make([][]string, 0, 1)
	expiresOn := ""
	if len(record.Certificates) > 0 {
		expiresOn = record.Certificates[0].ExpiresOn.Format("2006-01-02")
	}
	output = append(output, []string{
		record.ID,
		record.Type,
		record.Status,
		record.ValidationMethod,
		fmt.Sprintf("%d", record.ValidityDays),
		record.CertificateAuthority,
		expiresOn,
		strings.Join(record.Hosts, " "),
	})
	writeTable(c, output, "ID", "Type", "status", "validation_method", "validity_days", "certificate_authority", "expires_on", "hosts")
	return nil
}

func certificatePackDelete(c *cli.Context) error {
	zoneName := c.String("zone")
	if zoneName == "" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: the required flag \"zone\" was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	certificatePackId := c.String("id")
	if certificatePackId == "" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: the required flag \"id\" was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = api.DeleteCertificatePack(context.Background(), zoneID, certificatePackId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func certificatePacksList(c *cli.Context) error {
	zoneName := c.String("zone")
	if zoneName == "" {
		cli.ShowSubcommandHelp(c)
		err := errors.New("err: the required flag \"zone\" was empty or not provided")
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	records, err := api.ListCertificatePacks(context.Background(), zoneID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var output [][]string
	for _, r := range records {
		if c.String("id") != "" && r.ID != c.String("id") {
			continue
		}
		expiresOn := ""
		if len(r.Certificates) > 0 {
			expiresOn = r.Certificates[0].ExpiresOn.Format("2006-01-02")
		}
		output = append(output, []string{
			r.ID,
			r.Type,
			r.Status,
			r.ValidationMethod,
			fmt.Sprintf("%d", r.ValidityDays),
			r.CertificateAuthority,
			expiresOn,
			strings.Join(r.Hosts, " "),
		})
	}
	writeTable(c, output, "ID", "Type", "status", "validation_method", "validity_days", "certificate_authority", "expires_on", "hosts")
	return nil
}
