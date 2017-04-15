/*
    CMPS 2143 - Position.java class implementation
     
 */

public class Position implements IPosition {
    private int x_coord;
    private int y_coord;
     
    public Position()
    {
        x_coord = 0;
        y_coord = 0;
    }
    
    public Position(int x, int y)
    {
        x_coord = x;
        y_coord = y;
    }
    
    //Copy constructor
    public Position(Position otherPos)
    {
        this.x_coord = otherPos.x_coord;
        this.y_coord = otherPos.y_coord;
    }
    
    public int getX()
    {
        return x_coord;
    }
    
    public int getY()
    {
        return y_coord;
    }
    
    public void setX(int newX)
    {
        x_coord = newX;
    }
    
    public void setY(int newY)
    {
        y_coord = newY;
    }
    
    public void setCoord(int newX, int newY)
    {
        x_coord = newX;
        y_coord = newY;
    }
}
