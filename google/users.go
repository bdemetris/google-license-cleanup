package google

import (
	admin "google.golang.org/api/admin/directory/v1"
)

func (s *Service) getUsers(token string) (*admin.Users, error) {
	r, err := s.googleClient.Users.List().Customer("my_customer").MaxResults(500).
		OrderBy("email").PageToken(token).Do()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Service) GetAllUsers() ([]*admin.User, error) {
	var allUsers []*admin.User

	r, err := s.getUsers("")
	if err != nil {
		return nil, err
	}

	nextPageToken := r.NextPageToken

	for nextPageToken != "" {
		nr, err := s.getUsers(nextPageToken)
		if err != nil {
			return nil, err
		}
		allUsers = append(allUsers, nr.Users...)
		nextPageToken = nr.NextPageToken
	}

	return allUsers, nil
}
