package dsg

import (
      "math"
)

/** LList 负责分配空间; */

/** 滑动数组表 (带索引链表) */
type LList struct
{
      buf []Value
      // 首节点索引
      first int
      // 最后一个节点下一个的索引
      last int
}

func InitLList (nbuf int) * LList {

      var ll LList

      ll.buf = make ([]Value, nbuf)
      ll.first = 0
      ll.last = 0

      return &ll
}

func (ll * LList) Push (value Value) int {
      nlast := real_mod (ll.last + 1, len(ll.buf))
      if nlast == ll.first {
            return -1
      }
      ll.buf[ll.last] = value
      ll.last = nlast
      return 0
}

func (ll * LList) Pop () (value Value) {

      if ll.last == ll.first {
            return nil
      }

      llast := real_mod (ll.last - 1, len(ll.buf))
      value = ll.buf[llast]
      ll.last = llast

      return
}

func (ll * LList) PopFirst () (value Value) {

      if ll.last == ll.first {
            return nil
      }

      value = ll.buf[ll.first]
      ll.first = real_mod (ll.first + 1, len(ll.buf))
      return
}

func (ll * LList) GetLast () (value * Value) {
      if ll.last == ll.first {
            return nil
      }

      llast := real_mod (ll.last - 1, len(ll.buf))
      return &(ll.buf[llast])
}

func (ll * LList) GetFirst () (value * Value) {
      if ll.last == ll.first {
            return nil
      }
      return &(ll.buf[ll.first])
}

func (ll * LList) Get (ind int) (value * Value) {
      if ind < 0 || ind >= ll.GetN () {
            return nil
      }
      rind := real_mod (ll.first + ind, len(ll.buf))
      return &(ll.buf [rind])
}

func (ll * LList) GetN () int {
      return real_mod (ll.last - ll.first, len(ll.buf))
}

func (ll * LList) Flush () {
      ll.last = ll.first
}

func real_mod (a, b int) int {
      if b < 0 {
            b = int(math.Abs(float64(b)))
      }
      if a < 0 {
            a += (- ((a + 1) / b) + 1) * b
      } else {
            a = a % b
      }

      return a
}

