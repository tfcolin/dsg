package dsg

import "fmt"
import "testing"

func int_ktoi (key dsg.Key) (ind []int) {
      k := key.(int)
      ind = make ([]int, 3)
      ind[0] = k / 1000
      ind[1] = (k % 1000) / 100
      ind[2] = (k % 100) / 10

      return ind
}

func TestHash (t * testing.T) {
      var i int
      var logstr string

      n := []int{8, 10, 10}
      hash := dsg.InitTreeHash (n, int_ktoi, dsg.IntCompFunc)

      var key int
      var value dsg.Value

      for i = 0; i < 8000; i += 3 {
            key = i
            value = fmt.Sprintf ("%d_h", i)
            hash.AddData (key, value)
      }
      for i = 0; i < 8000; i += 200 {
            hash.RemoveData (i)
      }

      for i = 8000; i >= 0; i -= 50 {
            v := hash.Search (i)
            logstr += fmt.Sprintf ("Ind = %6d  ", i)
            if v == nil {
                  logstr += fmt.Sprintf ("Value = nil")
            } else {
                  logstr += fmt.Sprintf ("Value = %s", v.(string))
            }
            logstr += fmt.Sprintf ("\n")
      }

      hash.Clear()

      for i = 0; i < 5000; i += 50 {
            key = i
            value = fmt.Sprintf ("%d_hh", i)
            hash.AddData (key, value)
      }

      for i = 8000; i >= 0; i -= 50 {
            v := hash.Search (i)
            logstr += fmt.Sprintf ("Ind = %6d  ", i)
            if v == nil {
                  logstr += fmt.Sprintf ("Value = nil")
            } else {
                  logstr += fmt.Sprintf ("Value = %s", v.(string))
            }
            logstr += fmt.Sprintf ("\n")
      }

      t.Log (logstr)
}
