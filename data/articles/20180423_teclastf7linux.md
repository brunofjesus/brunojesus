---
Title: Running Linux on Teclast F7
Date: Sat, 23 Apr 2018 20:00:00 +0000
---

Recently I aquired a Teclast F7 Laptop [(review here)](20180324_teclastf7), one of the main problems of this laptop was the touchpad problem on Linux, **but there's a fix now!**.

---

**UPDATE:** Starting with kernel 4.20 you don't need to manually install the patch to make the touchpad work, you can install the latest kernels easily on ubuntu with [**ukuu**](https://github.com/teejee2008/ukuu).

---

I've got Xubuntu 18.04 running sucessfully on this laptop, the steps I'll show here will also work on other Ubuntu flavors.

## Getting ready
For this fix you need to compile the kernel, you may want to do this on another machine, you will need around 25GB of free space and this is a CPU intensive task. I've got no issues compiling on the teclast f7, but it is a slow process.

You need to install Xubuntu 18.04 (or other Ubuntu flavor), you can write the image on a USB drive and you are good to go.

**Do not mess with the BIOS, keep Windows as OS on the BIOS settings, otherwhise the brightness will not work correctly**

## Getting the Kernel and the Patch

Download the kernel from the source repository:

```sh
sudo apt-get source linux-image-$(uname -r)
```

Download the driver with the touchpad workaround:

```sh
wget https://raw.githubusercontent.com/brotfessor/sipodev/master/b/i2c-hid.c
```

Now you got the kernel sources and the patched driver, to patch the kernel simply replace the `i2c-hid.c` from `~/linux-4.15.0/drivers/hid/i2c-hid/` with the one you just downloaded.

## Build the kernel

First you need to dowload the required dependencies to build the kernel, you can get them with this command:

```sh
sudo apt-get build-dep linux-image-$(uname -r)
```

Now get into the kernel directory and run the following commands:

```sh
fakeroot debian/rules clean
fakeroot debian/rules binary-headers binary-generic binary-perarch
```

If you are building this from the Teclast F7 you can do something else for at least a couple of hours (I left mine overnight).

## Installing the kernel

If all went well you have sucessfully compiled the kernel! Now it's time for the installation.
The packages should be on your home folder, change to the home directory and execute the following command:

```sh
sudo dpkg -i linux*.deb
```

Now it's time to update grub and reboot:

```sh
sudo update-grub
systemctl reboot
```


## Prevent the system from updating the kernel

Since this patch is not official everytime you update the kernel from the repos you will be left with no touchpad, to fix this run the following commands:

```sh
sudo apt-mark hold linux-headers-generic linux-image-generic linux-headers-$(uname -r) linux-image-$(uname -r) linux-image-extra-$(uname -r)
```

**Important: you should repeat this process if you want to update the kernel**
