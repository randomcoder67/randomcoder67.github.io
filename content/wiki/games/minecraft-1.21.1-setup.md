# Minecraft 1.21.1 Setup - Vanilla Plus

I wanted to get back into Minecraft, so I have attemped to put together a modpack/setup that fine tunes the default experience, as well as fixing some issues I have.  
Below is a list of all the mods, data packs and resource packs. 

The changes I made are partially inspired by [Whitelight's](https://www.youtube.com/@Whitelight) recent Minecraft critique, I highly recommend watching it: [Another Serious Critique of Minecraft - Whitelight](https://www.youtube.com/watch?v=_KqeLT-EOe0)

## "Rules of Play"

There are some parts of modern Minecraft that I think are detrimental to the experience if the player overuses them. For this reason I created a loose set of rules for myself to follow during a playthrough. 

1. Don't abuse villager mechanics
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

Install fabric modloader for 1.21.1  
Download linked mods and place them in `~/.minecraft/mods`  
Download linked data packs and add them when creating a new world in Minecraft itself  
Download linked resource packs and place them in `~/.minecraft/resourcepacks`, then activate them inside the game.  

## Launcher

The default Minecraft launcher is not great, it's very slow and it's difficult to use modern and old version at the same time.  
[MultiMC](https://multimc.org/) is a good alternative launcher. It is avalible for Linux ([Arch PKGBUILD](https://github.com/MultiMC/multimc-pkgbuild)), and has been around for eons.  

## Mods

### Performance

* [Better Beds](https://www.curseforge.com/minecraft/mc-mods/better-beds) - Fewer block entities
* [Dynamic FPS](https://www.curseforge.com/minecraft/mc-mods/dynamic-fps) - Background FPS limit
* [Enhanced Block Entities](https://www.curseforge.com/minecraft/mc-mods/enhanced-block-entities) - Less chest lag
* [Entity Culling](https://www.curseforge.com/minecraft/mc-mods/dynamic-fps) - Less mob lag
* [Ferrite Core](https://www.curseforge.com/minecraft/mc-mods/ferritecore-fabric) - Less RAM usage
* [ImmediatelyFast](https://www.curseforge.com/minecraft/mc-mods/immediatelyfast) - More FPS
* [Not Enough Crashes](https://www.curseforge.com/minecraft/mc-mods/not-enough-crashes) - Better crash handling
* [Lithium](https://www.curseforge.com/minecraft/mc-mods/lithium) - Tick optimisation
* [Sodium](https://www.curseforge.com/minecraft/mc-mods/sodium) - FPS optimisation
* [Concurrent Chink Management Engine](https://modrinth.com/mod/c2me-fabric) - Better chunk loading
* [Cull Less Leaves](https://www.curseforge.com/minecraft/mc-mods/cull-less-leaves) - Cull leaves but not ugly

### Graphics

* [Iris Shaders](https://www.curseforge.com/minecraft/mc-mods/irisshaders) - Alternative to Optifine for shaders
* [Continuity](https://www.curseforge.com/minecraft/mc-mods/continuity) - Connected textures
* [Falling Leaves](https://modrinth.com/mod/fallingleaves) - Adds leaves that fall from trees
* [Visuality](https://modrinth.com/mod/visuality) - A bunch of little effects
* [Better Clouds](https://modrinth.com/mod/better-clouds)
* [Distant Horizons](https://modrinth.com/mod/distanthorizons) - LOD mod, enables much longer render distance
* [Particle Rain](https://modrinth.com/mod/particle-rain/) - Pretty weather effects

### Audio

* [AmbientSounds 6](https://www.curseforge.com/minecraft/mc-mods/ambientsounds) - Ambience
* [Sound Physics Remastered](https://www.curseforge.com/minecraft/mc-mods/sound-physics-remastered) - Realistic sounds
* [Charmonium](https://modrinth.com/mod/charmonium) - AmbientSounds 6 alternative, only install one
* [ExtraSounds](https://modrinth.com/mod/extrasounds) - UI sounds
* [Auditory Continued](https://modrinth.com/mod/auditory-continued) - Better block sounds
* [Cool Rain](https://modrinth.com/mod/coolrain) - Realistic rain sounds based on blocks

### UI

* [Obscure Tooltips](https://www.curseforge.com/minecraft/mc-mods/obscure-tooltips) - Fancy tooltips
* [Loot Journal](https://www.curseforge.com/minecraft/mc-mods/loot-journal-fabric) - Shows items your pick up
* [Status Effect Bars](https://www.curseforge.com/minecraft/mc-mods/status-effect-bars) - Duration bars
* [Better F3](https://modrinth.com/mod/betterf3) - Nicer looking and customisable F3 menu
* [Better Statistics Screen](https://modrinth.com/mod/better-stats)
* [Leave My Bars Alone](https://www.curseforge.com/minecraft/mc-mods/leave-my-bars-alone) - Prevent horse GUI from removing defualt GUI
* [Detail Armor Bar](https://modrinth.com/mod/detail-armor-bar)

### Inventory

* [ItemSwapper](https://modrinth.com/plugin/itemswapper) - Easy swapping of related items

### Utility

* [Mod Menu](https://www.curseforge.com/minecraft/mc-mods/modmenu) - A configuration menu for mods
* [Debugify](https://www.curseforge.com/minecraft/mc-mods/debugify) - Lots of bug fixes
* [Don't Clear Chat History](https://www.curseforge.com/minecraft/mc-mods/dont-clear-chat-history) - Keep chat history between world loads
* [Replay Mod](https://modrinth.com/mod/replaymod)
* [Logical Zoom](https://www.curseforge.com/minecraft/mc-mods/logical-zoom) - Zoom
* [Debugify](https://modrinth.com/mod/debugify) - Fixes a bunch of bugs
* [Carpet](https://modrinth.com/mod/carpet) - Technical Minecraft improvements
* [Fabrishot](https://modrinth.com/mod/fabrishot) - Large screenshots
* [Screenshot Viewer](https://modrinth.com/mod/screenshot-viewer)
* [Disable Custom World's Advice](https://modrinth.com/mod/dcwa)
* [Resource Pack Overrides](https://modrinth.com/mod/resource-pack-overrides)

### Mobs

* [More Mob Variants](https://modrinth.com/mod/more-mob-variants)
* [Villager Names](https://modrinth.com/mod/villager-names-serilum)
* [VillagerConfig](https://modrinth.com/mod/villagerconfig) - Allows customising villager trades via datapacks - see datapack section below for custom trades datapack

### World Generation

* [Geophilic](https://www.curseforge.com/minecraft/mc-mods/geophilic) - Subtle additions to biomes
* [MES - Moog's End Structures](https://modrinth.com/mod/mes-moogs-end-structures/gallery) - Lots of structure for the End

### Gameplay

* [You've Goat to be Kidding Me](https://modrinth.com/mod/goated) - Main feature: Ram Block, redstone component that can break blocks
* [Immersive Aircraft](https://www.curseforge.com/minecraft/mc-mods/immersive-aircraft)
* [Exposure](https://modrinth.com/mod/exposure) - Realistic film camera mod - not currently updated for 1.21.1
* [Linkart Refabricated](https://github.com/Flatkat/Linkart-Refabricated?tab=readme-ov-file) - Minecart linking
* [Trample No More](https://modrinth.com/mod/trample-no-more) - Crops won't be trampled when wearing leather or Feather Falling boots
* [Advancement Frames](https://www.curseforge.com/minecraft/mc-mods/advancement-frames) - Frames to display your advancements
* [Map Atlases](https://www.curseforge.com/minecraft/mc-mods/map-atlases) - Not updated to 1.21.1
* [Sword Blocking Mechanics](https://www.curseforge.com/minecraft/mc-mods/sword-blocking-mechanics) - Brings back sword blocking, imo shields are overpowered

### Dependancies

* [Architectury API](https://www.curseforge.com/minecraft/mc-mods/architectury-api)
* [Bookshelf](https://www.curseforge.com/minecraft/mc-mods/bookshelf)
* [Cloth Config API](https://www.curseforge.com/minecraft/mc-mods/cloth-config)
* [CreativeCore](https://www.curseforge.com/minecraft/mc-mods/creativecore)
* [Collective](https://www.curseforge.com/minecraft/mc-mods/collective)
* [Fabric API](https://www.curseforge.com/minecraft/mc-mods/fabric-api)
* [Forge Config API Port](https://www.curseforge.com/minecraft/mc-mods/forge-config-api-port-fabric)
* [Moonlight Lib](https://www.curseforge.com/minecraft/mc-mods/selene)
* [Prickle Fabric](https://www.curseforge.com/minecraft/mc-mods/prickle)
* [Puzzles Lib](https://www.curseforge.com/minecraft/mc-mods/puzzles-lib)
* [TCDCommons API](https://modrinth.com/mod/tcdcommons)
* [YetAnotherConfigLib](https://modrinth.com/mod/yacl)

## Datapacks

### My/Modified Datapacks

* [My/Modified Datapacks](https://github.com/randomcoder67/Minecraft-Tweaks-1.21.1)

* Includes changes to world generation (smaller "temperature regions")
* Changes to villager trading (inspired by the 1.20.2 experimental villager rebalance)
* More difficult firework rocket recipe (uses Dragon's Breath)
* Craftable bundle

* Also includes modified [Survivor's Elegy](https://modrinth.com/datapack/survivors-elegy/) and [The Creeper's Code](https://modrinth.com/datapack/the-creepers-code/), with only a few features

### Vanilla Tweaks

* [Vanilla Tweaks Datapacks](https://vanillatweaks.net/picker/datapacks/)

#### Convenience

* More Effective Tools - Gives some blocks proper tools to use

#### Informative

* Real Time Clock - Adds a clock which shows how long you have played in a world for in real time
* Coordinates HUD - Useful addition to HUD, had coordinates, facing direction and time
* Durability Ping - Stops you accidentally breaking tools

## Resource Packs

* [Fancy: GUI Overhaul](https://www.curseforge.com/minecraft/texture-packs/fancy-gui-overhaul)
* [Cubic Sun & Moon](https://modrinth.com/resourcepack/cubic-sun-moon)
* [Invisible Item Frames](https://www.curseforge.com/minecraft/texture-packs/invisible-item-frames)
* [xali's Enchanted Books](https://www.curseforge.com/minecraft/texture-packs/xalis-enchanted-books)
* [xali's Potions](https://www.curseforge.com/minecraft/texture-packs/xalis-potions)
* [Crops 3D](https://www.curseforge.com/minecraft/texture-packs/crops-3d)
* [Enhanced Boss Bars](https://www.curseforge.com/minecraft/texture-packs/enhanced-boss-bars)
* [Spawn Egg Faces](https://www.curseforge.com/minecraft/texture-packs/spawn-eggs-faces/gallery)
* [Simple 3D Chain](https://modrinth.com/resourcepack/simple-3d-chain)
* [Cyber's 3D Amethyst](https://www.curseforge.com/minecraft/texture-packs/cybers-3d-amethyst)
* [Better Lanterns](https://modrinth.com/resourcepack/better-lanterns)

### Vanilla Tweaks 

* [Vanilla Tweaks Resource Packs](https://vanillatweaks.net/picker/resource-packs/)

#### Aesthetic

* Colorful Enchanting Table Particles
* Warm Glow
* Flint Tipped Arrows

#### Utility

* Different Stems

#### GUI

* Ping Color Indicator
* Rainbow Experience Bar
* Classic Menu Panorama

#### Retro

* Old Door Sound
* Rename "Pottery Sherd" to "Pottery Shard"
* Old Level Up Sounds

#### Fun

* ilmango Golden Apples

#### Fixes

* Item Stitching Fix
* Updated Observer Texture
* Updated Stat Icon Textures
* Updated Spectator Icon Textures
* Redstone Wire Fix
* Big Dripleaf Stem Fix
* Small Dripleaf Stem Fix
* Consistent Blank Decorated Pot
* Consistent Buckets
* Proper Break Particles
* Iron Bars Fix
* Hoe Fix

## Shaders

* [Complementary Shaders](https://www.curseforge.com/minecraft/shaders/complementary-reimagined) - Not used when playing normally

## Other

* `doInsomnia` gamerule is set to false, as Phantoms attacking the player when they don't sleep only *encourages* overuse of beds. Also Phantoms just aren't fun
* `doFireTick` gamerule is set to false, as fire destroying your wood builds is just annoying

## Extra

* https://unmined.net/ - Mapping tool
* https://modrinth.com/mod/oritech - Tech mod, looks fun
