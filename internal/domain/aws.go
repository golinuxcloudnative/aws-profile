package domain

import "fmt"

type Profile struct {
	name           string
	sso_start_url  string
	sso_region     string
	sso_account_id string
	sso_role_name  string
	region         string
}

type ProfileOptions struct {
	Sso_start_url  string
	Sso_region     string
	Sso_account_id string
	Sso_role_name  string
	Region         string
}

func NewProfile(name string, opts *ProfileOptions) *Profile {
	profile := &Profile{
		name: name,
	}
	if opts == nil {
		return profile
	}
	if opts.Sso_start_url != "" {
		profile.sso_start_url = opts.Sso_start_url
	}
	if opts.Sso_region != "" {
		profile.sso_region = opts.Sso_region
	}
	if opts.Sso_account_id != "" {
		profile.sso_account_id = opts.Sso_account_id
	}
	if opts.Sso_role_name != "" {
		profile.sso_role_name = opts.Sso_role_name
	}
	if opts.Region != "" {
		profile.region = opts.Region
	}

	return profile
}

func (p Profile) Name() string      { return p.name }
func (p Profile) Region() string    { return p.region }
func (p Profile) Url() string       { return p.sso_start_url }
func (p Profile) Role() string      { return p.sso_role_name }
func (p Profile) AccountId() string { return p.sso_account_id }

func (p Profile) String() string {
	return fmt.Sprintf("name = %s, sso_start_url = %s, sso_region = %s, sso_account_id = %s, sso_role_name = %s, region = %s\n", p.name, p.sso_start_url, p.sso_region, p.sso_account_id, p.sso_role_name, p.region)
}
