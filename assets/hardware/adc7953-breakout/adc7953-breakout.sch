EESchema Schematic File Version 4
EELAYER 30 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title ""
Date ""
Rev ""
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Connector:Conn_01x19_Male J1
U 1 1 603B0761
P 2450 3800
F 0 "J1" H 2558 4881 50  0000 C CNN
F 1 "Conn_01x19_Male" H 2558 4790 50  0000 C CNN
F 2 "Connector_PinHeader_2.54mm:PinHeader_1x19_P2.54mm_Vertical" H 2450 3800 50  0001 C CNN
F 3 "~" H 2450 3800 50  0001 C CNN
	1    2450 3800
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x19_Male J2
U 1 1 603BF56D
P 4250 3800
F 0 "J2" H 4222 3732 50  0000 R CNN
F 1 "Conn_01x19_Male" H 4222 3823 50  0000 R CNN
F 2 "Connector_PinSocket_2.54mm:PinSocket_1x19_P2.54mm_Vertical" H 4250 3800 50  0001 C CNN
F 3 "~" H 4250 3800 50  0001 C CNN
	1    4250 3800
	-1   0    0    1   
$EndComp
$Comp
L power:GNDA #PWR0101
U 1 1 603C5753
P 6350 4950
F 0 "#PWR0101" H 6350 4700 50  0001 C CNN
F 1 "GNDA" H 6355 4777 50  0000 C CNN
F 2 "" H 6350 4950 50  0001 C CNN
F 3 "" H 6350 4950 50  0001 C CNN
	1    6350 4950
	1    0    0    -1  
$EndComp
Wire Wire Line
	5950 4700 5950 4950
Wire Wire Line
	5950 4950 6350 4950
Wire Wire Line
	6750 4700 6750 4950
Wire Wire Line
	6750 4950 6350 4950
Connection ~ 6350 4950
$Comp
L power:GNDA #PWR0102
U 1 1 603C84B2
P 5450 3800
F 0 "#PWR0102" H 5450 3550 50  0001 C CNN
F 1 "GNDA" H 5455 3627 50  0000 C CNN
F 2 "" H 5450 3800 50  0001 C CNN
F 3 "" H 5450 3800 50  0001 C CNN
	1    5450 3800
	1    0    0    -1  
$EndComp
$Comp
L ads7953sbdbtr:ADS7953SBDBTR U1
U 1 1 603A6FB5
P 6350 3700
F 0 "U1" H 6350 4765 50  0000 C CNN
F 1 "ADS7953SBDBTR" H 6350 4674 50  0000 C CNN
F 2 "Package_SO:TSSOP-38_4.4x9.7mm_P0.5mm" H 6550 4750 50  0001 C CNN
F 3 "https://www.ti.com/lit/ds/symlink/ads7950.pdf?HQS=dis-mous-null-mousermode-dsf-pf-null-wwe&DCM=yes&ref_url=https%3A%2F%2Fwww.mouser.com%2F&distId=26" H 6550 4750 50  0001 C CNN
	1    6350 3700
	1    0    0    -1  
$EndComp
Text GLabel 6750 2900 2    50   Input ~ 0
GPIO1
$Comp
L power:GNDA #PWR0103
U 1 1 603D0945
P 5450 3400
F 0 "#PWR0103" H 5450 3150 50  0001 C CNN
F 1 "GNDA" H 5455 3227 50  0000 C CNN
F 2 "" H 5450 3400 50  0001 C CNN
F 3 "" H 5450 3400 50  0001 C CNN
	1    5450 3400
	1    0    0    -1  
$EndComp
Wire Wire Line
	5450 3700 5450 3800
