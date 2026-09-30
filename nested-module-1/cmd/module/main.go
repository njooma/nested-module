package main

import (
	"nestedmodule1"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
	generic "go.viam.com/rdk/components/generic"
)

func main() {
	// ModularMain can take multiple APIModel arguments, if your module implements multiple models.
	module.ModularMain(resource.APIModel{ generic.API, nestedmodule1.NestedGenericComponent})
}
