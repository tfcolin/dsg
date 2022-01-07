package dsg

import "testing"
import "fmt"

func TestList (t * testing.T) {
      var i int
      var logstr string

      // for llist
      ll := InitLList (10)

      u1 := ll.Pop()
      u2 := ll.PopFirst()
      u3 := ll.GetFirst()
      u4 := ll.GetLast()

      if u1 == nil && u2 == nil && u3 == nil && u4 == nil {
            logstr += fmt.Sprintf ("All Nils.\n")
      }

      for i = 20; i > 0; i -= 2 {
            ll.Push (i)
      }

      for i = 0; i < ll.GetN(); i ++ {
            logstr += fmt.Sprintf("%6d", *ll.Get(i))
      }
      logstr += fmt.Sprintf ("\n")

      v1 := ll.Pop()
      v2 := ll.Pop()
      v3 := ll.PopFirst()
      v4 := *ll.GetFirst()
      v5 := *ll.GetLast()

      for i = 0; i < ll.GetN(); i ++ {
            logstr += fmt.Sprintf("%6d", *ll.Get(i))
      }
      logstr += fmt.Sprintf ("\n")

      logstr += fmt.Sprintf ("%6d%6d%6d%6d%6d\n", v1, v2, v3, v4, v5)

      ll.Flush()

      u1 = ll.Pop()
      u2 = ll.PopFirst()
      u3 = ll.GetFirst()
      u4 = ll.GetLast()

      if u1 == nil && u2 == nil && u3 == nil && u4 == nil {
            logstr += fmt.Sprintf ("All Nils.\n")
      }

      logstr += fmt.Sprintf ("Number of Elem. = %d\n", ll.GetN())

      t.Log (logstr)
}
