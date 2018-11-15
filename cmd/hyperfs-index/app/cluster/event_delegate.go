// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"github.com/hashicorp/memberlist"
)

// EventDelegate interface
type EventDelegate interface {
	Delegate() memberlist.EventDelegate
	Events() <-chan memberlist.NodeEvent
}

type eventDelegate struct {
	events chan memberlist.NodeEvent
}

// NewEventDelegate constructor
func NewEventDelegate(size int) EventDelegate {
	return &eventDelegate{
		events: make(chan memberlist.NodeEvent, size),
	}
}

func (e *eventDelegate) Delegate() memberlist.EventDelegate {
	return &memberlist.ChannelEventDelegate{
		Ch: e.events,
	}
}

func (e *eventDelegate) Events() <-chan memberlist.NodeEvent {
	return e.events
}
