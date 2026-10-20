---
Title: Posts are Back!
Date: Sun, 01 Nov 2026 20:00:00 +0100
---

Before having `brunojesus.pt`, there was `strang3quark.tk` (and also
`strang3quark.cf` when I lost the `tk` domain). At that time I was still
studying and decided to create a site from scratch inspired by some other sites
from folks at [nixers.net](https://nixers.net).

After finishing my studies, I abandoned `strang3quark.tk`. Later I created
`brunojesus.pt`, and since I wanted it to be a bit more professional, I
ended up discarding the old site and its contents.

![Screenshot: strang3quark.tk V1](/assets/media/20261101_posts_are_back/v1.png)
<small>Image: strang3quark.tk V1</small>

Since I usually have nothing interesting to post about, I ended up removing
the blog section from `brunojesus.pt`, and eventually it became just a front
page with a link to my CV.

Today that changes. I picked most of the `strang3quark.tk` content and moved
it here, and I also restored the posts I had removed from `brunojesus.pt`.
This is not just a restoration of old content, it's also a full rewrite of
the site.

As you might notice, this site looks a bit similar to the first iteration of
`strang3quark.tk`. I deliberately kept the same structure, with a few changes:
 - Switched to the [Catppuccin](https://catppuccin.com) Frappe and Latte color
 schemes, the same ones I use on my system.
 - Rewrote the entire site in Go. I previously used [Hugo](https://gohugo.io/),
 and I have no good reason for this change other than I wanted to do it 🤠

## How does this new site work?

If you take a look at the [source code](https://github.com/brunofjesus/brunojesus)
you'll notice that articles live in the `data/articles` folder. They are
written in `Markdown` and then converted to `HTML` by the
[generator](https://github.com/brunofjesus/brunojesus/blob/36665f38b13e6bf7474bdae792601b18d32a0d0b/cmd/generator/main.go),
which produces fully fledged HTML files. So, it generates a fully static website!

Then there's also the
[server](https://github.com/brunofjesus/brunojesus/blob/36665f38b13e6bf7474bdae792601b18d32a0d0b/cmd/server/main.go),
a simple `HTTP` server that serves the generated `HTML`. It's optional, I don't
strictly need it, but I'll use it for now.

## Was AI used?
All of the Go code was written by me (maybe the AI would have done a better
job 🤷‍♂️), however, the CSS changes were done by the AI, as well as the
conversion of the old Markdown and text files so they have the correct tags
and are properly parsed.

Still, I also leveraged AI to ask some questions about the code I was making
and used it to troubleshoot some stuff along the way, but minimal, nonetheless.

About the actual writing of articles, I'm really bad at that, but I'll try to
use AI only for proofreading and to do some minor optimizations.

## So, what's next?
Some of the old articles are outdated while others contain some errors. I'm
not planning to do much about the outdated ones, the publication dates are
public, so anyone who stumbles on them will know how old they are.

As for the wrong information, I will fix it and add notes at the end of each
affected article about the changes made over time.

I also have some ideas for stuff to post in the future, but I won't promise
I'll post often.

## What about the old site?
The source code for this site and the old `strang3quark.tk` is public, so
feel free to explore:
- [strang3quark.tk source code](https://gitlab.com/strang3quark/strang3quark-website)
- [brunojesus.pt source code](https://github.com/brunofjesus/brunojesus)
