package oneops

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// CloudsService handles communication with the cloud related
// methods of the OneOps API.
type CloudsService service

// Cloud represent OneOps cloud.
type Cloud struct {
	CiID             int          `json:"ciId,omitempty"`
	CiName           string       `json:"ciName,omitempty"`
	CiClassName      string       `json:"ciClassName,omitempty"`
	Impl             string       `json:"impl,omitempty"`
	NsPath           string       `json:"nsPath,omitempty"`
	CiGoid           string       `json:"ciGoid,omitempty"`
	Comments         string       `json:"comments,omitempty"`
	CiState          string       `json:"ciState,omitempty"`
	LastAppliedRfcID int          `json:"lastAppliedRfcId,omitempty"`
	CreatedBy        string       `json:"createdBy,omitempty"`
	UpdatedBy        interface{}  `json:"updatedBy,omitempty"`
	Created          int64        `json:"created,omitempty"`
	Updated          int64        `json:"updated,omitempty"`
	NsID             int          `json:"nsId,omitempty"`
	CiAttributes     CiAttributes `json:"ciAttributes,omitempty"`
	AttrProps        AttrProps    `json:"attrProps,omitempty"`
	AltNs            AltNs        `json:"altNs,omitempty"`
}

func (c Cloud) String() string {
	return Stringify(c)
}

// Get fetches a cloud by name.
func (s *CloudsService) Get(ctx context.Context, org, cloud string) (*Cloud, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	if cloud == "" {
		return nil, nil, errors.New("cloud name must be non-empty")
	}
	c := fmt.Sprintf("%v/clouds/%v", org, cloud)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	orgCloud := new(Cloud)
	resp, err := s.client.Do(ctx, req, orgCloud)
	if err != nil {
		return nil, resp, err
	}

	return orgCloud, resp, nil
}

// ListAll lists all clouds of organization.
func (s *CloudsService) ListAll(ctx context.Context, org string) ([]*Cloud, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	oc := fmt.Sprintf("%v/clouds", org)

	req, err := s.client.NewRequest("GET", oc, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgClouds []*Cloud
	resp, err := s.client.Do(ctx, req, &orgClouds)
	if err != nil {
		return nil, resp, err
	}

	return orgClouds, resp, nil
}
