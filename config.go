package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type UserConfig struct {
	Authorization string `json:"authorization"`
	GID           string `json:"gid"`
	CID           string `json:"cid"`
}

func LoadConfig() (UserConfig, error) {
	// Check if config.json exists
	if _, err := os.Stat("config.json"); os.IsNotExist(err) { // If it doesn't exist, create it
		// First we need to get config data from the user
		log.Println("config.json does not exist, creating it, please enter your authorization token")
		var auth string
		_, err := fmt.Scanln(&auth)
		if err != nil {
			return UserConfig{}, errors.New("failed to read authorization token from stdin")
		}

		// Create config.json
		uConfig := UserConfig{
			Authorization: "",
			GID:           "",
			CID:           "",
		}
		uConfig.Authorization = auth
		uConfigJson, err := json.Marshal(uConfig)
		if err != nil {
			return UserConfig{}, errors.New("error marshalling config.json")
		}
		err = os.WriteFile("config.json", []byte(uConfigJson), 0644)
		if err != nil {
			return UserConfig{}, errors.New("error writing config.json")
		}
		return uConfig, nil
	} else {
		// config.json exists, load it
		data, err := os.ReadFile("config.json")
		if err != nil {
			return UserConfig{}, errors.New("error reading config.json")
		}
		uConfig := UserConfig{}
		err = json.Unmarshal(data, &uConfig)
		if err != nil {
			return UserConfig{}, errors.New("error unmarshalling config.json")
		}
		return uConfig, nil
	}
}
