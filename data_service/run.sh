#!/bin/bash
echo "Bash version ${BASH_VERSION}..."
for i in {3041..3061}
  do 
    (go run client/client.go) > out/${i}.out 2>&1 &
 done