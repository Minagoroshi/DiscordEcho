package discordvoicego

type VoiceChannelJoinData struct {
	GuildID    *string `json:"guild_id"`
	ChannelID  *string `json:"channel_id"`
	SelfMute   bool    `json:"self_mute"`
	SelfDeaf   bool    `json:"self_deaf"`
	SelfCamera bool    `json:"self_video"`
}

type VoiceChannelJoinOp struct {
	Op   int         `json:"op"`
	Data interface{} `json:"d"`
}

type StreamConnect struct {
	Type            string `json:"type"`
	GuildId         string `json:"guild_id"`
	ChannelId       string `json:"channel_id"`
	PreferredRegion string `json:"preferred_region"`
}

type StreamStart struct {
	StreamKey string `json:"stream_key"`
	Paused    bool   `json:"paused"`
}
