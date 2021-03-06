/*
 * Aminer (C) 2014, 2015 Minio, Inc.
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

package main

import "github.com/minio/cli"

var commands = []cli.Command{
	findCmd,
	populateCmd,
}

var findCmd = cli.Command{
	Name:   "find",
	Usage:  "find all documents for a map",
	Action: runFindCmd,
}

var populateCmd = cli.Command{
	Name:   "populate",
	Usage:  "populate your mongo instance with new data",
	Action: runPopulateCmd,
}

var flags = []cli.Flag{
	cli.StringFlag{
		Name:  "server",
		Value: "localhost",
		Usage: "IP/HOSTNAME of your mongodb instance",
	},
	cli.StringFlag{
		Name:  "filter",
		Value: "50.204.118.154,10.134.253.170",
		Usage: "Provide command separated list of ip's to be filtered",
	},
}
