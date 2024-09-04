package utils

import "reflect"

func Contains(slice, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("contains: not slice")
	}

	for i := 0; i < s.Len(); i++ {
		//ShowInfoLogs(fmt.Sprintf("type %v - %v, item type %v - %v, %v", reflect.TypeOf(s.Index(i).Interface()), s.Index(i).Interface(), reflect.TypeOf(item), item, reflect.DeepEqual(s.Index(i).Interface(), item)))
		if reflect.DeepEqual(s.Index(i).Interface(), item) {
			return true
		}
	}

	return false
}
