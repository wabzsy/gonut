package gonut

import (
	"encoding/hex"
	"testing"
)

func TestFormatTemplate_ToBase64(t *testing.T) {
	tpl := NewFormatTemplate(GenRandomBytes(300))
	t.Logf("\n%s\n", tpl.ToBase64())
}

func TestFormatTemplate_ToBinary(t *testing.T) {
	tpl := NewFormatTemplate(GenRandomBytes(300))
	t.Logf("\n%s\n", hex.Dump(tpl.ToBinary()))
}

func TestFormatTemplate_ToCSharp(t *testing.T) {

}

func TestFormatTemplate_ToHex(t *testing.T) {
	//tpl := NewFormatTemplate(GenRandomBytes(300))
	tpl := NewFormatTemplate([]byte{0x00, 0x01, 0x02})
	t.Logf("\n%s\n", tpl.ToHex())
}

func TestFormatTemplate_ToPowerShell(t *testing.T) {

}

func TestFormatTemplate_ToPython(t *testing.T) {

}

func TestFormatTemplate_ToRubyC(t *testing.T) {
	tpl := NewFormatTemplate(GenRandomBytes(300))
	t.Logf("\n%s\n", tpl.ToRubyC())
}

func TestFormatTemplate_ToUUID(t *testing.T) {

}
