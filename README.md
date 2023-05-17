ARP List Program

This program scans the local network for ARP (Address Resolution Protocol) tables and writes the results to a file. The output file will contain a list corresponding IP addresses. Maybe nmap files for your network?

Prerequisites

This program requires Go to be installed on your system. You also need to have root privileges to run this program because it uses the sudo command to execute the arp-scan command.

To install Go on your system, follow the instructions on the official website.
To check if you have root privileges, run the following command:

sudo whoami

If the command returns your username, you have root privileges.

Installation
To install the program, clone the repository and run the following command in the terminal:

go install

This will compile and install the program in your Go bin directory.

Usage
To use the program, run the following command in the terminal:

sudo arpList

This will execute the program, scan the local network for ARP tables, and write the results to a file named output.txt in the current directory.

License

This program is licensed under the MIT License.