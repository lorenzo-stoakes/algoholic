#include <iostream>
#include <vector>

using namespace std;

const int N = 1e4;

vector<int>
isort(vector<int> ns)
{
    for(int i = 1; i < ns.size(); i++) {
        int key = ns[i];
        int j = i-1;
        for(; j >= 0 && key < ns[j]; j--) {
            ns[j+1] = ns[j];
        }
        ns[j+1] = key;
    }

    return ns;
}

int
main(void)
{
    vector<int> ns;

    for(int i = 0; i < N; i++) {
        ns.push_back(N-i);
    }

    ns = isort(ns);

    int expected = 1;
    for(vector<int>::const_iterator it = ns.begin(); it != ns.end(); it++) {
        int val = *it;

        if(val != expected) {
            cerr << "Value is " << val << ", expected " << expected << "." << endl;
            return EXIT_FAILURE;
        }
        expected++;
    }

    return EXIT_SUCCESS;
}
