# Minecraft 1.20.1 Setup - Vanilla Plus

I wanted to get back into Minecraft, so I have attemped to put together a modpack/setup that fine tunes the default experience, as well as fixing some issues I have.  
Below is a list of all the mods, data packs and resource packs. 

The changes I made are partially inspired by Whitelight's recent Minecraft critique, I highly recommend watching it: [Another Serious Critique of Minecraft - Whitelight](https://www.youtube.com/watch?v=_KqeLT-EOe0)

## "Rules of Play"

There are some parts of modern Minecraft that I think are detrimental to the experience if the player overuses them. For this reason I created a loose set of rules for myself to follow during a playthrough. 

1. Don't abuse villager mechanics
	* Village spawn rates have been reduced, so this should happen somewhat naturally
	* Try to avoid spending hours settings up perfect villager trading halls, as this bypasses a lot of the challenge in the game
2. Don't abuse beds
	* Not allowed to use beds on the first few nights
	* Try to work/do things outside at night, don't just sleep at 18:32 every night
	* Phantoms are disabled using a gamerule to encourage this
3. Don't spend all the time making automated farms
	* Try and get resources the indended way, at least for the early and mid game
	* If you want to make farms, make the less overpowered types (Guardian farm instead of Enderman farm etc)
4. Don't overuse/rush to Elytra
	* Try and use some of the other movement mechanics in the game, at least in the early and mid game
	* For example Horses and Minecarts (especially with the Linkart mod) are still very useful modes of transport
	* Consider only using Elytra as it was originally indended, as a glider (i.e. not using fireworks)

## Install Instructions

Install fabric modloader for 1.20.1  
Download linked mods and place them in `~/.minecraft/mods`  
Download linked data packs and add them when creating a new world in Minecraft itself  
Download linked resource packs and place them in `~/.minecraft/resourcepacks`, then activate them inside the game (order given in resource pack section)  

## Launcher

