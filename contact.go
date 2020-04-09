package gohubspot

import (
	"fmt"
)

type ContactsService service

type Contact struct {
	Vid          int                 `json:"vid"`
	CanonicalVid int                 `json:"canonical-vid"`
	MergedVids   []int               `json:"merged-vids"`
	PortalID     int                 `json:"portal-id"`
	IsContact    bool                `json:"is-contact"`
	ProfileToken string              `json:"profile-token"`
	ProfileURL   string              `json:"profile-url"`
	Properties   map[string]Property `json:"properties"`
}

type Contacts struct {
	Contacts []*Contact `json:"contacts"`
}

func (s *ContactsService) Create(properties Properties) (*IdentityProfile, error) {
	url := "/contacts/v1/contact"
	res := new(IdentityProfile)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) Update(contactID int, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", contactID)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) UpdateByEmail(email string, properties Properties) error {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s/profile", email)
	return s.client.RunPost(url, properties, nil)
}

func (s *ContactsService) CreateOrUpdateByEmail(email string, properties Properties) (*Vid, error) {
	url := fmt.Sprintf("/contacts/v1/contact/createOrUpdate/email/%s", email)

	res := new(Vid)
	err := s.client.RunPost(url, properties, res)
	return res, err
}

func (s *ContactsService) DeleteById(id int) (*ContactDeleteResult, error) {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d", id)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(url, res)
	return res, err
}

func (s *ContactsService) DeleteByEmail(email string) (*ContactDeleteResult, error) {
	url := fmt.Sprintf("/contacts/v1/contact/email/%s", email)

	res := new(ContactDeleteResult)
	err := s.client.RunDelete(url, res)
	return res, err
}

func (s *ContactsService) Merge(primaryID, secondaryID int) error {
	url := fmt.Sprintf("/contacts/v1/contact/merge-vids/%d/", primaryID)
	secondary := struct {
		SecondaryID int `json:"vidToMerge"`
	}{
		SecondaryID: secondaryID,
	}

	return s.client.RunPost(url, secondary, nil)
}

func (s *ContactsService) GetAll() ([]*Contact, error) {
	url := "/contacts/v1/lists/all/contacts/all"

	req, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Set("propertyMode", "value_only")
	query.Set("property", "firstname")
	query.Add("property", "lastname")
	query.Add("property", "num_notes")
	query.Add("property", "num_contacted_notes")
	req.URL.RawQuery = query.Encode()

	all := new(Contacts)
	err = s.client.Do(req, all)

	return all.Contacts, err
}

func (s *ContactsService) GetById(id int) (*Contact, error) {
	url := fmt.Sprintf("/contacts/v1/contact/vid/%d/profile", id)
	res := new(Contact)
	err := s.client.RunGet(url, res)
	return res, err
}

func (s *ContactsService) GetByToken(token string) (*Contact, error) {
	url := fmt.Sprintf("/contacts/v1/contact/utk/%s/profile", token)
	res := new(Contact)
	err := s.client.RunGet(url, res)
	return res, err
}
