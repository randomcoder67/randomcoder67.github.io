# Old Inventorio Changes

As part of my Minecraft 1.20.1 setup, I originally had a slightly modified version of Inventorio. I had a GitHub fork but then deleted it once I decided to remove it from the setup.

I will quickly list the changes I had made here for archival purposes.

## Code Changes

I sligtly modified it so that there are always 2 extra inventory slots (removing the requirement for Deep Pockets enchamtment)

This was done by replacing all lines identical to this:

``` kt
EnchantmentHelper.getEquipmentLevel(DeepPocketsEnchantment, player).coerceIn(0, 3)
```

With this:

``` kt
EnchantmentHelper.getEquipmentLevel(DeepPocketsEnchantment, player).coerceIn(2, 2)
```

In the `PlayerInventoryExtension.kt` files. This file is present numerous times because of the multi-Minecraft-version structure of the mods code, so make sure to change it in every file.

This change makes the mod think you always have Deep Pockets II, causing 2 extra inventory slots, and an extra row of offhand slots, to be present at all times. Of course if you wanted 1 or 3 extra inventory slots, you could use `coerceIn(1, 1)`, or `coerceIn(3, 3)` instead.

## Config

I removed the auto-tool-switching feature in the config as I find it confusing, this can be done in game via a mod such as ModMenu

 
