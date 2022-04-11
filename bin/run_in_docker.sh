if [ $# -lt 1 ]
then
    port=9999
else
    port=$1
fi
echo "running server on port $port!"
docker run -d -p 9999:$port unicorn_center #要用detech -p前加-d