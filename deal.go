package gohubspot

type DealsService service

type Deal struct {
	PortalId   int                 `json:"portalId"`
	DealId     int                 `json:"dealId"`
	IsDeleted  bool                `json:"isDeleted"`
	Properties map[string]Property `json:"properties"`
}

type Deals struct {
	Deals []*Deal `json:"deals"`
}

func (s *DealsService) GetAll() ([]*Deal, error) {
	url := "/deals/v1/deal/paged"
	all := new(Deals)
	err := s.client.RunGet(url, all)
	return all.Deals, err
}
