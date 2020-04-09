package gohubspot

type CompaniesService service

type Company struct {
	PortalId   int                 `json:"portalId"`
	CompanyId  int                 `json:"companyId"`
	IsDeleted  bool                `json:"isDeleted"`
	Properties map[string]Property `json:"properties"`
}

type Companies struct {
	Companies []*Company `json:"companies"`
}

func (s *CompaniesService) GetAll() ([]*Company, error) {
	url := "/companies/v2/companies/paged"

	req, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Set("propertyMode", "value_only")
	query.Set("property", "name")
	query.Add("property", "domain")
	query.Add("property", "website")
	query.Add("property", "city")
	query.Add("property", "state")
	query.Add("property", "country")
	query.Add("property", "industry")
	query.Add("property", "lifecyclestage")
	query.Add("property", "annualrevenue")
	query.Add("property", "num_notes")
	query.Add("property", "num_contacted_notes")
	query.Add("property", "notes_last_updated")
	query.Add("property", "notes_last_contacted")
	req.URL.RawQuery = query.Encode()

	all := new(Companies)
	err = s.client.Do(req, all)

	return all.Companies, err
}
