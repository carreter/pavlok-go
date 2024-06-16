package main

import (
	"flag"
	"fmt"
	"github.com/carreter/pavlok-go"
	"os"
	"strings"
)

func main() {
	apiKey := ""
	reason := ""
	stimulusValue := 0

	flag.StringVar(&apiKey, "apikey", "", "api key, required")
	flag.StringVar(&reason, "reason", "", "reason for the stimulus")
	flag.IntVar(&stimulusValue, "value", 50, "stimulus strength value")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <stimulus type>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if apiKey == "" {
		fmt.Println("Please specify an api key")
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("Invalid number of arguments (got: %v, want: 1)\n", len(args))
		flag.Usage()
		os.Exit(1)
	}

	stimulusType := pavlok.StimulusType(strings.ToLower(args[0]))
	if stimulusType != pavlok.Zap && stimulusType != pavlok.Beep && stimulusType != pavlok.Vibe {
		fmt.Printf("Invalid stimulus type: %s. Acceptable stimulus types are: zap, beep, vibe\n", stimulusType)
	}

	stimulus := pavlok.Stimulus{
		Type:   stimulusType,
		Value:  stimulusValue,
		Reason: reason,
	}

	c := pavlok.NewClient(apiKey)

	err := c.SendStimulus(stimulus)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
