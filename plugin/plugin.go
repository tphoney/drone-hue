// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"crypto/tls"
	"errors"
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

	HubIP       string `envconfig:"PLUGIN_HUB_IP"`
	HubToken    string `envconfig:"PLUGIN_HUB_TOKEN"`
	TargetType  string `envconfig:"PLUGIN_TARGET_TYPE"`
	Target      string `envconfig:"PLUGIN_TARGET"`
	Payload     string `envconfig:"PLUGIN_PAYLOAD"`
	FailPayload string `envconfig:"PLUGIN_FAIL_PAYLOAD"`
}

// ValidateAndSetArgs checks parameters are set and gives default values for optional.
func ValidateAndSetArgs(args Args) (validatedArgs Args, err error) {
	if args.HubIP == "" || args.HubToken == "" {
		err = errors.New("hub_ip and hub_token must be set in settings.")
	}
	if args.TargetType == "" {
		args.TargetType = "groups"
	}
	if args.Target == "" {
		args.Target = "0"
	}
	if args.Payload == "" {
		args.Payload = `{"alert":"lselect"}`
	}
	if args.FailPayload == "" {
		args.FailPayload = `{"alert":"lselect"}`
	}
	validatedArgs = args
	return args, err
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	// write code here
	logrus.Debugln("Starting ...")
	args, err := ValidateAndSetArgs(args)
	if err != nil {
		logrus.Errorln("Wrong parameters passed.")
		return err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var body *strings.Reader
	if len(args.Failed.Steps) == 0 {
		body = strings.NewReader(
			args.Payload)
		logrus.Debugf("Previous steps have all passed. Using %s", args.Payload)
	} else {
		body = strings.NewReader(
			args.FailPayload)
		logrus.Debugf("Previous steps have a failure. Using %s", args.FailPayload)
	}
	request, err := http.NewRequest(
		http.MethodPut,
		"https://"+args.HubIP+"/api/"+args.HubToken+"/"+args.TargetType+"/"+args.Target+"/action",
		body,
	)
	if err != nil {
		logrus.Errorln("Something went wrong building the request.")
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("Something went wrong making the request.")
		return err
	}
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("Something went wrong reading the response.")
		return err
	}
	logrus.Debugf("response %s", string(respBody))

	return nil
}
