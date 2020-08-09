##Error in a concurrency/parallelism
This code shows us using of array `counter` in functions that run in goroutines.
Incorrectness is that functions should increment array to **2000000**, but it does up to **1893443-2000000**.
So, this program show race condition in what one goroutine ahead of another.
####For example:
This process can be bad in banking.	
So we have two operations:
- deposit
- pay for something
	
In this example we run **2** goroutines for **2** transactions:
1. deposit
2. pay for something

On the bank account is **1000$** 
1. deposit **500$**
2. pay **300$**

We run 2 goroutines and then two functions get shared value (**1000$**). Then process it.
Deposit func changes bank account to **1500$**, but pay function has old value, and it overrides bank account value.
Changes it to **700$**. At the end of the story user lost **500$**
To prevent this we should use Mutex/atomic/channels.