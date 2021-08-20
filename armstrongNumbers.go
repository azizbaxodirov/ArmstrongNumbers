package armstrongnumbers

func ArmstrongNumbers(between, to int) (result []int) {

	for i := between; i <= to; i++ {
		temp := i
		var sum int
		var nums []int

		for {
			nums = append(nums, temp%10)
			temp = temp / 10

			if temp == 0 {
				break
			}
		}

		for j := 0; j < len(nums); j++ {
			digit := nums[j]
			for g := 0; g < len(nums)-1; g++ {
				digit = digit * nums[j]
			}

			sum += digit
		}

		if sum == i {
			result = append(result, i)
		}
	}
	return result
}
