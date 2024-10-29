package braillebyte

import (
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    []byte
		expected string
	}{
		{[]byte{0x01}, "⠁"},
		{[]byte{0x02, 0x03}, "⠂⠃"},
		{[]byte{0x00, 0x05}, "_⠅"},
		{[]byte{0xFF}, "⣿"},
	}

	for _, test := range tests {
		result := Encode(test.input)
		if result != test.expected {
			t.Errorf("Encode(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}

	t.Log(Encode([]byte("Hallo Welt")))
	t.Log(string(Decode("⡈⡡⡬⡬⡯⠠⡗⡥⡬⡴")))
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected []byte
	}{
		{"⠁", []byte{0x01}},
		{"⠂⠃", []byte{0x02, 0x03}},
		{"_⠅", []byte{0x00, 0x05}},
		{"⣿", []byte{0xFF}},
	}

	for _, test := range tests {
		result := Decode(test.input)
		if string(result) != string(test.expected) {
			t.Errorf("Decode(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Encode([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Decode("⠁⠂⠃⠄⠅")
	}
}
