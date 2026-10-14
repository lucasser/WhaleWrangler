/*
Copyright © 2026 Flames Fireclaw

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	"forgejo.islandianet.net/flames/whale-wrangler/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts specified docker services",
	Long: `Specify services to start as list of names.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		start(args)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start(services []string) {

	var config helpers.Services
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Errorf("error unmarshaling config: %w", err)
	}

	startService(config, services)
}

func startService(config helpers.Services, services []string) {
	for _, service := range services {
		if dockerService, ok := config.Docker[service]; ok {
			helpers.RunCommand(config.WorkingDirectory, "docker compose up -d", dockerService)
		} else {
			fmt.Printf("%s doesn't exist in config file", service)
		}
	}
}
