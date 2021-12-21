<?php
echo date($_GET['format'] ?? $argv[2],$_GET['timestamp']??$argv[1]);