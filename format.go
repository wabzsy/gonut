package gonut

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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
	buffer := bytes.NewBufferString("unsigned char buf[] = \n")
	// TODO
	return buffer.Bytes()
}

func (f *FormatTemplate) ToPython() []byte {
	buffer := bytes.NewBufferString("buf   = \"\"\n")
	// TODO
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

func NewFormatTemplate(data []byte) *FormatTemplate {
	return &FormatTemplate{Data: data}
}
