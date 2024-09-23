# **Rhythm Generator**

## **Description**

This program offers two types of algorithms to generate, print and playback patterns of evenly spread beats (on sets) in a given number of steps (on sets + off sets). The default algorithm is the Euclidean Rhythm algorithm and the optional is Custom algorithm. Both algorithms produce similar results. In some cases the Custom algorithm produces less even but musically interesting patterns.  Rhythm Generator also offers an optional `fill` function that bridges the gap between beats that are separated by two or more off sets. The fill option produces results with more on set values, adding a type of *double time feel* to the new pattern while outlining the original. The results are notated with the capital letter `X` for onsets, lower case `o` for offsets and `x` for filled off sets.

## **Usage**

### Steps

Dial in the `Steps` box the appropriate integer value for the desired steps.

### Beats

Dial in the `Beats` box the appropriate integer value for the desired beats. Beats must be a smaller integer than steps.

### BPM

BPM must be an integer value between 1-500. The steps are interpreted as eight notes against a quarter note pulse. This means that each BPM pulse is equal to two steps.

### Double Time

If `Double Time` is checked the steps are interpreted as 16 notes. The threshold for bpm values in this setting is 250.

### Click

Click is in unison with the BPM unless `Double Time` is checked in which case the click is played in half time. In other words the steps are interpreted as 16 notes which is in aliment with the double time definition. Click can be turned on and off during playback.

### Play Offsets

The default setting for this option is on. While playing you can toggle the play offsets checkbox. If not enabled the offsets are interpreted as rests while if enabled the offsets are interpreted as closed hihats.

### Play and Stop 

When you hit play you see the type of Pattern ie `Pattern: Custom Filled Algorithm`. The pattern is repeated until stop is pressed and the iteration count is indicated on the `bar` label. Below that you can see the notation of the generated pattern with `X` indicating the beats, `o` the offsets and `x` the filled steps. The `X`'s trigger a snare rimshot sound, the `o`'s trigger a closed hi hat sound and the `x`'s trigger a snare side stick sound. After the playback is stopped you can restart it by pressing play again or first change any of the parameters and then play the new pattern.

### Choose Algorithm Type

By default the algorithm for pattern generation is Euclidean Rhythm. You can change to Custom by ticking the `Custom Algorithm` box. 

### Remove Symmetry

The `Remove Symmetry` option breaks the symmetry of a pattern by inverting the last part of the pattern. This produces more musically meaningful results. To see that in action dial 12 Steps 9 beats in your preferred Bpm. If a non symmetrical pattern is not achievable the check box will be disabled during play back and reenabled when stopped.

### Fill steps:  

Tick the fill steps box to add more movement to any pattern. If the fill steps function doesn't apply to the generated pattern the checkbox is automatically un checked.

### Play Fills

If there is a filled steps pattern generated `Play Fills` checkbox is automatically checked. If unchecked the filled steps are interpreted as rests creating some negative space in the pattern. If enabled the offsets are interpreted as softer values (a side stick sound). See [Play/Stop](#play-and-stop) and [Fill Steps](#fill-steps) for more info.

### Invert Left/Right:

Inverts the generated pattern one step left or right. After inverting while playing the inversion while be played starting from the next bar. You can see the direction and invertion value as possitive integers for right and negative intergers for left.

## **Compatibility**

- Go v1.23.1

## **Installation**

You can run the program in a go environment or build using the `build.sh` file provided to run as an executable.
In any case first follow these steps:

1. Clone repository https://github.com/vstefanopoulos/rhythm-generator.git
2. Make sure you have go 1.23.1 version or later installed on your machine 
3. Open terminal at repo folder

### **Instructions for Go environment**

Run the terminal command `$go run main.go`

### **Instructions for building executable**

1. `$chmod +x build.sh`
2. Run `$./build.sh`
3. Opt to delete or keep go source files
4. Run `$./main.bin` 


NOTE: When building or run in go environment the message `$ld: warning: ignoring duplicate libraries: '-lobjc'` might pop up in terminal. To my knowledge this doesn't affect the program's functionality
