# DiscordVCEcho

## Usage

### Flags (Not Required)
```bash
 -a string
        Authorization
  -c string
        Channel ID 
  -g string
        Guild ID 
        
 Example Usage: ./DiscordVCEcho -a "Bot Token" -g "Guild ID" -c "Channel ID"       

```
Comes with built in CLI Prompts if flag data is not provided. Flag data is not required. If you want a static Channel ID and Guild ID, you can specify them in the config.json. If you want to use the CLI prompts, you can leave the flags, and cid & gid config fields blank.

Credits: 

https://discord.com/developers/docs/topics/opcodes-and-status-codes

https://github.com/bwmarrin/discordgo
