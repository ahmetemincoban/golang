/*
--- BASIC TYPES (BUILT-IN TYPES)
NUMERIC TYPES
uint8       the set of all unsigned  8-bit integers (0 to 255) = 256
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
int8        the set of all signed  8-bit integers (-128 to 127) = 256
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers
complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts
byte        alias for uint8
rune        alias for int32
STRING TYPES
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
The number of bytes is called the length of the string and is never negative.
BOOLEAN TYPES
A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.
--- COMPOSITE TYPES
Slice - Struct - Pointer - Function - Interface - Map - Channel - Array
*/

package main

import "fmt"

func main() {

	name:="Ahmet"
	secondName:="Emin"
	surname:="ST"
	age:=30
	isMale:=true
	      
	
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", secondName)
	fmt.Printf("%T\n", surname)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMale)
}

/*
--------------------GENERAL------------------------------------
%v	the value in a default format when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value

--------------------Boolean------------------------------------
%t	the word true or false

--------------------Integer------------------------------------
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"

--------Floating-point and complex constituents-----------------
%b	decimalless scientific notation with exponent a power of two, in the manner of strconv.FormatFloat with the 'b' format, e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

----String and slice of bytes (treated equivalently with these verbs)----
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte

--------------------Slice------------------------------------
%p	address of 0th element in base 16 notation, with leading 0x

--------------------Pointer------------------------------------
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers, formatting the value exactly as if it were an integer.

--------------------The default format for %v is:------------------------------------
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

*/