package UnionFindPkg

type UnionFind struct {
	IDs   []int
	Sizes []int
}

func (us *UnionFind) ParentOf(x int) int {
	if x < 0 || x >= len(us.IDs) {
		panic("invalid x")
	}

	return us.IDs[x]
}

func MakeUnionFind(size int) *UnionFind {
	IDs := make([]int, size)
	Sizes := make([]int, size)

	for i := range IDs {
		IDs[i] = i
	}

	for i := range Sizes {
		Sizes[i] = 1
	}

	return &UnionFind{
		IDs:   IDs,
		Sizes: Sizes,
	}
}

func (us *UnionFind) Unify(x int, y int) {
	len := len(us.IDs)

	if x < 0 || x >= len {
		panic("invalid x")
	}

	if y < 0 || y >= len {
		panic("invalid y")
	}

	xRoot := us.Find(x)
	yRoot := us.Find(y)

	if us.Sizes[xRoot] >= us.Sizes[yRoot] {
		// merge into xRoot
		us.Sizes[xRoot] += us.Sizes[yRoot]
		us.Sizes[yRoot] = 0
		us.IDs[yRoot] = xRoot
	} else {
		// merge into yRoot
		us.Sizes[yRoot] += us.Sizes[xRoot]
		us.Sizes[xRoot] = 0
		us.IDs[xRoot] = yRoot
	}
}

func (us *UnionFind) Find(x int) int {
	cur := x

	for cur != us.ParentOf(cur) {
		cur = us.ParentOf(cur)
	}

	root := cur

	cur = x

	for cur != us.ParentOf(cur) {
		us.IDs[cur] = root
		cur = us.ParentOf(cur)
	}

	return root
}
