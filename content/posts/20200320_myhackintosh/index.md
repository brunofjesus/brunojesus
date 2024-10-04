---
title: "My Hackintosh (Catalina)"
date: "2020-03-20"
summary: "How I built my hackintosh"
description: "Description and instructions to build / replicate my hackintosh"
toc: true
readTime: true
autonumber: true
math: true
tags: ["macos","hardware"]
showTags: false
---

Since I started working on my new job, in which I was given a MacBook, I started to struggle when switching between 
computers with other operating systems, mostly because of different key shortcuts (looking at you IntelliJ).
Also I enjoy using macOS, so I decided to replace my Linux distro for macOS on my home machine.

## Hardware

- **Motherboard**: [MSI Bazooka H370M](https://www.msi.com/Motherboard/H370M-BAZOOKA.html)
- **Processor**: [Intel® Core™ i5-8400](https://ark.intel.com/content/www/br/pt/ark/products/126687/intel-core-i5-8400-processor-9m-cache-up-to-4-00-ghz.html)
- **Disk (NVMe)**: [Samsung SSD 970 EVO Plus 500GB](https://www.samsung.com/us/computing/memory-storage/solid-state-drives/ssd-970-evo-plus-nvme-m-2-500gb-mz-v7s500b-am/)
- **Graphics**: [Sapphire Radeon RX 580 NITRO+ 4GB OC](https://www.sapphiretech.com/en/consumer/nitro-rx-580-4g-g5)
- **Memory**: 32GB 4x8GB [Aegis DDR4-3000MHz CL16-18-18-38 1.35V](https://www.gskill.com/product/165/185/1536026353/F4-3000C16D-16GISBAegis-DDR4DDR4-3000MHz-CL16-18-18-38-1.35V16GB-(2x8GB))
- **PSU**: [Cooler Master MasterWatt Lite 600W 80 PLUS](https://www.coolermaster.com/catalog/power-supplies/masterwatt-lite-full-range/masterwatt-lite-600w-full-range/)

## Requirements
The latest BIOS versions aren't working (at least for the installation). I get it to work with the version 7B24v10.
You can dowload it [here](https://download.msi.com/bos_exe/mb/7B24v10.zip).

The NVMe memory does not work well with some older firmware, I needed to get the firmware version 2B2QEXM7 to be able to,
install macOS. You can get that firmware [here](https://s3.ap-northeast-2.amazonaws.com/global.semi.static/SAMSUNG_SSD_970EVO_Plus_ISO_190513/7ES357322A6707A720E1A71EF11A3BE1EED819E011D317626415F0281A78151C/Samsung_SSD_970_EVO_Plus_2B2QEXM7.iso),
or [here](https://www.tonymacx86.com/attachments/samsung_ssd_970_evo_plus_2b2qexm7-iso.407927/).

EFI folder with required drivers and kernel extensions. You can get it [here](EFI.zip).

Access to a device running macOS. If you don't have on you can install it in a virtual machine using Linux.
More information [here](https://github.com/foxlet/macOS-Simple-KVM).

Unibeast for creating the flash drive. You can download it [here](https://www.tonymacx86.com/resources/unibeast-10-0-0-catalina.448/).

## How to proceed

### Create a Bootable USB
 - Download macOS Catalina from AppStore
 - Format your USB drive as Mac OS Extended (Journaled) using Disk Utility
 - Create a UEFI bootable disk with Unibeast
 - Replace the EFI folder with the one you downloaded.

### Disconnect all unneded peripherals from the computer
This included external disks, other flashdrives, wifi and bluetooth dongles, etc.

### BIOS Setup
**Settings -> Advanced**
- **Super IO Configuration**: Disabled
- **USB Configuration**: Enable all
- **Integrated Graphics Configuration**
  - **Initiate Graphic Adapter**: PEG
  - **IGD Multi-Monitor**: Disabled
  
**Settings -> Boot**
- Boot configuration
  - **Boot mode select**: UEFI
- Fixed boot order
  - **Boot option #1**: UEFI USB Key
  
**OC SET-UP**
- Overclocking
  - **Extreme Memory Profile (XMP)**: Enabled
  - CPU Features
    - **LIMIT CPUID Maximum**: Disabled
    - **Intel VT-D Tech**: Disabled
    - **C1E Support**: Disabled
    
### Installation

- Format your hard drive as APFS and use GUID Partition Map.
- Your computer will reboot, when it boots to Clover select **'Boot macOS Install Prebooter from Preboot'**.
- Your computer will reboot again, now you can select **'Boot macOS from "your hard disk name"'**.
- If all goes well you will boot into macOS.
- Download Clover Configurator [here](https://mackie100projects.altervista.org/download-clover-configurator/).
- Open Clover Convigurator, mount EFI, and mount the EFI partition on from your hard drive and from your USB stick.
- Copy your hard drives's EFI folder to your disk's EFI partition.
- Turn off your computer, unplug the USB drive and boot from your disk!
- Enjoy!


---

The information for this article was based on the following guides:
 - https://www.tonymacx86.com/threads/bazookamac-1-3-i5-8400-h370m-bazooka-vega-56-catalina-compatible.266512/
 - https://www.tonymacx86.com/threads/success-portable-desktop-i7-8086k-msi-h370m-bazooka-2-radeon-rx580-nvme-pcie-m-2-startup.279704/
 - Some more info that I got while troubleshooting
