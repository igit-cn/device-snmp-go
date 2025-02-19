// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2021 Canonical Ltd
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * SPDX-License-Identifier: Apache-2.0'
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"

	hooks "github.com/canonical/edgex-snap-hooks/v2"
)

var cli *hooks.CtlCli = hooks.NewSnapCtl()

// installProfiles copies the profile configuration.toml files from $SNAP to $SNAP_DATA.
func installConfig() error {
	var err error

	path := "/config/device-snmp/res/configuration.toml"
	destFile := hooks.SnapData + path
	srcFile := hooks.Snap + path

	if err = os.MkdirAll(filepath.Dir(destFile), 0755); err != nil {
		return err
	}

	if err = hooks.CopyFile(srcFile, destFile); err != nil {
		return err
	}

	return nil
}

func installDevices() error {
	var err error

	path := "/config/device-snmp/res/devices/device.snmp.trendnet.TPE082WS.toml"
	destFile := hooks.SnapData + path
	srcFile := hooks.Snap + path

	if err = os.MkdirAll(filepath.Dir(destFile), 0755); err != nil {
		return err
	}

	if err = hooks.CopyFile(srcFile, destFile); err != nil {
		return err
	}

	return nil
}

func installDevProfiles() error {
	var err error

	profs := [...]string{"patlite", "switch.dell.N1108P-ON", "trendnet.TPE082WS"}

	for _, v := range profs {
		path := fmt.Sprintf("/config/device-snmp/res/profiles/device.snmp.%s.yaml", v)
		destFile := hooks.SnapData + path
		srcFile := hooks.Snap + path

		if err := os.MkdirAll(filepath.Dir(destFile), 0755); err != nil {
			return err
		}

		if err = hooks.CopyFile(srcFile, destFile); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var err error

	if err = hooks.Init(false, "edgex-device-snmp"); err != nil {
		fmt.Println(fmt.Sprintf("edgex-device-snmp::install: initialization failure: %v", err))
		os.Exit(1)
	}

	err = installConfig()
	if err != nil {
		hooks.Error(fmt.Sprintf("edgex-device-snmp:install: %v", err))
		os.Exit(1)
	}

	err = installDevices()
	if err != nil {
		hooks.Error(fmt.Sprintf("edgex-device-snmp:install: %v", err))
		os.Exit(1)
	}

	err = installDevProfiles()
	if err != nil {
		hooks.Error(fmt.Sprintf("edgex-device-snmp:install: %v", err))
		os.Exit(1)
	}

}
