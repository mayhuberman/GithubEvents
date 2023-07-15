package services

import (
	"fmt"
	"githubEventsListener/models"
	"reflect"
	"testing"
)

var expectedEmails = map[string]bool{
	"mockuser1@example.com": true,
	"mockuser2@example.com": true,
}
var expectedActor = "mock2"
var expectedUrl = "url2"

func TestHappyFlow(t *testing.T) {
	ActorsSize = 1
	UrlsSize = 1

	mockEventGetter := &MockEventGetter{}
	mockEvents := mockEventGetter.GetEvents()
	fmt.Println(mockEvents)

	ExtractData(mockEvents)

	if !reflect.DeepEqual(EventTypes, expectedEventTypes()) {
		t.Errorf("Result of event types (%v) does not match expected (%v)", EventTypes, expectedEventTypes())
	}

	if !reflect.DeepEqual(Emails, expectedEmails) {
		t.Errorf("Result of email (%v) does not match expected (%v)", Emails, expectedEmails)
	}

	if ActorsList.Size() != ActorsSize {
		t.Errorf("Expected actors list size: %d, Got: %d", ActorsSize, ActorsList.Size())
	}
	actor := ActorsList.GetHeadDataAndMoveNext()
	if actor != expectedActor {
		t.Errorf("Result of actor (%v) does not match expected (%v)", actor, expectedActor)
	}

	if UrlsList.Size() != UrlsSize {
		t.Errorf("Expected urls list size: %d, Got: %d", UrlsSize, UrlsList.Size())
	}
	url := UrlsList.GetHeadDataAndMoveNext()
	if url != expectedUrl {
		t.Errorf("Result of url (%v) does not match expected (%v)", url, expectedUrl)
	}
}

func expectedEventTypes() map[string]int {
	return map[string]int{
		"push": 2,
	}
}

// Mock implementation of GetEvents
type MockEventGetter struct{}

func (m *MockEventGetter) GetEvents() []models.Event {
	// Mock implementation to return a predefined list of events
	events := []models.Event{
		{
			ID:   "1",
			Type: "push",
			Actor: models.Actor{
				ID:           1,
				Login:        "mock1",
				DisplayLogin: "mock1",
				GravatarID:   "gravatar_id",
				URL:          "url1",
				AvatarURL:    "avatar_url"},
			Repo: models.Repo{
				ID:         1,
				Name:       "repo_name1",
				URL:        "url1",
				StarsCount: 5},
			Payload: models.Payload{
				Commits: []models.Commit{
					{
						Author: models.Author{
							Email: "mockuser1@example.com",
						},
					},
				},
			},
		},
		{
			ID:   "2",
			Type: "push",
			Actor: models.Actor{
				ID:           2,
				Login:        expectedActor,
				DisplayLogin: "mock2",
				GravatarID:   "gravatar_id",
				URL:          "url2",
				AvatarURL:    "avatar_url"},
			Repo: models.Repo{
				ID:         2,
				Name:       "repo_name2",
				URL:        expectedUrl,
				StarsCount: 3},
			Payload: models.Payload{
				Commits: []models.Commit{
					{
						Author: models.Author{
							Email: "mockuser2@example.com",
						},
					},
				},
			},
		},
	}

	return events
}

//TODO::
//Due to lack of time, I wrote only one test, but I would have added tests that verify the returned JSON data.
//I would have checked for a different order in the list and verified the repo stars method.
