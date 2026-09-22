package rttstar

import (
	k "CODE/PATHPLANNING/KDTREE"
	ty "CODE/PATHPLANNING/TYPES"
	wm "CODE/WORLDMODEL/V1"
	"math/rand/v2"
)

type RRTStarTree struct { // the tree used by the rrt* algorithm
	Root  *ty.RRTNode
	Nodes []*ty.RRTNode
}

func RRTstar(qinit, qgoal ty.PointRRT, N int, rad, deltaQ, goalRadius, min, max float64) (*RRTStarTree, *ty.RRTNode) {
	// func RRTstar(qinit ty.PointRRT, N int, rad, deltaQ, min, max float64) *RRTStarTree {
	start := &ty.RRTNode{Point: qinit, Parent: nil, Children: []*ty.RRTNode{}, Cost: 0} // generates the starting point
	tree := &RRTStarTree{Root: nil, Nodes: []*ty.RRTNode{}}                             // generates an empty rrt* tree

	return rRTStarHilf(start, qgoal, N, tree, rad, deltaQ, min, max, goalRadius)
}

func rRTStarHilf(qinit *ty.RRTNode, qgoal ty.PointRRT, N int, T *RRTStarTree, rad, deltaQ, goalRadius, min, max float64) (*RRTStarTree, *ty.RRTNode) {

	T.insertNode(nil, qinit)       // insert the starting point into the rrt* tree
	KTree := ty.KDTree{Root: nil}  // generates an empty kd tree
	k.KDInsertPoint(&KTree, qinit) // insert the starting point into the kd tree

	var bestGoal *ty.RRTNode // the path leading to this point is the best one found

	for count := 0; count < N; count++ {

		var qrand *ty.RRTNode

		if rand.Float64() < 0.1 { // Goal Bias (10%)
			qrand = &ty.RRTNode{Point: qgoal}
		} else {
			qrand = randompoint(min, max)
		}

		qnearest := k.NearestNeighbor(qrand, KTree) // find the nearest neighbor using the kd tree
		if qnearest == nil {                        // if there isnt any, skip
			continue
		}
		qnew := Steer(qnearest, qrand, deltaQ) // qnews distance to qnearest is at most deltaQ
		if qnew == nil {
			continue
		}
		if !wm.ObstacleFree(qnearest.Point, qnew.Point) { // check if path is empty
			continue
		}
		qparent := ChooseParent(qnew, qnearest, rad, KTree) // determining the best parent from Q_neat (difference RRT and RRT*)
		if qparent == nil {
			continue
		}

		T.insertNode(qparent, qnew) // insert into rrt* tree

		k.KDInsertPoint(&KTree, qnew) // insert qnew now, otherwise the point will find itself as its nearest neighbor

		T.Rewire(qnew, rad, KTree) // check if a already existing node can find a better path using the new connection

		// check if goal can be reached
		distGoal := ty.DistanceBetweenPoints(qnew.Point, qgoal)
		if distGoal <= goalRadius && wm.ObstacleFree(qnew.Point, qgoal) { // if the point lays within the searched radius and the path is reachable
			goalCost := qnew.Cost + distGoal                 // then the cost will be calculated
			if bestGoal == nil || goalCost < bestGoal.Cost { // if the path is better than the ones before
				bestGoal = &ty.RRTNode{Point: qgoal, Parent: qnew, Children: nil, Cost: goalCost} // take it
			}
		}
	}

	return T, bestGoal // after weve done it for enough iterations, we can stop
}

func ChooseParent(qrand *ty.RRTNode, qnearest *ty.RRTNode, rad float64, KT ty.KDTree) *ty.RRTNode {
	if qrand == nil || qnearest == nil {
		return nil
	}
	// nearest neighbor as first parent
	qmin := qnearest
	cmin := qnearest.Cost + ty.DistanceBetweenPoints(qnearest.Point, qrand.Point)
	Qnear := k.RadiusSearch(qrand, KT, rad) // find all points located within a radius rad of qrand

	for _, candidate := range Qnear { // check every neighbor
		if candidate == nil {
			continue
		}
		if !wm.ObstacleFree(candidate.Point, qrand.Point) { // check if path empty
			continue
		}
		cnew := candidate.Cost + ty.DistanceBetweenPoints(candidate.Point, qrand.Point) // determining the cost
		if cnew < cmin {                                                                // check if this parent is better
			cmin = cnew
			qmin = candidate
		}
	}
	return qmin
}

