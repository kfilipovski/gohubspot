package gohubspot

type DealsService service

type Deal struct {
	PortalId   int        `json:"portalId"`
	DealId     int        `json:"dealId"`
	IsDeleted  bool       `json:"isDeleted"`
	Properties Properties `json:"properties"`
}

type Deals struct {
	Deals []Company `json:"deals"`
}

func (s *DealsService) GetAll() (*Deals, error) {
	url := "/deals/v1/deal/paged"
	all := new(Deals)
	err := s.client.RunGet(url, all)
	return all, err
}
