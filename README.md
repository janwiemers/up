<div align="center">
  <img src="/logo.svg />
</div>

# `up`

Provides a super simplistic uptime monitor, no fancy stuff just barebones up across various different protocols. Its aim is to be setup easily and operated without effort.

## TODOs

- [ ] Implement Command line flag to fetch config file path

## Building and running the application

To build the server component

```
make build.server
make run.server

// or
make buildAndRun.server
```

To build the cli component

```
make build.cli
make run.cli

// or
make buildAndRun.cli
```

## Monitor Configuration

`up` uses a yaml syntax to describe its monitors.
the default path is `config.example.yaml`

**Example**

```
---
monitors:
  my_app_1:
    name: My App 1
    target: `string` (default = "")
    expectation: `string` (default = 200)
    protocol: `string` (default = http)
    interval: `time.Dration` (default = 5m)
    label: `string` (default = "")
```

Supported protocols are `http` (can be http or https) as well as `dns`
