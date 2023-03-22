package ffmpeg

import (
	"fmt"
	"strings"
)

func Build(url string, headers []string) string {
	qheaders := make([]string, len(headers))
	for i, h := range headers {
		qheaders[i] = fmt.Sprintf("'%s'", h)
	}

	joinedH := strings.Join(qheaders, "$'\\r\\n'")
	return fmt.Sprintf("ffmpeg -headers %s -i %s", joinedH, url)
}
