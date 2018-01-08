// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Mon,  8 Jan 2018]

package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/eddiefisher/ponger/src/version"
)

func init() {
	logrus.Printf("commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
}

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		logrus.Warnln(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Warnln(err)
			conn.Close()
			break
		}
		err = handler(conn)
		if err != nil {
			logrus.Warnln(err)
			conn.Close()
			continue
		}
	}
}

func handler(conn net.Conn) (err error) {
	for i := 0; i < 3; i++ {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	}
	err = errors.New("end of for")
	return
}
