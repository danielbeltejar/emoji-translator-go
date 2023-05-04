package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"strings"
)

type TranslatorStruct struct {
	translated string
}

func Translator(inputText string) *TranslatorStruct {
	t := &TranslatorStruct{}

	if strings.Contains(inputText, "🥵") {
		inputText = strings.ReplaceAll(strings.ReplaceAll(inputText, " ", ""), "\n", "")
		decodedText := strings.ReplaceAll(strings.ReplaceAll(inputText, "🥵", "0"), "🥶", "1")
		byteArr := make([]byte, len(decodedText)/8+1)
		for i := 0; i < len(decodedText); i += 8 {
			end := i + 8
			if end > len(decodedText) {
				end = len(decodedText)
			}
			b := decodedText[i:end]
			var val byte
			for j := 0; j < len(b); j++ {
				val = (val << 1) | (b[j] - '0')
			}
			byteArr[i/8] = val
		}
		byteArr = byteArr[:len(byteArr)-1]
		decodedTextBytes, err := zlib.NewReader(bytes.NewReader(byteArr))
		if err != nil {

		}
		defer decodedTextBytes.Close()

		var decodedTextBuffer bytes.Buffer
		decodedTextBuffer.ReadFrom(decodedTextBytes)
		t.translated = decodedTextBuffer.String()
	} else {
		encodedText := []byte(inputText)
		var compressedData bytes.Buffer
		compressor, _ := zlib.NewWriterLevel(&compressedData, zlib.BestCompression)
		compressor.Write(encodedText)
		compressor.Close()

		bitString := ""
		for _, b := range compressedData.Bytes() {
			bitString += fmt.Sprintf("%08b", b)
		}
		encodedText = []byte(strings.ReplaceAll(strings.ReplaceAll(bitString, "0", "🥵"), "1", "🥶"))
		t.translated = string(encodedText)
	}

	return t
}

func (t *TranslatorStruct) String() string {
	return t.translated
}
