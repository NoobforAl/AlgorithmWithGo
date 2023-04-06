package main

import (
	"log"
)

/*
*
*
* search Insert
*
* args :
*   s string
*
*
* return value:
*   string
*
*
 */
func longestPalindrome(s string) string {
	// this answer take more memory!
	if len(s) < 2 {
		return s
	}

	table := make([][]bool, len(s))
	for i := range table {
		table[i] = make([]bool, len(s))
		table[i][i] = true
	}

	var start, max, i, k int = 0, 1, 0, 3

	for ; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			start = i
			max = 2
		}
	}

	for ; k <= len(s); k++ {
		i = 0
		for ; i < (len(s) - k + 1); i++ {
			j := i + k - 1
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true

				if k > max {
					start = i
					max = k
				}
			}
		}
	}

	return s[start : start+max]
}

func main() {
	// simple test
	if ans := longestPalindrome("aaaacrc"); ans != "aaaa" {
		log.Fatalf("Worn Answer! answer equal `aaaa` not equal : %s", ans)
	}
	log.Println("Answer is True!")
}

/*
//* bad answer
//* O( n^3 )

```go
func isPalindrome(s *string, i, j int) bool {
    for k := 0; k < ( (j - i) / 2 ) + 1; k++ {
        if (*s)[i+k] != (*s)[j-k] { return false }
    }
    return true
}


func longestPalindrome(s string) string {
    if len(s) < 2 { return s }
    var max, start int = 1, 0
    var flag bool

    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            flag = isPalindrome(&s, i, j)
            if flag && (j-i+1) > max {
                start = i
                max = j-i+1
            }
        }
    }
    return s[start:start+max]
}
```
*/

/*
//* better answer
//* O ( n^2 )
```go
func longestPalindrome(s string) string {
    answ := ""
    for i := 0; i < len(s); i++ {
        r := i
        for r < len(s) && s[r] == s[i] {
            r++
        }
        l := i - 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            l--
            r++
        }
        if r - l - 1 > len(answ) {
            answ = s[l+1:r]
        }
    }
    return answ
}
```
*/
