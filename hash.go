package dsg

/** Hash 不负责分配空间. */

/** Hash 叶节点结构 */
type Node struct {
      key Value
      value Value
      next * Node
}

/** Hash 非叶节点结构, 若其子节点为叶节点, 则 len(sublist) = 0,
 *  head 指向首个子节点, tail 指向最后一个子节点.
 *  若其子节点为非叶节点, 则 head = tail = nil, sublist 指向子节点数组, 
 *  n 为子节点个数.*/
type Chain struct {
      sublist []Chain
      head * Node
      tail * Node
}

/** Hash 结构 */
type Hash struct {
      // 根结点
      list * Chain
      ktoi KeyToIndexFunc
      comp CompFunc
}

func init_nhash (ktoi KeyToIndexFunc, comp CompFunc) * Hash {
      var hash Hash
      var chain Chain

      hash.list = &chain
      hash.ktoi = ktoi
      hash.comp = comp

      return &hash
}

func InitHash (n int, ktoi KeyToIndexFunc, comp CompFunc) * Hash {
      var i int

      if n <= 0 {
            return nil;
      }

      hash := init_nhash (ktoi, comp)

      hash.list.sublist = make ([]Chain, n)
      hash.list.head, hash.list.tail = nil, nil

      for i = 0; i < n; i ++ {
            cc := &(hash.list.sublist[i])
            cc.sublist = make ([]Chain, 0)
            cc.head = nil
            cc.tail = nil
      }

      return hash
}

func (hash * Hash) InitSubHash (ind []int, n int) {
      var i int

      cc := hash.GotoChain (ind)
      if len(cc.sublist) != 0 || cc.head != nil || cc.tail != nil {
            return
      }

      cc.sublist = make ([]Chain, n)
      cc.head, cc.tail = nil, nil

      for i = 0; i < n; i ++ {
            sc := &(cc.sublist[i])
            sc.sublist = make ([]Chain, 0)
            sc.head = nil
            sc.tail = nil
      }
}

func InitTreeHash (n []int, ktoi KeyToIndexFunc, comp CompFunc) * Hash {
      if len(n) <= 0 {
            return nil
      }

      hash := InitHash (n[0], ktoi, comp)
      if len(n) <= 1 {
            return hash
      }

      var ind []int = make ([]int, len(n))
      var isend bool = false
      var ii int
      var cd int = 0

      for ;!isend; {
            for ii = cd + 1; ii < len(n); ii ++ {
                  hash.InitSubHash (ind[0:ii], n[ii]);
            }

            cd = len(n) - 2
            for {
                  ind[cd] ++
                  if ind[cd] < n[cd] {
                        break
                  }
                  if cd == 0 {
                        isend = true
                        break
                  }
                  ind[cd] = 0
                  cd --
            }
      }

      return hash
}

/** 跳转到 hash 的某个非叶节点 
 * \param ind 位置索引
 */
func (hash * Hash) GotoChain (ind []int) * Chain {
      var i int
      cc := hash.list
      for i = 0; i < len(ind); i ++ {
            if len(cc.sublist) == 0 || ind[i] >= len(cc.sublist) {
                  return nil
            }
            cc = &(cc.sublist[ind[i]])
      }
      return cc
}

/** 通过 key 查找 hash 的某个非叶节点, 
 * \param hash 待查找的 hash
 * \param key 用于查找的 key
 */
func (hash * Hash) SearchChain (key Value) * Chain {
      ind := hash.ktoi (key)
      cc := hash.GotoChain (ind)
      return cc;
}

/** 通过 key 查找 hash 的某个叶节点, 
 * \param hash 待查找的 hash
 * \param key 用于查找的 key
 * \param last 返回所得叶节点的前一个节点 (若为该树枝上的首个节点, 则返回 nil)
 * \param pcc 返回所得叶节点的父节点.
 * \param result 返回查找所得的节点.
 */
func (hash * Hash) SearchNode (key Value) (last * Node, pcc * Chain, result * Node) {
      cc := hash.SearchChain (key)
      if cc == nil {
            return nil, nil, nil
      }

      pcc = cc

      var cn * Node = cc.head
      var cl * Node = nil

      last = cl
      if cn == nil {
            return nil, nil, nil
      }

      for ;hash.comp(key, cn.key) != 0; {
            cl = cn
            cn = cn.next
            if cn == nil {
                  cl = nil
                  break
            }
      }

      last = cl
      result = cn

      return
}

/** 向 hash 中增加一个叶节点.
 * \param hash 待增加hash
 * \param key 增加节点的 key
 * \param value 增加节点的 value
 */
func (hash * Hash) AddData (key Value, value Value) {

      cc := hash.SearchChain (key);
      if cc == nil {
            return
      }
      if len(cc.sublist) != 0 {
            return
      }

      var cn Node = Node{key: key, value: value, next: nil}

      if cc.tail == nil {
            cc.head = &cn
            cc.tail = &cn
      } else {
            cc.tail.next = &cn;
            cc.tail = &cn;
      }
}

/** 删除 hash 中以 key 为索引的节点
 * \param hash 待删除节点所在的 hash
 * \param key 待删除节点的索引 (key)
 */
func (hash * Hash) RemoveData (key Value) {
      var cc * Chain
      var cl, cn * Node

      cl, cc, cn = hash.SearchNode (key)

      if cn == nil {
            return
      }

      if cl != nil {
            cl.next = cn.next
      } else {
            cc.head = cn.next
      }
      if cn.next == nil {
            cc.tail = cl
      }
}

/** 核心函数, 利用 Hash 完成查找, 
 * \param hash 已经建立好的 Hash
 * \param key 待查找的 key
 * \param value 返回查找得到的 value (nil: 没找到) 
 */
func (hash * Hash) Search (key Value) (value Value) {
      _, _, cn := hash.SearchNode (key)

      if cn != nil {
            value = cn.value
      } else {
            value = nil
      }

      return
}

func (hash * Hash) Clear () {
      if hash == nil || hash.list == nil {
            return
      }
      clear_sub (hash.list)
}

func clear_sub (cc * Chain) {
      if len(cc.sublist) == 0 {
            cc.head = nil
            cc.tail = nil
      } else {
            for i := 0; i < len(cc.sublist); i ++ {
                  clear_sub (&(cc.sublist[i]))
            }
      }
}

