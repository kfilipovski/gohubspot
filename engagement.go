package gohubspot

type EngagementsService service

type EngagementType string

const (
	EngagementTypeNote    EngagementType = "NOTE"
	EngagementTypeEmail   EngagementType = "EMAIL"
	EngagementTypeTask    EngagementType = "TASK"
	EngagementTypeMeeting EngagementType = "MEETING"
	EngagementTypeCall    EngagementType = "CALL"
)

type Engagement struct {
	Id         int        `json:"id"`
	PortalId   int        `json:"portalId"`
	DealId     int        `json:"dealId"`
	Active     bool       `json:"active"`
	Properties Properties `json:"properties"`
}

type Engagements struct {
	Results []*Engagement `json:"results"`
}

func (s *EngagementsService) GetAll() ([]*Engagement, error) {
	url := "/engagements/v1/engagements/paged"
	all := new(Engagements)
	err := s.client.RunGet(url, all)
	return all.Results, err
}
