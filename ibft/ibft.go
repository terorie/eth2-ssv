package ibft

import "github.com/bloxapp/ssv/ibft/types"

type IBFT struct {
	Instances   []*Instance
	persistence types.Persister
}
