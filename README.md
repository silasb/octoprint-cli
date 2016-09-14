# octoprint-cli

## WIP - 9/13/2016 - Usable, but be wary

Simple command line client for Octoprint.

```
git clone https://github.com/silasb/octoprint-cli
cd octoprint-cli
go build -o octoprint

# list files
./octoprint --host http://10.5.5.15:5000 --key 1234 files
# upload file
./octoprint --host http://10.5.5.15:5000 --key 1234 upload examples/test.gcode
```
