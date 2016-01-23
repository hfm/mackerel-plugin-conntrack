mackerel-plugin-conntrack [![Build Status](https://travis-ci.org/hfm/mackerel-plugin-conntrack.svg?branch=master)](https://travis-ci.org/hfm/mackerel-plugin-conntrack)
===

nf\_conntrack custom metrics plugin for mackerel.io agent.

Synopsis
---

```sh
mackerel-plugin-conntrack [-tempfile=<tempfile>] [-version]
```

```console
$ mackerel-plugin-conntrack -h
Usage of mackerel-plugin-conntrack:
  -tempfile string
        Temp file name (default "/tmp/mackerel-plugin-conntrack")
  -version
        Print version information and quit.
```

Example of mackerel-agent.conf

```toml
[plugin.metrics.conntrack]
command = "/path/to/mackerel-plugin-conntrack"
```
