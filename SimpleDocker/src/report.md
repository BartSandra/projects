## Part 1. Готовый докер

* Скачиваем докер образ nginx через `$ docker pull nginx`:

![docker_images](Images/1.png)

* Проверяем наличие докер образа через `$ docker images`:

![docker_images](Images/2.png)

* Запустить докер образ через `$ docker run -d [image_id|repository]`:

![docker_images](Images/3.png)

* Проверяем, что образ запустился через `$ docker ps`:

![docker_images](Images/4.png)

* Посмотреть информацию о контейнере через `$ docker inspect[container_id|container_name]`:

![docker_images](Images/5.png)

* По выводу команды определить и поместить в отчёт:
    - размер контейнера:

    ![docker_images](Images/6.png)

    - список замапленных портов:

    ![docker_images](Images/7.png)

    - ip контейнера:
    
    ![docker_images](Images/8.png)

* Остановить докер образ через `$ docker stop [container_id|container_name]`

* Проверить, что образ остановился через `$ docker ps`:

![docker_images](Images/9.png)

* Запустить докер с замапленными портами 80 и 443 на локальную машину через команду `$ docker run -p 80:80 -p 433:433 -d nginx`:

![docker_images](Images/10.png)

* Проверить, что в браузере по адресу `$ localhost:80` доступна стартовая страница nginx:

![docker_images](Images/11.png)

* Перезапустить докер контейнер через `$ docker restart [container_id|container_name]`:

* Проверим, что контейнер запустился `$ docker ps`:

![docker_images](Images/12.png)


## Part 2. Операции с контейнером

* Прочитать конфигурационный файл nginx.conf внутри докер контейнера через команду `$ docker exec [container_id|container_name] cat /etc/nginx/nginx.conf`

![docker_images](Images/13.png)

* Создать на локальной машине файл nginx.conf. Настроить в нем по пути /status отдачу страницы статуса сервера nginx:

![docker_images](Images/14.png)

* Скопировать созданный файл nginx.conf внутрь докер образа через команду docker `$ docker cp nginx.conf [container_id|container_name]:/etc/nginx` Перезапустить nginx внутри докер образа через команду `$ docker exec [container_id|container_name] nginx -s reload`

![docker_images](Images/15-16.png)


* Проверить, что по адресу localhost:80/status отдается страничка со статусом сервера nginx:

![docker_images](Images/17.png)

* Экспортировать контейнер в файл container.tar через команду `$ docker export [container_id|container_name] > container.tar`

* Остановить контейнер:

![docker_images](Images/18.png)

* Удалить образ через `$ docker rmi -f [image_id|repository]`, не удаляя перед этим контейнеры:

![docker_images](Images/19.png)

* Удалить остановленный контейнер:

![docker_images](Images/20.png)

* Импортировать контейнер обратно через команду `$ docker import -c 'cmd ["nginx", "-g", "daemon off;"]' container.tar nginx:latest`

* Запустить импортированный контейнер:

![docker_images](Images/21.png)

* Проверим, что по адресу `localhost:80/status` отдается страничка со статусом сервера nginx:

![docker_images](Images/22.png)

