func twoSum(numbers []int, target int) []int {
    mp := make(map[int]int)
    for i := 0; i < len(numbers); i++ {
        tmp := target - numbers[i]
        if val, exists := mp[tmp]; exists {
            return []int{val, i + 1}
        }
        mp[numbers[i]] = i + 1
    }
    return nil
}
