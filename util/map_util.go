package util

// func description:  判断map是否为空或者nil
// input: map
// output: false: 有数据：
func IsMapEmpty(mapData map[string]interface{}) bool {
	if mapData == nil || len(mapData) == 0 {
		return true
	}

	return false
}
