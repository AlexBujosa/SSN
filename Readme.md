# Index

1. **About the algorithm**
2. **Installation**
3. **User Documentation**
4. **Testing Plan**
5. **SSN**

# 1. **About the algorithm**

- The valid SSN (Social Security Number) must satisfy the following conditions:
  It should have 9 digits.
  It should be divided into 3 parts by hyphen (-).
  The first part should have 3 digits and should not be 000, 666, or between 900 and 999.
  The second part should have 2 digits and it should be from 01 to 99.
  The third part should have 4 digits and it should be from 0001 to 9999.

# 2. **Installation**

Firstable, you need to install the lastest version of Go. Using this command

- Go to downloads in the terminal and the you write wget https://go.dev/dl/go1.18.2.linux-amd64.tar.gz
- Then you run tar -xzvf go1.18.2.linux-amd64.tar.gz
- Optional (if you want to folders inside your go folder yo can use ll go, the meaning of ll is list all <'Name of folder '>)
- To see all the commands that go gives you, you write .go/bin/go in your terminal
- Now to use all the commands that you saw in the step before, everywhere you are. You need to move your go folder to /usr/local/. Using the next command to move: sudo mv go /usr/local/ and then you try using go in your terminal if the command is not recognized you need access to env-variable
- - write in your terminal -> vi ~/.bashrc
- - then press shift + g at the same time, this redirect you to bottom of the file and press O and then enter
- - Now you need to write #Add go path (this is a comment)
- - then you write export PATH=$PATH:/usr/local/bin, then press control c to get out and then you write :wq and press enter
- - then you write source ~/.bashrc and press enter
- - and finally you write go, and then this commands displays all the go commands, you can use

**For using the programm you need to run the next commands**

- cd Documents, to locate inside Documents folder then sudo mkdir (name_folder)
- then cd name_folder
- then you write git clone https://github.com/AlexBujosa/decimalAromano
- cd decimalAromano
- then write go run decimalAromano.go
- Finally this go to display the console application

# 3. **User Documentation**

[User Documentation](Docs/UserDocumentation.md)

# 4. **Testing Plan**

[Testing Plan](Docs/TestingPlan.md)
