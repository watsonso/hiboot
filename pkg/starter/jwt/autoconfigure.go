// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

// Package jwt provides the hiboot starter for injectable jwt dependency
package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hidevopsio/hiboot/pkg/app"
	mw "github.com/iris-contrib/middleware/jwt"
)

type configuration struct {
	app.Configuration

	Properties Properties `mapstructure:"jwt"`
	middleware *Middleware
	token      Token
}

func init() {
	app.AutoConfiguration(newConfiguration)
}

func newConfiguration() *configuration {
	return &configuration{}
}

func (c *configuration) Middleware(jwtToken Token) *Middleware {
	return NewJwtMiddleware(mw.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//log.Debug(token)
			return jwtToken.VerifyKey(), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodRS256,
	})
}

// JwtToken
// TODO: token or jwtToken
func (c *configuration) Token() Token {
	t := new(jwtToken)
	t.Initialize(&c.Properties)
	return t
}
