<?php
class DirectionInstruction
{
    public $name;
    public $deltaX;
    public $deltaY;

    function __construct($name, $deltaX, $deltaY)
    {
        $this->name = $name;
        $this->deltaX = $deltaX;
        $this->deltaY = $deltaY;
    }
}
;

$data = file_get_contents("./data.txt");
$lines = explode("\n", $data);
$array = array();
for ($y = 0; $y < count($lines); $y++) {
    echo "<tr>";
    for ($x = 0; $x < strlen($lines[$y]); $x++) {
        if (ctype_space($lines[$y][$x])) {
            continue;
        }

        $array[$y][$x] = $lines[$y][$x];
    }
    echo "</tr>";
}

$directions = [
    new DirectionInstruction("right", 1, 0),
    new DirectionInstruction("left", -1, 0),
    new DirectionInstruction("up", 0, -1),
    new DirectionInstruction("down", 0, 1),
    new DirectionInstruction("right-up diagonal", 1, -1),
    new DirectionInstruction("left-up diagonal", -1, -1),
    new DirectionInstruction("right-down diagonal", 1, 1),
    new DirectionInstruction("left-down diagonal", -1, 1)
];


$xmases = 0;
$crossmasses = 0;

for ($y = 0; $y <= count($array) - 1; $y++) {
    for ($x = 0; $x <= count($array[$y]) - 1; $x++) {
        $char = $array[$y][$x];

        if ($char === 'X') {
            foreach ($directions as $dir) {
                if (check($array, $x, $y, $dir->deltaX, $dir->deltaY)) {
                    echo "'XMAS' found at: (" . $x . ", " . $y . ") towards (" .
                        $x + $dir->deltaX . ", " . $y + $dir->deltaY . "), with a " . $dir->name . "<br>";

                    $xmases++;
                }
            }
        }

        if ($char === 'A') {
            if (checkCrossMas($array, $x, $y, $directions)) {
                echo 'Cross MAS found at: ' . $x . ", " . $y . "<br>";
                $crossmasses++;
            }
        }
    }
}

function checkCrossMas($array, $startX, $startY, $directionInstructions)
{
    $diagonals = array_filter($directionInstructions, "diagonal");

    $passes = 0;

    // could make this enum but lazy
    // 0 - up
    // 1 - right
    // 2 - bottom
    // 3 - left
    $mDir = 0;

    for (; $mDir < 4; $mDir++) {

        foreach ($diagonals as $dir) {
            if (simple_check($array, $startX, $startY, $dir->deltaX, $dir->deltaY, mapDirToCharForCross($dir, $mDir), 1)) {
                $passes++;
            }
        }

        if($passes === 4){
            return true;
        }else{
            $passes = 0;
        }
    }

    return false;
}

function mapDirToCharForCross($dirInst, $mDir)
{
    switch ($mDir) {
        case 0: {
            if (str_contains($dirInst->name, "up")) {
                return 'M';
            }
            break;
        }
        case 1: {
            if (str_contains($dirInst->name, "right")) {
                return 'M';
            }
            break;
        }
        case 2: {
            if(str_contains($dirInst->name, "down")){
                return 'M';
            }
            break;
        }
        case 3: {
            if(str_contains($dirInst->name, "left")){
                return 'M';
            }
            break;
        }

    }

    return 'S';
}

function diagonal($var)
{
    return str_contains($var->name, "diagonal");
}

function check($array, $startX, $startY, $deltaX, $deltaY)
{
    $chars = ['M', 'A', 'S'];
    for ($i = 1; $i <= count($chars); $i++) {
        if (!simple_check($array, $startX, $startY, $deltaX, $deltaY, $chars[$i - 1], $i)) {
            return false;
        }
    }

    return true;
}

function simple_check($array, $startX, $startY, $deltaX, $deltaY, $char, $currentStep)
{
    $newX = $startX + $currentStep * $deltaX;
    $newY = $startY + $currentStep * $deltaY;

    // check bounds
    if ($newX < 0 || $newY < 0 || $newY >= count($array) || $newX >= count($array[$newY])) {
        return false;
    }

    // check character
    if ($array[$newY][$newX] !== $char) {
        return false;
    }

    return true;
}

?>


<table border="1">
    <?php
    // Add the contents of $array into the table
    foreach ($array as $row) {
        echo "<tr>";
        foreach ($row as $cell) {
            echo "<td>" . htmlspecialchars($cell) . "</td>";
        }
        echo "</tr>";
    }
    ?>
</table>

<?php
echo "<br>Total XMASes: " . $xmases;
echo "<br>Total CrossMASes: " . $crossmasses;
?>