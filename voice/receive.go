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
func ReceivePCM(vConn *discordgo.VoiceConnection, packetChannel chan *discordgo.Packet) {
	var err error

	if packetChannel == nil { // if the channel is nil, return
		VConnLogger.Log("Packet Channel is nil", nil)
		return
	}

	for {
		if vConn.Ready == false || vConn.OpusRecv == nil {
			VConnLogger.Log(fmt.Sprintf("Not Ready to receive %+vConn : %+vConn", vConn.Ready, vConn.OpusSend), errors.New("Discordgo not ready to receive opus packets. Exitting Application"))
			log.Fatalf("Discordgo not ready to receive opus packets. %+vConn : %+vConn,  Exitting Application \n", vConn.Ready, vConn.OpusSend)
		}

		packet, ok := <-vConn.OpusRecv
		if !ok {
			VConnLogger.Log("OpusRecv channel closed", nil)
			return
		}

		if speakers == nil {
			speakers = make(map[uint32]*gopus.Decoder)
		}

		_, ok = speakers[packet.SSRC]
		if !ok {
			speakers[packet.SSRC], err = gopus.NewDecoder(48000, 2)
			if err != nil {
				VConnLogger.Log("error creating opus decoder", err)
				continue
			}
		}

		packet.PCM, err = speakers[packet.SSRC].Decode(packet.Opus, 960, false)
		if err != nil {
			VConnLogger.Log("Error decoding opus data", err)
			continue
		}

		packetChannel <- packet
	}
}
