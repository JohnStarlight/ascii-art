package internal

import (
	"fmt"
	"os"
	"strings"
)

func PrintAscii(lines []string, filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Could not open file:", err)
		os.Exit(1)
	}
	bannerLines := strings.Split(string(data), "\n")

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			if i > 0 {
				fmt.Println()
			}
			continue
		}
		for row := 1; row <= 8; row++ {
			for _, r := range lines[i] {
				fmt.Print(bannerLines[(int(r)-32)*9+row])
			}
			fmt.Println()
		}
	}
}
