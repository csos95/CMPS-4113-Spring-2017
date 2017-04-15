/*NAME: [Name]		DATE: 09/01/2015
FILE: Prog2.java
PROGRAM SPECIFICATIONS
NARRATIVE DESCRIPTION:
This program is designed to create a Robot to pick up rocks with positions
and weight from an input file. The max number, weight, initial position, 
direction of robot will be asked of the user and default values will be
used if none are entered. The input and output file names will also be asked of
the user. The program will exit if the input file entered does not exist and
will create a new outfile if the file entered does not exist. The robot will  
pick the rock closest to it first and its position changed to position of
the rock it picks until max number or weight is reached. The rocks picked up
will have their information printed to the outfile. The total distance by the
robot, weight and number of rocks picked will also be printed.

INTERFACE:
INTRODUCTION SCREEN: Welcome to NASA2 ROBOT program

EXIT SCREEN: Thanks for using NASA2 ROBOT program

INPUTS:	Robot position, direction, max number, weight options.
        Input, output file names

OUTPUTS: Information of rock(s) picked up, total number, weight, distance 

CONSTANTS:

 */

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.InputMismatchException;
import java.util.NoSuchElementException;
import java.util.Scanner;
import java.util.*;

public class Prog2 {

    /** 
        * Purpose: gets output file from user
        * Requires: none
        * Returns: A BufferedWriter file
    */
    public static BufferedWriter getOutfile()
    {
        Scanner reader = new Scanner(System.in);
        System.out.printf("Enter name/path to outfile: ");
        //Store name for outfile
        String filename = reader.nextLine();
        
        File userFile = new File(filename);
        userFile.setWritable(true);
        FileWriter fw = null;
        BufferedWriter outFile = null;
        try
        {
            //System.out.printf("In try...%n");
            if (!userFile.exists())
            {
                //Attempt to create outfile if it does not exist
                System.out.printf("File doesn't exist. Attempting to create file...%n");
                //if file is created successfully
                if(userFile.createNewFile())
                {
                    System.out.printf("File %s created successfully.%n", userFile.getAbsolutePath());
                }
            }
            fw = new FileWriter(userFile);
            outFile = new BufferedWriter(fw);
        }
        catch(IOException e)
        {
            //Interruption with creating file
            System.out.printf("Problem creating outfile.%n");
            e.printStackTrace();
        }
        finally
        {
            return outFile;
        }
    }
    
    /** 
	* Purpose: asks user for max number, weight, position and direction of robot
	* Requires: none
        * Returns: A Robot with all user's requirements, or default values if no
                   requirements are entered
    */
    public static Robot getRobot()
    {
        //Prompts for coordinate, direction
        System.out.printf("Welcome to NASA2 ROBOT program.%n%n");
        System.out.printf("Do you wish to enter a coordinate for your Robot?%n" +
                "Enter Y for yes or N for no: ");
        Scanner reader = new Scanner(System.in);
        String option = reader.nextLine();
        Position coord = new Position();        
        Robot robot = new Robot();
        
        System.out.printf("%nEnter direction (N, S, E, W) for robot: ");
        String direction = reader.nextLine();
        char dir;
        //prompt for max number of rocks
        System.out.printf("%nDo you wish to enter a max number of rocks to be picked?%n" +
                "Enter Y for yes or N for no: ");
        String maxNumChoice = reader.nextLine();
        int maxNum = robot.MAX_ROCKS;
        
        //prompt for max weight of rocks
        System.out.printf("%nDo you wish to Enter a max weight for rocks to be picked?%n" +
                "Enter Y for yes and N for no: ");
        String maxWeightChoice = reader.nextLine();
        double maxWeight = robot.MAX_WEIGHT;
        
        //check direction entry
        if (!direction.equals(""))
        {
            dir = direction.charAt(0);
            if (dir != 'N' && dir != 'S' && dir != 'E' && dir != 'W')
            {
                System.out.printf("Invalid direction entered. Default direction 'N' in use%n");
            }
        }
        else
        {
            System.out.printf("Default direction N in use%n");
            dir = 'N';
        }        
        
        //check coordinate entry
        if (option.equalsIgnoreCase("Y"))
        {
            //ask for coordinate
            try
            {
                System.out.printf("%nEnter robot's coordinate: ");
                coord = new Position (reader.nextInt(), reader.nextInt());
                System.out.printf("Coordinate in use:  (%d,%d).%n", coord.getX(), coord.getY());
            }
            catch (InputMismatchException e)
            {
                //If input entered is a mismatch
                System.out.printf("Invalid input entered. Default coordinate (0,0) in use.%n");
                //e.printStackTrace();
            }
            catch (NoSuchElementException e)
            {
                //If incomplete input entered; one or more missing
                System.out.printf("Incomplete input enetered.  Default coordinate (0,0) in use.%n");
                //e.printStackTrace();                
            }
        }
        else 
        {
            System.out.printf("%nDefault coordinate (0,0) in use.%n");
        }
        
        //check entry for max number of rocks
        if (!maxNumChoice.equalsIgnoreCase("Y") && !maxNumChoice.equalsIgnoreCase("N"))
        {
            System.out.printf("Invalid option entered. Default max Number to be picked is %d.", maxNum);
        }
        else
        {
            if (maxNumChoice.equalsIgnoreCase("Y"))
            {
                //if yes, ask for max number of rocks
                System.out.printf("%nEnter max Number of Rocks to be picked: ");
                try
                {
                    maxNum = reader.nextInt();
                    System.out.printf("Max number of rocks %d in use.%n", maxNum);
                }
                catch (InputMismatchException e)
                {
                    //If input entered is a mismatch
                    System.out.printf("Invalid option entered. Default number " +
                            "of rocks to be picked %d in use.%n", maxNum);
                    //e.printStackTrace();
                }
                catch (NoSuchElementException e)
                {
                    //If incomplete input entered; one or more missing
                    System.out.printf("Invalid option entered. Default number " +
                            "of rocks to be picked %d in use.%n", maxNum);
                    //e.printStackTrace();                
                }
            }
            else
            {
                System.out.printf("%nDefault number of rocks to be picked %d in use.%n", maxNum);
            }
        }
        //check entry for max weight
        if (!maxWeightChoice.equalsIgnoreCase("Y") && !maxWeightChoice.equalsIgnoreCase("N"))
        {
            System.out.printf("Invalid option entered. Default max Wieght %.2f in use.", maxWeight);
        }
        else
        {
            //if yes, prompt for max weight
            if (maxNumChoice.equalsIgnoreCase("Y"))
            {
                System.out.printf("%nEnter max weight of rocks to be picked: ");
                try
                {
                    maxWeight = reader.nextDouble();
                    System.out.printf("Max weight of rocks %.2f in use.%n", maxWeight);
                }
                catch (InputMismatchException e)
                {
                    //If input entered is a mismatch
                    System.out.printf("Invalid option entered. Default weight " +
                            "of rocks to be picked %.2f in use.", maxWeight);
                    //e.printStackTrace();
                }
                catch (NoSuchElementException e)
                {
                    //If incomplete input entered; one or more missing
                    System.out.printf("Invalid option entered. Default weight " +
                            "of rocks to be picked %.2f in use.", maxWeight);
                    //e.printStackTrace();                
                }
            }
            else
            {
                System.out.printf("Default weight of rocks to be picked %.2f in use.%n", maxWeight);
            }
        }
        //return new robot
        return new Robot(coord, dir, maxNum, maxWeight);
    }
    
