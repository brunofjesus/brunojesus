---
title: "Ubuntu on HP Pro C640 Chromebook"
date: "2022-05-17"
summary: "Hacking a Chromebook in order to run Ubuntu on it"
description: "Hacking a Chromebook in order to run Ubuntu on it"
toc: true
readTime: true
autonumber: true
math: true
tags: ["chromebook","linux","ubuntu","hardware"]
showTags: false
---

## Introduction

Some time ago I got a cheap used HP Pro C640 Chromebook, although it is capable of running **Crostini**
I wasn't a huge fan of being locked to a Google account, so I decided to hack  the shit out of it and install a
**Linux Distro**, and after some trial and error I was successful!

![C640 Running Ubuntu](c640.jpeg)

## Disclamer

This guide will convert your Chromebook into a "generic" machine, this means that **you will no longer be able to run ChromeOS after you flash the MrChromebox firmware**.

All of your data will be erased, make sure to backup all of your data before following this guide.

The steps listed in this guide are the ones I followed in order to install **Ubuntu** on my **HP Pro C640 Chromebook**, 
the fact that it worked for me is no guarantee that it will work on your machine as well!

## Requirements
- An HP Pro C640 Chromebook
- 2 flash drives (one for the firmware backup and another one with Ubuntu 22.04)
- A screwdriver and some old plastic card (or similar), for the back cover removal
- Time and patience

## Enable development mode

Before flashing the new firmware you need to gain root access to the ChromeOS system, you can do that by enabling **Developer Mode**.
Enabling this will wipe all your data, so make sure you have a backup.

 - Shut down the Chromebook
 - Press and hold the **esc**, **refresh** and **power** keys at the same time until a white screen saying "**Chrome OS is missing or damaged**" appears.
 - Press **Ctrl** and **D** simultaneously in order for the boot process to continue.
 - You will get a message saying **OS verification is OFF**, press **Ctrl** and **D** simultaneously in order for the boot process to continue.
 - The system will boot into Chrome OS and ask you for your google credentials. **DO NOT LOGIN**

## Disable the firmware write protection

Usually Chromebooks do not allow you to write in the firmware flash memory without messing with hardware, the C640 is no exception,
you will need to disconnect the battery in order to remove the firmware write protection.

This official guide gives an explanation on how to do it (removing the back cover was not an easy task on my machine!)

[![HP C640 Teardown](https://img.youtube.com/vi/CTN52CBXzaE/0.jpg)](https://www.youtube.com/watch?v=CTN52CBXzaE)

## Flash MrChromebox firmware

In order to be able to boot Linux operating systems we need to install a custom firmware, I used **MrChromebox** for that.

You can check [this page](https://mrchromebox.tech/#fwscript) for more info.

- In the **Chrome OS** setup screen (after boot), press **Ctrl** + **Alt** + **F2** (Third button of the top row / Forward arrow) simultaneously.
- Login as `chronos`, it should have no password

- Run the following commands (press enter after each one):
```sh
cd; curl -LO mrchromebox.tech/firmware-util.sh
sudo install -Dt /usr/local/bin -m 755 firmware-util.sh
sudo firmware-util.sh
```

You should see something similar to the image bellow, this image is not of my computer (sorry, forgot to take a picture):

![MrChromebox](mrchromebox.png)

---

If the "[WP]" is red instead of green it means that the firmware flashing is locked. See *[Disable the firmware write protection]*.

---

- Select option 2
  - **After this step you can no longer go back to ChromeOS, make sure you backup you firmware (the script does that as long as you have a flash drive)**

- After you reboot your computer it should no longer be able to boot into ChromeOS

## Install Ubuntu

You can install any system you want, but I cannot assure you that it will work or that the further steps also work or are enough to have a stable environment.

- Create an [Ubuntu 22.04](https://ubuntu.com/download/desktop/thank-you?version=22.04&architecture=amd64) installation flash drive, you can use [etcher](https://www.balena.io/etcher/) for that.

- Insert the flash and turn on the Chromebook, when it turns on press **esc**, open the **Boot menu** and then select your **usb flash drive name**.

- Install Ubuntu like you would do on any regular computer.

- **Do not forget to install all the updates on the first boot.**

## Fix sound

By default, the sound should only be coming from the headphone jack, to fix that we need to specify the default device.

- Login into Ubuntu

- Open a terminal **Ctrl** + **Alt** + **T**

- Run this in the terminal `sudo nano /etc/pulse/default.pa`

- Add the following at the end of the file
```shell
# Chromebook configuration
load-module module-alsa-sink device=hw:0,0
load-module module-alsa-sink device=hw:0,5
load-module module-alsa-sink device=hw:0,3
load-module module-alsa-source device=hw:0,1
set-default-sink 1
set-default-source 3
update-sink-proplist alsa_output.pci-0000_00_1f.3-platform-sof_rt5682.stereo-fallback device.description="Headphones"
update-sink-proplist alsa_output.hw_0_5 device.description="Internal speakers"
update-sink-proplist alsa_output.hw_0_3 device.description="HDMI"
```

- To save the file press **Ctrl** + **X**, then it should ask for saving the file, press **Y** and then **Enter**

- I noticed that by some reason this settings are not loaded automatically during startup, I found a little workaround for that
  
  - Press the **Everything Key (The magnifier key)** and write **Startup**, then open the **Startup Applications** app.

  - Click add, and **add** the following:
    ![Restart pulseaudio](restart_pulseaudio.png)

  - **(Optional)** Add the login sound:
    ![Login sound](login_sound.png)
    - `canberra-gtk-play --id="system-ready"`

- Reboot
 
## Fix suspension

I noticed that the computer would not suspend properly, it would just immediately wake up from sleep, the following steps fixed it.

- In the terminal execute `sudo nano /etc/default/grub`
- On the `GRUB_CMDLINE_LINUX_DEFAULT` add `mem_sleep_default=deep`
  - It should look something like `GRUB_CMDLINE_LINUX_DEFAULT="quiet splash mem_sleep_default=deep"`
- Save the file

- Execute `sudo update-grub`

- Reboot, everything should work fine now.

## Final notes
With all those steps everything should work as intended.

Feel free to contact me if you have any suggestion to further improve this article.

If you have any doubt I suggest looking at [/r/chrultrabook/ subreddit](https://www.reddit.com/r/chrultrabook/).

### Sources
 - https://mrchromebox.tech/#fwscript
 - https://www.reddit.com/r/chrultrabook/comments/umi82r/comet_lake_audio_sofrt5682_and_deep_sleep/
