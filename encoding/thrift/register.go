// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package thrift

import (
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport"

	"go.uber.org/thriftrw/protocol"
	"go.uber.org/thriftrw/wire"
	"golang.org/x/net/context"
)

// Register calls the Registry's Register method.
//
// This function exists for backwards compatibility only. It will be removed
// in a future version.
//
// Deprecated: Use the Registry's Register method directly.
func Register(r transport.Registry, rs []transport.Registrant) {
	r.Register(rs)
}

// Handler represents a Thrift request handler. It speaks in raw Thrift payloads.
//
// Users should use the server package generated by the code generator rather
// than using this directly.
type Handler interface {
	Handle(ctx context.Context, reqMeta yarpc.ReqMeta, body wire.Value) (Response, error)
}

// HandlerFunc is a convenience type alias for functions that act as Handlers.
type HandlerFunc func(context.Context, yarpc.ReqMeta, wire.Value) (Response, error)

// Handle forwards the request to the underlying function.
func (f HandlerFunc) Handle(ctx context.Context, reqMeta yarpc.ReqMeta, body wire.Value) (Response, error) {
	return f(ctx, reqMeta, body)
}

// Service is a generic Thrift service implementation.
type Service struct {
	// Name of the Thrift service. This is the name specified for the service
	// in the IDL.
	Name    string
	Methods map[string]Handler
}

// BuildRegistrants builds a list of Registrants from a Thrift service
// specification.
func BuildRegistrants(s Service, opts ...RegisterOption) []transport.Registrant {
	var rc registerConfig
	for _, opt := range opts {
		opt.applyRegisterOption(&rc)
	}

	proto := protocol.Binary
	if rc.Protocol != nil {
		proto = rc.Protocol
	}

	rs := make([]transport.Registrant, 0, len(s.Methods))
	for methodName, handler := range s.Methods {
		rs = append(rs, transport.Registrant{
			Procedure: procedureName(s.Name, methodName),
			Handler: thriftHandler{
				Handler:           handler,
				Protocol:          proto,
				DisableEnveloping: rc.DisableEnveloping,
			},
		})
	}
	return rs
}
