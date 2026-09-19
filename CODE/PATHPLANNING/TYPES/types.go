package types

import "math"

// A POINT IN 3-DIMENSIONAL SPACE

type PointRRT struct {
	X, Y, Z float64
}

// BLATT (BINÄR)

type KDNode struct {
	Value *RRTNode // points to a RRTNODE, no duplicates
	Left  *KDNode
	Right *KDNode
}

// KD TREE

type KDTree struct {
	Root *KDNode
}

// NODE of an RRT TREE

type RRTNode struct {
	Point    PointRRT
	Parent   *RRTNode
	Children []*RRTNode
	Cost     float64 // pathcost from start up to this point
}

// calculate the distance between two points in 3 dimensional space

func DistanceBetweenPoints(p1, p2 PointRRT) float64 {
	dx := p1.X - p2.X // abs between to points using the Euklidischen Distanz im 3D-Raum
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz) // -- math.Pow has a longer runtime as dx*dx
	// -- return math.Sqrt(math.Pow((p1.X-p2.X), 2)+math.Pow((p1.Y-p2.Y), 2)+math.Pow((p1.Z-p2.Z), 2))
}
