config file defines list of services.

everything takes a config file param

### subcommands:

disable: takes a list of services and marks them as disabled

enable: takes a list of services and marks them as enabled

refresh: starts enabled services and stops disabled services

restart: restarts all running services?

start: starts specified services

stop: stops specified/all services

update: updates specified/all services according to some config params


## pseudocode

keep it as a viper config, iterate thru that

```go
services[] = args. [] for all.

all = services.empty

for (service : Viper.services) {
	if (all || services.contains(service)) {
		action(service)
	}
}

action(service, cmd) {
	cd service.homedir
	exec cmd
	get and output logs?
}
```