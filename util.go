package gbm

import "strconv"

var mb = map[rune][]rune{
	48:[]rune{48},
	49:[]rune{49},
	50:[]rune{49,48},
	51:[]rune{49,49},
	52:[]rune{49,48,48},
	53:[]rune{49,48,48},
	54:[]rune{49,49,48},
	55:[]rune{49,49,49},
	56:[]rune{49,48,48,48},
	57:[]rune{49,48,48,49},
}

func TrimFrontChar(a []rune, char rune)([]rune){
	alen := len(a)
	res := []rune{}
	flag := true
	if alen ==1{
		return a
	}
	for _,val := range a{
		if val == char && flag {
			continue
		}
		flag = false
		res = append(res,val)
	}
	if alen>1 && len(res)==0{
		res = []rune{ char }
	}
	return res
}

func BNMultip(a,b []rune) (r []rune)  {
	alen := len(a)
	blen := len(b)
	if alen>blen{
		a,b = b,a
		alen,blen = blen,alen
	}
	r = []rune{48}
	sec :=make([]rune,0,10)
	var carryflag int64 = 0
	var cellnum int64 = 0
	for i:=(alen-1);i>=0;i--{
		sec = make([]rune,0,10)
		if a[i]==48{
			sec = append(sec,48)
		}else{
			for j:=(blen-1);j>=0;j--{
				cellnum = ma[ a[i] ] * ma[ b[j] ] + carryflag
				carryflag = 0
				if cellnum>9{
					carryflag = cellnum/10
					cellnum = cellnum - carryflag*10
					sec = append(sec,[]rune( strconv.FormatInt(cellnum,10) )[0] )
					if j==0{
						sec = append(sec,[]rune( strconv.FormatInt(carryflag,10) )[0] )
						carryflag = 0
					}
				}else{
					sec = append(sec,[]rune( strconv.FormatInt(cellnum,10) )[0] )
				}
			}
			sec = SReverseRune(sec)
			for k:=(alen-i);k>1;k--{
				sec = append(sec,48)
			}
		}
		sec = TrimFrontChar(sec,48)
		r = BNAdd(r,sec)
	}
	return r
}

func BBMultip(a,b []rune) (r []rune)  {
	alen := len(a)
	blen := len(b)
	if alen>blen{
		a,b = b,a
		alen,blen = blen,alen
	}
	r = make([]rune,0,10)
	r = append(r,48)
	sec :=make([]rune,0,10)
	for i:=(alen-1);i>=0;i--{
		sec = make([]rune,0,10)
		if a[i]==48{
			sec = append(sec,48)
		}else{
			for j:=(blen-1);j>=0;j--{
				if b[j]==49{
					sec = append(sec,49)
				}else{
					sec = append(sec,48)
				}
			}
			sec = SReverseRune(sec)
			for k:=(alen-i);k>1;k--{
				sec = append(sec,48)
			}
		}
		sec = TrimFrontChar(sec,48)
		r = BBAdd(r,sec)
	}
	return r
}

func BBCompare(a,b []rune) (r rune)  {
	alen := len(a)
	blen := len(b)
	if alen>blen{
		r = 62
	}else if alen<blen{
		r = 60
	}else{
		for i:=0;i<alen;i++{
			if a[i]==b[i]{
				r = 61
			}else if a[i]==49 && b[i]==48{
				r = 62
				break
			}else{
				r = 60
				break
			}
		}
	}
	return r
}

func BBMinus(a,b []rune) (r []rune)  {
	alen := len(a)
	blen := len(b)
	r = make([]rune,0,10)
	var carryflag rune = 48
	var j int = 0
	for i:=(alen-1);i>=0;i--{
		j++
		if j<=blen{
			if carryflag==48{
				if a[i]==48&& b[blen-j]==48{
					r = append(r,48)
					carryflag = 48
				}else if a[i]==48&& b[blen-j]==49{
					r = append(r,49)
					carryflag = 49
				}else if a[i]==49&& b[blen-j]==48{
					r = append(r,49)
					carryflag = 48
				}else{
					r = append(r,48)
					carryflag = 48
				}
			}else{
				if a[i]==48&& b[blen-j]==48{
					r = append(r,49)
					carryflag = 49
				}else if a[i]==48&& b[blen-j]==49{
					r = append(r,48)
					carryflag = 49
				}else if a[i]==49&& b[blen-j]==48{
					r = append(r,48)
					carryflag = 48
				}else{
					r = append(r,49)
					carryflag = 49
				}
			}
		}else{
			if carryflag==48{
				r = append(r,a[i])
				carryflag = 48
			}else{
				if a[i]==48{
					r = append(r,49)
					carryflag = 49
				}else{
					r = append(r,48)
					carryflag = 48
				}
			}
		}
	}
	r = SReverseRune(r)
	r = TrimFrontChar(r,48)
	return r
}

