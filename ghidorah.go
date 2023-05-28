package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type GhidorahReg struct {
	Name             string `json:name`
	ExternAccessible bool   `json:extern_facing`
	Internal         bool   `json:internal`
	Port             string `json:port`
}

// initialize connection to Ghidorah
func ConnectToGhidorah(args GhidorahReg, path string) {

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(args)
	if err != nil {
		LogFatal("Could not encode Registration Data", err)
	}

	req, err := http.NewRequest(http.MethodGet, path+"/register", &buf)
	if err != nil {
		LogFatal("Could not create service registation request", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		LogFatal("Could not connect to Ghidorah", err)
	}

	if resp.StatusCode != 200 {
		LogFatal("Could not register service", resp.Status)
	}

	LogInfo("Connected to Ghidorah")
}
