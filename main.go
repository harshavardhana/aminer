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

import (
	"net/http"
	"time"

	"github.com/minio/cli"
	"gopkg.in/mgo.v2"
)

// LogMessage is a serializable json log message
type LogMessage struct {
	StartTime     time.Time
	Duration      time.Duration
	StatusMessage string // human readable http status message
	ContentLength string // human readable content length

	// HTTP detailed message
	HTTP struct {
		ResponseHeaders http.Header
		Request         *http.Request
	}
}

var db *mgo.Collection

func connectToMongo(c *cli.Context) *mgo.Session {
	session, err := mgo.Dial(c.GlobalString("server"))
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	// make this configurable
	db = session.DB("test").C("downloads")
	return session
}

func main() {
	app := cli.NewApp()
	app.Usage = "A miner for your minio access logs"
	app.Version = "0.0.1"
	app.Commands = commands
	app.Flags = flags
	app.Author = "Minio.io"

	app.RunAndExitOnError()
}