Connection ~ 5450 3800
$Comp
L power:GNDA #PWR0104
U 1 1 603D338A
P 5450 3100
F 0 "#PWR0104" H 5450 2850 50  0001 C CNN
F 1 "GNDA" H 5455 2927 50  0000 C CNN
F 2 "" H 5450 3100 50  0001 C CNN
F 3 "" H 5450 3100 50  0001 C CNN
	1    5450 3100
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0105
U 1 1 603D3FBD
P 7250 3200
F 0 "#PWR0105" H 7250 2950 50  0001 C CNN
F 1 "GNDA" H 7255 3027 50  0000 C CNN
F 2 "" H 7250 3200 50  0001 C CNN
F 3 "" H 7250 3200 50  0001 C CNN
	1    7250 3200
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0106
U 1 1 603D4D35
P 7250 3700
F 0 "#PWR0106" H 7250 3450 50  0001 C CNN
F 1 "GNDA" H 7255 3527 50  0000 C CNN
F 2 "" H 7250 3700 50  0001 C CNN
F 3 "" H 7250 3700 50  0001 C CNN
	1    7250 3700
	1    0    0    -1  
$EndComp
Text GLabel 5950 2900 0    50   Input ~ 0
GPIO2
Text GLabel 5950 3000 0    50   Input ~ 0
GPIO3
Text GLabel 5950 3200 0    50   Input ~ 0
REFP
Wire Wire Line
	5450 3100 5950 3100
Wire Wire Line
	5450 3400 5950 3400
Wire Wire Line
	5450 3700 5950 3700
Wire Wire Line
	5450 3800 5950 3800
Text GLabel 5950 3500 0    50   Input ~ 0
MXO
Text GLabel 5950 3600 0    50   Input ~ 0
AINP
Text GLabel 5950 4300 0    50   Input ~ 0
CH11
Text GLabel 5950 4400 0    50   Input ~ 0
CH10
Text GLabel 5950 4500 0    50   Input ~ 0
CH9
Text GLabel 5950 4600 0    50   Input ~ 0
CH8
Text GLabel 6750 3000 2    50   Input ~ 0
GPIO0
Text GLabel 6750 3100 2    50   Input ~ 0
+VBD
Wire Wire Line
	6750 3200 7250 3200
Wire Wire Line
	6750 3700 7250 3700
Text GLabel 6750 3300 2    50   Input ~ 0
SDO
Text GLabel 6750 3400 2    50   Input ~ 0
SDI
Text GLabel 6750 3500 2    50   Input ~ 0
SCLK
Text GLabel 6750 3600 2    50   Input ~ 0
CS
Text GLabel 6750 3900 2    50   Input ~ 0
CH0
Text GLabel 6750 4000 2    50   Input ~ 0
CH1
Text GLabel 6750 4100 2    50   Input ~ 0
CH2
Text GLabel 6750 4200 2    50   Input ~ 0
CH3
Text GLabel 6750 4300 2    50   Input ~ 0
CH4
Text GLabel 6750 4400 2    50   Input ~ 0
CH5
Text GLabel 6750 4500 2    50   Input ~ 0
CH6
Text GLabel 6750 4600 2    50   Input ~ 0
CH7
Text GLabel 2650 2900 2    50   Input ~ 0
GPIO2
Text GLabel 2650 3000 2    50   Input ~ 0
GPIO3
$Comp
L power:GNDA #PWR0107
U 1 1 603DF0BA
P 3100 3100
F 0 "#PWR0107" H 3100 2850 50  0001 C CNN
F 1 "GNDA" H 3105 2927 50  0000 C CNN
F 2 "" H 3100 3100 50  0001 C CNN
F 3 "" H 3100 3100 50  0001 C CNN
	1    3100 3100
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0108
U 1 1 603E3DEA
P 3100 3400
F 0 "#PWR0108" H 3100 3150 50  0001 C CNN
F 1 "GNDA" H 3105 3227 50  0000 C CNN
F 2 "" H 3100 3400 50  0001 C CNN
F 3 "" H 3100 3400 50  0001 C CNN
	1    3100 3400
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0109
U 1 1 603E40A6
P 3400 4850
F 0 "#PWR0109" H 3400 4600 50  0001 C CNN
F 1 "GNDA" H 3405 4677 50  0000 C CNN
F 2 "" H 3400 4850 50  0001 C CNN
F 3 "" H 3400 4850 50  0001 C CNN
	1    3400 4850
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0110
U 1 1 603E4548
P 3650 3200
F 0 "#PWR0110" H 3650 2950 50  0001 C CNN
F 1 "GNDA" H 3655 3027 50  0000 C CNN
F 2 "" H 3650 3200 50  0001 C CNN
F 3 "" H 3650 3200 50  0001 C CNN
	1    3650 3200
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0111
U 1 1 603E462B
P 3650 3700
F 0 "#PWR0111" H 3650 3450 50  0001 C CNN
F 1 "GNDA" H 3655 3527 50  0000 C CNN
F 2 "" H 3650 3700 50  0001 C CNN
F 3 "" H 3650 3700 50  0001 C CNN
	1    3650 3700
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0112
U 1 1 603E47A9
P 3100 3800
F 0 "#PWR0112" H 3100 3550 50  0001 C CNN
F 1 "GNDA" H 3105 3627 50  0000 C CNN
F 2 "" H 3100 3800 50  0001 C CNN
F 3 "" H 3100 3800 50  0001 C CNN
	1    3100 3800
	1    0    0    -1  
