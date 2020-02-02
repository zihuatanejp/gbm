package gbm

import (
	"errors"
	"strings"
)

type Number interface {
	AscendPower(n int) (Decimal)
	DescendPower(n int)(Decimal)
}

type Int struct {
	OgnData   []rune
	NegaFlag bool
	TenData []rune
	BinData []rune
}

type Decimal struct {
	OgnData []rune
	NegaFlag bool
	FirstPart Int
	SecondPart []rune
}

func InitInt(ogndata string) (Int,error) {
	var res Int = Int{NegaFlag:false}
	ogn := []rune(ogndata)
	tenrunes :=[]rune{}
	for ind,val := range ogn{
		if ind == 0 && (!strings.ContainsRune("+-0123456789",val) ){
			return res,errors.New("not a number")
		}
		if ind == 0 && strings.ContainsRune("+-",val){
			if val==45{
				res.NegaFlag = true
			}
			continue
		}
		if !strings.ContainsRune("0123456789",val){
			break
		}
		tenrunes = append(tenrunes,val)
	}
	if len(tenrunes)>0{
		res.OgnData = ogn
		res.TenData = TrimFrontChar(tenrunes,48)
		BinData,err := ConvToBin(res.TenData)
		if err!=nil{
			return res,errors.New("not a number")
		}
		res.BinData =BinData
		return res,nil
	}
	return res,errors.New("not a number")
}

func InitDecimal(ogndata string) (Decimal,error){
	var res Decimal = Decimal{NegaFlag:false}
	ogn := []rune(ogndata)
	firstpart := []rune{}
	secondpart := []rune{}
	seesecondpart := 0
	for ind,val := range ogn{
		if ind == 0 && (!strings.ContainsRune("+-0123456789",val) ){
			return res,errors.New("not a number")
		}
		if ind == 0 && strings.ContainsRune("+-",val){
			if val==45{
				res.NegaFlag = true
			}
			continue
		}
		if val ==46{
			seesecondpart++
			continue
		}
		if (!strings.ContainsRune("0123456789",val) ) || seesecondpart > 1{
			break
		}
		if seesecondpart ==0{
			firstpart = append(firstpart,val)
		}
		if seesecondpart==1{
			secondpart = append(secondpart,val)
		}
	}
	if len(firstpart)>0{
		res.OgnData = ogn
		res.FirstPart,_= InitInt( string(firstpart) )
		if len(secondpart)>0{
			if len(secondpart)>8{
				secondpart = secondpart[:8]
			}
			res.SecondPart = secondpart
		}else{
			res.SecondPart = []rune{48}
		}
		return res,nil
	}
	return res,errors.New("not a number")
}

func (a Int) AscendPower(n int) (Decimal){
	var res Decimal
	newognrune := []rune{}
	if a.NegaFlag{
		newognrune = append(newognrune,45)
	}
	for _,val := range a.TenData{
		newognrune = append(newognrune,val)
	}
	for i:=0;i<n;i++{
		newognrune = append(newognrune,48)
	}
	res,_ = InitDecimal( string(newognrune) )
	return res
}

func (a Decimal) AscendPower(n int) (Decimal){
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag{
		newogn = append(newogn,45)
	}
	secondpartlen := len(a.SecondPart)
	firstpart := []rune{}
	secondpart := []rune{}
	if secondpartlen>n{
		secondpart = a.SecondPart[n:]
		firstpart = a.SecondPart[0:n]
	}else{
		secondpart = []rune{'0'}
		firstpart= append([]rune{},a.SecondPart[0:]...)
		for i:=0;i<(n- secondpartlen);i++{
			firstpart = append( firstpart , 48 )
		}
	}
	if len( a.FirstPart.TenData )==1 && a.FirstPart.TenData[0]==48{
		firstpart = TrimFrontChar( firstpart ,48)
	}else{
		firstpart = append(  a.FirstPart.TenData ,firstpart... )
	}

	newogn = append(newogn,firstpart...)
	newogn = append(newogn,46)
	newogn = append(newogn,secondpart...)
	res,_ = InitDecimal( string(newogn) )
	return res
}

func (a Int) DescendPower(n int)(Decimal){
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag{
		newogn = append(newogn,45)
	}
	alen := len(a.TenData)
	firstpart := []rune{}
	secondpart := []rune{}
	ind := 0
	if alen>n{
		ind = alen-n
		firstpart = a.TenData[0:ind]
		secondpart = a.TenData[ind:alen]
	}else{
		firstpart = []rune{'0'}
		for i:=0;i<(n-alen);i++{
			secondpart = append(secondpart,48)
		}
		secondpart = append(secondpart,a.TenData...)
	}
	newogn = append(newogn,firstpart...)
	newogn = append(newogn,46)

	if len( secondpart )>0 {
		secondpart = SReverseRune( TrimFrontChar( SReverseRune(secondpart), 48) )
		newogn = append(newogn,secondpart...)
	}else{
		newogn = append(newogn,'0')
	}
	res,_ = InitDecimal( string(newogn) )
	return res
}

func (a Decimal) DescendPower(n int)(Decimal)  {
	var res Decimal
	newogn := []rune{}
	if a.NegaFlag{
		newogn = append(newogn,45)
	}
	firstpartlen := len(a.FirstPart.TenData)
	firstpart := []rune{}
	secondpart := []rune{}
	ind := firstpartlen-n
	if firstpartlen>n{
		firstpart = a.FirstPart.TenData[0:ind]
		secondpart = append( secondpart , a.FirstPart.TenData[ind:]... )
	}else{
		firstpart = []rune{'0'}
		for i:=0;i < n-firstpartlen;i++{
			secondpart = append(secondpart,48)
		}
		secondpart = append( secondpart , a.FirstPart.TenData[0:]... )
	}

	secondpart = append( secondpart , a.SecondPart... )

	newogn = append(newogn,firstpart...)
	newogn = append(newogn,46)
	if len( secondpart )>0 {
		secondpart = SReverseRune( TrimFrontChar( SReverseRune(secondpart), 48) )
		newogn = append(newogn,secondpart...)
	}else{
		newogn = append(newogn,48)
	}
	res,_ = InitDecimal( string(newogn) )
	return res
}