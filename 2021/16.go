package main

import (
	"fmt"
	"strconv"
	"strings"
)

type packet struct {
	version int64
	typeId  int64
	value   int64
}

var (
	input      = "2052ED9802D3B9F465E9AE6003E52B8DEE3AF97CA38100957401A88803D05A25C1E00043E1545883B397259385B47E40257CCEDC7401700043E3F42A8AE0008741E8831EC8020099459D40994E996C8F4801CDC3395039CB60E24B583193DD75D299E95ADB3D3004E5FB941A004AE4E69128D240130D80252E6B27991EC8AD90020F22DF2A8F32EA200AC748CAA0064F6EEEA000B948DFBED7FA4660084BCCEAC01000042E37C3E8BA0008446D8751E0C014A0036E69E226C9FFDE2020016A3B454200CBAC01399BEE299337DC52A7E2C2600BF802B274C8848FA02F331D563B3D300566107C0109B4198B5E888200E90021115E31C5120043A31C3E85E400874428D30AA0E3804D32D32EED236459DC6AC86600E4F3B4AAA4C2A10050336373ED536553855301A600B6802B2B994516469EE45467968C016D004E6E9EE7CE656B6D34491D8018E6805E3B01620C053080136CA0060801C6004A801880360300C226007B8018E0073801A801938004E2400E01801E800434FA790097F39E5FB004A5B3CF47F7ED5965B3CF47F7ED59D401694DEB57F7382D3F6A908005ED253B3449CE9E0399649EB19A005E5398E9142396BD1CA56DFB25C8C65A0930056613FC0141006626C5586E200DC26837080C0169D5DC00D5C40188730D616000215192094311007A5E87B26B12FCD5E5087A896402978002111960DC1E0004363942F8880008741A8E10EE4E778FA2F723A2F60089E4F1FE2E4C5B29B0318005982E600AD802F26672368CB1EC044C2E380552229399D93C9D6A813B98D04272D94440093E2CCCFF158B2CCFE8E24017CE002AD2940294A00CD5638726004066362F1B0C0109311F00424CFE4CF4C016C004AE70CA632A33D2513004F003339A86739F5BAD5350CE73EB75A24DD22280055F34A30EA59FE15CC62F9500"
	packets    = []*packet{}
	hexIntoBin = map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
)

func main() {
	bin := ""
	for i := 0; i < len(input); i++ {
		bin += hexIntoBin[input[i:i+1]]
	}

	var res *packet
	for bin != strings.Repeat("0", len(bin)) {
		res, bin = parse(bin)
	}

	sum := 0
	for _, p := range packets {
		sum += int(p.version)
	}
	fmt.Println("Part One: ", sum)
	fmt.Println("Part Two: ", res.value)

}

func parse(bin string) (*packet, string) {
	version := bin[0:3]
	ver, _ := strconv.ParseInt(version, 2, 64)
	bin = bin[3:]

	typeId := bin[0:3]
	typ, _ := strconv.ParseInt(typeId, 2, 64)
	bin = bin[3:]

	num := int64(0)
	inside := []*packet{}

	if typ == 4 {
		end := false
		literal := ""
		for !end {
			// proc += 5
			sym := bin[0:5]
			bin = bin[5:]
			literal += sym[1:]
			if sym[0:1] == "0" {
				end = true
			}
		}
		num, _ = strconv.ParseInt(literal, 2, 64)
		// bin = bin[(proc/4+1)*4-proc:]
	} else {
		lenType := bin[0:1]
		lenStr := ""
		if lenType == "0" {
			lenStr = bin[1:16]
			bin = bin[16:]
			// proc += 16
		} else {
			lenStr = bin[1:12]
			bin = bin[12:]
			// proc += 12
		}
		length, _ := strconv.ParseInt(lenStr, 2, 64)
		if lenType == "0" {
			subbin := bin[:length]
			bin = bin[length:]
			for len(subbin) != 0 {
				var res *packet
				res, subbin = parse(subbin)
				inside = append(inside, res)
			}
		} else {
			for i := 0; i < int(length); i++ {
				var res *packet
				res, bin = parse(bin)
				inside = append(inside, res)
			}
		}
		// fmt.Println(ver, typ, lenStr, length)
		switch typ {
		case 0:
			for _, p := range inside {
				num += p.value
			}
		case 1:
			num = 1
			for _, p := range inside {
				num *= p.value
			}
		case 2:
			num = 1000000000
			for _, p := range inside {
				if p.value < num {
					num = p.value
				}
			}
		case 3:
			for _, p := range inside {
				if p.value > num {
					num = p.value
				}
			}
		case 5:
			if inside[0].value > inside[1].value {
				num = 1
			}
		case 6:
			if inside[0].value < inside[1].value {
				num = 1
			}
		case 7:
			if inside[0].value == inside[1].value {
				num = 1
			}
		}
	}

	parsed := &packet{version: ver, typeId: typ, value: num}
	packets = append(packets, parsed)
	return parsed, bin
}
