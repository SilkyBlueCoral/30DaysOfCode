/**
 * Created by silk on 7/18/17.
 */
public class Difference {
    private int[] elements;
    public int maximumDifference;

    Difference(int[] a){
        this.elements=a;

    }


    public void computeDifference(){

        maximumDifference=0;

        for (int x=0; x<elements.length;x++){

            for (int y=0; y<elements.length;y++){

                int diff=Math.abs(elements[x]-elements[y]);
                if (diff > maximumDifference) {
                    maximumDifference=diff;
                }
            }


        }



    }

}

