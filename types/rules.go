package types

import (
	"github.com/subeenchung/kaleidoscope/types"
)

type Rule struct {

	Name string `yaml:"name"`
	Interval string `yaml:"interval"`
	Index string `yaml:"index"`
	SearchDSL types.DSL `yaml:"search"`

} 

