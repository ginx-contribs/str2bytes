# str2bytes

str2bytes is a package support convert `string` to `[]byte`, and convert `[]byte` to `string`, both of them without memory copy.

## install
```bash
ge get github.com/ginx-contribs/str2bytes@latest
```

## usage
```go
package main

import (
	"github.com/ginx-contribs/str2bytes"
	"log"
)

func main() {
	bytes := str2bytes.Str2Bytes("hello world")
	log.Println(bytes)
	str := str2bytes.Bytes2Str(bytes)
	log.Println(str)
}
```