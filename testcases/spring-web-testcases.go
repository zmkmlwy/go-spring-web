/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testcases

import (
	"container/list"
	"errors"
	"net/http"

	"github.com/go-spring/go-spring-parent/spring-error"
	"github.com/go-spring/go-spring-parent/spring-utils"
	"github.com/go-spring/go-spring-web/spring-web"
)

///////////////////// filter ////////////////////////

type NumberFilter struct {
	l *list.List
	n int
}

func NewNumberFilter(n int, l *list.List) *NumberFilter {
	return &NumberFilter{
		l: l,
		n: n,
	}
}

func (f *NumberFilter) Invoke(ctx SpringWeb.WebContext, chain *SpringWeb.FilterChain) {

	defer func() {
		ctx.LogInfo("::after", f.n)
		f.l.PushBack(f.n)
	}()

	ctx.LogInfo("::before", f.n)
	f.l.PushBack(f.n)

	chain.Next(ctx)
}

///////////////////// service ////////////////////////

type Service struct {
	store map[string]string
}

func NewService() *Service {
	return &Service{
		store: make(map[string]string),
	}
}

func (s *Service) Get(ctx SpringWeb.WebContext) {

	key := ctx.QueryParam("key")
	ctx.LogInfo("/get", "key=", key)

	val := s.store[key]
	ctx.LogInfo("/get", "val=", val)

	ctx.String(http.StatusOK, val)
}

func (s *Service) Set(ctx SpringWeb.WebContext) {

	var param struct {
		A string `form:"a" json:"a"`
	}

	if err := ctx.Bind(&param); err != nil {
		panic(err)
	}

	ctx.LogInfo("/set", "param="+SpringUtils.ToJson(param))

	s.store["a"] = param.A
}

func (s *Service) Panic(ctx SpringWeb.WebContext) {
	panic("this is a panic")
}

///////////////////// rpc service ////////////////////////

type RpcService struct{}

func (s *RpcService) OK(ctx SpringWeb.WebContext) interface{} {
	return "123"
}

func (s *RpcService) Err(ctx SpringWeb.WebContext) interface{} {
	panic("err")
}

func (s *RpcService) Panic(ctx SpringWeb.WebContext) interface{} {

	err := errors.New("panic")
	isPanic := ctx.QueryParam("panic") == "1"
	SpringError.ERROR.Panic(err).When(isPanic)

	return "ok"
}
