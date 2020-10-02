#include <iostream>
#include <stack>
##include <deque>
using namespace std;

//bracket class
class bracket
{
public:
    bracket(int pos, char character);
    ~bracket();
    int pos;
    char character;
};

bracket::bracket(int where, char what)
{
    pos = where;
    character = what;
}

//deconstructor for our bracket
bracket::~bracket() {}

//type of brackets to look at
char brackets[6] = {'{', '[', '(', ')', ']', '}'};

int main()
{
    //Here , we take the input from user
    string input;
    getline(cin, input);

    //our stack of bracket
    stack<bracket> keepTrack;

    for (int i = 0; i < input.size(); i++)
    {
        for (int j = 0; j < 6; j++)
        {
            //checking if it mathces the input data matches any of the brackets
            if (input[i] == brackets[j])
            {

                //is it an opeing bracket
                if (j < 3)
                {
                    keepTrack.push(bracket(i + 1, input[i]));
                }
                else
                {
                    //is it a closing one
                    if (keepTrack.empty())
                    {
                        cout << i + 1 << endl;
                        exit(0);
                    }
                    else
                    {
                        if (keepTrack.top().character != brackets[5 - j])
                        {
                            cout << i + 1 << endl;
                            exit(0);
                        }
                        keepTrack.pop();
                    }
                }
            }
        }
    }

    if (!keepTrack.empty())
    {
        cout << keepTrack.top().pos << endl;
    }
    else
    {
        cout << "Success" << endl;
    }

    return 0;
}
