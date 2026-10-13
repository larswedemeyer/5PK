#include <iostream>
#include <vector> // WICHTIG: Stellt vector bereit; vector can shrink and grow in size, otherwise behaves like an array
#include <cmath>
#include <algorithm>
using namespace std;
// CNN

// Input -> Conv -> ReLU -> Pooling -> Upsampling -> Flatten -> Dense / Fully connected layer -> Softmax
// TO DO: Input
// Backtracking :C

// Input

vector<vector<float>> input(vector<vector<float>> a) {
    return a;
}

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

vector<vector<float>> convolution(vector<vector<float>>& previmage,int bias,int kernel,float weights[3][3]) {
    if (previmage.empty() || previmage[0].empty()) return {};
    
    int n = previmage.size();       // Zeilen
    int m = previmage[0].size();    // Spalten
    int outeredge = kernel / 2;     // 0er Spalten
    
    vector<vector<float>> newimage(n, vector<float>(m));
    for (int y = 0; y < n; ++y) {
        for (int x = 0; x < m; ++x) {
            float sum = 0;
            for (int i = 0; i < kernel; ++i) {
                for (int j = 0; j < kernel; ++j) {
                    int imageY = y + i - outeredge;
                    int imageX = x + j - outeredge;
                    // Pos. innerhalb freigegebenem Bereich
                    if (imageY >= 0 && imageY < n && imageX >= 0 && imageX < m) {
                        sum += weights[i][j] * previmage[imageY][imageX];
                    }
                }
            }
            newimage[y][x] = sum + bias;
        }
    }
    return newimage;
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
vector<vector<float>> maxPooling(vector<vector<float>>&previmage, int kernel) {

    if (previmage.empty() || previmage[0].empty()) return {};

    int n = previmage.size(); // Zeilen
    int m = previmage[0].size(); // Spalten

    vector<vector<float>> newimage(newHeightPooling(n, kernel), vector<float>(newWidthPooling(m, kernel)));

    for (int y = 0; y+kernel <= n; y += kernel) { // für die alten Zeilen
        for (int x = 0; x+kernel <= m; x += kernel) { // Element jeder Spalte
            float highestValue = max({previmage[y][x], previmage[y+1][x], previmage[y][x+1], previmage[y+1][x+1]});
            newimage[y/kernel][x/kernel] = highestValue; // an entsprechende Stelle nach Pooling
                }
        }
    return newimage;
}


// Width of the new picture after pooling

// at first kernel only 2 wide
int newWidthPooling(int width, int kernel) {
    return (width+1) / kernel;
}


// Hight of the new picture after pooling

int newHeightPooling(int height, int kernel) {
    return (height+1) / kernel;
}


// Upsampling

// after the Convolution and Pooling the image should be brought back to its original size

vector<vector<float>> upsampling(const vector<vector<float>>& image) {
    if (image.empty() || image[0].empty()) return {};

    int n = image.size(); // Zeilen
    int m = image[0].size(); // Spalten

    vector<vector<float>> newimage(2 * n, vector<float>(2 * m)); // leeres Grid 2mx2n, durch starre Speicherplatz reserviert

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

vector<float> flatten(vector<vector<float>>& image){
    vector<float> newimage = {};
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
// alle Eingabewerte des 1D Vektors mit dem Gewicht eines Neurons multiplizieren, aufsummieren, bias addieren

class DenseLayer {
private:
    vector<vector<float>> Gewichte;
    vector<float> Bias;

public:
    vector<float> forward(const vector<float>& input) {
        
        float n = Bias.size();

        vector<float> output(n, 0.0f);

        for (int j = 0; j < n; j++) {
            output[j] = Bias[j];

            for (int i = 0; i < input.size(); i++) {
                output[j] += input[i] * Gewichte[j][i];
            }
        }

        return output;
    }
};


//Softmax

// softmax(zi) = exp(zi) sum j=1 to d exp zj ; 1 ≤ i ≤ d 
float softmax(const vector<float>& zahlen, int K, int i){ // & pointer mäßig
    
    if (zahlen.empty() || i < 0 || i >= K || i >= zahlen.size()) { // abfangen
        return 0.0f; // f = float
    }
    
    float numerator = exp(zahlen[i]);

    float denominator = 0.0f;

    for (int j = 0; j < K; ++j ){ // K die Klassen zwischen denen Unterschieden wird (Ziel, nicht Ziel)
        denominator += exp(zahlen[j]);
    }
    return numerator / denominator;
}

int main() {
    cout << reLu(0.04);
    return 0;
}