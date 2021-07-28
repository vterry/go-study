# Polymorphism Study

In this example I've tried to implement a possible solution to the problem bellow.

**The problem:** 

We need to create a solution to evaluate candidate's skills to a Dev position.

The candidate must fill a formulary with those informations:
* Name
* Mail address
* Level of knowledge from 0 to 10 in the follow skills:
    *  JAVA
    * C#
    * Golang
    * Javascript
    * VUE
    * PHP

After candidate fill the formulary, we must send an email to candidate following the rules:

* If candidate knowledge in JAVA, C# and Golang is between 7 and 10, send an email with a Backend message.

* If candidate knowledge in Javascript, VUE and PHP is between 7 and 10, send an email with a Frontend message.

* If candidate knowledge is between 7 and 10 in all skills, send an email with a Backend and Frontend message.

* If candidate knowledge don't fit with those requirements, send a generic email.