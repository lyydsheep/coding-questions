package questions

import (
	"context"
	"fmt"
	"sync"
)

// SolveQ7 创建10个goroutine，id分别是 0,1,2,3,4,5,6,7,8,9，每个 goroutine 打印以 id 结尾的数字
// 例如 id 为 1 的goroutine 打印 1, 11, 21...
// 升序输出 0～100000
func SolveQ7(ctx context.Context) {
	const (
		N = 100000
		n = 10
	)
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		// 条件变量
		cond = sync.NewCond(&mu)
	)

	cur := 0
	wg.Add(n)
	for i := range n {
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			for cur <= N {
				// 判断是否能够输出
				// 循环检查，避免虚假唤醒
				for cur <= N && cur%10 != id {
					// 不能输出，让出 mu并阻塞挂起
					cond.Wait()
				}

				if cur > N {
					return
				}

				// 可以输出
				fmt.Println(id, cur)
				// 向下传递
				cur++
				cond.Broadcast()
			}
		}(i)
	}

	wg.Wait()
}
