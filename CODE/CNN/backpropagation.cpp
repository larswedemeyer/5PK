#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>
#include <string>

#define STB_IMAGE_IMPLEMENTATION
#include "stb_image.h"

using namespace std;

/* Backpropagation

// Loss Funktion Ausgabe Softmax

float losssoftmax(float expected, float given) {
    // formel c klassen H(P,Q) = - Sum C C=1 (yc log (yc^)) yc richtige yc^ vorhergesagte (Wahrscheinlichkeit)
    // je höher die Wahrscheinlichkeit der richtigen Klasse, desto niedriger der Loss

    // anfangs nur mit einer Klasse (Ball)

    float n = - (expected * log(given));
    return n;
}

// Gradient Descent

// wneu = walt - n 6L / 6w (w = Gewicht, L = Loss, n = Learning Rate, 6 = partielle Ableitung)

*/

// abstrakt und vereinfacht

float vorhersage(float w, float x){ // w ist das gewicht, x der eingabewert
    float yv; // Vorhergesagter Wert
    yv = w * x; // Berechnung des Wertes durch das CNN / ein Neuron
    return yv; // Ausgabe des Wertes
}

float loss(float yv, float y) {
    float L; // die Loss funktion
    float u = yv - y; // Unterschied zwischen vorhersage und gewünschter Ausgabe
    L = 1/2 * (u*u); // Berechnung des Fehlers mithilfe einer quadratischen Loss Funktion
    return L; // Ausgabe der Funktion
}

// nun kann man mithilfe der loss Funktion Gewichte jeweils härter bestrafen wenn sie außerhalb des gewünschten Bereiches liegen

// es gilt zu schauen in welche Richtung das Gewicht verändert werden muss damit der Fehler entsprechend kleiner wird

// da der Fehler vom Gewicht abhängt lässt sich die Funktion als Funktion des Gewichts schreiben 

// L(w) = 1/2 * (wx - y)^2              siehe oben bei der Vorhersage für yv = wx

// nun wollen wir mithilfe der Ableitung (Leibnizschreibweise) bestimmen wie sich der Fehler bei kleinen Veränderungen verhält

// dL / dw = (yv - y)x

// ABLEITEN:
// L(w) = 1/2 * (wx-y)^2
// wir setzen u = wx-y
// somit haben wir L(u) = 1/2 * u^2
// die Ableitung der äußeren Funktion dL/ du = u
// die Ableitung der inneren Funktion du/ dw = x
// nach der Kettenregel multiplizieren wir dL/du * du/dw = dL/dw
// somit erhalten wir durch einsetzen dL/ dw = (wx-y)x
// wx = yv und somit können wir schreiben: dL/ dw = (yv -y)x
// Das ist der gesuchte Gradient

// Ein positiver Gradient bedeutet, wenn wir das Gewicht erhöhen steigt der Fehler lokal
// Ein negativer Gradient bedeutet, wenn wir das Gewicht erniedrigen sinkt der Fehler lokal
// dementsprechend muss das Gewicht angepasst werden

// dann wird das Gewicht mithilfe der Lernrate n wneu = walt - n * dL / dw aktualisiert