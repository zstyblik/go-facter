# go-facter

go-facter is a loose implementation of Puppet Labs [facter] in golang. The main target are platforms where there isn't possible or feasible to install Ruby, eg. [CoreOS]. Also, you can run it in Docker Container and still get data from the Host itself.

In theory, go-facter can be used as a library of sort to build custom facts.

## Licence

BSD 3-Clause ("BSD New" or "BSD Simplified") licence.

## Environment variables

- `HOST_ETC` - specify alternative path to `/etc` directory
- `HOST_PROC` - specify alternative path to `/proc` mountpoint
- `HOST_SYS` - specify alternative path to `/sys` mountpoint

## Requirements

- go v1.5 or newer is required

## Build

```
go get github.com/KittenConnect/go-facter/...
cd ~/go/src/github.com/KittenConnect/go-facter
go build ./cmd/facter
```

[facter]: https://github.com/puppetlabs/facter
[CoreOS]: https://coreos.com/
