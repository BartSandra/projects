#!/bin/bash

touch metrics.html
exec 1>./metrics.html

cpu1=$( top -b | head -3 | tail +3  | awk '{ print $2 }' | sed 's/,/./' )
cpu2=$( top -b | head -3 | tail +3  | awk '{ print $4 }' | sed 's/,/./' )
cpu=$( echo "$cpu1 + $cpu2" | bc )
echo -e "# HELP my_node_cpu_seconds_total Seconds the cpus spent in each mode.\n"
echo -e "# TYPE my_node_cpu_seconds_total counter\n"
echo -e "my_node_cpu_seconds_total $cpu\n"

ram=$(( $( sudo free | grep Mem | awk '{ print $3 }' ) * 1000 )) # RAM (общий объем оперативной памяти)
echo -e "# HELP my_node_memory_MemTotal_bytes Memory information field MemTotal_bytes.\n"
echo -e "# TYPE my_node_memory_MemTotal_bytes gauge\n"
echo -e "my_node_memory_MemTotal_bytes $ram\n"

mem_size=$(( $( sudo df | grep -w / | awk '{ print $2 }' ) * 1000 )) # общий объем памяти
echo -e "# HELP my_node_filesystem_size_bytes Filesystem size in bytes.\n"
echo -e "# TYPE my_node_filesystem_size_bytes gauge\n"
echo -e "my_node_filesystem_size_bytes $mem_size\n"

mem_used=$(( $( sudo df | grep -w / | awk '{ print $3 }' ) * 1000 )) # используется
echo -e "# HELP my_node_filesystem_used_bytes\n"
echo -e "# TYPE my_node_filesystem_used_bytes gauge\\n"
echo -e "my_node_filesystem_used_bytes $mem_used\n"

mem_free=$(( $( sudo df | grep -w / | awk '{ print $4 }' ) * 1000 )) # доступно
echo -e "# HELP my_node_filesystem_avail_bytes Filesystem space available to non-root users in bytes.\n"
echo -e "# TYPE my_node_filesystem_avail_bytes gauge\n"
echo -e "my_node_filesystem_avail_bytes $mem_free"