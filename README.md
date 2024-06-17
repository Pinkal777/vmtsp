# vmtsp
Challenge Problem-Multiple Traveling Salesman problem

Due to its computational complexity, the multiple traveling salesman problem is considered one of the NP-hard problems.

There are multiple approaches to solve this problem 
one of them is using Clarke and Wright’s Savings Algorithm 
greedy approach is being used by Clarke and Wright’s Savings Algorithm algorithm by calculating the cost/savings in advance for all the possible routes
and targeting from the max saving route, we will assign the route to the Driver.

Makefile is created to ensure ease of use for Golang program.

-To build the binary of this project, clone the repo and run the "make build" command as follows. that will create executable binary vmtsp.exe
![image](https://github.com/Pinkal777/vmtsp/assets/22688340/00bc782a-477f-4fd0-836a-a9effe29a425)
![image](https://github.com/Pinkal777/vmtsp/assets/22688340/dd9b192f-143f-4c99-84be-76c9ecc643e7)


-After building - the executable program binary can be supplied to the given Python test suite using the following command.
This will iterate through all the problem files inside trainingProblems and evaluate against the program binary.

i.e   	py .\evaluateShared.py --cmd .\vmtsp.exe --problemDir trainingProblems

![image](https://github.com/Pinkal777/vmtsp/assets/22688340/5c383f70-5a20-46aa-bbea-fa9ae618bcba)
![image](https://github.com/Pinkal777/vmtsp/assets/22688340/67038221-866c-48e7-897c-658325ceb80d)
![image](https://github.com/Pinkal777/vmtsp/assets/22688340/372f5151-5266-4bb6-a1dd-9aa2dfbb69c5)



-The program can also be run independently against a single file using "go run . problem1.txt" without any test suite script 
![image](https://github.com/Pinkal777/vmtsp/assets/22688340/4428f7b8-6d76-49ba-8b37-9f7a17288ea2)



:References:
https://www.tandfonline.com/doi/full/10.1080/21642583.2019.1674220
https://aswani.ieor.berkeley.edu/teaching/FA15/151/lecture_notes/ieor151_lec18.pdf

https://www.ncbi.nlm.nih.gov/pmc/articles/PMC3870871/
