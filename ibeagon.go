package ibeagon

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type IBeaconPacket struct {
	Interf string
	Uuid   string
	Major  int
	Minor  int
	Power  int
	Rssi   int
}

func NewIBPacket(line string) *IBeaconPacket {
	vals := strings.Split(line, " ")
	return &IBeaconPacket{
		Uuid:  parseUuid(vals),
		Major: parseMajor(vals),
		Minor: parseMinor(vals),
		Power: parsePower(vals),
		Rssi:  parseRssi(vals)}
}

func IsValid(str string) bool {
	r, err := regexp.Compile(`^04\ 3E\ 2A\ 02\ 01\ .{26}\ 02\ 01\ .{14}\ 02\ 15`)
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	return r.MatchString(str)
}

func (ibp *IBeaconPacket) MapKey() string {
	return fmt.Sprintf("IBE_%05d%05d", ibp.Major, ibp.Minor)
}

func parseUuid(vals []string) string {
	return strings.Join(
		[]string{
			strings.Join(vals[23:27], ""),
			strings.Join(vals[27:29], ""),
			strings.Join(vals[29:31], ""),
			strings.Join(vals[31:33], ""),
			strings.Join(vals[33:39], "")},
		"-")
}

func parseMajor(vals []string) int {
	major, _ := strconv.ParseInt(strings.Join(vals[39:41], ""), 16, 0)
	return int(major)
}

func parseMinor(vals []string) int {
	minor, _ := strconv.ParseInt(strings.Join(vals[41:43], ""), 16, 0)
	return int(minor)
}

func parsePower(vals []string) int {
	power, _ := strconv.ParseInt(vals[43], 16, 0)
	return int(power) - 256
}

func parseRssi(vals []string) int {
	rssi, _ := strconv.ParseInt(vals[44], 16, 0)
	return int(rssi) - 256
}

func (ibp *IBeaconPacket) ToString() string {
	return fmt.Sprintf("INT %v UUID %v MAJOR %v MINOR %v RSSI %d", ibp.Interf, ibp.Uuid, ibp.Major, ibp.Minor, ibp.Rssi)
}
