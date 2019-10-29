# Go Quiz
## Key of Learning
- Go Routines
- Channel
- CLI Flags
## Description
This is the 1st exercise from https://gophercises.com/

The goal is to build a cli quiz app with a timer on it.

Questions & it's answer are stored inside csv file. So we need to parse it.

This assesment is divided into 2 part.

### 1st Part
We just create a normal quiz app, without a timer.
After user finished answering the question, the app will calculate the scores.

My approach to finishing the 1st part :
1. read strings inside `problems.csv` with `ioutil` package, resulting a `byte slice`
2. convert it to strings. And then use `csv` package to parse it into a slice of records.
3. And then we loop the records to get user's answers
4. Calculate the score.

### 2nd Part
2nd Part is pretty tough. Because we have to implement a timer on it.
Default timer is 10 seconds.
But user can modify it using `-d` flags

My approach to finishing the 2nd part :
1. The general idea is the same from 1st part.
2. use `flag` package to setup the CLI flags.
3. I used `go routine` and `channel` to use the timer.
4. If the timer is ended, it force the app to submmit the user's answer, whether he is finished or not
5. Calculate the score