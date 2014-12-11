package west

import (
	"encoding/json"
)

// Marshal marshals the object into bytes.
var Marshal = json.Marshal

// Unmarshal unmsrahsls from the bytes into the object.
var Unmarshal = json.Unmarshal
