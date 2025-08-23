package algorithm

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ACM 输入输出
var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func readInts() []int {
	str, _ := reader.ReadString('\n')
	if str[len(str)-1] == '\n' {
		str = str[:len(str)-1]
	}
	strs := strings.Split(str, " ")
	nums := make([]int, len(strs))
	for i := range strs {
		nums[i], _ = strconv.Atoi(strs[i])
	}
	return nums
}

func readInt() int {
	str, _ := reader.ReadString('\n')
	if str[len(str)-1] == '\n' {
		str = str[:len(str)-1]
	}
	num, _ := strconv.Atoi(str)
	return num
}
