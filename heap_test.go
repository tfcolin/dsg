package dsg

import "math"
import "testing"

func PrintHeap (hp * dsg.Heap, t * testing.T)  {
      var i, j, ii, m int
      var n int = hp.GetSize()

      if n == 0 {
            t.LogF ("Empty Heap\n")
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
                  t.Logf ("%6.2f", *hp.Get(ii))
                  for m = 0; m < step - 1; m ++ {t.Logf ("      ")}
                  ii ++
            }
            t.Logf ("\n")

            step /= 2
      }
}

func TestHeap (t * testing.T) {

      var i int
      var data []dsg.Value = []dsg.Value{2.2,5.1,8.3,1.3,4.9,7.2,3.1,6.2,9.0,10.1,11.6,12.7,0.1,-2.8,12.5,5.4};
      var n int = 16
      hp := dsg.InitFloatHeap ()

      for i = 0; i < n; i ++ {
            hp.Add (data[i])
            t.Logf ("After add %6.2f\n", data[i])
            PrintHeap(hp)
      }

      t.Logf ("Pri:\n")
      for i = 0; i < n; i ++ {
            value := hp.Pop().(float64)
            t.Logf ("Extract: %6.2f\n", value)
            PrintHeap (hp)
      }
      t.Logf ("\n")

      dsg.FloatHeapSort (data)

      t.Logf ("Data after sort =\n")
      for i = 0; i < n; i ++ {
            t.Logf ("%6.2f", data[i])
      }
      t.Logf ("\n")

}
