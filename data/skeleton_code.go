package data

var GO_CODE = `package %s

import "fmt"
import "testing"

func main() {
    fmt.Println("Hello, world!")
}

func %s() {
	// %s
}

func Test%sBasic(t *testing.T) {

}
`
