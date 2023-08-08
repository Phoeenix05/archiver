#!/bin/zsh

sudo chown :wheel service.plist
launchctl load service.plist
launchctl start io.github.phoeenix05.archiver
