// Package goNewton package implements an API wrapper for newton API (https://github.com/aunyks/newton-api)
package goNewton

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The Newton struct is an empty interface, it has all functionalities of the Newton API
type Newton struct {
}

// Response struct to destructure the JSON response into a Go type
type response struct {
	Operation  string `json:"operation"`
	Expression string `json:"expression"`
	Result     string `json:"result"`
}

// The url endpoint for the API
const (
	url = "https://newton.now.sh/%s/%s" //Format of URL "https://newton.now.sh/:operation/:expression"
)

// Simplify returns the simplified expression as a string
func (n *Newton) Simplify(exp string) (string, error) {
	const endPoint = "simplify"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Factor returns a string with factors of the expression.
// Returns the expression if factorisation is not possible.
func (n *Newton) Factor(exp string) (string, error) {
	const endPoint = "factor"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Derive return a string with the derivation of the expression
// Returns the expression if derivation is not possible.
func (n *Newton) Derive(exp string) (string, error) {
	const endPoint = "derive"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Integrate return a string with the integration on the expression
// Returns the expression if integration is not possible.
func (n *Newton) Integrate(exp string) (string, error) {
	const endPoint = "integrate"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Zeros return a string with the number of zeros in the expression
func (n *Newton) Zeros(exp string) (string, error) {
	const endPoint = "zeros"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Tangent returns a string with the tangential equation to the expression
func (n *Newton) Tangent(expr string, num int) (string, error) {
	const endPoint = "tangent"
	exp := fmt.Sprintf("%d|%s", num, expr)
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// TangentStr returns a string with the tangential equation to the expression
// Uses the Newton API string format "number|expression"
// The '|' is used to separate the number and expression
func (n *Newton) TangentStr(exp string) (string, error) {
	const endPoint = "tangent"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Area returns a string with the area under the curve of the given expression
// The expression of the curve and the starting and ending point are separated
func (n *Newton) Area(expr string, start int, end int) (string, error) {
	const endPoint = "area"
	exp := fmt.Sprintf("%d:%d|%s", start, end, expr)
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// AreaStr returns a string with the area under the curve of the given expression
// Uses the Newton API string format "start:end|expression"
// The starting and ending are given as numbers separated by a ':' and then the expression after a '|'
func (n *Newton) AreaStr(exp string) (string, error) {
	const endPoint = "area"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Cos returns a strign of the cosine of the given expresion
func (n *Newton) Cos(exp string) (string, error) {
	const endPoint = "cos"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Sin returns a string of the sine of the given expression
func (n *Newton) Sin(exp string) (string, error) {
	const endPoint = "sin"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Tan returns a string of the tan(gent) of the given expression
func (n *Newton) Tan(exp string) (string, error) {
	const endPoint = "tan"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// CosInverse returns a string of the cos inverse of the given expression
func (n *Newton) CosInverse(exp string) (string, error) {
	const endPoint = "arccos"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// SinInverse return a string of the sine inverse of the given expression
func (n *Newton) SinInverse(exp string) (string, error) {
	const endPoint = "arcsin"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// TanInverse returns a string of the tan inverse of te given expression
func (n *Newton) TanInverse(exp string) (string, error) {
	const endPoint = "arctan"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Absolute return a string of the absolute value of the given expression
func (n *Newton) Absolute(exp string) (string, error) {
	const endPoint = "abs"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// Log return a string with the log of a number with base, both as parameters
func (n *Newton) Log(num int, base int) (string, error) {
	const endPoint = "log"
	exp := fmt.Sprintf("%d|%d", base, num)
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// LogStr return a string of the log of a number
// Using the Newton API string format , "base|number"
func (n *Newton) LogStr(exp string) (string, error) {
	const endPoint = "log"
	call := fmt.Sprintf(url, endPoint, exp)
	res, err := http.Get(call)
	if err != nil {
		return "", err
	}
	var data response
	json.NewDecoder(res.Body).Decode(&data)
	return data.Result, nil
}

// New return a new Newton object
func New() *Newton {
	return new(Newton)
}
