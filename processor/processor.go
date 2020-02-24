package processor

import "mdcdntools/common"

type Processor interface {
	Execute(config common.ArgsConfig) (bool, error)
}
