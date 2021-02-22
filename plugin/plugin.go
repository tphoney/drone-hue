// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	// TODO replace or remove
	Param1 string `envconfig:"PLUGIN_PARAM1"`
	Param2 string `envconfig:"PLUGIN_PARAM2"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// write code here
	logrus.Infoln("starting ...")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	body := strings.NewReader(
		`{"alert":"lselect"}`)
	request, err := http.NewRequest(
		http.MethodPut,
		"https://192.168.0.115/api/zpSMgrRfIA-JC8BQQrqGquobI-SsT0v7hsm5gV7R/groups/5/action",
		body,
	)
	if err != nil {
		logrus.Errorf("something went wrong building the request %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("something went wrong making the request %v", err)
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("something went wrong reading the response %v", err)
	}
	logrus.Infof("response %v", string(respBody))

	return nil
}
