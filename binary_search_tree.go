package binarysearchtree

type stack struct {
  items []*BinarySearchTree
}
func (s *stack) push(item *BinarySearchTree) {
  s.items = append(s.items, item)
}
func (s *stack) pop() *BinarySearchTree {
  if len(s.items) == 0 {
    return nil
  }
  bst := s.items[len(s.items)-1]
  s.items = s.items[:len(s.items)-1]
  return bst
}

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
  bst := new(BinarySearchTree)
  bst.data = i
  return bst
}

func (bst *BinarySearchTree) length() int {
  left := 0
  right := 0
  if bst.left != nil {
    left = bst.left.length() 
  }
  if bst.right != nil {
    right = bst.right.length()
  }
  if left > right {
    return left + 1
  }
  return right + 1
}
func (bst *BinarySearchTree) lengthR() int {
  if bst.right != nil {
    return bst.right.length()
  }
  return 0
}
func (bst *BinarySearchTree) lengthL() int {
  if bst.left != nil {
    return bst.left.length()
  }
  return 0
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) *BinarySearchTree {
  root := bst
  for {
    if i <= bst.data {
      if bst.left == nil {
        bst.left = new(BinarySearchTree)
        bst.left.data = i
        break
      } else {
        bst = bst.left
      }
    } else {
      if bst.right == nil {
        bst.right = new(BinarySearchTree)
        bst.right.data = i
        break
      } else {
        bst = bst.right
      }
    }
  }

  // rebalance
  left := 0
  right := 0
  if root.left != nil {
    left = root.lengthL()
  } 
  if root.right != nil {
    right = root.lengthR()
  } 

  var newRoot *BinarySearchTree
  if left > right && left - right > 1 {
    if root.left.lengthL() > root.left.lengthR() {
      newRoot = root.left
      root.left = newRoot.right
      newRoot.right = root
    } else {
      newRoot = root.left.right
      root.left.right = newRoot.left
      newRoot.left = root.left
      root.left = newRoot.right
      newRoot.right = root
    }
    swap(root, newRoot)
    root.right = newRoot
  } else if right > left && right - left > 1 {
    if root.right.lengthL() > root.right.lengthR() {
      newRoot = root.right.left
      root.right.left = newRoot.right
      newRoot.right = root.right
      root.right = newRoot.left
      newRoot.left = root
    } else {
      newRoot = root.right
      root.right = newRoot.left
      newRoot.left = root
    }
    swap(root, newRoot)
    root.left = newRoot
  }
  return root
}

func swap(bst1, bst2 *BinarySearchTree) {
  bst1.right, bst2.right = bst2.right, bst1.right
  bst1.left, bst2.left = bst2.left, bst1.left
  bst1.data, bst2.data = bst2.data, bst1.data
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
  stack := new(stack)
  var output []int

  for bst != nil {
    stack.push(bst) 
    if bst.left != nil {
      bst = bst.left
      continue
    }
    for {
      bst = stack.pop()
      if bst == nil {
        break
      }
      output = append(output, bst.data)
      if bst.right == nil {
        continue
      }
      bst = bst.right
      break
    }
  } 
  return output
}
