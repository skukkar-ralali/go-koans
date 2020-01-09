package go_koans

func aboutBasics() {
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // in it there is nothing false

	var i int = __int__
	assert(i == -1) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == -1)

	assert(-5%2 == __int__)
	assert(-5*2/10 == __int__)
	assert(-1 == __int__)

	var x int
	assert(-1+x == __int__) // zero values are valued in Go

	var f float32
	assert(-1.0+f == __float32__) // for types of all types

	var s string
	s = "impossibly lame value"
	assert(s == __string__) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	c.s = "impossibly lame value"
	assert(-1+c.x == __int__)       // and types within composite types
	assert(-1.0+c.f == __float32__) // which match the other types
	assert(c.s == __string__)       // in a typical way
}
