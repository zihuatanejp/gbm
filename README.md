# gbm
go base math lib.

## Description

* Support infinite number operation, 
there is no limit on the numerical range of 32bit, int64bit.

* No confusing short, long, float, double, complex...
  Only human-friendly integers and decimals.
  Easy to format structure design，also provides convenient formatting methods.
  Strong compatibility. Easy to initialize,Simple to use method.
  
 * Decimal use precision can be as small as 8 digits after the decimal point, 
 and does not use the ieee754 standard, 
 so there is no annoying floating point operation deviation problem,
  and it will meet the human habit and automatically handle the number of decimal places.
 
 * Supports basic addition, subtraction, multiplication,
 division and remainder power and comparison operations
 
 * Supports negative number operations, 
 including negative powers of negative numbers
 (Except for the margin surplus, 
 the negative margin surplus scheme with differences between languages is not used. 
 Only the margin surplus scheme with no differences between positive numbers is used,
  but the fault tolerance is processed. The negative sign will be ignored directly, 
  so don't worry about errors)
  
 ### Usage
 
 ```go
num1,err := InitInt("12")
num2,err := InitInt("-12.3")
num3,err := InitDecimal("-12.3")
/*
num1:{ RawData:12, NegaFlag:false, TenData:[1,2], BinData:[1,1,0,0] }
num2:{ RawData:-12.3, NegaFlag:true, TenData:[1,2], BinData:[1, 1, 0, 0] }
num3:{ RawData:-12.3, NegaFlag:true, FirstPart:num1, SecondPart:[3] }
*/


t9,_ := InitInt("8759480")
t9.FmtInt(",",3) // "8,759,480"

t8,_ := InitDecimal("3.4955")
t8.FmtDecimal("fixed",2)// "3.49" 
t8.FmtDecimal("fixed",6)// "3.495500"
t8.FmtDecimal("max",0)// "3" 
t8.FmtDecimal("max",1)// "3.4"

t9,_ := InitDecimal("8759480.6977")
NumberFmt(t9,"-",3,"fixed",2)// "8-759-480.69"

t1,_ := InitInt("13")
t2,_ := InitInt("-15")
NumberCompare(t1,t2)// ">"
t3,err := NumberCompareBool(t1,t2,">")// true

t12,_ := InitInt("68438")
t13,_ := InitInt("234")
t12.AddInt(t13).RawData// "68672"

t10,_ := InitInt("40249")
t11,_ := InitInt("37252")
t10.SubInt(t11).RawData// "2997"

t5,_ := InitDecimal("9999999999999999999999999999999999999999999999999999999999999999999.99")
t6,_ := InitDecimal("0.01")
NumberAdd(t5,t6).(Decimal).RawData// "10000000000000000000000000000000000000000000000000000000000000000000.0"

t3,_ := InitDecimal("-1.1")
NumberMultip(t3,t3).(Decimal).RawData// "1.21"

t9,_ := InitInt("2")
t10,_ := InitInt("10")
NumberPower(t9,t10).RawData// "1024"

t1,_ := InitDecimal("10.0")
t2,_ := InitInt("3")
NumberDivis(t1,t2).RawData// "3.33333333"

```
  
 ## License
 
 Apache License 2.0
 
## Author

zihuatanejp 