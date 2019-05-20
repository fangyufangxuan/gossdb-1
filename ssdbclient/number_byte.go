package ssdbclient

var (
	byt = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

//toNum
//
//将byte的内容转换为数字，只要解析ssdb协议的长度时使用
func toNum(bs []byte) int {
	re := 0
	for _, v := range bs {
		if v > '9' || v < '0' {
			return re
		}
		re = re*10 + byt[v]
	}
	return re
}
