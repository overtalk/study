#! bash/sh

now=$(date +%Y-%m-%d-%T)
count=0

download()
{
    for line in `cat url.txt`
    do
        mkdir $now
        echo "download-----rdb-"$count
        curl $line -o $now/rdb-$count
        count=$[$count+1]
    done
}


