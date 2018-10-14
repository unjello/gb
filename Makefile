.PHONY: all upload build clean inst uninst

all: upload
upload:
	twine upload dist/*
build:
	python setup.py sdist bdist_wheel
clean:
	rm -rf build/ dist/ generate_build.egg-info
inst:
	pip install --user dist/gb-`cat generate_build/__init__.py | cut -d\' -f 2`-py3-none-any.whl
uninst:
	pip uninstall -y generate_build
