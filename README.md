<div align="center">
  <img height="100px" src="/logo.svg" />
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

## How to contribute

I'm really glad you're reading this.

Submitting changes

- Fork this repository
- Create a branch in your fork
- Send a Pull Request along wih a clear description of your changes as well as specs

Always write a clear log message for your commits. One-line messages are fine for small changes, but bigger changes should look like this:

```
\$ git commit -m "A brief summary of the commit

> A paragraph describing what changed and its impact."
> Coding conventions
> Start reading our code and you'll get the hang of it. We optimize for readability:
```
