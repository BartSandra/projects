# Report

## 1. Part 1. Installation of the OS

- the Ubuntu version

`$ cat /etc/issue`

![linux](images/linux1_1.png)

## 2. Part 2. Creating a user

- 2.1 Add a user other than the one created during installation.

`$ sudo adduser nikita`

![linux](images/linux2_3.png)

- 2.2 Result of cat /etc/passwd command output

`$ cat /etc/passwd`

![linux](images/linux2_2.png)

`$ cat /etc/passwd | grep nikita`

![linux](images/linux2_4.png)

- 2.3 The user must be added to adm group

`$ sudo usermod -G adm nikita`

`$ cat /etc/group | grep adm`

![linux](images/linux2_5.png)

## 3. Part 3. Setting up the OS network

- 3.1 Set the machine name as user-1

`$ sudo hostname user-1`

`$ hostnamectl`

![linux](images/linux3_1.png)

- 3.2 Set the time zone corresponding to your current location.

`$ timedatectl`

![linux](images/linux3_2.png)

`$ timedatectl list-timezones | grep Moscow`

`$ sudo timedatectl set-timezone Europe/Moscow`

`$ timedatectl`

![linux](images/linux3_3.png)

- 3.3 Output the names of the network interfaces using a console command.

`$ ip link show`

![linux](images/linux3_4.png)

or `$ ifconfig -a`

![linux](images/linux3_5.png)

>lo(local loopback) – виртуальный интерфейс, присутствующий по умолчанию в любом Linux. Он используется для отладки сетевых программ и запуска серверных приложений на локальной машине. С этим интерфейсом всегда связан адрес 127.0.0.1. У него есть dns-имя – localhost.

- 3.4 To get the ip address of the device you are working on from the DHCP server.

`$ hostname -I`

![linux](images/linux3_6.png)

or `$ ip r`

![linux](images/linux3_7.png)

>DHCP (Dynamic Host Configuration Protocol — протокол динамической настройки узла) — сетевой протокол, позволяющий сетевым устройствам автоматически получать IP-адрес и другие параметры, необходимые для работы в сети TCP/IP.

- 3.5 The external ip address of the gateway (ip) 

`$ wget -O - -q icanhazip.com`

![linux](images/linux3_8.png)

and the internal IP address of the gateway, aka default ip address (gw)

`$ ip route`

![linux](images/linux3_9.png)

- 3.6 Set static (manually set, not received from DHCP server) ip, gw, dns settings (use public DNS servers, e.g. 1.1.1.1 or 8.8.8.8).

`$ sudo vim /etc/netplan/00-installer-config.yaml`

Changing the file /etc/netplan/00-installer-config.yaml

![linux](images/linux3_10.png)

`$ sudo netplan apply` - applied changes in netplan

`$ reboot` - reboot the virtual machine

We check that the static network settings (ip, gw, dns) correspond to those set in the previous point.

`$ inconfig`

![linux](images/linux3_11.png)

Check if remote hosts are successfully pinged 1.1.1.1 

`$ ping 1.1.1.1`

![linux](images/linux3_12.png)

and ya.ru

`$ ping ya.ru`

![linux](images/linux3_13.png)

# Part 4. OS Update

Update the system packages to the latest version

`$ sudo apt-get dist-upgrade`

![linux](images/linux4.png)

# Part 5. Using the sudo command

>Команда sudo предоставляет возможность пользователям выполнять команды от имени суперпользователя root, либо других пользователей.

Change the OS hostname from the username created in Part 2 (using sudo).

`$ sudo hostnamectl set-hostname user-2`

`$ hostnamectl`

![linux](images/linux5.png)

`$ hostname`

![linux](images/linux5_1.png)

## Part 6. Installing and configuring the time service

Set up automatic time synchronization service

`$ sudo timedatectl`

`$ timedatectl show`

![linux](images/linux6.png)

## Part 7. Installing and using text editors

## - VIM

`$ sudo apt install vim` - Install VIM text editor

`$ vim test_vim.txt` - create a test_vim.txt file

![linux](images/linux7_1.png)

- esc 
- :wq - Exit with saving
- Enter 

![linux](images/linux7_2.png)

- esc
- :q! - Exit without saving
- Enter

![linux](images/linux7_3.png)

- /'What we want to find' - Search

![linux](images/linux7_5.png)

- s/'WHAT'/'WHAT' - Replacement

