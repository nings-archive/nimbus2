declare install_status
declare build_status

compile () {
    echo "- Removing old nimbus2..."
    rm ~/go/src/nimbus2/main 2> /dev/null
    remove_status=$?
    if [ "$remove_status" = "1" ]; then
        echo "File \"nimbus2\" not found"
    fi

    go install nimbus2/lib
    install_status=$?

    echo "- Building main.go"
    go build ~/go/src/nimbus2/main.go
    build_status=$?
}

while true; do
    compile
    if [ "$install_status" == "0" ] && [ "$build_status" == "0" ]; then
        echo "- Success!"
        break
    fi
    echo "- There was an error in the install or build stages"
    echo "- Try again: <enter>; Abort: <ctrl-c>"
    read text
done
