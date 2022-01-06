<?php

interface Programmer {
    public function WriteHelloWorld();
}

class PHPProgrammer implements Programmer {
    public function WriteHelloWorld()
    {
        return "echo \"Hello World\"";
    }
}

function main()
{
    $pProgrammer = new PHPProgrammer();
    echo $pProgrammer->WriteHelloWorld();
}

main();