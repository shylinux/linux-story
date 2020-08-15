#!/bin/bash
# git &>/dev/null || yum install -y git

[ -f ~/.ish/plug.sh ] || [ -f ./.ish/plug.sh ] || git clone https://github.com/shylinux/intshell ./.ish
[ "$ISH_CONF_PRE" != "" ] || source ./.ish/plug.sh || source ~/.ish/plug.sh

require show.sh
require help.sh
require miss.sh

ish_miss_prepare_compile
ish_miss_prepare_install
# ish_miss_prepare_develop

# ish_miss_prepare_volcanos
# ish_miss_prepare learning
# ish_miss_prepare_icebergs
# ish_miss_prepare toolkits
