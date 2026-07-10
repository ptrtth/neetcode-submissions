func hasDuplicate(nums []int) bool {
    
    unique := map[int]bool{}

    for _, num := range nums {
        if unique[num] {
            return true

        }
        unique[num] = true
    }

    return false
}
