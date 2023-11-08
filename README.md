## ascii-art-output by ayessenb and yotarov
* `ayessenb` 
* `yotarov` 

### Objectives

Ascii-art-output is a program which specifies the name of file and outputs the provided string inside of that file converted to graphic format using the specified banner.
Moreover, the program can convert the graphical representation of the letters to text.
The argument will be a flag, --output=<fileName.txt>, in which --output is the flag and  <fileName> is the file name that will contain the output of the program.
The argument will be a flag, --reverse=<fileName.txt>, in which --reverse is the flag and <fileName> is the file name. The program must then print this string in normal text.


- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Usage: go run . [OPTION]
- Ascii-art-reverse project implemented, the program accept the correctly formatted [OPTION] [STRING].
- Ascii-art-fs project implemented, the program accept the correctly formatted [STRING] [BANNER].
- The program must still able to run with a single [STRING] argument.


### Banner Format

* shadow
* standard
* thinkertoy

#### Description:


* The program reads the banner file where the height of each letter is 8 lines
* The first line is not considered
* Each character in banner file is separated by one '\n'
* Since Banner contains the ascii values between 32 and 126 then the other runes are not accepted.
* If the Banner file will be changed the program will give the Error message
* If the program will get wrong input format, the program will give the Error message
* If the second argument is not "standard" or "shadow" or "thinkertoy" then the program will give the Error messages
* If the input string argument doesn't contain ascii value then the program gives the Error message.(between 32 and 126 (both included))


#### Improved skills:
* go lang programming skills 
* usage of string libraries of Go
* handling new line char.
* usage of many functions and files


## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@github.com:compygirl/ascii-art-output.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd ascii-art-output/output
```
Run a program:
```CMD/Terminal 
go run . --output=go.txt "Salem Alem" standard
```


Output: __Salem Alem__
+ go.txt
```bash
  _____           _                                       _                     $
 / ____|         | |                              /\     | |                    $
| (___     __ _  | |   ___   _ __ ___            /  \    | |   ___   _ __ ___   $
 \___ \   / _` | | |  / _ \ | '_ ` _ \          / /\ \   | |  / _ \ | '_ ` _ \  $
 ____) | | (_| | | | |  __/ | | | | | |        / ____ \  | | |  __/ | | | | | | $
|_____/   \__,_| |_|  \___| |_| |_| |_|       /_/    \_\ |_|  \___| |_| |_| |_| $
                                                                                $
                                                                                $
```


Run a program:
```CMD/Terminal 
go run . --output=go.txt "Salem Alem" thinkertoy
```


Output: __Salem Alem__
+ go.txt
```bash
                                               
 o-o       o                   O   o           
|          |                  / \  |           
 o-o   oo  | o-o o-O-o       o---o | o-o o-O-o 
    | | |  | |-' | | |       |   | | |-' | | | 
o--o  o-o- o o-o o o o       o   o o o-o o o o 
                                               
                                               


```

Run a program:
```CMD/Terminal 
go run . --output=go.txt "Salem\nAlem" shadow 
```


Output: __Salem Alem__
+ go.txt
```bash
                                             
  _|_|_|          _|                         
_|         _|_|_| _|   _|_|   _|_|_|  _|_|   
  _|_|   _|    _| _| _|_|_|_| _|    _|    _| 
      _| _|    _| _| _|       _|    _|    _| 
_|_|_|     _|_|_| _|   _|_|_| _|    _|    _| 
                                             
                                             
                                    
  _|_|   _|                         
_|    _| _|   _|_|   _|_|_|  _|_|   
_|_|_|_| _| _|_|_|_| _|    _|    _| 
_|    _| _| _|       _|    _|    _| 
_|    _| _|   _|_|_| _|    _|    _| 
                                    
                                    


```


Run a program:
```CMD/Terminal 
go run . --output=go.txt "Salem\ 
Alem\n" standard
```

Output: __Salem Alem__
+ go.txt
```bash
  _____           _                     __            
 / ____|         | |                    \ \           
| (___     __ _  | |   ___   _ __ ___    \ \          
 \___ \   / _` | | |  / _ \ | '_ ` _ \    \ \         
 ____) | | (_| | | | |  __/ | | | | | |    \ \        
|_____/   \__,_| |_|  \___| |_| |_| |_|     \_\       
                                                      
                                                      
            _                     
    /\     | |                    
   /  \    | |   ___   _ __ ___   
  / /\ \   | |  / _ \ | '_ ` _ \  
 / ____ \  | | |  __/ | | | | | | 
/_/    \_\ |_|  \___| |_| |_| |_| 
                                  
                                  



```


## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
10.05.2023.
