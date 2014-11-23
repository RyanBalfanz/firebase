package firebase

type firebaseEntity struct {
	Data map[string]interface{}
}

func NewFirebaseEntity() *firebaseEntity {
	// d := new(firebaseEntity)
	// d.Data = make(map[string]interface{})
	// return d
	return &firebaseEntity{make(map[string]interface{})}
}

func NewFirebaseEntityWithData(data map[string]interface{}) *firebaseEntity {
	return &firebaseEntity{Data: data}
}
