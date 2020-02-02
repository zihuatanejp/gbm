package gbm

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
