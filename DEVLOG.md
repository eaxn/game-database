# The Devlog aka How I made this

## 02/31/2022 - I started the project
## 03/31/2022 - I created a task system
- I created a task system
  - every task has a repeating mechanism (by duration, for example 1 day or by date, for example every monday) => the option (by date) is not available yet
  - every task is a seperated go routine (so you can run multiple tasks at the same time)
  - every task is built in a Task struct, for avoiding any errors (for example infinite loops), which cause the whole program to crash or to be stopped (by spamming the go routines)
