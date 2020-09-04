# Catan Service

This is the Catan service

Generated with

```
micro new --namespace=boardgame.catan --type=service catan
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: boardgame.catan.service.catan
- Type: service
- Alias: catan

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./catan-service
```

Build a docker image
```
make docker
```