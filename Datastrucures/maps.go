package Datastrucures

import "unsafe"

v = m[k]
// compiles to something similiar to the following
v = runtime.mapaccess(mt, m, k)
// which has the following signature

type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      *mapextra
}

// map bucket
type bmap struct {
	tophash [8]uint8
	keys    [8]keyType
	values  [8]valueType
}


func mapaccess(t *maptype, m *hmap, k unsafe.Pointer) unsafe.Pointer
