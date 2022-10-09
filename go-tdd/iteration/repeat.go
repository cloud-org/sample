// author: ashing
// time: 2020/4/16 1:48 下午
// mail: axingfly@gmail.com
// Less is more.

package iteration

// repeat string
func Repeat(character string, num int) string {
	//return ""
	//return "xxxxx"
	repeated := ""
	// var repeated string
	for i := 0; i < num; i++ {
		//  repeated = repeated + character
		repeated += character
	}
	return repeated
}
