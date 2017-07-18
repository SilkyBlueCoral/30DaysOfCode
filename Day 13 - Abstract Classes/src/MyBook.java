/**
 * Created by silk on 7/18/17.
 */

public class MyBook extends Book {

    /**
     *   Class Constructor
     *
     *   @param title The book's title.
     *   @param author The book's author.
     *   @param price The book's price.
     **/


    int price;

    MyBook(String title,String author, int price){

        super(title, author);
        this.price=price;

    }


    @Override
    void display() {
        System.out.println("Title: "+this.title+"\nAuthor: "+this.author+"\nPrice: "+this.price);
    }
}

