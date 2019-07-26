#!/bin/bash
set -Eeuo pipefail
trap "echo ERR trap fired!" ERR

echo -e "\033[1;37mInstalling tusk task runner...\033[0m"
curl -sL https://git.io/tusk | bash -s -- -b .direnv/bin latest

#echo -e "\033[1;37mBuilding the development environment...\033[0m"
#tusk env.build
