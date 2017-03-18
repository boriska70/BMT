#!/bin/bash

mkdir -p ./.cover
profile="./.cover/cover.out"
declare -a testsuite_failures=()
echo "mode: count" >"$profile"


declare -a folders=( $(go list ./... | grep -v vendor) )

for folder in "${folders[@]}"; do
  echo Testing $folder
#  IFS='/' read -ra names <<< "$folder"
#  namesLen=${#names[@]}
#  let "namesLen--"
#  pkg=$(echo ${names[$namesLen]})
  f=$(echo ./.cover/$(echo $folder | tr / -).cover)
  tf=$(echo ./.cover/$(echo $folder | tr / -)_tests.xml)
  go test -cover -covermode="count" -coverprofile="$f" $folder | go-junit-report > "$tf"
  #check for failures in order to set correct exit code after all
  grep -F "testsuite " "$tf"
  if [ $? -eq 0 ]; then
      grep -F "failures=\"0\"" "$tf"
        if [ $? -ne 0 ]; then
          testsuite_failures+=($folder)
        fi
  fi
done
#create full cover profile
grep -h -v "^mode:" ./.cover/*.cover >>"$profile"

echo failures: $testsuite_failures
if [ "${#testsuite_failures[@]}" -eq "0" ]; then
  echo all tests passed
  exit 0
else
  echo "FAILED TESTSUITE(S) FOUND: $testsuite_failures"
  exit 13
fi
