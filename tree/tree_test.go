package tree_test

import "github.com/bagashiz/go-dsa/tree"

var treesInt = []*tree.BinaryNode[int]{
	{
		Data: 20,
		Right: &tree.BinaryNode[int]{
			Data: 50,
			Right: &tree.BinaryNode[int]{
				Data:  100,
				Right: nil,
				Left:  nil,
			},
			Left: &tree.BinaryNode[int]{
				Data: 30,
				Right: &tree.BinaryNode[int]{
					Data:  45,
					Right: nil,
					Left:  nil,
				},
				Left: &tree.BinaryNode[int]{
					Data:  29,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &tree.BinaryNode[int]{
			Data: 10,
			Right: &tree.BinaryNode[int]{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &tree.BinaryNode[int]{
				Data: 5,
				Right: &tree.BinaryNode[int]{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	},
	{
		Data: 20,
		Right: &tree.BinaryNode[int]{
			Data:  50,
			Right: nil,
			Left: &tree.BinaryNode[int]{
				Data: 30,
				Right: &tree.BinaryNode[int]{
					Data: 45,
					Right: &tree.BinaryNode[int]{
						Data:  49,
						Right: nil,
						Left:  nil,
					},
					Left: nil,
				},
				Left: &tree.BinaryNode[int]{
					Data:  29,
					Right: nil,
					Left: &tree.BinaryNode[int]{
						Data:  21,
						Right: nil,
						Left:  nil,
					},
				},
			},
		},
		Left: &tree.BinaryNode[int]{
			Data: 10,
			Right: &tree.BinaryNode[int]{
				Data:  15,
				Right: nil,
				Left:  nil,
			},
			Left: &tree.BinaryNode[int]{
				Data: 5,
				Right: &tree.BinaryNode[int]{
					Data:  7,
					Right: nil,
					Left:  nil,
				},
				Left: nil,
			},
		},
	},
}
