// Copyright The Cryostat Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

const appNameEnv = "APP_NAME"
const operatorVersionEnv = "OPERATOR_VERSION"
const coreImageEnv = "CORE_IMG"
const datasourceImageEnv = "DATASOURCE_IMG"
const grafanaImageEnv = "GRAFANA_IMG"
const reportsImageEnv = "REPORTS_IMG"

// This program generates a const_generated.go file containing image tag
// constants for each container image deployed by the operator, along with
// other constants. These constants are populated using environment variables.
func main() {
	// Fill in image tags struct from the environment variables
	consts := struct {
		AppName            string
		OperatorVersion    string
		CoreImageTag       string
		DatasourceImageTag string
		GrafanaImageTag    string
		ReportsImageTag    string
	}{
		AppName:            getEnvVar(appNameEnv),
		OperatorVersion:    getEnvVar(operatorVersionEnv),
		CoreImageTag:       getEnvVar(coreImageEnv),
		DatasourceImageTag: getEnvVar(datasourceImageEnv),
		GrafanaImageTag:    getEnvVar(grafanaImageEnv),
		ReportsImageTag:    getEnvVar(reportsImageEnv),
	}

	// Create the source file to generate
	file, err := os.Create("const_generated.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = fileTemplate.Execute(file, consts)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvVar(name string) string {
	val, pres := os.LookupEnv(name)
	if !pres {
		log.Fatalf("Environment variable \"%s\" is not defined", name)
	}
	return val
}

var fileTemplate = template.Must(template.New("").Parse(`// Code generated by const_generator.go; DO NOT EDIT.
package controllers

// User facing name of the operand application
const AppName = "{{ .AppName }}"

// Version of the Cryostat Operator
const OperatorVersion = "{{ .OperatorVersion }}"

// Default image tag for the core application image
const DefaultCoreImageTag = "{{ .CoreImageTag }}"

// Default image tag for the JFR datasource image
const DefaultDatasourceImageTag = "{{ .DatasourceImageTag }}"

// Default image tag for the Grafana dashboard image
const DefaultGrafanaImageTag = "{{ .GrafanaImageTag }}"

// Default image tag for the Grafana dashboard image
const DefaultReportsImageTag = "{{ .ReportsImageTag }}"
`))
