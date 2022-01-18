package dsg

type AvlNode struct {
      value Value;
      left * AvlNode;
      right * AvlNode;
      comp_func CompFunc;
      bf int;
}

type AvlTree struct {
      root * AvlNode;
      comp_func CompFunc;
}

func InitAvlTree (comp_func CompFunc) * AvlTree {
      return &(AvlTree{comp_func: comp_func})
}

func (tr * AvlTree) Add (v Value) {
      if tr.root == nil {
            tr.root = InitAvlNode (v, tr.comp_func)
      } else {
            tr.root, _ = tr.root.Add (v)
      }
}

func (tr * AvlTree) Remove (v Value) {
      if tr.root != nil {
            tr.root, _ = tr.root.Remove (v)
      }
}

func InitAvlNode (v Value, comp_func CompFunc) * AvlNode {
      return &(AvlNode{value: v, comp_func: comp_func}) 
}

func (root * AvlNode) Add (v Value) (new_root * AvlNode, new_deeper int) {
      var deeper int
      if root.comp_func (v, root.value) < 0 {
            if root.left == nil {
                  root.left = InitAvlNode (v, root.comp_func)
            } else {
                  root.left, deeper = root.left.Add (v)
                  if deeper == 0 {
                        return root, 0
                  }
            }
            if root.bf == -1 {
                  new_root = root.BalanceLeft()
                  new_deeper = 0
                  return
            }
            root.bf --
            new_deeper = root.bf
            new_root = root
            return
      }

      if root.right == nil {
            root.right = InitAvlNode (v, root.comp_func)
      } else {
            root.right, deeper = root.right.Add (v)
            if deeper == 0 {
                  return root, 0
            }
      }
      if root.bf == 1 {
            new_root = root.BalanceRight()
            new_deeper = 0
            return
      }
      root.bf ++
      new_deeper = root.bf
      new_root = root
      return
}

func (root * AvlNode) BalanceLeft () * AvlNode {
      var p, q * AvlNode
      if root.left.bf < 0 {
            p = root.left
            root.left = p.right
            root.bf = 0
      } else {
            q = root.left
            p = q.right
            q.right = p.left
            root.left = p.right
            p.left = q
            root.bf = 0
            q.bf = 0
            if p.bf == -1 {
                  root.bf = 1
            }
            if p.bf == 1 {
                  q.bf = -1
            }
      }
      p.right = root
      p.bf = 0
      return p
}

func (root * AvlNode) BalanceRight () * AvlNode {
      var p, q * AvlNode
      if root.right.bf > 0 {
            p = root.right
            root.right = p.left
            root.bf = 0
      } else {
            q = root.right
            p = q.left
            q.left = p.right
            root.right = p.left
            p.right = q
            root.bf = 0
            q.bf = 0
            if p.bf == 1 {
                  root.bf = -1
            }
            if p.bf == -1 {
                  q.bf = 1
            }
      }
      p.left = root
      p.bf = 0
      return p
}

func (root * AvlNode) Remove (v Value) (new_root * AvlNode, junk * AvlNode) {
      comp_res := root.comp_func (v, root.value)
      if comp_res == 0 {
            junk = root
            if root.right == nil {
                  return root.left, junk
            }
            oldbf := root.right.bf
            root.right, new_root = root.right.RemoveLeftMostDescendant ()
            new_root.left = root.left
            new_root.right = root.right
            new_root.bf = root.bf
            new_root = new_root.RestoreRightBalance(oldbf)
            return 
      } else if comp_res < 0 {
            if root.left == nil {
                  return root, nil
            }
            oldbf := root.left.bf
            root.left, junk = root.left.Remove (v)
            new_root = root.RestoreLeftBalance (oldbf)
            return 
      } else {
            if root.right == nil {
                  return root, nil
            }
            oldbf := root.right.bf
            root.right, junk = root.right.Remove (v)
            new_root = root.RestoreRightBalance (oldbf)
            return 
      }
}

func (root * AvlNode) RemoveLeftMostDescendant () (new_root * AvlNode, junk * AvlNode) {
      left_child := root.left
      if left_child == nil {
            junk = root
            new_root = root.right
            return
      }
      oldbf := left_child.bf
      root.left, junk = left_child.RemoveLeftMostDescendant ()
      new_root = root.RestoreLeftBalance (oldbf)
      return
}

func (root * AvlNode) RestoreLeftBalance (oldbf int) * AvlNode {
      left_child := root.left
      if left_child == nil {
            root.bf ++
      } else if (left_child.bf != oldbf && left_child.bf == 0) {
            root.bf ++
      }
      if root.bf > 1 {
            return root.BalanceLeft()
      }
      return root
}


func (root * AvlNode) RestoreRightBalance (oldbf int) * AvlNode {
      right_child := root.right
      if right_child == nil {
            root.bf --
      } else if (right_child.bf != oldbf && right_child.bf == 0) {
            root.bf --
      }
      if root.bf < -1 {
            return root.BalanceRight()
      }
      return root
}
