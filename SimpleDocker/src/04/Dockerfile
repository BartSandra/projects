FROM nginx
WORKDIR /home/
COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./server.c .
COPY ./start.sh .
RUN apt-get update 
RUN apt-get install -y gcc spawn-fcgi libfcgi-dev
ENTRYPOINT ["sh", "./start.sh"]