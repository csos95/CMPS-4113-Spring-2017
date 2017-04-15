
/*NAME: [Name]			DATE: 09/10/2015
FILE: Prog1.java
PROGRAM SPECIFICATIONS
NARRATIVE DESCRIPTION:
This program is designed to read sales and receipt Records from the console
and make sales from the receipt Record(s) with a 20% markup in a LIFO order,
Hence receipt Records received later are sold first. Appropriate information
pertaining to each sale and receipt Record are printed to both the console 
and the outfile provided. A new file is created if the specified outfile does
not already exist.

INTERFACE:
INTRODUCTION SCREEN: Welcome to Widget store program

EXIT SCREEN: Thanks for using Widget store program

INPUTS:	Name of outfile, sales and/or receipt Records

OUTPUTS: Information pertaining to each sale(s) and/or Record(s)

CONSTANTS:

 */

import java.util.*;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;

public class Prog1 {
    /** 
	 * Purpose: Prompts user to start entering records
	 * Requires: none
         * Returns: none
    */
    public static void promptForRecord()
    {
        System.out.printf("%nEnter Records below \"Trans-type Date/time Qty Price\"%n"
                + "Push Return twice to stop:%n");
    }
    
    /** 
	 * Purpose: gets output file from user
	 * Requires: none
         * Returns: a BufferedWriter file
    */
    public static BufferedWriter getOutfile()
    {
        Scanner reader = new Scanner(System.in);
        System.out.printf("Welcome to Widget store program.%n%n");
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
	 * Purpose: reads records from console and returns it
	 * Requires: none
         * Returns: record read
    */
    public static Record readRecord()
    {
        Scanner reader = new Scanner(System.in);
        System.out.printf("%n");
        char tType; String dateTime; int qty; double cost;
        String line = reader.nextLine();
        Scanner ls = new Scanner(line);
        //If user enteres data
        if (!line.equals(""))
        {
            try
            {              
                String type = ls.next();            
                tType = type.charAt(0);
                dateTime = ls.next();            
                qty = ls.nextInt(); 
                if (tType == 'r')
                {
                    //get cost and return new receipt Record
                    cost = ls.nextDouble();
                    return (new Record (dateTime, qty, cost, tType));
                }
                else
                {
                    //return record without cost
                    return (new Record (dateTime, qty, tType));
                }
                              
            }
            catch (InputMismatchException e)
            {
                //If input entered is a mismatch
                System.out.printf("Invalid input entered.%n");
                e.printStackTrace();
            }
            catch (NoSuchElementException e)
            {
                //If incomplete input entered; one or more missing
                System.out.printf("Incomplete input enetered for Record.%n");
                e.printStackTrace();                
            }
            
            return null;
            
        }
        else
            return null;
        
    }
    
    /** 
	 * Purpose: prints message for invalid symbol 
	 * Requires: a BufferedWriter file
         * Returns: none
    */
    public static void printInvalidSymbol(BufferedWriter outFile)
    {
        System.out.printf("INVALID SYMBOL ENTERED.%n");
        try
        {
            outFile.newLine();
            outFile.write("INVALID SYMBOL ENTERED.");
        }
        catch(IOException e)
        {
            //interruption writing to file
            e.printStackTrace();
            System.out.printf("Writing to file interrupted.&n");
        }
    }
    
    /** 
	 * Purpose: prints information for receipt
	 * Requires: price, quantity and a BufferedWriter file
         * Returns: none
    */
    public static void printReceiptInfo (double price, int qty, BufferedWriter outFile)
    {        
        System.out.printf("%d WIDGETS RECEIVED AT $%.2f. - COST: $%.2f%n", qty, price, (qty * price));
        try
        {
            printDashes(outFile); 
            outFile.newLine();
            
            outFile.write(qty + "\t\t" + "RECEIVED" + "\t\t$" + String.format("%.2f", price) + "\t\t$" + String.format("%.2f", (qty * price)));
        }
        catch(IOException e)
        {
            //interruption writing to file
            System.out.printf("Writing to file interrupted.%n");
            e.printStackTrace();            
        }
                
    }
    
    /** 
	 * Purpose: prints information for sales record
	 * Requires: quantity, sale price (including 20% markup) and a BufferedWriter output file
         * Returns: price for total sale
    */
    public static double printSaleInfo(int qty, double markup_sale, BufferedWriter outFile)
    {
        double total_sale = markup_sale * qty;
        
        System.out.printf("%d WIDGETS SOLD AT $%.2f. - SALE: $%.2f%n", qty, markup_sale, total_sale);
        try
        {
            outFile.newLine();            
            outFile.write(qty + "\t\t" + "SOLD" + "\t\t\t$" + String.format("%.2f", markup_sale) + "\t\t$" + String.format("%.2f", total_sale));
        }
        catch (IOException e)
        {
            //interruption writing to file
            e.printStackTrace();
            System.out.printf("Writing to file interrupted.%n");
        }
        finally
        {
            return total_sale;
        }
    }
    
    /** 
	 * Purpose: sells widgets (if available) from stack
	 * Requires: stack records to sell from, record for widgets to be sold, BufferedWriter output file 
         * Returns: remaining quantity of widgets unsold
    */
    public static int makeSales (RecordStack receiptStack, Record saleRecord, BufferedWriter outFile)
    {
        Record tempRecord;
        int qty = saleRecord.getQty();
        double total = 0;
        
        printDashes(outFile);        
        
        //while we have more to sell AND there's more to be bought
        while (!receiptStack.isEmpty() && qty > 0)
        {
            //Get latest Record received
            tempRecord = receiptStack.pop();
            //make sales from latest Record
            qty -= tempRecord.getQty();
            double price = tempRecord.getPrice();
            price += (price * 0.2);
            //check if we removed more than we had to
            if (qty < 0)
            {
                //print sale info and get amount of sales
                total += printSaleInfo(tempRecord.getQty() + qty, price, outFile);
                //if we removed more, change Record's qty to remaining qty and push back to stack
                tempRecord.changeQty(-qty);
                receiptStack.push(tempRecord);
                qty = 0;
            }
            else
            {
                total += printSaleInfo(tempRecord.getQty(), price, outFile);
            }
            
        }
        printSaleResult(total, outFile);
        //return remaining quantity
        return qty;
        
    }
     
    /** 
	 * Purpose: prints sales information to output file and console
	 * Requires: total price of sales made, BufferedWriter output file
         * Returns: none
    */
    public static void printSaleResult(double total, BufferedWriter outFile)
    {
        System.out.printf("TOTAL FOR THIS SALE: $%.2f%n", total);        
        try
        {
            outFile.newLine();
            outFile.write("TOTAL SALE \t\t\t\t\t\t$" + String.format("%.2f", total));
        }
        catch (IOException e)
        {
            e.printStackTrace();
            System.out.printf("Writing to file interrupted.%n");
        }
    }
    
    /** 
	 * Purpose: prints message for insufficient widgets
	 * Requires: quantity of widgets, BufferedWriter output file
         * Returns: none
    */
    public static void printInsuffWidg(int qty, BufferedWriter outFile)
    {
        System.out.printf("%d WIDGETS NOT AVAILABLE FOR SALE.%n", qty);
        try
        {
            outFile.newLine();
            outFile.write(qty + " WIDGETS NOT AVAILABLE FOR SALE.");
        }
        catch (IOException e)
        {
            e.printStackTrace();
            System.out.printf("Writing to file interrupted.%n");
        }
    }
    
    /** 
	 * Purpose: prints header for output file
	 * Requires: BufferedWriter output file
         * Returns: none
    */
    public static void printOutfileHead(BufferedWriter output)
    {
        try
        {
            output.write("CMPS 2433 - [Name]");
            output.newLine();
            output.write("PROGRAM 1 OUTFILE");
            output.newLine(); output.newLine();
            output.write("Welcome to Widget store program");
            output.newLine(); printDashes(output); output.newLine();
            output.write("QUANTITY\tTRANS-TYPE\t\tPRICE\t\tTOTAL");
        }
        catch(IOException e)
        {
            e.printStackTrace();
            System.out.printf("Writing to outfile interrupted.&n");
        }
        finally{}
                
    }
    
    /** 
	 * Purpose: prints dashes to output file
	 * Requires: BufferedWriter output file
         * Returns: none
    */
    public static void printDashes(BufferedWriter outFile)
    {
        try
        {
            outFile.newLine();
            outFile.write("----------------------------------------------------------------");
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
    }
    
    /** 
	 * Purpose: prints exit information to console and output file
	 * Requires: BufferedWriter output file
         * Returns: none
    */
    public static void printExit(BufferedWriter outFile)
    {        
        try
        {
            printDashes(outFile);
            outFile.newLine(); outFile.newLine();
            outFile.write("Thanks for using Widget store program.");
            outFile.close();
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
        System.out.printf("Thanks for using Widget store program.%n");
    }
    
    //MAIN FUNCTION
    public static void main(String[] args) {        
		    
        RecordStack myReceiptStack = new RecordStack ();
            
        //Output stream to use for writing output to file
        BufferedWriter output = getOutfile();
              
        promptForRecord();
        printOutfileHead(output);
        
        //while a null Record is not returned                
        for (Record myRecord = readRecord(); myRecord != null; myRecord = readRecord())
        {
            if (myRecord.getType() == 'r')
            {
                //push receipt Record and print info
                if (!myReceiptStack.isFull())
                {
                    myReceiptStack.push(myRecord);
                    printReceiptInfo(myRecord.getPrice(), myRecord.getQty(), output);
                }                        
            }
            else if (myRecord.getType() == 's')
            {
                //makes sales, print info                
                int qty = makeSales(myReceiptStack, myRecord, output);                    
                if (qty > 0)
                {
                    //If not all were sold
                    printInsuffWidg(qty, output);
                }
                                       
            }
            else
            {
                //If wrong symbol entered
                printInvalidSymbol(output);
            }
                                   
        }
                
        printExit(output);
               
}
   
    
}
