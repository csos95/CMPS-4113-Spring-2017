/*
    CMPS 2143 - RecordStack.java Implementation
 */

public class RecordStack implements IRecordStack {
    
    private Record Inventory [];
    public final int MAXSIZE = 20;	
    private int maxStack = MAXSIZE;
    private int top = -1;
    
    //Default Contructor
    public RecordStack ()
    {
        Inventory = new Record [MAXSIZE];
    }
        
    //Constructor to instantiate with number of Records
    public RecordStack (int numRecords)
    {
        maxStack = numRecords;
        Inventory = new Record [maxStack];
    }
    
    //Copy constructor
    public RecordStack (RecordStack otherStack)
    {
        maxStack = otherStack.maxStack;
        top = -1;
        Inventory = new Record [maxStack];
        copyStack(otherStack);
    }
    
    //private method for copy constructor
    private void copyStack (RecordStack otherStack)
    {
        //check if otherStack is not empty
        if (otherStack.top > -1)
        {
            //pop from otherStack
            Record tempRecord = otherStack.Inventory[otherStack.top];
            otherStack.top--;
            //recurse
            copyStack(otherStack);
            //on getting back, push Record on both RecordStacks
            top++;
            Inventory[top] = tempRecord;
            otherStack.top++;
            otherStack.Inventory[otherStack.top] = tempRecord;
        }        
    }
    
    //Check if Stack is full
    public boolean isFull ()
    {
        return (top == (maxStack - 1));
    }
        
    //Check if Stack is Empty    
    public boolean isEmpty ()
    {
        return (top < 0);
    }
        
    //Add Record to Stack    
    public void push (Record item)
    {
        if (top != (maxStack - 1))
        {
            top++;
            Inventory[top] = item;
        }
        else
        {
            System.out.println ("STACK IS FULL!");
        }
        
    }
        
    //Remove Record from top of Stack
    public Record pop ()
    {
        if (top != -1)
        {
            Record temp = Inventory[top];
            top--;
            return temp;                
        }
        else
        {
            System.out.println ("STACK IS EMPTY!");
            return null;
        }
    }    
        
}
