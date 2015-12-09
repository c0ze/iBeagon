package ibeagon

import (
	"strings"
	"testing"
)

// UUID: B9407F30-F5F8-466E-AFF9-25556B57FE6D MAJOR: 2731 MINOR: 34771 POWER: -81 RSSI: -88
// UUID: B9407F30-F5F8-466E-AFF9-25556B57FE6D MAJOR: 25110 MINOR: 65216 POWER: -81 RSSI: -81
const Packet1 = `04 3E 2A 02 01 00 01 16 62 C0 FE 30 C6 1E 02 01 06 1A FF 4C 00 02 15 B9 40 7F 30 F5 F8 46 6E AF F9 25 55 6B 57 FE 6D 62 16 FE C0 AF AF`

// UUID: B9407F30-F5F8-466E-AFF9-25556B57FE6D MAJOR: 2731 MINOR: 34771 POWER: -81 RSSI: -86
const Packet2 = `04 3E 2A 02 01 00 01 AB 0A D3 87 1C DF 1E 02 01 06 1A FF 4C 00 02 15 B9 40 7F 30 F5 F8 46 6E AF F9 25 55 6B 57 FE 6D 0A AB 87 D3 AF AA`

// invalid packet
const Packet3 = `04 3E 2A 02 01 00 01 AB 0A D3 87 1C DF 1E 02 01 06 1A FF 4C 00 12 15 B9 40 7F 30 F5 F8 46 6E AF F9 25 55 6B 57 FE 6D 0A AB 87 D3 AF AA`

func TestParseUuid(t *testing.T) {
	uuid := parseUuid(strings.Split(Packet1, " "))
	expected_uuid := "B9407F30-F5F8-466E-AFF9-25556B57FE6D"
	if uuid != expected_uuid {
		t.Errorf("Parsing uuid {%v} failed for packet1: %v", expected_uuid, uuid)
	}

	uuid = parseUuid(strings.Split(Packet2, " "))
	if uuid != expected_uuid {
		t.Errorf("Parsing uuid {%v} failed for packet2: %v", expected_uuid, uuid)
	}
}

func TestParseMajor(t *testing.T) {
	major := parseMajor(strings.Split(Packet1, " "))
	expected_major := 25110
	if major != expected_major {
		t.Errorf("Parsing major {%v} failed for packet1: %v", expected_major, major)
	}

	major = parseMajor(strings.Split(Packet2, " "))
	expected_major = 2731
	if major != expected_major {
		t.Errorf("Parsing major {%v} failed for packet2: %v", expected_major, major)
	}
}

func TestParseMinor(t *testing.T) {
	minor := parseMinor(strings.Split(Packet1, " "))
	expected_minor := 65216
	if minor != expected_minor {
		t.Errorf("Parsing major {%v} failed for packet1: %v", expected_minor, minor)
	}

	minor = parseMinor(strings.Split(Packet2, " "))
	expected_minor = 34771
	if minor != expected_minor {
		t.Errorf("Parsing major {%v} failed for packet2: %v", expected_minor, minor)
	}
}

func TestParsePower(t *testing.T) {
	power := parsePower(strings.Split(Packet1, " "))
	expected_power := -81
	if power != expected_power {
		t.Errorf("Parsing power {%v} failed for packet1: %v", expected_power, power)
	}

	power = parsePower(strings.Split(Packet2, " "))
	if power != expected_power {
		t.Errorf("Parsing power {%v} failed for packet2: %v", expected_power, power)
	}
}

func TestParseRssi(t *testing.T) {
	rssi := parseRssi(strings.Split(Packet1, " "))
	expected_rssi := -81
	if rssi != expected_rssi {
		t.Errorf("Parsing rssi {%v} failed for packet1: %v", expected_rssi, rssi)
	}

	rssi = parseRssi(strings.Split(Packet2, " "))
	expected_rssi = -86
	if rssi != expected_rssi {
		t.Errorf("Parsing rssi {%v} failed for packet2: %v", expected_rssi, rssi)
	}
}

func TestIsValid(t *testing.T) {
	if !IsValid(Packet1) {
		t.Errorf("Validation failed for packet1")
	}

	if !IsValid(Packet2) {
		t.Errorf("Validation failed for packet2")
	}

	if IsValid(Packet3) {
		t.Errorf("Validation passed for packet3")
	}
}

func TestToString(t *testing.T) {
	ibp := NewIBPacket(Packet1)
	packet1String := "INT  UUID B9407F30-F5F8-466E-AFF9-25556B57FE6D MAJOR 25110 MINOR 65216 RSSI -81"
	if packet1String != ibp.ToString() {
		t.Errorf("ToString failed for packet1 \nexpected: %v\ngot: %v", packet1String, ibp.ToString())
	}

	ibp = NewIBPacket(Packet2)
	packet2String := "INT  UUID B9407F30-F5F8-466E-AFF9-25556B57FE6D MAJOR 2731 MINOR 34771 RSSI -86"
	if packet2String != ibp.ToString() {
		t.Errorf("ToString failed for packet2 \nexpected: %v\ngot: %v", packet2String, ibp.ToString())
	}
}

func TestMapKey(t *testing.T) {
	ibp := NewIBPacket(Packet1)
	packet1MapKey := "IBE_2511065216"
	if packet1MapKey != ibp.MapKey() {
		t.Errorf("MapKey failed for packet1 \nexpected: %v\ngot: %v", packet1MapKey, ibp.MapKey())
	}

	ibp = NewIBPacket(Packet2)
	packet2MapKey := "IBE_0273134771"
	if packet2MapKey != ibp.MapKey() {
		t.Errorf("MapKey failed for packet2 \nexpected: %v\ngot: %v", packet2MapKey, ibp.MapKey())
	}
}
