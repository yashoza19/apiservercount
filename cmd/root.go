/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "apiservercount",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.apiservercount.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func kubeConfig() (*rest.Config, error) {
	config, err := ctrl.GetConfig()
	if err != nil {
		// returned when there is no kubeconfig
		if errors.Is(err, clientcmd.ErrEmptyConfig) {
			return nil, fmt.Errorf("please provide kubeconfig before retrying: %v", err)
		}

		// returned when the kubeconfig has no servers
		if errors.Is(err, clientcmd.ErrEmptyCluster) {
			return nil, fmt.Errorf("malformed kubeconfig. Please check before retrying: %v", err)
		}

		// any other errors getting kubeconfig would be caught here
		return nil, fmt.Errorf("error getting kubeocnfig. Please check before retrying: %v", err)
	}
	return config, nil
}
