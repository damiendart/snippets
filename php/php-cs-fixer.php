<?php

/*
 * Copyright (C) 2022 Damien Dart, <damiendart@pobox.com>.
 * This file is distributed under the MIT licence. For more information,
 * please refer to the accompanying "LICENCE" file.
 */

/*
 * A generic ".php-cs-fixer.php" containing a bunch of house rules for
 * personal projects using PHP 8 and above.
 */

declare(strict_types=1);

$header = <<<'HEADER'
Copyright (C) 2022 Damien Dart, <damiendart@pobox.com>.
This file is distributed under the MIT licence. For more information,
please refer to the accompanying "LICENCE" file.
HEADER;

return (new PhpCsFixer\Config())
    ->setRules(
        [
            '@PhpCsFixer' => true,
            '@PSR12' => true,
            'array_syntax' => ['syntax' => 'short'],
            'concat_space' => ['spacing' => 'one'],
            'declare_strict_types' => true,
            'header_comment' => [
                'header' => $header,
                'location' => 'after_open',
                'separate' => 'both',
            ],
            'no_unused_imports' => true,
            'multiline_whitespace_before_semicolons' => [
                'strategy' => 'no_multi_line',
            ],
            'ordered_imports' => ['sort_algorithm' => 'alpha'],
            'trailing_comma_in_multiline' => [
                'elements' => ['arrays', 'arguments', 'parameters'],
            ],
            'void_return' => true,
        ],
    )
    ->setFinder(PhpCsFixer\Finder::create()->in(__DIR__)->name('*.php'));
