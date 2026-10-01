#include <iostream>
#include <vector> // WICHTIG: Stellt std::vector bereit; vector can shrink and grow in size, otherwise behaves like an array
#include <cmath>

// CNN

// Input -> Conv -> ReLU -> Pooling -> Flatten -> Dense / Fully connected layer -> Softmax
// Backtracking :C

// Input
std::vector<float> image;

/*
int width = ...;
int height = ...;
int channels = ...;
merken
*/

// -> bekommt ein schon in vector umgewandeltes Bild mit den Maßen

// Conv
class ConvLayer {
private:
    float weights[3][3]; // gibt an wie viele Gewichte gespeichert werden können, hier 3x3 Feld
    float bias; // main soll nicht die Gewichte ändern, nur die jeweiligen Funktionen
public:
    void forward(); // Convolution
};

// ReLU
void relu() {
    // ReLU
}

// MaxPooling -> größter Wert wird genommen
std::vector<float> maxpooling(std::vector<float>&previmage, int width, int height, int kernel) {
    std::vector<float> newimage = {};

    // die Zeile u. Spalte die nicht vollständig vom kernel erfasst wird, wird ignoriert

    // ss = startspalte, sz = startzeile
     
    for (int ss = 0; ss+kernel <= height; ss += kernel) { // kernel verschieben vertikal
    for (int sz = 0; sz+kernel <= width; sz += kernel) { // kernel verschieben horizontal
        float loc = image[ss * width + sz];
        for (int y = 0; y < kernel; ++y) { // kernel durchgehen
            for (int x = 0; x < kernel; ++x) {

                int index = (ss + y) * width + (ss + x);
                if (previmage[index] > loc) { // kernel max nehmen
                loc = previmage[index];
            }
        }
    }
        newimage.push_back(loc); // kernel max eintragen
        }
    }
    return newimage;
}
// im[1,2,3,4,5,6,7,8,9] w3 h3


// Dense / Fully Connected Layer
class DenseLayer {
private:
    // Gewichte
    // Bias

public:
    void forward();
};

// Flatten
void flatten(){
    // Flatten
}

//Softmax

// softmax(zi) = exp(zi) sum j=1 to d exp zj ; 1 ≤ i ≤ d 
float softmax(const std::vector<float>& zahlen, int K, int i){ // & pointer mäßig
    
    if (zahlen.empty() || i < 0 || i >= K || i >= zahlen.size()) { // abfangen
        return 0.0f; // f = float
    }
    
    float numerator = std::exp(zahlen[i]);

    float denominator = 0.0f;

    for (int j = 0; j < K; ++j ){ // K die Klassen zwischen denen Unterschieden wird (Ziel, nicht Ziel)
        denominator += std::exp(zahlen[j]);
    }
    return numerator / denominator;
}

int main() {
    relu();
    // maxpooling();
    DenseLayer dense;
    dense.forward();
    return 0;
}