package lib

import (
	"log"
	"strconv"
)

type Number64 interface {
	~int64 | ~float64
}

type Number32 interface {
	~int | ~int32 | ~float32
}

type Number interface {
	Number32 | Number64
}

func NumbToString[N Number](numb N) string {
	i := int(numb)
	a := strconv.Itoa(i)
	return a
}

func StringToNumber[N Number](s string) N {
	i, parseIntErr := strconv.Atoi(s)
	if parseIntErr != nil {
		f, parseFloatErr := strconv.ParseFloat(s, 64)
		if parseFloatErr != nil {
			log.Println(parseFloatErr)
			return 0
		}
		return N(f)
	}
	return N(i)
}
