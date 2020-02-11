package gbm

import (
	"errors"
	"strconv"
	"strings"
)

type Number interface {
	AscendPower(n int) Decimal
	DescendPower(n int) Decimal
}

type Int struct {
	RawData  string
	NegaFlag bool
	TenData  []rune
	BinData  []rune
}

type Decimal struct {
	RawData    string
	NegaFlag   bool
	FirstPart  Int
	SecondPart []rune
}

func InitInt(rawdata string) (Int, error) {
	var res Int = Int{NegaFlag: false}
	rawrunes := []rune(rawdata)
	tenrunes := []rune{}
	for ind, val := range rawrunes {
		if ind == 0 && (!strings.ContainsRune("+-0123456789", val)) {
			return res, errors.New("not a number")
		}
		if ind == 0 && strings.ContainsRune("+-", val) {
			if val == 45 {
				res.NegaFlag = true
			}
			continue
		}
		if !strings.ContainsRune("0123456789", val) {
			break
		}
		tenrunes = append(tenrunes, val)
	}
	if len(tenrunes) > 0 {
		res.RawData = rawdata
		res.TenData = TrimFrontChar(tenrunes, 48)
		BinData, err := ConvToBin(res.TenData)
		if err != nil {
			return res, errors.New("not a number")
		}
		res.BinData = BinData
		return res, nil
	}
	return res, errors.New("not a number")
}

func InitDecimal(rawdata string) (Decimal, error) {
	var res Decimal = Decimal{NegaFlag: false}
	rawrunes := []rune(rawdata)
	firstpart := []rune{}
	secondpart := []rune{}
	seesecondpart := 0
	for ind, val := range rawrunes {
		if ind == 0 && (!strings.ContainsRune("+-0123456789", val)) {
			return res, errors.New("not a number")
		}
		if ind == 0 && strings.ContainsRune("+-", val) {
			if val == 45 {
				res.NegaFlag = true
			}
			continue
		}
		if val == 46 {
			seesecondpart++
			continue
		}
		if (!strings.ContainsRune("0123456789", val)) || seesecondpart > 1 {
			break
		}
		if seesecondpart == 0 {
			firstpart = append(firstpart, val)
		}
		if seesecondpart == 1 {
			secondpart = append(secondpart, val)
		}
	}
	if len(firstpart) > 0 {
		res.RawData = rawdata
		res.FirstPart, _ = InitInt(string(firstpart))
		if len(secondpart) > 0 {
			if len(secondpart) > 8 {
				secondpart = secondpart[:8]
			}
			res.SecondPart = secondpart
		} else {
			res.SecondPart = []rune{48}
		}
		return res, nil
	}
	return res, errors.New("not a number")
}

func (a Int) ToDecimal() Decimal {
	var r = Decimal{
		RawData:    a.RawData,
		NegaFlag:   a.NegaFlag,
		FirstPart:  a,
		SecondPart: []rune{48},
	}
	return r
}

func (a Decimal) ToInt() Int {
	var r = Int{
		RawData:  a.RawData,
		NegaFlag: a.NegaFlag,
		TenData:  a.FirstPart.TenData,
		BinData:  a.FirstPart.BinData,
	}
	return r
}

func (a Int) FmtInt(separator string, split int) string {
	res := []rune{}
	i := 0
	sep := SReverseRune([]rune(separator))
	tenb := SReverseRune(a.TenData)
	for _, val := range tenb {
		if i == split {
			i = 0
			for _, v := range sep {
				res = append(res, v)
			}
		}
		res = append(res, val)
		i++
	}
	if a.NegaFlag {
		res = append(res, '-')
	}
	res = SReverseRune(res)
	r := string(res)
	return r
}

func (a Decimal) FmtDecimal(mode string, n int) string {
	r := []rune{}
	if a.NegaFlag {
		r = append(r, '-')
	}
	for _, val := range a.FirstPart.TenData {
		r = append(r, val)
	}
	if mode == "fixed" {
		if n > 0 {
			r = append(r, '.')
			i := 0
			for _, val := range a.SecondPart {
				i++
				r = append(r, val)
				if i == n {
					break
				}
			}
			for ; i < n; i++ {
				r = append(r, '0')
			}
			return string(r)
		} else {
			return string(r)
		}
	} else if mode == "max" {
		if n > 0 {
			r = append(r, '.')
			i := 0
			for _, val := range a.SecondPart {
				i++
				r = append(r, val)
				if i == n {
					break
				}
			}
			return string(r)
		} else {
			return string(r)
		}
	} else {
		r = append(r, '.')
		for _, val := range a.SecondPart {
			r = append(r, val)
		}
		return string(r)
	}
}

