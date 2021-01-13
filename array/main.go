package main

/*
	计算container-with-most-water
	函数是 c = (k-i)*(a[k]+a[i]-|a[k]-a[i]|)/2
*/

func main() {

}

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	if len(height) == 1 {
		return height[0] * 1
	}

	var area, t1, t2, tmp int
	t1 = 0
	t2 = len(height) - 1
	for t1 < t2 {
		if height[t2] > height[t1] {
			tmp = height[t1] * (t2 - t1)
			t1++
		} else {
			tmp = height[t2] * (t2 - t1)
			t2--
		}

		if tmp > area {
			area = tmp
		}
	}

	return area
}
