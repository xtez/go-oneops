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
	CiID             int          `json:"ciId,omitempty"`
	CiName           string       `json:"ciName,omitempty"`
	CiClassName      string       `json:"ciClassName,omitempty"`
	Impl             string       `json:"impl,omitempty"`
	NsPath           string       `json:"nsPath,omitempty"`
	CiGoid           string       `json:"ciGoid,omitempty"`
	Comments         string       `json:"comments,omitempty"`
	CiState          string       `json:"ciState,omitempty"`
	LastAppliedRfcID int          `json:"lastAppliedRfcId,omitempty"`
	CreatedBy        interface{}  `json:"createdBy,omitempty"`
	UpdatedBy        string       `json:"updatedBy,omitempty"`
	Created          int64        `json:"created,omitempty"`
	Updated          int64        `json:"updated,omitempty"`
	NsID             int          `json:"nsId,omitempty"`
	CiAttributes     CiAttributes `json:"ciAttributes,omitempty"`
	AttrProps        AttrProps    `json:"attrProps,omitempty"`
	AltNs            AltNs        `json:"altNs,omitempty"`
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
