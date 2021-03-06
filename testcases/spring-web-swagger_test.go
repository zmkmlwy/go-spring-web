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

package testcases_test

import (
	"container/list"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/go-spring/go-spring-web/spring-echo"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-swagger"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring-web/testcases"
)

func TestSwagger(t *testing.T) {

	l := list.New()
	f2 := testcases.NewNumberFilter(2, l)
	f5 := testcases.NewNumberFilter(5, l)
	f7 := testcases.NewNumberFilter(7, l)

	get := func(ctx SpringWeb.WebContext) {
		ctx.LogInfo("invoke get()")
		ctx.String(http.StatusOK, "1")
	}

	server := SpringWeb.NewWebServer()

	// 添加第一个 web 容器
	{
		c1 := SpringGin.NewContainer()
		server.AddWebContainer(c1)
		c1.SetPort(8080)

		c1.GET(SpringSwagger.GET("/get", get, f2, f5, f7).Doc("get doc").Build())
	}

	// 添加第二个 web 容器
	{
		c2 := SpringEcho.NewContainer()
		server.AddWebContainer(c2)
		c2.SetPort(9090)

		c2.GET(SpringSwagger.GET("/get", get, f2, f5, f7).Doc("get doc").Build())
	}

	// 启动 web 服务器
	server.Start()

	time.Sleep(time.Millisecond * 100)
	fmt.Println()

	resp, _ := http.Get("http://127.0.0.1:8080/get?key=a")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("code:", resp.StatusCode, "||", "resp:", string(body))
	fmt.Println()

	resp, _ = http.Get("http://127.0.0.1:9090/get?key=a")
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("code:", resp.StatusCode, "||", "resp:", string(body))
	fmt.Println()

	server.Stop(context.TODO())

	time.Sleep(time.Millisecond * 50)
}
