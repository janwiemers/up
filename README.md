<div align="center">
  <img height="100px" src="/assets/logo.svg" />
</div>

# `up`

[![Build](https://github.com/janwiemers/up/workflows/CI/badge.svg)](https://github.com/janwiemers/up/actions)

Provides a super simplistic uptime monitor, no fancy stuff just barebones up across various different protocols. Its aim is to be setup easily and operated without effort.
`up` starts a go routine for every monitor defined in the `config.yaml`. This ensures that even if one monitor crashes the other will survive.

When a monitor fails `up` will retry it the amount of times defined in `MAX_RETRY`. If it fails the defined amount of times it will mark the specific monitor as `degraded`.

## Table of Contents

- [`up`](#up)
  - [Table of Contents](#table-of-contents)
  - [Building and running the application](#building-and-running-the-application)
    - [local](#local)
    - [Docker](#docker)
  - [Monitor Configuration](#monitor-configuration)
  - [Configuration](#configuration)
    - [General](#general)
    - [Monitor configuration](#monitor-configuration-1)
    - [Email](#email)
  - [CLI](#cli)
  - [How to contribute](#how-to-contribute)

## Building and running the application

To build the server component

### local

```sh
make build.up
make run.up

// or
make buildAndRun.up
```

### Docker

```sh
make build.docker
make run.docker

// or
make buildAndRun.docker
```

## Monitor Configuration

`up` uses a yaml syntax to describe its monitors.
the default path is `config.example.yaml`

**Examples**

```yaml
---
# A Monitor that checks the availability of a specific A Record
- name: DNS Monitor
  target: 8.8.8.8
  expectation: cloudflare.com
  protocol: dns
  interval: 30s
  label: production

# A Monitor that checks if a specific TCP port is open
- name: TCP Monitor
  target: google.com
  protocol: tcp
  interval: 30s
  label: production

# A Monitor that checks if a specific HTTP endpoint is giving a given response
- name: TCP Monitor
  target: https://google.com
  protocol: http
  expectation: 200
  interval: 30s
  label: production

# A Monitor that checks if a specific HTTP endpoint is giving a given response
- name: ICMP Monitor
  target: https://google.com
  protocol: icmp
  interval: 10s
  label: production
```

| Property    | Description                         | Type           | Default  |
| ----------- | ----------------------------------- | -------------- | -------- |
| name        | The displayed name of the monitor   | `string`       | `""`     |
| target      | The URL or IP of the monitor        | `string`       | `""`     |
| expectation | Expected Answer                     | `string`       | `"200"`  |
| protocol    | Service Type                        | `string`       | `"http"` |
| interval    | Service Type                        | `time.Dration` | `5m`     |
| interval    | Service Type                        | `time.Dration` | `5m`     |
| label       | Label to add additional information | `stringn`      | `""`     |

Supported protocols are `http` (can be http or https) as well as `dns`

## Configuration

There are several variables that make the configuration of `up`

### General

| VAR              | Description                            | Default                 |
| ---------------- | -------------------------------------- | ----------------------- |
| PORT             | The port the server is starting on     | `8080`                  |
| UP_BASE_URL      | URL of the running server              | `http://localhost:8080` |
| DB_PATH          | Path to the sqlite DB                  | `./up.db`               |
| DB_CLEANUP_AFTER | Time after which the DB deletes checks | `24h`                   |

### Monitor configuration

| VAR               | Description                                          | Default                 |
| ----------------- | ---------------------------------------------------- | ----------------------- |
| MAX_RETRY         | Max retry count before marking a monitor as degraded | `3`                     |
| MONITOR_FILE_PATH | Path to load the config from                         | `./config.example.yaml` |

### Email

| VAR                        | Description                      | Default |
| -------------------------- | -------------------------------- | ------- |
| EMAIL_TO                   | Email or DL to send the email to | `""`    |
| EMAIL_SENDER_PASSWORD      | Password of the smtp login       | `""`    |
| EMAIL_SENDER_FROM          | Email of the smpt login          | `""`    |
| EMAIL_SENDER_HOST          | SMPT Host                        | `""`    |
| NOTIFICATIONS_ENABLE_EMAIL | Enable email notifications       | `false` |

### Prometheus

`up` exports prometheus compatible metrics. There are a default metrics that belong to GO as well as GIN and two custom metrics that one can use to monitor the state and amount of the monitors

| metric            | Description                                         | Default |
| ----------------- | --------------------------------------------------- | ------- |
| active_monitors   | provides a counter with the current loaded monitors | `0`     |
| degraded_monitors | provides a gauge with the current degraded monitors | `0`     |

## CLI

This repository additionally contains a terminal client to receive the data from `up`

<div align="center">
  <img src="/assets/cli.png" />
</div>

To run the CLI you'll need to a have a `~/.up` file. The file follows the yaml syntax and currently has the following options
| VAR | Description | Default | Example |
| --- | --------------------- | ------- | --------------- |
| url | host of the up server | `""` | `locahost:8080` |

```sh
make build.cli
make run.cli

// or
make buildAndRun.cli
```

## How to contribute

I'm really glad you're reading this.

Submitting changes

- Fork this repository
- Create a branch in your fork
- Send a Pull Request along wih a clear description of your changes as well as specs

Always write a clear log message for your commits. One-line messages are fine for small changes, but bigger changes should look like this:

```
$ git commit -m "A brief summary of the commit

> A paragraph describing what changed and its impact."
> Coding conventions
> Start reading our code and you'll get the hang of it. We optimize for readability:
```
