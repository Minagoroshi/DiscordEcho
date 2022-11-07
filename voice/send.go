package discordvoicego

import (
	discordgo "github.com/Minagoroshi/discordgopluscamera"
	"layeh.com/gopus"
)

// SendPCM encodes PCM data into Opus then send that to Discordgo
func SendPCM(v *discordgo.VoiceConnection, pcm <-chan []int16) {
	if pcm == nil {
		return
	}

	var err error

	opusEncoder, err = gopus.NewEncoder(FRAMERATE, CHANNELS, gopus.Audio)
	if err != nil {
		VConnLogger.Log("Failed to create new opus encoder", err)
		return
	}

	for {
		// read pcm from chan, exit if channel is closed.
		receive, ok := <-pcm
		if !ok {
			VConnLogger.Log("PCM Channel closed", nil)
			return
		}

		// try encoding pcm frame with Opus
		opus, err := opusEncoder.Encode(receive, FRAMESIZE, MAXBYTES)
		if err != nil {
			VConnLogger.Log("Encoding Error", err)
			return
		}

		// send encoded opus data to the sendOpus channel
		v.OpusSend <- opus
	}
}
