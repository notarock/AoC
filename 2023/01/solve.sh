total=0
while read line; do

    num=$(echo $line \
	| sed 's/one/o1ne/g' \
	| sed 's/two/tw2o/g' \
	| sed 's/three/thr3ee/g' \
	| sed 's/four/fo4ur/g' \
	| sed 's/five/fi5ve/g' \
	| sed 's/six/si6x/g' \
	| sed 's/seven/sev7en/g' \
	| sed 's/eight/eigh8t/g' \
	| sed 's/nine/ni9ne/g' \
	| grep -o '[0-9]' \
	| xargs | sed 's/ //g' \
	| sed -E 's/(.).+(.)/\1\2/')

    if [[ $num -lt 10 ]]; then
	num="$num$num"
    fi
    echo "Num is = $num"

    total=$(($total+$num))

done <input.txt

echo "Total = $total"
