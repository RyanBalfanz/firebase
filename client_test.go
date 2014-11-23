package firebase

import (
	"testing"
)

const FIREBASE_URL string = "https://redirect-service.firebaseio.com/"

func TestClient(t *testing.T) {
	client := NewClient(FIREBASE_URL)
	if client.BaseUrl() != FIREBASE_URL {
		t.Fail()
	}
}

func TestClientGetEntity(t *testing.T) {
	client := NewClient(FIREBASE_URL)
	// var entity = NewFirebaseEntity()
	entity, err := client.GetEntity("example")
	if err != nil {
		t.Fail()
	}

	if entity.Data["url"] != "http://example.com" {
		t.Errorf("%s", entity.Data)
	}
}
