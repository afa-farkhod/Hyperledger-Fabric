# Black Money Usage Tracker Built on Hyperledger Fabric
Implementation of Hyperledger Fabric with the demonstration and its usage against the black money counterfeit.

_____________________________________________________
Major problems which are need to be addressed in less economically developed countries in the Global Village are black money and
counterfeit. I implemented Hyperledger Fabric with real demonstrative web application to solve the previously mentioned issue.

**1.1	Problem Statement:**
 
Major issues of our today's topic are black money and their negative consequences on the government’s economy and ordinary people 
lifestyle. Counterfeit money is used illegally and its usage without paying any tax in the domestic market brings inflation in
the economy. Moreover, little work has been done to the analysis of the impact of the black market exchange rate on the long-run
demand for money in economies that have black market activities for their currencies, like less economically developed countries:
India, Pakistan, Uruguay and others. To be precise, foreign currency exchange rate in black market as double as the price in the 
governmental official bank.      

**1.2	Solution:**

Ideal solution for the problem stated above is to come up with an efficient monitoring system which should have no any
vulnerabilities against tracking counterfeit money and their usage in the domestic market. Monitoring should be done efficiently and
without any delay. What is important, the solution should ensure that fake money is detected and inform the user as well as authority 
about its existence in the system. So that preventive action could be taken immediately. Better solution for this case would be the
implementation of the Hyperledger Fabric technology. By making the ledger of all transactions including money serial numbers, our system
automatically detects any kind of counterfeit money and prevents their usage in the market.    

**1.3 Network setup and web application run:**

- Bring up the Hyperledger Network, for this we have to be in money-blockchain/money-app directory: 

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test1.PNG?raw=true)

- Install node modules inside the money-app directory:

<$cd...money-blockchain/money-app>

<$npm install -g>

- Then register admin:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test2.PNG?raw=true)

- Register user:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test3.PNG?raw=true)

- Run the server:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test4.PNG?raw=true)

- Web API on web browser:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test5.PNG?raw=true)

- Query all money packets:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test6.png?raw=true)

- Query ledger result on the terminal:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test7.png?raw=true)

- Query a specific money packet on web:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test8.PNG?raw=true)

- Specific transaction information on terminal:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test9.PNG?raw=true)

- Create new packet to the ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test10.PNG?raw=true)

- Generate unique id transaction:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test11.PNG?raw=true)

- New packet appears on the ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test12.PNG?raw=true)

- New record is added to the ledger on terminal environment:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test13.PNG?raw=true)

- Changing money holder in the ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test14.PNG?raw=true)

- Assigning transaction id:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test15.PNG?raw=true)

- Query all money packets after updating the owner information:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test16.PNG?raw=true)

- Owner update to the newest one on the terminal ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test17.PNG?raw=true)














