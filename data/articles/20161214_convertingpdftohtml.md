---
Title: Converting PDF files to HTML
Date: Fri, 14 Dec 2016 16:45:00 +0000
---

It's a shame that we don't still have a decent PDF reader in Linux, I was in need of a PDF reader that let me copy and paste without messing all the formatting (the Firefox PDF reader didn't worked well) and I didn't want to install a PDF reader with lots of dependencies, I like things simple, that's why I don't use any Desktop Environment, just i3-wm, simple applications, and scripts (lots of scripts). 

I've tried XPDF, it's a decent PDF reader but the copy and paste operation is a little weird.

I was tired of PDF readers and I needed a solution, that's when I found **poppler**, it's a PDF renderer who let's you convert PDF files to other formats.

You can install **poppler** package in Arch Linux with `pacman -S poppler`.

I've made a simple script that uses poppler to convert the file to HTML, saves it in `/tmp/<filename>` and then opens it in your default browser.

```sh
#! /bin/bash
# convert_pdftohtml.sh
# Copyright (C) 2016 Bruno Jesus (aka strang3quark) <bruno.fl.jesus@gmail.com>
#
# Distributed under terms of the MIT license.
#

PDFPATH=$1;
PDFFILE=$(basename $PDFPATH);


if [ "$2" == "--format" ]; then
	FORMAT="-s";
	HTMLFILE="index-html.html";
else
	FORMAT="";
	HTMLFILE="index.html";
fi

mkdir /tmp/$PDFFILE;

pdftohtml -p $FORMAT $PDFPATH /tmp/$PDFFILE/index.html

$BROWSER /tmp/$PDFFILE/$HTMLFILE
```


----------


**The usage is very simple:**

```sh
convert_pdftohtml.sh myfile.pdf # this will remove all the weird formatting
convert_pdftohtml.sh myfile.pdf --format # this will keep all the formatting
```

If you have some alternatives or suggestions please contact me.
