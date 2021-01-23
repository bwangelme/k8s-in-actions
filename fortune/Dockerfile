FROM ubuntu:latest
ADD fortuneloop.sh /bin/fortuneloop.sh
RUN apt-get update; apt-get -y install fortune; chmod a+x /bin/fortuneloop.sh
ENTRYPOINT ["/bin/fortuneloop.sh"]
CMD ["10"]
