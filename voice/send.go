package discordvoicego

import (
	discordgo "github.com/Minagoroshi/discordgopluscamera"
	"layeh.com/gopus"
)

// SendPCM encodes PCM data into Opus then send that to Discordgo
func SendPCM(v *discordgo.VoiceConnection, pcmChan <-chan []int16) {
	var err error

	if pcmChan == nil {
		VConnLogger.Log("PCM channel is nil", nil)
		return
	}

	encoder, err = gopus.NewEncoder(FRAMERATE, CHANNELS, gopus.Audio)
	if err != nil {
		VConnLogger.Log("Failed to create new opus encoder", err)
		return
	}

	for {
		// read pcmChan from chan, exit if channel is closed.
		pcmData, ok := <-pcmChan
		if !ok {
			VConnLogger.Log("PCM Channel closed", nil)
			return
		}

		// try encoding pcmChan frame with Opus
		opus, err := encoder.Encode(pcmData, FRAMESIZE, MAXBYTES)
		if err != nil {
			VConnLogger.Log("Encoding Error", err)
			return
		}

		// send encoded opus pcmData to the sendOpus channel
		v.OpusSend <- opus
	}
}
