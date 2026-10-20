---
Title: EURKey on Ubuntu
Date: Mon, 18 Mar 2019 21:40:00 +0000
---

Hello single person that reads this posts.

Last week I decided to build a desktop computer, this computer will mainly be used for Development and I should not "play" with the system, so I decided to go with Ubuntu, it's not my favourite distro but it's perfect for this use case.

I'm using the default desktop environment (Gnome with Gnome-Shell), it is a bit RAM heavy, but the machine has lot's of RAM, so it's not a big problem.

After installation I dived into the `gnome-control-center`, got into the **Regions and Language**, and I didn't found the EURKey layout, even tough it is installed, I could switch temporarily with `setxkbmap eu` but I didn't wanted that, I was looking to configure it properly.

So it seems that the Gnome folks really like to hide things from us for the sake of **"symplicity"**, so I had to search a bit.

## Here it goes

 1. Make sure the `gnome-control-center` is closed

 2. Run this on your terminal:  
	```gsettings set org.gnome.desktop.input-sources show-all-sources true```

 3. Log out and log back in *(or just restart the gnome-shell)*

 4. Open `gnome-control-center`, go to **Regions and Language**, choose your language and you should find **EurKEY (US Based layout with european letters**.

And it is that simple!
