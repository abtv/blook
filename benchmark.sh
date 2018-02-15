delim="----------------"
echo "Building blook"
go build
echo $delim

file=large.log
file_from=10000000
file_to=40000000
search_from=17000400
search_to=17000700
expected_res=301

if [ ! -e "$file" ] ; then
    echo "File $file doesn't exist. Creating it, please wait"
    seq -f %1.f $file_from $file_to > "$file"
    echo $delim
fi

echo "Benchmarking. Please, wait: the whole benchmark can take some time"
echo $delim
echo "blook results:"
blook_res=$(time ./blook $search_from $search_to large.log | wc -l)
echo $delim
echo "sed results:"
sed_res=$(time sed -ne '/'$search_from'/,/'$search_to'/ p' large.log | wc -l)
echo $delim
echo "awk results:"
awk_res=$(time awk -v from=$search_from -v to=$search_to '$1>=from && $1<=to' large.log | wc -l)
echo $delim

if [[ $blook_res -eq $expected_res ]] && [[ $sed_res -eq $expected_res ]] && [[ $awk_res -eq $expected_res ]] ; then
  echo "Done"
else
  echo "Wrong: lines count should be the same for all the commands"
fi
