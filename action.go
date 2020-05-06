package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	eventFile := os.Getenv("GITHUB_EVENT_PATH")
	if len(eventFile) == 0 {
		failOnError(errors.New("Not in an action"))
	}
	eventRaw, err := ioutil.ReadFile(eventFile)
	failOnError(err)
	event := &GitHubEvent{}
	err = json.Unmarshal(eventRaw, event)
	failOnError(err)

	// fail the action
	failOnError(errors.New("Not implemented"))
	os.Exit(1)
}

func failOnError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e.Error())
		os.Exit(1)
	}
}
