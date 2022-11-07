package main

import (
	voice "DiscordEcho/voice"
	"flag"
	"fmt"
	discordgo "github.com/Minagoroshi/discordgopluscamera"
	"log"
)

func main() {

	uConfig, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	var (
		Authorization = flag.String("a", uConfig.Authorization, "Authorization")
		GID           = flag.String("g", uConfig.GID, "Guild ID")
		CID           = flag.String("c", uConfig.CID, "Channel ID")
	)
	flag.Parse()

	if *Authorization == "" {
		var a string
		log.Println("Authorization not found, Please enter your authorization token")
		_, err := fmt.Scanln(&a)
		if err != nil {
			log.Fatal(err)
		}
		Authorization = &a
	}

	voice.NewLogger()
	voice.VConnLogger.Log("Connecting to Discord", nil)
	pConn, pSession := connect(*Authorization, *GID, *CID)
	defer func() {
		pConn.Disconnect()
		pSession.Close()
	}()
	voice.VConnLogger.Log("Connected to Discord", nil)
	SpoofStream(pConn)
	go echo(pConn)

	/*	This is the code for the relay bot
		voice.VConnLogger.Log("Connecting to Discord", nil)
		lConn, lSession := connect("", *GID, "")
		defer func() {
			lConn.Disconnect()
			lSession.Close()
		}()
		voice.VConnLogger.Log("Connected to Discord", nil)*/

	// Wait for key press to exit
	fmt.Print("Press ENTER to exit")
	fmt.Scanln()
	log.Println("Exiting gracefully")

}

func echo(conn *discordgo.VoiceConnection) {

	receiveChan := make(chan *discordgo.Packet, 2)
	go voice.ReceivePCM(conn, receiveChan)

	sendChan := make(chan []int16, 2)
	go voice.SendPCM(conn, sendChan)

	conn.Speaking(true)
	defer conn.Speaking(false)
	for {
		packet, ok := <-receiveChan
		if !ok {
			voice.VConnLogger.Log("Receive channel closed", nil)
			return
		}

		sendChan <- packet.PCM
	}
}

func relay(listener, player *discordgo.VoiceConnection) {
	receiveChan := make(chan *discordgo.Packet, 2)
	go voice.ReceivePCM(listener, receiveChan)

	sendChan := make(chan []int16, 2)
	go voice.SendPCM(player, sendChan)

	player.Speaking(true)
	defer player.Speaking(false)
	for {
		packet, ok := <-receiveChan
		if !ok {
			voice.VConnLogger.Log("Receive channel closed", nil)
			return
		}

		sendChan <- packet.PCM
	}
}

func connect(authorization, gid, cid string) (*discordgo.VoiceConnection, *discordgo.Session) {

	// Connect to Discord
	discord, err := discordgo.New(authorization)
	if err != nil {
		log.Fatal(err)
	}

	// Open Websocket
	err = discord.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to voice channel.
	if gid == "" {
		var g string
		log.Println("Please enter Guild ID")
		_, err := fmt.Scanln(&g)
		if err != nil {
			log.Fatal(err)
		}
		cid = g
	}
	if cid == "" {
		var c string
		log.Println(" Please enter Channel ID")
		_, err := fmt.Scanln(&c)
		if err != nil {
			log.Fatal(err)
		}
		cid = c
	}

	vConn, err := discord.ChannelVoiceJoin(gid, cid, false, false, true)
	if err != nil {
		log.Fatal(err)
	}

	return vConn, discord
}

func SpoofStream(conn *discordgo.VoiceConnection) {
	data := voice.VoiceChannelJoinOp{18, voice.StreamConnect{
		Type:            "guild",
		GuildId:         "997653892856815686",
		ChannelId:       "1025646630315249696",
		PreferredRegion: "us-east",
	}}
	conn.Session.WsMutex.Lock()
	err := conn.Session.WsConn.WriteJSON(data)
	if err != nil {
		log.Fatal(err)
	}
	conn.Session.WsMutex.Unlock()

	data = voice.VoiceChannelJoinOp{22, voice.StreamStart{
		StreamKey: "guild:997653892856815686:1025646630315249696:1006308267292635297",
		Paused:    false,
	}}
	conn.Session.WsMutex.Lock()
	err = conn.Session.WsConn.WriteJSON(data)
	if err != nil {
		log.Fatal(err)
	}
	conn.Session.WsMutex.Unlock()

}
