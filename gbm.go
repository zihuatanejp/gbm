package gbm

import (
	"errors"
	"strings"
)

type Number interface {
	AscendPower(n int) Decimal
	DescendPower(n int) Decimal
}

type Int struct {
	OgnData  []rune
	NegaFlag bool
	TenData  []rune
	BinData  []rune
}

type Decimal struct {
	OgnData    []rune
	NegaFlag   bool
	FirstPart  Int
	SecondPart []rune
}

func InitInt(ogndata string) (Int, error) {
	var res Int = Int{NegaFlag: false}
	ogn := []rune(ogndata)
	tenrunes := []rune{}
	for ind, val := range ogn {
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
		res.OgnData = ogn
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

func InitDecimal(ogndata string) (Decimal, error) {
	var res Decimal = Decimal{NegaFlag: false}
	ogn := []rune(ogndata)
	firstpart := []rune{}
	secondpart := []rune{}
	seesecondpart := 0
	for ind, val := range ogn {
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
		res.OgnData = ogn
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
		OgnData:    a.OgnData,
		NegaFlag:   a.NegaFlag,
		FirstPart:  a,
		SecondPart: []rune{'0'},
	}
	return r
}

func (a Decimal) ToInt() Int {
	var r = Int{
		OgnData:  a.OgnData,
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

func (a Int) AscendPower(n int) Decimal {
	var res Decimal
	newognrune := []rune{}
	if a.NegaFlag {
		newognrune = append(newognrune, 45)
	}
	for _, val := range a.TenData {
		newognrune = append(newognrune, val)
	}
	for i := 0; i < n; i++ {
		newognrune = append(newognrune, 48)
	}
	res, _ = InitDecimal(string(newognrune))
	return res
}

func (a Decimal) AscendPower(n int) Decimal {
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag {
		newogn = append(newogn, 45)
	}
	secondpartlen := len(a.SecondPart)
	firstpart := []rune{}
	secondpart := []rune{}
	if secondpartlen > n {
		secondpart = a.SecondPart[n:]
		firstpart = a.SecondPart[0:n]
	} else {
		secondpart = []rune{'0'}
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

	newogn = append(newogn, firstpart...)
	newogn = append(newogn, 46)
	newogn = append(newogn, secondpart...)
	res, _ = InitDecimal(string(newogn))
	return res
}

func (a Int) DescendPower(n int) Decimal {
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag {
		newogn = append(newogn, 45)
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
	newogn = append(newogn, firstpart...)
	newogn = append(newogn, 46)

	if len(secondpart) > 0 {
		secondpart = SReverseRune(TrimFrontChar(SReverseRune(secondpart), 48))
		newogn = append(newogn, secondpart...)
	} else {
		newogn = append(newogn, '0')
	}
	res, _ = InitDecimal(string(newogn))
	return res
}

func (a Decimal) DescendPower(n int) Decimal {
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag {
		newogn = append(newogn, 45)
	}
	firstpartlen := len(a.FirstPart.TenData)
	firstpart := []rune{}
	secondpart := []rune{}
	ind := firstpartlen - n
	if firstpartlen > n {
		firstpart = a.FirstPart.TenData[0:ind]
		secondpart = append(secondpart, a.FirstPart.TenData[ind:]...)
	} else {
		firstpart = []rune{'0'}
		for i := 0; i < n-firstpartlen; i++ {
			secondpart = append(secondpart, 48)
		}
		secondpart = append(secondpart, a.FirstPart.TenData[0:]...)
	}

	secondpart = append(secondpart, a.SecondPart...)

	newogn = append(newogn, firstpart...)
	newogn = append(newogn, 46)
	if len(secondpart) > 0 {
		secondpart = SReverseRune(TrimFrontChar(SReverseRune(secondpart), 48))
		newogn = append(newogn, secondpart...)
	} else {
		newogn = append(newogn, 48)
	}
	res, _ = InitDecimal(string(newogn))
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

func (a Int) AddInt(b Int) Int {
	if (!a.NegaFlag) && b.NegaFlag {
		b1 := Int{b.OgnData, false, b.TenData, b.BinData}
		return a.SubInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Int{a.OgnData, false, a.TenData, a.BinData}
		return b.SubInt(a1)
	}
	negaflag := false
	newogn := []rune{}
	if a.NegaFlag && b.NegaFlag { // 负数+负数 -1 + -3 = -4
		newogn = append(newogn, '-')
		negaflag = true
	}
	bindata := BBAdd(a.BinData, b.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	res := Int{
		OgnData:  newogn,
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	return res
}

func (a Int) AddDecimal(b Decimal) Decimal {
	if (!a.NegaFlag) && b.NegaFlag {
		b1 := Decimal{b.OgnData, false, b.FirstPart, b.SecondPart}
		return a.SubDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Int{a.OgnData, false, a.TenData, a.BinData}
		return b.SubInt(a1)
	}
	newogn := []rune{}
	negaflag := false
	if a.NegaFlag && b.NegaFlag { // 负数+负数 -1 + -3 = -4
		newogn = append(newogn, '-')
		negaflag = true
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	bindata := BBAdd(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	resint := Int{
		OgnData:  newogn,
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
		b1 := Decimal{b.OgnData, false, b.FirstPart, b.SecondPart}
		return a.SubDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		a1 := Decimal{a.OgnData, false, a.FirstPart, a.SecondPart}
		return b.SubDecimal(a1)
	}
	newogn := []rune{}
	negaflag := false
	if a.NegaFlag && b.NegaFlag { // 负数+负数 -1 + -3 = -4
		newogn = append(newogn, '-')
		negaflag = true
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	bindata := BBAdd(a1.FirstPart.BinData, b1.FirstPart.BinData)
	tenrunes, _ := ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	resint := Int{
		OgnData:  newogn,
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	resdecimal := resint.DescendPower(8)
	return resdecimal
}

func (a Int) SubInt(b Int) Int {
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Int{b.OgnData, false, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Int{b.OgnData, true, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	comp := BBCompare(a.BinData, b.BinData)
	var r Int
	if a.NegaFlag && b.NegaFlag { //负数-负数 -2 - -1= -2 +1 = 1-2 = -1        -1 - -2 = -1 + 2 = 2-1 = 1     -1 - -1 =-1+1 =0
		if comp == '=' {
			r, _ = InitInt("0")
			return r
		}
		a1 := Int{a.OgnData, false, a.TenData, a.BinData}
		b1 := Int{b.OgnData, false, b.TenData, b.BinData}
		return b1.SubInt(a1)
	}
	negaflag := false
	newogn := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == '=' {
			r, _ = InitInt("0")
			return r
		}
		if comp == '<' {
			negaflag = true
			newogn = append(newogn, '-')
			bindata = BBMinus(b.BinData, a.BinData)
		}
		if comp == '>' {
			bindata = BBMinus(a.BinData, b.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	r = Int{
		OgnData:  newogn,
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	return r
}

func (a Int) SubDecimal(b Decimal) Decimal {
	var r Decimal
	if (!a.NegaFlag) && (b.NegaFlag) {
		b1 := Decimal{b.OgnData, false, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Decimal{b.OgnData, true, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubInt(a)
	}
	negaflag := false
	newogn := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == '<' {
			negaflag = true
			newogn = append(newogn, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == '>' {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	resint := Int{
		OgnData:  newogn,
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
		b1 := Int{b.OgnData, false, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) { //负数-正数 -2 - 1  = -2 + -1 =  -3
		b1 := Int{b.OgnData, true, b.TenData, b.BinData}
		return a.AddInt(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubDecimal(a)
	}
	negaflag := false
	newogn := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == '<' {
			negaflag = true
			newogn = append(newogn, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == '>' {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	resint := Int{
		OgnData:  newogn,
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
		b1 := Decimal{b.OgnData, false, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	if a.NegaFlag && (!b.NegaFlag) {
		b1 := Decimal{b.OgnData, true, b.FirstPart, b.SecondPart}
		return a.AddDecimal(b1)
	}
	a1 := a.AscendPower(8)
	b1 := b.AscendPower(8)
	comp := BBCompare(a1.FirstPart.BinData, b1.FirstPart.BinData)
	if a.NegaFlag && b.NegaFlag {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		a.NegaFlag = false
		b.NegaFlag = false
		return b.SubDecimal(a)
	}
	negaflag := false
	newogn := []rune{}
	var bindata, tenrunes []rune
	if (!a.NegaFlag) && (!b.NegaFlag) {
		if comp == '=' {
			r, _ = InitDecimal("0")
			return r
		}
		if comp == '<' {
			negaflag = true
			newogn = append(newogn, '-')
			bindata = BBMinus(b1.FirstPart.BinData, a1.FirstPart.BinData)
		}
		if comp == '>' {
			bindata = BBMinus(a1.FirstPart.BinData, b1.FirstPart.BinData)
		}
	}
	tenrunes, _ = ConvToTen(bindata)
	newogn = append(newogn, tenrunes...)
	resint := Int{
		OgnData:  newogn,
		NegaFlag: negaflag,
		TenData:  tenrunes,
		BinData:  bindata,
	}
	r = resint.DescendPower(8)
	return r
}
