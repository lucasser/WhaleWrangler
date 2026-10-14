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
package helpers

import (
	"github.com/spf13/viper"
)

type Services struct {
	WorkingDirectory string                   `mapstructure:"working_directory"`
	CaddySystem      CaddySystemService       `mapstructure:"caddy_service"`
	Docker           map[string]DockerService `mapstructure:"docker_services"`
	Caddy            map[string]CaddyService  `mapstructure:"caddy_services"`
}

type DockerService struct {
	ProjectDir string    `mapstructure:"project_directory"`
	Enabled    bool      `mapstructure:"enabled"`
	Caddy      Caddyfile `mapstructure:"caddyfile"`
}

type CaddyService struct {
	Enabled bool      `mapstructure:"enabled"`
	Caddy   Caddyfile `mapstructure:"caddyfile"`
}

func readConfig() {

}

func getStringOrElse(v *viper.Viper, arg string, fallback string) string {
	result := v.GetString(arg)

	if result == "" {
		return fallback
	}

	return result
}

// func parseDockerConfig(v *viper.Viper) *DockerService {
// 	return &DockerService{
// 		Name:       v.GetString("name"),
// 		ProjectDir: getStringOrElse(v, "project_directory"),
// 		Enabled:    v.GetBool("enabled"),
// 		Caddyfile:  *parseCaddyConfig(v.Sub("caddyfile")),
// 	}
// }

// func parseCaddyConfig(v *viper.Viper) *CaddyService {
// 	return &CaddyService{
// 		Name:      v.GetString("name"),
// 		Enabled:   v.GetBool("enabled"),
// 		Caddyfile: *parseCaddyConfig(v.Sub("caddyfile")),
// 	}
// }
