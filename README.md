# `up`

Provides a super simplistic uptime monitor, no fancy stuff just barebones up across various different protocols. Its aim is to be setup easily and operated without effort.

## TODOs

- [ ] Extract config path path into `ENV` vars
- [ ] Extract db path path into `ENV` vars

## Monitor Configuration

`up` uses a yaml syntax to describe its monitors.

**Example**

```
---
monitors:
  my_app_1:
    name: My App 1
    address: `string` (default = "")
    expectation: `string` (default = 200)
    protocol: `string` (default = http)
    interval: `time.Dration` (default = 5m)
```

Supported protocols are `http` (can be http or https) as well as `dns`
