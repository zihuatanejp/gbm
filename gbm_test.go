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

func TestInt_FmtInt(t *testing.T) {
	t1,_ := InitInt("23444332245")
	if t1.FmtInt("ip",2) !="2ip34ip44ip33ip22ip45"{
		t.Error("23444332245 -> 2ip34ip44ip33ip22ip45")
	}
	t2,_ := InitInt("-3273934723828449274")
	if t2.FmtInt(",",3)!="-3,273,934,723,828,449,274"{
		t.Error("-3273934723828449274 -> -3,273,934,723,828,449,274 fail")
	}
	t3,_ := InitInt("3")
	t4 := t3.FmtInt(",",3)
	if t4!="3"{
		t.Error("t4 fail")
	}
	t5 := t3.FmtInt("ok",2)
	if t5!="3"{
		t.Error("t5 fail")
	}
	t6,_ := InitInt("367")
	t7 := t6.FmtInt("you",2)
	if t7!="3you67"{
		t.Error("367 -> 3you67 fail")
	}
	t8,_ := InitInt("768")
	if t8.FmtInt(",",1)!="7,6,8"{
		t.Error("768 -> 7,6,8 fmt fail")
	}
}

func TestDecimal_FmtDecimal(t *testing.T) {
	t1,_ := InitDecimal("3.23")
	if t1.FmtDecimal("fixed",3)!="3.230"{
		t.Error("3.23 -> 3.230 fail")
	}
	if t1.FmtDecimal("fixed",2)!="3.23"{
		t.Error("3.23 -> 3.23 fail")
	}
	if t1.FmtDecimal("fixed",1)!="3.2"{
		t.Error("3.23 -> 3.2 fail")
	}
	if t1.FmtDecimal("fixed",4)!="3.2300"{
		t.Error("3.23 -> 3.2300 fail")
	}
	if t1.FmtDecimal("fixed",0)!="3"{
		t.Error("3.23 -> 3 fail")
	}
	if t1.FmtDecimal("max",1)!="3.2"{
		t.Error("3.23 -> 3.2 max fail")
	}
}

func TestNumberFmt(t *testing.T) {
	t1,_ := InitInt("32")
	if NumberFmt(t1,",",1,"fixed",3)!="3,2.000"{
		t.Error("32 -> 3,2.000 format fail")
	}
	t2,_ := InitDecimal("235323.897")
	t3 := NumberFmt(t2,",",3,"max",2)
	if t3!="235,323.89"{
		t.Error("235323.897 -> 235,323.89 format error")
	}
}

func TestInt_AscendPower(t *testing.T) {
	t1,_ := InitInt("5")
	if string( t1.AscendPower(0).FirstPart.TenData )!="5"{
		t.Error("5 Rise 0 powers firstpart wrong")
	}
	if string( t1.AscendPower(0).SecondPart )!="0"{
		t.Error("5 Rise 0 power secondpart wrong")
	}
	if string( t1.AscendPower(1).FirstPart.TenData )!="50"{
		t.Error("5 Rise 1 powers firstpart wrong")
	}
	if string( t1.AscendPower(1).SecondPart )!="0"{
		t.Error("5 Rise 1 power secondpart wrong")
	}
	if string( t1.AscendPower(2).FirstPart.TenData )!="500"{
		t.Error("5 Rise 2 powers firstpart wrong")
	}
	if string( t1.AscendPower(2).SecondPart )!="0"{
		t.Error("5 Rise 2 power secondpart wrong")
	}
	if string( t1.AscendPower(3).FirstPart.TenData )!="5000"{
		t.Error("5 Rise 3 powers firstpart wrong")
	}
	if string( t1.AscendPower(3).SecondPart )!="0"{
		t.Error("5 Rise 3 power secondpart wrong")
	}
	t2,_ := InitInt("-32")
	if string( t2.AscendPower(2).OgnData ) !="-3200"{
		t.Error("-32 Rise 2 powers == -3200 fail")
	}
	if t2.AscendPower(2).NegaFlag!=true{
		t.Error("-32 Rise 2 powers negative sign wrong")
	}
}