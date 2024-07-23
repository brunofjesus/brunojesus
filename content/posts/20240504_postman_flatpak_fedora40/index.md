---
title: "Fix Postman Flatpak not opening after login"
date: "2023-09-05"
summary: "Make Postman run on Fedora 40"
description: "Make Postman run on Fedora 40"
toc: true
readTime: true
autonumber: true
math: true
tags: ["linux","fedora","flatpak","fix"]
showTags: false
---

## The problem

While setting up my new personal laptop with Fedora 40 I noticed that the Postman's flatpak version refused to start 
after I logged in, after running it from the command line with `flatpak run com.getpostman.Postman`  I noticed the following error:

```
Main~handleUncaughtError - Uncaught errors [Error: ENOENT: no such file or directory, open '/home/user/.var/app/com.getpostman.Postman/config/Postman/proxy/postman-proxy-ca.crt'] {
  errno: -2,
  code: 'ENOENT',
  syscall: 'open',
  path: '/home/user/.var/app/com.getpostman.Postman/config/Postman/proxy/postman-proxy-ca.crt'
}

```

## The fix

After searching through the comments on Gnome Software I found a comment from Alexandre DeBusschere that solved the issue:
 - Go to `~/.var/app/com.getpostman.Postman/config/Postman/proxy`
 - Run the following command to generate a valid certificate:

   `
openssl req -subj '/C=US/CN=Postman Proxy' -new -newkey rsa:2048 -sha256 -days 365 -nodes -x509 -keyout postman-proxy-ca.key -out postman-proxy-ca.crt
   `

The application should work just fine now!
