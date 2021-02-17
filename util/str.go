package util

import (
	"bytes"
	"strings"
)

func StringBuilder(args ...string) string{
	var builder strings.Builder
	for _, str := range args{
		builder.WriteString(str)
	}
	return builder.String()
}

func StringBuf(args ...string) string{
	var buf bytes.Buffer
	for _, str := range args{
		buf.WriteString(str)
	}
	return buf.String()
}
