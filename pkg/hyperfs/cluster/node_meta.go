// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

// NodeType type
type NodeType string

// NodeType enums
const (
	NodeTypeAPI     NodeType = "api"
	NodeTypeIndex   NodeType = "index"
	NodeTypeStorage NodeType = "storage"
)

// NodeMeta struct
type NodeMeta struct {
	Cluster    string
	Region     string
	Type       NodeType
	Properties map[string]interface{}
}
