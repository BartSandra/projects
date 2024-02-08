#!/bin/bash

codes=("200" "201" "400" "401" "403" "404" "500" "501" "502" "503")
methods=("GET" "POST" "PUT" "DELETE" "PATCH")
agents=("Mozilla" "Google Chrome" "Opera" "Safari" "Internet Explorer" "Microsoft Edge" "Crawler and bot" "Library and net tool")
arr_site=("/request.html" "/request1.html" "/request2.html" "/request3.html" "/request4.jpg" "/request5.html" "/request6.html")

for ((days=1; days<=5; days++))
do
    logs=$(shuf -i 100-1000 -n 1)
    minutes=$((1440 / logs))
    
    for ((l=0; l<logs; l++))
    do
        ip=$(shuf -i 1-255 -n 1).$(shuf -i 0-255 -n 1).$(shuf -i 0-255 -n 1).$(shuf -i 0-255 -n 1)
        time=$(($minutes * l + $(shuf -i 0-$minutes -n 1)))
        date=$(date +"%d/%b/%Y:%H:%M:%S %z" -d "$((5 - days)) days ago +$time minutes 0")
        gen_ref_url="www.somesite.$(cat /dev/urandom | tr -dc A-Z-a-z-0-9 | head -c5).fun"
        
        log_file="access$days.log"
        echo "$ip - - [$date] \"${methods[$(shuf -i 0-4 -n 1)]} ${arr_site[$(shuf -i 0-6 -n 1)]} HTTP/1.1\" ${codes[$(shuf -i 0-9 -n 1)]} $(shuf -i 0-126 -n 1) \"$gen_ref_url\" \"${agents[$(shuf -i 0-6 -n 1)]}\" \"-\"" >> $log_file
    done
done


# 200 - успешный запрос
# 201 - успешный запрос и создание нового ресурса
# 400 - некорректный запрос от клиента
# 401 - требуется аутентификация для доступа к ресурсу
# 403 - доступ к ресурсу запрещен
# 404 - запрашиваемый ресурс не найден
# 500 - внутренняя ошибка сервера
# 501 - сервер не поддерживает функциональность, необходимую для обработки запроса
# 502 - сервер получил некорректный ответ при попытке обратиться к другому серверу
# 503 - сервер временно недоступен