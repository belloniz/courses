package computer

type Spec struct {
	Maker string
	Price int
	model string // can't be accessed, because this field is not exported (first letter is lower case)
}
