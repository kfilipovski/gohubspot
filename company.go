package gohubspot

type CompaniesService service

type Company struct {
	PortalId   int        `json:"portalId"`
	CompanyId  int        `json:"companyId"`
	IsDeleted  bool       `json:"isDeleted"`
	Properties Properties `json:"properties"`
}

type Companies struct {
	Companies []*Company `json:"companies"`
}

func (s *CompaniesService) GetAll() ([]*Company, error) {
	url := "/companies/v2/companies/paged"
	all := new(Companies)
	err := s.client.RunGet(url, all)
	return all.Companies, err
}
