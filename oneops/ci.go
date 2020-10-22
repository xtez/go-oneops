package oneops

// CiAttributes represents atrributes of OneOps configuration item.
type CiAttributes struct {
	Adminstatus  string `json:"adminstatus,omitempty"`
	Auth         string `json:"auth,omitempty"`
	Autorepair   string `json:"autorepair,omitempty"`
	Autoreplace  string `json:"autoreplace,omitempty"`
	Autoscale    string `json:"autoscale,omitempty"`
	Availability string `json:"availability,omitempty"`
	Codpmt       string `json:"codpmt,omitempty"`
	Debug        string `json:"debug,omitempty"`
	Description  string `json:"description,omitempty"`
	Dpmtdelay    string `json:"dpmtdelay,omitempty"`
	GlobalDNS    string `json:"global_dns,omitempty"`
	Location     string `json:"location,omitempty"`
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

func (ci CiAttributes) String() string {
	return Stringify(ci)
}
