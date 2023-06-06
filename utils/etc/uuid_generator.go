package etc

import uid "github.com/google/uuid"

func GenerateRandomUUID() string {
	return uid.NewString()
}
