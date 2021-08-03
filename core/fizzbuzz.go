// Package core provides the domain layer.
package core

import (
	"fmt"
	"strconv"
)

// FizzBuzzLimitMax is the maximum allowed limit value here 100_000 is an
// arbitrary value, this nust be depend of product owner.
const FizzBuzzLimitMax = 100_000

// FizzBuzzParams is a struct with all the parameters to FizzBuzz function
type FizzBuzzParams struct {
	Limit int    `json:"limit"` // limit of the loop
	Int1  int    `json:"int1"`  // all multiples of [Int1] are replaced by [Str1] (classic fizzbuzz: 3)
	Str1  string `json:"str1"`  // replacement string for multiples of Int1 (classic fizzbuzz: "fizz")
	Int2  int    `json:"int2"`  // all multiples of [Int2] are replaced by [Str2] (classic fizzbuzz: 5)
	Str2  string `json:"str2"`  // replacement string for multiples of [Int2] (classic fizzbuzz: "buzz")
}

// HasError checks if the current FizzBuzzParams structure has an error and
// returns it otherwise this method returns nil.
//
// [Limit], [Int1], [Int2] must be a positive integer.
// [Limit] can't be greater than MAX_LIMIT constant
func (params FizzBuzzParams) HasError() error {
	if params.Limit < 1 {
		return fmt.Errorf("Bad parameter: 'limit' must be a positive number - got (%d)", params.Limit)
	}
	if params.Limit > FizzBuzzLimitMax {
		return fmt.Errorf("Bad parameter: maximum 'limit' reached - got (%d) max (%d)", params.Limit, FizzBuzzLimitMax)
	}
	if params.Int1 < 1 {
		return fmt.Errorf("Bad parameter: multiple must be a positive number - got (int1: %d)", params.Int1)
	}
	if params.Int2 < 1 {
		return fmt.Errorf("Bad parameter: multiple must be a positive number - got (int2: %d)", params.Int2)
	}

	return nil
}

// DefaultFizzBuzzParams are the parameters of the original fizzbuzz.
var DefaultFizzBuzzParams = FizzBuzzParams{
	Limit: 100,
	Int1:  3,
	Str1:  "fizz",
	Int2:  5,
	Str2:  "buzz",
}

// FizzBuzz function take a [p FizzBuzzParams] and return a list of strings
// with numbers from 1 to [p.Limit], where:
//  - all multiples of [p.Int1] are replaced by [p.Str1],
//  - all multiples of [p.Int2] are replaced by [p.Str2],
//  - all multiples of [p.Int1] and [p.Int2] are replaced by [p.Str1][p.Str2]
//
// if [p FizzBuzzParams] is invalid an error is returned.
//
// Example of use :
//  fizzbuzz, err := core.FizzBuzz(
//      core.FizzBuzzParams{
//          Limit: 10,
//          Int1:  3,
//          Str1:  "fizzzz",
//          Int2:  5,
//          Str2:  "buzzzz",
//      },
//  )
func FizzBuzz(p FizzBuzzParams) ([]string, error) {
	if err := p.HasError(); err != nil {
		return nil, err
	}

	res := make([]string, p.Limit)

	for i := 1; i <= p.Limit; i++ {
		var s string
		if i%p.Int1 == 0 {
			s += p.Str1
		}
		if i%p.Int2 == 0 {
			s += p.Str2
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		res[i-1] = s
	}
	return res, nil
}
