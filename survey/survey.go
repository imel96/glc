package survey

func CountBigShoppers(N int, s []int) int {
	buyer := N

	for _, purchases := range s {
		buyer -= (N - purchases)
	}
	if buyer < 0 {
		return 0

	} else {
		return buyer
	}
}
