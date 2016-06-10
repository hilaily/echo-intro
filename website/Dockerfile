FROM reg.lab.st/armor:0.1.0
MAINTAINER Vishal Rana <vr@labstack.com>

ADD armor.json /etc
ADD public /www

EXPOSE 80

CMD ["-c", "/etc/armor.json"]
