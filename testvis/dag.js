var dag = {
    "h": 5,
    "w": 3,
    "n": {
        "cmd/smlchk": {
            "n": "cmd/smlchk",
            "x": 2,
            "y": 2,
            "i": [
                "godep"
            ],
            "o": []
        },
        "cmd/smldag": {
            "n": "cmd/smldag",
            "x": 2,
            "y": 4,
            "i": [
                "godep"
            ],
            "o": []
        },
        "godep": {
            "n": "godep",
            "x": 1,
            "y": 2,
            "i": [
                "goload"
            ],
            "o": [
                "cmd/smlchk",
                "cmd/smldag"
            ]
        },
        "goimp": {
            "n": "goimp",
            "x": 0,
            "y": 4,
            "i": [],
            "o": []
        },
        "goload": {
            "n": "goload",
            "x": 0,
            "y": 2,
            "i": [],
            "o": [
                "godep"
            ]
        },
        "~": {
            "n": "~",
            "x": 0,
            "y": 0,
            "i": [],
            "o": []
        }
    }
};