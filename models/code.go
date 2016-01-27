package models

import (
	"bufio"
	"os"
)

// Codes struct contains list of codes
type Codes struct {
	Data []string
}

// Init returns list of passcodes
func (m *Codes) Init() {
	configFile, _ := os.Open("config/codes.txt")
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m.Data = append(m.Data, scanner.Text())
	}
}

// Find returns result of searching in slice
func (m *Codes) Find(passcode string) bool {
	for _, code := range m.Data {
		if code == passcode {
			return true
		}
	}

	return false
}
