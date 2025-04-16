#!/bin/bash
dbDomains=`ls domain/*.go`
for eachFile in $dbDomains
do
    echo ${eachFile##*/};
    mockgen -source=domain/${eachFile##*/} -destination=mocks/repositories/${eachFile##*/} -package=mocks_infrastructures
done