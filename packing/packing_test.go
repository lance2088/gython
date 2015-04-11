package packing

import "testing"

func TestPack(t *testing.T) {
	magic := 62212

	packed, err := Pack("HccL", magic, '\r', '\n', int64(10000))
	if err != nil {
		t.Errorf("Got error from pack %s", err)
	}

	expected := []byte{'\x04', '\xf3', '\r', '\n', '\x10', '\'', '\x00', '\x00'}
	if string(packed) != string(expected) {
		t.Errorf("%v is not equal to %v", packed, expected)
	}
}

func TestPack_Mismatch(t *testing.T) {
	_, err := Pack("Hcc", 1)
	if err == nil {
		t.Errorf("Expected error, got none")
	}
}

func TestPack_UnknownFormat(t *testing.T) {
	_, err := Pack("8")
	if err == nil {
		t.Errorf("Expected error, got none")
	}
}

func TestUnpack(t *testing.T) {
	packed := []byte{'\x04', '\xf3', '\r', '\n', '\x10', '\'', '\x00', '\x00'}
	results, err := Unpack("HccL", packed)

	if err != nil {
		t.Errorf("Got error from unpacking %s", err)
	}

	magic, linefeed, newline, num := results[0], results[1], results[2], results[3]
	if magic != int32(62212) {
		t.Errorf("Unpack magic %v != 62212", magic)
	}

	if linefeed != '\r' {
		t.Errorf("Unpack linefeed %v != '\\r'", linefeed)
	}

	if newline != '\n' {
		t.Errorf("Unpack newline %v != '\\r'", newline)
	}

	if num != int64(10000) {
		t.Errorf("Unpack num %v != 10000", num)
	}
}
