package questions

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 从第一个非叶子节点开始
	n := len(arr)
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 将最大值放到最后
	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
