/***************************************
	Nathan Durst
	October 5, 2015
	Contemporary Programming Languages
	This file uses a CoinSorter class to
	place individual coin values in an ArrayList
	Then after it calls on CoinSorter methods,
	it prints the number of rolls and coins to a file
****************************************/

import java.util.Random;
import java.util.Scanner;
import javax.swing.JOptionPane;
import java.io.*;
import java.util.ArrayList;

public class Source {

	public static void main(String[] args)
			throws Exception, IOException {
		//Calls function to return a file name
		// and prints coin values to the file that
		// will later be read from
		String inFile = getFileName();
		String outFile = getFileName();
		printValues(inFile);

		//Declares main variables including file stream variables
		// and ArrayList variable
		PrintWriter out = new PrintWriter(new FileWriter(outFile));
		Scanner scanner = new Scanner(new File(inFile));
		CoinSorter coins = new CoinSorter();
		ArrayList<Integer> list = new ArrayList<Integer>();

		try
		{	//prints heading to file
			showIntro(out);

			//executes while not end of file
			while (scanner.hasNext())
			{
				//stores first integer in "val"
				// then loops while integer is not -1
				int val = scanner.nextInt();
				while (val != -1)
				{
					//appends integer to end of list
					// and retrieves next integer from file
					list.add(val);
					val = scanner.nextInt();
				}
				//calls function to copy list to array
				// and sorts in ascending order
				coins.serialize(list);
				//Divides coins into appropriate rolls
				coins.fillRolls();
				//Counts number of rolls and spare coins
				coins.ComputeTotal();
				//converts CoinSorter obj to a string
				// so it can be printed to a file
				coins.toString();
				//prints CoinSorter obj to file
				printTotals(coins, out);
				//clears array list so it can be reinitialized
				list.clear();
			}
		}
		catch(Exception e){
			System.out.println(inFile +" does not exist!");
		}
		//print closing message and close files
		showClosing(out);
		scanner.close();
		out.close();
	}

	//PURPOSE: Prints heading to output file
	//REQUIRES: PrintWriter variable for printing
	//RETURNS: nothing
	public static void showIntro(PrintWriter out)
	{
		out.println("Nathan Durst");
		out.println("CoinSorter output\n");

		out.println("Welcome to the CoinSorter 3000");
		out.println("Please insert some coins:\n");
	}

	//PURPOSE: Shows dialog box to prompt user for filenames
	//REQUIRES: nothing
	//RETURNS: String representing the name of the file
	public static String getFileName()
	{
		//prompts user for name of file
		String myFile = JOptionPane.showInputDialog("Name the file"
						+ " you would like to use");
		return myFile;
	}

	//PURPOSE: Randomly prints integer coin values 1, 5, 10, and 25
	//REQUIRES: String representing the name of file
	//RETURNS: nothing
	public static void printValues(String inFile) throws Exception
	{
		//stopper used for modular arithmetic
		// a -1 is printed after ever 160 integers
		int stopper = 0;
		Random random = new Random();
		PrintWriter out = new PrintWriter(new FileWriter(inFile));
		try
		{	//40 rows
			for (int i = 0; i<40; i++)
			{	//20 integers per row (800 total integers)
				for (int j = 0; j<20; j++)
				{
					//creates an integer 1-4
					int coinVal = random.nextInt(4);
					//each integer represents a coin value
					// 1, 5, 10, or 25
					if (coinVal == 1)
						out.print(1);
					else if (coinVal == 2)
						out.print(5);
					else if (coinVal == 3)
						out.print(10);
					else
						out.print(25);
					out.print(" ");
					stopper ++;
					//prints a -1 after 160 integers
					if (stopper % 160 == 0)
						out.print(-1);
				}
				out.print('\n');
			}
		}
		catch (Exception e) {
			System.out.println(inFile + "does not exist!");
		}
		out.close();
	}

	//PURPOSE: converst CoinSorter obj to a string and prints it
	//REQUIRES: CoinSorter obj and PrintWriter obj
	//RETURNS: nothing
	public static void printTotals(CoinSorter coins, PrintWriter out)
	{
		String output;
		//converts obj to a string
		output = coins.toString();
		out.println(output);
		out.println("Yum! Now enter more coins if you have them!\n");
	}

	//PURPOSE:prints closing message
	//REQUIRES: PrintWriter obj
	//RETURNS: nothing
	public static void showClosing(PrintWriter out)
	{
		out.println("Looks like you ran out of money!");
		out.println("Thank you for using the CoinSorter 3000");
		out.println("Have a great day!!");
	}
}
