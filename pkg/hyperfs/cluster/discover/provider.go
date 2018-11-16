// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package discover

// Provider interface
type Provider interface {
	Name() string
	Addrs() ([]string, error)
	Close() error
}
