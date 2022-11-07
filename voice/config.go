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
	CHANNELS  int = STEREO              // 1 for mono, 2 for stereo
	FRAMERATE int = 48000               // audio sampling rate
	FRAMESIZE int = 960                 // uint16 size of each audio frame
	MAXBYTES  int = (FRAMESIZE * 2) * 2 // max size of opus data
)

var (
	speakers    map[uint32]*gopus.Decoder
	opusEncoder *gopus.Encoder
	mutex       sync.Mutex
)
