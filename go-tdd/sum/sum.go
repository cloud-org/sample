// author: ashing
// time: 2020/4/16 3:53 下午
// mail: axingfly@gmail.com
// Less is more.

package sum

// calc sum
func Sum(num []int) int {
	sum := 0
	//for i := 0; i < len(num); i++ {
	//	sum += num[i]
	//}
	for _, number := range num {
		sum += number
	}
	return sum
}

// []int
// []int []int
// []int []int []int
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // numbers[1:] 取出除去第一个(下标为0)的所有元素
			sums = append(sums, Sum(tail))
		}
	}

	return
}
