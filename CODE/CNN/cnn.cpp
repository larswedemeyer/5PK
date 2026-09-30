#include <iostream>
#include <vector> // WICHTIG: Stellt std::vector bereit; vector can shrink and grow in size, otherwise behaves like an array

// CNN

// Input -> Conv -> ReLU -> Pooling -> Flatten -> Dense / Fully connected layer -> Softmax
// Backtracking :C

/* Fib

using namespace std;

void fibonacci() {
    int n;
    cout << "Enter the number of terms: ";
    cin >> n;

    long long t1 = 0, t2 = 1, nextTerm = 0; // long long datentyp; danach variablen des types

    cout << "Fibonacci Series: ";

    for (int i = 1; i <= n; ++i) {
        if(i == 1) {
            cout << t1 << ", ";
            continue;
        }
        if(i == 2) {
            cout << t2 << ", ";
            continue;
        }
        
        nextTerm = t1 + t2;
        t1 = t2;
        t2 = nextTerm;
        
        cout << nextTerm << ", ";
    }
    return 0;
}
*/

// Input
std::vector<float> image;

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

// Pooling
void pooling() {
    // Pooling
}

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

void softmax(){
    // Softmax
}

int main() {
    relu();
    pooling();
    DenseLayer dense;
    dense.forward();
    return 0;
}