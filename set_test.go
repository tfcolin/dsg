package dsg

import "testing"
import "fmt"

func PrintSet (set * Set, name string) (logstr string) {
      var i int
      var labels []int = set.GetAllLabel ()

      logstr += fmt.Sprintf ("%s = { ", name)
      for i = 0; i < len(labels); i ++ {
            logstr += fmt.Sprintf ("%5d", labels[i])
      }
      logstr += fmt.Sprintf (" }\n")
      return 
}

func TestSet (t *testing.T) {
      var i int
      size := 21

      var s1 []bool = []bool{false,false,true,true,true,false,false,true,false,false,true,true,false,true,false,true,true,false,true,false,true}
      var s2 []bool = []bool{true,false,false,false,true,false,true,false,true,false,true,true,false,false,false,true,true,true,true,false,false}

      set1 := InitSet (size)
      set2 := InitSet (size)
      setu := InitSet (size)
      seti := InitSet (size)
      setc := InitSet (size)
      setd := InitSet (size)
      sete := InitSet (size)
      setf := InitSet (size)

      setf.Full ()

      for i = 0; i < size; i ++ {
            set1.SetLabel (i, s1[i]);
            set2.SetLabel (i, s2[i]);
      }

      setu.Union (set1, set2)
      seti.Inter (set1, set2)
      setc.Co (set1)
      setd.Diff (set1, set2)

      t.Log(PrintSet (set1, "S1"))
      t.Log(PrintSet (set2, "S2"))
      t.Log(PrintSet (setu, "Union"))
      t.Log(PrintSet (seti, "Intersection"))
      t.Log(PrintSet (setc, "Complement"))
      t.Log(PrintSet (setd, "Difference"))
      t.Log(PrintSet (sete, "Empty"))
      t.Log(PrintSet (setf, "Full"))

      set2.Empty ()

      t.Log(PrintSet (set2, "AfterEmpty"))
}
