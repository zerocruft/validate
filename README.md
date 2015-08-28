# zerocruft/validate
Quickly validate your struct
Version 0.0.1

```
go get github.com/zerocruft/validate
```


```go

func main() {
    b := Boom{}
    b.Name = "Nok"
    
    err, infractions := validate.ValidateRequired(b)
    if err != nil {
        log.Error(err)
        log.Print(infractions)
    }
}

type Boom struct {
	Name string `validate:"required"`
	Age  int
	Bam  string `validate:"required" json:"boom"`
}
```
