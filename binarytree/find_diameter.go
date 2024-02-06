package main

type treeInfo struct {
	height   int
	diameter int
}

func getTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{
			0, 0,
		}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeInfo.height + rightTreeInfo.height

	maxDiameterSoFar := max(leftTreeInfo.diameter, rightTreeInfo.diameter)

	currentDiameter := max(longestPathThroughRoot, maxDiameterSoFar)

	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)

	return treeInfo{
		currentHeight, currentDiameter,
	}
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).diameter
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}
