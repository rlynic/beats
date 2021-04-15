// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by 'make imports' - DO NOT EDIT.

/*
Package include imports all input packages so that they register
their factories with the global registry. This package can be imported in the
main package to automatically register all of the standard supported inputs
modules.
*/
package include

import (
	//_ "github.com/elastic/beats/filebeat/input/docker"
	_ "github.com/elastic/beats/filebeat/input/log"
	//_ "github.com/elastic/beats/filebeat/input/redis"
	//_ "github.com/elastic/beats/filebeat/input/stdin"
	//_ "github.com/elastic/beats/filebeat/input/syslog"
	//_ "github.com/elastic/beats/filebeat/input/tcp"
	//_ "github.com/elastic/beats/filebeat/input/udp"
)
