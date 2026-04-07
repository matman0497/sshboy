# sshboy

Simple CLI for storing SSH hosts and connecting to them quickly.

## Install

```bash
go install github.com/matman0497/sshboy@latest
```

This places the `sshboy` binary in your `GOBIN` (defaults to `~/go/bin`). Make sure that directory is on your `PATH`.

## Usage

Run `sshboy init` once to create `~/.sshboy/inventory.yaml`.

Common commands:

- `sshboy add` — store a new server in the inventory.
- `sshboy list` — prints all saved servers.
- `sshboy edit <name>` — updates the host or user for an existing entry.
- `sshboy connect <name>` — opens an `ssh` session to the server.
- `sshboy ping <name>` — runs `ping` against the saved host to check reachability.
- `sshboy interactive` — launches a TUI
- `sshboy version` — prints the CLI version.

