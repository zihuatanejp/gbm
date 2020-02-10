package gbm

import "testing"

func TestInitInt(t *testing.T) {
	t1, err := InitInt("0")
	if err != nil {
		t.Error("init 0 to int error", err)
	}
	if t1.NegaFlag != false {
		t.Error("init 0 to int minus sign error")
	}
	if t1.RawData != "0" {
		t.Error("init 0 to int origin string wrong")
	}
	t2, err := InitInt("-1")
	if err != nil {
		t.Error("init -1 to int error", err)
	}
	if t2.NegaFlag != true {
		t.Error("init -1 to int minus sign error")
	}
	if t2.RawData != "-1" {
		t.Error("init -1 to int origin string wrong")
	}
	if string(t2.TenData) != "1" {
		t.Error("init -1 to int ten-base-string wrong")
	}
	if string(t2.BinData) != "1" {
		t.Error("init -1 to int binary-string wrong")
	}
	t3, err := InitInt("-1.0")
	if err != nil {
		t.Error("init -1.0 to int error", err)
	}
	if t3.NegaFlag != true {
		t.Error("init -1.0 to int minus sign error")
	}
	if t3.RawData != "-1.0" {
		t.Error("init -1.0 to int origin string wrong")
	}
	if string(t3.TenData) != "1" {
		t.Error("init -1.0 to int ten-base-string wrong")
	}
	if string(t3.BinData) != "1" {
		t.Error("init -1.0 to int binary-string wrong")
	}
	t4, err := InitInt("2.0")
	if err != nil {
		t.Error("init 2.0 to int error", err)
	}
	if t4.NegaFlag != false {
		t.Error("init 2.0 to int minus sign error")
	}
	if t4.RawData != "2.0" {
		t.Error("init 2.0 to int origin string wrong")
	}
	if string(t4.TenData) != "2" {
		t.Error("init 2.0 to int ten-base-string wrong")
	}
	if string(t4.BinData) != "10" {
		t.Error("init 2.0 to int binary-string wrong")
	}
	t5, err := InitInt("433")
	if err != nil {
		t.Error("init 433 to int error", err)
	}
	if t5.NegaFlag != false {
		t.Error("init 433 to int minus sign error")
	}
	if t5.RawData != "433" {
		t.Error("init 433 to int origin string wrong")
	}
	if string(t5.TenData) != "433" {
		t.Error("init 433 to int ten-base-string wrong")
	}
	if string(t5.BinData) != "110110001" {
		t.Error("init 433 to int binary-string wrong")
	}
	t6, _ := InitInt("324")
	if string(t6.BinData) != "101000100" {
		t.Error("init 324 to int binary-string wrong")
	}
}

func TestInitDecimal(t *testing.T) {
	_, err := InitDecimal("0.1")
	if err != nil {
		t.Error("0.1 -> decimal fail", err)
	}
	_, err = InitDecimal("-0.1")
	if err != nil {
		t.Error("-0.1 -> decimal fail", err)
	}
	_, err = InitDecimal("x0.1")
	if err == nil {
		t.Error("x0.1 -> decimal should wrong")
	}
	t1, err := InitDecimal("3.5")
	if t1.NegaFlag == true {
		t.Error("3.5 -> decimal NegaFlag wrong", err)
	}
	if string(t1.FirstPart.TenData) != "3" {
		t.Error("3.5 -> decimal FirstPart wrong", err)
	}
	if string(t1.FirstPart.BinData) != "11" {
		t.Error("3.5 -> decimal FirstPart's BinData wrong", err)
	}
	if string(t1.SecondPart) != "5" {
		t.Error("3.5 -> decimal SecondPart wrong")
	}

	t2, err := InitDecimal("1.0")
	if string(t2.FirstPart.TenData) != "1" {
		t.Error("1.0 -> decimal FirstPart wrong", err)
	}
	if string(t2.SecondPart) != "0" {
		t.Error("1.0 -> decimal SecondPart wrong", err)
	}
	t3, err := InitDecimal("1")
	if string(t3.FirstPart.TenData) != "1" {
		t.Error("1 -> decimal FirstPart wrong", err)
	}
	if string(t3.SecondPart) != "0" {
		t.Error("1 -> decimal SecondPart wrong", err)
	}

	t4, err := InitDecimal("-0.0002")
	if t4.NegaFlag != true {
		t.Error("-0.0002 -> decimal NegaFlag wrong", err)
	}
	if string(t4.FirstPart.TenData) != "0" {
		t.Error("-0.0002 -> decimal FirstPart wrong", err)
	}
	if string(t4.SecondPart) != "0002" {
		t.Error("-0.0002 -> decimal SecondPart wrong", err)
	}
}

