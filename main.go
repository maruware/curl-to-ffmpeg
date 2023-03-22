package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maruware/curl-to-ffmpeg/curl"
	"github.com/maruware/curl-to-ffmpeg/ffmpeg"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()

	c, err := curl.Parse(input)

	if err != nil {
		panic(err)
	}

	cmd := ffmpeg.Build(c.Url, c.Headers)
	fmt.Printf(cmd)
}
