package dsg

import "testing"

func PrintSet (set * dsg.Set, name string) {
      var i int
      var labels []int = set.GetAllLabel ()

      t.Logf ("%s = { ", name)
      for i = 0; i < len(labels); i ++ {
            t.Logf ("%5d", labels[i])
      }
      t.Logf (" }\n")
}

func main () {
      var i int
      size := 21

      var s1 []bool = []bool{false,false,true,true,true,false,false,true,false,false,true,true,false,true,false,true,true,false,true,false,true}
      var s2 []bool = []bool{true,false,false,false,true,false,true,false,true,false,true,true,false,false,false,true,true,true,true,false,false}

      set1 := dsg.InitSet (size)
      set2 := dsg.InitSet (size)
      setu := dsg.InitSet (size)
      seti := dsg.InitSet (size)
      setc := dsg.InitSet (size)
      setd := dsg.InitSet (size)
      sete := dsg.InitSet (size)
      setf := dsg.InitSet (size)

      setf.Full ()

      for i = 0; i < size; i ++ {
            set1.SetLabel (i, s1[i]);
            set2.SetLabel (i, s2[i]);
      }

      setu.Union (set1, set2)
      seti.Inter (set1, set2)
      setc.Co (set1)
      setd.Diff (set1, set2)

      PrintSet (set1, "S1")
      PrintSet (set2, "S2")
      PrintSet (setu, "Union")
      PrintSet (seti, "Intersection")
      PrintSet (setc, "Complement")
      PrintSet (setd, "Difference")
      PrintSet (sete, "Empty")
      PrintSet (setf, "Full")
      set2.Empty ()
      PrintSet (set2, "AfterEmpty")
}
