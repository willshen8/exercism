package flatten

func Flatten(input interface{}) []interface{} {
	var results = []interface{}{}
	switch e := input.(type) {
	case []interface{}:
		for _, v := range e {
			results = append(results, Flatten(v)...)
		}
	case interface{}:
		results = append(results, e)
	}
	return results
}
