<?php

class Pet
{
    public function speak()
    {
        echo "....\n";
    }

    public function speakTo(string $name)
    {
        $this->speak();
        echo $name;
    }
}

class  Dog extends Pet
{
    public function speak()
    {
        echo "Wang!\n";
    }
}

function main()
{
    $dog = new Dog();
    $dog->speak();
    $dog->speakTo("肥龙");
}

main();