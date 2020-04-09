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
	query.Set("properties", "name")
	query.Add("properties", "domain")
	query.Add("properties", "website")
	query.Add("properties", "city")
	query.Add("properties", "state")
	query.Add("properties", "country")
	query.Add("properties", "industry")
	query.Add("properties", "annualrevenue")
	query.Add("properties", "num_notes")
	query.Add("properties", "num_contacted_notes")
	query.Add("properties", "notes_last_updated")
	query.Add("properties", "notes_last_contacted")
	req.URL.RawQuery = query.Encode()

	all := new(Companies)
	err = s.client.Do(req, all)

	return all.Companies, err
}
