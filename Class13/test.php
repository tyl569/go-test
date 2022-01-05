<?php

run();

function timeSpend($method, int $n)
{
    $start = time();
    $ret = $method($n);
    $executeTime = time() - $start;
    echo "time spent:" . $executeTime . "\n";
    return $ret;
}

function slow(int $n): int
{
    sleep(1);
    return $n;
}

function run()
{
    echo timeSpend("slow", 10);
}
