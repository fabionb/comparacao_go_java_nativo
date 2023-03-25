set term png small size 800,600
set output "memory.png"
set ylabel "RSS"
set ytics nomirror
set yrange [0:*]
plot "memory.log" using 1 with lines axes x1y1 title "RSS"

