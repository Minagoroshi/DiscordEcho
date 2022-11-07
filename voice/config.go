package discordvoicego

import (
	"layeh.com/gopus"
	"sync"
)

const (
	MONO   = 1
	STEREO = 2
)

const (
	CHANNELS  int = STEREO // dual channel
	FRAMERATE int = 48000  // 48kHz
	FRAMESIZE int = 960
	MAXBYTES  int = (FRAMESIZE * 2) * 2
)

var (
	speakers map[uint32]*gopus.Decoder
	encoder  *gopus.Encoder
	mutex    sync.Mutex
)
