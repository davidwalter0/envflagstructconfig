package flag

import (
	"fmt"
	"os"
	"testing"
  "reflect"
)

func TestParseSliceInt(t *testing.T) {
	T := reflect.TypeOf(SliceIntValue{}).Elem()
  fmt.Printf("%v %T\n", T,T)

	os.Clearenv()
  fmt.Println("Test for int SliceInt")
	var sliceInt = new(SliceIntValue)
	Var(sliceInt, "sliceInt", "use SliceInt")
  switch T.Kind() {
  case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
    sliceInt.Set("1,2,3")
  case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
    sliceInt.Set("1,2,3")
  case reflect.Float32, reflect.Float64:
    sliceInt.Set("1.1,2.2,3.3")
  case reflect.Bool:
    sliceInt.Set("false,true")
  }
	fmt.Printf("%v %T\n", *sliceInt, *sliceInt)
	// Parse()
	// Usage()
}

