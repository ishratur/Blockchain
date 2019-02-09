package blockchain

import (
	"bytes"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	// TODO
	zerobytes := make([]byte, 32)
	
	hash := chain.Chain[0].Hash

	//checking if : blk.Hash = blk.CalcHash()
	if !bytes.Equal(hash, chain.Chain[0].CalcHash()) {
		return false
	}
	//checking if the first block has a Valid hash
	if !chain.Chain[0].ValidHash() {
		return false
	}
	// checking if the prehvoius hash of the first block is zero byte
	if !bytes.Equal(chain.Chain[0].PrevHash, zerobytes) {
		return false
	}
	//checking if the the generation of the first block is zero
	if chain.Chain[0].Generation != 0 {
		return false
		
	}

	for i := 1; i < len(chain.Chain); i++ {
		if chain.Chain[i].Difficulty != chain.Chain[0].Difficulty {
			return false
		}

		if chain.Chain[i].Generation != chain.Chain[0].Generation + uint64(i) {
			return false
		}

		if !bytes.Equal(chain.Chain[i].PrevHash, hash) {
			return false
		}
		//update the hash values
		hash = chain.Chain[i].Hash

		if !bytes.Equal(hash, chain.Chain[i].CalcHash()) {
			return false
		}

		if !chain.Chain[i].ValidHash() {
			return false
		}
	}
	return true
}
