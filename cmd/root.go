// Copyright © 2017 Farhad Farahi <farhad.farahi@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mrsrvman/orazabbix/orametrics"
	"os"
	goflag "flag"
)

var (
	cfgFile          string
	connectionString string
	zabbixHost       string
	zabbixPort       int
	hostName         string
	localFile	bool
	useRAC		bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "orazabbix",
	Short: "Oracle database monitoring on Zabbix",
	Long:  `Golang implementation of Oracle database monitoring on Zabbix`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: runCmd,
}

func runCmd(cmd *cobra.Command, args []string) {
	goflag.CommandLine.Parse([]string{})
	orametrics.Init(connectionString, zabbixHost, zabbixPort, hostName,localFile,useRAC)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orazabbix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("version", "v", false, "Prints version information")
	RootCmd.Flags().StringVarP(&connectionString, "connectionstring", "c", "system/oracle@localhost:1521/xe", "ConnectionString to the Database, Format: username/password@ip:port/sid")
	RootCmd.Flags().StringVarP(&zabbixHost, "zabbix", "z", "localhost", "Zabbix Server/Proxy Hostname or IP address")
	RootCmd.Flags().IntVarP(&zabbixPort, "port", "p", 10051, "Zabbix Server/Proxy Port")
	RootCmd.Flags().StringVarP(&hostName, "host", "H", "server1", "Hostname of the monitored object in zabbix server")
	RootCmd.Flags().BoolVarP(&localFile, "local", "l", false, "Do not send information to the server. Use local file")
	RootCmd.Flags().BoolVarP(&useRAC, "RAC", "R", false, "Do not send information to the server. Use local file")
	RootCmd.PersistentFlags().AddGoFlagSet(goflag.CommandLine)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".orazabbix" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".orazabbix")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
