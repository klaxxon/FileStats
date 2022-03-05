# Simple file statistics
Provides counts by byte value, a histogram by byte value and the bit distribution.
<h2>Run against the main.go source file</h2>
<pre>
$>go run main.go main.go

Starting: /home/jgettys/Development/go/bin/dlv dap --check-go-version=false --listen=127.0.0.1:34723 --log-dest=3 from /home/jgettys/Development/filestats
DAP server listening at: 127.0.0.1:34723
File Statistics
File main.go, total bytes 1796

Counts by byte value
         00     01     02     03     04     05     06     07     08     09     0A     0B     0C     0D     0E     0F
00 :      0      0      0      0      0      0      0      0      0    119     99      0      0      0      0      0 
10 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
20 :    202      2     44      1      0     14      1      0     38     38      3     29     17      1     29      6 
30 :     22     12      7      1     11      2     18      0      3      1     20     18     15     27      3      1 
40 :      1      5     10     14      1      0      4      0      3      0      0      0      1      0      2      3
50 :     19      0      4      2      0      0      5      1      5      1      0     13      7     13      0      1 
60 :      0     83     38      7      6     54     58      8      7     57      0      3     40     34     71     53 
70 :      9      0     56     50    121     35     11      2      5     33      0     18      0     18      0      0 
80 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
90 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
A0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
B0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
C0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
D0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
E0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 
F0 :      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0      0 

Histogram by byte value
     00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F
00 :                             +  =
10 :                                                
20 :  @     :        .        :  :     .  .     .   
30 :  .  .              .           .  .  .  .      
40 :           .                                    
50 :  .                                .     .      
60 :     =  :        :  -        -        :  :  -  :
70 :        :  :  +  :           .     .     .      
80 :                                                
90 :                                                
A0 :                                                
B0 :                                                
C0 :                                                
D0 :                                                
E0 :                                                
F0 :                                                

Counts by bit
01 = 780  43.430%
02 = 714  39.755%
04 = 694  38.641%
08 = 855  47.606%
10 = 590  32.851%
20 = 1463  81.459%
40 = 992  55.234%
80 = 0  0.000%



