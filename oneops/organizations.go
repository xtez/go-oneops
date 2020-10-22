package oneops

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// OrganizationsService handles communication with the Organization related
// methods of the OneOps API.
type OrganizationsService service

// Organization represents OneOps organization.
type Organization struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	CmsID        int       `json:"cms_id,omitempty"`
	Assemblies   bool      `json:"assemblies,omitempty"`
	Catalogs     bool      `json:"catalogs,omitempty"`
	Services     bool      `json:"services,omitempty"`
	Announcement string    `json:"announcement,omitempty"`
	FullName     string    `json:"full_name,omitempty"`
}

func (o Organization) String() string {
	return Stringify(o)
}

// Get fetches an organization by name.
func (s *OrganizationsService) Get(ctx context.Context, org string) (*Organization, *http.Response, error) {
	o := fmt.Sprintf("account/organizations/%v", org)

	req, err := s.client.NewRequest("GET", o, nil)
	if err != nil {
		return nil, nil, err
	}

	organization := new(Organization)
	resp, err := s.client.Do(ctx, req, organization)
	if err != nil {
		return nil, resp, err
	}

	return organization, resp, nil
}

// ListAll lists all organizations accessible by authenticated user.
func (s *OrganizationsService) ListAll(ctx context.Context) ([]*Organization, *http.Response, error) {
	o := "account/organizations"

	req, err := s.client.NewRequest("GET", o, nil)
	if err != nil {
		return nil, nil, err
	}
	var orgs []*Organization
	resp, err := s.client.Do(ctx, req, &orgs)
	if err != nil {
		return nil, resp, err
	}

	return orgs, resp, nil
}
