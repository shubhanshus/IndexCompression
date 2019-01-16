package IndexCompression

import (
"bytes"
"io"
"math"
)

func encode(w io.Writer, n uint64) error {
	bytes := 0
	switch {
	case n < 128:
		bytes = 1
		n = (n << 1)
	case n < 16834:
		bytes = 2
		n = (n << 2) | 1
	case n < 2097152:
		bytes = 3
		n = (n << 3) | 3
	default:
		bytes = 4
		n = (n << 4) | 7
	}
	d := [4]byte{
		byte(n), byte(n >> 8), byte(n >> 16), byte(n >> 24),
	}
	_, err := w.Write(d[:bytes])
	return err
}


func decode(w io.Writer, n uint64) error {
	bytes := 0
	switch {
	case n < 128:
		bytes = 1
		n = (n << 1)
	case n < 16834:
		bytes = 2
		n = (n << 2) | 1
	case n < 2097152:
		bytes = 3
		n = (n << 3) | 3
	default:
		bytes = 4
		n = (n << 4) | 7
	}
	d := [4]byte{
		byte(n), byte(n>>8), byte(n>>16), byte(n>>24),
	}
	_, err := w.Write(d[:bytes])
	return err
}

func encodev2 (n int) []byte{
	var b []byte
	if n==0{
		return b
	}
	i:=0
	i= int(math.Log(float64(n))/math.Log(128)) + 1
	ret:=make([]byte, i)
	j:=i-1
	for j>=0{
		ret[j] = byte(n%128)
		j= j-1
		n/=128
	}
	ret [i-1] +=128
	return ret
}


func encodeArray(numbers []int) []byte{
	//var byt []byte
	b := new(bytes.Buffer)

	for _,num :=range numbers{
		b.Write(encodev2(num))
	}

	return b.Bytes()
}



func decodev2(byteArray []byte) []int{

	var numbers []int
	n:=0
	for _, b := range byteArray {
		if (b & 0xff) <128{
			n =128 *n +int(b);
		}else{
			num:=(128*n + ((int(b)-128) & 0xff))
			numbers= append(numbers, num)
			n=0;
		}
	}
	return numbers
}



//
//func main() {
//	xs := []int{25202150, 25202187, 25202199, 25202208, 25202320, 25202335, 25202458, 25203494, 25203567, 25203569, 25205062, 25205082}
//	//var b bytes.Buffer
//	//for _, x := range xs {
//	//	if err := encode(&b, x); err != nil {
//	//		panic(err)
//	//	}
//	//}
//	//fmt.Println(b.Bytes())
//
//	var by []byte
//
//	by=encodeArray(xs)
//
//	fmt.Println(by)
//
//	nums:=decodev2(by)
//	fmt.Println(nums)
//
//
//
//}
