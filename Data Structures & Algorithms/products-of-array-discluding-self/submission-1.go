func productExceptSelf(nums []int) []int {
    product := 1
    result  := []int{}
    zeroes  := 0

    for _,num := range nums {
        if num == 0 {
            zeroes += 1
            continue
        }
        product *= num
    }

    if zeroes > 1 {
        for range nums {
            result = append(result, 0)
        }
    } else {
        for _,num := range nums {
            if zeroes == 1 {
                if num != 0 {
                    result = append(result, 0)
                } else {
                    result = append(result, product)
                }
            } else {
                result = append(result, product / num)
            }

        }
    }

    return result
}