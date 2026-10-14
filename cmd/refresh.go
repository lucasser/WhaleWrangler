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

	"forgejo.islandianet.net/flames/docker-manager/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// refreshCmd represents the refresh command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Starts all enabled services, and stops all disabled ones; cleans up caddyfiles",
	Long: `For each service: if enabled start it, if caddyfile not present generate it;
	if disabled stop it and remove caddyfile.
	restart caddy service`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		pull, err := cmd.Flags().GetBool("pull")
		if err != nil {
			fmt.Println(err)
			return;
		}
		refresh(pull)
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// refreshCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	refreshCmd.Flags().BoolP("pull", "p", false, "Pull updated images")
}

func refresh(pull bool) {

	var config helpers.Services
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Errorf("error unmarshaling config: %w", err)
	}

	fmt.Println()

	helpers.ReloadCaddy(config)

	startStopDockerServices(config, pull)
}

func startStopDockerServices(config helpers.Services, pull bool) {
	startCommand := "docker compose up -d"
	if pull {
		startCommand = "docker compose up --pull always -d"
	}
	for service, dockerService := range config.Docker {
		fmt.Println(service)
		switch dockerService.Enabled {
		case true:
			helpers.RunCommand(config.WorkingDirectory, startCommand, dockerService)
			break
		case false:
			helpers.RunCommand(config.WorkingDirectory, "docker compose down", dockerService)
			break
		default:
			fmt.Printf("Service %s has no 'Enabled' parameter\n", service)
		}
	}
}