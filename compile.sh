WinBashErr='interrupted system call'
declare install_status
declare build_status

compile () {
    printf "Removing old main...\n"
    rm ~/go/src/nimbus2/main 2> /dev/null
    local remove_status=$?
    if [ "$remove_status" = "1" ]; then
        printf "  WARN: File \"main\" not found\n"
    fi

    printf "Installing lib/...\n"
    local installOutErr="$(go install nimbus2/lib 2>&1)"; install_status=$?
    while [[ "$installOutErr" == *"$WinBashErr"* ]]; do
        printf "  ERR: $installOutErr\n"
        installOutErr="$(go install nimbus2/lib 2>&1)"; install_status=$?
    done
    if [ ! "$installOutErr" == "" ]; then
        printf "  ERR: $installOutErr\n"
    fi

    printf "Building main.go...\n"
    local buildOutErr="$(go build ~/go/src/nimbus2/main.go 2>&1)"; build_status=$?
    while [[ "$buildOutErr" == *"$WinBashErr"* ]]; do
        printf "  ERR: $buildOutErr\n"
        buildOutErr="$(go build ~/go/src/nimbus2/main.go 2>&1)"; build_status=$?
    done
    if [ ! "$buildOutErr" == "" ]; then
        printf "  ERR: $buildOutErr\n"
    fi
}

while true; do
    compile
    if [ "$install_status" == "0" ] && [ "$build_status" == "0" ]; then
        printf "Compiled without interrupted system call.\n"
        break
    fi
    printf "There was an error in the install or build stages\n"
    printf "Try again: <enter>; Abort: <ctrl-c>\n"
    read text
done
