# zerocruft/validate
Quickly validate your struct

Version 0.0.1
[![Build Status](https://travis-ci.org/zerocruft/validate.svg?branch=master)](https://travis-ci.org/zerocruft/validate)

```
go get github.com/zerocruft/validate
```


```go

func main() {
    b := Boom{}
    b.Name = "Nok"
    
    err, infractions := validate.Required(b)
    if err != nil {
        log.Print(err)
        log.Print(infractions)
    }
}

type Boom struct {
	Name string `validate:"required"`
	Age  int
	Bam  string `validate:"required" json:"boom"`
}
```
