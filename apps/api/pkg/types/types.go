package types

// Term represents the lease term season.
type Term string

const (
	TermWinter Term = "Winter"
	TermSpring Term = "Spring"
	TermSummer Term = "Summer"
	TermFall   Term = "Fall"
)

// Gender represents supported gender values.
type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
)