    /** 
	* Purpose: gets input file from user
	* Requires: none
        * Returns: A Scanner object, or null if file does not exist
    */
    public static Scanner getInFile()
    {
        Scanner reader = new Scanner(System.in);
        System.out.printf("%nEnter name of input file: ");
        String filename = reader.nextLine();
        
        File userFile = new File(filename);
        userFile.setReadable(true);
        Scanner infile = null;
                        
        try
        {
            if (!userFile.exists())
            {
                System.out.printf("Input File entered does not exist.%n");
            }
            else
            {
                infile = new Scanner(userFile);
            }
            
        }
        catch(IOException e)
        {
            //Interruption with creating file
            System.out.printf("Problem creating infile.%n");
            e.printStackTrace();
        }
        finally
        {
            return infile;
        }
    }
    
    /** 
	* Purpose: reads all rock information from input file and stores in ArrayList
        * Requires: none
        * Returns: An ArrayList of Rocks
    */
    public static ArrayList getRocks(Scanner inFile)
    {
        ArrayList rockList = new ArrayList();
        int count = 0;
        //while there's at least one more line in file
        while (inFile.hasNextLine())
        {         
            //read line
            String line = inFile.nextLine();
            Scanner reader = new Scanner(line);
            count++;
            try
            {
                //read variables from line
                int xVal = reader.nextInt();
                int yVal = reader.nextInt();
                double weight = reader.nextDouble();
                //add new rock to list if all variables are read properly
                rockList.add(new Rock(new Position(xVal, yVal), weight));
            }
            catch (InputMismatchException e)
            {
                //If input entered is a mismatch
                System.out.printf("Invalid input on line %d%n", count);
                //e.printStackTrace();
            }   
            catch (NoSuchElementException e)
            {
                //If incomplete input entered; one or more missing
                System.out.printf("Incomplete input on line %d%n", count);
                //e.printStackTrace();                
            }
        }
     return rockList;   
    }
    
    /** 
	* Purpose: computes distance between two positions
	* Requires: two positions to compute their distance
        * Returns: the computed distance
    */	
    public static double computeDistance(Position pos1, Position pos2)
    {   
        //get y squared
        double y_sqr = pos1.getY() - pos2.getY();
        y_sqr *= y_sqr;
        //get x squared
        double x_sqr = pos1.getX() - pos2.getX();
        x_sqr *= x_sqr;
        //return distance
        return Math.sqrt(y_sqr + x_sqr);
    }
    
