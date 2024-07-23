---
title: "Apache Guacamole on FreeNAS 11.3"
date: "2020-05-29"
summary: "Host your own web RDP and VNC client on FreeNAS"
description: "Host your own web RDP and VNC client on FreeNAS"
toc: true
readTime: true
autonumber: true
math: true
tags: ["homelab","nas", "selfhost"]
showTags: false
---

## Description

I'm running on some VMs on a FreeNAS machine which I want to access via multiple protocols, mainly RDP, the problem is,
sometimes I may not have an RDP client configured, that's where [Apache Guacamole](https://guacamole.apache.org) comes
in handy.

I created a jail in FreeNAS to install Guacamole on it, but this should work without any problems on a FreeBSD system.

## Installation

 - Install the packages
 
   ```sh
   pkg update
   pkg install guacamole-client guacamole-server
   ```
   
 - Copy the provided `user-mapping.xml.sample`
 
   ```sh
   cp /usr/local/etc/guacamole-client/user-mapping.xml.sample /usr/local/etc/guacamole-client/user-mapping.xml
   ```

 - Edit the file with your favourite text editor (e.g ```vim```, ```vi```), it should look something like this without the comments

   ```xml
   <?xml version="1.0" encoding="UTF-8"?>
   <user-mapping>
       <authorize username="youruser" password="yourpassword">
       
           <connection name="My VM 1">
               <protocol>rdp</protocol>
               <param name="hostname">192.168.1.190</param>
               <param name="port">3389</param>
               <param name="username">vm-user</param>
               <param name="ignore-cert">true</param>
           </connection>
       </authorize>
   </user-mapping>
   ```
        
   - This is just an example, please change your username, password and add your own collections, you can add as many as you like.
    
    
 - Start the services on startup
 
   ```sh
   echo 'guacd_enable="YES"' >> /etc/rc.conf
   echo 'tomcat9_enable="YES"' >> /etc/rc.conf
   ```

            
 - Restart the system/jail and then go to your browser and navigate to ```http://your-machine-ip:8080/guacamole```
