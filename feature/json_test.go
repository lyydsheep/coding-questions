package feature

import (
	"encoding/json"
	"io"
	"log"
	"runtime"
	"strings"
	"testing"
)

// 比较 json/v1 和 json/v2 的性能差异

const (
	N = 1_000_000
)

// TestV1Encoding Total bytes allocated during Encode: 8399320 bytes (8.01 MiB)
func TestV1Encoding(t *testing.T) {
	//	初始化待编码的切片
	in := make([]struct{}, N)
	// 仅执行编码操作，不保存实际内容
	out := io.Discard

	// 清除 sync.Pool对象，避免影响结果
	for range 5 {
		runtime.GC()
	}

	// 记录编码前的内存
	var beforeStats runtime.MemStats
	runtime.ReadMemStats(&beforeStats)

	// 编码
	encoder := json.NewEncoder(out)
	if err := encoder.Encode(in); err != nil {
		log.Fatalf("v1 Encode failed: %v", err)
	}

	// 记录编码后的内存
	var afterStats runtime.MemStats
	runtime.ReadMemStats(&afterStats)

	allocBytes := afterStats.TotalAlloc - beforeStats.TotalAlloc
	log.Printf("Total bytes allocated during Encode: %d bytes (%.2f MiB)", allocBytes, float64(allocBytes)/1024/1024)
}

// TestV1Decoding Total bytes allocated during Decode: 8387344 bytes (8.00 MiB)
func TestV1Decoding(t *testing.T) {
	str := "[" + strings.TrimSuffix(strings.Repeat("{},", N), ",") + "]"
	in := strings.NewReader(str)

	// 预分配容量，避免 slice 扩容影响结果
	out := make([]struct{}, 0, N)

	// 清除 sync.Pool对象，避免影响结果
	for range 5 {
		runtime.GC()
	}

	// 记录解码前的内存
	var beforeStats runtime.MemStats
	runtime.ReadMemStats(&beforeStats)

	// 解码
	decoder := json.NewDecoder(in)
	if err := decoder.Decode(&out); err != nil {
		log.Fatalf("v1 Decode failed: %v", err)
	}

	// 记录编码后的内存
	var afterStats runtime.MemStats
	runtime.ReadMemStats(&afterStats)

	allocBytes := afterStats.TotalAlloc - beforeStats.TotalAlloc
	log.Printf("Total bytes allocated during Dncode: %d bytes (%.2f MiB)", allocBytes, float64(allocBytes)/1024/1024)
}
