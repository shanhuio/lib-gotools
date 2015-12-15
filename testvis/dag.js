var dag = {
    "assert": [],
    "cond": [
        "sched",
        "import",
        "thread",
        "lock",
        "assert"
    ],
    "import": [],
    "init": [
        "import",
        "sched",
        "ptable",
        "int_handler"
    ],
    "int_handler": [
        "import",
        "sched",
        "threading",
        "assert"
    ],
    "join": [
        "lock",
        "cond",
        "thread",
        "sched",
        "assert",
        "import"
    ],
    "lock": [
        "thread",
        "sched",
        "assert",
        "import"
    ],
    "mailbox": [
        "lock",
        "cond",
        "assert"
    ],
    "ptable": [
        "import"
    ],
    "queue_test": [
        "thread",
        "assert"
    ],
    "run": [
        "sched",
        "init",
        "threading",
        "assert",
        "import"
    ],
    "sched": [
        "thread",
        "import",
        "assert",
        "ptable"
    ],
    "sem": [
        "import",
        "sched",
        "assert",
        "thread"
    ],
    "thread": [
        "import",
        "assert"
    ],
    "threading": [
        "import",
        "ptable",
        "sched",
        "join",
        "thread",
        "assert"
    ],
    "whale": [
        "lock",
        "cond",
        "assert"
    ]
};