func NumberFmt(a Number, separator string, split int, mode string, n int) string {
	firstpart := []rune{}
	secondpart := []rune{}
	switch a1 := a.(type) {
	case Int:
		{
			i := 0
			sep := SReverseRune([]rune(separator))
			tenb := SReverseRune(a1.TenData)
			for _, val := range tenb {
				if i == split {
					i = 0
					for _, v := range sep {
						firstpart = append(firstpart, v)
					}
				}
				firstpart = append(firstpart, val)
				i++
			}
			if a1.NegaFlag {
				firstpart = append(firstpart, '-')
			}
			firstpart = SReverseRune(firstpart)
			if mode == "fixed" {
				if n > 0 {
					secondpart = append(secondpart, '.')
					for i := 0; i < n; i++ {
						secondpart = append(secondpart, '0')
					}
				}
			}
		}
	case Decimal:
		{
			i := 0
			sep := SReverseRune([]rune(separator))
			tenb := SReverseRune(a1.FirstPart.TenData)
			for _, val := range tenb {
				if i == split {
					i = 0
					for _, v := range sep {
						firstpart = append(firstpart, v)
					}
				}
				firstpart = append(firstpart, val)
				i++
			}
			if a1.NegaFlag {
				firstpart = append(firstpart, '-')
			}
			firstpart = SReverseRune(firstpart)
			if mode == "fixed" {
				if n > 0 {
					secondpart = append(secondpart, '.')
					i := 0
					for _, val := range a1.SecondPart {
						i++
						secondpart = append(secondpart, val)
						if i == n {
							break
						}
					}
					for ; i < n; i++ {
						secondpart = append(secondpart, '0')
					}
				}
			} else if mode == "max" {
				if n > 0 {
					secondpart = append(secondpart, '.')
					i := 0
					for _, val := range a1.SecondPart {
						i++
						secondpart = append(secondpart, val)
						if i == n {
							break
						}
					}
				}
			} else {
				secondpart = append(secondpart, '.')
				for _, val := range a1.SecondPart {
					secondpart = append(secondpart, val)
				}
			}
		}
	}
	res := append(firstpart, secondpart...)
	return string(res)
}

