// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
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
//

package timestamp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// TimestampResponseHandlerFunc turns a function with the right signature into a timestamp response handler
type TimestampResponseHandlerFunc func(TimestampResponseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TimestampResponseHandlerFunc) Handle(params TimestampResponseParams) middleware.Responder {
	return fn(params)
}

// TimestampResponseHandler interface for that can handle valid timestamp response params
type TimestampResponseHandler interface {
	Handle(TimestampResponseParams) middleware.Responder
}

// NewTimestampResponse creates a new http.Handler for the timestamp response operation
func NewTimestampResponse(ctx *middleware.Context, handler TimestampResponseHandler) *TimestampResponse {
	return &TimestampResponse{Context: ctx, Handler: handler}
}

/* TimestampResponse swagger:route POST /api/v1/tsr timestamp timestampResponse

Generates a timestamp response

*/
type TimestampResponse struct {
	Context *middleware.Context
	Handler TimestampResponseHandler
}

func (o *TimestampResponse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewTimestampResponseParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
