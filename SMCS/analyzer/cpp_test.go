package analyzer

import (
	"testing"
	"fmt"
)

func Test_CPP_Hello(t *testing.T) {
	program := `#include <iostream>
				using namespace std;
				int main() {
					cout << "Hello, World!" << endl;
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 6
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 6
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_Add(t *testing.T) {
	program := `#include <iostream>
				using namespace std;
				int main() {
					int x = 1;
					int y = 2;
					int z = x + y;
					cout << z << endl;
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 9
		case "Lines of Documentation":
			expected = 0
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 9
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_Document(t *testing.T) {
	program := `/*This is a super awesome program that does cool stuff.
				By: [user]*/
				#include <iostream>
				using namespace std;
				int main() {
					int x = 1;
					int y = 2;
					int z = x + y;
					cout << z << endl;
					//useless comment
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 9
		case "Lines of Documentation":
			expected = 3
		case "Blank Lines":
			expected = 0
		case "Total Lines":
			expected = 12
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_P2a(t *testing.T) {
	program := `// CMPS 1044 - Tina Johnson
				// Program 2 - Catering Service
				// 2/20/17
				// This program asks the user for the amount of people and prices for a catering
				// service. It asks the user for number of adults and children that are coming
				// , the prices for adult meals, dessert price, room fee, tax rate, and deposit.
				// It then runs several calculations to determine the totals for all meals,
				// dessert total, tax total, subtotal, and overall total. Lastly, it prints
				// out the results for the user to view.

				#include<iostream>
				#include<iomanip>
				#include<fstream>

				using namespace std;

				int main()
				{

					// User entered variables
					int NumAdult, NumChild;
					// More user entered variables
					double AdultMeal, ChildMeal, DessertCost, RoomFee, TaxRate, Deposit;
					//Variables for calculations
					double AdultMealAmt, ChildMealAmt, DessertTotal, TaxTotal, Subtotal, Total;

					ofstream outfile;
					outfile.open("output.txt");

					//Asks user to enter variables
					cout << "How many adults will there be?: ";
					cin >> NumAdult;
					cout << "How many children will there be?: ";
					cin >> NumChild;
					cout << "Cost for adult meal? : ";
					cin >> AdultMeal;
					cout << "Cost for one dessert?: ";
					cin >> DessertCost;
					cout << "Cost for room fee?: ";
					cin >> RoomFee;
					cout << "Tax Rate? : ";
					cin >> TaxRate;
					cout << "Total deposit?: ";
					cin >> Deposit;

					cout << "\n\n";


					outfile << "CMPS 1044 - Tina Johnson\n";
					outfile << "Program 2 - Catering Service\n";
					outfile << "2/20/17\n";
					outfile << "This program asks the user for the amount of people and prices for a catering\n";
					outfile << "service. It asks the user for number of adults and children that are coming\n";
					outfile << ", the prices for adult meals, dessert price, room fee, tax rate, and deposit.\n";
					outfile << "It then runs several calculations to determine the totals for all meals,\n";
					outfile << "dessert total, tax total, subtotal, and overall total. Lastly, it prints\n";
					outfile << "out the results for the user to view.\n\n";

					//Caculations used for displayed results
					ChildMeal = AdultMeal * .6;
					AdultMealAmt = AdultMeal * NumAdult;
					ChildMealAmt = ChildMeal * NumChild;
					DessertTotal = (NumAdult + NumChild) * DessertCost;
					TaxTotal = (AdultMealAmt + ChildMealAmt + DessertTotal) * TaxRate;
					Subtotal = AdultMealAmt + ChildMealAmt + DessertTotal + TaxTotal + RoomFee;
					Total = Subtotal - Deposit;

					//Redisplay of user entered numbers
					outfile  << "    **** Catering Service **** " << endl;
					outfile  << "Number of adults:" << setw(13) << NumAdult << endl;
					outfile  << "Number of children:" << setw(11) << NumChild << endl;
					outfile  << fixed << setprecision(2);
					outfile  << "Cost per adult meal:" << setw(4) << "$" << setw(7) << AdultMeal << endl;
					outfile  << "Cost per child meal:" << setw(4) << "$" << setw(7) << ChildMeal << endl;
					outfile  << "Cost per dessert:" << setw(7) << "$" << setw(7) << DessertCost << endl;
					outfile  << "Room fee:" <<setw(15) << "$" << setw(7) << RoomFee << endl;
					outfile  << "Tax rate:" << setw(15)<< "$" << setw(7) << TaxRate << "\n\n";

					// Displays calculations
					outfile  << "Total for adult meals:"<< setw(2) << "$" << setw(7) << AdultMealAmt << endl;
					outfile  << "Total for child meals:"<< setw(2)<< "$" << setw(7) << ChildMealAmt << endl;
					outfile  << "Total for dessert:" << setw(6) << "$" << setw(7) << DessertTotal << endl;
					outfile  << "Tax amount:" << setw(13) << "$" << setw(7) << TaxTotal << endl;
					outfile  << "Room Fee:"<< setw(15) << "$" << setw(7) << RoomFee << "\n\n";
					// More calculation display
					outfile  << "Subtotal:" << setw(15) <<"$" << setw(7) << Subtotal << endl;
					outfile  << "Less deposit: " << setw(10) << "$" << setw(7) << Deposit << endl;
					outfile  << "Balance Due: "<< setw(11) << "$"  << setw(7) << Total << endl;

					outfile.close();
					system("pause");
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 63
		case "Lines of Documentation":
			expected = 17
		case "Blank Lines":
			expected = 13
		case "Total Lines":
			expected = 93
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_P3a(t *testing.T) {
	program := `/*
				CMPS 1044 Johnson
				02/28/17
				Project 3-1 Roman Numerals
				The user will enter a number between 1 and 10 and the program will
				print out a numeral # using switch statements.
				*/

				#include<iostream>
				#include<fstream>
				#include<cstring>

				using namespace std;

				int main()
				{
					int digit, count = 1;
					//bool quit = false;

					ofstream outfile;
					outfile.open("TimesNewRoman.txt"); //Opens TimesNewRoman.txt textfile.

					outfile << "CMPS 1044 Johnson\n";
					outfile << "02 / 28 / 17\n";
					outfile << "Project 3 - 1 Roman Numerals\n";
					outfile << "The user will enter a number between 1 and 10 and the program will ";
					outfile << "print out a numeral # using switch statements.\n\n";

					while (count < 12)
					{

						cout << "Please enter an integer between 1-10: \n";

						cin >> digit;

						outfile << "Trial " << count << endl;

						switch (digit)
						{
						case 1:
							outfile << 'I';

							break;
						case 2:
							outfile << "II";

							break;
						case 3:
							outfile << "III";

							break;
						case 4:
							outfile << "IV";

							break;
						case 5:
							outfile << 'V';

							break;
						case 6:
							outfile << "VI";

							break;
						case 7:
							outfile << "VII";

							break;
						case 8:
							outfile << "VIII";

							break;
						case 9:
							outfile << "IX";

							break;
						case 10:
							outfile << 'X';

							break;
						default:
							outfile << "Please enter a valid choice.\n";
							break;
						}

						outfile << endl << endl;

						count++;
						//goto Endwhile;
					}



						outfile.close(); // closes TimesNewRoman.txt

					system("pause"); // Please ignore the Red Squiggly, the program will compile and run fine.
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 62
		case "Lines of Documentation":
			expected = 12
		case "Blank Lines":
			expected = 26
		case "Total Lines":
			expected = 97
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_P4a(t *testing.T) {
	program := `// CMPS 1044 - Tina Johnson
				// Program 4 - Salary Mutiplier
				// 3 / 27 / 17
				// This program is designed to calulate how much money someone will get when their
				// daily salary of one penny per day was doubled each day.
				// The program will ask the user how many days they worked. The program will then
				// put the user entered number into a for loop where it will double .01 as
				// many times as the user entered number. The program will then display, in a table
				// format, how much they earned for each day as well as the sum of the
				// salaries for all of the days.

				#include<iostream>
				#include<fstream>
				#include<iomanip>

				using namespace std;

				int main()
				{
					ofstream outfile;
					outfile.open("output.txt");

					// Start variable used as counter
					// Days variable is a user entered variable
					int days, start;
					double penny = .01;
					double sum = 0;

					outfile << "CMPS 1044 - Tina Johnson\n";
					outfile << "Program 4 - Salary Mutiplier\n";
					outfile << "3 / 27 / 17\n";
					outfile << "This program is designed to calulate how much money someone will get when their\n";
					outfile << "daily salary of one penny per day was doubled each day.\n";
					outfile << "The program will ask the user how many days they worked. The program will then \n";
					outfile << "put the user entered number into a for loop where it will double .01 as\n";
					outfile << "many times as the user entered number. The program will then display, in a table\n";
					outfile << "format, how much they earned for each day as well as the sum of the\n";
					outfile << "salaries for all of the days.\n\n";

					cout << "Enter the number of days that you worked at the job : ";
					cin >> days;
					cout << endl;

					// Checks if the user entered a valid number that is >= 1
					while (days <= 0)
					{
						cout << "Please enter a number that is greater than or equal to 1 : ";
						cin >> days;
						outfile << endl;
					}

					// Header of Table
					outfile << "   Day" << setw(20) << "Pay\n";
					outfile << "-----------------------------------\n";

					// Primary loop used for salary multiplication calculations
					for (start = 1; start <= days; start++)
						{
							outfile << fixed << setprecision(2);

							outfile << setw(5) << start << setw(23) << penny << endl;
							sum += penny;
							penny *= 2;
						}
					outfile << endl;
					// Sum of pay for each day in table
					outfile << setw(7) << "Total" << setw(10) << "$" << setw(11) << sum << "\n\n";

					outfile.close();
					system("pause");
					return 0;
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 45
		case "Lines of Documentation":
			expected = 16
		case "Blank Lines":
			expected = 11
		case "Total Lines":
			expected = 72
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}

func Test_CPP_Mutant(t *testing.T) {
	program := `/**************************************************************
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
				/****
				adkasdbkas*/
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
				}`

	a := NewAnalyzer()
	var expected int

	analysis := a.Analyze("c", program, []string{"Lines of Code", "Lines of Documentation", "Blank Lines", "Total Lines", "Number of Functions"})
	for _, result := range analysis.Results {
		switch result.Metric {
		case "Lines of Code":
			expected = 130
		case "Lines of Documentation":
			expected = 32
		case "Blank Lines":
			expected = 7
		case "Total Lines":
			expected = 169
		case "Number of Functions":
			expected = 1
		}
		if expected != result.Value {
			t.Error(fmt.Sprintf("Expected %d %s, got %d %s", expected, result.Metric, result.Value, result.Metric))
		}
	}
}