func NumberCompare(a,b Number) (comp string){
	switch a1:= a.(type) {
	case Int:{
		if b1,ok:=b.(Int);ok{
			if (!a1.NegaFlag) &&(b1.NegaFlag){
				return ">"
			}
			if (a1.NegaFlag) &&(!b1.NegaFlag){
				return "<"
			}
			if (!a1.NegaFlag) &&(!b1.NegaFlag){
				return string( BBCompare(a1.BinData,b1.BinData) )
			}
			if a1.NegaFlag && b1.NegaFlag {
				return string( BBCompare(b1.BinData,a1.BinData) )
			}
		}
		if b2,ok:= b.(Decimal);ok{
			if (!a1.NegaFlag) &&(b2.NegaFlag){
				return ">"
			}
			if (a1.NegaFlag) &&(!b2.NegaFlag){
				return "<"
			}
			if (!a1.NegaFlag) &&(!b2.NegaFlag){
				comp := BBCompare(a1.BinData,b2.FirstPart.BinData)
				if (comp ==62) ||(comp==60){
					return string(comp)
				}
				if strings.ContainsAny("123456789", string(b2.SecondPart) ){
					return "<"
				}
				return string(comp)
			}
			if a1.NegaFlag && b2.NegaFlag {
				comp := BBCompare(b2.FirstPart.BinData , a1.BinData)
				if (comp ==62) ||(comp==60){
					return string( comp)
				}
				if strings.ContainsAny("123456789", string(b2.SecondPart) ){
					return ">"
				}
				return string(comp)
			}
		}
	}
	case Decimal:{
		if b3,ok:=b.(Int);ok{
			comp := NumberCompare( b3, a1 )
			if comp ==">"{
				return "<"
			}
			if comp =="<"{
				return ">"
			}
			return string( comp)
		}
		if b4,ok:= b.(Decimal);ok{
			if (!a1.NegaFlag) &&(b4.NegaFlag){
				return ">"
			}
			if (a1.NegaFlag) &&(!b4.NegaFlag){
				return "<"
			}
			if (!a1.NegaFlag) &&(!b4.NegaFlag){
				comp := BBCompare(a1.FirstPart.BinData,b4.FirstPart.BinData)
				if (comp ==62) ||(comp==60){
					return string(comp)
				}
				t1 := '0'
				t2 := '0'
				for i:=0;i<8;i++{
					if i< len(a1.SecondPart){
						t1 = a1.SecondPart[i]
					}else {
						t1 = '0'
					}
					if i < len(b4.SecondPart){
						t2 = b4.SecondPart[i]
					}else {
						t2 = '0'
					}
					tn1,err := strconv.ParseInt( string(t1),10,64 )
					if err!=nil{
						tn1 = 0
					}
					tn2,err := strconv.ParseInt( string(t2),10,64 )
					if err!=nil{
						tn2 = 0
					}
					if tn1 > tn2{
						comp = 62
						break
					}
					if tn1 < tn2{
						comp = 60
						break
					}
				}
				return string(comp)
			}
			if a1.NegaFlag && b4.NegaFlag {
				comp := BBCompare(b4.FirstPart.BinData,a1.FirstPart.BinData)
				if (comp ==62) ||(comp==60){
					return string(comp)
				}
				t1 := '0'
				t2 := '0'
				for i:=0;i<8;i++{
					if i< len(a1.SecondPart){
						t1 = a1.SecondPart[i]
					}else {
						t1 = '0'
					}
					if i < len(b4.SecondPart){
						t2 = b4.SecondPart[i]
					}else {
						t2 = '0'
					}
					tn1,err := strconv.ParseInt( string(t1),10,64 )
					if err!=nil{
						tn1 = 0
					}
					tn2,err := strconv.ParseInt( string(t2),10,64 )
					if err!=nil{
						tn2 = 0
					}
					if tn1 > tn2{
						comp = 60
						break
					}
					if tn1 < tn2{
						comp = 62
						break
					}
				}
				return string(comp)
			}
		}
	}
	}
	return comp
}

func NumberCompareBool(a Number,b Number,s string)(bool,error){
	comp := NumberCompare(a,b)
	if (s==">") || (s=="<") || (s=="="){
		if string(comp) == s{
			return true,nil
		}else{
			return false,nil
		}
	} else if s==">="{
		if (comp==">"|| comp=="="){
			return true,nil
		}else{
			return false,nil
		}
	} else if s=="<="{
		if (comp=="<")||(comp=="="){
			return true,nil
		}else{
			return false,nil
		}
	} else{
		return false,errors.New("parametor c is wrong")
	}
}

func (a Int) AscendPower(n int) Decimal {
	var res Decimal
	newrawrunes := []rune{}
	if a.NegaFlag {
		newrawrunes = append(newrawrunes, 45)
	}
	for _, val := range a.TenData {
		newrawrunes = append(newrawrunes, val)
	}
	for i := 0; i < n; i++ {
		newrawrunes = append(newrawrunes, 48)
	}
	res, _ = InitDecimal(string(newrawrunes))
	return res
}

func (a Decimal) AscendPower(n int) Decimal {
	var res Decimal
	newraw := []rune{}
	if a.NegaFlag {
		newraw = append(newraw, 45)
	}
	secondpartlen := len(a.SecondPart)
	firstpart := []rune{}
	secondpart := []rune{}
	if secondpartlen > n {
		secondpart = a.SecondPart[n:]
		firstpart = a.SecondPart[0:n]
	} else {
		secondpart = []rune{48}
		firstpart = append([]rune{}, a.SecondPart[0:]...)
		for i := 0; i < (n - secondpartlen); i++ {
			firstpart = append(firstpart, 48)
		}
	}
	if len(a.FirstPart.TenData) == 1 && a.FirstPart.TenData[0] == 48 {
		firstpart = TrimFrontChar(firstpart, 48)
	} else {
		firstpart = append(a.FirstPart.TenData, firstpart...)
	}

	newraw = append(newraw, firstpart...)
	newraw = append(newraw, 46)
	newraw = append(newraw, secondpart...)
	res, _ = InitDecimal(string(newraw))
	return res
}

