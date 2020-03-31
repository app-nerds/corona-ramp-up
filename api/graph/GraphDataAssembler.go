/*
 * Copyright (c) 2020. App Nerds LLC All Rights Reserved
 */

package graph

import (
	"github.com/app-nerds/corona-ramp-up/api/collate"
)

type GraphDataAssembler interface {
	Assemble(regions []collate.Region, startPoint collate.StartPoint) GraphDataCollection
}

type GraphDataAssemblerConfig struct {
	Collator collate.ICollator
}
