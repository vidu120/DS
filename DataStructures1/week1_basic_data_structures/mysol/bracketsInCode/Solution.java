import java.util.*;

class Node{
    char data;
    Node next;

    Node(char key){
        this.data = key;
        this.next = null;
    }
}

class SinglylinkedlistStack{
    private Node head;
    //constructor for linked list
    SinglylinkedlistStack(){
        this.head = null;
    }
    void push(char data){
        if (head == null){
            head = new Node(data);
        }else{
            Node newNode = new Node(data);
            newNode.next = head;
            head = newNode;
        }
    }
    char pop(){
        Node keep = head;
        head = head.next;
        return keep.data;
    }

    boolean empty(){
        if (head == null){
            return true;
        }
        return false;
    }

    void print(){
        Node keep = head;
        while (keep != null){
            System.out.print(keep.data);
            keep = keep.next;
        }
        System.out.println();
    }
}


class Solution{
    public static void main(String[] args){
        SinglylinkedlistStack list = new SinglylinkedlistStack();
        Scanner scanner = new Scanner(System.in);
    }
}