func (a Int) DescendPower(n int) Decimal {
	var res Decimal
	newraw := []rune{}
	if a.NegaFlag {
		newraw = append(newraw, 45)
	}
	alen := len(a.TenData)
	firstpart := []rune{}
	secondpart := []rune{}
	ind := 0
	if alen > n {
		ind = alen - n
		firstpart = a.TenData[0:ind]
		secondpart = a.TenData[ind:alen]
	} else {
		firstpart = []rune{'0'}
		for i := 0; i < (n - alen); i++ {
			secondpart = append(secondpart, 48)
		}
		secondpart = append(secondpart, a.TenData...)
	}
	newraw = append(newraw, firstpart...)
	newraw = append(newraw, 46)

	if len(secondpart) > 0 {
		secondpart = SReverseRune(TrimFrontChar(SReverseRune(secondpart), 48))
		if len(secondpart) > 8 {
			secondpart = secondpart[:8]
		}
		newraw = append(newraw, secondpart...)
	} else {
		newraw = append(newraw, 48)
	}
	res, _ = InitDecimal(string(newraw))
	return res
}

func (a Decimal) DescendPower(n int) Decimal {
	var res Decimal
	newraw := []rune{}
	if a.NegaFlag {
		newraw = append(newraw, 45)
	}
	firstpartlen := len(a.FirstPart.TenData)
	firstpart := []rune{}
	secondpart := []rune{}
	ind := firstpartlen - n
	if firstpartlen > n {
		firstpart = a.FirstPart.TenData[0:ind]
		secondpart = append(secondpart, a.FirstPart.TenData[ind:]...)
	} else {
		firstpart = []rune{48}
		for i := 0; i < n-firstpartlen; i++ {
			secondpart = append(secondpart, 48)
		}
		secondpart = append(secondpart, a.FirstPart.TenData[0:]...)
	}

	secondpart = append(secondpart, a.SecondPart...)

	newraw = append(newraw, firstpart...)
	newraw = append(newraw, 46)
	if len(secondpart) > 0 {
		secondpart = SReverseRune(TrimFrontChar(SReverseRune(secondpart), 48))
		if len(secondpart) > 8 {
			secondpart = secondpart[:8]
		}
		newraw = append(newraw, secondpart...)
	} else {
		newraw = append(newraw, 48)
	}
	res, _ = InitDecimal(string(newraw))
	return res
}

func NumberAdd(a, b Number) (r Number) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r = a1.AddInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r = a1.AddDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r = b3.AddDecimal(a1)
			}
			if b4, ok := b.(Decimal); ok {
				r = b4.AddDecimal(a1)
			}
		}
	}
	return r
}

func NumberSub(a, b Number) (r Number) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r = a1.SubInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r = a1.SubDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r = a1.SubInt(b3)
			}
			if b4, ok := b.(Decimal); ok {
				r = a1.SubDecimal(b4)
			}
		}
	}
	return r
}

func NumberMultip(a, b Number) (r Number) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r = a1.MultipInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r = a1.MultipDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r = a1.MultipInt(b3)
			}
			if b4, ok := b.(Decimal); ok {
				r = a1.MultipDecimal(b4)
			}
		}
	}
	return r
}

func NumberPower(a, b Number) (r Decimal) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r = a1.PowerInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r = a1.PowerDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r = a1.PowerInt(b3)
			}
			if b4, ok := b.(Decimal); ok {
				r = a1.PowerDecimal(b4)
			}
		}
	}
	return r
}

func NumberDivis(a, b Number) (r Decimal) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r = a1.DivisInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r = a1.DivisDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r = a1.DivisInt(b3)
			}
			if b4, ok := b.(Decimal); ok {
				r = a1.DivisDecimal(b4)
			}
		}
	}
	return r
}