func TestInt_FmtInt(t *testing.T) {
	t1, _ := InitInt("23444332245")
	if t1.FmtInt("ip", 2) != "2ip34ip44ip33ip22ip45" {
		t.Error("23444332245 -> 2ip34ip44ip33ip22ip45")
	}
	t2, _ := InitInt("-3273934723828449274")
	if t2.FmtInt(",", 3) != "-3,273,934,723,828,449,274" {
		t.Error("-3273934723828449274 -> -3,273,934,723,828,449,274 fail")
	}
	t3, _ := InitInt("3")
	t4 := t3.FmtInt(",", 3)
	if t4 != "3" {
		t.Error("t4 fail")
	}
	t5 := t3.FmtInt("ok", 2)
	if t5 != "3" {
		t.Error("t5 fail")
	}
	t6, _ := InitInt("367")
	t7 := t6.FmtInt("you", 2)
	if t7 != "3you67" {
		t.Error("367 -> 3you67 fail")
	}
	t8, _ := InitInt("768")
	if t8.FmtInt(",", 1) != "7,6,8" {
		t.Error("768 -> 7,6,8 fmt fail")
	}
}

func TestDecimal_FmtDecimal(t *testing.T) {
	t1, _ := InitDecimal("3.23")
	if t1.FmtDecimal("fixed", 3) != "3.230" {
		t.Error("3.23 -> 3.230 fail")
	}
	if t1.FmtDecimal("fixed", 2) != "3.23" {
		t.Error("3.23 -> 3.23 fail")
	}
	if t1.FmtDecimal("fixed", 1) != "3.2" {
		t.Error("3.23 -> 3.2 fail")
	}
	if t1.FmtDecimal("fixed", 4) != "3.2300" {
		t.Error("3.23 -> 3.2300 fail")
	}
	if t1.FmtDecimal("fixed", 0) != "3" {
		t.Error("3.23 -> 3 fail")
	}
	if t1.FmtDecimal("max", 1) != "3.2" {
		t.Error("3.23 -> 3.2 max fail")
	}
}

func TestNumberFmt(t *testing.T) {
	t1, _ := InitInt("32")
	if NumberFmt(t1, ",", 1, "fixed", 3) != "3,2.000" {
		t.Error("32 -> 3,2.000 format fail")
	}
	t2, _ := InitDecimal("235323.897")
	t3 := NumberFmt(t2, ",", 3, "max", 2)
	if t3 != "235,323.89" {
		t.Error("235323.897 -> 235,323.89 format error")
	}
}

func TestInt_AscendPower(t *testing.T) {
	t1, _ := InitInt("5")
	if string(t1.AscendPower(0).FirstPart.TenData) != "5" {
		t.Error("5 Rise 0 powers firstpart wrong")
	}
	if string(t1.AscendPower(0).SecondPart) != "0" {
		t.Error("5 Rise 0 power secondpart wrong")
	}
	if string(t1.AscendPower(1).FirstPart.TenData) != "50" {
		t.Error("5 Rise 1 powers firstpart wrong")
	}
	if string(t1.AscendPower(1).SecondPart) != "0" {
		t.Error("5 Rise 1 power secondpart wrong")
	}
	if string(t1.AscendPower(2).FirstPart.TenData) != "500" {
		t.Error("5 Rise 2 powers firstpart wrong")
	}
	if string(t1.AscendPower(2).SecondPart) != "0" {
		t.Error("5 Rise 2 power secondpart wrong")
	}
	if string(t1.AscendPower(3).FirstPart.TenData) != "5000" {
		t.Error("5 Rise 3 powers firstpart wrong")
	}
	if string(t1.AscendPower(3).SecondPart) != "0" {
		t.Error("5 Rise 3 power secondpart wrong")
	}
	t2, _ := InitInt("-32")
	if t2.AscendPower(2).RawData != "-3200" {
		t.Error("-32 Rise 2 powers == -3200 fail")
	}
	if t2.AscendPower(2).NegaFlag != true {
		t.Error("-32 Rise 2 powers negative sign wrong")
	}
}

