package oneops

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// EnvironmentsService handles communication with the environment related
// methods of the OneOps API.
type EnvironmentsService service

// Environment represent OneOps environment.
type Environment struct {
	CiID             int                                `json:"ciId,omitempty"`
	CiName           string                             `json:"ciName,omitempty"`
	CiClassName      string                             `json:"ciClassName,omitempty"`
	Impl             string                             `json:"impl,omitempty"`
	NsPath           string                             `json:"nsPath,omitempty"`
	CiGoid           string                             `json:"ciGoid,omitempty"`
	Comments         string                             `json:"comments,omitempty"`
	CiState          string                             `json:"ciState,omitempty"`
	LastAppliedRfcID int                                `json:"lastAppliedRfcId,omitempty"`
	CreatedBy        string                             `json:"createdBy,omitempty"`
	UpdatedBy        interface{}                        `json:"updatedBy,omitempty"`
	Created          int64                              `json:"created,omitempty"`
	Updated          int64                              `json:"updated,omitempty"`
	NsID             int                                `json:"nsId,omitempty"`
	CiAttributes     CiAttributes                       `json:"ciAttributes,omitempty"`
	AttrProps        AttrProps                          `json:"attrProps,omitempty"`
	AltNs            AltNs                              `json:"altNs,omitempty"`
	Clouds           map[int]EnvironmentCloudAttributes `json:"clouds,omitempty"`
}

// EnvironmentCloudAttributes represents of cloud association attributes of Environment.
type EnvironmentCloudAttributes struct {
	Adminstatus string `json:"adminstatus,omitempty"`
	DpmtOrder   string `json:"dpmt_order,omitempty"`
	PctScale    string `json:"pct_scale,omitempty"`
	Priority    string `json:"priority,omitempty"`
}

func (e Environment) String() string {
	return Stringify(e)
}

func (eca EnvironmentCloudAttributes) String() string {
	return Stringify(eca)
}

// Get fetches a environment by name.
func (s *EnvironmentsService) Get(ctx context.Context, org, assembly, environment string) (*Environment, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	if assembly == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	if environment == "" {
		return nil, nil, errors.New("environment name must be non-empty")
	}
	te := fmt.Sprintf("%v/assemblies/%v/transition/environments/%v", org, assembly, environment)

	req, err := s.client.NewRequest("GET", te, nil)
	if err != nil {
		return nil, nil, err
	}

	transitionEnvironment := new(Environment)
	resp, err := s.client.Do(ctx, req, transitionEnvironment)
	if err != nil {
		return nil, resp, err
	}

	return transitionEnvironment, resp, nil
}

// ListAll lists all environments of organization.
func (s *EnvironmentsService) ListAll(ctx context.Context, org, assembly string) ([]*Environment, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	if assembly == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	te := fmt.Sprintf("%v/assemblies/%v/transition/environments", org, assembly)

	req, err := s.client.NewRequest("GET", te, nil)
	if err != nil {
		return nil, nil, err
	}

	var transitionEnvironments []*Environment
	resp, err := s.client.Do(ctx, req, &transitionEnvironments)
	if err != nil {
		return nil, resp, err
	}

	return transitionEnvironments, resp, nil
}
