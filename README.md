# **Rhythm Generator**

## **Description**

This program offers two types of algorithms to generate, print and playback patterns of evenly spread beats (on sets) in a given number of steps (on sets + off sets). The default algorithm is the Euclidean Rhythm algorithm and the optional is Custom algorithm. Both algorithms produce similar results. In some cases the Custom algorithm produces less even but musically interesting patterns. After the pattern is generated in the case where symmetry occurs the program breaks the symmetry by inverting the last part of the pattern. This produces more musically meaningfull results. To see that in action dial 10 Steps 4 beats in your prefered Bpm. Rhythm Generator also offers an optional `fill` function that bridges the gap between beats that are separated by two or more off sets. The fill option produces results with more on set values, adding a type of *double time feel* to the new pattern while outlining the original. The results are notated with the capital letter `X` for onsets, lower case `o` for offsets and `x` for filled off sets.

## **Usage**

### Steps

Dial in the `Steps` box the appropriate integer value for the desired steps

### Beats

Dial in the `Beats` box the appropriate integer value for the desired beats. Beats must be a smaller integer than steps

### BPM

BPM must be an integer value between 1-300. The steps are interpreted as eight notes against a quarter note pulse. This means that each BPM pulse is equal to two steps.

### Play/Stop

When you hit play you see the type of Pattern ie `Pattern: Custom Filled Algorithm`. The pattern is repeated until stop is pressed and the iteration count is indicated on the `bar` label. Below that you can see the notation of the generated pattern with `X` indicating the beats, `o` the offsets and `x` the filled steps. The `X`'s trigger a snare rimshot sound, the `o`'s trigger a closed hi hat sound and the `x`'s trigger a snare side stick sound. After the playback is stopped you can restart it by pressing play again or first change any of the parameters and then play the new pattern.

### Invert Left/Right:

Inverts the generated pattern one step left or right. After inverting you can audition the inverted pattern using play however if you change any other parameters (steps, beats, bpm, fill or Custom Algorithm) the pattern will be reset and any inversions need to be reapplied

### Choose Algorithm Type

By default the algorithm for pattern generation is Euclidean Rhythm. You can change to Custom by ticking the `Use Custom Algorithm` box. For each algorithm the Asymetrical function is applied.

### Fill steps:  

Tick the fill steps box to fill an sparse beats populated pattern or to add more movement to any pattern. In most cases the number of beats are more than the selected ones.

## **Compatibility**

- macOS 10.15 Catalina or later
- Go v1.22.2

## **Installation**

You can run the program in a go enviroment or build using the `build.sh` file provided to run as an executable.
In any case first follow these steps:

1. Clone repository https://github.com/vstefanopoulos/rhythm-generator.git
2. Make sure you have go 1.22.2 version or later installed on your machine 
3. Open terminal at repo folder

### **Instructions for Go enviroment**

Run the terminal command `$go run main.go`

### **Instructions for building executable**

1. Run `$chmod +x build.sh`
2. Run `$./build.sh`
3. Run `$./main.bin`. 

NOTE: When building or run in go enviroment the message `$ld: warning: ignoring duplicate libraries: '-lobjc'` might pop up in terminal. To my knowledge this doesn't affect the program's functionality