package oneops

import "time"

// OrganizationsService handles communication with the Organization related
// methods of the OneOps API.
type OrganizationsService service

// Organization represents OneOps Organization
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
