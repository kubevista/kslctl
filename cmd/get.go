/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if args[0] == "webhook" {
			fmt.Println(getWebhooks())
		}

		if args[0] == "argocd-password" {
			fmt.Println(getArgoCDPassword())
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getWebhooks() string {

	topCommand := "kubectl"
	arg0 := "get"
	arg1 := "ingress"
	arg2 := "-o"
	arg3 := "jsonpath='{.items[0].status.loadBalancer.ingress[0].hostname}'"

	cmd := exec.Command(topCommand, arg0, arg1, arg2, arg3)
	hook, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	return string(hook[1 : len(hook)-1])
}

func getArgoCDPassword() string {

	topCommand := "kubectl"
	arg0 := "get"
	arg1 := "secret"
	arg2 := "argocd-initial-admin-secret"
	arg3 := "-n"
	arg4 := "argocd"
	arg5 := "-o"
	arg6 := "jsonpath='{.data.password}'"

	cmd := exec.Command(topCommand, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	hook, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	decodedPassword, err := base64.StdEncoding.DecodeString(string(hook[1 : len(hook)-1]))
	if err != nil {
		panic(err)
	}
	return string(decodedPassword)
}
