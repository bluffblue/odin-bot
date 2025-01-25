# Discord Moderation Bot

A powerful Discord moderation bot built with Go, featuring essential moderation commands and comprehensive logging system.

## Features

### Moderation Commands
- `/ban` - Ban a user from the server (Admin only)
- `/unban` - Unban a previously banned user (Admin only)
- `/kick` - Kick a user from the server (Admin only)
- `/warn` - Issue warnings to users (Admin only)
- `/unwarn` - Remove warnings from users (Admin only)
- `/clearmsg [amount]` - Clear specified number of messages (1-100) (Admin only)

### Security & Logging
- All commands are restricted to users with administrator permissions
- Ephemeral responses - Command responses are only visible to the user who executed them
- Comprehensive logging system that tracks:
  - Command execution details
  - User information
  - Timestamp
  - Success/Failure status
  - Action details

## Installation

1. Clone the repository
```bash
git clone <repository-url>
cd odin-bot
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
Create a `.env` file in the root directory with the following:
```env
BOT_TOKEN=your_bot_token
GUILD_ID=your_guild_id
LOG_CHANNEL_ID=your_log_channel_id
```

- `BOT_TOKEN`: Your Discord bot token from Discord Developer Portal
- `GUILD_ID`: The ID of your Discord server
- `LOG_CHANNEL_ID`: The ID of the channel where command logs will be sent

## Project Structure

```
odin-bot/
├── src/
│   ├── commands/    # Command handlers
│   │   └── commands.go
│   ├── modals/      # Command definitions and logging
│   │   ├── modals.go
│   │   └── logs.go
│   └── main.go      # Bot initialization and setup
├── .env             # Environment variables
├── .gitignore       # Git ignore file
├── go.mod           # Go modules file
└── README.md        # Project documentation
```

## Running the Bot

1. Make sure all environment variables are set in `.env`
2. Run the bot:
```bash
go run src/main.go
```

## Security Considerations

- Keep your `.env` file secure and never commit it to version control
- Only grant administrator permissions to trusted users
- All command executions are logged for security audit purposes
- Command responses are ephemeral to maintain privacy

## Logging System

The bot includes a comprehensive logging system that:
- Logs all command executions to a dedicated channel
- Provides detailed information about each command execution
- Uses formatted code blocks for better readability
- Includes success/failure status with emoji indicators
- Maintains an audit trail of all moderation actions

Log Format:
```
Command: [command_name]
User: [username] (user_id)
Channel: [channel_id]
Time: [timestamp]
Status: ✅ Success/❌ Failed
Details: [action_details]
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
