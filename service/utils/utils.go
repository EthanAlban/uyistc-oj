package utils

import (
	"fmt"
	"net"
	"strings"
)

// 这个包里是一些个公共的工具

// GetExternalIp 获取本机对外ip
func GetExternalIp() {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}

// PrintSysLogo 打印系统图形
func PrintSysLogo() {
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, "  ___       ___     ___         ___     ___     _____________     _______________\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |  \\  \\     |   |   |   |   |   _______   |   |_______________|\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   \\  \\    |   |    ———    |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   |\\  \\   |   |    ___    |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   | \\  \\  |   |   |   |   |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   |  \\  \\ |   |   |   |   |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |_____|   |   |   |   \\  \\|   |   |   |   |  |_______|  |     | | |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |_____________|   |___|    \\______|   |___|   |_____________|        |  |      V 1.0.0\n", 0x1B)
}
