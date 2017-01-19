//This package is used for the fraction example.
//It only implements the things used in the example and is
//no meant to be used as a stand-alone package.
package fraction

import "fmt"

//Fraction represents a fraction with an integer numerator and denominator.
//The names of the variables a lower-case.
//This means that they are private and cannot be accessed outside of this package.
//This is also true for naming of functions.
type Fraction struct {
	numerator int
	denominator int
}

//New returns a pointer to new Fraction
func New(num, denom int) *Fraction {
	return &Fraction{numerator: num, denominator: denom}
}

//String method is called by functions that use the stringer interface.
//Examples are the fmt and log print commands.
func (f *Fraction) String() string {
	//Use the fmt.Sprintf method to create a string representation of the fraction.
	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
}

//Add method receives a pointer to a fraction.
//Functions with receivers are similar to class methods in OO languages.
//To add two fractions foo and bar together and store in foo, use foo.Add(bar)
func (f *Fraction) Add(frac *Fraction) {
	f.numerator = f.denominator * frac.numerator + f.numerator * frac.denominator
	f.denominator *= frac.denominator
	f.simplify()
}

//Addr method takes two pointers to a fraction.
//To add two fractions foo and bar together and return the result, use fraction.Add(foo, bar)
func Addr(a, b *Fraction) *Fraction {
	f := &Fraction{numerator: a.denominator * b.numerator + a.numerator * b.denominator,
		denominator: a.denominator * b.denominator}
	f.simplify()
	return f
}

//GCD returns the greatest common divisor of two integers
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

//simplify a fraction
func (f *Fraction) simplify() {
	gcd := GCD(f.numerator, f.denominator)
	f.numerator /= gcd
	f.denominator /= gcd
}