// check if we can find a better path for already existing nodes using qrand
func (T *RRTStarTree) Rewire(qrand *ty.RRTNode, rad float64, KT ty.KDTree) {
	if qrand == nil {
		return
	}
	Qnear := k.RadiusSearch(qrand, KT, rad) // finding Qnears
	for _, qnear := range Qnear {           // checking neighbors
		if qnear == nil {
			continue
		}
		if qnear == qrand { // cant be connected to itself
			continue
		}
		if !wm.ObstacleFree( // path empty?
			qrand.Point,
			qnear.Point,
		) {
			continue
		}
		newCost := qrand.Cost + ty.DistanceBetweenPoints(qrand.Point, qnear.Point) // cost if it would change the parent
		if newCost < qnear.Cost {                                                  // path better -> rewire
			if qnear.Parent != nil {
				oldParent := qnear.Parent
				for i, child := range oldParent.Children {
					if child == qnear {
						oldParent.Children = append(oldParent.Children[:i], oldParent.Children[i+1:]...)
						break
					}
				}
			}
			qnear.Parent = qrand                           // new parent
			qrand.Children = append(qrand.Children, qnear) // register qnear as a child of qrand
			qnear.Cost = newCost                           // new cost
			updateChildCosts(qnear)                        // cost of the children have to be updated as well
		}
	}
}

func updateChildCosts(node *ty.RRTNode) {
	if node == nil {
		return
	}
	for _, child := range node.Children {
		if child == nil {
			continue
		}
		child.Cost = node.Cost + ty.DistanceBetweenPoints(node.Point, child.Point) // Kosten des Parents + Kosten des Parents zum Kind
		updateChildCosts(child)                                                    // Rekursiv auch die Kinder aktualisieren
	}
}

func (T *RRTStarTree) insertNode(qparent *ty.RRTNode, qson *ty.RRTNode) {
	if qson == nil {
		return
	}

	if T.Root == nil && qparent == nil { // falls noch kein root qson zum root
		T.Root = qson
	} else {
		qson.Parent = qparent                                                          // parent des neuen Knoten machen
		qson.Cost = qparent.Cost + ty.DistanceBetweenPoints(qson.Point, qparent.Point) // kosten Start bis qson
		qparent.Children = append(qparent.Children, qson)                              // als kind speichern
	}

	T.Nodes = append(T.Nodes, qson)
}

// Steer bewegt sich von qfrom in Richtung qto, maximal um deltaQ, dadurch RRT-Schritt nicht beliebig lang
func Steer(qfrom *ty.RRTNode, qto *ty.RRTNode, deltaQ float64) *ty.RRTNode {
	if qfrom == nil || qto == nil {
		return nil
	}
	distance := ty.DistanceBetweenPoints(qfrom.Point, qto.Point)
	if distance <= deltaQ { // Punkt liegt kürzer entfernt als deltaQ
		return &ty.RRTNode{Point: qto.Point, Parent: nil, Children: []*ty.RRTNode{}, Cost: 0}
	}

	// Richtungsvektor von qfrom nach qto
	dx := (qto.Point.X - qfrom.Point.X) / distance
	dy := (qto.Point.Y - qfrom.Point.Y) / distance
	dz := (qto.Point.Z - qfrom.Point.Z) / distance

	// Neuer Punkt liegt genau deltaQ entfernt.
	newPoint := ty.PointRRT{X: qfrom.Point.X + dx*deltaQ, Y: qfrom.Point.Y + dy*deltaQ, Z: qfrom.Point.Z + dz*deltaQ}
	return &ty.RRTNode{Point: newPoint, Parent: nil, Children: []*ty.RRTNode{}, Cost: 0}
}

// generates random values for x,y,z in a given intervall
func randompoint(min, max float64) *ty.RRTNode {
	p := ty.PointRRT{X: randomfloat(min, max), Y: randomfloat(min, max), Z: randomfloat(min, max)}
	return &ty.RRTNode{Point: p, Parent: nil, Children: []*ty.RRTNode{}, Cost: 0}
}

// RangeFloat generates a random number in a given intervall [min, max)
func randomfloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func ExtractPath(goal *ty.RRTNode) []ty.PointRRT {
	path := []ty.PointRRT{}
	current := goal
	for current != nil {
		path = append(path, current.Point)
		current = current.Parent
	}
	// umdrehen
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
