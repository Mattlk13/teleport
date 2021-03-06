/*
Copyright 2015 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shell

import (
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type ShellSuite struct {
}

var _ = check.Suite(&ShellSuite{})

func (s *ShellSuite) TestGetShell(c *check.C) {
	shell, err := GetLoginShell("root")
	c.Assert(err, check.IsNil)
	c.Assert(shell == "/bin/bash" || shell == "/bin/sh", check.Equals, true)

	shell, err = GetLoginShell("non-existent-user")
	c.Assert(err, check.IsNil)
	c.Assert(shell, check.Equals, DefaultShell)

	shell, err = GetLoginShell("nobody")
	c.Assert(err, check.IsNil)
	c.Assert(shell, check.Matches, ".*(nologin|false)")
}
