package leetcode

import (
	"fmt"
	"gotest.tools/v3/assert"
	"io/ioutil"
	"testing"
)

var BigTxt string
var BigEnTxt string

func init() {
	b, err := ioutil.ReadFile("data_cn.txt") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	BigTxt = string(b) // convert content to a 'string'

	b, err = ioutil.ReadFile("hobbit.txt") // just pass the file name

	if err != nil {
		fmt.Print(err)
	}

	BigEnTxt = string(b) // convert content to a 'string'
}

//----------------------------------------------------------------------

func Test_AC_Index_BigTxt(t *testing.T) {
	AC := NewAC([]string{BigTxt})
	index, b := AC.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
	assert.Equal(t, b, 10)
}
func Test_StrBF_Index_BigTxt(t *testing.T) {
	StrBF := NewStrBF(BigTxt)
	index := StrBF.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}
func Test_StrBM_Index_BigTxt(t *testing.T) {
	StrBM := NewStrBM(BigTxt)
	index := StrBM.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}
func Test_StrHorspool_Index_BigTxt(t *testing.T) {
	StrHorspool := NewStrHorspool(BigTxt)
	index := StrHorspool.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}
func Test_StrKMP_Index_BigTxt(t *testing.T) {
	StrKMP := NewStrKMP(BigTxt)
	index := StrKMP.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}
func Test_StrRK_Index_BigTxt(t *testing.T) {
	var hash *ExclusiveOr
	StrRK := NewStrRK(BigTxt, hash)
	index := StrRK.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}
func Test_StrSunday_Index_BigTxt(t *testing.T) {
	StrSunday := NewStrSunday(BigTxt)
	index := StrSunday.Index([]byte("人界与这个世界的屏障被击穿"))
	assert.Equal(t, index, 10)
}

//----------------------------------------------------------------------

func Test_AC_Index_BigEnTxt(t *testing.T) {
	AC := NewAC([]string{"surprised"})
	index, search := AC.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 10)
	assert.Equal(t, search, 10)
}

//func Benchmark_AC_Index_BigEnTxt(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		AC := NewAC([]string{"surprised"})
//		index, search := AC.Index([]byte(BigEnTxt))
//		assert.Equal(b, index, 10)
//		assert.Equal(b, search, 10)
//	}
//}

func Test_StrBF_Index_BigEnTxt(t *testing.T) {
	StrBF := NewStrBF("surprised")
	index := StrBF.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

func Benchmark_Test_StrBF_Index_BigEnTxt(b *testing.B) {
	StrBF := NewStrBF("surprised")
	for i := 0; i < b.N; i++ {
		index := StrBF.Index([]byte(BigEnTxt))
		assert.Equal(b, index, 11513)
	}
}
func Test_StrBM_Index_BigEnTxt(t *testing.T) {
	StrBM := NewStrBM("surprised")
	index := StrBM.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

func Benchmark_Test_StrBM_Index_BigEnTxt(b *testing.B) {
	StrBM := NewStrBM("surprised")
	for i := 0; i < b.N; i++ {
		index := StrBM.Index([]byte(BigEnTxt))
		assert.Equal(b, index, 11513)
	}
}
func Test_StrHorspool_Index_BigEnTxt(t *testing.T) {
	StrHorspool := NewStrHorspool("surprised")
	index := StrHorspool.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

func Benchmark_Test_StrHorspool_Index_BigEnTxt(b *testing.B) {
	StrHorspool := NewStrHorspool("surprised")
	for i := 0; i < b.N; i++ {
		index := StrHorspool.Index([]byte(BigEnTxt))
		assert.Equal(b, index, 11513)
	}
}
func Test_StrKMP_Index_BigEnTxt(t *testing.T) {
	StrKMP := NewStrKMP("surprised")
	index := StrKMP.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

func Benchmark_Test_StrKMP_Index_BigEnTxt(b *testing.B) {
	StrKMP := NewStrKMP("surprised")
	for i := 0; i < b.N; i++ {
		index := StrKMP.Index([]byte(BigEnTxt))
		assert.Equal(b, index, 11513)
	}
}
func Test_StrRK_Index_BigEnTxt(t *testing.T) {
	var hash *ExclusiveOr
	StrRK := NewStrRK("surprised", hash)
	index := StrRK.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

//func Benchmark_Test_StrRK_Index_BigEnTxt(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		var hash *ExclusiveOr
//		StrRK := NewStrRK("surprised", hash)
//		index := StrRK.Index([]byte(BigEnTxt))
//		assert.Equal(b, index, 11513)
//	}
//}

func Test_StrSunday_Index_BigEnTxt(t *testing.T) {
	StrSunday := NewStrSunday("surprised")
	index := StrSunday.Index([]byte(BigEnTxt))
	assert.Equal(t, index, 11513)
}

func Benchmark_Test_StrSunday_Index_BigEnTxt(b *testing.B) {
	StrSunday := NewStrSunday("surprised")
	for i := 0; i < b.N; i++ {
		index := StrSunday.Index([]byte(BigEnTxt))
		assert.Equal(b, index, 11513)
	}
}

//----------------------------------------------------------------------

func Test_AC_Index_Normal(t *testing.T) {
	AC := NewAC([]string{"abc", "abe"})
	index, b := AC.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
	assert.Equal(t, b, 10)
}
func Test_StrBF_Index_Normal(t *testing.T) {
	StrBF := NewStrBF("十六")
	index := StrBF.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
func Test_StrBM_Index_Normal(t *testing.T) {
	StrBM := NewStrBM("十六")
	index := StrBM.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
func Test_StrHorspool_Index_Normal(t *testing.T) {
	StrHorspool := NewStrHorspool("十六")
	index := StrHorspool.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
func Test_StrKMP_Index_Normal(t *testing.T) {
	StrKMP := NewStrKMP("十六")
	index := StrKMP.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
func Test_StrRK_Index_Normal(t *testing.T) {
	var hash *ExclusiveOr
	a := ExclusiveOr(0)
	hash = &a
	StrRK := NewStrRK("十六", hash)
	index := StrRK.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
func Test_StrSunday_Index_Normal(t *testing.T) {
	StrSunday := NewStrSunday("十六")
	index := StrSunday.Index([]byte("大激动地说sad拿都拿的石窟和大家卡第十六届大连科技"))
	assert.Equal(t, index, 10)
}
