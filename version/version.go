package version

var (
	// TMCore is the current version of Tendermint Core.
	// It's the Semantic Version of the software.
	// Must be a string because scripts like dist.sh read this file.
	TMCore  = "0.24.0"
	Version = TMCore

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

// ABCIVersion is the version of the ABCI library
const ABCIVersion = "0.14.0"

// Protocol is used for implementation agnostic versioning.
type Protocol int64

var (
	// P2PProtocol versions all p2p behaviour and msgs.
	P2PProtocol Protocol = 4

	// BlockProtocol versions all block data structures and processing.
	BlockProtocol Protocol = 7
)

// App includes the protocol and software version for the application.
// This information is included in ResponseInfo. The App.Protocol can be
// updated in ResponseEndBlock.
type App struct {
	Protocol Protocol
	Software string
}

// Consensus captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
type Consensus struct {
	Block Protocol `json:"block"`
	App   Protocol `json:"app"`
}

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
