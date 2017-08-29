#!/bin/bash

options="--v=3"
options="$optihns --threadNum=12"
options="$options --host=http://127.0.0.1:28080"
options="$options --kind=memory --memory=100 --delay=150"
#options="$options --kind=cpu --cpu=100"
options="$options --time=100"
echo "$options"
./_output/webclient $options
