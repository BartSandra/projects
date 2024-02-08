## 1. Установка готового dashboard

* Установим готовый дашборд `Node Exporter Quickstart and Dashboard` с официального сайта Grafana Labs 

https://grafana.com/grafana/dashboards/13978-node-exporter-quickstart-and-dashboard/

* Добавим дашборд в Grafana: `Dashboards->Manage->Import` и загружаем скачанный .json файл

![part_8](./Images/1.png)

## 2. Тесты добавленного dashboard

* Запустим bash-скрипт из `Part 2`

![part_8](./Images/3.png)

* Проверим результаты работы

![part_8](./Images/2.png)

* Запустим команду

`$ stress -c 2 -i 1 -m 1 --vm-bytes 32M -t 60s`

* Проверим результаты работы

![part_8](./Images/4.png)

## 3. Настройка статической маршрутизации между двумя машинами

* Опишем сетевой интерфейс первой машины

![part_8](./Images/7.png)

`$ sudo neplan apply`

* Опишем сетевой интерфейс второй машины

![part_8](./Images/8.png)

`$ sudo neplan apply`

* Проверим, что машины пингуются

![part_8](./Images/9.png)

![part_8](./Images/10.png)

## 4. Тест нагрузки сети

* Первая машина выступает в роли сервера, выполним команду 

`$ iperf3 -s`

![part_8](./Images/12.png)

* Вторая машина выступает в роли клиента, выполним команду 
    
`$ iperf3 -c 172.24.116.8`

![part_8](./Images/11.png)

* Посмотрим на нагрузку сетевого интерфейса

![part_8](./Images/13.png)
