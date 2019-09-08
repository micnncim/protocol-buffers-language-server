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
	"sync"

	"github.com/go-language-server/uri"
	"go.uber.org/atomic"
)

var (
	sessionIndex = &atomic.Int64{}
	viewIndex    = &atomic.Int64{}
)

// Session represents a single connection from a client.
// This is the level at which things like open files are maintained on behalf
// of the client.
// A session may have many active views at any given time.
type Session interface {
	// AddView creates a new View, adds it to the Session and returns it.
	AddView(ctx context.Context, view View)

	// View returns a view with a matching name, if the session has one.
	View(name string) (View, bool)

	// ViewOf returns a view corresponding to the given URI.
	ViewOf(uri uri.URI) (View, bool)

	// Views returns the set of active views built by this session.
	Views() []View

	// Shutdown the session and all views it has created.
	Shutdown(ctx context.Context)

	// DidOpen is invoked each time a file is opened in the editor.
	DidOpen(ctx context.Context, uri uri.URI)

	// DidSave is invoked each time an open file is saved in the editor.
	DidSave(uri uri.URI)

	// DidClose is invoked each time an open file is closed in the editor.
	DidClose(uri uri.URI)

	// IsOpen can be called to check if the editor has a file currently open.
	IsOpen(uri uri.URI) bool

	SetOverlay(uri uri.URI, data []byte) (isFirstChange bool)

	GetOverlay(uri uri.URI) (*overlay, bool)
}

type session struct {
	id int64

	views   []View
	viewMap map[uri.URI]View
	viewMu  *sync.RWMutex

	overlayMu *sync.RWMutex
	overlays  map[uri.URI]*overlay

	openFiles   map[uri.URI]bool
	openFilesMu *sync.RWMutex
}

// overlay is an overlay for changed files.
type overlay struct {
	session   Session
	uri       uri.URI
	data      []byte
	hash      string
	unchanged bool
}

// NewSession returns Session.
func NewSession() Session {
	return &session{
		id:          sessionIndex.Add(1),
		viewMap:     make(map[uri.URI]View),
		viewMu:      &sync.RWMutex{},
		openFiles:   make(map[uri.URI]bool),
		openFilesMu: &sync.RWMutex{},
	}
}

var _ Session = (*session)(nil)

func (s *session) AddView(ctx context.Context, view View) {
	s.viewMu.Lock()

	s.views = append(s.views, view)
	// we always need to drop the view map
	s.viewMap = make(map[uri.URI]View)

	s.viewMu.Unlock()
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

func (s *session) ViewOf(uri uri.URI) (v View, ok bool) {
	s.viewMu.RLock()
	v, ok = s.viewMap[uri]
	s.viewMu.RUnlock()
	return v, ok
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

func (s *session) Shutdown(ctx context.Context) {
	s.viewMu.Lock()
	defer s.viewMu.Unlock()

	s.views = nil
	s.viewMap = nil
}

func (s *session) DidOpen(ctx context.Context, uri uri.URI) {
	s.openFilesMu.Lock()
	defer s.openFilesMu.Unlock()

	s.openFiles[uri] = true
}

// TODO: Implement.
func (s *session) DidSave(uri uri.URI) {
	panic("implement me")
}

func (s *session) DidClose(uri uri.URI) {
	s.openFilesMu.Lock()
	delete(s.openFiles, uri)
	s.openFilesMu.Unlock()
}

func (s *session) IsOpen(uri uri.URI) bool {
	s.openFilesMu.RLock()
	defer s.openFilesMu.RUnlock()

	open, ok := s.openFiles[uri]
	if !ok {
		return false
	}
	return open
}

func (s *session) SetOverlay(uri uri.URI, data []byte) (isFirstChange bool) {
	s.overlayMu.Lock()
	defer s.overlayMu.Unlock()

	if data == nil {
		delete(s.overlays, uri)
		return
	}

	o := s.overlays[uri]

	s.overlays[uri] = &overlay{
		session:   s,
		uri:       uri,
		data:      data,
		hash:      hashContent(data),
		unchanged: o == nil,
	}

	isFirstChange = o != nil && o.unchanged
	return
}

func (s *session) GetOverlay(uri uri.URI) (overlay *overlay, ok bool) {
	s.overlayMu.RLock()
	overlay, ok = s.overlays[uri]
	s.overlayMu.RUnlock()
	return
}

func hashContent(content []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(content))
}
