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

package SpringWeb

import (
	"net/http"
)

// Router 路由分组
type Router struct {
	mapping  WebMapping
	basePath string
	filters  []Filter
}

// NewRouter Router 的构造函数
func NewRouter(mapping WebMapping, basePath string, filters []Filter) *Router {
	return &Router{
		mapping:  mapping,
		basePath: basePath,
		filters:  filters,
	}
}

// Request 注册任意 HTTP 方法处理函数
func (r *Router) Request(method string, path string, fn Handler) *Mapper {
	return r.mapping.Request(method, r.basePath+path, fn, r.filters...)
}

// GET 注册 GET 方法处理函数
func (r *Router) GET(path string, fn Handler) *Mapper {
	return r.Request(http.MethodGet, path, fn)
}

// POST 注册 POST 方法处理函数
func (r *Router) POST(path string, fn Handler) *Mapper {
	return r.Request(http.MethodPost, path, fn)
}

// PATCH 注册 PATCH 方法处理函数
func (r *Router) PATCH(path string, fn Handler) *Mapper {
	return r.Request(http.MethodPatch, path, fn)
}

// PUT 注册 PUT 方法处理函数
func (r *Router) PUT(path string, fn Handler) *Mapper {
	return r.Request(http.MethodPut, path, fn)
}

// DELETE 注册 DELETE 方法处理函数
func (r *Router) DELETE(path string, fn Handler) *Mapper {
	return r.Request(http.MethodDelete, path, fn)
}

// HEAD 注册 HEAD 方法处理函数
func (r *Router) HEAD(path string, fn Handler) *Mapper {
	return r.Request(http.MethodHead, path, fn)
}

// OPTIONS 注册 OPTIONS 方法处理函数
func (r *Router) OPTIONS(path string, fn Handler) *Mapper {
	return r.Request(http.MethodOptions, path, fn)
}
