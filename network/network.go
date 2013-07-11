package network

type Network struct {
    NetworkID byte
}

const ( 
    MainID = iota
    TestID = iota
)

var MainNetwork, TestNetwork *Network

func NewNetwork(network_id byte) (*Network) {
    return &Network {NetworkID: network_id}
}

func init() {
    MainNetwork = NewNetwork(MainID)
    TestNetwork = NewNetwork(TestID)
}
