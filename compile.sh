echo "- Removing old nimbus2..."
rm ~/go/src/nimbus2/nimbus2

sh ~/go/src/nimbus2/lib/install.sh

echo "- Building nimbus2.go..."
go build ~/go/src/nimbus2/nimbus2.go
