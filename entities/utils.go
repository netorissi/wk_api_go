package entities

import (
	"bytes"
	"encoding/base32"
	"encoding/json"

	"github.com/google/uuid"
)

func MapToJson(objmap map[string]string) string {
	if b, err := json.Marshal(objmap); err != nil {
		return ""
	} else {
		return string(b)
	}
}

var encoding = base32.NewEncoding("12ndrfg8ejkmc44xot1999sza333h769")

// NewUUID is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewUUID() (string, error) {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	id, err := uuid.New().MarshalBinary()
	if err != nil {
		return "", err
	}
	encoder.Write(id)
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String(), nil
}
