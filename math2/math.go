package math

func SegmentIndex(segs []int, distance int) (index, indexUsed, remains int) {
	var sum int

	index = -1
	for i := 0; i < len(segs); i++ {
		if distance <= sum+segs[i] && index == -1 {
			indexUsed = distance - sum
			index = i
		}

		sum += segs[i]
	}
	remains = sum - distance
	return
}

func Round(f float64) int64 {
	i := int64(f)
	if f-float64(i) >= 0.5 {
		i += 1
	}
	return i
}
