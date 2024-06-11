package api

import (
	"strconv"
	"strings"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		namedQuery = "where nation = $3 and gender = $1 and age = $2 "
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			//["male", 34, "uzbek"]
			//   0     1     2
			i++
		}
	}

	return namedQuery, args
}
