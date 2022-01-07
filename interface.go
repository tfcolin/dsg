package dsg

type Value interface {}
type Key interface {}

type CompFunc func (v1 Value, v2 Value) int
type KeyToIndexFunc func (key Key) (ind []int)