$EndComp
Text GLabel 2650 3200 2    50   Input ~ 0
REFP
Text GLabel 2650 3500 2    50   Input ~ 0
MXO
Text GLabel 2650 3600 2    50   Input ~ 0
AINP
Text GLabel 2650 4300 2    50   Input ~ 0
CH11
Text GLabel 2650 4400 2    50   Input ~ 0
CH10
Text GLabel 2650 4500 2    50   Input ~ 0
CH9
Text GLabel 2650 4600 2    50   Input ~ 0
CH8
Text GLabel 4050 2900 0    50   Input ~ 0
GPIO1
Text GLabel 4050 3000 0    50   Input ~ 0
GPIO0
Text GLabel 4050 3100 0    50   Input ~ 0
+VBD
Text GLabel 4050 3300 0    50   Input ~ 0
SDO
Text GLabel 4050 3400 0    50   Input ~ 0
SDI
Text GLabel 4050 3500 0    50   Input ~ 0
SCLK
Text GLabel 4050 3600 0    50   Input ~ 0
CS
Text GLabel 4050 3900 0    50   Input ~ 0
CH0
Text GLabel 4050 4000 0    50   Input ~ 0
CH1
Text GLabel 4050 4100 0    50   Input ~ 0
CH2
Text GLabel 4050 4200 0    50   Input ~ 0
CH3
Text GLabel 4050 4300 0    50   Input ~ 0
CH4
Text GLabel 4050 4400 0    50   Input ~ 0
CH5
Text GLabel 4050 4500 0    50   Input ~ 0
CH6
Text GLabel 4050 4600 0    50   Input ~ 0
CH7
Wire Wire Line
	4050 4700 3400 4700
Wire Wire Line
	3400 4700 3400 4850
Wire Wire Line
	2650 4700 3400 4700
Connection ~ 3400 4700
Wire Wire Line
	3650 3700 4050 3700
Wire Wire Line
	3650 3200 4050 3200
Wire Wire Line
	3100 3100 2650 3100
Wire Wire Line
	3100 3400 2650 3400
Wire Wire Line
	2650 3800 3100 3800
Wire Wire Line
	2650 3700 3100 3700
Wire Wire Line
	3100 3700 3100 3800
Connection ~ 3100 3800
$Comp
L power:VAA #PWR0113
U 1 1 60405435
P 5850 3300
F 0 "#PWR0113" H 5850 3150 50  0001 C CNN
F 1 "VAA" V 5868 3427 50  0000 L CNN
F 2 "" H 5850 3300 50  0001 C CNN
F 3 "" H 5850 3300 50  0001 C CNN
	1    5850 3300
	0    -1   -1   0   
$EndComp
$Comp
L power:VAA #PWR0114
U 1 1 60408FC1
P 6850 3800
F 0 "#PWR0114" H 6850 3650 50  0001 C CNN
F 1 "VAA" V 6868 3927 50  0000 L CNN
F 2 "" H 6850 3800 50  0001 C CNN
F 3 "" H 6850 3800 50  0001 C CNN
	1    6850 3800
	0    1    1    0   
$EndComp
Wire Wire Line
	6750 3800 6850 3800
Wire Wire Line
	5850 3300 5950 3300
