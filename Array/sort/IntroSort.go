package sort

import (
	"time"
)

//内省排序，基于快速排序的一种混合排序算法，不具有稳定性。
//复杂度为O(NlogN) & O(logN)。
//主要限制了QuickSort的最坏情况，适合递归实现(没有爆栈风险)。
func IntroSort(list []int) {
	var life = uint(12)
	for sz := len(list); sz != 0; sz /= 2 {
		life++
	}
	magic = uint(time.Now().Unix())
	doIntroSort(list, life)
}
func doIntroSort(list []int, life uint) {
	if len(list) < sz_limit {
		InsertSort(list)
	} else if life == 0 {
		HeapSort(list)
	} else {
		var knot = partition(list)
		doIntroSort(list[:knot], life-1)
		doIntroSort(list[knot+1:], life-1)
	}
}
