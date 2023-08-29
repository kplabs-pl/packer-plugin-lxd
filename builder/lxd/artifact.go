// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lxd

import (
	"fmt"
	"os"
)

type Artifact struct {
	id string
	f   string
	// StateData should store data such as GeneratedData
	// to be shared with post-processors
	StateData map[string]interface{}
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return []string{ a.f }
}

func (a *Artifact) Id() string {
	return a.id
}

func (a *Artifact) String() string {
	return fmt.Sprintf("image: %s in file %s", a.id, a.f)
}

func (a *Artifact) State(name string) interface{} {
	return a.StateData[name]
}

func (a *Artifact) Destroy() error {
	return os.RemoveAll(a.f)
}
