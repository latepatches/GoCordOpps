# GoCordOpps

Small Go scripts for testing Discord automation using the Discord HTTP API.

## Included

- `MessageSpammer.go` — sends a message to a channel repeatedly
- `ChannelSpammer.go` — creates webhooks and posts spam through them
- `GroupSpammer.go` — creates a DM group with a target user and deletes it quickly

## Requirements

- Go 1.18+
- A Discord bot token or user token
- A target channel ID or user ID

## Usage

```bash
go run MessageSpammer.go
```

```bash
go run ChannelSpammer.go
```

```bash
go run GroupSpammer.go
```

Each script will prompt for the required Discord IDs and settings.

## Warning

This project is for learning, testing, and authorized internal use only. Do not use it against servers, users, or communities without permission. Discord may rate-limit or restrict accounts for abusive behavior.

## Notes

- Uses direct Discord API requests with Go's HTTP client
- Not affiliated with Discord
- No formal license is included in this repo yet
