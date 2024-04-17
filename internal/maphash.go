package internal

// maphash implementation from github.com/dolthub/maphash

// Copyright 2022 Dolthub, Inc.
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

import (
	"math/rand"
	"unsafe"
)

// Hasher hashes values of type K.
// Uses runtime AES-based hashing.
type Hasher[K comparable] struct {
	hash hashfn
	seed uintptr
}

// NewHasher creates a new Hasher[K] with a random seed.
func NewHasher[K comparable]() Hasher[K] {
	return Hasher[K]{
		hash: getRuntimeHasher[K](),
		seed: newHashSeed(),
	}
}

// NewSeed returns a copy of |h| with a new hash seed.
func NewSeed[K comparable](h Hasher[K]) Hasher[K] {
	return Hasher[K]{
		hash: h.hash,
		seed: newHashSeed(),
	}
}

// Hash hashes |key|.
func (h Hasher[K]) Hash(key K) uint64 {
	return uint64(h.Hash2(key))
}

// Hash2 hashes |key| as more flexible uintptr.
func (h Hasher[K]) Hash2(key K) uintptr {
	// promise to the compiler that pointer
	// |p| does not escape the stack.
	p := noescape(unsafe.Pointer(&key))
	return h.hash(p, h.seed)
}

type hashfn func(unsafe.Pointer, uintptr) uintptr

func getRuntimeHasher[K comparable]() (h hashfn) {
	a := any(make(map[K]struct{}))
	i := (*mapiface)(unsafe.Pointer(&a))
	h = i.typ.hasher
	return
}

func newHashSeed() uintptr {
	return uintptr(rand.Int())
}

// noescape hides a pointer from escape analysis. It is the identity function
// but escape analysis doesn't think the output depends on the input.
// noescape is inlined and currently compiles down to zero instructions.
// USE CAREFULLY!
// This was copied from the runtime (via pkg "strings"); see issues 23382 and 7921.
//
//go:nosplit
//go:nocheckptr
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

type mapiface struct {
	typ *maptype
	val *hmap
}

// go/src/runtime/type.go
type maptype struct {
	typ    _type
	key    *_type
	elem   *_type
	bucket *_type
	// function for hashing keys (ptr to key, seed) -> hash
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8
	elemsize   uint8
	bucketsize uint16
	flags      uint32
}

// go/src/runtime/map.go
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	// hash seed
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	// true type is *mapextra
	// but we don't need this data
	extra unsafe.Pointer
}

// go/src/runtime/type.go
type tflag uint8
type nameOff int32
type typeOff int32

// go/src/runtime/type.go
type _type struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}
