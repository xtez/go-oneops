package oneops

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// AssembliesService handles communication with the assembly related
// methods of the OneOps API.
type AssembliesService service

// Assembly represent OneOps Assembly.
type Assembly struct {
	CiID             int          `json:"ciId"`
	CiName           string       `json:"ciName"`
	CiClassName      string       `json:"ciClassName"`
	Impl             string       `json:"impl"`
	NsPath           string       `json:"nsPath"`
	CiGoid           string       `json:"ciGoid"`
	Comments         string       `json:"comments"`
	CiState          string       `json:"ciState"`
	LastAppliedRfcID int          `json:"lastAppliedRfcId"`
	CreatedBy        interface{}  `json:"createdBy"`
	UpdatedBy        string       `json:"updatedBy"`
	Created          int64        `json:"created"`
	Updated          int64        `json:"updated"`
	NsID             int          `json:"nsId"`
	CiAttributes     CiAttributes `json:"ciAttributes"`
	AttrProps        AttrProps    `json:"attrProps"`
	AltNs            AltNs        `json:"altNs"`
}

func (a Assembly) String() string {
	return Stringify(a)
}

// Get fetches a Assembly by name.
func (s *AssembliesService) Get(ctx context.Context, org, assembly string) (*Assembly, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	if assembly == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	a := fmt.Sprintf("%v/assemblies/%v", org, assembly)

	req, err := s.client.NewRequest("GET", a, nil)
	if err != nil {
		return nil, nil, err
	}

	orgAssembly := new(Assembly)
	resp, err := s.client.Do(ctx, req, orgAssembly)
	if err != nil {
		return nil, resp, err
	}

	return orgAssembly, resp, nil
}

// ListAll lists all clouds of organization.
func (s *AssembliesService) ListAll(ctx context.Context, org string) ([]*Assembly, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	oa := fmt.Sprintf("%v/assemblies", org)

	req, err := s.client.NewRequest("GET", oa, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgAssembly []*Assembly
	resp, err := s.client.Do(ctx, req, &orgAssembly)
	if err != nil {
		return nil, resp, err
	}

	return orgAssembly, resp, nil
}
