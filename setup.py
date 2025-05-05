from setuptools import setup

setup(
    name="gerrit",
    version="0.1.0",
    py_modules=["gerrit"],
    install_requires=[
        "Click",
    ],
    entry_points={
        "console_scripts": [
            "gerrit = gerrit:cli",
        ],
    },
)
