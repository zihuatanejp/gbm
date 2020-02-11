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


```
  
 ## License
 
 Apache License 2.0
 
## Author

zihuatanejp 