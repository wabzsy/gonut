package gonut

import (
	"encoding/binary"
	"testing"
)

// ******************************
// test vectors for SPECK-64/128
var (
	// 128-bit key
	key64 = [16]uint8{0x00, 0x01, 0x02, 0x03, 0x08, 0x09, 0x0a, 0x0b, 0x10, 0x11, 0x12, 0x13, 0x18, 0x19, 0x1a, 0x1b}
	// 64-bit plain text
	plain64 = [8]uint8{0x74, 0x65, 0x72, 0x3b, 0x2d, 0x43, 0x75, 0x74}
	// 64-bit cipher text
	cipher64 = [8]uint8{0x48, 0xa5, 0x6f, 0x8c, 0x8b, 0x02, 0x4e, 0x45}
)

// ******************************
// test vectors for Maru hash
//

// maru_tv_t
// donut/hash.c
// typedef struct _maru_tv_t { ... }
type maru_tv_t struct {
	str  string
	hash uint64
}

// maru_tv
// donut/hash.c
// maru_tv_t maru_tv[MARU_MAX_STR] = { ... }
var maru_tv = [MARU_MAX_STR]maru_tv_t{
	{"", 0x8E63EC0D29F27D07},
	{"C", 0x19C7DC40E602AC8E},
	{"73", 0x5197B6ACC87EF423},
	{"NY4", 0x3BC2F21615A953C5},
	{"X9TM", 0xC9EC6B72BF5273D6},
	{"H339P", 0x6B60077EF084C1E2},
	{"TMCT3N", 0x33374AA7206F00FC},
	{"RF4M66W", 0xF7B91D9C42A886C5},
	{"XTCX43NN", 0x615D4FB7A2246376},
	{"C6XCYXF9F", 0x80D4B6324A24CEB6},
	{"RR3TN69H9M", 0xE6369CFF4F98B4F8},
	{"F9C9YNTMYYR", 0xF173A1158A4D80A9},
	{"FPW779364RYH", 0x517A4E86DF00BB97},
	{"WHN4N9CT7YF7C", 0xFCBA9541CD7765A5},
	{"633H6CTRC64FWR", 0x79EEC9CC663EDDC1},
	{"WPNX66993HPWNYX", 0x139CFA0D49AF17DC},
	{"H66C3Y9F677WP96N", 0xEFF27A644D53171A},
	{"TY3YX7N3FPN7YNWT4", 0x5361C6DBF89D0B47},
	{"YF496N7XH4HYHRN6WM", 0x71451CE666D8E9A4},
	{"TWP4M739RYTCTFMMCC7", 0xC17E5C46E2BEAD},
	{"4FHNWP4MR9TT9Y6HYWCX", 0x1E40C5A64B8ECE85},
	{"TCMHT3TF7T4TRCWCF6RPF", 0xD02290F438AA84A9},
	{"4WW63CTHPR36MN7P3WXTHT", 0x79FBDCFC2ECE09FF},
	{"NFFMNM3CF3NXY6P9MCC7YPX", 0x10B7C56D102D623B},
	{"R74YN9MX7PMP364HYNYCR9FY", 0x86EC8AA614611458},
	{"94X3NFT7W4FPTX3MCTY99HMPR", 0x7929169892B04FC1},
	{"66R379FR67W7T7H79WTCF37H6Y", 0xBEA85FD3754045D8},
	{"4F6HFPT3NRN7WPP766RFCXR43RX", 0x76E410266A6830A},
	{"HP374TWWPMYRTTWC6Y6T4C4T4HP4", 0x6A7509443FF48F74},
	{"TTX3966P63XPYMPM6XM994TX9X9X3", 0xA8AFC37C137AE14},
	{"34TX7XRX4WH7T6PW439TNRY77FHPT4", 0xE64667746B53394},
	{"P3WRCCT4PXN3H6PMNR3YXY6X379MTXH", 0xCE981B296791D5C7},
	{"73YF99H9XXTYC6XF6CXTPCM4YXYN33R4", 0xE39B1FC51BED4BA9},
	{"T7R3PR43C93MH4TRT6M644T7RCXMX4WM6", 0xE25C6B39CB28FDB5},
	{"44M69YYFX3C9H9M6P46933PW34RRCM9NXX", 0x84DBD675600E871E},
	{"7TW7P76CXMFC3HFTFHMWXWW33TWTPT6PYP3", 0xE00F660E699F9231},
	{"6CH6W7WYHH7HXT7RTMW4FRCN39HR997F6FWN", 0x32AEC917C63A878E},
	{"WF3HNPT37XPPXYFXR447F7RWF3C69H74CT6R6", 0x4B4FCD2604496365},
	{"WR9CCHH9NNNXCXYMXMFFW6YYC7449M4HYXM4MC", 0x330DA5A18A58C952},
	{"C9993Y93PTWYRNP46PYN763RNFYP4PN4WWHR9CM", 0x5249D03CD52C71DB},
	{"XNPXCY47FWWMTFF6R7RWNX79MC4YN43M9RYCC4RX", 0xF11E92E74F70C4F4},
	{"3RPX6MFNCWPMPT3M467HWCYCHH9PM7R6MFWXYXNTN", 0x86C8BD9789AA71BE},
	{"633MWRYMCYF3TTCRP6HXTR9TTX43T3MYXPPWMWFMF6", 0xA5F12FC7418615CC},
	{"RW7NN6Y4779639NNTHN6TR939F3799FNWTH4FP46RNY", 0x37EADA4549B0B96F},
	{"MCTRH9PTXRXFPMYRP9Y4TC4PCRX4W9YFCW649R3YN33W", 0x8E5BADDF84BB9779},
	{"F7RT3H36PNH6CF9TWCRNYMFWX9M9MTTNY6C7X43HC4PN9", 0x8040D317E8DCD294},
	{"7HTCNCTNHRFY6HCRTYPTHYP3H9T4PR96RWRY7NRPTMH936", 0x151FB43ECC51AFA1},
	{"NYRMPMYCWXTPRCYM9TX6Y4XTMY9RT3PHH49PP36H7XR7WPN", 0x650A8724A052DFC5},
	{"6PX3XWM3X973693M4F373R3N9FNC4CCTXN3CTTYMP3NW4C49", 0x9E48D8154522BFDD},
	{"74RNX9MTPW7FNY9WMTXNPPRMR97PRPPRCN3CMHPNWFFW44R3C", 0xDC1ABDA05084DCBA},
	{"47NC63XRTXXPWHN76H9XF9R7TTHWR6T7XMF9TMCHP4FX4WCYTT", 0x5DC075A21ECF2DD8},
	{"TYMXHXC4N6XXTR4T7X37PHWTYXFF9M7MXP6477RW4FM7P9PFXFR", 0x2CD151D5D71FA785},
	{"MRPW7NCXPT4N7YN3WN7P9WYNY3PPR464WR7P7PP37MXFF9FC7WTH", 0x7B88469D5AFE14D9},
	{"9XW6RYX6NTYC4NCRR7YRTWM7HWNFXRT4P396CYMFPRNTRW3X69R39", 0x81E069528C3C9BEE},
	{"9NMRWF4W34MHWTPP74RY34YWMT94H76HTRX34MR7C9MF696M3TXMN3", 0x4D19A0CB3BC48BFF},
	{"HF6FM9RMT3NPMR37TX3FPTFYRFNXTMHWTF7WN94YNP4TMP3FNHM3N9F", 0x30CEBE63BE4E30F1},
	{"NW7CCWFTTNFPMTY3H6X96HX6MXY67W3RPTRCCWHWPYPC7PPRF74PH7RC", 0x64F3DF1E551B22BE},
	{"HFH43PM9TNCCW79XCMW79HYCN4HY6MT9MFFRYRXYX4H3P9T9FHF6NWC3C", 0xEF36678895FBB3A8},
	{"N7WH9WYMNHYY3C3RRFTW3RNYH3C646C97FTPT3MH7TMW6MTC4PT44NWCWH", 0x1B75D90E82D98E1D},
	{"663F4T7PMWN996R9FYWRY3Y33HCNFH6PRWF9TPHN363YFFF6C9CHTP3XNXP", 0x25767AD747B833D6},
	{"3P7934TX6CFHPM6TWY6H4CXT47P4XRMFTPNMCFP9H9F4MPFWWF9XRMPHFCYX", 0x1F3E15CB56A60E93},
	{"WW6YN7NXTH9TRT4PYW9W3WTNP9XMHP6Y3NPX7R93Y763M9HRHWTN93W3M9WX3", 0x744735578C4F6EF2},
	{"HT6R4P6P7T4YFYYX3H3F49XYMPCPMWNT6R3PHTM47PTHTRCN9XMFHHYTH7TMPY", 0x559EA0D5309795E6},
	{"NHP9Y96YYF44H7NN33WYYC364CY3W4FNF6F7WTHN6WFF6RXXRWNRFF4T9XF934N", 0xBE7F06CC36982F52},
}

func TestMaru(t *testing.T) {
	iv := binary.LittleEndian.Uint64(cipher64[:])
	for _, item := range maru_tv {
		if item.hash != Maru([]byte(item.str), iv) {
			t.Errorf("Maru test %+v failed.\n", item)
		}
	}
}

func TestSpeck(t *testing.T) {
	if Speck(key64, binary.LittleEndian.Uint64(plain64[:])) != binary.LittleEndian.Uint64(cipher64[:]) {
		t.Errorf("SPECK-64/128 Test FAILED.\n")
	}
}
