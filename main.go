package main

import (
	"fmt"
	"golang.zx2c4.com/wireguard/tun"
	"log"
	"strconv"
)

func allocateTUN() (tun.Device, error) {
	for i := 0; i <= 254; i++ {
		name := "utun" + strconv.Itoa(i)
		tunDevice, err := tun.CreateTUN(name, 1420)
		if err == nil {
			log.Printf("Successfully allocated TUN device: %s", name)
			return tunDevice, nil
		}
	}
	return nil, fmt.Errorf("no available TUN devices")
}

func main() {
	log.Println("Starting TUN device allocation...")
	tunDevice, err := allocateTUN()
	if err != nil {
		log.Fatalf("Failed to create TUN: %v", err)
	}
	defer tunDevice.Close()

	tunDeviceName, err := tunDevice.Name()
	if err != nil {
		log.Fatalf("Failed to get TUN name: %v", err)
	}

	log.Println("TUN interface created with automatic name:", tunDeviceName)

	//packet := make([]byte, 2048)
	//for {
	//	n, err := tunDevice.Read(packet, 0)
	//	if err != nil {
	//		log.Printf("Error reading from TUN: %v", err)
	//		continue
	//	}
	//
	//	log.Printf("Received packet: % x", packet[:n])
	//
	//	// Echo the same packet back
	//	_, err = tunDevice.Write(packet[:n], 0)
	//	if err != nil {
	//		log.Printf("Error writing to TUN: %v", err)
	//	}
	//}
}
