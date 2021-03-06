/**************************************************************
//                   Project #1 Mutant Flatworld Explorers
//                    Name: Da Dong
//                    Data Structures Date: 02/04/2015
//***************************************************************
//      This program is a pracitce to use vector class instead of 2D array.
//      It is a simple Robot simulation that creates a board
//          and moves the robot around using a sequence of commands.
//          When a robot runs off the board it is called LOST
//          and the position is block for future robots
//          so they will not jump off at the same location.
//      The first line of input is the upper - right coordinates of the rectangular world,
//        the lower - left coordinates are assumed to be 0, 0.
//
//*****************************************************************

//*******************************************************************
//                      ObjectName::vector
//                    Parameters: none
//                     Complexity: none
//********************************************************************/
#include<iostream>
#include<vector>
#include<cstring>

using namespace std;
int main()
{
    int MaxX,MaxY,x,y, len,z;
    int count = 0;
    char Dir;
    //declare a character in order to record the direction before lost
    char L = ' ';
    //declare a series of instructions
    char str[100];

    //set a map
    vector<int>v1(51);
    //set a vector to store the loction that the robot has been previously lost
    vector< pair <int, int> > v;
    
    //input the Max size of the rectangular
    cin >> MaxX >> MaxY;

    //using for loops to set the value into the vectors
    for (int i = MaxY; i >= 0; i--)
    {
        v1[0] = i;
        for (int c = 1; c <= MaxX; c++)
            v1[c] = c;
    }

    while (cin >> x >> y >> Dir >> str)
    {
        //using c-string to know the length of the instructions
        len = strlen(str);

        //using while loop to analyze the instructions
        z = 0;
        while (z < len)
        {
            if (str[z] == 'R')
            {
                if (Dir == 'E')
                    Dir = 'S';
                else if (Dir == 'S')
                    Dir = 'W';
                else if (Dir == 'W')
                    Dir = 'N';
                else
                    Dir = 'E';
            }
            else if (str[z] == 'L')
            {
                if (Dir == 'E')
                    Dir = 'N';
                else if (Dir == 'S')
                    Dir = 'E';
                else if (Dir == 'W')
                    Dir = 'S';
                else
                    Dir = 'W';
            }
            else
            {
                if (Dir == 'E')
                    x++;
                else if (Dir == 'S')
                    y--;
                else if (Dir == 'W')
                    x--;
                else
                    y++;
            }

            //avoid losting at the same point.
            if (count > 0)
            {
                for (int n = 0; n < count; n++)
                    if (pair<int, int>(x, y) == v[n] && Dir == L)
                        z ++;
            }
            //use if statements to identify if the robot is lost.
            if (x > MaxX && y > MaxY)
            {
                cout << MaxX << " " << MaxY << " " << Dir;
                v.push_back(pair<int, int>(MaxX, MaxY));
                count++;
                z = len;
                L = Dir;
            }
            else if (x > MaxX && y < 0)
            {
                cout << MaxX << " " << 0 << " " << Dir;
                v.push_back(pair<int, int>(MaxX, 0));
                count++;
                z = len;
                L = Dir;
            }
            else if (y < 0 && x < 0)
            {
                cout << 0 << " " << 0 << " " << Dir;
                v.push_back(pair<int, int>(0, 0));
                count++;
                z = len;
                L = Dir;
            }
            else if (y < 0)
            {
                cout << x << " " << 0 << " " << Dir;
                v.push_back(pair<int, int>(x, 0));
                count++;
                z = len;
                L = Dir;
            }
            else if (y > MaxY)
            {
                cout << x << " " << MaxY << " " << Dir;
                v.push_back(pair<int, int>(x, MaxY));
                count++;
                z = len;
                L = Dir;
            }
            else if (x > MaxX)
            {
                cout << MaxX << " " << y << " " << Dir;
                v.push_back(pair<int, int>(MaxX, y));
                count++;
                z = len;
                L = Dir;
            }
            else if (x < 0)
            {
                cout << 0 << " " << y << " " << Dir;
                v.push_back(pair<int, int>(0, y));
                count++;
                z = len;
                L = Dir;
            }
            if (x < 0 || x>MaxX || y <0 || y>MaxY)
                cout << " LOST" << '\n';
            z++;
        }
        if (x >= 0 && x <= MaxX && y >= 0 && y <= MaxY)
            cout << x << " " << y << " " << Dir << '\n';
    }
    return 0;
}