func TestDecimal_AscendPower(t *testing.T) {
	t1, _ := InitDecimal("-3.567")
	if t1.AscendPower(0).NegaFlag != true {
		t.Error("-3.567 Rise 0 powers negative sign wrong")
	}
	if string(t1.AscendPower(0).FirstPart.TenData) != "3" {
		t.Error("-3.567 Rise 0 powers firstpart wrong")
	}
	if string(t1.AscendPower(0).SecondPart) != "567" {
		t.Error("-3.567 Rise 0 powers secondpart wrong")
	}
	t2, _ := InitDecimal("3.567")
	if t2.AscendPower(1).NegaFlag != false {
		t.Error("3.567 Rise 1 powers negative sign wrong")
	}
	if t2.AscendPower(1).RawData != "35.67" {
		t.Error("3.567 Rise 1 powers == 35.67 fail")
	}
	if t2.AscendPower(2).RawData != "356.7" {
		t.Error("3.567 Rise 2 powers == 356.7 fail")
	}
	if t2.AscendPower(3).RawData != "3567.0" {
		t.Error("3.567 Rise 3 powers == 3567.0 fail")
	}
	if t2.AscendPower(4).RawData != "35670.0" {
		t.Error("3.567 Rise 4 powers == 35670.0 fail")
	}
	if t2.AscendPower(5).RawData != "356700.0" {
		t.Error("3.567 Rise 5 powers == 356700.0 fail")
	}
	t3, _ := InitDecimal("-0.003")
	if t3.AscendPower(1).RawData != "-0.03" {
		t.Error("-0.003 Rise 1 powers == -0.03 fail")
	}

	if t3.AscendPower(2).RawData != "-0.3" {
		t.Error("-0.003 Rise 2 powers == -0.3 fail")
	}
	if t3.AscendPower(3).RawData != "-3.0" {
		t.Error("-0.003 Rise 3 powers == -3.0 fail")
	}
	if t3.AscendPower(4).RawData != "-30.0" {
		t.Error("-0.003 Rise 4 powers == -30.0 fail")
	}
	if t3.AscendPower(5).RawData != "-300.0" {
		t.Error("-0.003 Rise 5 powers == -300.0 fail")
	}
	t4, _ := InitDecimal("10.003")
	if t4.AscendPower(1).RawData != "100.03" {
		t.Error("10.003 Rise 1 powers == 100.03 fail")
	}
	if t4.AscendPower(2).RawData != "1000.3" {
		t.Error("10.003 Rise 2 powers == 1000.3 fail")
	}
	if t4.AscendPower(3).RawData != "10003.0" {
		t.Error("10.003 Rise 3 powers == 10003.0 fail")
	}
	if t4.AscendPower(4).RawData != "100030.0" {
		t.Error("10.003 Rise 4 powers == 100030.0 fail")
	}
	if t4.AscendPower(5).RawData != "1000300.0" {
		t.Error("10.003 Rise 5 powers == 1000300.0 fail")
	}
}

