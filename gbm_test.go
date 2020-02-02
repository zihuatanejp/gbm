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
}