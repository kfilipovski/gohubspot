package gohubspot

type EngagementsService service

type Engagement struct {
	Id          int      `json:"id"`
	PortalId    int      `json:"portalId"`
	Active      bool     `json:"active"`
	Type        string   `json:"type"`
	BodyPreview string   `json:"bodyPreview"`
	Source      string   `json:"source"`
	SourceId    string   `json:"sourceId"`
	CreatedAt   UnixTime `json:"createdAt"`
	LastUpdated UnixTime `json:"lastUpdated"`
	Timestamp   UnixTime `json:"timestamp"`
}

type Associations struct {
	ContactIds []int `json:"contactIds"`
}

type EngagementResult struct {
	Engagement   Engagement             `json:"engagement"`
	Associations Associations           `json:"associations"`
	Metadata     map[string]interface{} `json:"metadata"`
}

type Engagements struct {
	Results []*EngagementResult `json:"results"`
}

func (s *EngagementsService) GetAll() ([]*EngagementResult, error) {
	url := "/engagements/v1/engagements/paged"
	all := new(Engagements)
	err := s.client.RunGet(url, all)
	return all.Results, err
}
