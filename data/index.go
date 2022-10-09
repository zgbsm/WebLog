package data

import (
	"github.com/google/cel-go/cel"
	"sync"
	"time"
)

type Info struct {
	Start     time.Time   `json:"start"`
	Rule      cel.Program `json:"rule"`
	Requested bool        `json:"requested"`
}

var Data map[string]Info
var DataLock sync.Mutex
