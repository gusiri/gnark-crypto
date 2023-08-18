// Copyright 2020 Consensys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package secp256k1

func processChunkG1Jacobian[B ibg1JacExtended](chunk uint64,
	chRes chan<- g1JacExtended,
	c uint64,
	points []G1Affine,
	digits []uint16,
	sem chan struct{}) {

	if sem != nil {
		// if we are limited, wait for a token in the semaphore
		<-sem
	}

	var buckets B
	for i := 0; i < len(buckets); i++ {
		buckets[i].setInfinity()
	}

	// for each scalars, get the digit corresponding to the chunk we're processing.
	for i, digit := range digits {
		if digit == 0 {
			continue
		}

		// if msbWindow bit is set, we need to subtract
		if digit&1 == 0 {
			// add
			buckets[(digit>>1)-1].addMixed(&points[i])
		} else {
			// sub
			buckets[(digit >> 1)].subMixed(&points[i])
		}
	}

	// reduce buckets into total
	// total =  bucket[0] + 2*bucket[1] + 3*bucket[2] ... + n*bucket[n-1]

	var runningSum, total g1JacExtended
	runningSum.setInfinity()
	total.setInfinity()
	for k := len(buckets) - 1; k >= 0; k-- {
		if !buckets[k].ZZ.IsZero() {
			runningSum.add(&buckets[k])
		}
		total.add(&runningSum)
	}

	if sem != nil {
		// release a token to the semaphore
		// before sending to chRes
		sem <- struct{}{}
	}

	chRes <- total
}

// we declare the buckets as fixed-size array types
// this allow us to allocate the buckets on the stack
type bucketg1JacExtendedC2 [2]g1JacExtended
type bucketg1JacExtendedC3 [4]g1JacExtended
type bucketg1JacExtendedC4 [8]g1JacExtended
type bucketg1JacExtendedC5 [16]g1JacExtended
type bucketg1JacExtendedC6 [32]g1JacExtended
type bucketg1JacExtendedC7 [64]g1JacExtended
type bucketg1JacExtendedC8 [128]g1JacExtended
type bucketg1JacExtendedC9 [256]g1JacExtended
type bucketg1JacExtendedC10 [512]g1JacExtended
type bucketg1JacExtendedC11 [1024]g1JacExtended
type bucketg1JacExtendedC12 [2048]g1JacExtended
type bucketg1JacExtendedC13 [4096]g1JacExtended
type bucketg1JacExtendedC14 [8192]g1JacExtended
type bucketg1JacExtendedC15 [16384]g1JacExtended

type ibg1JacExtended interface {
	bucketg1JacExtendedC2 |
		bucketg1JacExtendedC3 |
		bucketg1JacExtendedC4 |
		bucketg1JacExtendedC5 |
		bucketg1JacExtendedC6 |
		bucketg1JacExtendedC7 |
		bucketg1JacExtendedC8 |
		bucketg1JacExtendedC9 |
		bucketg1JacExtendedC10 |
		bucketg1JacExtendedC11 |
		bucketg1JacExtendedC12 |
		bucketg1JacExtendedC13 |
		bucketg1JacExtendedC14 |
		bucketg1JacExtendedC15
}
