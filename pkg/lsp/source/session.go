// Copyright 2019 The Protocol Buffers Language Server Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"context"
	"crypto/sha1"
	"fmt"
	"strings"
	"sync"

	"github.com/go-language-server/uri"
	"go.uber.org/atomic"
)

var (
	sessionIndex = &atomic.Int64{}
	viewIndex    = &atomic.Int64{}
)

// Session represents a single connection from a client.
// A session just manages views and does not access files directly.
type Session interface {
	// View returns a view with a matching name, if the session has one.
	View(name string) (View, bool)

	// ViewOf returns a view corresponding to the given URI.
	ViewOf(uri uri.URI) View

	// Views returns the set of active views built by this session.
	Views() []View

	// AddView creates a new View, adds it to the Session and returns it.
	AddView(ctx context.Context, view View)

	// RemoveView removes a View with a matching name.
	RemoveView(ctx context.Context, view View) error

	// Shutdown the session and all views it has created.
	Shutdown(ctx context.Context)
}

type session struct {
	id int64

	views   []View
	viewMap map[uri.URI]View
	viewMu  *sync.RWMutex
}

var _ Session = (*session)(nil)

// NewSession returns Session.
func NewSession() Session {
	return &session{
		id:      sessionIndex.Add(1),
		viewMap: make(map[uri.URI]View),
		viewMu:  &sync.RWMutex{},
	}
}

func (s *session) View(name string) (View, bool) {
	s.viewMu.RLock()
	defer s.viewMu.RUnlock()

	for _, view := range s.views {
		if view.Name() == name {
			return view, true
		}
	}

	return nil, false
}

func (s *session) ViewOf(uri uri.URI) View {
	s.viewMu.RLock()
	// uri is folder and matches one of viewMap.
	v, ok := s.viewMap[uri]
	if ok {
		return v
	}
	s.viewMu.RUnlock()

	v = s.bestView(uri)
	s.viewMap[uri] = v

	return v
}

func (s *session) Views() []View {
	s.viewMu.RLock()
	defer s.viewMu.RUnlock()

	views := make([]View, 0, len(s.views))
	for _, view := range s.views {
		views = append(views, view)
	}

	return views
}

func (s *session) AddView(ctx context.Context, view View) {
	s.viewMu.Lock()

	s.views = append(s.views, view)
	// we always need to drop the view map
	s.viewMap = make(map[uri.URI]View)

	s.viewMu.Unlock()
}

func (s *session) RemoveView(ctx context.Context, view View) error {
	s.viewMu.Lock()
	defer s.viewMu.Unlock()
	// we always need to drop the view map
	s.viewMap = make(map[uri.URI]View)
	for i, v := range s.views {
		if v == view {
			s.views[i] = s.views[len(s.views)-1]
			s.views[len(s.views)-1] = nil
			s.views = s.views[:len(s.views)-1]
			return nil
		}
	}
	return fmt.Errorf("view %s for %v not found", view.Name(), view.Folder())
}

func (s *session) Shutdown(context.Context) {
	s.viewMu.Lock()
	defer s.viewMu.Unlock()

	s.views = nil
	s.viewMap = nil
}

// bestView finds the best view to associate a given URI with.
// viewMu must be held when calling this method.
func (s *session) bestView(uri uri.URI) View {
	// we need to find the best view for this file
	var longest View
	for _, view := range s.views {
		if longest != nil && len(longest.Folder()) > len(view.Folder()) {
			continue
		}
		if strings.HasPrefix(string(uri), string(view.Folder())) {
			longest = view
		}
	}
	if longest != nil {
		return longest
	}
	// TODO: are there any more heuristics we can use?
	return s.views[0]
}

func hashContent(content []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(content))
}
