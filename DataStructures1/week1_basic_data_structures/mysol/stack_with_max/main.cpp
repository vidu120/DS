#include<iostream>
#include<stack>

using namespace std;



//new stack with max element functionality inherited from the stack<int> class
class maxImplementedStack : public stack<int>{
    public:
        int Max();
        void Push(int data);
        void Pop();
};

int maxImplementedStack::Max(){
    return this->top();
}
void maxImplementedStack::Pop(){
    this->pop();
}

void maxImplementedStack::Push(int data){
    if(this->empty()){
        this->push(data);
    }else{
        if (this->top() < data){
            this->push(data);
        }else{
            this->push(this->top());
        }
    }
}



int main(){
    int queries;
    cin >> queries;

    string queryType;
    int data;

    //this is the stack of the max elements
    maxImplementedStack maxElemList;

    while(queries--){
        cin >> queryType;
        if(queryType == "push"){
            cin >> data;
            maxElemList.Push(data);
        }else if(queryType == "pop"){
            maxElemList.Pop();
        }else{
            cout << maxElemList.Max() << endl;
        }
    }
}