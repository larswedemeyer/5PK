#include <iostream>
#include <vector> // WICHTIG: Stellt std::vector bereit; vector can shrink and grow in size, otherwise behaves like an array
#include <cmath>

// CNN

// Input -> Conv -> ReLU -> Pooling -> Upsampling -> Flatten -> Dense / Fully connected layer -> Softmax
// TO DO: Input -> Conv -> Flatten -> Dense / Fully connected layer
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

// stride = 1; kernel 3x3
// with padding (0 an nicht erreichten Bereichen einfügen); Kernel mit festgelegten Werten über Bild schieben
// width height starten bei 1
std::vector<float> convolution(std::vector<float>&previmage, int width, int height, int bias, float weights[3][3]) {
    std::vector<float> newimage = {};

    // die Zeile u. Spalte die nicht vollständig vom kernel erfasst wird, wird zu Null (erstmal mit ignorieren implementiert)

    // ss = startspalte, sz = startzeile
    // werte im kernel mit previmage multiplizieren, aufsummieren und bias addieren, dann zu neuem image hinzufügen

    for (int ss = 0; ss+3 <= height; ss += 1) { // kernel verschieben vertikal
    for (int sz = 0; sz+3 <= height; sz += 1) { // kernel verschieben horizontal
        // float loc = previmage[ss * width + sz]; 
        float sum = 0;
        for (int y = 0; y < 3; y++) { // kernel durchgehen
        for (int x = 0; x < 3; x++) { // kernel durchgehen
           sum += weights[y][x] * previmage[(ss+y) * width + (sz+x)];
        }
        }
        sum += bias;
        newimage.push_back(sum); // berechnete eintragen in neue
    }
    }
    return newimage;
}


// Width of the new picture after pooling

// we know that the kernel is the size of 3x3
int newWidthConv(int width) {
    if (width >= 2) {
    return width-2;
    }
    return 0;
}


// Hight of the new picture after pooling

int newHeightConv(int height) {
    if (height >= 2) {
    return height-2;
    }
    return 0;
}


// ReLU

float reLu(float x) {
    if (x > 0) {
        return x;
    }
    return 0;
}


// MaxPooling -> größter Wert wird genommen
// stride (verschiebung) so groß wie kernel (also 2)
std::vector<float> maxPooling(std::vector<float>&previmage, int width, int height, int kernel) {
    std::vector<float> newimage = {};

    // die Zeile u. Spalte die nicht vollständig vom kernel erfasst wird, wird ignoriert

    // ss = startspalte, sz = startzeile
     
    for (int ss = 0; ss+kernel <= height; ss += kernel) { // kernel verschieben vertikal
    for (int sz = 0; sz+kernel <= width; sz += kernel) { // kernel verschieben horizontal
        float loc = previmage[ss * width + sz]; // previ or i?
        for (int y = 0; y < kernel; ++y) { // kernel durchgehen
            for (int x = 0; x < kernel; ++x) {

                int index = (ss + y) * width + (sz + x);
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


// Width of the new picture after pooling

int newWidthPooling(int width, int kernel) {
    return width / kernel;
}


// Hight of the new picture after pooling

int newHeightPooling(int height, int kernel) {
    return height / kernel;
}


// Upsampling

// after the Convolution and Pooling the image should be brought back to its original size

#include <vector>

std::vector<std::vector<float>> upsampling(const std::vector<std::vector<float>>& image) {
    if (image.empty() || image[0].empty()) return {};

    int n = image.size(); // Zeilen
    int m = image[0].size(); // Spalten

    std::vector<std::vector<float>> newimage(2 * n, std::vector<float>(2 * m)); // leeres Grid 2mx2n

    for (int y = 0; y < n; ++y) { // für die alten Zeilen
        for (int x = 0; x < m; ++x) { // Element jeder Spalte
            float val = image[y][x];

            newimage[2 * y][2 * x]         = val; // 2 x 2 direkt ins Zielbild
            newimage[2 * y][2 * x + 1]     = val;
            newimage[2 * y + 1][2 * x]     = val;
            newimage[2 * y + 1][2 * x + 1] = val;
        }
    }

    return newimage;
}

// Flatten

std::vector<float> flatten(std::vector<std::vector<float>>& image){
    std::vector<float> newimage = {};
    int n = image.size(); // Zeilen
    int m = image[0].size(); // Elemente erste Zeile

    for (int y = 0; y < n; ++y ) {
    for (int x = 0; x < m; ++x ) {
        newimage.push_back(image[y][x]); // kernel max eintragen
    }
}
return newimage;
}


// Dense / Fully Connected Layer

class DenseLayer {
private:
    // Gewichte
    // Bias

public:
    void forward();
};


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
    std::cout << reLu(0.04);
    // maxpooling();
    //DenseLayer dense;
    //dense.forward();
    return 0;
}