package dsg

import "io"
import "fmt"

type Set struct {
      label []byte
      size int
}

func InitSet (size int) * Set {
      var bsize int = (size - 1) / 8 + 1;
      var res Set = Set{size: size, label: make ([]byte, bsize)}
      return &res;
}

func (s1 * Set) Empty () {
      var i int
      for i = 0; i < len(s1.label); i ++ {
            s1.label[i] = 0
      }
}

func (s1 * Set) Full () {
      s1.Empty ()
      s1.Co (s1)
}

func (sres * Set) Union (s1 * Set, s2 * Set) {
      var i int
      for i = 0; i < len(s1.label); i ++ {
            sres.label[i] = s1.label[i] | s2.label[i]
      }
}

func (sres * Set) Inter (s1 * Set, s2 * Set) {
      var i int
      for i = 0; i < len(s1.label); i ++ {
            sres.label[i] = s1.label[i] & s2.label[i]
      }
}

func (sres * Set) Diff (s1 * Set, s2 * Set) {
      var i int
      for i = 0; i < len(s1.label); i ++ {
            sres.label[i] = s1.label[i] &^ s2.label[i]
      }
}

func (sres * Set) Co (s1 * Set) {
      var i int
      for i = 0; i < len(s1.label); i ++ {
            sres.label[i] = ^s1.label[i]
      }
}

func (set * Set) GetLabel (ind int) bool {
      i := ind / 8;
      ii := ind % 8;

      var res byte = (set.label[i] & (1 << ii))
      if res != 0 {
            return true
      } else {
            return false
      }
}

func (set * Set) SetLabel (ind int, label bool) {
      i := ind / 8
      ii := ind % 8
      var l byte

      if label {
            l = 1
      } else {
            l = 0
      }

      set.label[i] &^= 1 << ii
      set.label[i] |= l << ii;
}

func (set * Set) GetSize () int {
      return set.size
}

func (set * Set) GetNLabel () int {
      var i int
      var res int = 0
      var size int = set.size

      for i = 0; i < size; i ++ {
            if (set.GetLabel (i)) {
                  res ++
            }
      }

      return res
}

func (set * Set) GetAllLabel () (labels []int) {
      var i int
      var size int = set.size

      labels = make ([]int, 0, size)

      for i = 0; i < size; i ++ {
            if (set.GetLabel (i)) {
                  labels = append (labels, i)
            }
      }

      return
}

func (set * Set) IsEmpty () bool {
      return (set.GetNLabel() == 0)
}

func (oset * Set) CopyFrom (sset * Set) (is_err bool) {
      if oset.size != sset.size {return true}
      copy (oset.label, sset.label)
      return false
}

func (set * Set) Save (fsav io.Writer) {
      fmt.Fprintf (fsav, "%d ", set.size)
      for _, i := range set.label {
            fmt.Fprintf (fsav, "%v ", i)
      }
      fmt.Fprintf (fsav, "\n")
}

func (set * Set) Load (fload io.Reader) {
      fmt.Fscan (fload, &(set.size))
      set.label = make ([]byte, set.size)
      for i := 0; i < set.size; i ++ {
            fmt.Fscan (fload, &(set.label[i]))
      }
}
