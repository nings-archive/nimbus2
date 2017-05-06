declare install_status
declare build_status

compile () {
    echo "- Removing old nimbus2..."
    rm ~/go/src/nimbus2/nimbus2 2> /dev/null
    remove_status=$?
    if [ "$remove_status" = "1" ]; then
        echo "File \"nimbus2\" not found"
    fi

    sh ~/go/src/nimbus2/lib/install.sh
    install_status=$?

    echo "- Building nimbus2.go..."
    go build ~/go/src/nimbus2/nimbus2.go
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
