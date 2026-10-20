---
Title: Sharing keyboard and mouse over network
Date: Tue, 14 Jul 2015 11:50:00 +0000
---

I wanted to use both my laptop and the desktop simultaneously, It was really "painfull" to use two keyboards, so I searched for some kind of software that let me use **one** keyboard and mouse in **multiple** computers at the same time.

The solution already exists and it's called **Synergy**!

"Synergy combines your desktop devices together in to one cohesive experience. It's software for sharing your mouse and keyboard between multiple computers on your desk. It works on Windows, Mac OS X and Linux."

For **synergy** to work you need one computer to work as a server (the computer in which the keyboard and mouse are connnected), the other computers will connect to the server and use the shared keyboard and mouse.

## Server Configuration

### Naming your computers
First you need to give a name to your clients in the `/etc/hosts` file, like this:

```text
#<ip-address>	<hostname.domain.org>	<hostname>
127.0.0.1	localhost.localdomain	localhost	custompc
::1		localhost.localdomain	localhost	custompc

192.168.1.51	laptoppc
```

 * The server is named **custompc** and the client is named **laptoppc**


### Synergy Configuration
The configuration file for synergy is located in `/etc/synergy.conf`, create it if you don't have it:

```ini
section: screens
	custompc:
	laptoppc:
end

section: links
	#  The laptop is on the right of the desktop screen
	custompc:
		right = laptoppc

	laptoppc:
		left  = custompc

end

section: options
	relativeMouseMoves = false
	screenSaverSync = true
	win32KeepForeground = false
	switchCorners = none 
	switchCornerSize = 0
	#switch computers with a keybinding
	keystroke(Control+Alt+l) = switchInDirection(right)
	keystroke(Control+Alt+h) = switchInDirection(left)
end
```

Now we just need to run the synergy server:

```sh
synergys -f
```


## Client Configuration

There's not much to do in the client, you just need to to run it:

```sh
synergyc 192.168.2.50
```

 * I used 192.168.2.50, you need to replace it with your server address.
 

That's all, if you have some suggestions or corrections feel free to contact me.
