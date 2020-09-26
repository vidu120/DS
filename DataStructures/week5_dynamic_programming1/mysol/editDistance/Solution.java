import java.util.Scanner;

public class Solution {
    public static Integer[][] editDistanceTrack(Integer[][] trace ,String firstString  ,String secondString ){
        



        return trace;
    }


    public static void main(String[] args){

        //scanner fot taking inputs
        Scanner scanner = new Scanner(System.in);

        //This is the inputs from the user
        String firstString = scanner.nextLine();
        String secondString = scanner.nextLine();

        Integer[][] trace = new Integer[100][100];

        //scanner part is done
        scanner.close();

        trace = editDistanceTrack(trace, firstString , secondString);
    }
}
