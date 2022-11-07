package discordvoicego

import (
	"errors"
	"fmt"
	discordgo "github.com/Minagoroshi/discordgopluscamera"
	"layeh.com/gopus"
	"log"
)

// ReceivePCM recieves from the discordgo.VoiceConnection.OpusRecv channel.
// It then decodes the Opus data into PCM and sends it to the provided channel.
func ReceivePCM(v *discordgo.VoiceConnection, c chan *discordgo.Packet) {
	if c == nil {
		return
	}

	var err error

	for {
		if v.Ready == false || v.OpusRecv == nil {
			VConnLogger.Log(fmt.Sprintf("Not Ready to receive %+v : %+v", v.Ready, v.OpusSend), errors.New("Discordgo not ready to receive opus packets. Exitting Application"))
			log.Fatalf("Discordgo not ready to receive opus packets. %+v : %+v,  Exitting Application \n", v.Ready, v.OpusSend)
		}

		p, ok := <-v.OpusRecv
		if !ok {
			return
		}

		if speakers == nil {
			speakers = make(map[uint32]*gopus.Decoder)
		}

		_, ok = speakers[p.SSRC]
		if !ok {
			speakers[p.SSRC], err = gopus.NewDecoder(48000, 2)
			if err != nil {
				VConnLogger.Log("error creating opus decoder", err)
				continue
			}
		}

		p.PCM, err = speakers[p.SSRC].Decode(p.Opus, 960, false)
		if err != nil {
			VConnLogger.Log("Error decoding opus data", err)
			continue
		}

		c <- p
	}
}
