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
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Caddyfile struct {
	Name    string       `mapstructure:"name"`
	Entries []Caddyentry `mapstructure:"params"`
}

type Caddyentry struct {
	Name      string `mapstructure:"name"`
	Ip        string `mapstructure:"server_ip"`
	Subdomain string `mapstructure:"subdomain"`
	Port      int    `mapstructure:"port"`
	Flags     string `mapstructure:"flags"`
}

type CaddySystemService struct {
	Base_dir       string `mapstructure:"base_dir"`
	Services_dir   string `mapstructure:"services_dir"`
	Reload_command string `mapstructure:"reload_command"`
}

func GenerateCaddyfile(abs_services_dir string, caddyfile Caddyfile) {
	path := fmt.Sprintf("%s%s.caddy", abs_services_dir, caddyfile.Name)
	file, err := os.Create(path)
	if err != nil {
		panic("Failed to create file: " + err.Error())
	}
	defer file.Close()

	contents := ""

	for _, entry := range caddyfile.Entries {
		contents += fmt.Sprintf("@%[1]s host %[2]s.{env.DOMAIN}\nhandle @%[1]s {\n	reverse_proxy %[3]s:%[4]d %[5]s\n}\n",
			entry.Name, entry.Subdomain, entry.Ip, entry.Port, entry.Flags)
	}

	//fmt.Print(contents)

	_, err = file.WriteString(contents)
	if err != nil {
		panic(err)
	}
}

func ReloadCaddy(config Services) {
	working_dir := config.CaddySystem.Base_dir + config.CaddySystem.Services_dir
	for _, service := range config.Docker {
		if service.Enabled {
			GenerateCaddyfile(working_dir, service.Caddy)
		} else {
			os.Remove(fmt.Sprintf("%s%s.caddy", working_dir, service.Caddy.Name))
		}
	}

	for _, service := range config.Caddy {
		if service.Enabled {
			GenerateCaddyfile(working_dir, service.Caddy)
		} else {
			os.Remove(fmt.Sprintf("%s%s.caddy", working_dir, service.Caddy.Name))
		}
	}

	reloadCaddy(config.CaddySystem)
}

func reloadCaddy(caddyService CaddySystemService) {
	// Change the current working directory to the service's home directory
	path := caddyService.Base_dir
	if err := os.Chdir(path); err != nil {
		fmt.Printf("Error changing directory to %s: %v\n", path, err)
		return
	}

	commandArray := strings.Fields(caddyService.Reload_command)
	// Execute the command
	cmdExec := exec.Command(commandArray[0], commandArray[1:]...)
	cmdExec.Stdout = os.Stdout
	cmdExec.Stderr = os.Stderr

	fmt.Printf("Reloading Caddy in %s\n", path)

	err := cmdExec.Run()
	if err != nil {
		fmt.Printf("Failed with error: %v\n", err)
	} else {
		fmt.Println("Command executed successfully.")
	}
}
