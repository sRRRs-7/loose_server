package utils

func StarContains(arr []int64, user_id int64) []int64 {
	for i := range arr {
		if arr[i] == user_id {
			arr[i] = arr[len(arr)-1]
			return arr[:len(arr)-1]
		}
	}
	arr = append(arr, user_id)
	return arr
}
