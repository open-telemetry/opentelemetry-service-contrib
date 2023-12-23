// Code generated by mdatagen. DO NOT EDIT.

package main

import (
	"os"
	"path/filepath"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/configschema"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"
)

//go:generate go run ./component_config_gen.go

func main() {
	f := apachereceiver.NewFactory()
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := configschema.GenerateConfigDoc(filepath.Join(path, "..", ".."), f); err != nil {
		panic(err)
	}
	if err := configschema.GenerateMetadatada(f, filepath.Join(path, "..", ".."), filepath.Join(path, "..", "..")); err != nil {
		panic(err)
	}
}
