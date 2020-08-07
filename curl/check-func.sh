#!/bin/bash
 
# function check_response
# usage: check_response $got $expected $method $admin $url
check_response () {
  echo "checking method=$3, admin=$admin, url=$url"
  echo "response codes: got=$1; want=$2"
  if [ $1 != $2 ]; then
    echo -e '\033[1;31m  -> ERROR <- \e[0m'
  else
    echo -e '\e[32m ° O o . OK . o O °\e[0m'
  fi
}
# function check_response
 
 
if [ -z "$1" ]; then
  echo "usage: $0 inputfile server-url"
  echo "inputfile: the input file has the following format:"
  echo "      lines starting with # are ignored"
  echo "      the following space separated columns are needed: method, user:pw, url, response_code"
  echo "server-url: if the second parameter is missing, server-url defaults to https://localhost:8820"
  exit 1
fi
 
 
inFile="$1"
if  [ -z "$2" ]; then
  server="https://localhost:8820"
else
  server="$2"
fi
 
while read -r line
do
  line="$line"
  if [[ $line == \#* ]]; then
    continue
  fi
  method=$(echo $line | awk '{print $1}')
  chain=$(echo $line | awk '{print $2}')
  admin=$(echo $line | awk '{print $3}')
  url=$(echo $line | awk '{print $4}')
  url=$(sed -e "s/\$chain/$chain/g" <<< $url)
  expected=$(echo $line | awk '{print $NF}')
 
  R=`curl --cacert ../tls/public.pem -o - -s -w " %{http_code}\n" -X $method -u $admin "$server$url" `
  echo "Result of $line:"
  echo "$R"
  echo "/Result"
  got=`echo $R | awk  '{print $NF}' `
  check_response $got $expected $method $admin $url
 
done < "$inFile"
