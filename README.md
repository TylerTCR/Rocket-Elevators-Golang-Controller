# Rocket-Elevators-Golang-Controller
This is Tyler's Rocket Elevators' Commercial Controller coded in Golang.

A brief run-down on how the controller works; When someone at the lobby presses a floor they want to go to, it will find the best column and elevator within that column to pick up the person and take them to the floor they want. However, if someone is at another floor and they call an elevator, it will find the best elevator within that column to take them to the lobby.

### Installation

With golang installed on your computer, all you need to do is initialize the module:

`go mod init Rocket-Elevators-Commercial-Controller`

The code to run the scenarios is included, and can be executed with:

`go run . <SCENARIO-NUMBER>`

### Running the tests

To test this controller with scenarios, run the following in the terminal:

`go test`

To get more details about each test, simply add the `-v` flag at the end like so: 

`go test -v` 