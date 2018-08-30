package main

var setupText = "import setuptools\n" +
	"with open(\"README.md\", \"r\") as fh:\n" +
	"    long_description = fh.read()\n" +
	"setuptools.setup(\n" +
	"    name=\"%s\",\n" +
	"    version=\"%s\",\n" +
	"    author=\"%s\",\n" +
	"    author_email=\"%s\",\n" +
	"    description=\"%s\",\n" +
	"    long_description=\"%s\",\n" +
	"    long_description_content_type=\"text/markdown\",\n" +
	"    url=\"%s\",\n" +
	"    packages=setuptools.find_packages(),\n" +
	"    classifiers=(\n" +
	"         \"Programming Language :: Python :: 3\",\n" +
	"         \"License :: OSI Approved :: MIT License\",\n" +
	"         \"Operating System :: OS Independent\",\n" +
	"	 ),\n" +
	")"
