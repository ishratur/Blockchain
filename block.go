package blockchain

import (
	"encoding/hex"
	"crypto/sha256"
	//"strings"
	"fmt"

)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	// TODO
	var block Block
	block.Generation=0
	block.Difficulty=difficulty
	block.Data=""
	block.PrevHash= make([]byte, 32)
	//block.Proof=56231
	return block
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	var block Block
	block.Generation= prev_block.Generation+1
	block.Difficulty=prev_block.Difficulty
	block.Data=data
	block.PrevHash=prev_block.Hash
	return block
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	// TODO
	
	hash_string := hex.EncodeToString(blk.PrevHash) + ":"
	hash_string += fmt.Sprint(blk.Generation) + ":"
	hash_string += fmt.Sprint(blk.Difficulty) + ":"
	hash_string += fmt.Sprint(blk.Data) + ":"
	hash_string += fmt.Sprint(blk.Proof)
	new_hash := sha256.Sum256([]byte(hash_string))
	hash_return := new_hash[:]
	return hash_return


}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	// TODO
	//blk.Hash=blk.CalcHash()
	nBytes:= int(blk.Difficulty/8)
	nBits:=uint(blk.Difficulty%8)
	var length = len(blk.Hash)
	var a bool=false
	var zerobyte byte = 0
	//fmt.Println(blk.Hash)

	for i := 0; i < nBytes; i++ {
		if blk.Hash[length-i-1]==zerobyte {
			a=true
			
		}			
		
	}

	/*fmt.Println((1<<nBits))
	fmt.Println(blk.Hash[length-(nBytes+1)])*/
	

	if blk.Hash[length-(nBytes+1)]%(1<<nBits)==0 {
			a=true
			
	}else{
			a=false
		}


	return a


}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
	//bit form of the hash
	//a:=fmt.Sprintf("%08b",blk.Hash)
	//fmt.Println(a)
}

