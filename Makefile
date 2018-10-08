.PHONY: all upload build clean inst uninst

all: upload
upload:
	python setup.py sdist bdist_wheel upload
build:
	python setup.py sdist bdist_wheel
clean:
	rm -rf build/ dist/ gb.egg-info
inst:
	pip install --user dist/gb-`cat gb/__init__.py | cut -d\' -f 2`-py2.py3-none-any.whl
uninst:
	pip uninstall -y gb
