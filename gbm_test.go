package gbm

import "testing"

func TestInitInt(t *testing.T) {
	t1,err := InitInt("0")
	if err!=nil{
		t.Error("init 0 to int error",err)
	}
	if t1.NegaFlag != false{
		t.Error("init 0 to int minus sign error")
	}
	if string(t1.OgnData)!="0"{
		t.Error("init 0 to int origin string wrong")
	}
	t2,err := InitInt("-1")
	if err!=nil{
		t.Error("init -1 to int error",err)
	}
	if t2.NegaFlag != true{
		t.Error("init -1 to int minus sign error")
	}
	if string(t2.OgnData)!="-1"{
		t.Error("init -1 to int origin string wrong")
	}
	if string(t2.TenData)!="1"{
		t.Error("init -1 to int ten-base-string wrong")
	}
	if string(t2.BinData)!="1"{
		t.Error("init -1 to int binary-string wrong")
	}
	t3,err := InitInt("-1.0")
	if err!=nil{
		t.Error("init -1.0 to int error",err)
	}
	if t3.NegaFlag != true{
		t.Error("init -1.0 to int minus sign error")
	}
	if string(t3.OgnData)!="-1.0"{
		t.Error("init -1.0 to int origin string wrong")
	}
	if string(t3.TenData)!="1"{
		t.Error("init -1.0 to int ten-base-string wrong")
	}
	if string(t3.BinData)!="1"{
		t.Error("init -1.0 to int binary-string wrong")
	}
	t4,err := InitInt("2.0")
	if err!=nil{
		t.Error("init 2.0 to int error",err)
	}
	if t4.NegaFlag != false{
		t.Error("init 2.0 to int minus sign error")
	}
	if string(t4.OgnData)!="2.0"{
		t.Error("init 2.0 to int origin string wrong")
	}
	if string(t4.TenData)!="2"{
		t.Error("init 2.0 to int ten-base-string wrong")
	}
	if string(t4.BinData)!="10"{
		t.Error("init 2.0 to int binary-string wrong")
	}
	t5,err := InitInt("433")
	if err!=nil{
		t.Error("init 433 to int error",err)
	}
	if t5.NegaFlag != false{
		t.Error("init 433 to int minus sign error")
	}
	if string(t5.OgnData)!="433"{
		t.Error("init 433 to int origin string wrong")
	}
	if string(t5.TenData)!="433"{
		t.Error("init 433 to int ten-base-string wrong")
	}
	if string(t5.BinData)!="110110001"{
		t.Error("init 433 to int binary-string wrong")
	}
	t6,_:= InitInt("324")
	if string(t6.BinData)!="101000100"{
		t.Error("init 324 to int binary-string wrong")
	}
}

func TestInitDecimal(t *testing.T) {
	_,err := InitDecimal("0.1")
	if err !=nil{
		t.Error("0.1 -> decimal fail",err)
	}
	_,err = InitDecimal("-0.1")
	if err !=nil{
		t.Error("-0.1 -> decimal fail",err)
	}
	_,err = InitDecimal("x0.1")
	if err ==nil{
		t.Error("x0.1 -> decimal should wrong")
	}
	t1,err := InitDecimal("3.5")
	if t1.NegaFlag==true{
		t.Error("3.5 -> decimal NegaFlag wrong",err)
	}
	if string( t1.FirstPart.TenData )!="3"{
		t.Error("3.5 -> decimal FirstPart wrong",err)
	}
	if string( t1.FirstPart.BinData )!="11"{
		t.Error("3.5 -> decimal FirstPart's BinData wrong",err)
	}
	if string(t1.SecondPart)!="5"{
		t.Error("3.5 -> decimal SecondPart wrong")
	}

	t2,err := InitDecimal("1.0")
	if string(t2.FirstPart.TenData)!="1"{
		t.Error("1.0 -> decimal FirstPart wrong",err)
	}
	if string(t2.SecondPart)!="0"{
		t.Error("1.0 -> decimal SecondPart wrong",err)
	}
	t3,err := InitDecimal("1")
	if string(t3.FirstPart.TenData)!="1"{
		t.Error("1 -> decimal FirstPart wrong",err)
	}
	if string(t3.SecondPart)!="0"{
		t.Error("1 -> decimal SecondPart wrong",err)
	}

	t4,err := InitDecimal("-0.0002")
	if t4.NegaFlag!=true{
		t.Error("-0.0002 -> decimal NegaFlag wrong",err)
	}
	if string( t4.FirstPart.TenData )!="0"{
		t.Error( "-0.0002 -> decimal FirstPart wrong",err )
	}
	if string( t4.SecondPart )!="0002"{
		t.Error( "-0.0002 -> decimal SecondPart wrong",err )
	}
}