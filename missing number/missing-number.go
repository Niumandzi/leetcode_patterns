func missingNumber(nums []int) int {
    n := len(nums)
    sum := n*(n + 1)/2 

    for _, num := range nums {
        sum -= num
    }

    return sum
}

func missingNumberSecond(nums []int) int {
    ans := 0
    for i,n := range nums {
        ans ^= (i+1)
        ans ^= n
    }
    return ans
}
