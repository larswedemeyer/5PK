// --> VERSION 1, VERWENDET WERDEN STARRE REGELN

// VERSION 2, HYBRID-MODELL

// VERSION 3, ECHTES RL

// NUR ENTSCHEIDUNG WAS ZU TUN, WIE BEWEGT WIRD WIRD GETRENNT BESTIMMT

package verhaltensweise

import (
	wm "CODE/WORLDMODEL/V1"
)

// BALL SUCHEN
// ZUM BALL LAUFEN
// 		BALL AUFHEBEN

// BRAUCHT BALL EVERSEEN
// DISTANCE TOBALL

func Decisions(w *wm.World) {
	if !w.Ball.EverSeen {
		SearchBall()
	} else {
		abs, _ := w.DistanceToBall()
		if abs < 0.5 {
			PickBallUp()
		} else {
			GoToBall()
		}
	}
}

func PickBallUp() {}

func SearchBall() {}

func GoToBall() {

}

// Verweist bei der Bewegung auf RRT