    /** 
	* Purpose: gets index of closest rock to a specific position
	* Requires: a position, ArrayList of Rocks 
        * Returns: index of closest Rock to the position passed in
    */
    public static int getClosestRock(Position position, ArrayList rockList)
    {
        //set initial min distance to largest possible distance
        double minDist = Math.sqrt(20000); 
        //get size of ArrayList, set starting point to 0, index of closest rock
        int listSize = rockList.size(), count = 0, minIndex = -1;
        //while we still have more rocks to check 
        while (count < listSize)
        {
            //get next rock
            Rock rock = (Rock) rockList.get(count);
            //get distance from position to rock
            double distance = computeDistance(position, rock.getPosition());
            //update minimum distance if applicable
            if (distance < minDist)
            {
                minIndex = count;
                minDist = distance;
            }
            count++;
        }
        //return index of closest rock
        return minIndex;
    }

    /** 
	* Purpose: prints dashes to output file for formatting
	* Requires: BufferedWriter file
        * Returns: none
    */public static void printDashes(BufferedWriter outFile)
    {
        try
        {
            outFile.newLine();
            outFile.write("---------------------------------------------------"
                    + "---------------");
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
    }

    /** 
	* Purpose: prints header to output file 
	* Requires: BufferedWriter file
        * Returns: none
    */
    public static void printOutfileHead(BufferedWriter output)
    {
        try
        {
            output.write("CMPS 2413 - [Name]"); output.newLine();
            output.write("PROGRAM 2: NASA2 OUTFILE"); output.newLine(); output.newLine();
            output.write("Welcome to  NASA2 ROBOT program");
            output.newLine(); printDashes(output); output.newLine();
            output.write("ROBOT POSITION\tROCK POSITION\tROCK WEIGHT\tDISTANCE TRAVELLED");            
            output.newLine();
        }
        catch(IOException e)
        {
            e.printStackTrace();
            System.out.printf("Writing to outfile interrupted.&n");
        }
        finally{}
                
    }
    
    /** 
	 * Purpose: prints exit information to console and output file
	 * Requires: BufferedWriter output file
         * Returns: none
    */
    public static void printExit(BufferedWriter outFile, int maxN, double maxW, double total)
    {        
        try
        {
            outFile.newLine(); 
            outFile.write("TOTAL ROCKS COLLECTED\t\t\t\t" + maxN); outFile.newLine();
            outFile.write("TOTAL WEIGHT COLLECTED\t\t\t\t" + String.format("%.2f", maxW));
            outFile.newLine(); 
            outFile.write("TOTAL DISTANCE \t\t\t\t\t" + String.format("%.2f", total));
            printDashes(outFile);
            outFile.newLine(); outFile.newLine();
            outFile.write("Thanks for using  NASA2 ROBOT program.");
            outFile.close();
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
        System.out.printf("%nThanks for using  NASA2 ROBOT program.%n");
    }
    
    /** 
	* Purpose: prints information after a rock is picked
	* Requires: BufferedWriter file, robot's position (before picking rock),
                    Rock's position
        * Returns: none
    */
    public static void printInfo(BufferedWriter output, Position robPos, Rock rock)
    {
        try
        {
            String rob = posToString(robPos);
            String robT = rob.length() > 7 ? "\t" : "\t\t";
            String roc = posToString(rock.getPosition());
            String rocT = roc.length() > 7 ? "\t" : "\t\t";
            output.write(posToString(robPos) + robT + posToString(rock.getPosition()) 
                    + rocT + String.format("%.2f", rock.getWeight()) + "\t\t" +
                    String.format("%.2f", computeDistance(robPos, rock.getPosition())));
            output.newLine();
        }
        catch (IOException e){
            e.printStackTrace();
        }
    }
    
    /** 
	* Purpose: converts a position to String (x,y) format and returns it 
	* Requires: A Position
        * Returns: (x,y) format for Position
    */
    public static String posToString(Position pos)
    {
        return "(" + pos.getX() + "," + pos.getY() + ")";
    }
    
    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        
        //Get robot info from user
        Robot marsRobot = getRobot();
        
        //Get input file
        Scanner inFile = getInFile();
        //Get output file
        BufferedWriter outFile = getOutfile();
        printOutfileHead(outFile);
        
        if (inFile != null)
        {
            //get all rocks and store in array list
            ArrayList rockList = getRocks(inFile);
            //activate robot
            marsRobot.activate();
            //compute all distances for robot to rock(s), get index of closest rock
            int index = getClosestRock(marsRobot.getPosition(), rockList);
            //while max rocks isn't reached and rockList isn't empty
            while(!marsRobot.maxRocksReached() && !rockList.isEmpty())
            {
                //get closest rock
                Rock rock = (Rock) rockList.remove(index);
                //get robot's position
                Position robPos = marsRobot.getPosition();
                //pick up rock
                if (marsRobot.pickRock(rock))
                {
                    //print info                        
                    printInfo(outFile, robPos, rock);
                }
                //get index of next closest rock
                index = getClosestRock(marsRobot.getPosition(), rockList);
            }
            marsRobot.deactivate();
        }          
        printExit(outFile, marsRobot.getNumRocks(), marsRobot.getWieightRocks(),
                marsRobot.getTotalDistance());              
        
    }
    
}
