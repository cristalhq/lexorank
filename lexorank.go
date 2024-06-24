package lexorank

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	minChar = '0'
	maxChar = 'z'
)

// Rank returns a new Lexorank between prev and next.
// Uses 0-9A-Za-z alphabet.
func Rank(prev, next string) (string, error) {
	if prev == "" {
		prev = string(minChar)
	}
	if next == "" {
		next = string(maxChar)
	}
	if !isValid(prev) || !isValid(next) {
		return "", errors.New("lexorank: incorrect prev or next")
	}

	rank := make([]byte, 0, max(len(prev), len(next)))

	// will take upto |prev| + |next| + C iterations
	for i := 0; ; i++ {
		prev := getChar(prev, i, minChar)
		next := getChar(next, i, maxChar)

		if prev == next {
			rank = append(rank, prev)
			continue
		}

		mid := avg(prev, next)
		if mid == prev || mid == next {
			rank = append(rank, prev)
			continue
		}

		rank = append(rank, mid)
		break
	}

	if r := string(rank); r < next {
		return r, nil
	}
	return prev, nil
}

// RankN returns n Lexoranks between prev and next (via simple iteration).
// Uses Rank under the hood.
func RankN(prev, next string, n int) ([]string, error) {
	idx, err := Rank(prev, next)
	if err != nil {
		return nil, err
	}

	suffixRankLen := len(strconv.Itoa(n))
	suffixRankFormat := fmt.Sprintf("%%0%dd", suffixRankLen)

	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, idx+fmt.Sprintf(suffixRankFormat, i))
	}
	return res, nil
}
