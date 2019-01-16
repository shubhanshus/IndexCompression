package IndexCompression


import (
"fmt"
"math"
"os"
"strconv"
)

func encodeBIC(nums []int64, low int64, high int64) string{
	if len(nums) ==0{
		return ""
	}
	if len(nums)==1{
		encode_int(nums[0], low, high)
	}

	h:= int64(len(nums)/2)
	return encode_int(nums[h],low+h,high-int64(len(nums))+h+1)+encodeBIC(nums[:h], low, nums[h])+encodeBIC(nums[h+1:], nums[h]+1, high)
}

func encode_int(val int64, low int64, high int64) string{
	//fmt.Println("val:",val," low:",low," high:",high)

	if val<low || val>high{
		print("error val is incorrect")
		os.Exit(3)
	}
	nx:= val - low
	hx:= high -low
	nbits := int(math.Log2(float64(hx)))
	unused:= (1<<uint64((nbits+1))) - hx
	if (nx >= unused){
		nbits += 1
		nx += unused
	}
	//fmt.Println(nbits)
	//fmt.Println(nx)
	var s string
	for i := 1; i <= nbits; i++ {
		s = strconv.FormatInt(nx & 0x01, 10)+ s
		nx >>= 1
	}
	return s
}


func decode_int(b string,low int64, high int64) (int64, string){
	nbits := uint(math.Log2(float64(high-low)))
	if (nbits == 0){
		return low, b
	}
	if uint(len(b))<nbits{
		print("error encoding too short")
		os.Exit(3)
	}
	var nx int64
	if i, err := strconv.ParseInt(b[:nbits], 2, 64); err != nil {
		fmt.Println(err)
	} else {
		nx = i
	}
	b= b[nbits:]
	hx:= high - low
	unused:= int64((1<<(nbits+1)) - hx)
	if (nx >= unused){
		if len(b)<1{
			print("error encoding too short")
			os.Exit(3)
		}
		if i, err := strconv.ParseInt(string(b[0]), 2, 64); err != nil {
			fmt.Println(err)
		} else {
			nx = (nx << 1) +i
		}
		b = b[1:]
		nx -= unused
	}

	return low+nx, b
}


func decodeBIC(b string, n int64, low int64, high int64) ([]int64, string){

	var output []int64

	if n ==0 {
		return output, b
	}
	if n == 1{
		k, b := decode_int(b, low, high)
		return append(output, k), b
	}
	h := n/2
	f1 := h
	f2 := n-h-1

	mid, b := decode_int(b, low+f1, high-f2)
	v1, b := decodeBIC(b, f1, low, mid)
	v2, b := decodeBIC(b, f2, mid+1, high)

	output=append(output,v1...)
	output=append(output,mid)
	output=append(output,v2...)

	return output, b

}

//func main(){
//	test:=[]int64{25203470, 25203474, 25203475, 25203488, 25205135}
//	fmt.Println("input array ",test)
//	encoded:=encodeBIC(test,25203470,25205136)
//	fmt.Println("encoded string:",encoded)
//
//	output,_:=decodeBIC(encoded,5,25203470,25205136)
//
//	fmt.Println("output",output)
//
//}