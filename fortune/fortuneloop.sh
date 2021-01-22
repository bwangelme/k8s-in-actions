#!/bin/bash
#
# Author: bwangel<bwangel.me@gmail.com>
# Date:1月,04,2021 09:10


trap "exit" SIGINT
INTERNAL=$1
echo "Configured to generate new fortune every $INTERNAL seconds"
mkdir /var/htdocs
while :
do
    echo $(date) Writing fortune to /var/htdocs/index.html
    /usr/games/fortune > /var/htdocs/index.html
    sleep $INTERNAL
done
