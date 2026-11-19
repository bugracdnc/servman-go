# servman

[![Tests](https://github.com/bugracdnc/servman-go/actions/workflows/test.yml/badge.svg)](https://github.com/bugracdnc/servman-go/actions/workflows/test.yml)

A tiny CLI for managing [runit](http://smarden.org/runit/) services (Void Linux, antiX, Artix-runit…) without remembering which symlink goes where.

```console
$ servman list
Listing services: enabled, disabled
01) NetworkManager      # green: enabled
02) cronie              # blue:  enabled, but marked down on boot
03) nginx               # red:   installed, not enabled

$ servman enable nginx
nginx enabled.
```

## Commands

| Command | Alias | What it does |
|---|---|---|
| `servman list [-e\|-d]` | `l` | List services; `-e` only enabled, `-d` only disabled |
| `servman search <text> [-e\|-d]` | `se` | List services whose name contains `text` |
| `servman enable <name>` | `en` | Symlink `/etc/sv/<name>` into `/var/service` |
| `servman disable <name>` | `d` | Remove the symlink from `/var/service` |
| `servman down <name>` | | Keep it enabled but don't start it on boot (creates `down`) |
| `servman up <name>` | | Start it on boot again (removes `down`) |

Names are case-insensitive. Changes run through `sudo`.

## Install

```sh
git clone https://github.com/bugracdnc/servman-go
cd servman-go
make install    # builds and copies servman to /usr/bin (sudo)
```

Requires Go 1.23+. Shell completion: `servman completion bash|zsh|fish`.

## Development

```sh
go test ./...   # tests run against a fake /etc/sv + /var/service in a temp dir
```
