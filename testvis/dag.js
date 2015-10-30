var dag = {
    "height": 7,
    "width": 11,
    "nodes": {
        "assert.g": {
            "x": 2,
            "y": 1,
            "ins": [],
            "outs": [
                "queue.g"
            ]
        },
        "cond.g": {
            "x": 6,
            "y": 4,
            "ins": [
                "lock.g"
            ],
            "outs": [
                "join.g",
                "mailbox.g"
            ]
        },
        "context.g": {
            "x": 1,
            "y": 3,
            "ins": [
                "int_frame.g",
                "page_table.g"
            ],
            "outs": [
                "thread.g"
            ]
        },
        "int_frame.g": {
            "x": 0,
            "y": 2,
            "ins": [],
            "outs": [
                "context.g"
            ]
        },
        "int_handler.g": {
            "x": 9,
            "y": 2,
            "ins": [
                "sched.g"
            ],
            "outs": [
                "main.g"
            ]
        },
        "interrupt.g": {
            "x": 3,
            "y": 4,
            "ins": [],
            "outs": [
                "sched.g"
            ]
        },
        "join.g": {
            "x": 7,
            "y": 5,
            "ins": [
                "cond.g"
            ],
            "outs": [
                "threading.g"
            ]
        },
        "lock.g": {
            "x": 5,
            "y": 4,
            "ins": [
                "sched.g"
            ],
            "outs": [
                "cond.g"
            ]
        },
        "mailbox.g": {
            "x": 8,
            "y": 4,
            "ins": [
                "cond.g"
            ],
            "outs": [
                "test_init.g"
            ]
        },
        "main.g": {
            "x": 10,
            "y": 2,
            "ins": [
                "int_handler.g",
                "print_str.g",
                "test_init.g"
            ],
            "outs": []
        },
        "page_table.g": {
            "x": 0,
            "y": 4,
            "ins": [],
            "outs": [
                "context.g"
            ]
        },
        "print_str.g": {
            "x": 9,
            "y": 0,
            "ins": [],
            "outs": [
                "main.g"
            ]
        },
        "queue.g": {
            "x": 3,
            "y": 2,
            "ins": [
                "assert.g",
                "thread.g"
            ],
            "outs": [
                "queue_test.g",
                "sched.g"
            ]
        },
        "queue_test.g": {
            "x": 4,
            "y": 1,
            "ins": [
                "queue.g"
            ],
            "outs": []
        },
        "sched.g": {
            "x": 4,
            "y": 3,
            "ins": [
                "interrupt.g",
                "queue.g"
            ],
            "outs": [
                "int_handler.g",
                "lock.g",
                "sem.g"
            ]
        },
        "sem.g": {
            "x": 8,
            "y": 2,
            "ins": [
                "sched.g"
            ],
            "outs": [
                "test_init.g"
            ]
        },
        "test_init.g": {
            "x": 9,
            "y": 4,
            "ins": [
                "mailbox.g",
                "sem.g",
                "threading.g"
            ],
            "outs": [
                "main.g"
            ]
        },
        "thread.g": {
            "x": 2,
            "y": 3,
            "ins": [
                "context.g"
            ],
            "outs": [
                "queue.g"
            ]
        },
        "threading.g": {
            "x": 8,
            "y": 6,
            "ins": [
                "join.g"
            ],
            "outs": [
                "test_init.g"
            ]
        }
    }
};