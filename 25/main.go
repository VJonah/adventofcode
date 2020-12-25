package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	devices := parseInput("data.txt")
	fmt.Println(devices[0].findLoopSize())
	fmt.Println(devices[1].findLoopSize())
	fmt.Println(getEncryptionKey(devices[0], devices[1]))
}

func getEncryptionKey(d1, d2 *Device) int {
	encryptionKey1 := applyLoop(d1, d2.LoopSize)
	encryptionKey2 := applyLoop(d2, d1.LoopSize)
	if encryptionKey1 == encryptionKey2 {
		return encryptionKey2
	}
	return -1
}

type Device struct {
	SubjectN int
	LoopSize int
	Key      int
}

func applyLoop(d *Device, loopSize int) int {
	value := 1
	subjectN := d.Key
	for i := 0; i < loopSize; i++ {
		value *= subjectN
		value = value % 20201227
	}
	return value
}

func (d *Device) findLoopSize() int {
	value := 1
	for i := 0; i > -1; i++ {
		if value == d.Key {
			d.LoopSize = i
			return i
		}
		value *= d.SubjectN
		value = value % 20201227
	}
	return -1
}

func parseInput(href string) []*Device {
	data, _ := ioutil.ReadFile(href)
	lines := strings.Split(string(data), "\n")
	fmt.Println(lines)
	devices := []*Device{}
	device1Key, _ := strconv.Atoi(lines[0])
	device2Key, _ := strconv.Atoi(lines[1])
	devices = append(devices, &Device{Key: device1Key, SubjectN: 7})
	devices = append(devices, &Device{Key: device2Key, SubjectN: 7})
	return devices
}
