**This whole setup revolves around using caddy as a reverse proxy. if you are using something else, this project is not for you.**

# Overview

## What is this?

WhaleWrangler is a way to manage many docker compose services using a config file. Quickly turning on what you need and disabling what is not needed at the moment.

## How does it work?

Services are declared in the yaml [config file](examples/config.yml). The script will read from the config file every time it runs, and use the file to manipulate compose services.

# Install

Currently WhaleWrangler is not packaged anywhere, so clone the repo, then run `go build -o whalewrangler forgejo.islandianet.net/flames/whale-wrangler`. You will get an executable which you can then run.

# Usage

## Configure your environment

### Config file

You will need a `config.yaml` file which defines the services you have and where they are located. See [examples/config.yml] for a reference.

### Caddy

You can use caddy as a docker, or native. The important thing is your caddy should source a main Caddyfile containing something allong the lines of

```caddy
https://*.{env.DOMAIN} {
	tls {
		dns cloudflare {env.CLOUDFLARE_API_TOKEN}
		resolvers [your dns server]
	}

	log {
		output stdout
	}

	import /etc/caddy/services/*.caddy
}
```

With env.DOMAIN defined in an env file as your base domain

The import statement should point to the directory defined in `config.caddy_service.services_dir`

### Compose Files

## Use the script

The main thing you will do is run `whalewrangler refresh -c [CONFIG FILE]`. This will start all enabled services, stop all disabled ones, and generate the correct caddyfiles.

you can also run `whalewrangler start or stop -c [CONFIG FILE] serviceName` to start or stop a single service.