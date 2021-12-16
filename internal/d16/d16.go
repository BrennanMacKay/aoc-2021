package d16

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"math"
	"strconv"
	"strings"
)

var hexMap = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type packet struct {
	version int
	id      int

	value   int
	packets []packet
}

func Day16(args []string) int {
	switch args[0] {
	case "p1":
		return part1(args[1:])
	case "p2":
		return part2(args[1:])
	default:
		fmt.Println("Unknown part", args[0])
		return 1
	}
}

func part1(args []string) int {
	input := tools.Read(args[0])
	indexToRun, _ := strconv.Atoi(args[1])
	fmt.Println(input[indexToRun])
	bits := hexToBinary(input[indexToRun])
	fmt.Println(bits)

	packet, _ := parsePacket(bits)
	fmt.Println(packet)

	fmt.Println(sumVersions(packet))

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	indexToRun, _ := strconv.Atoi(args[1])
	fmt.Println(input[indexToRun])
	bits := hexToBinary(input[indexToRun])
	fmt.Println(bits)

	packet, _ := parsePacket(bits)
	fmt.Println(packet)

	fmt.Println(computePacket(packet))

	return 0

}

func computePacket(packet packet) int {
	switch packet.id {
	case 0: //sum
		sum := 0
		for _, p := range packet.packets {
			sum += computePacket(p)
		}
		return sum
	case 1: //product
		prod := 1
		for _, p := range packet.packets {
			prod *= computePacket(p)
		}
		return prod
	case 2: //min
		min := math.MaxInt
		for _, p := range packet.packets {
			res := computePacket(p)
			if res < min {
				min = res
			}
		}
		return min
	case 3: //max
		max := 0
		for _, p := range packet.packets {
			res := computePacket(p)
			if res > max {
				max = res
			}
		}
		return max
	case 4: //literal
		return packet.value
	case 5: //greater than
		if computePacket(packet.packets[0]) > computePacket(packet.packets[1]) {
			return 1
		} else {
			return 0
		}
	case 6: //less than
		if computePacket(packet.packets[0]) < computePacket(packet.packets[1]) {
			return 1
		} else {
			return 0
		}
	case 7: //equal to
		a := computePacket(packet.packets[0])
		b := computePacket(packet.packets[1])
		fmt.Println("EQUAL?", a, b)

		if a == b {
			return 1
		} else {
			return 0
		}
	default:
		panic("Unknown type id!")
	}
}

func sumVersions(packet packet) int {
	sum := packet.version
	for _, p := range packet.packets {
		sum += sumVersions(p)
	}

	return sum
}

func binaryToInt(input string) int {
	i, _ := strconv.ParseInt(input, 2, 64)
	return int(i)
}

func parsePacket(input string) (packet, string) {
	fmt.Println(input)
	packet := packet{}
	packet.version, packet.id = binaryToInt(input[0:3]), binaryToInt(input[3:6])
	input = input[6:]

	switch packet.id {
	case 4:
		packet.value, input = parseLiteral(input)
	default:
		packet.packets, input = parseOperator(input)
	}

	return packet, input
}

func parseOperator(input string) ([]packet, string) {
	lengthType := input[0]
	input = input[1:]

	var packets []packet
	switch lengthType {
	case '0':
		packets, input = parseBitLength(input)
	case '1':
		packets, input = parsePacketCount(input)
	}

	return packets, input
}

func parseBitLength(input string) ([]packet, string) {
	lengthInBits := binaryToInt(input[0:15])
	input = input[15:]

	consumed := 0
	packets := make([]packet, 0)
	for consumed < lengthInBits {
		packet, newInput := parsePacket(input)
		packets = append(packets, packet)
		consumed += len(input) - len(newInput)
		input = newInput
	}

	return packets, input
}

func parsePacketCount(input string) ([]packet, string) {
	packetCount := binaryToInt(input[0:11])
	input = input[11:]

	consumed := 0
	packets := make([]packet, 0)
	for consumed < packetCount {
		packet, newInput := parsePacket(input)
		packets = append(packets, packet)
		consumed++
		input = newInput
	}

	return packets, input
}

func parseLiteral(input string) (int, string) {
	group := input[0:5]
	input = input[5:]
	builder := strings.Builder{}
	for {
		builder.WriteString(group[1:5])
		if group[0] == '0' {
			break
		} else {
			group = input[0:5]
			input = input[5:]
		}
	}

	bits := builder.String()
	value := binaryToInt(bits)
	return value, input
}

func hexToBinary(input string) string {
	builder := strings.Builder{}
	for _, h := range []rune(input) {
		builder.WriteString(hexMap[h])
	}

	return builder.String()
}