var ma = map[rune]int64{48: 0, 49: 1, 50: 2, 51: 3, 52: 4, 53: 5, 54: 6, 55: 7, 56: 8, 57: 9}

func BBAdd(a,b []rune) (r []rune) {
	alen := len(a)
	blen := len(b)
	if alen>blen{
		a,b = b,a
		alen,blen = blen,alen
	}
	r = make([]rune,0,10)
	var carryflag rune = 48
	var j int = 0
	for i:=(alen-1);i>=0;i--{
		j++
		if a[i]==48 && b[blen-j]==48{
			if carryflag == 48{
				r = append(r,48)
			}else{
				r = append(r,49)
			}
			carryflag = 48
		}else if a[i]==49 && b[blen-j]==49 {
			if carryflag == 48{
				r = append(r,48)
			}else{
				r = append(r,49)
			}
			carryflag = 49
		}else{
			if carryflag == 48{
				r = append(r,49)
				carryflag = 48
			}else{
				r = append(r,48)
				carryflag = 49
			}
		}
	}
	var abtimes = blen-alen
	if abtimes!=0{
		for i:=abtimes-1;i>=0;i--{
			if b[i]== 48 && carryflag==48{
				r = append(r,48)
				carryflag = 48
			}else if b[i]==48 && carryflag==48{
				r = append(r,48)
				carryflag = 49
			}else{
				r = append(r,49)
				carryflag = 48
			}
		}
	}
	if carryflag==49{
		r = append(r,49)
	}
	r = SReverseRune(r )
	return r
}

func BBDivis(a,b []rune)( []rune){
	x := []rune{'0'}
	y := []rune{'1'}
	sec := make([]rune,0,10)
	x1 := make([]rune,0,10)
	var compare rune
	for{
		x = BBAdd(x,y)
		sec = BBMultip(x,b)
		compare = BBCompare(a,sec)
		if compare==61{
			return x
		}else if compare==62{
			x1 = BBAdd(x,[]rune{49})
			sec = BBMultip( x1, b )
			compare = BBCompare(a,sec)
			if compare==61{
				return x1
			}else if compare==60{
				return x
			}else{
				y = BBMultip(y,[]rune{49,48})
			}
		}else{
			x = BBMinus(x,y)
			y = []rune{49}
		}
	}
}

func SReverseRune(s []rune) (r []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	r = s
	return r
}

func BNAdd(a,b []rune)(r []rune){
	alen := len(a)
	blen := len(b)
	if alen>blen{
		a,b = b,a
		alen,blen = blen,alen
	}
	r = make([]rune,0,10)
	var carryflag int64 = 0
	var j int = 0
	var secnum int64 = 0
	for i:=(alen-1);i>=0;i--{
		j++
		secnum = ma[ a[i] ] + ma[ b[blen-j] ] + carryflag
		if secnum > 9 {
			secnum = secnum - 10
			carryflag = 1
		} else {
			carryflag = 0
		}
		r = append(r,[]rune(strconv.FormatInt(secnum,10) )[0] )
	}
	var abtimes = blen-alen
	if abtimes!=0{
		for i:=abtimes-1;i>=0;i--{
			secnum = ma[ b[i] ] + carryflag
			if secnum > 9 {
				secnum = secnum - 10
				carryflag = 1
			} else {
				carryflag = 0
			}
			r = append(r,[]rune(strconv.FormatInt(secnum,10) )[0] )
		}
	}
	if carryflag==1{
		r = append(r,49 )
	}
	r = SReverseRune(r)
	return r
}

func BBMod(a,b []rune)(r []rune){
	times := BBDivis(a,b)
	r = BBMinus(a, BBMultip(b,times) )
	return r
}