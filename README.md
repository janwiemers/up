<div align="center">
  <img height="100px" src="/logo.svg" />
</div>

# `up`

[![Build](https://github.com/janwiemers/up/workflows/CI/badge.svg)](https://github.com/janwiemers/up/actions)

Provides a super simplistic uptime monitor, no fancy stuff just barebones up across various different protocols. Its aim is to be setup easily and operated without effort.

## TODOs

- [ ] Implement Command line flag to fetch config file path

## Building and running the application

To build the server component

### local

```sh
make build.server
make run.server

// or
make buildAndRun.server
```

### Docker

```sh
make build.doker
make run.doker

// or
make buildAndRun.doker
```

### CLI

To build the cli component

```sh
make build.cli
make run.cli

// or
make buildAndRun.cli
```

## Monitor Configuration

`up` uses a yaml syntax to describe its monitors.
the default path is `config.example.yaml`

**Example**

```yaml
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