func TestInt_DescendPower(t *testing.T) {
	t2, _ := InitInt("55")
	if t2.DescendPower(0).RawData != "55.0" {
		t.Error("55 descend 0 powers fail")
	}
	if t2.DescendPower(1).RawData != "5.5" {
		t.Error("55 descend 1 powers == 5.5 fail")
	}
	if t2.DescendPower(2).RawData != "0.55" {
		t.Error("55 descend 2 powers == 0.55 fail")
	}
	if t2.DescendPower(3).RawData != "0.055" {
		t.Error("55 descend 3 powers == 0.055 fail")
	}
	if t2.DescendPower(4).RawData != "0.0055" {
		t.Error("55 descend 4 powers == 0.0055 fail")
	}
	if string(t2.DescendPower(4).FirstPart.TenData) != "0" {
		t.Error("55 descend 4 powers == 0.0055 firstpart wrong")
	}
	if string(t2.DescendPower(4).SecondPart) != "0055" {
		t.Error("55 descend 4 powers == 0.0055 secondpart wrong")
	}
	t5, _ := InitInt("-23")
	if t5.DescendPower(0).NegaFlag != true {
		t.Error("-23 descend 0 powers negative sign wrong")
	}
	if t5.DescendPower(0).RawData != "-23.0" {
		t.Error("-23 descend 0 powers fail")
	}
	if t5.DescendPower(0).FirstPart.RawData != "23" {
		t.Error("-23 descend 0 powers firstpart wrong")
	}
	if string(t5.DescendPower(0).SecondPart) != "0" {
		t.Error("-23 descend 0 powers secondpart wrong")
	}
	if t5.DescendPower(1).NegaFlag != true {
		t.Error("-23 descend 0 powers negative sign wrong")
	}
	if t5.DescendPower(1).RawData != "-2.3" {
		t.Error("-23 descend 1 powers fail")
	}
	if t5.DescendPower(1).FirstPart.RawData != "2" {
		t.Error("-23 descend 1 powers firstpart wrong")
	}
	if string(t5.DescendPower(1).SecondPart) != "3" {
		t.Error("-23 descend 1 powers secondpart wrong")
	}
	if t5.DescendPower(2).NegaFlag != true {
		t.Error("-23 descend 2 powers negative sign wrong")
	}
	if t5.DescendPower(2).RawData != "-0.23" {
		t.Error("-23 descend 2 powers fail")
	}
	if t5.DescendPower(2).FirstPart.RawData != "0" {
		t.Error("-23 descend 2 powers firstpart wrong")
	}
	if string(t5.DescendPower(2).SecondPart) != "23" {
		t.Error("-23 descend 2 powers secondpart wrong")
	}
	if t5.DescendPower(3).NegaFlag != true {
		t.Error("-23 descend 3 powers negative sign wrong")
	}
	if t5.DescendPower(3).RawData != "-0.023" {
		t.Error("-23 descend 3 powers fail")
	}
	if t5.DescendPower(3).FirstPart.RawData != "0" {
		t.Error("-23 descend 3 powers firstpart wrong")
	}
	if string(t5.DescendPower(3).SecondPart) != "023" {
		t.Error("-23 descend 3 powers secondpart wrong")
	}
	if t5.DescendPower(4).NegaFlag != true {
		t.Error("-23 descend 4 powers negative sign wrong")
	}
	if t5.DescendPower(4).RawData != "-0.0023" {
		t.Error("-23 descend 4 powers fail")
	}
	if t5.DescendPower(4).FirstPart.RawData != "0" {
		t.Error("-23 descend 4 powers firstpart wrong")
	}
	if string(t5.DescendPower(4).SecondPart) != "0023" {
		t.Error("-23 descend 4 powers secondpart wrong")
	}
}

func TestDecimal_DescendPower(t *testing.T) {
	t1, _ := InitDecimal("-3.567")
	if t1.DescendPower(0).NegaFlag != true {
		t.Error("-3.567 descend 0 powers negative sign wrong")
	}
	if string(t1.DescendPower(0).FirstPart.TenData) != "3" {
		t.Error("-3.567 descend 0 powers firstpart wrong")
	}
	if string(t1.DescendPower(0).SecondPart) != "567" {
		t.Error("-3.567 descend 0 powers secondpart wrong")
	}
	t2, _ := InitDecimal("3.567")
	if t2.DescendPower(1).NegaFlag != false {
		t.Error("3.567 descend 1 powers negative sign wrong")
	}
	if t2.DescendPower(1).RawData != "0.3567" {
		t.Error("3.567 descend 1 powers == 0.3567 fail")
	}
	t3, _ := InitDecimal("10.003")
	if t3.DescendPower(1).RawData != "1.0003" {
		t.Error("10.003 descend 1 powers == 1.0003 fail")
	}
	if t3.DescendPower(2).RawData != "0.10003" {
		t.Error("10.003 descend 2 powers ==0.10003 fail")
	}
	if t3.DescendPower(3).RawData != "0.010003" {
		t.Error("10.003 descend 3 powers ==0.010003 fail")
	}
	if t3.DescendPower(4).RawData != "0.0010003" {
		t.Error("10.003 descend 4 powers ==0.0010003 fail")
	}
	if t3.DescendPower(5).RawData != "0.00010003" {
		t.Error("10.003 descend 5 powers ==0.00010003 fail")
	}
	t4, _ := InitDecimal("356.7")
	if t4.DescendPower(1).RawData != "35.67" {
		t.Error("356.7 descend 1 powers == 35.67 fail")
	}
	if t4.DescendPower(2).RawData != "3.567" {
		t.Error("356.7 descend 2 powers == 3.567 fail")
	}
	if t4.DescendPower(3).RawData != "0.3567" {
		t.Error("356.7 descend 3 powers == 0.3567 fail")
	}
	if t4.DescendPower(4).RawData != "0.03567" {
		t.Error("356.7 descend 4 powers == 0.03567 fail")
	}
	if t4.DescendPower(5).RawData != "0.003567" {
		t.Error("356.7 descend 5 powers == 0.003567 fail")
	}
}

