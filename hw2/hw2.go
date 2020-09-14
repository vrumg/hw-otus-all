package hw2

import (
	"errors"
	"strconv"
)

//Unpack Compose unpacked string from given packed variant
func Unpack(packedStr string) (string, error) {

	var retErr error
	var retString string

	//in case string does not contain '\' symbol
	var prevSym rune
	for i, sym := range packedStr {
		str := string(sym)
		intSym, err := strconv.Atoi(str)

		//is digit and first item
		if err == nil && i == 0 {
			retErr = errors.New("first rune is digit")
			return retString, retErr
		}

		//is digit
		if err == nil {
			//duplicate previous symbol by digit value times
			for ; intSym > 1; intSym-- {
				retString = retString + string(prevSym)
			}
			continue
		}

		retString = retString + string(sym)
		prevSym = sym
	}

	//log.Println(len(retString), retErr)

	return retString, retErr
}
