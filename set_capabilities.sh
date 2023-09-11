#!/bin/bash
sudo setcap cap_net_raw,cap_net_admin,cap_net_bind_service+eip /home/jmbenlloch/go/myproject/build/bin/*
