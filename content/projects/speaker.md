# Speaker Made from Old Laptop Speakers

This project was inspired by [DIY Perks video](https://www.youtube.com/watch?v=WLP_L7Mgz6M) on uses for old laptops. I decided to try and build a speaker for my University room, using some old laptop speakers and a Raspberry Pi

## Components

### Speakers

Laptop speakers will be loud enough for my use case, as I am in a University room, so I don't want it to make too much noise and annoy the neighbours.  
I bought a broken Toshiba Satellite Pro L300D for £20, as review said it had decent speakers.

### Raspberry Pi

I decided to use a Raspberry Pi, instead of just the Bluetooth reciever DIY Perks uses, for two main reasons:
1. Bluetooth is annoying, and bluetooth recieveres have a tendancy to make really loud "DEVICE CONNECTED" noises, which I find annoying
2. If I use a Raspberry Pi, the speaker can be it's own distinct music player, and doesn't rely on other devices being connected. It would be nice to, for example, come back into my room and have music playing, even if I took both my phone and laptop with me.

Our University computer science department gives each student some storage space on an SSH server, so I can use that to transfer music back and forth to the Raspberry Pi without physically connecting to it.  
It can also be used to control the music playback, although a simpler solution to this would be nice (discussed later).

Originally I purchased a Raspberry Pi Zero 2 W, and planned to use the GPIO headphone jack functionality to get audio out to the amplifier board. I realised after recieving it however that this audio needs filtering. This would require a bunch more componentes, so I decided to just return it and buy a Pi 3 A+, as it was only about £5 more expensive, and has a headphone jack.

### Case

For the case I am using PVC Foam Board, as it is easy to glue together, and pretty cheap.  
I got [5 Sheets from Amazon](https://www.amazon.co.uk/dp/B08M3FYVYW)

### Other Components

#### Amplifier

This cheap amplifier from Amazon was used: [Ticfox Digital Amplifier Board 5Wx2](https://www.amazon.co.uk/dp/B09YYGBZ59)  
5 amplifier are actually included, I didn't realise that until they arrived, so I guess I have a lot of spares now.

#### 

## Build

## Software

I wrote a Bash script to control the music playback.

## Final Result
