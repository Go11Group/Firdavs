package postgres

import (
	"strconv"
	"strings"
)

// ReplaceQueryParams funksiyasi nomlangan so'rov parametrlarini PostgreSQL parametrlariga o'zgartiradi.
// Kiritilgan nomdagi so'rov string, o'zgartirilgan so'rov string va parametrlar ro'yxati qaytariladi.
func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int           // Parametrlar indeksi
		args []interface{} // Parametrlar ro'yxati
	)

	i = 1
	for k, v := range params {
		// Parametr nomini va $1, $2, ... indekslarni aniqlash uchun
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v) // Parametr qiymatini ro'yxatga qo'shish
			i++
		}
	}

	return namedQuery, args
}
