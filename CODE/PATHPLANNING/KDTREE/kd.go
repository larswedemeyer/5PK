package kd

// the K-D TREE
import (
	ty "CODE/PATHPLANNING/TYPES"
	"math"
)

// building a K-D Tree by adding individual Points

// The tree is created dynamically; the slowdown can be neglected considering the random distribution of points
func KDInsertPoint(t *ty.KDTree, p *ty.RRTNode) {
	if t == nil || p == nil {
		return // Exception Handling
	}
	t.Root = insertPointHilf(t.Root, p, 0) // starting in the first Dimension and by the root, the point will be inserted into the Tree
}

func insertPointHilf(node *ty.KDNode, rrtnode *ty.RRTNode, d int) *ty.KDNode { // node current KD node, rrtnode node which is to be inserted, d dimension
	if node == nil { // empty node -> add Point
		return &ty.KDNode{Value: rrtnode, Left: nil, Right: nil}
	}

	dimension := d % 3 // determining the current dimension

	switch dimension { // depending on the dimension, compare whether the nn needs to go to the left or the right
	case 0: // X-Axis
		if rrtnode.Point.X < node.Value.Point.X { // if the rrtode is smaller it has to go left
			node.Left = insertPointHilf(node.Left, rrtnode, d+1) // -> increasing the dimension by one
		} else { // if bigger right
			node.Right = insertPointHilf(node.Right, rrtnode, d+1)
		}
	case 1: // Y-Axis
		if rrtnode.Point.Y < node.Value.Point.Y {
			node.Left = insertPointHilf(node.Left, rrtnode, d+1)
		} else {
			node.Right = insertPointHilf(node.Right, rrtnode, d+1)
		}
	case 2: // Z-Axis
		if rrtnode.Point.Z < node.Value.Point.Z {
			node.Left = insertPointHilf(node.Left, rrtnode, d+1)
		} else {
			node.Right = insertPointHilf(node.Right, rrtnode, d+1)
		}
	}

	return node // return Point, finish
}

func NearestNeighbor(p *ty.RRTNode, t ty.KDTree) *ty.RRTNode {
	if p == nil || t.Root == nil { // Exception Handling
		return nil
	}
	best := nearestNeighborHilf(p.Point, t.Root, nil, math.Inf(1), 0) // p.Point = point coords, t.Root = starting Point, nil = current best, math.Inf(1) = pos infinite, 0 = current dimension

	return best // return the best points
}

func nearestNeighborHilf(target ty.PointRRT, node *ty.KDNode, best *ty.RRTNode, bestDistance float64, d int) *ty.RRTNode {

	if node == nil { // when the search is finished, best will be returned
		return best
	}

	distance := ty.DistanceBetweenPoints(target, node.Value.Point) // Abs of the current point to the target Point

	if distance < bestDistance { // is the current node the best?
		best = node.Value       // if yes update
		bestDistance = distance // update best distance
	}

	dimension := d % 3 // determining through which dimension the hyperplane is splitting the area

	var targetValue float64
	var nodeValue float64

	switch dimension {

	case 0: // dimension
		targetValue = target.X         // X Coords of the searched Point
		nodeValue = node.Value.Point.X //  current Point X coords

	case 1:
		targetValue = target.Y
		nodeValue = node.Value.Point.Y

	case 2:
		targetValue = target.Z
		nodeValue = node.Value.Point.Z
	}

	var primary *ty.KDNode
	var secondary *ty.KDNode

	if targetValue < nodeValue { // if the coords of the searched Point are smaller than the one of the current one, then it is located on the left side
		primary = node.Left    // so we have to continue our search in the left part of the tree in any case
		secondary = node.Right // on the right side only under specific cinditions
	} else { // in the other case same but reversed
		primary = node.Right
		secondary = node.Left
	}

	best = nearestNeighborHilf(target, primary, best, bestDistance, d+1) // searching in the primary tree

	if best != nil { // if there is a best, determine the best Distance
		bestDistance = ty.DistanceBetweenPoints(target, best.Point)
	}

	planeDistance := math.Abs(targetValue - nodeValue) // Abs of the searched point to the hyperplane

	if planeDistance < bestDistance { // if the other side could contain a better point search in it too
		best = nearestNeighborHilf(target, secondary, best, bestDistance, d+1)
	}
	return best
}

// K-D NÄHESTEN PUNKT SUCHEN
// -- DIMENSION NOCH ZYKLISCH; SPÄTER NACH SPREAD

// liefert den einen am nächstgelegensten Punkt
func SearchPoint(p *ty.KDNode, T ty.KDTree, e float64, k int) []*ty.KDNode { // ges Punkt, Baum, abs, anzahl
	L := []*ty.KDNode{}
	return searchPointInKDTreeHilf(p, L, T.Root, e, k, 0)
}

