// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Settings for the plugin.
type Settings struct {
	// Fill in the data structure with appropriate values
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	// Validation of the settings.
	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	// Implementation of the plugin.
	fmt.Println("starting ...")
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
		fmt.Printf("something went wrong building the request %v", err)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("something went wrong making the request %v", err)
	}
	respBody, _ := ioutil.ReadAll(response.Body)

	fmt.Printf("response %v", string(respBody))

	return nil
}
