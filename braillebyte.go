/*
This Package provides simple functions to encode and decode bytes to a Braille-like representation.

This is just a fun project and does not reflect a real library for converting text to readable Braille.
It is intended for educational and entertainment purposes only.
*/

package braillebyte

import "strings"

// Encode takes a slice of bytes and returns a Braille representation as a string.
// Input: in []byte - The bytes to encode.
// Output: string - The encoded Braille string.
func Encode(in []byte) string {
	var result strings.Builder

	for _, b := range in {
		if brailleChar, ok := bytetobraille[b]; ok {
			result.WriteString(brailleChar)
		} else {
			result.WriteString("_")
		}
	}

	return result.String()
}

// Decode takes a Braille string and returns a slice of bytes representing the original data.
// Input: in string - The Braille string to decode.
// Output: []byte - The decoded bytes.
func Decode(in string) []byte {
	var result []byte

	for _, brailleChar := range in {
		for b, char := range bytetobraille {
			if string(brailleChar) == char {
				result = append(result, b)
				break
			}
		}
	}

	return result
}

var bytetobraille map[byte]string = map[byte]string{
	// 0x00: "⠀", // U+2800 <-- this might be stupid if its arranged in blocks
	0x00: "_", // U+2800
	0x01: "⠁", // U+2801
	0x02: "⠂", // U+2802
	0x03: "⠃", // U+2803
	0x04: "⠄", // U+2804
	0x05: "⠅", // U+2805
	0x06: "⠆", // U+2806
	0x07: "⠇", // U+2807
	0x08: "⠈", // U+2808
	0x09: "⠉", // U+2809
	0x0A: "⠊", // U+280A
	0x0B: "⠋", // U+280B
	0x0C: "⠌", // U+280C
	0x0D: "⠍", // U+280D
	0x0E: "⠎", // U+280E
	0x0F: "⠏", // U+280F
	0x10: "⠐", // U+2810
	0x11: "⠑", // U+2811
	0x12: "⠒", // U+2812
	0x13: "⠓", // U+2813
	0x14: "⠔", // U+2814
	0x15: "⠕", // U+2815
	0x16: "⠖", // U+2816
	0x17: "⠗", // U+2817
	0x18: "⠘", // U+2818
	0x19: "⠙", // U+2819
	0x1A: "⠚", // U+281A
	0x1B: "⠛", // U+281B
	0x1C: "⠜", // U+281C
	0x1D: "⠝", // U+281D
	0x1E: "⠞", // U+281E
	0x1F: "⠟", // U+281F
	0x20: "⠠", // U+2820
	0x21: "⠡", // U+2821
	0x22: "⠢", // U+2822
	0x23: "⠣", // U+2823
	0x24: "⠤", // U+2824
	0x25: "⠥", // U+2825
	0x26: "⠦", // U+2826
	0x27: "⠧", // U+2827
	0x28: "⠨", // U+2828
	0x29: "⠩", // U+2829
	0x2A: "⠪", // U+282A
	0x2B: "⠫", // U+282B
	0x2C: "⠬", // U+282C
	0x2D: "⠭", // U+282D
	0x2E: "⠮", // U+282E
	0x2F: "⠯", // U+282F
	0x30: "⠰", // U+2830
	0x31: "⠱", // U+2831
	0x32: "⠲", // U+2832
	0x33: "⠳", // U+2833
	0x34: "⠴", // U+2834
	0x35: "⠵", // U+2835
	0x36: "⠶", // U+2836
	0x37: "⠷", // U+2837
	0x38: "⠸", // U+2838
	0x39: "⠹", // U+2839
	0x3A: "⠺", // U+283A
	0x3B: "⠻", // U+283B
	0x3C: "⠼", // U+283C
	0x3D: "⠽", // U+283D
	0x3E: "⠾", // U+283E
	0x3F: "⠿", // U+283F
	0x40: "⡀", // U+2840
	0x41: "⡁", // U+2841
	0x42: "⡂", // U+2842
	0x43: "⡃", // U+2843
	0x44: "⡄", // U+2844
	0x45: "⡅", // U+2845
	0x46: "⡆", // U+2846
	0x47: "⡇", // U+2847
	0x48: "⡈", // U+2848
	0x49: "⡉", // U+2849
	0x4A: "⡊", // U+284A
	0x4B: "⡋", // U+284B
	0x4C: "⡌", // U+284C
	0x4D: "⡍", // U+284D
	0x4E: "⡎", // U+284E
	0x4F: "⡏", // U+284F
	0x50: "⡐", // U+2850
	0x51: "⡑", // U+2851
	0x52: "⡒", // U+2852
	0x53: "⡓", // U+2853
	0x54: "⡔", // U+2854
	0x55: "⡕", // U+2855
	0x56: "⡖", // U+2856
	0x57: "⡗", // U+2857
	0x58: "⡘", // U+2858
	0x59: "⡙", // U+2859
	0x5A: "⡚", // U+285A
	0x5B: "⡛", // U+285B
	0x5C: "⡜", // U+285C
	0x5D: "⡝", // U+285D
	0x5E: "⡞", // U+285E
	0x5F: "⡟", // U+285F
	0x60: "⡠", // U+2860
	0x61: "⡡", // U+2861
	0x62: "⡢", // U+2862
	0x63: "⡣", // U+2863
	0x64: "⡤", // U+2864
	0x65: "⡥", // U+2865
	0x66: "⡦", // U+2866
	0x67: "⡧", // U+2867
	0x68: "⡨", // U+2868
	0x69: "⡩", // U+2869
	0x6A: "⡪", // U+286A
	0x6B: "⡫", // U+286B
	0x6C: "⡬", // U+286C
	0x6D: "⡭", // U+286D
	0x6E: "⡮", // U+286E
	0x6F: "⡯", // U+286F
	0x70: "⡰", // U+2870
	0x71: "⡱", // U+2871
	0x72: "⡲", // U+2872
	0x73: "⡳", // U+2873
	0x74: "⡴", // U+2874
	0x75: "⡵", // U+2875
	0x76: "⡶", // U+2876
	0x77: "⡷", // U+2877
	0x78: "⡸", // U+2878
	0x79: "⡹", // U+2879
	0x7A: "⡺", // U+287A
	0x7B: "⡻", // U+287B
	0x7C: "⡼", // U+287C
	0x7D: "⡽", // U+287D
	0x7E: "⡾", // U+287E
	0x7F: "⡿", // U+287F
	0x80: "⢀", // U+2880
	0x81: "⢁", // U+2881
	0x82: "⢂", // U+2882
	0x83: "⢃", // U+2883
	0x84: "⢄", // U+2884
	0x85: "⢅", // U+2885
	0x86: "⢆", // U+2886
	0x87: "⢇", // U+2887
	0x88: "⢈", // U+2888
	0x89: "⢉", // U+2889
	0x8A: "⢊", // U+288A
	0x8B: "⢋", // U+288B
	0x8C: "⢌", // U+288C
	0x8D: "⢍", // U+288D
	0x8E: "⢎", // U+288E
	0x8F: "⢏", // U+288F
	0x90: "⢐", // U+2890
	0x91: "⢑", // U+2891
	0x92: "⢒", // U+2892
	0x93: "⢓", // U+2893
	0x94: "⢔", // U+2894
	0x95: "⢕", // U+2895
	0x96: "⢖", // U+2896
	0x97: "⢗", // U+2897
	0x98: "⢘", // U+2898
	0x99: "⢙", // U+2899
	0x9A: "⢚", // U+289A
	0x9B: "⢛", // U+289B
	0x9C: "⢜", // U+289C
	0x9D: "⢝", // U+289D
	0x9E: "⢞", // U+289E
	0x9F: "⢟", // U+289F
	0xA0: "⢠", // U+28A0
	0xA1: "⢡", // U+28A1
	0xA2: "⢢", // U+28A2
	0xA3: "⢣", // U+28A3
	0xA4: "⢤", // U+28A4
	0xA5: "⢥", // U+28A5
	0xA6: "⢦", // U+28A6
	0xA7: "⢧", // U+28A7
	0xA8: "⢨", // U+28A8
	0xA9: "⢩", // U+28A9
	0xAA: "⢪", // U+28AA
	0xAB: "⢫", // U+28AB
	0xAC: "⢬", // U+28AC
	0xAD: "⢭", // U+28AD
	0xAE: "⢮", // U+28AE
	0xAF: "⢯", // U+28AF
	0xB0: "⢰", // U+28B0
	0xB1: "⢱", // U+28B1
	0xB2: "⢲", // U+28B2
	0xB3: "⢳", // U+28B3
	0xB4: "⢴", // U+28B4
	0xB5: "⢵", // U+28B5
	0xB6: "⢶", // U+28B6
	0xB7: "⢷", // U+28B7
	0xB8: "⢸", // U+28B8
	0xB9: "⢹", // U+28B9
	0xBA: "⢺", // U+28BA
	0xBB: "⢻", // U+28BB
	0xBC: "⢼", // U+28BC
	0xBD: "⢽", // U+28BD
	0xBE: "⢾", // U+28BE
	0xBF: "⢿", // U+28BF
	0xC0: "⣀", // U+28C0
	0xC1: "⣁", // U+28C1
	0xC2: "⣂", // U+28C2
	0xC3: "⣃", // U+28C3
	0xC4: "⣄", // U+28C4
	0xC5: "⣅", // U+28C5
	0xC6: "⣆", // U+28C6
	0xC7: "⣇", // U+28C7
	0xC8: "⣈", // U+28C8
	0xC9: "⣉", // U+28C9
	0xCA: "⣊", // U+28CA
	0xCB: "⣋", // U+28CB
	0xCC: "⣌", // U+28CC
	0xCD: "⣍", // U+28CD
	0xCE: "⣎", // U+28CE
	0xCF: "⣏", // U+28CF
	0xD0: "⣐", // U+28D0
	0xD1: "⣑", // U+28D1
	0xD2: "⣒", // U+28D2
	0xD3: "⣓", // U+28D3
	0xD4: "⣔", // U+28D4
	0xD5: "⣕", // U+28D5
	0xD6: "⣖", // U+28D6
	0xD7: "⣗", // U+28D7
	0xD8: "⣘", // U+28D8
	0xD9: "⣙", // U+28D9
	0xDA: "⣚", // U+28DA
	0xDB: "⣛", // U+28DB
	0xDC: "⣜", // U+28DC
	0xDD: "⣝", // U+28DD
	0xDE: "⣞", // U+28DE
	0xDF: "⣟", // U+28DF
	0xE0: "⣠", // U+28E0
	0xE1: "⣡", // U+28E1
	0xE2: "⣢", // U+28E2
	0xE3: "⣣", // U+28E3
	0xE4: "⣤", // U+28E4
	0xE5: "⣥", // U+28E5
	0xE6: "⣦", // U+28E6
	0xE7: "⣧", // U+28E7
	0xE8: "⣨", // U+28E8
	0xE9: "⣩", // U+28E9
	0xEA: "⣪", // U+28EA
	0xEB: "⣫", // U+28EB
	0xEC: "⣬", // U+28EC
	0xED: "⣭", // U+28ED
	0xEE: "⣮", // U+28EE
	0xEF: "⣯", // U+28EF
	0xF0: "⣰", // U+28F0
	0xF1: "⣱", // U+28F1
	0xF2: "⣲", // U+28F2
	0xF3: "⣳", // U+28F3
	0xF4: "⣴", // U+28F4
	0xF5: "⣵", // U+28F5
	0xF6: "⣶", // U+28F6
	0xF7: "⣷", // U+28F7
	0xF8: "⣸", // U+28F8
	0xF9: "⣹", // U+28F9
	0xFA: "⣺", // U+28FA
	0xFB: "⣻", // U+28FB
	0xFC: "⣼", // U+28FC
	0xFD: "⣽", // U+28FD
	0xFE: "⣾", // U+28FE
	0xFF: "⣿", // U+28FF
}
