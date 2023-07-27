package gonut

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
)

type FormatTemplate struct {
	Data []byte
}

func (f *FormatTemplate) ToBinary() []byte {
	return f.Data
}

func (f *FormatTemplate) ToBase64() []byte {
	return []byte(base64.StdEncoding.EncodeToString(f.Data))
}

func (f *FormatTemplate) ToHex() []byte {
	return []byte(hex.EncodeToString(f.Data))
}

func (f *FormatTemplate) ToRubyC() []byte {
	buffer := bytes.NewBufferString("unsigned char buf[] = ")
	rows := Convert1d2d(f.Data, 16)
	for _, row := range rows {
		buffer.WriteString("\n\"")
		for _, c := range row {
			buffer.WriteString(fmt.Sprintf("\\x%02x", c))
		}
		buffer.WriteString("\"")
	}

	buffer.WriteString(";\n")
	return buffer.Bytes()
}

func (f *FormatTemplate) ToPython() []byte {
	buffer := bytes.NewBufferString("buf =  b\"\"")
	rows := Convert1d2d(f.Data, 16)
	for _, row := range rows {
		buffer.WriteString("\nbuf += b\"")
		for _, c := range row {
			buffer.WriteString(fmt.Sprintf("\\x%02x", c))
		}
		buffer.WriteString("\"")
	}

	buffer.WriteString("\n")
	return buffer.Bytes()
}

func (f *FormatTemplate) ToPowerShell() []byte {
	buffer := bytes.NewBufferString("[Byte[]] $buf = ")
	// TODO
	return buffer.Bytes()
}

func (f *FormatTemplate) ToCSharp() []byte {
	buffer := bytes.NewBufferString(fmt.Sprintf("byte[] my_buf = new byte[%d] {\n", len(f.Data)))
	// TODO
	return buffer.Bytes()
}

func (f *FormatTemplate) ToUUID() []byte {
	buffer := bytes.NewBufferString("TODO")
	// TODO
	return buffer.Bytes()
}

func Convert1d2d(data []byte, n int) [][]byte {
	length := len(data)
	count := int(math.Ceil(float64(length) / float64(n)))
	rows := make([][]byte, count)
	for i := 0; i < count; i++ {
		start := i * n
		end := start + n
		if end > length {
			end = length
		}
		rows[i] = data[start:end]
	}
	return rows
}

func NewFormatTemplate(data []byte) *FormatTemplate {
	return &FormatTemplate{Data: data}
}