The default Minecraft launcher is not great, it's very slow and it's difficult to use modern and old version at the same time.  
[MultiMC](https://multimc.org/) is a good alternative launcher. It is avalible for Linux ([Arch PKGBUILD](https://github.com/MultiMC/multimc-pkgbuild)), and has been around for eons.  
There is also [PolyMC](https://polymc.org/) and [Prism Launcher](https://prismlauncher.org/), which seem to be basically identical but apparently PrismMC was forked from PolyMC because of some drama. There is a bunch of boring political drama around the two, and MultiMC seems to do everything you would need, so just stick to that.

## Mods

### Graphics/Audio

* [Ambient Sounds](https://www.curseforge.com/minecraft/mc-mods/ambientsounds) - Self explanatory, adds nice ambient sounds to the game
* [Better Clouds](https://modrinth.com/mod/better-clouds) - More realistic but still vanilla looking clouds
* [Cave Dust](https://modrinth.com/mod/cave-dust) - Basalt Delta like particle to all caves
* [Continuity](https://modrinth.com/mod/continuity) - Optifine's connected textures support for Fabric/Sodium
* [Explosive Enhancement](https://modrinth.com/mod/explosive-enhancement) - Pretty explosions
* [Falling Leaves](https://modrinth.com/mod/fallingleaves) - Falling leaf particles for every tree type
* [Not Enough Animations](https://www.curseforge.com/minecraft/mc-mods/not-enough-animations) - Third person equivalent of first person animations
* [Visuality](https://modrinth.com/mod/visuality) - Adds lots more small effects to the game

(Screenshots)

### Inventory

* [Clean Tooltips](https://modrinth.com/mod/clean-tooltips) - Easier to read item tooltips
* [Detail Armor Bar](https://www.curseforge.com/minecraft/mc-mods/detail-armor-bar) - Colour of armour bar icons varies depending on the type of armour you are wearing
* [GUI Clock](https://www.curseforge.com/minecraft/mc-mods/gui-clock) - Vanilla friendly way to get the accurate in game time - feels less cheaty than using a datapack to always display it
* [GUI Compass](https://www.curseforge.com/minecraft/mc-mods/gui-compass) - Same idea as GUI Clock but for coordinaties. I use the Better F3 to remove coordinates from the F3 menu so this is the only way to see them
* [Leave My Bars Alone](https://www.curseforge.com/minecraft/mc-mods/leave-my-bars-alone) - Stops important hud bars from being hidden when riding a mob
* [Pick Up Notifier](https://modrinth.com/mod/pick-up-notifier) - Pick up notifer for item picked up from the ground
* [Traveller's Titles](https://www.curseforge.com/minecraft/mc-mods/travelers-titles-fabric) - Shows RPG like titles when entering a biome (reminds me of [Cube World](https://www.youtube.com/watch?v=kkzafkGcswM))

### Mobs

* [More Mob Variants](https://modrinth.com/mod/more-mob-variants) - Adds variants of mobs which haven't been updated in a long time
* [VillagerConfig](https://modrinth.com/mod/villagerconfig) - Allows custimising villager trades via datapacks - see datapack section below for custom trades datapack

### Additions

* [Better Trims](https://www.curseforge.com/minecraft/mc-mods/better-trims) - Gives trims more of an ingame purpose by adding effects to them
* [Block Runner](https://www.curseforge.com/minecraft/mc-mods/block-runner-forge) - The player moves faster on paths, gives them more of a point
* [Exposure](https://modrinth.com/mod/exposure) - Film camera mod, allows you to take photos in game and place them in item frames
* [Linkart](https://modrinth.com/mod/linkart) - Allows Minecarts to be connected with chains
* [Antique Atlas 4](https://modrinth.com/mod/antique-atlas-4) - Atlas mod in the style of an old timey explorer map. Not too overpowered so should fit well into vanilla

### Utility

* [Better Third Person](https://modrinth.com/mod/better-third-person) - Simple mod which makes the third person camera much more usable
* [Fabrishot](https://www.curseforge.com/minecraft/mc-mods/fabrishot) - Allows taking larger sized screenshots
* [Screenshot Viewer](https://www.curseforge.com/minecraft/mc-mods/screenshot-viewer) - Ingame screenshot viewer so you don't have to open the screenshots folder all the time
* [Mod Menu](https://modrinth.com/mod/modmenu) - Simple mod configuration tool for Fabric mods
* [Better F3](https://www.curseforge.com/minecraft/mc-mods/betterf3) - More readable debug screen, also allows customising it. My included config disables the coordinates, requiring the player to craft a compass to see their coordinates
* [Better Statistics Screen](https://www.curseforge.com/minecraft/mc-mods/better-stats) - Nicer statistics screen
* [Disable Custom Worlds Advice](https://www.curseforge.com/minecraft/mc-mods/fabric-disable-custom-worlds-advice)- Gets rid of the infuriating "this world has experimental features" confirmation when loading a world with custom settings
* [Resource Pack Overrides](https://www.curseforge.com/minecraft/mc-mods/resource-pack-overrides) - Allows more control over which resource packs

### World Generation

#### Terrain/Biomes

* [Geophillic](https://modrinth.com/datapack/geophilic/) - Adds small extra details to biomes

#### Structures

* [YUNG's Better Mineshafts](https://www.curseforge.com/minecraft/mc-mods/yungs-better-mineshafts-fabric) - Makes mineshafts much more interesting
* [YUNG's Better Strongholds](https://www.curseforge.com/minecraft/mc-mods/yungs-better-strongholds-fabric) - Makes strongholds much more interesting

### Performance

* [Sodium](https://modrinth.com/mod/sodium) - FPS optimisation mod - far better than Optifine (in terms of performance, compatibility, update speed and the fact it's open source!)
* [Lithium](https://modrinth.com/mod/lithium) - TPS (ticks per second) optimisation mod
* [Entity Culling](https://www.curseforge.com/minecraft/mc-mods/entityculling) - Doesn't render entities which the player cannot see. Leads to drastic FPS improvements in areas like villager trading halls or large farms
* [FerriteCore](https://modrinth.com/mod/ferrite-core) - Memory optimisation mod
* [Enchanced Block Entities](https://www.curseforge.com/minecraft/mc-mods/enhanced-block-entities) - Makes some block entities use block models, leading large FPS gains when a lot of chests are being rendered for example
* [Debugify](https://www.curseforge.com/minecraft/mc-mods/debugify) - Fixes a bunch of bugs, including some chunk loading ones which can lead to performance improvements

### Dependancies

* [Architectury](https://modrinth.com/mod/architectury-api)
* [Bookshelf Fabric](https://www.curseforge.com/minecraft/mc-mods/bookshelf)
* [CITResewn](https://modrinth.com/mod/cit-resewn)
* [Cloth-Config](https://modrinth.com/mod/cloth-config)
* [Collective](https://modrinth.com/mod/collective)
* [Configured - Fabric](https://www.curseforge.com/minecraft/mc-mods/configured-fabric)
* [Creative Core Fabric](https://www.curseforge.com/minecraft/mc-mods/creativecore)
* [Fabric API](https://modrinth.com/mod/fabric-api)
* [Forge Config API Port](https://www.curseforge.com/minecraft/mc-mods/forge-config-api-port-fabric)
* [Geckolib Fabric](https://www.curseforge.com/minecraft/mc-mods/geckolib)
* [Indium](https://modrinth.com/mod/indium)
* [Moonlight](https://modrinth.com/mod/moonlight)
* [PuzzlesLib](https://modrinth.com/mod/puzzles-lib)
* [Resourceful Config Fabric](https://modrinth.com/mod/resourceful-config)
* [Resourceful Lib Fabric](https://modrinth.com/mod/resourceful-lib)
* [Yet Another Config Lib](https://modrinth.com/mod/yacl)
* [Yungs API](https://www.curseforge.com/minecraft/mc-mods/yungs-api-fabric)
* [Entity Model Features](https://modrinth.com/mod/entity-model-features)
* [Entity Model Textures](https://modrinth.com/mod/entitytexturefeatures)

## Datapacks

* [Vanilla Tweaks Datapack](https://vanillatweaks.net/picker/datapacks/):
	* Real Time Clock - Adds a clock which shows how long you have played in a world for in real time
* [Vanilla Tweaks Crafting Tweaks](https://vanillatweaks.net/picker/crafting-tweaks/)
	* Craftable Bundles - Makes the bundle item craftable with rabbit hide as intended
* [Villager Trade Rebalance Port](https://github.com/randomcoder67/Minecraft-1.20.1-Villager-Trade-Rebalance-Port) - Everyone hated it but the [1.20.2 villager trade rebalance](https://minecraft.wiki/w/Villager_Trade_Rebalance) experimental feature was a massive improvement, and was something we really needed. This datapack, in combination with the VillagerConfig mod, ports the changes to 1.20.1, as well as adding similar changes to the Toolsmith and Weaponsmith villagers.
* [Custom Worldgen - Older Style](https://github.com/randomcoder67/Minecraft-1.20.1-Custom-World-Gen---Older-Minecraft-Style) - This is a simple datapack I made to tweaks some aspects of the world generation. At the moment it has two features:
	* In modern Minecraft, villages are far too common, taking all the excitement out of finding one and making abusing their mechanics far too easy. This datapack drastically reduces the number of villages which spawn, encouraging either further exploration or building your own village with cursed zombie villagers!
	* This datapack also reduces the level of temperature clustering in biomes. Temperature based biome clustering was introduced in 1.7 and in my opinion is makes the world far more tedious to explore for the sake of "realism"
* [Survivor's Elegy Modified](https://github.com/randomcoder67/Survivors-Elegy-rc67-Fork) - Survivor's Elegy is a datapack intendedto improve vanilla Minecraft. This is my fork only including the features I want to use. I highly recommend checking out the original and all of it's features.
	* The features included in my fork are listed on the linked GitHub page
* [The Creepers Code Modified](https://github.com/randomcoder67/the-creepers-code-fork-rc67) - The Creepers Code is similar to Survivor's Elegy in style. I forked it and only included the features I wanted to use. As with Survivor's Elegy, I highly recommend checking out the original to see all of the features.
	* Similarly to Survivor's Elegy, features I included are listed on the linked GitHub page

## Resource Packs

* [Enhanced Boss Bars](https://www.curseforge.com/minecraft/texture-packs/enhanced-boss-bars) - Makes the boss bars look less generic
* [Better Lanterns](https://www.curseforge.com/minecraft/texture-packs/better-lanterns) - 3D Lanterns and Chains
* [Better Skeletons](https://www.curseforge.com/minecraft/texture-packs/better-skeletons) - Variety in skeleton textures
* [Crops 3D](https://www.curseforge.com/minecraft/texture-packs/crops-3d) - 3D Crops
* [Cubic Sun & Moon](https://modrinth.com/resourcepack/cubic-sun-moon) - Makes the Sun and Moon cubes
* [Enchanced Item Frame](https://www.curseforge.com/minecraft/texture-packs/enhanced-item-frame) - Fancy item frames
* [Invisible Item Frames](https://www.curseforge.com/minecraft/texture-packs/invisible-item-frames) - Makes item frames themselves invisible, only showing the item inside
* [Fancy GUI Overhaul](https://www.curseforge.com/minecraft/texture-packs/fancy-gui-overhaul) - Pretty GUI for interactable blocks
* [xali's Enchanted Books](https://www.curseforge.com/minecraft/texture-packs/xalis-enchanted-books) - Makes enchanted books have different appearances depending on the enchantment
* [xali's Potions](https://www.curseforge.com/minecraft/texture-packs/xalis-potions) - Gives different potions unique textures
* [Skeleton Physics](https://www.curseforge.com/minecraft/texture-packs/skeleton-physics) - Skeletons fall apart into bones when killed
* [Soft Weather](https://www.curseforge.com/minecraft/texture-packs/soft-weather) - Makes the weather effects less in-your-face

### Vanilla Tweaks 

* [Vanilla Tweaks Resource Packs](https://vanillatweaks.net/picker/resource-packs/)

#### Aesthetic

* Red Iron Golem Flowers
* Colorful Enchanting Table Particles
* Unique Dyes
* Flint Tipped Arrows
* Moss Carpet Overhang

#### Utility

* Diminishing Tools
* Different Stems
* Clear Banner Patterns

#### 3D

* 3D Chains

#### HUD

* Ping Color Indicator
* Rainbow Experience Bar

#### Menu Panoramas

* Classic

#### Retro

* Classic Minecraft Logo (no "Java Edition" subtitle)
* Rename "Pottery Sherd" to "Pottery Shard"

#### Fixes

* Item Stitching Fix
* Updated Observer Texture
* Updated Toast Textures
* Updated Spectator Icon Textures
* Big Dripleaf Stem Fix
* Small Dripleaf Stem Fix
* Consistent Buckets
* Consistent Tadpole Bucket
* Proper Break Particles
* Soul Soil Soul Campfire
* Iron Bars Fix
* Hoe Fix

## Other

* `doInsomnia` gamerule is set to false, as Phantoms attacking the player when they don't sleep only *encourages* overuse of beds. Also Phantoms just aren't fun
* `doFireTick` gamerule is set to false, as fire destroying your wood builds is just annoying

## Notes

Originally I had a modified version of the mod [Inventorio](https://modrinth.com/mod/inventorio) on this list, however I removed it. For more details see [Inventorio Fork](inventorioFork.md)
