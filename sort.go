package imp

func SortBubble(arr []int) {
	if len(arr) == 0 {
		return
	}

	f := false
	for i := 0; i < len(arr); i++ {
		f = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				f = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !f {
			break
		}
	}
}

func SortInsert(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
}

func SortSelect(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i<len(arr); i++ {
		f := i
		for j:=i+1; j<len(arr); j++ {
			if arr[j]<arr[f] {
				f = j
			}
		}
		arr[i], arr[f] = arr[f], arr[i]
	}
}