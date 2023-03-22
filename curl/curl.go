package curl

import (
	"github.com/google/shlex"
)

type Curl struct {
	Url     string
	Headers []string
}

func Parse(input string) (*Curl, error) {
	args, err := shlex.Split(input)
	if err != nil {
		return nil, err
	}

	curl := Curl{}

	nextIsHeader := false

	for _, v := range args {
		// entry
		if v == "curl" {
			continue
		}
		// flag
		if v == "--compressed" {
			continue
		}

		// value
		if v == "-H" {
			nextIsHeader = true
			continue
		}
		if nextIsHeader {
			curl.Headers = append(curl.Headers, v)
			nextIsHeader = false
			continue
		}

		curl.Url = v
	}
	return &curl, nil
}