func searchPointInKDTreeHilf(p *ty.KDNode, L []*ty.KDNode, TNode *ty.KDNode, e float64, k int, d int) []*ty.KDNode {
	if TNode == nil {
		return L // ende des durchsuchens
	}

	if ty.DistanceBetweenPoints(TNode.Value.Point, p.Value.Point) < e { // wenn der abs des aktuellen Punktes zum Gesuchten Punkt kleiner als e ist, zu Liste hinzufügen
		L = append(L, TNode)
	}

	// k benutzen

	dimension := d % 3 // Durch die aktuelle Dimension läuft das Hyperplane, nach dem wird geschaut auf welcher Seite des Raums der Punkt liegt
	var a, b float64

	switch dimension {
	case 0: // schauen auf welcher Seite der Punkt liegt
		a = p.Value.Point.X
		b = TNode.Value.Point.X
	case 1:
		a = p.Value.Point.Y
		b = TNode.Value.Point.Y
	case 2:
		a = p.Value.Point.Z
		b = TNode.Value.Point.Z
	}
	var primaryTree *ty.KDNode
	var secondaryTree *ty.KDNode

	if a < b { // left from the root
		primaryTree = TNode.Left
		secondaryTree = TNode.Right
	} else {
		secondaryTree = TNode.Left
		primaryTree = TNode.Right
	}

	L = searchPointInKDTreeHilf(p, L, primaryTree, e, k, d+1)

	worstNode := highestDistanceBetweenPoints(L, p)

	if len(L) < k && math.Abs(a-b) < e { // there are still empty places in L, the Hyperplane is nearer than allowed
		L = searchPointInKDTreeHilf(p, L, secondaryTree, e, k, d+1)
	} else if worstNode != nil && ty.DistanceBetweenPoints(worstNode.Value.Point, p.Value.Point) > math.Abs(a-b) && math.Abs(a-b) <= e { // schauen ob der am weitesten entfernteste näher ist als das Hyperplane && ob Hyplerplane weiter weg ist als erlaubt
		L = searchPointInKDTreeHilf(p, L, secondaryTree, e, k, d+1)
	}

	return L
}

// RadiusSearch returns the Pointss which lay in a radius Q_near around the goalpoint
func RadiusSearch(p *ty.RRTNode, t ty.KDTree, radius float64) []*ty.RRTNode {
	result := []*ty.RRTNode{} // die ergebnisse sind eine Leere Liste

	if p == nil || t.Root == nil { // if there isnt any point or tree, it cannot be that a point lays within the rad
		return result
	}

	return radiusSearchHilf(p.Point, t.Root, radius, result, 0) // the empty list is filled with points which lay within the rad
}

func radiusSearchHilf(target ty.PointRRT, node *ty.KDNode, radius float64, result []*ty.RRTNode, d int) []*ty.RRTNode {

	if node == nil { // end of the search
		return result
	}

	distance := ty.DistanceBetweenPoints(target, node.Value.Point) // Abs to the current point

	if distance <= radius { // does the point lay within the rad?
		result = append(result, node.Value) // if yes, its valid and can be added
	}

	dimension := d % 3 // dimension

	var targetValue float64
	var nodeValue float64

	switch dimension {

	case 0: // case of the dimension
		targetValue = target.X         // targetValue coord of the searched point
		nodeValue = node.Value.Point.X // nodeValue coord of the current point

	case 1:
		targetValue = target.Y
		nodeValue = node.Value.Point.Y

	case 2:
		targetValue = target.Z
		nodeValue = node.Value.Point.Z
	}

	if targetValue < nodeValue { // if the searched point is smaller than the current one, continue the search on the left side
		result = radiusSearchHilf(target, node.Left, radius, result, d+1)

		if math.Abs(targetValue-nodeValue) <= radius { // if points of the other side could lay within the rad, continue the search there aswell
			result = radiusSearchHilf(target, node.Right, radius, result, d+1) // continue the search on the right side
		}
	} else { // in the other case, the same way but turned around (a bit bad formulated)
		result = radiusSearchHilf(target, node.Right, radius, result, d+1)
		if math.Abs(targetValue-nodeValue) <= radius {
			result = radiusSearchHilf(target, node.Left, radius, result, d+1)
		}
	}

	return result
}

// from a list of points, return the one which is the one with the highest distance
func highestDistanceBetweenPoints(ps []*ty.KDNode, p *ty.KDNode) *ty.KDNode {
	if len(ps) == 0 || p == nil {
		return nil
	}

	maxNode := ps[0]
	maxDist := ty.DistanceBetweenPoints(p.Value.Point, ps[0].Value.Point)

	for _, node := range ps[1:] {
		if node == nil {
			continue
		}
		dist := ty.DistanceBetweenPoints(p.Value.Point, node.Value.Point)
		if dist > maxDist {
			maxDist = dist
			maxNode = node
		}
	}

	return maxNode
}
