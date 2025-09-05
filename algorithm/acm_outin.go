package algorithm

import (
	"bufio"
	"os"
)

// ACM 输入输出
var (
	Reader *bufio.Reader
	Writer *bufio.Writer
)

func init() {
	Reader = bufio.NewReader(os.Stdin)
	Writer = bufio.NewWriter(os.Stdout)
}