func NumberMod(a, b Number) (r Int, err error) {
	switch a1 := a.(type) {
	case Int:
		{
			if b1, ok := b.(Int); ok {
				r, err = a1.ModInt(b1)
			}
			if b2, ok := b.(Decimal); ok {
				r, err = a1.ModDecimal(b2)
			}
		}
	case Decimal:
		{
			if b3, ok := b.(Int); ok {
				r, err = a1.ModInt(b3)
			}
			if b4, ok := b.(Decimal); ok {
				r, err = a1.ModDecimal(b4)
			}
		}
	}
	return r, err
}

func (a Int) AddInt(b Int) Int {
	if (!a.NegaFlag) && b.NegaFlag {
		b1 := Int{b.RawData, false, b.TenData, b.BinData}
		return a.SubInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Int{a.RawData, false, a.TenData, a.BinData}
		return b.SubInt(a1)
	}
	negaflag := false
	newraw := []rune{}
	if a.NegaFlag && b.NegaFlag {
		newraw = append(newraw, 45)
		negaflag = true
	}
	bindata := BBAdd(a.BinData, b.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	res := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	return res
}

func (a Int) AddDecimal(b Decimal) Decimal {
	if (!a.NegaFlag) && b.NegaFlag {
		b1 := Decimal{b.RawData, false, b.FirstPart, b.SecondPart}
		return a.SubDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Int{a.RawData, false, a.TenData, a.BinData}
		return b.SubInt(a1)
	}
	newraw := []rune{}
	negaflag := false
	if a.NegaFlag && b.NegaFlag {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	bindata := BBAdd(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	resdecimal := resint.DescendPower(8)
	return resdecimal
}

func (a Decimal) AddInt(b Int) Decimal {
	return b.AddDecimal(a)
}

func (a Decimal) AddDecimal(b Decimal) Decimal {
	if (!a.NegaFlag) && b.NegaFlag {
		b1 := Decimal{b.RawData, false, b.FirstPart, b.SecondPart}
		return a.SubDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Decimal{a.RawData, false, a.FirstPart, a.SecondPart}
		return b.SubDecimal(a1)
	}
	newraw := []rune{}
	negaflag := false
	if a.NegaFlag && b.NegaFlag {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	bindata := BBAdd(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	resdecimal := resint.DescendPower(8)
	return resdecimal
}

func (a Int) SubInt(b Int) Int {
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Int{b.RawData, false, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Int{b.RawData, true, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	comp := BBCompare(a.BinData, b.BinData)
	var r Int
	if a.NegaFlag && b.NegaFlag {
		if comp == 61 {
			r, _ = InitInt("0")
			return r
		}
		a1 := Int{a.RawData, false, a.TenData, a.BinData}
		b1 := Int{b.RawData, false, b.TenData, b.BinData}
		return b1.SubInt(a1)
	}
	negaflag := false
	newraw := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == 61 {
			r, _ = InitInt("0")
			return r
		}
		if comp == 60 {
			negaflag = true
			newraw = append(newraw, 45)
			bindata = BBMinus(b.BinData, a.BinData)
		}
		if comp == 62 {
			bindata = BBMinus(a.BinData, b.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	r = Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	return r
}

func (a Int) SubDecimal(b Decimal) Decimal {
	var r Decimal
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Decimal{b.RawData, false, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Decimal{b.RawData, true, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubInt(a)
	}
	negaflag := false
	newraw := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == 60 {
			negaflag = true
			newraw = append(newraw, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == 62 {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Decimal) SubInt(b Int) Decimal {
	var r Decimal
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Int{b.RawData, false, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Int{b.RawData, true, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubDecimal(a)
	}
	negaflag := false
	newraw := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == 60 {
			negaflag = true
			newraw = append(newraw, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == 62 {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Decimal) SubDecimal(b Decimal) Decimal {
	var r Decimal
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Decimal{b.RawData, false, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Decimal{b.RawData, true, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubDecimal(a)
	}
	negaflag := false
	newraw := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == 61 {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == 60 {
			negaflag = true
			newraw = append(newraw, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == 62 {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Int) MultipInt(b Int) Int {
	var r Int
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	bindata := BBMultip(a.BinData, b.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	r = Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	return r
}

func (a Int) MultipDecimal(b Decimal) Decimal {
	var r Decimal
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	b1 := b.AscendPower(8)
	bindata := BBMultip(a.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Decimal) MultipInt(b Int) Decimal {
	var r Decimal = b.MultipDecimal(a)
	return r
}

func (a Decimal) MultipDecimal(b Decimal) Decimal {
	var r Decimal
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	bindata := BBMultip(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(16)
	return r
}

func (a Int) PowerInt(b Int) Decimal {
	var r Decimal
	if (!b.NegaFlag) && (len(b.TenData) == 1) {
		if b.TenData[0] == 48 {
			r, _ = InitDecimal("1.0")
			return r
		}
		if b.TenData[0] == 49 {
			return a.ToDecimal()
		}
	}
	var i int64 = 1
	var limit int64
	resint := a
	limit, _ = strconv.ParseInt(string(b.TenData), 10, 64)
	for ; i < limit; i++ {
		resint = resint.MultipInt(a)
	}
	if b.NegaFlag {
		r, _ = InitDecimal("1.0")
		r = r.DivisInt(resint)
	} else {
		r = resint.ToDecimal()
	}
	return r
}

func (a Int) PowerDecimal(b Decimal) Decimal {
	var r Decimal = a.PowerInt(b.ToInt())
	return r
}

func (a Decimal) PowerInt(b Int) Decimal {
	var r Decimal
	if (!b.NegaFlag) && (len(b.TenData) == 1) {
		if b.TenData[0] == 48 {
			r, _ = InitDecimal("1.0")
			return r
		}
		if b.TenData[0] == 49 {
			return a
		}
	}
	var i int64 = 1
	var limit int64
	resdecimal := a
	limit, _ = strconv.ParseInt(string(b.TenData), 10, 64)
	for ; i < limit; i++ {
		resdecimal = resdecimal.MultipDecimal(a)
	}
	if b.NegaFlag {
		r, _ = InitDecimal("1.0")
		r = r.DivisDecimal(resdecimal)
	} else {
		r = resdecimal
	}
	return r
}

func (a Decimal) PowerDecimal(b Decimal) Decimal {
	var r Decimal = a.PowerInt(b.ToInt())
	return r
}

func (a Int) DivisInt(b Int) Decimal {
	var r Decimal
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(8)
	bindata := BBDivis(a1.FirstPart.BinData, b.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Int) DivisDecimal(b Decimal) (r Decimal) {
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(16)
	b1 := b.AscendPower(8)
	bindata := BBDivis(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Decimal) DivisInt(b Int) Decimal {
	var r Decimal
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(16)
	b1 := b.AscendPower(8)
	bindata := BBDivis(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Decimal) DivisDecimal(b Decimal) (r Decimal) {
	newraw := []rune{}
	negaflag := false
	if (a.NegaFlag && (!b.NegaFlag)) || (b.NegaFlag && (!a.NegaFlag)) {
		newraw = append(newraw, 45)
		negaflag = true
	}
	a1 := a.AscendPower(16)
	b1 := b.AscendPower(8)
	bindata := BBDivis(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newraw = append(newraw, tenrunes...)
	resint := Int{
		RawData:  string(newraw),
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}

func (a Int) ModInt(b Int) (r Int, e error) {
	bbcomp := BBCompare(a.BinData, b.BinData)
	if bbcomp == 60 {
		e = errors.New("mod,but a<b,fail")
		return r, e
	}
	if bbcomp == 61 {
		r, _ := InitInt("0")
		return r, nil
	}
	c := a.DivisInt(b).ToInt()
	a.NegaFlag = false
	b.NegaFlag = false
	c.NegaFlag = false
	r = a.SubInt(c.MultipInt(b))
	return r, nil
}

func (a Int) ModDecimal(b Decimal) (Int, error) {
	b1 := b.ToInt()
	r, err := a.ModInt(b1)
	return r, err
}

func (a Decimal) ModInt(b Int) (Int, error) {
	a1 := a.ToInt()
	r, err := a1.ModInt(b)
	return r, err
}

func (a Decimal) ModDecimal(b Decimal) (Int, error) {
	a1 := a.ToInt()
	b1 := b.ToInt()
	r, err := a1.ModInt(b1)
	return r, err
}
