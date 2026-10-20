---
Title: Reading feeds in the terminal
Date: Fri, 25 Dec 2016 14:30:00 +0000
---

I live in the terminal and I also like to read my feeds there.

For a long time I used [newsbeuter](http://newsbeuter.org/), it was ok, but I didn't like to much and I also wanted to learn how to create nCurses programs, that's why I created [simpleRSS](https://gitlab.com/strang3quark/simpleRSS).

It's far from finished and certanly far from perfect but it's pretty usable right now.

The feeds can be organized by categories and you can switch through the menus with h,j,k,l just like in vim.

**Feeds Screen**

![Feeds Screen](/assets/media/20161225_readingfeedsterminal/feedlist.png)


**Articles List**

![Articles List](/assets/media/20161225_readingfeedsterminal/articlelist.png)


**Article Reading / Links**

![Article Reading](/assets/media/20161225_readingfeedsterminal/articlereading1.png)
![Article Links](/assets/media/20161225_readingfeedsterminal/articlereading2.png)

I plan to support image display with `w3mimgdisplay` like [ranger](http://ranger.nognu.org):

![Ranger Image Support](http://ranger.nongnu.org/screenshots/w3mimgpreview.png)

But before that I still need to do a major code revision (that code is a mess) and also implement some features like: 

 - Logs (for debug mode)
 - Bookmarks
 - Export articles to Markdown or HTML.

If you are interested in using it or contributing take a look [here](https://github.com/strang3quark/simpleRSS).
