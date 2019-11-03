FROM jenkins/jenkins:2.202


RUN mkdir -p /usr/share
ADD go1.13.4.linux-amd64.tar.gz /usr/share/


COPY install.sh /root/
RUN chmod 777 /root/install.sh \
     &&  /root/install.sh


EXPOSE 8080

USER root

