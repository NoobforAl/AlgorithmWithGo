package main

import (
	"log"
)

/*
*
*
* fullAdder
*
* args :
*   c1, c2 byte
*   flag bool
*
*
* return value:
*   byte, bool
*
*
 */
func fullAdder(c1, c2 byte, flag bool) (byte, bool) {
	switch {
	case c1 == '1' && c2 == '1':
		if flag {
			return '1', flag
		} else {
			return '0', true
		}

	case c1 != c2:
		if flag {
			return '0', flag
		} else {
			return '1', false
		}

	default:
		if flag {
			return '1', false
		} else {
			return '0', false
		}
	}
}

/*
*
*
* add Binary
*
* args :
*   a, b string
*
*
* return value:
*   string
*
*
 */
func addBinary(a, b string) string {
	var (
		ans, c1, c2 byte
		res         string
		flag        bool
		i           int
	)

	for ; i < len(a) && i < len(b); i++ {
		c1, c2 = a[len(a)-(1+i)], b[len(b)-(1+i)]
		ans, flag = fullAdder(c1, c2, flag)
		res = string(ans) + res
	}

	for ; i < len(a); i++ {
		if flag {
			c1 = a[len(a)-(1+i)]
			ans, flag = fullAdder(c1, '0', flag)
			res = string(ans) + res
			continue
		}
		res = string(a[len(a)-(1+i)]) + res
	}

	for ; i < len(b); i++ {
		if flag {
			c1 = b[len(b)-(1+i)]
			ans, flag = fullAdder(c1, '0', flag)
			res = string(ans) + res
			continue
		}
		res = string(b[len(b)-(1+i)]) + res
	}

	if flag {
		return "1" + res
	}
	return res
}

func main() {
	// simple test
	if ans := addBinary("11", "1"); ans != "100" {
		log.Fatalf("Worn Answer! answer equal `100` not equal : %v", ans)
	}
	log.Println("Answer is True!")
}
