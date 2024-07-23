---
title: "Fix XFCE Desktop Visible After Suspend Before Lock Screen"
date: "2023-09-05"
summary: "Prevent session from being visible when waking from sleep"
description: "Prevent session from being visible when waking from sleep"
toc: true
readTime: true
autonumber: true
math: true
tags: ["linux","debian","xfce","fix"]
showTags: false
---

## The problem

My **Thinkpad X230** running **Debian 12** with **XFCE** + **LightDM** was showing the screen for around 2 seconds before actually locking it.

This issue only happens when waking from sleep by opening the lid.

## The fix

After looking aroud I found a simple fix for this, we just need to specify the graphics driver on the **Xorg** configuration. The solution only works for laptops with **Intel Graphics**.

Create this file `/etc/X11/xorg.conf.d/20-intel.conf` and put the following contents there:

```
Section "Device"
  Identifier "Intel Graphics"
  Driver "Intel"
EndSection
```

The problem should go away after a reboot!

---

<small>Source:</small> `https://www.reddit.com/r/debian/comments/wxgjlq/lightdm_lightlocker_shows_screen_on_resume/`
