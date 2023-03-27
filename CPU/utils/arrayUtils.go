package utils

func ArrUInt32ExitNumber(arr *[]uint32, target int) bool {
	for i, _ := range *arr {
		if target == i {
			return true
		}
	}
	return false
}
