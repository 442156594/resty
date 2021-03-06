// +build go1.7

// Copyright (c) 2015-2016 Jeevanandam M (jeeva@myjeeva.com)
// 2016 Andrew Grigorev (https://github.com/ei-grad)
// All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"time"
)

// Request type is used to compose and send individual request from client
// go-resty is provide option override client level settings such as
//		Auth Token, Basic Auth credentials, Header, Query Param, Form Data, Error object
// and also you can add more options for that particular request
//
type Request struct {
	URL        string
	Method     string
	QueryParam url.Values
	FormData   url.Values
	Header     http.Header
	UserInfo   *User
	Token      string
	Body       interface{}
	Result     interface{}
	Error      interface{}
	Time       time.Time
	RawRequest *http.Request
	SRV        *SRVRecord

	client           *Client
	bodyBuf          *bytes.Buffer
	isMultiPart      bool
	isFormData       bool
	setContentLength bool
	isSaveResponse   bool
	outputFile       string
	proxyURL         *url.URL
	multipartFiles   []*File
	ctx              context.Context
}

// SetContext method sets the context.Context for current Request. It allows
// to interrupt the request execution if ctx.Done() channel is closed.
// See https://blog.golang.org/context article and the "context" package
// documentation.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

func (r *Request) addContextIfAvailable() {
	if r.ctx != nil {
		r.RawRequest = r.RawRequest.WithContext(r.ctx)
	}
}

func (r *Request) isContextCancelledIfAvailable() bool {
	if r.ctx != nil {
		if r.ctx.Err() != nil {
			return true
		}
	}
	return false
}