![linux](images/linux7_4.png)

## - JOE

`$ sudo apt install joe` - Install JOE text editor

`$ joe test_joe.txt` - create a test_joe.txt file

![linux](images/linux7_6.png)

- Ctrl+K, Q, Y - Exit with saving

![linux](images/linux7_7.png)

![linux](images/linux7_8.png)

- Ctrl+K, Q, N - Exit without saving

![linux](images/linux7_9.png)

- Ctrl+K, F - Search

![linux](images/linux7_10.png)

- Ctrl+K, F, R, Y - Replacement

![linux](images/linux7_11.png)

![linux](images/linux7_12.png)

## - MCEDIT

`$ sudo apt install mcedit` - Install MCEDIT text editor

`$ mcedit test_mcedit.txt` - create a test_mcedit.txt file

![linux](images/linux7_13.png)

- F2 /YES /F10 - Exit with saving

![linux](images/linux7_14.png)

- F10 /NO - Exit without saving

![linux](images/linux7_15.png)

- F7 /'What are we looking for' /OK || FIND ALL - Search

![linux](images/linux7_16.png)

- F4 /'WHAT'/'WHAT'  /REPLACE || ALL - Replacement

![linux](images/linux7_19.png)

![linux](images/linux7_17.png)

![linux](images/linux7_18.png)

## Part 8. Installing and basic setup of the SSHD service

`$ netstat -tan`

![linux](images/linux8_1.png)

- 8.1 Install the SSHd service

`$ sudo apt-get install ssh`

`$ sudo apt install openssh-server`

![linux](images/linux8_2.png)

- 8.2  Add an auto-start of the service whenever the system boots

`$ sudo systemctl enable ssh`

`$ systemctl status ssh`

![linux](images/linux8_3.png)

- 8.3 Reset the SSHd service to port 2022.

`$ sudo nano /etc/ssh/sshd_config`

Edit the file /etc/ssh/sshd_config

![linux](images/linux8_4.png)

- 8.4 Show the presence of the sshd process using the ps command. To do this, you need to match the keys to the command.

>ps - Вывести информацию об активных процессах;

>ps -e или ps -A(a) - Вывести информацию обо всех процессах;

>ps -a - Выбрать все процессы, кроме фоновых;

>ps -d(g) - Вывести информацию обо всех процессах, кроме лидеров групп;

>ps -N - Выбрать все процессы кроме указанных;

>ps -С - Выбирать процессы по имени команды;

>ps -G - Выбрать процессы по ID группы;

>ps -t - Выдавать информацию только о процессах, ассоциированных с терминалами из заданного списка_терминалов. Терминал - это либо имя файла-устройства, например ttyномер или console, либо просто номер, если имя файла начинается с tty;

>ps -p 'pid' - Выбрать процессы PID;;

>ps --ppid - Выбрать процессы по PID родительского процесса;

>ps -s - Выбрать процессы по ID сессии;

>ps -u - Выбрать процессы пользователя;

>ps -ef - Вывести полный список;

Show presence of sshd process

`$ ps aux | grep sshd`

![linux](images/linux8_8.png)

or `$ ps -e | grep sshd`

![linux](images/linux8_9.png)

- 8.5 Reboot the system.

`$ service ssh restart`

![linux](images/linux8_6.png)

`$ reboot`

`$ netstat -tan`

![linux](images/linux8_7.png)

>-t - Отображает только соединения TCP

>-a - Вывод всех активных подключений TCP и прослушиваемых компьютером портов TCP и UDP

>-n - Вывод активных подключений TCP с отображением адресов и номеров портов в числовом формате без попыток определения имен

>Proto: Протокол соединения

>recv-Q: Очередь получения сети

>send-Q: Сетевая очередь отправки

>Local Address: IP-адрес локального компьютера и номер используемого порта

>Foreign Address: IP-адрес и номер порта удаленного компьютера, подключенного к данному сокету

>State: Состояние TCP-соединения

>0.0.0.0: Клиентские устройства, такие как ПК, показывают адрес 0.0.0.0, когда они не подключены к какой-либо сети TCP / IP. Устройство может получить этот адрес по умолчанию, если оно не в сети. В случае сбоев назначения адреса, он может быть автоматически назначен DHCP. На случай, если ваше устройство настроено на этот адрес, оно не может общаться с любыми другими устройствами в сети через IP.

# Part 9. Installing and using the top, htop utilities:

