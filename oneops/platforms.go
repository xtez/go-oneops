package oneops

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// PlatformsService handles communication with the platform related
// methods of the OneOps API.
type PlatformsService service

// Platform represent OneOps platform.
type Platform struct {
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

func (p Platform) String() string {
	return Stringify(p)
}

// Get fetches a platform by name.
func (s *PlatformsService) Get(ctx context.Context, org, assembly, platform string) (*Platform, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	if assembly == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	if platform == "" {
		return nil, nil, errors.New("platform name must be non-empty")
	}
	ap := fmt.Sprintf("%v/assemblies/%v/design/platforms/%v", org, assembly, platform)

	req, err := s.client.NewRequest("GET", ap, nil)
	if err != nil {
		return nil, nil, err
	}

	assemblyPlatform := new(Platform)
	resp, err := s.client.Do(ctx, req, assemblyPlatform)
	if err != nil {
		return nil, resp, err
	}

	return assemblyPlatform, resp, nil
}

// ListAll lists all assemblies of organization.
func (s *PlatformsService) ListAll(ctx context.Context, org, assembly string) ([]*Platform, *http.Response, error) {
	if org == "" {
		return nil, nil, errors.New("org name must be non-empty")
	}
	if assembly == "" {
		return nil, nil, errors.New("assembly name must be non-empty")
	}
	ap := fmt.Sprintf("%v/assemblies/%v/design/platforms", org, assembly)

	req, err := s.client.NewRequest("GET", ap, nil)
	if err != nil {
		return nil, nil, err
	}

	var assemblyPlatforms []*Platform
	resp, err := s.client.Do(ctx, req, &assemblyPlatforms)
	if err != nil {
		return nil, resp, err
	}

	return assemblyPlatforms, resp, nil
}
