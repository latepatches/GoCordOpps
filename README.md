# GoCordOpps

GoCordOpps is a small Go-based Discord utility collection that interacts with Discord's HTTP API to automate a few repetitive actions:

- send repeated messages to a channel
- create webhooks in a channel and spam via them
- create and immediately delete a DM group with a target user

This project is intended for learning, testing, and authorized internal use only. Discord API usage can trigger rate limits, account restrictions, or moderation actions if abused. Use only on accounts and servers you own or are explicitly authorized to test.

## Repository contents

- `MessageSpammer.go` — sends an arbitrary message repeatedly to a Discord channel.
- `ChannelSpammer.go` — creates webhooks inside a channel, saves their URLs, then posts spam messages through those webhooks.
- `GroupSpammer.go` — creates a group DM with a target user and a placeholder user, then deletes the channel quickly.

## Requirements

- Go 1.18 or newer
- A valid Discord token
- A target Discord channel ID or user ID

## Setup

```bash
go mod init gocordopps
# or just run the files directly with Go
```

If needed, install Go from the official Go website and ensure it is available in your PATH.

## Usage

### 1) Message spammer

```bash
go run MessageSpammer.go
```

Prompts:

- Token
- Channel ID
- Message
- Amount

Example flow:

```text
[?] Token :
[?] Channel ID :
[?] Message :
[?] Amount :
```

This script sends a POST request to:

```text
https://discord.com/api/v9/channels/<channel_id>/messages
```

for the requested number of times.

### 2) Channel webhook spammer

```bash
go run ChannelSpammer.go
```

Prompts:

- Message
- Channel ID
- Token

What it does:

- creates several webhooks in the target channel
- collects the webhook URLs
- posts the same message repeatedly through each webhook

### 3) Group creator / leaver

```bash
go run GroupSpammer.go
```

Prompts:

- Token
- Victim ID
- Amount

This script creates a DM group using the target user and a placeholder account, then immediately deletes that channel after the API confirms creation.

## Notes

- The code uses the Discord API directly with standard Go HTTP requests.
- Requests are sent concurrently using `sync.WaitGroup` for parallel execution.
- Rate limits and Discord account protections may prevent successful spam attempts over time.
- Some scripts create temporary channels or webhooks, which may trigger moderation alerts if used improperly.

## Warning

Use this repository responsibly and only in situations where you are authorized to test or interact with the target system. Unauthorized use against other users, communities, or Discord servers may violate platform policies and local laws.

## License

This repository does not currently include a formal license file. If you are using or distributing this code, confirm the project owner's licensing requirements before publishing or reusing it publicly.
