package dsg

func Float64CompFunc (v1 Value, v2 Value) int {
      fv1 := v1.(float64)
      fv2 := v2.(float64)
      if fv1 < fv2 {
            return -1
      } else if fv1 == fv2 {
            return 0
      } else {
            return 1
      }
}

func IntCompFunc (v1 Value, v2 Value) int {
      iv1 := v1.(int)
      iv2 := v2.(int)
      if iv1 < iv2 {
            return -1
      } else if iv1 == iv2 {
            return 0
      } else {
            return 1
      }
}

func InitIntHeap () * Heap {
      return InitHeap (IntCompFunc)
}

func InitFloatHeap () * Heap {
      return InitHeap (Float64CompFunc)
}

func FloatHeapSort (data []Value) {
      HeapSort (data, Float64CompFunc)
}

func IntHeapSort (data []Value) {
      HeapSort (data, IntCompFunc)
}
