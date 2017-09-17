# Grumpy cat

This program is used to kill mouse while you're typing on keyboard. Why? On
KDE slimbooks, touchpad is recognized as a classic mouse, and as such can't be
managed with touchpad system settings, which means it can't use syndaemon to
disable touchpad while typing on keyboard (and this is very annoying).

What grumpy cat is doing:

* monitor your keyboard
* unload mouse module while typingk
* load module again after 750 ms timeout (editable in source, as `timeoutNanoseconds`)

This approach is brutal, it means you keep unloading/loading driver, so it's
clearly not a long term solution.

The problem has been reported to KDE and ubuntu (by slimbook themselves), you
can follow discussion on those pages:

* https://forum.kde.org/viewtopic.php?f=309&t=140461&p=381265#p381265
* https://bugs.launchpad.net/ubuntu/+source/linux/+bug/1688625

If you use grumpy cat, make sure to subscribe to those threads to know when it's
not needed anymore.


## Install

```
go get github.com/oelmekki/grumpy_cat
```

If you're not used to golang, you can find a build for amd64 in the release page.


## Usage

To use grumpy cat, you need to provide the device path for your keyboard and
its module name.

To find your device path:

```
grep keyboard /var/log/Xorg.0.log
```

It will look something like `/dev/input/event0`.

You can list active modules with `lsmod`. If you use a KDE slimbook, it
probably will be `psmouse`.

Finally, you need to run grumpy_cat as root, since it needs to both read
keyboard and manage modules.

Example:

```
sudo grumpy_cat /dev/input/event3 psmouse
```
