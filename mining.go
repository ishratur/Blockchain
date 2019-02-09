package blockchain

import (
	"../work_queue"
	
	
	
	
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	Start uint64
	End uint64
	Blk Block
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO
	
	if chunks>=(end-start) {
		chunks=8
		
	}
	queue := work_queue.Create(uint(workers), uint(chunks))

	_range := (end-start+1)/chunks
	var worker work_queue.Worker
	mr := MiningResult{}
	mw := miningWorker{}
	mw.Blk = blk
	for i := start; i <= end; i+=_range {
		mw.Start = i
		mw.End = i+_range
		if i+_range > end {
			mw.End = end  
		}
		worker = mw
		queue.Enqueue(worker)

		result := (<- queue.Results).(MiningResult)
		if result.Found {
			mr = result
			queue.Shutdown()
			break
		}
	}

	return mr
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

func (mw miningWorker) Run() interface{} {
	for i := mw.Start; i <= mw.End; i++ {
			mw.Blk.SetProof(i)
		if mw.Blk.ValidHash() {
			return MiningResult{i, true}
		}
	}
	return MiningResult{0, false}
}

