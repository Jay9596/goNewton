# goNewton

goNewton is a simple Go API wrapper for the [newton micro-service.](https://newton.now.sh/)  
[GitHub](https://github.com/aunyks/newton-api)  

Newton does anything from numerical calculation to symbolic math parsing.

# Installation
`go get github.com/jay9596/goNewton`

# Documentation
- [Newton API](https://github.com/aunyks/newton-api)
- Go wrapper API [GoDoc](https://godoc.org/github.com/Jay9596/goNewton) 

## Usage
Basic Trignometric calculation  
```newton := goNewton.New()
    cos, _ := newton.Cos("pi")
    sin, _ := newton.Sin("0")
    tan, _ := newton.Tan("45")
```
Error handling  
Error handling is very simple, just check it error returned is _nil_.
```
    tangent, err := newton.Tangent("x^3",2)
    if err != nil {
        // error handling
    }
```
Some merhods require additional parameters, in the newton API these are includede via "|" pipe, or ":" colon depending upon the endpoint. Here two different methods are present to call these endpoints.
```
    newton.Area("x^3",2,4) // Area of expression from start value to end value => 60
    newton.AreaStr("2:4|x^3") // => 60

    newton.Log(8,2) // log(8)<sub>2</sub> => 3
    newton.LogStr("2|8") // => 3
    
    newton.Tangent("x^3",2) // Tangent of expression at point => "12 x + -16"
    newton.TangentStr("2|x^3") // => "12 x + -16"
```
The methods with 'Str' in names use the newton API string format to send additional values.  


For more information about the _endpoints_ and _micro-service_, checkout the [newton Github repo](https://github.com/aunyks/newton-api)