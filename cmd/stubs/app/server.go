/*
 * Copyright (c) 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"github.com/baidu/openless/cmd/stubs/options"
	"github.com/baidu/openless/pkg/stubs"
)

func Run(stubsOptions *options.StubsOptions) error {
	addr := fmt.Sprintf(":%d", stubsOptions.Port)

	handler := wrapRouter(stubsOptions)
	return fasthttp.ListenAndServe(addr, handler)
}

func wrapRouter(stubsOptions *options.StubsOptions) func(ctx *fasthttp.RequestCtx) {
	return stubs.InstallHandler(stubsOptions).HandleRequest
}