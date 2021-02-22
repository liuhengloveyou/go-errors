errors
======
For replacement offical package of `errors`


#### example
```go
package errors

import (
	"fmt"
	"testing"
)

var (
	ErrParseTest  = TN(10001, "hello {{.param1}}")
	ErrParseTest2 = Error{Code: 10002, Message: "test error"}
)

func TestAll(t *testing.T) {
	e1 := ErrParseTest.New(Params{"param1": "world aaa"})
	fmt.Println(e1.Error())

	fmt.Println(ErrParseTest2.Code, ErrParseTest2.Error())
}


```

#### example output
```bash
$ go-errors git:(main) ✗ go test
10001 hello world aaa
10002 10002 test error
```