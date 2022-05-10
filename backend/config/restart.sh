#!/bin/sh
process_name="nft_http"
main_pid=`ps -aux | grep ${process_name}  | grep -v grep | awk '{print $2}'`
kill -9 ${main_pid}
chmod +777 ./${process_name}
./${process_name}  >> run.log 2>&1 &
