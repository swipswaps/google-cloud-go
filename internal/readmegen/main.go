package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	regionStart = "[//]: # (START "

	// Regions defined in README
	supportedAPIRegion = "supported-apis"
)

func main() {
	f, err := os.Open("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	w := bytes.NewBuffer(nil)
	skipTo := ""
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if skipTo == "" {
			w.WriteString(s)
		}
		if strings.HasPrefix(s, regionStart) {
			i := strings.Index(s, "(")
			j := strings.Index(s, ")")
			// +7 for the trailing "START "
			region := s[i+7 : j]
			skipTo = strings.Replace(s, "START", "END", 1)
			log.Printf("Found region: %q", region)
			w.WriteString("\n")
			ProcessRegion(w, region)
			w.WriteString("\n")
		}
		if s == skipTo {
			skipTo = ""
			w.WriteString(s)
		}
	}
	//fmt.Println(w.String())
}

// ProcessRegion something.
func ProcessRegion(w *bytes.Buffer, region string) error {
	switch region {
	case supportedAPIRegion:
		return processSupportedAPIRegion(w)
	default:
		return fmt.Errorf("unknown region found")
	}

}
