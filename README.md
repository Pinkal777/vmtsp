# vmtsp
Challenge Problem 

Multiple Traveling salesman problem is considered as one of the NP hard problem due to its computational complexity.

There are multiple approches to solve this problem 
one of them is using Clarke and Wright’s Savings Algorithm 
greedy approch is beeing used by this Clarke and Wright’s Savings Algorithm algorithm by calculating the cost/savings in advance for all the possible routes
and targating from the max saving route we will assign route to Driver.

for smooth operation of the program MakeFile is created

To build binary of this project,checkout repo and run "make build" command as follows. that will ceate executable binary vmtsp.exe
//image needed

After building - executable program binary can be supplied to given python test suite using following command.
this will iterate through all the problem files inside trainingProblems and evaluate against the program binary.
i.e   	py .\evaluateShared.py --cmd .\vmtsp.exe --problemDir trainingProblems
//image needed


program can also be run independently against sigle file using "go run . problem1.txt" without any test suite script 
//image needed



Reference:
saving algorithm
https://www.tandfonline.com/doi/full/10.1080/21642583.2019.1674220
https://aswani.ieor.berkeley.edu/teaching/FA15/151/lecture_notes/ieor151_lec18.pdf
https://www.ncbi.nlm.nih.gov/pmc/articles/PMC3870871/