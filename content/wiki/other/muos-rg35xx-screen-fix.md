# MuOS RG35XX Plus Screen Issue Fix

## The Issue

On the newer revisions of the RG35XX devices from Anbernic, MuOS Pixie has an issue with the screen timings, causing horizontal lines to appear on the bottom of the screen over time

* [Example 1](https://ibb.co/1GxjqSCd)
* [Example 2](https://www.reddit.com/media?url=https%3A%2F%2Fpreview.redd.it%2F504177hj16ke1.jpg%3Fwidth%3D320%26crop%3Dsmart%26auto%3Dwebp%26s%3D4af0907a67bf38dfc6977cd23dad085c68181c36)

The issue is more noticeable on high brightness levels, and gets worse the longer the device is used for  
It seems to occur on panel with the model number 355608-A9

## The Fix

* DISCLAIMERS: 
	* I have only tested this on the RG35XX Plus
	* I make no guarantees that this guide works
	* Editing these files can potentially damage the screen
	* If you enter commands incorrectly, you could overwrite the wrong drive
	* Or possibly corrupt your MuOS install
		* I recommend backing it up with `dd` first (`sudo dd if=/dev/sdX bs=1M status=progress of=backup.iso`)

The fix is to change the screen timings. By default MuOS sets the screen refresh rate to 59.98Hz. This guide changes that to match Knulli, which runs the screen at 59.94Hz. In theory this introduces more stutter, but it is unlikely to be noticeable.

### Updated Windows Methods

GitHub user andromalandro made a program to apply the patch on a Windows machine: [https://github.com/andromalandro/FexFlasher](https://github.com/andromalandro/FexFlasher)

### Linux Version

1. Take the SD card out of the handheld and put in in your computer
2. Mount `/dev/sdX5` where X is the device letter of the SD card
3. Backup the folder `mountPoint/opt/muos/device/rg35xx-plus/package/` (replace `rg35xx-plus` with your device)
4. Navigate to `mountPoint/opt/muos/device/rg35xx-plus/package/` (replace `rg35xx-plus` with your device)
5. Open the `sunxi.dts` file in a text editor, and find the section labelled `lcd0@01c0c000`
6. Replace that entire section (from the opening `{` to the closing `}`) with the text from the section below
7. Save and exit the file
8. Run `dtc -I dts -O dtb -o sunxi.fex sunxi.dts`
9. Run `~/Downloads/dragonsecboot -pack boot_package.cfg`  
	9.1. This program can be found here: [https://github.com/MustardOS/tool/blob/main/dragonsecboot](https://github.com/MustardOS/tool/blob/main/dragonsecboot)  
	9.2. Download this file and make it executable with `chmod +x dragonsecboot`  
10. Copy the `boot_package.fex` file to another place on your computer (`~/Downloads` for example)
11. Unmount the SD card, but leave in computer (`cd ~`, `sudo umount /dev/sdX5`)
12. Run `sudo dd if=boot_package.fex of=/dev/sdX bs=1k seek=16400`
	* IMPORTANT: Make sure you ge the right `sdX` device
	* IMPORTANT: Make sure you enter the correct seek offset `16400`
13. Remove the SD card, place back in device and boot it up

## Text to Replace

[On GitHub](https://gist.github.com/randomcoder67/c0660e56337ab3d935810c310be31324)  
Copy this text:

```
		lcd0@01c0c000 {
			compatible = "allwinner,sunxi-lcd0";
			pinctrl-names = "active", "sleep";
			status = "okay";
			lcd_used = <0x01>;
			lcd_driver_name = "fog_fj035fhd05_v1";
			lcd_backlight = <0x32>;
			lcd_if = <0x00>;
			lcd_hv_if = <0x00>;
			lcd_x = <0x280>;
			lcd_y = <0x1e0>;
			lcd_width = <0x96>;
			lcd_height = <0x5e>;
			lcd_dclk_freq = <0x18>;
			lcd_pwm_used = <0x01>;
			lcd_pwm_ch = <0x00>;
			lcd_pwm_freq = <0xc350>;
			lcd_pwm_pol = <0x01>;
			lcd_pwm_max_limit = <0xff>;
			lcd_hbp = <0x2e>;
			lcd_ht = <0x302>;
			lcd_hspw = <0x14>;
			lcd_vbp = <0x0f>;
			lcd_vt = <0x208>;
			lcd_vspw = <0x04>;
			lcd_lvds_if = <0x00>;
			lcd_lvds_colordepth = <0x00>;
			lcd_lvds_mode = <0x00>;
			lcd_frm = <0x01>;
			lcd_hv_clk_phase = <0x02>;
			lcd_hv_sync_polarity = <0x03>;
			lcd_gamma_en = <0x00>;
			lcd_bright_curve_en = <0x00>;
			lcd_cmap_en = <0x00>;
			deu_mode = <0x00>;
			lcdgamma4iep = <0x16>;
			smart_color = <0x5a>;
			lcd_pin_power;
			lcd_power;
			lcd_bl_en;
			lcd_gpio_0 = <0x4b 0x08 0x09 0x01 0xffffffff 0xffffffff 0x01>;
			lcd_gpio_1 = <0x4b 0x08 0x0a 0x01 0xffffffff 0xffffffff 0x01>;
			lcd_gpio_2 = <0x4b 0x08 0x08 0x01 0xffffffff 0xffffffff 0x01>;
			lcd_gpio_3 = <0x4b 0x08 0x0e 0x01 0xffffffff 0xffffffff 0x01>;
			lcd_gpio_4 = <0x4b 0x08 0x0f 0x01 0xffffffff 0xffffffff 0x01>;
			pinctrl-0 = <0xa2>;
			pinctrl-1 = <0xa3>;
			linux,phandle = <0x164>;
			phandle = <0x164>;
		};
```
