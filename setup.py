# Always prefer setuptools over distutils
from setuptools import setup, find_packages
from os import path
from io import open
from generate_build import __version__ as gb_version

here = path.abspath(path.dirname(__file__))

with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

def get_requires(filename):
    requirements = []
    with open(filename, "rt") as req_file:
        for line in req_file.read().splitlines():
            if not line.strip().startswith("#"):
                requirements.append(line)
    return requirements


project_requirements = get_requires("generate_build/requirements.txt")
dev_requirements = get_requires("generate_build/requirements_dev.txt")

setup(
    name='generate_build',
    version=gb_version,
    description='Yet another build generator for C++',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/unjello/gb',
    author='Andrzej Lichnerowicz',
    author_email='andrzej@lichnerowicz.pl',
    classifiers=[  
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'Topic :: Software Development :: Build Tools',
        'License :: CC0 1.0 Universal (CC0 1.0) Public Domain Dedication',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.6',
    ],
    keywords=['C/C++', 'tool', 'c', 'c++', 'cpp', 'build', 'developer'],
    packages=find_packages(exclude=['contrib', 'docs', 'tests']),
    install_requires=project_requirements,
    extras_require={  
        'dev': dev_requirements,
        'test': dev_requirements
    },
    package_data={
        '': ['LICENSE'],
    },
    entry_points={
        'console_scripts': [
            'gb=generate_build.gb:run',
        ],
    },
    project_urls={
        'Bug Reports': 'https://github.com/unjello/gb/issues',
        'Source': 'https://github.com/unjello/gb/',
    },
)
