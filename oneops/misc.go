package oneops

// AttrProps represents attribute properties of OneOps object.
type AttrProps struct{}

func (ap AttrProps) String() string {
	return Stringify(ap)
}

// AltNs represents alternate namespace of OneOps object.
type AltNs struct{}

func (an AltNs) String() string {
	return Stringify(an)
}