- 9.1 Install and run the top and htop utilities

`$ sudo apt install htop`

`$ top`

![linux](images/linux9_1.png)

- uptime - 2:07;

- количество авторизованных пользователей - 1;

- общая загрузка системы - 0.00, 0.00, 0.00;

- общее количество процессов - 90;

- загрузка cpu - 0.0%;

- загрузка памяти - 157Mб;

`$ top -o %MEM`

- pid процесса занимающего больше всего памяти - 623 

![linux](images/linux9_3.png)

`$ top -o %CPU`

- pid процесса, занимающего больше всего процессорного времени - 1785 

![linux](images/linux9_4.png)

- 9.2

## sort by PID

`$ htop --sort-key PID`

![linux](images/linux9_5.png)

## sort by PERCENT_CPU

`$ htop --sort-key PERCENT_CPU`

![linux](images/linux9_6.png)

## sort by PERCENT_MEM

`$ htop --sort-key PERCENT_MEM`

![linux](images/linux9_7.png)

## sort by TIME

`$ htop --sort-key TIME`

![linux](images/linux9_8.png)

## Filter 'sshd'

`$ htop`

- F4

![linux](images/linux9_9.png)

## Search 'syslog'

- F3

![linux](images/linux9_10.png)

## with hostname, clock and uptime output added

- F2

![linux](images/linux9_11.png)

## Part 10. Using the fdisk utility

`$ sudo fdisk -l`

![linux](images/linux9_12.png)

- Название диска: VBOX HARDDISK

- Размер: 50.47 GiB

- Количество секторов: 105816064

- Размер swap: 3.8G

`$ free -h`

![linux](images/linux9_13.png)

## Part 11. Using the df utility

`$ df`

![linux](images/linux11_1.png)

- Размер раздела - 51758528;

- Размер занятого пространства - 6928420;

- Размер свободного пространства - 42168476;

- Процент использования - 15%;

- Единица измерения в выводе - килобайт.

`$ df -Th`

![linux](images/linux11_2.png)

- Размер раздела - 50G;

- Рразмер занятого пространства - 6.7G;

- Размер свободного пространства - 41G;

- Процент использования - 15%

- Тип файловой системы для раздела - ext4.

## Part 12. Using the du utility

Output the size of the /home, /var, /var/log folders

in human readable format

`$ sudo du -sh /var/log /home /var`  

![linux](images/linux11_3.png)

in bytes

`$ sudo du -s --block-size=1 /var/log /home /var` 

![linux](images/linux11_4.png)

Output the size of all contents in /var/log (not the total, but each nested element using *)

`$ sudo du -sh /var/log/*` 

![linux](images/linux11_5.png)

## Part 13. Installing and using the ncdu utility

Install the ncdu utility

`$ sudo apt install ncdu`

Output the size of the /home, /var, /var/log folders

![linux](images/linux13_4.png)

`$ sudo ncdu /home`

![linux](images/linux13_1.png)

`$ sudo ncdu /var`

![linux](images/linux13_2.png)

`$ sudo ncdu /var/log`

![linux](images/linux13_3.png)

## Part 14. Working with system logs

- 14.1 /var/log/dmesg

Open for viewing:

`$ sudo vim /var/log/dmesg`

![linux](images/linux13_5.png)

- 14.2 /var/log/syslog

`$ sudo vim /var/log/syslog`

![linux](images/linux13_6.png)

- 14.3 /var/log/auth.log

`$ sudo vim /var/log/auth.log`

![linux](images/linux13_7.png)

- Последняя успешная авторизация: Jun 18 19:46:29;

- Имя пользователя: finchmar;

- Метод входа в систему: pam-unix.

- 14.4 Restart the SSHd service

`$ sudo systemctl restart ssh`

`$ cat /var/log/syslog | grep ssh`

![linux](images/linux13_8.png)

## Part 15. Using the CRON job scheduler

Using the job scheduler, run the uptime command in every 2 minutes

`$ sudo crontab -e`

![linux](images/linux13_9.png)

Find lines in the system logs (at least two in a given time range) about the execution.

`$ grep -i cron /var/log/syslog`

![linux](images/linux13_11.png)

Display a list of current jobs for CRON

`$ sudo crontab -l`

![linux](images/linux13_12.png)

Remove all tasks from the job scheduler

`$ sudo crontab -r`

`$ sudo crontab -l`

![linux](images/linux13_13.png)



