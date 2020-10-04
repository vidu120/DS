#include<iostream>

using namespace std;

//our node class a
class Node{
    public:
        int data;
        Node *next;
        Node();
        Node(int in);
        ~Node();
};
//constructors for our Node class
Node::Node(){
    next = nullptr;
}
Node::Node(int in){
    data = in;
    next = nullptr;
}

Node::~Node(){}

class stackWithMax{
    private:
        Node *headMaxList;
    public:
        //a different stack to keep track of the max value as we add elements in the stack
        stackWithMax();
        ~stackWithMax();
        void push(int data);
        bool empty();
        int pop();
        int max(); 
};

stackWithMax::stackWithMax(){
     headMaxList = nullptr;
}

stackWithMax::~stackWithMax(){}

void stackWithMax::push(int data){
    if (headMaxList == nullptr){
        headMaxList->data = data;
        headMaxList->next = nullptr;
        return;
    }
    //for keeping track of the max element stacks
    Node maxElem;
    if(headMaxList->data < data){
        maxElem = Node(data);   
    }else{
        maxElem  = Node(headMaxList->data);
    }
    maxElem.next = headMaxList;
    headMaxList = &maxElem;
}

bool stackWithMax::empty(){
    if(headMaxList == nullptr){
        return true;
    }
    return false;
}

int stackWithMax::pop(){
    if(headMaxList == nullptr){
        cout << "Error - No elements in the stack"<< endl;
        return 0;
    }
    Node *top = headMaxList;
    headMaxList = headMaxList->next;
    int popped = top->data;
    free(top);
    return popped;
}

int stackWithMax::max(){
    return pop();
}


int main(){
    int queries;
    cin >> queries;

    string queryType;
    int data;

    stackWithMax max = stackWithMax();

    // while (queries--)
    // {   
    //     cin >> queryType;
    //     if (queryType == "push"){
    //         cin >> data;
    //         max.push(data);
    //     }else if(queryType == "top"){
    //         max.pop();
    //     }else{
    //         cout << max.max() << endl;
    //     }
    // }
    max.push(1);
    cout << max.max() << endl;
    
}