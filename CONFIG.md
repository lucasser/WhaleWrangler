## Config File Reference

The config yaml file defines a working directory, and contains an array of services.
```yaml
working_directory: base directory of docker compose projects
```
The entire setup revolves around caddy. Set it up here.

```yaml
caddy_service:
  base_dir: directory of your caddy composefile if docker, or config if native.
  services_dir: the directory the service caddyfiles will live in relative to base_dir. You will need to import the individual .caddy files from there in your Caddyfile.
  reload_command: The command to reload your caddy. if using docker, it'll be `docker exec caddy caddy reload --adapter caddyfile --config /etc/caddy/Caddyfile` if not, `caddy reload --adapter caddyfile --config [PATH TO CADDYFILE]`
```
If using caddy as a docker service, make sure to declare it as one in the next step.

There are currently two supported types of services: docker compose, and just a caddyfile

### Docker compose config params:

```yaml
Name of the service/project:
	project_directory: directory with compose file relative to working_directory
	enabled: boolean whether service should be autostarted
	caddyfile:
		name: filename of generated .caddy file.
		params: list of caddyfile entries
			name of service:
				server_ip: name of docker container or dns of server on caddy LAN
				subdomain: subdomain to use on WAN
				port: port docker container or server is listening on
```

### Caddy file only config params:

```yaml
Name of the service/project:
    enabled: boolean whether caddy route is enabled
    caddyfile:
      params:
        name: filename of generated .caddy file.
		params: list of caddyfile entries
			name of service:
				server_ip: dns of server on caddy LAN to route to
				subdomain: subdomain to use on WAN
				port: port to route to
				flags: special parameters. Will be appended as {flags} after the reverse_proxy line

```

## Generated `.caddy` file

```caddy
@NAME host SUBDOMAIN.{env.DOMAIN}
handle @NAME {
	reverse_proxy http://NAME:PORT
}
```