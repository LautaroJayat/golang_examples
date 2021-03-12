# Wait Groups

The cool thing about wait groups is that they allow us to block the execution of code until the count goes to 0.
The trick is just to add one to the count and make sure to trigger the .done() function when the child rutine finish it's job.