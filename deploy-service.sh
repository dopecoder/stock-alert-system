#!/bin/sh
go build
go install
sudo cp trigger.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl disable trigger.service
sudo systemctl enable trigger.service
sudo systemctl start trigger.service
sudo systemctl status trigger.service