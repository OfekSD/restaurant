while [ 1=1 ]
do
    nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
done