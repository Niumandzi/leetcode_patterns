func containsDuplicate(nums []int) bool {
    keys := make(map[int]struct{})

    for i := range nums {
        if _, in := keys[nums[i]]; in {
            return true
        }
        keys[nums[i]]=struct{}{}
    }

    return false
}
