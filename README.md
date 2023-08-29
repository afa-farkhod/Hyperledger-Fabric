# [Black Money Usage Tracker Built on Hyperledger Fabric](https://www.kci.go.kr/kciportal/ci/sereArticleSearch/ciSereArtiView.kci?sereArticleSearchBean.artiId=ART002594524)

_____________________________________________________
- Major problems which are need to be addressed in less economically developed countries in the Global Village are black money and
counterfeit. I implemented Hyperledger Fabric with real demonstrative web application to solve the previously mentioned issue.

### [Hyperledger Fabric](https://www.hyperledger.org/use/fabric)

- A fundamental element of a Hyperledger Fabric blockchain network is the set of peer nodes (or, simply, peers). Peers are fundamental because they manage ledgers and smart contracts. Starting in Hyperledger Fabric v2.4, peers also manage transaction proposals and endorsements by running the Fabric Gateway service. Recall that a ledger immutably records all of the transactions generated by smart contracts (which in Hyperledger Fabric are contained in chaincode, more on this later) and endorsed by the required organizations. Smart contracts and ledgers encapsulate the processes and information, respectively, that are shared by channel peers. These aspects of peers make them a good starting point for understanding a Fabric network.

<p align="center">
  <img src="https://user-images.githubusercontent.com/24220136/231626996-af1607f9-0001-4855-ac8c-c1ddb3a4338a.png" alt="Image">
</p>

**1.1.	Problem Statement**
 
- Major issues of our today's topic are black money and their negative consequences on the government’s economy and ordinary people 
lifestyle. Counterfeit money is used illegally and its usage without paying any tax in the domestic market brings inflation in
the economy. Moreover, little work has been done to the analysis of the impact of the black market exchange rate on the long-run
demand for money in economies that have black market activities for their currencies, like less economically developed countries:
India, Pakistan, Uruguay and others. To be precise, foreign currency exchange rate in black market as double as the price in the 
governmental official bank.      

**1.2.	Solution**

- Ideal solution for the problem stated above is to come up with an efficient monitoring system which should have no any
vulnerabilities against tracking counterfeit money and their usage in the domestic market. Monitoring should be done efficiently and
without any delay. What is important, the solution should ensure that fake money is detected and inform the user as well as authority 
about its existence in the system. So that preventive action could be taken immediately. Better solution for this case would be the
implementation of the Hyperledger Fabric technology. By making the ledger of all transactions including money serial numbers, our system
automatically detects any kind of counterfeit money and prevents their usage in the market.    

**1.3. Network setup and web application run**

- Bring up the Hyperledger Network, for this we have to be in `money-blockchain/money-app` directory: 

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test1.PNG?raw=true)

- Install node modules inside the money-app directory:
  ```
  $ cd...money-blockchain/money-app
  $ npm install -g
  ```
- Then register admin:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test2.PNG?raw=true)

- Register user:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test3.PNG?raw=true)

- Run the server:

<p align="center">
  <img src="https://github.com/af4092/money-blockchain/blob/master/images/Test4.PNG" alt="Image">
</p>

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

<p align="center">
  <img src="https://github.com/af4092/money-blockchain/blob/master/images/Test10.PNG" alt="Image">
</p>

- Generate unique id transaction:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test11.PNG?raw=true)

- New packet appears on the ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test12.PNG?raw=true)

- New record is added to the ledger on terminal environment:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test13.PNG?raw=true)

- Changing money holder in the ledger:

<p align="center">
  <img src="https://github.com/af4092/money-blockchain/blob/master/images/Test14.PNG" alt="Image">
</p>

- Assigning transaction id:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test15.PNG?raw=true)

- Query all money packets after updating the owner information:

<p align="center">
  <img src="https://github.com/af4092/money-blockchain/blob/master/images/Test16.PNG" alt="Image">
</p>

- Owner updates to the newest one on the terminal ledger:

![alt text](https://github.com/af4092/money-blockchain/blob/master/images/Test17.PNG?raw=true)

## [Publication](https://www.dbpia.co.kr/journal/articleDetail?nodeId=NODE09354936&language=ko_KR&hasTopBanner=true)


<p align="center">
  <img src="https://github.com/af4092/Hyperledger-Fabric/assets/24220136/16a6b45e-e06f-408b-80a8-8e5e86f86348" alt="Image">
</p>
