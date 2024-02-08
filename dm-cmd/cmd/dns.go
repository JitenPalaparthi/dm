package cmd

import (
	"log"

	"github.com/kitecyber/dm/dm-cmd/manager"
	"github.com/spf13/cobra"
)

var configType string
var scope string
var primaryDNS string
var secondaryDNS string
var iface string
var show string

func init() {

	dnsCmd.Flags().StringVarP(&scope, "scope", "s", "system", "two types of the scopes. system|command.command is used to set through system based commands")
	dnsCmd.Flags().StringVarP(&primaryDNS, "pd", "", "", "provide primary dns")
	dnsCmd.Flags().StringVarP(&secondaryDNS, "sd", "", "", "provide secondary dns")
	dnsCmd.Flags().StringVarP(&iface, "interface", "i", "system", "provide interfaces based on the system")

	showCmd.Flags().StringVarP(&scope, "scope", "s", "system", "two types of the scopes. system|command")
	showCmd.Flags().StringVarP(&iface, "interface", "i", "system", "provide interfaces based on the system")

	rootCmd.AddCommand(dnsCmd)
	dnsCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays dns information",
	Long:  "Displays current dns information based on interface or system level",
	Run: func(cmd *cobra.Command, args []string) {
		var idm manager.IDNSDeviceManager
		if scope == "system" {
			idm = new(manager.GlobalDNS)
			pd, sd, err := idm.GetDNS("system")
			println("Primary DNS:\t", pd, "\nSeconday DNS:\t", sd)
			if err != nil {
				log.Fatalln(err)
			}
		} else if scope == "command" {
			if iface == "" {
				log.Fatalln("interface cannot be empty")
			}
			idm = new(manager.CommandDNS)
			pd, sd, err := idm.GetDNS(iface)
			println("Primary DNS:", pd, "\nSeconday DNS:", sd)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			log.Fatalln("undefined scope.Scope can be system|command")
		}
	},
}
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "dns is to configure dns",
	Long:  `dns is to configure dns, settings to be supplied`,
	Run: func(cmd *cobra.Command, args []string) {
		if primaryDNS == "" || secondaryDNS == "" {
			log.Fatalln("primary and secondary dns ips must be given")
		}

		var idm manager.IDNSDeviceManager
		if scope == "system" {
			idm = new(manager.GlobalDNS)
			err := idm.SetDNS("", primaryDNS, secondaryDNS)
			if err != nil {
				log.Fatalln(err)
			}
			println("dns has been set")
		} else if scope == "command" {
			if iface == "" {
				log.Fatalln("interface cannot be empty")
			}
			idm = new(manager.CommandDNS)

			err := idm.SetDNS(iface, primaryDNS, secondaryDNS)
			if err != nil {
				log.Fatalln(err)
			}
			println("dns has been set")
		} else {
			log.Fatalln("undefined scope.Scope can be system|command")
		}
	},
}