func TestNumberAdd(t *testing.T) {
	t1,_ := InitInt("68438")
	t2,_ := InitInt("234")
	if NumberAdd(t1,t2).(Int).RawData!="68672"{
		t.Error("68438 + 234 = 68672 fail")
	}
	t3,_ := InitDecimal("3.156")
	t4,_ := InitDecimal("1.544")
	if NumberAdd(t3,t4).(Decimal).RawData!="4.7"{
		t.Error("3.156 + 1.544 = 4.7 fail")
	}
}

func TestNumberSub(t *testing.T) {
	t1,_ := InitInt("7289")
	t2,_ := InitInt("4266")
	t3,_ := InitInt("3023")
	if NumberSub(t1,t2).(Int).RawData!=t3.RawData{
		t.Error("7289 - 4266 = 3023 fail")
	}
	if NumberSub(t1,t3).(Int).RawData!=t2.RawData{
		t.Error("7289 - 3023 = 4266 fail")
	}
	t4,_ := InitInt("577")
	t5,_ := InitInt("67")
	if NumberSub(t4,t5).(Int).RawData!="510"{
		t.Error("577 - 67 = 510 fail")
	}
	t6,_ := InitInt("2505")
	t7,_ := InitInt("700")
	if NumberSub(t6,t7).(Int).RawData!="1805"{
		t.Error(" 2505 - 700 = 1805 fail")
	}
	t8,_ := InitInt("31258")
	t9,_ := InitInt("28100")
	if NumberSub(t8,t9).(Int).RawData!="3158"{
		t.Error("31258 - 28100 = 3158 fail")
	}
	t10,_ := InitInt("3158")
	if NumberSub(t8,t10).(Int).RawData!="28100"{
		t.Error("31258 - 3158 =  28100 fail")
	}
	t11,_ := InitInt("40249")
	t12,_ := InitInt("37252")
	if NumberSub(t11,t12).(Int).RawData!="2997"{
		t.Error("40249 - 37252 = 2997 fail")
	}
	tt1,_ := InitDecimal("4.7")
	tt2,_ := InitDecimal("3.156")

	if tt1.SubDecimal(tt2).RawData!="1.544"{
		t.Error("4.7 - 3.156 = 1.544 fail")
	}
	tt4,_:=InitInt("-0")
	if tt1.SubInt(tt4).RawData!="4.7"{
		t.Error("4.7 - -0 =4.7 fail")
	}
}

func TestNumberMultip(t *testing.T) {
	t1,_ := InitInt("577")
	t2,_ := InitInt("67")
	if NumberMultip(t1,t2).(Int).RawData!="38659"{
		t.Error("577 * 67 = 38659 fail")
	}
	t3,_ := InitDecimal("-1.1")
	if NumberMultip(t3,t3).(Decimal).RawData!="1.21"{
		t.Error("-1.1 * -1.1 = 1.21 fail")
	}
	t4,_ := InitDecimal("1.08")
	t5,_ := InitInt("34266")
	if NumberMultip(t4,t5).(Decimal).RawData!="37007.28"{
		t.Error("1.08 * 34266 = 37007.28 fail")
	}
	t6,_ := InitInt("0")
	if NumberMultip(t5,t6).(Int).RawData!="0"{
		t.Error("34266 * 0 = 0")
	}
}

