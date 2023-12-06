package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Вывод ошибки,", err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	fmt.Println(input)
}
