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

package tchserver

import (
	"time"

	"github.com/yarpc/yarpc-go"
	"github.com/yarpc/yarpc-go/crossdock-go"
	"github.com/yarpc/yarpc-go/crossdock/client/random"
	"github.com/yarpc/yarpc-go/encoding/json"

	"golang.org/x/net/context"
)

func runJSON(t crossdock.T, rpc yarpc.RPC) {
	assert := crossdock.Assert(t)
	checks := crossdock.Checks(t)

	headers := yarpc.NewHeaders().With("hello", "json")
	token := random.String(5)

	resBody, resMeta, err := jsonCall(rpc, headers, token)
	if skipOnConnRefused(t, err) {
		return
	}
	if checks.NoError(err, "json: call failed") {
		assert.Equal(token, resBody, "body echoed")
		assert.Equal(headers, resMeta.Headers(), "headers echoed")
	}
}

type jsonEcho struct {
	Token string `json:"token"`
}

func jsonCall(rpc yarpc.RPC, headers yarpc.Headers, token string) (string, yarpc.CallResMeta, error) {
	client := json.New(rpc.Channel(serverName))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reqMeta := yarpc.NewReqMeta(ctx).Procedure("echo").Headers(headers)
	reqBody := &jsonEcho{Token: token}

	var resBody jsonEcho
	resMeta, err := client.Call(reqMeta, reqBody, &resBody)
	return resBody.Token, resMeta, err
}