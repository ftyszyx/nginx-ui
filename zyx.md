# 编译

go build -tags=jsoniter -ldflags "$LD_FLAGS -X 'github.com/0xJacky/Nginx-UI/settings.buildTime=$(date +%s)'" -o nginx-ui -v main.go

-tags=jsoniter - This flag tells the compiler to use the jsoniter library instead of the standard encoding/json library. JsonIter is a high-performance drop-in replacement for Go's standard JSON library
-ldflags - Allows passing arguments to the linker. In this case:
$LD_FLAGS - References existing linker flags defined elsewhere
-X 'github.com/0xJacky/Nginx-UI/settings.buildTime=$(date +%s)' - Sets a variable at link time:
Injects the Unix timestamp ($(date +%s)) into the buildTime variable in the settings package
-o nginx-ui - Specifies the output filename for the compiled binary
-v - Enables verbose output during the build process
