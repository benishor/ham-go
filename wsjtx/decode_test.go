package wsjtx_test

import (
	"testing"

	"github.com/tzneal/ham-go/adif"
	"github.com/tzneal/ham-go/wsjtx"
)

func TestDecodeLoggedADIF(t *testing.T) {
	msg, err := wsjtx.Decode([]byte{0xad, 0xbc, 0xcb, 0xda, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x06, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x00, 0x00, 0x01, 0x31, 0x0a, 0x3c, 0x61, 0x64, 0x69, 0x66, 0x5f, 0x76, 0x65, 0x72, 0x3a, 0x35, 0x3e, 0x33, 0x2e, 0x31, 0x2e, 0x30, 0x0a, 0x3c, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x69, 0x64, 0x3a, 0x36, 0x3e, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x0a, 0x3c, 0x45, 0x4f, 0x48, 0x3e, 0x0a, 0x3c, 0x63, 0x61, 0x6c, 0x6c, 0x3a, 0x36, 0x3e, 0x57, 0x42, 0x36, 0x46, 0x57, 0x53, 0x20, 0x3c, 0x67, 0x72, 0x69, 0x64, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x3a, 0x34, 0x3e, 0x44, 0x4d, 0x31, 0x32, 0x20, 0x3c, 0x6d, 0x6f, 0x64, 0x65, 0x3a, 0x33, 0x3e, 0x46, 0x54, 0x38, 0x20, 0x3c, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x3a, 0x33, 0x3e, 0x2b, 0x30, 0x37, 0x20, 0x3c, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x63, 0x76, 0x64, 0x3a, 0x33, 0x3e, 0x2d, 0x30, 0x32, 0x20, 0x3c, 0x71, 0x73, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x38, 0x3e, 0x32, 0x30, 0x32, 0x30, 0x30, 0x34, 0x32, 0x33, 0x20, 0x3c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x6e, 0x3a, 0x36, 0x3e, 0x30, 0x32, 0x31, 0x35, 0x34, 0x35, 0x20, 0x3c, 0x71, 0x73, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x3a, 0x38, 0x3e, 0x32, 0x30, 0x32, 0x30, 0x30, 0x34, 0x32, 0x33, 0x20, 0x3c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x3a, 0x36, 0x3e, 0x30, 0x32, 0x31, 0x36, 0x34, 0x35, 0x20, 0x3c, 0x62, 0x61, 0x6e, 0x64, 0x3a, 0x33, 0x3e, 0x32, 0x30, 0x6d, 0x20, 0x3c, 0x66, 0x72, 0x65, 0x71, 0x3a, 0x39, 0x3e, 0x31, 0x34, 0x2e, 0x30, 0x37, 0x36, 0x39, 0x33, 0x37, 0x20, 0x3c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x69, 0x67, 0x6e, 0x3a, 0x35, 0x3e, 0x57, 0x34, 0x54, 0x4e, 0x4c, 0x20, 0x3c, 0x6d, 0x79, 0x5f, 0x67, 0x72, 0x69, 0x64, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x3a, 0x36, 0x3e, 0x45, 0x4d, 0x36, 0x34, 0x4f, 0x52, 0x20, 0x3c, 0x74, 0x78, 0x5f, 0x70, 0x77, 0x72, 0x3a, 0x33, 0x3e, 0x32, 0x35, 0x57, 0x20, 0x3c, 0x45, 0x4f, 0x52, 0x3e})
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	_, ok := msg.(*wsjtx.LoggedADIF)
	if !ok {
		t.Errorf("expected an ADIF message")
	}
}
func TestDecodeQSO(t *testing.T) {
	msg, err := wsjtx.Decode([]byte{0xad, 0xbc, 0xcb, 0xda, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x06, 0x57, 0x53, 0x4a, 0x54, 0x2d, 0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x85, 0x53, 0x00, 0x7d, 0x33, 0x38, 0x01, 0x00, 0x00, 0x00, 0x06, 0x57, 0x42, 0x36, 0x46, 0x57, 0x53, 0x00, 0x00, 0x00, 0x04, 0x44, 0x4d, 0x31, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd6, 0xcc, 0x09, 0x00, 0x00, 0x00, 0x03, 0x46, 0x54, 0x38, 0x00, 0x00, 0x00, 0x03, 0x2b, 0x30, 0x37, 0x00, 0x00, 0x00, 0x03, 0x2d, 0x30, 0x32, 0x00, 0x00, 0x00, 0x03, 0x32, 0x35, 0x57, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x25, 0x85, 0x53, 0x00, 0x7c, 0x4a, 0x5a, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x57, 0x34, 0x54, 0x4e, 0x4c, 0x00, 0x00, 0x00, 0x06, 0x45, 0x4d, 0x36, 0x34, 0x4f, 0x52, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	qlog := msg.(*wsjtx.QSOLogged)
	exp := "20200423"
	got := adif.UTCDate(qlog.QSOOff)
	if got != exp {
		t.Errorf("expected date = %s, got %s", exp, got)
	}
	exp = "0216"
	got = adif.UTCTime(qlog.QSOOff)
	if got != exp {
		t.Errorf("expected time = %s, got %s", exp, got)
	}
}
