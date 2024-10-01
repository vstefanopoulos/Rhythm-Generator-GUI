# **Rhythm Generator**

## **Description**

This program offers two types of algorithms to generate, print and playback patterns of evenly spread beats (onsets) in a given number of steps (onsets + offsets). The default algorithm is the Euclidean Rhythm algorithm and the optional is Custom algorithm. Both algorithms produce similar results. In some cases the Custom algorithm produces less even but musically interesting patterns. Rhythm Generator also offers an optional `Fill Steps` function that bridges the gap between beats that are separated by two or more off sets. The fill option produces results with more on set values, adding a type of *double time feel* to the new pattern while outlining the original. The results are notated with the capital letter `X` for onsets, lower case `o` for offsets and `x` for filled off sets.

## **Usage**

### Beats

Dial in the `Beats` box the appropriate integer value for the desired beats. Beats must be a smaller integer than steps. While playing you can change the beats by dialling in a new value and hitting `enter`. The new pattern will be auditioned on the next bar.

### Steps

Dial in the `Steps` box the appropriate integer value for the desired steps. While playing you can change the steps by dialling in a new value and hitting `enter`. The new pattern will be auditioned on the next bar.

### BPM

BPM must be an integer value between 1-500. The steps are interpreted as eight notes against a quarter note pulse. This means that each BPM pulse is equal to two steps. BPM can be changed while playing by dialling the new bpm value and hitting `enter`. The new tempo will take effect in the next bar.

### Double Time

If `Double Time` is checked the steps are interpreted as 16 notes. The threshold for bpm values in this setting is 250. Double time can be toggled while playing. The new tempo will take effect in the next bar.

### Click

Click is in unison with the BPM unless `Double Time` is checked in which case the click is played in half time. In other words the steps are interpreted as 16 notes which is in aliment with the double time definition. Click can be turned on and off during playback.

### Accent Downbeat

If enabled a lower pitched sound is triggered on the first step of each pattern.

### Mute Offsets

While playing you can toggle the mute offsets checkbox. If enabled the offsets are interpreted as rests.

### Play and Stop 

When you hit play audition the generated pattern is repeated until stop is pressed and the iteration count is indicated on the `bar` label. Below that you can see the notation of the generated pattern with `X` indicating the beats, `o` the offsets and `x` the filled steps. The `X`'s trigger a snare rimshot sound, the `o`'s trigger a closed hi hat sound and the `x`'s trigger a snare side stick sound. After the playback is stopped you can restart it by pressing play again or change any of the parameters and then play the new pattern.

### Choose Algorithm Type

By default the algorithm for pattern generation is Euclidean Rhythm. You can change to Custom by ticking the `Custom Algorithm` box. 

### Remove Symmetry

The `Remove Symmetry` option breaks the symmetry of a pattern (if there is any) by inverting the last part of the pattern. This produces more musically meaningful results. To see that in action dial `Steps: 15` and `Beats: 6` in your preferred Bpm and compare that result. The Rs box remains ticked in the UI if symmetry not removed and `✓Ok` appears next to it if symmetry was removed. If not applicable but remove symmetry checked the algorithm will be applied again to the next generated pattern

### Fill Steps:  

Toggle the `Fill Steps` box to add more movement to a pattern. If the fill steps function doesn't apply to the generated pattern the button remains checked and fill steps algorithm will be applied to the next generated pattern if applicable. If after checked and if filled pattern is possible `✓Ok` appears next to it. 

### Mute Fills

If checked the filled steps are interpreted as rests creating some negative space in the pattern. If enabled the offsets are interpreted as softer values (a side stick sound). See [Play/Stop](#play-and-stop) and [Fill Steps](#fill-steps) for more info.

### Invert Left/Right:

Inverts the generated pattern one step left or right. After inverting while playing the new inverted pattern will be played on the next bar. You can see the direction and inversion value as positive integers for right and negative integers for left. 

## **Compatibility**

- Go v1.23.1

## **Installation**

You can run the program in a go environment or build using the `build.sh` file provided to run as an executable.
In any case first follow these steps:

1. Clone repository https://github.com/vstefanopoulos/rhythm-generator.git
2. Make sure you have go 1.23.1 version or later installed on your machine 
3. Open terminal or bash at repo folder

### **Instructions for Go environment**

Run the terminal command `$go run main.go`

### **Instructions for building executable**

1. `$chmod +x build.sh`
2. Run `$./build.sh`
3. Opt to delete or keep go source files
4. Run `$./main.bin` 


NOTE: When building or run in go environment the message `$ld: warning: ignoring duplicate libraries: '-lobjc'` might pop up in terminal. To my knowledge this doesn't affect the program's functionality
