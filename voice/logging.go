package discordvoicego

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

var VConnLogger VoiceLogger

type VoiceLogger struct {
	LogFile *os.File   // The file to write logs to
	Logger  log.Logger // The logger to use
}

func NewLogger() (VoiceLogger, error) {
	// make a filename using the current date and time
	filename := "discordvoicego_" + strconv.Itoa(int(time.Now().UnixNano())) + ".log"
	logFile, err := os.Create(filename)
	if err != nil {
		return VoiceLogger{}, errors.New("error creating log file")
	}

	var vLogger VoiceLogger

	vLogger.LogFile = logFile
	vLogger.Logger = *log.New(logFile, "DiscordVoiceGO ", log.LstdFlags)
	VConnLogger = vLogger

	return VConnLogger, nil
}

func (v *VoiceLogger) Log(msg string, err error) {
	if err != nil {
		v.Logger.Println(" {ERROR} "+msg, err)
	} else {
		v.Logger.Println(msg)
	}
}
