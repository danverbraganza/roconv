# Testing needs to be a lot better.

diff <( go run roconv.go -mode=1 1 2 5 9 10 19 20 1999 1000000 10000000 ) <( echo "I
II
V
IX
X
XIX
XX
MCMXCIX
M̅
M̅M̅M̅M̅M̅M̅M̅M̅M̅M̅"
)

diff <( go run roconv.go -mode=I I MCMXCIX XXV VL M̅) <( echo "1
1999
25
45
1000000"
)

echo "If you see no output from diff, all the tests passed."


