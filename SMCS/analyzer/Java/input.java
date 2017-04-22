/*
    CMPS 2143 - Record.java class implementation
     
 */

public class Record {
    private String date_time;
    private int quantity;
    private double price;
    char type;
    
    //Default Constructor
    public Record ()
    {
        date_time = "";
        quantity = 0;
        price = 0;
        type = '\0';
    }
    
    //Constructor
    public Record (String dat_tym, int qty, double price2, char tType)
    {
        date_time = dat_tym;
        quantity = qty;
        price = price2;
        type = tType;
    }
    
    //Constructor 
    public Record (String dat_tym, int qty, char tType)
    {
        date_time = dat_tym;
        quantity = qty;
        price = 0;
        type = tType;
    }
    
    /** 
	 * Returns the date_time stamp of a Record
	 * 
	 * @return date_time of record
    */
    public String getDateTime()
    {
        return date_time;
    }
    
    /** 
	 * Returns the quantity of a Record
	 * 
	 * @return quantity of record
    */
    public int getQty()
    {
        return quantity;
    }
    
    /** 
	 * Returns price of a Record
	 * 
	 * @return price of record
    */
    public double getPrice()
    {
        return price;
    }
    
    /** 
	 * Returns the type of a Record
	 * 
	 * @return type of record
    */
    public char getType()
    {
        return type;
    }
    
    /** 
	 * Changes the quantity in a Record
	 * 
	 * @param newQty to be changed to
    */
    public void changeQty(int newQty)
    {
        quantity = newQty;
    }
    
    /** 
	 * Changes the dateTime
	 * 
	 * @param newDateTime to be changed to
    */
    public void changeDateTime(String newDateTime)
    {
        date_time = newDateTime;
    }
    
    /** 
	 * Changes the price
	 * 
	 * @param newPrice to be changed to
    */
    public void changePrice(double newPrice)
    {
        price = newPrice;
    }
    
    /** 
	 * Changes the type
	 * 
	 * @param newType to be changed to
    */
    public void changeType(char newType)
    {
        type = newType;
    }
}