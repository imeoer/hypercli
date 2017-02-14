package types

import (
	"time"

	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/strslice"
)

type Func struct {
  // Func name, required, unique, immutable, max length: 255, format: [a-z0-9]([-a-z0-9]*[a-z0-9])?
	Name string `json:"Name"`

	// Container size, optional, default: s4
	Size string `json:"Size,omitempty"`

	// Name of the container image, required, immutable
	Image string `json:"Image"`

	// Command to run when starting the container, optional, immutable
	Command strslice.StrSlice `json:"Command,omitempty"`

	// List of environment variable to set in the container, optional, format: ["VAR=value", ...]
	Env *[]string `json:"Env,omitempty"`

	// The response headers of http endpoint, optional, format: ["key=value", ...]
	Header *[]string `json:"Header,omitempty"`

	// The maximum number of concurrent container, optional, default: -1 (container quota), min: 1, max: container quota
	MaxConcurrency int `json:"MaxConcurrency,omitempty"`

	// The maximum number of func call which waiting for completed, optional, default: -1 (unlimit), min: 1, max: unlimit
	MaxLimit int `json:"MaxLimit,omitempty"`

	// The UUID of func
	UUID string `json:"UUID,omitempty"`

	// The created time
	Created time.Time `json:"Created,omitempty"`
}

type FuncListOptions struct {
	Filters filters.Args
}
