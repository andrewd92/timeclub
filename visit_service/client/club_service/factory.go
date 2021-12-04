package club_service

var client ClubClient

func Instance() ClubClient {
	if nil == client {
		client = &ClubClientImpl{}
	}

	return client
}

func SetMockInstance(mockClient ClubClient) {
	client = mockClient
}
