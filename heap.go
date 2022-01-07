package dsg

type Heap struct {
      cmf CompFunc
      data []Value
}

func InitHeap (cmf CompFunc) * Heap {
      const init_block = 1024

      var hp Heap
      hp.data = make ([]Value, 0, init_block)
      hp.cmf = cmf

      return &hp
}

func (hp * Heap) Get (ind int) * Value {
      return &(hp.data[ind])
}

/* return -1 if no child */
func (hp * Heap) GetLCInd (ind int) int {
      res := ind * 2 + 1;
      if res >= len(hp.data) {
            return -1
      } else {
            return res
      }
}

/* return -1 if no child */
func (hp * Heap) GetRCInd (ind int) int {
      res := ind * 2 + 2;
      if res >= len(hp.data) {
            return -1
      } else {
            return res
      }
}

/* return -1 if no par (root) */
func (hp * Heap) GetParInd (ind int) int {
      res := 0;
      if ind == 0 {
            res = -1
      } else {
            res = (ind - 1) / 2
      }
      return res;
}

func (hp * Heap) GetSize () int {
      return len(hp.data)
}

/** 向 heap 中添加元素 */
func (hp * Heap) Add (value Value) {
      pos := len(hp.data)
      hp.data = append (hp.data, value)

      for ;pos > 0; {
            ppos := hp.GetParInd (pos)
            pv := *(hp.Get(ppos))
            if hp.cmf (value, pv) < 0 {
                  hp.data[pos] = pv
                  pos = ppos
            } else {
                  break
            }
      }
      hp.data[pos] = value
}

/** 重建以索引 root 为根结点的子树重建堆 */
func (hp * Heap) Rebuild(root int) {
      pos := root
      value := *(hp.Get(root))

      for ;pos < len(hp.data); {
            lpos := hp.GetLCInd (pos)
            rpos := hp.GetRCInd (pos)
            var mcpos int

            if lpos != -1 {
                  if rpos != -1 && hp.cmf (*(hp.Get(rpos)), *(hp.Get(lpos))) < 0 {
                        mcpos = rpos;
                  } else {
                        mcpos = lpos
                  }

                  if (hp.cmf (value, *(hp.Get(mcpos))) < 0) {
                        hp.data[pos] = value
                        return
                  }

                  hp.data[pos] = *(hp.Get(mcpos))
                  pos = mcpos
            } else {
                  hp.data[pos] = value
                  return
            }
      }

      return
}

/** 删除 heap 的根元素 */
func (hp * Heap) Pop() (value Value) {
      n := len(hp.data)
      if n == 0 {
            return nil
      }

      value = *(hp.Get(0))
      vlast := *(hp.Get(n - 1))
      hp.data = hp.data[0:(n - 1)]

      if len(hp.data) == 0 {
            return value
      }

      hp.data[0] = vlast
      hp.Rebuild(0)

      return value
}


type PQueue Heap

/* 初始化优先队列 */
func InitPQueue (cmf CompFunc) * PQueue {
      return (* PQueue)(InitHeap (cmf))
}

/* 获得优先队列长度 */
func (pq * PQueue) GetSize () int {
      return (* Heap)(pq).GetSize()
}

/* 向优先队列中添加元素 */
func (pq * PQueue) Add (value Value) {
      (* Heap)(pq).Add(value)
}

/* 从优先队列中取出最小元素 
 * \return -1: 队列空; 0: 成功
 */
func (pq * PQueue) Pop () (value Value) {
      return (* Heap)(pq).Pop()
}

/* 清空优先队列 */
func (pq * PQueue) Flush () {
      hp := (* Heap)(pq)
      hp.data = hp.data[0:0]
}

/* 堆排序: 从大到小. */
func HeapSort (data []Value, cmf CompFunc) {

      var i int
      var hp Heap

      hp.data = data
      hp.cmf = cmf

      max := len(hp.data)

      for i = max / 2 - 1; i >= 0; i -- {
            hp.Rebuild(i)
      }

      for i = max - 1; i > 0; i -- {
            hp.data[i], hp.data[0] = hp.data[0], hp.data[i]
            hp.data = hp.data[0:i]
            hp.Rebuild(0)
      }
}

