package pretty

import (
	"encoding/json"
)

func Print(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
