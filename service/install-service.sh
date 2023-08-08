#!/bin/zsh

go build

# Stop and unload the service
# launchctl stop io.github.phoeenix05.archiver
# launchctl unload service.plist

sudo chown :wheel service.plist
launchctl load service.plist
launchctl start io.github.phoeenix05.archiver
