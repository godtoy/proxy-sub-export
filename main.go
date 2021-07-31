package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
)

func getInterfaces() ([]net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return interfaces, nil
}
func getIpAddress(i net.Interface) ([]string, error) {
	byName, err := net.InterfaceByName(i.Name)
	if err != nil {
		return nil, err
	}
	addresses, err := byName.Addrs()
	addrs := make([]string, 0)
	for _, v := range addresses {
		addrs = append(addrs, v.String())
	}
	return addrs, nil
}

var ipv4Reg, _ = regexp.Compile(`\d+\.\d+\.\d+\.\d+.*`)

func main() {
	interfaces, err := getInterfaces()
	if err != nil {
		fmt.Println("Getting interface error:", err)
		return
	}
	for i, v := range interfaces {
		fmt.Println(i, " : ", v.Name)
	}
	number := -1
	for {
		fmt.Print("Which NIC do you want to search：")
		var no string
		_, err = fmt.Scanln(&no)
		if err != nil {
			fmt.Println(err)
			continue
		}
		vv, err := strconv.Atoi(no)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if vv >= 0 && vv < len(interfaces) {
			number = vv
			break
		}
	}
	if number < 0 {
		return
	}

	ips, err := getIpAddress(interfaces[number])
	if err != nil {
		fmt.Println("Getting ip address error: ", err)
		return
	}

	f, err := os.OpenFile("ip.txt", os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println("Sync to file error: ", err)
		return
	}
	defer f.Close()
	for _, i := range ips {
		fmt.Println(i)
		if !ipv4Reg.MatchString(i) {
			continue
		}
		_, _ = f.WriteString(i + "\r\n")
	}
	fmt.Println("Success, save data to ip.txt")
}
