package main

import (
	"github.com/aquasecurity/fanal/cache"
	fanalTypes "github.com/aquasecurity/fanal/types"
	"github.com/aquasecurity/trivy/pkg/scanner/local"
	trivyTypes "github.com/aquasecurity/trivy/pkg/types"
)

type (
	scannerSetup struct {
		fscache       cache.FSCache
		localScanner  local.Scanner
		scanOptions   trivyTypes.ScanOptions
		dockerOptions fanalTypes.DockerOption
	}

	optionSet struct {
		input string
		m     map[string]bool
	}
)
