func maxArea(height []int) int {
    i, j := 0, len(height)-1
    max := 0

    for i < j {
        h := int(math.Min(float64(height[i]), float64(height[j])))
        max = int(math.Max(float64(max), float64(h*(j-i))))
        if height[i] < height[j] {
            i++
        } else {
            j--
        }
    }

    return max
}