$Comp
L power:VAA #PWR0115
U 1 1 6040A51D
P 3950 3800
F 0 "#PWR0115" H 3950 3650 50  0001 C CNN
F 1 "VAA" V 3968 3927 50  0000 L CNN
F 2 "" H 3950 3800 50  0001 C CNN
F 3 "" H 3950 3800 50  0001 C CNN
	1    3950 3800
	0    -1   -1   0   
$EndComp
$Comp
L power:VAA #PWR0116
U 1 1 6040ACB8
P 2750 3300
F 0 "#PWR0116" H 2750 3150 50  0001 C CNN
F 1 "VAA" V 2768 3427 50  0000 L CNN
F 2 "" H 2750 3300 50  0001 C CNN
F 3 "" H 2750 3300 50  0001 C CNN
	1    2750 3300
	0    1    1    0   
$EndComp
Wire Wire Line
	3950 3800 4050 3800
Wire Wire Line
	2750 3300 2650 3300
$Comp
L power:VAA #PWR0117
U 1 1 6041DBFA
P 2300 1250
F 0 "#PWR0117" H 2300 1100 50  0001 C CNN
F 1 "VAA" H 2317 1423 50  0000 C CNN
F 2 "" H 2300 1250 50  0001 C CNN
F 3 "" H 2300 1250 50  0001 C CNN
	1    2300 1250
	1    0    0    -1  
$EndComp
$Comp
L power:GNDA #PWR0118
U 1 1 6041E575
P 2350 1750
F 0 "#PWR0118" H 2350 1500 50  0001 C CNN
F 1 "GNDA" H 2355 1577 50  0000 C CNN
F 2 "" H 2350 1750 50  0001 C CNN
F 3 "" H 2350 1750 50  0001 C CNN
	1    2350 1750
	1    0    0    -1  
$EndComp
$Comp
L Device:C C1
U 1 1 6041F70F
P 1700 1500
F 0 "C1" H 1815 1546 50  0000 L CNN
F 1 "1uF" H 1815 1455 50  0000 L CNN
F 2 "Capacitor_SMD:C_0603_1608Metric" H 1738 1350 50  0001 C CNN
F 3 "~" H 1700 1500 50  0001 C CNN
	1    1700 1500
	1    0    0    -1  
$EndComp
$Comp
L Device:C C2
U 1 1 60420D0A
P 2050 1500
F 0 "C2" H 2165 1546 50  0000 L CNN
F 1 "1uF" H 2165 1455 50  0000 L CNN
F 2 "Capacitor_SMD:C_0603_1608Metric" H 2088 1350 50  0001 C CNN
F 3 "~" H 2050 1500 50  0001 C CNN
	1    2050 1500
	1    0    0    -1  
$EndComp
$Comp
L Device:C C3
U 1 1 60421148
P 3750 1500
F 0 "C3" H 3865 1546 50  0000 L CNN
F 1 "10uF" H 3865 1455 50  0000 L CNN
F 2 "Capacitor_SMD:C_0603_1608Metric" H 3788 1350 50  0001 C CNN
F 3 "~" H 3750 1500 50  0001 C CNN
	1    3750 1500
	1    0    0    -1  
$EndComp
Wire Wire Line
	1700 1350 2050 1350
Wire Wire Line
	2300 1250 2300 1350
Wire Wire Line
	2300 1350 2050 1350
Connection ~ 2050 1350
Wire Wire Line
	1700 1650 2050 1650
Wire Wire Line
	2050 1650 2350 1650
Wire Wire Line
	2350 1650 2350 1750
Connection ~ 2050 1650
$Comp
L power:GNDA #PWR0119
U 1 1 60425B8E
P 3750 1750
F 0 "#PWR0119" H 3750 1500 50  0001 C CNN
F 1 "GNDA" H 3755 1577 50  0000 C CNN
F 2 "" H 3750 1750 50  0001 C CNN
F 3 "" H 3750 1750 50  0001 C CNN
	1    3750 1750
	1    0    0    -1  
$EndComp
Text GLabel 3750 1100 2    50   Input ~ 0
REFP
Wire Wire Line
	3750 1350 3750 1100
Wire Wire Line
	3750 1650 3750 1750
$EndSCHEMATC