func TestNumberPower(t *testing.T) {
	t1,_ := InitInt("10")
	t2,_ := InitInt("3")
	if NumberPower(t1,t2).RawData!="1000"{
		t.Error("10 ** 3 = 1000")
	}
	t3,_ := InitDecimal("10.0")
	t4,_ := InitDecimal("3.0")
	if NumberPower(t3,t4).RawData!="1000.0"{
		t.Error("10.0 ** 3.0 = 1000.0 fail")
	}
	if NumberPower(t3,t2).RawData!="1000.0"{
		t.Error("10.0 ** 3 = 1000.0 fail")
	}
	t5,_ := InitDecimal("1.15")
	t6,_ := InitInt("22")
	if NumberPower(t5,t6).RawData!="21.64474527"{
		t.Error("1.15 ** 22 = 21.64474527 fail")
	}
}

func TestInt_AddInt(t *testing.T) {
	t1, _ := InitInt("0")
	if t1.AddInt(t1).RawData != "0" {
		t.Error("0+0=0 fail")
	}
	t2, _ := InitInt("1")
	if t1.AddInt(t2).RawData != "1" {
		t.Error("0+1=1 fail")
	}
	if t2.AddInt(t1).RawData != "1" {
		t.Error("1+0=1 fail")
	}
	t3, _ := InitInt("3")
	if t3.AddInt(t2).RawData != "4" {
		t.Error("3+1=4 fail")
	}
	t4, _ := InitInt("-1")
	if t3.AddInt(t4).RawData != "2" {
		t.Error("3 + -1 =2 fail")
	}
	if t4.AddInt(t3).RawData != "2" {
		t.Error("-1 + 3 =2 fail")
	}
	t5, _ := InitInt("-3")
	if t5.AddInt(t2).RawData != "-2" {
		t.Error("-3 + 1 = -2 fail")
	}
	if t5.AddInt(t4).RawData != "-4" {
		t.Error("-3 + -1 = -4 fail")
	}
	t6, _ := InitInt("20")
	t7, _ := InitInt("11")
	t8, _ := InitInt("-16")
	if t6.AddInt(t7).RawData != "31" {
		t.Error("20+11=31 fail")
	}
	if t6.AddInt(t8).RawData != "4" {
		t.Error("20 + -16 =4 fail")
	}
	if t7.AddInt(t8).RawData != "-5" {
		t.Error("11 + -16 = -5 fail")
	}
	t9, _ := InitInt("23")
	t10, _ := InitInt("39")
	if t9.AddInt(t10).RawData != "62" {
		t.Error("23 + 39 = 62 fail")
	}
	t11, _ := InitInt("-39")
	if t9.AddInt(t11).RawData != "-16" {
		t.Error("23 + -39 = -16 fail")
	}
}

func TestInt_SubInt(t *testing.T) {
	t1, _ := InitInt("3")
	t2, _ := InitInt("1")
	t3, _ := InitInt("-1")
	t4, _ := InitInt("-2")
	if t1.SubInt(t2).RawData != "2" {
		t.Error("3 - 1 =2 fail")
	}
	if t2.SubInt(t1).RawData != "-2" {
		t.Error("1 - 3 = -2 fail")
	}
	if t1.SubInt(t3).RawData != "4" {
		t.Error("3 - -1 = 4 fail")
	}
	if t4.SubInt(t2).RawData != "-3" {
		t.Error("-2 - 1 = -3 fail")
	}
	if t4.SubInt(t3).RawData != "-1" {
		t.Error("-2 - -1 = -1 fail")
	}
	if t3.SubInt(t4).RawData != "1" {
		t.Error("-1 - -2 = 1 fail")
	}
	if t3.SubInt(t3).RawData != "0" {
		t.Error("-1 - -1 = 0 fail ")
	}
	t8, _ := InitInt("68")
	t9, _ := InitInt("38")
	if t8.SubInt(t9).RawData != "30" {
		t.Error("68 - 38 = 30 fail")
	}
}
