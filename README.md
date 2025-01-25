# Discord Moderation Bot

A Discord moderation bot built with Go, featuring essential moderation commands for server management.

## Features

- `/ban` - Ban users with optional reason
- `/unban` - Unban previously banned users
- `/kick` - Kick users with optional reason
- `/warn` - Issue warnings to users with optional reason
- `/unwarn` - Remove warnings from users

## Requirements

- Go 1.21 or higher
- Discord Bot Token
- Server ID (Guild ID)

## Installation

1. Clone this repository:
```bash
git clone [your-repository-url]
cd [repository-name]
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create `.env` file in the root directory:
```env
BOT_TOKEN=your_discord_bot_token_here
GUILD_ID=your_guild_id_here
```

## Configuration

1. Create a new Discord Application at [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a Bot for your application
3. Enable necessary Privileged Gateway Intents:
   - SERVER MEMBERS INTENT
   - MESSAGE CONTENT INTENT
4. Copy your bot token and paste it in the `.env` file
5. Invite the bot to your server with required permissions:
   - Ban Members
   - Kick Members
   - Send Messages
   - Use Slash Commands

## Running the Bot

```bash
go run src/main.go
```

## Security

- Never commit your `.env` file
- Keep your bot token private
- Regularly rotate your bot token if compromised
- Implement proper permission checks for commands

