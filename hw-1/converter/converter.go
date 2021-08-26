package converter

import (
  //"strconv"
  "fmt"
)
type Converter struct{

}
func (c Converter) ConvertToString(value interface{}) string{
  var s string

  /*switch value.(type) {
  case int:
    s=strconv.FormatInt(int64(value),10)
  case int64:
    s=strconv.FormatInt(value,10)
  case uint:
    s=strconv.FormatUint(uint64(value),10)
  case uint:
    s=strconv.FormatUint(value,10)
  case bool:
    s=strconv.FormatBool(value)
  case float32:
    s=strconv.FormatFloat(value,'E', -1, 32)
  case float32:
    s=strconv.FormatFloat(value,'E', -1, 64)
	default:
		fmt.Println("unsupported value")
  }*/


  s = fmt.Sprint(value)

  return s
}
