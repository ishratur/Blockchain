package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
	"encoding/hex"
)

// TODO: some useful tests of Blocks
func TestCalcHash(t *testing.T) {
	b0 := Initial(2)
	b0.SetProof(242278)
	h0 := "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000"
	assert.Equal(t, hex.EncodeToString(b0.Hash), h0)

	b1 := b0.Next("message")
	b1.SetProof(75729)
	h1 := "02b09bde9ff60582ef21baa4bef87a95dfcd67efaf258e6df60463da0a940000"
	assert.Equal(t, hex.EncodeToString(b1.Hash), h1)
	fmt.Println()
}

func TestBlockChain(t *testing.T) {
	b0 := Initial(2)
	b0.SetProof(242278)
	b1 := b0.Next("this is an interesting message")
	b1.SetProof(41401)

	chain := Blockchain{}
	chain.Add(b0)
	chain.Add(b1)
	assert.Equal(t, chain.IsValid(), true)

	b2 := b1.Next("this is not interesting")
	b2.SetProof(195955)
	assert.Equal(t, chain.IsValid(), true)
	fmt.Println()

}

func TestInitialTwo(t *testing.T) {
	b0 := Initial(2)
	b0.Mine(1)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	assert.Equal(t, b0.ValidHash(), true)

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	assert.Equal(t, b1.ValidHash(), true)

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	assert.Equal(t, b2.ValidHash(), true)
	fmt.Println()

}

func TestInitialThree(t *testing.T) {
	
	b0 := Initial(3)
	b0.Mine(5)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	assert.Equal(t, b0.ValidHash(), true)
	fmt.Println("One down, 2 to go!")

	b1 := b0.Next("this is an interesting message")
	b1.Mine(5)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	assert.Equal(t, b1.ValidHash(), true)
	fmt.Println("Almost there!")

	b2 := b1.Next("this is not interesting")
	b2.Mine(5)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	assert.Equal(t, b2.ValidHash(), true)
	fmt.Println()


	
}



func TestInitialFour(t *testing.T){
	b0 := Initial(19)
	b0.SetProof(87745)
	b1 := b0.Next("hash example 1234")
	b1.SetProof(1407891)
	assert.Equal(t, b1.ValidHash(), true)
	fmt.Println()

	

}
func TestInitialFive(t *testing.T){
	b0 := Initial(19)
	b0.SetProof(87745)
	b1 := b0.Next("hash example 1234")
	b1.SetProof(346082)
	assert.Equal(t, b1.ValidHash(), false)
	fmt.Println()




}
func TestInitialSix(t *testing.T) {
	
	b0 := Initial(35)
	b0.Mine(5)
	fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
	assert.NotEqual(t, b0.ValidHash(), false)
	fmt.Println("One down, 2 to go!")

	b1 := b0.Next("this is an interesting message")
	b1.Mine(5)
	fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
	assert.NotEqual(t, b1.ValidHash(), false)
	fmt.Println("Almost there!")

	b2 := b1.Next("this is not interesting")
	b2.Mine(5)
	fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
	assert.NotEqual(t, b2.ValidHash(), false)
	fmt.Println()
	fmt.Println()
	

	fmt.Println("CONGRATS! ALL SIX TESTS PASSED!")
	
}