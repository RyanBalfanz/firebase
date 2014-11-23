package firebase

import (
	// "github.com/RyanBalfanz/firebase"
	"testing"
)

func TestCreateEntity(t *testing.T) {
	var e = NewFirebaseEntity()
	e.Data["foo"] = "bar"
	if e.Data["foo"] != "bar" {
		t.Fail()
	}
}

// func TestCreateEntityWithData(t *testing.T) {
// 	var e = firebase.NewFirebaseEntityWithData(make(map[string]interface{}))
// 	e.Data["foo"] = "bar"
// 	if e.Data["foo"] != "bar" {
// 		t.Fail()
// 	}

// 	var data = make(map[string]interface{})
// 	data["foo"] = "bar"
// 	var e2 = firebase.NewFirebaseEntityWithData(data)
// 	if e2.Data["foo"] != "bar" {
// 		t.Fail()
// 	}
// }
