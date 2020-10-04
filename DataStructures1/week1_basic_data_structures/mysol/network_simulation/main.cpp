#include<iostream>
#include<queue>

using namespace std;


class packet
{
public:
    int timeOfArrival;
    int timeForProcessing;
    packet(int a , int b);
    ~packet();
};

packet::packet(int a , int b)
{
    timeOfArrival = a;
    timeForProcessing = b;
}
packet::~packet()
{
}

//This is our packets queue
queue<packet> packets;
queue<int> packetsIndex;


void findTimeOfProccessing(int *ptr , int noOfPackets , int bufferSize){
    int timeOfArrival , timeForProccessing , timeSpent;
    bool firstOne = false;
    for(int i = 0;i < noOfPackets ; i++){
        //take details about the packet from the cin stream
        cin >> timeOfArrival;
        cin >> timeForProccessing;

        //for the counter to work properly
        if(!firstOne){
            timeSpent = timeOfArrival;
            firstOne = true;
        }

        if(packets.size() != bufferSize){
            packets.push(packet(timeOfArrival , timeForProccessing));
            packetsIndex.push(i);
        }else{
            if(packets.front().timeForProcessing + timeSpent <= timeOfArrival){

                //first update the subsequent packet's time for processing
                *(ptr + packetsIndex.front()) = timeSpent;
                packetsIndex.pop();
                 
                //update the time spent till now
                timeSpent += packets.front().timeForProcessing;
                packets.pop();

                packets.push(packet(timeOfArrival, timeForProccessing));
                packetsIndex.push(i);
            }else{
                *(ptr+ i) = -1;
            }
        }
    }

    while(packets.size() != 0){
        if(timeSpent < packets.front().timeOfArrival){
            timeSpent = packets.front().timeOfArrival;
        }
        //first update the subsequent packet's time for processing
        *(ptr + packetsIndex.front()) = timeSpent;
        packetsIndex.pop();
            
        //update the time spent till now
        timeSpent += packets.front().timeForProcessing;
        packets.pop();
    }

}


int main(){

    //taking our buffer size from the user
    int bufferSize , noOfPackets;
    cin >> bufferSize;
    cin >> noOfPackets;

    //time for processing or for declination
    int *ptr = (int *)(malloc(noOfPackets * sizeof(int)));

    //pass the above pointer for storing the values associated with each packet
    findTimeOfProccessing(ptr , noOfPackets , bufferSize);

    for(int i = 0; i < noOfPackets ; i++){
        cout << *(ptr + i) << endl;
    }

    return 0;
}   