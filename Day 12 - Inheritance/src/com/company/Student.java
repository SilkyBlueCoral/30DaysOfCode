package com.company;

class Student extends Person {
    private int[] testScores;

    /*
    *   Class Constructor
    *
    *   @param firstName - A string denoting the Person's first name.
    *   @param lastName - A string denoting the Person's last name.
    *   @param id - An integer denoting the Person's ID number.
    *   @param scores - An array of integers denoting the Person's test scores.
    */
    // Write your constructor here
    Student(String firstName, String lastName, int id, int[] scores) {
        // explicit call to superclass constructor:

        super(firstName, lastName, id);

        this.testScores = scores;
    }
    /*
    *   Method Name: calculate
    *   @return A character denoting the grade.
    */
    // Write your method here

    public char calculate() {

        int numScores = testScores.length;
        int addedScores = 0;
        for (int x = 0; x < numScores; x++) {
            addedScores += testScores[x];

        }

        int average = addedScores / numScores;
        char grade = ' ';

        if (average < 40) {
            grade = 'T';
        } else if (40 <= average && average < 55) {
            grade = 'D';
        } else if (55 <= average && average < 70) {
            grade = 'P';

        } else if (70 <= average && average < 80) {
            grade = 'A';
        } else if (80 <= average && average < 90) {
            grade = 'E';
        } else if (90 <= average && average <= 100) {
            grade = 'O';
        }


        return grade;
    }
}


