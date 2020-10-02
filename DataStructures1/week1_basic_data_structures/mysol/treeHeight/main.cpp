#include <iostream>
#include <map>
#include <iterator>

using namespace std;

//this is our node class for storing the depth and the parent to a particular node
class Node
{
public:
    //two constructors
    Node(int input);
    Node();
    //deconstructor
    ~Node();

    //our members
    int parent;
};

Node::Node(int input)
{
    parent = input;
}

Node::Node(){
}

Node::~Node()
{
}

//maps for storing the recursice calls to not call them again and again
map<int, int> stored;
map<int, int>::iterator itr;

void input(Node *tree, int noOfNodes)
{
    int takein;
    while (noOfNodes--)
    {
        cin >> takein;
        tree->parent = takein;
        tree += 1;
    }
}

int blastOff(Node *tree , int index){
    //if it is present in the map itself
    itr = stored.find((tree + index)->parent);
    if(itr != stored.end()){
        return 1 + itr->second;
    }
    //if it is the root then that is the end
    if((tree + index)->parent == -1){
        stored.insert(pair<int , int>(index, 1));
        return 1;
    }

    //first of all store the value for the current index
    int store = 1 + blastOff(tree , (tree + index)->parent);

    //map it into the map<stored>
    stored.insert(pair<int ,int>(index , store));

    //return te current stored value
    return store;
}

int calculateDepth(Node *tree , int noOfNodes){

    //keep tracking of the start pointer
    int max = INT32_MIN;
    int keeps;
    for(int i = 0; i < noOfNodes ; i++){
        keeps = blastOff(tree , i);
        if(max < keeps){
            max = keeps;
        }
    }
    return max;
}


int main()
{

    //taking no of nodes fromm the user
    int noOfNodes;
    cin >> noOfNodes;

    //our tree nodes
    Node tree[100000];
    Node *inputptr = tree;
    input(inputptr, noOfNodes);
    
    //printing out the answer
    cout << calculateDepth(inputptr , noOfNodes) << endl;
    return 0;
}
