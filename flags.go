package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// Flags is command flags
type Flags struct {
	Name         string
	Distribution string
	Path         string
	Version      string
	PathPattern  string
	EventType    string
	IsNewVersion bool
	Region       string
}

// Scan load floag from context to struct
func (f *Flags) Scan(ctx *cli.Context) *Flags {
	f.Name = ctx.String("name")
	f.Distribution = ctx.String("distribution")
	f.Path = ctx.String("path")
	f.Version = ctx.String("lambda-version")
	f.PathPattern = ctx.String("path-pattern")
	f.EventType = ctx.String("event-type")
	f.Region = ctx.String("region")
	f.IsNewVersion = ctx.Bool("publish-new-version")

	return f
}

// IsSetupFromSourceCode is upload lambda@edge from path and setup cloudfront
func (f *Flags) IsSetupFromSourceCode() bool {
	return f.Name != "" && f.Distribution != "" && f.Path != "" && f.PathPattern != "" && f.EventType != ""
}

// IsSetupFromVersion is setup cloud front from lambda@edge from version
func (f *Flags) IsSetupFromVersion() bool {
	return f.Path == "" && f.Name != "" && f.Distribution != "" && f.Version != "" && f.PathPattern != "" && f.EventType != ""
}

// IsUpdateFunctionCode is update function code
func (f *Flags) IsUpdateFunctionCode() bool {
	fmt.Printf("%s %s", f.Name, f.Path)
	return f.Name != "" && f.Path != ""
}
