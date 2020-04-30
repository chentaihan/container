package btree

type IBTree interface {
	Clone() *BTree
	ReplaceOrInsert(item Item) Item
	Delete(item Item) Item
	DeleteMin() Item
	DeleteMax() Item
	Get(key Item) Item
	Min() Item
	Max() Item
	Has(key Item) bool
	Len() int
	Clear(addNodesToFreelist bool)
}

// Item represents a single object in the binaryTree.
type Item interface {
	// Less tests whether the current item is less than the given argument.
	//
	// This must provide a strict weak ordering.
	// If !a.Less(b) && !b.Less(a), we treat this to mean a == b (i.e. we can only
	// hold one of either a or b in the binaryTree).
	Less(than Item) bool
}
