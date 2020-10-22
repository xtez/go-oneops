package oneops

// AttrProps represents attribute properties of OneOps object.
type AttrProps struct {
	Updated Updated `json:"updated,omitempty"`
}

func (ap AttrProps) String() string {
	return Stringify(ap)
}

// AltNs represents alternate namespace of OneOps object.
type AltNs struct{}

func (an AltNs) String() string {
	return Stringify(an)
}

// Updated represents update information of OneOps object.
type Updated struct {
	Adminstatus  string `json:"adminstatus,omitempty"`
	Autorepair   string `json:"autorepair,omitempty"`
	Autoreplace  string `json:"autoreplace,omitempty"`
	Autoscale    string `json:"autoscale,omitempty"`
	Availability string `json:"availability,omitempty"`
	Codpmt       string `json:"codpmt,omitempty"`
	Debug        string `json:"debug,omitempty"`
	Description  string `json:"description,omitempty"`
	Dpmtdelay    string `json:"dpmtdelay,omitempty"`
	GlobalDNS    string `json:"global_dns,omitempty"`
	Logging      string `json:"logging,omitempty"`
	Monitoring   string `json:"monitoring,omitempty"`
	Owner        string `json:"owner,omitempty"`
	Profile      string `json:"profile,omitempty"`
	Restricted   string `json:"restricted,omitempty"`
	Site         string `json:"site,omitempty"`
	Subdomain    string `json:"subdomain,omitempty"`
	Tags         string `json:"tags,omitempty"`
	Verify       string `json:"verify,omitempty"`
}
