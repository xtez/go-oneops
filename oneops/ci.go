package oneops

// CiAttributes represents atrributes of OneOps configuration item.
type CiAttributes struct {
	Adminstatus string `json:"adminstatus,omitempty"`
	Auth        string `json:"auth,omitempty"`
	Description string `json:"description,omitempty"`
	Location    string `json:"location,omitempty"`
}

func (ci CiAttributes) String() string {
	return Stringify(ci)
}
