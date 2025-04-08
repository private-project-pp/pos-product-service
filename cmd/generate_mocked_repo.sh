#!/bin/bash
dbDomains=`ls domain/*.go`
for eachFile in $dbDomains
do
    echo ${eachFile##*/};
    mockgen -source=domain/${eachFile##*/} -destination=mocks/repositories/${eachFile##*/} -package=mocks_repository
done
# if [ "$1" = "ALL" ]; then
#     echo "Semua $1"
# elif [ "$1" != "" ]; then
#     echo "Spesific $1"
# else 
#     echo "Kosong cak"
# fi