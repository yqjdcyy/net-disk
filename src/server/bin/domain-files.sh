APP_HOME=/data/service/webapps/files
RUN_PATH=/data/service/run
APP_NAME=files
LOG_PATH=/data/service/logs/files
LOG_FILE=$LOG_PATH/$APP_NAME.log
PID_FILE=$RUN_PATH/$APP_NAME.pid

APP_OPTS=" -c=/data/service/webapps/files/properties/gateway.properties "
for arg in "$@"
do
  if [ "$arg" = "start" -o "$arg" = "restart" ]; then
    continue;
  else
    APP_OPTS=${APP_OPTS}" "${arg}    
  fi
done

CMD="$APP_HOME/$APP_NAME $APP_OPTS"

psid=0

getpid() {
  psid=`cat $PID_FILE`
}
resetpid() {
  `echo "" > $PID_FILE`
}
setpid() {
  `echo "$psid" > $PID_FILE`
}
getprocesscount() {
  getpid
  if [ -z $psid ];then
    return 0
  else
    pcount=$(ps aux |grep "$psid" |grep "$APP_NAME" |wc -l)
    return $pcount
  fi
}

start() {
      echo "Starting $APP_NAME ..."
      echo "checking $APP_NAME status"
      status
      isstart=$? 
      if [ $isstart -eq '1' ];then
        return
      fi

      nohup $CMD > $LOG_FILE 2>&1 &
      psid=$!
	echo $psid
      setpid
      echo "================================"
      echo "info: $APP_NAME start success!"
      echo "================================"
      
}
stop() {
   
   echo "Stopping $APP_NAME ..."
    getpid
    if [ -z $psid ]; then
      echo "================================"
      echo "error: $APP_NAME already stopped!"
      echo "================================"
    else
      kill -9 $psid
      resetpid
      echo "================================"
      echo "info: $APP_NAME stopped!"
      echo "================================"
    fi
}

help() {
      $APP_HOME/$APP_NAME --help
}

status() {
      getprocesscount
      pcount=$?

      if [ $pcount -eq 1 ];then
        echo "$APP_NAME already started"
        return 1
      else
        echo "$APP_NAME not started"
        return 0
      fi
}

case "$1" in
   'start')
      start
      ;;
   'stop')
     stop
     ;;
   'restart')
     stop
     start
     ;;
   'help')
     help
     ;;
   'status')
     status
     ;;
  *)
     echo "Usage: $0 {start|stop|restart|help|status}"
     exit 1
esac
exit 0
