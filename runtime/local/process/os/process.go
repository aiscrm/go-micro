// Package os runs processes locally
package os

import (
	"github.com/aiscrm/go-micro/v2/runtime/local/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
