package dsg

import "math"
import "testing"
import "fmt"

func PrintHeap (hp * Heap) (logstr string) {
      var i, j, ii, m int
      var n int = hp.GetSize()

      if n == 0 {
            logstr += fmt.Sprintf ("Empty Heap\n")
            return
      }

      var k int = int(math.Floor(math.Log2(float64(n))) + 1)
      if k > 5 {
            k = 5
      }

      var step int = int(math.Pow (2, float64(k - 1)));

      ii = 0
      for i = 0; i < k; i ++ {
            var ni int = int(math.Pow (2, float64(i)))
            for j = 0; j < ni; j ++ {
                  if (ii >= hp.GetSize ()) {
                        break;
                  }
                  logstr += fmt.Sprintf ("%6.2f", *hp.Get(ii))
                  for m = 0; m < step - 1; m ++ {logstr += fmt.Sprintf ("      ")}
                  ii ++
            }
            logstr += fmt.Sprintf ("\n")

            step /= 2
      }

      return 
}

func TestHeap (t * testing.T) {

      var i int
      var data []Value = []Value{2.2,5.1,8.3,1.3,4.9,7.2,3.1,6.2,9.0,10.1,11.6,12.7,0.1,-2.8,12.5,5.4};
      var n int = 16
      hp := InitFloatHeap ()

      for i = 0; i < n; i ++ {
            hp.Add (data[i])
            t.Logf ("After add %6.2f", data[i])
            t.Log(PrintHeap(hp))
      }

      t.Logf ("Pri:")
      for i = 0; i < n; i ++ {
            value := hp.Pop().(float64)
            t.Logf ("Extract: %6.2f", value)
            t.Log(PrintHeap (hp))
      }

      FloatHeapSort (data)

      t.Logf ("Data after sort =")
      var logstr string
      for i = 0; i < n; i ++ {
            logstr += fmt.Sprintf ("%6.2f", data[i])
      }
      t.Log (logstr)

}
